// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/mmio"
	"embedded/rtos"
	"math/bits"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/mem/dtcm"
	"github.com/embeddedgo/imxrt/hal/irq"

	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/usb"
	"github.com/embeddedgo/imxrt/p/usbphy"
)

const leNum = len(usb.Periph{}.ENDPTCTRL)

type noteNext struct {
	note rtos.Note
	next uintptr // *noteNext
}

func (nn *noteNext) uintptr() uintptr {
	return uintptr(unsafe.Pointer(nn))
}

const maxCtrlData = 256 // BUG: may be too small for the config descriptors

// Control transfer data structures.
type ctds struct {
	dtd  DTD // control data stage TD, 32 B, requires 32 alignment
	std  DTD // control status stage TD, 32 B, requires 32 alignment
	data [maxCtrlData]byte
}

// We use DTCM for buffers that require specific alignment and to avoid cache
// maintenance operations for transactions performed in the IRQ handler mode.
type dtcmem struct {
	qhs [leNum * 2]dQH // queue heads, 1024 B, requires 4096 alignment
	isr ctds           // used by ISR in the control data stage
	thr ctds           // used by handleControRequests goroutine
}

var dtcmCache [2]*dtcmem // cache the allocated DTCM for both USB controllers

// A Device represents an USB Device Controler Driver (DCD).
type Device struct {
	u      *usb.Periph
	phy    *usbphy.Periph
	des    map[uint32]string
	dtcm   *dtcmem
	atdtwm sync.Mutex
	config atomic.Uint32
	cr     ControlRequest // for control requests in ISR

	// For threads waiting for the configured state.
	cwl   atomic.Uintptr // *wait
	cwlmu sync.Mutex

	// For control request in the thread mode.
	crno rtos.Note
	crst uint32
	crsa [leNum][2]uint32
	crhm map[uint32]func(cr *ControlRequest) int
}

/*
func (d *Device) Print(he uint8) {
	qh := &d.dtcm.qhs[he]
	mmio.MB()
	fmt.Printf(
		"ectrl=%#x qh[%d]: mult=%d zlt=%d maxpkt=%4d ios=%d curr=%#x next=%#x\n",
		d.u.ENDPTCTRL[he/2].Load(), he, qh.config>>30&3, qh.config>>29&1, qh.config>>16&0x3ff, qh.config>>15&1, qh.current, qh.next,
	)
}
*/

// NewDevice returns a new device controler driver for USB controller 1 or 2.
func NewDevice(controller int) *Device {
	d := new(Device)
	switch controller {
	case 1:
		d.u = usb.USB1()
		d.phy = usbphy.USBPHY1()
	case 2:
		d.u = usb.USB2()
		d.phy = usbphy.USBPHY2()
	default:
		return nil
	}
	controller--
	m := dtcmCache[controller]
	if m == nil {
		m = dtcm.New[dtcmem](4096)
		m.isr.dtd.next = dtdEnd
		m.isr.std.next = dtdEnd
		m.thr.dtd.next = dtdEnd
		m.thr.std.next = dtdEnd
		dtcmCache[controller] = m
	}
	for he := range m.qhs {
		qh := &m.qhs[he]
		qh.head = dtdEnd
		qh.next = dtdEnd
	}
	d.dtcm = m
	d.cr.Data = m.isr.data[:] // cannot be set in ISR because of write barriers
	d.crhm = make(map[uint32]func(r *ControlRequest) int)
	return d
}

// Controller returnt the controller number used this device driver.
func (d *Device) Controller() int {
	switch d.u {
	case usb.USB1():
		return 1
	case usb.USB2():
		return 2
	}
	return -1
}

// Init initializes the USB device controler and the driver itself.
func (d *Device) Init(intPrio int, descriptors map[uint32]string, forceFullSpeed bool) {
	// Ungate all necessary clocks.
	ccm.CCM().CCGR6.SetBits(ccm.CG6_0 | ccm.CG6_11) // usboh3 | anadig (CCMA)
	ccm_analog.CCM_ANALOG().PLL_USB1_SET.Store(ccm_analog.PLL_USB_EN_USB_CLKS)

	u, phy := d.u, d.phy

	// Reset
	phy.CTRL_SET.Store(usbphy.SFTRST)
	u.USBCMD.SetBits(usb.RST)
	for u.USBCMD.LoadBits(usb.RST) != 0 {
	}
	phy.CTRL_CLR.Store(usbphy.SFTRST | usbphy.CLKGATE)

	// Enable power to PHY and select device mode.
	phy.PWD.Store(0)
	if forceFullSpeed {
		u.PORTSC1.SetBits(usb.PFSC)
	}
	u.USBMODE.Store(usb.CM_2 | usb.SLOM) // device mode, setup lockout disabled

	// Setup QHs for EP0.
	d.dtcm.qhs[0].setConf(64, dqhIOS) // Rx (host OUT)
	d.dtcm.qhs[1].setConf(64, 0)      // Tx (host IN)
	mmio.MB()

	u.ASYNC_ENDPTLISTADDR.Store(uint32(uintptr(unsafe.Pointer(&d.dtcm.qhs[0]))))

	d.des = descriptors

	// Enable interrupts
	ui := irq.USB_OTG1
	if u == usb.USB2() {
		ui = irq.USB_OTG2
	}
	ui.Enable(intPrio, 0)
	u.USBINTR.Store(usb.UE | usb.PCE | usb.URE | usb.SLE)
	// UEE not enabled. Transfer errors handled by goroutines waiting for IOC.

	go handleControRequests(d)
}

// field serves two functions. In case of OUT direction it contains the data
// received from the host, otherwise (IN direction) it serves as the output
// buffer, for sending data to the host.
type ControlRequest struct {
	LE      int8   // logical endpoint number
	Request uint16 // bRequest<<8 | bmRequestType
	Value   uint16 // wValue
	Index   uint16 // wIndex
	Data    []byte // len(Data) = wLength
}

// Handle registers the handler for the control requests adressed to the logical
// endpoint number le. All handlers must be registered before enabling the
// device.
func (d *Device) Handle(le int8, request uint16, handler func(cr *ControlRequest) int) {
	key := uint32(le)<<16 | uint32(request)
	if handler != nil {
		d.crhm[key] = handler
	} else {
		delete(d.crhm, key)
	}
}

// Enable enables the device controller.
func (d *Device) Enable() {
	time.Sleep(10 * time.Millisecond) // ensure a long enough disconnect state
	d.u.USBCMD.SetBits(usb.RS)
}

// Disable disables the device controller.
func (d *Device) Disable() {
	d.config.Store(0)
	d.u.USBCMD.ClearBits(usb.RS)
}

func handleControRequests(d *Device) {
	var cr ControlRequest
	for {
		d.crno.Sleep(-1)
		d.crno.Clear()
		crst := uint16(atomic.SwapUint32(&d.crst, 0))
		for cr.LE = 0; crst != 0; cr.LE, crst = cr.LE+1, crst>>1 {
			n := bits.TrailingZeros16(crst)
			cr.LE, crst = cr.LE+int8(n), crst>>uint(n) // skip the zero bits
			crsa := d.crsa[cr.LE]
			key := uint32(cr.LE)<<16 | crsa[0]&0xffff
			handler := d.crhm[key]
			n = parseSetup(&cr, crsa)
			cr.Data = d.dtcm.thr.data[:n]
			if handler == nil {
				badControlRequest(d, &cr)
				continue
			}
			execContorHandler(d, &d.dtcm.thr, &cr, handler)
		}
	}
}

// Config returns the configuration number selected during the USB enumeration
// process or zero if the device is not in the configured state.
func (d *Device) Config() int {
	return int(d.config.Load())
}

// WaitConfig waits for the selection of the cn configuration number during the
// USB enumeration process. Use cn=0 to wait for the configured state (any
// configuration number).
func (d *Device) WaitConfig(cn int) {
	for {
		cnf := d.config.Load()
		if cnf != 0 && (cnf == uint32(cn) || cn == 0) {
			return
		}
		d.cwlmu.Lock()
		var (
			cwl uintptr
			cw  noteNext
		)
		for {
			cwl := d.cwl.Load()
			cw.next = cwl
			if d.cwl.CompareAndSwap(cwl, cw.uintptr()) {
				break
			}
		}
		cnf = d.config.Load()
		if cnf != 0 && !d.cwl.CompareAndSwap(cw.uintptr(), cwl) {
			// ISR removed cw, must keep reference to cw until recieving a note
			cnf = 0
		}
		d.cwlmu.Unlock()
		if cnf == 0 {
			cw.note.Sleep(-1)
		}
	}
}

// Endpoint direction.
const (
	OUT uint8 = 0 // output endpoint (Rx for device, Tx for host)
	IN  uint8 = 1 // input endpoint (Tx for device, Rx for host)
)

// HE returns the hardware endpiont number for the given logical endpoint and
// direction.
func HE(le int8, dir uint8) uint8 { return uint8(le<<1) | dir }

// LE returns the logical endpoint number and its direction for the given
// hardware endpoint.
func LE(he uint8) (le int8, dir uint8) { return int8(he >> 1), he & 1 }
