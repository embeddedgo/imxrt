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

	"github.com/embeddedgo/imxrt/hal/dtcm"
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
// maintenance operations for EP0 transactions in ISR.
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
	pmu    sync.Mutex // prime mutex
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
func (d *Device) Print(i int) {
	qh := &d.dtcm.qhs[i]
	mmio.MB()
	fmt.Printf(
		"%#x qh[%d]: mult=%d zlt=%d maxpkt=%4d ios=%d current=%#x\n",
		d.u.ENDPTCTRL[i/2].Load(), i, qh.config>>30&3, qh.config>>29&1, qh.config>>16&0x3ff, qh.config>>15&1, qh.current,
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
	}
	d.dtcm = m
	d.cr.Data = m.isr.data[:] // cannot be set in ISR because of write barriers
	d.crhm = make(map[uint32]func(r *ControlRequest) int)
	return d
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
	u.USBINTR.Store(usb.UE | usb.UEE | usb.PCE | usb.URE | usb.SLE)

	go handleControRequests(d)
}

// ControlRequest represents an USB Control Request. The LE field is the logical
// endpoint number the request is adressed to. The remaining fields are derived
// from the Setup Stage. The length of the Data field is also specified by the
// setup packed but its content is related to the optional Data Stage. The Data
// field serves two functions. In case of OUT direction it contains the data
// received from the host, otherwise (IN direction) it serves as the output
// buffer, for sending data to the host.
type ControlRequest struct {
	LE      int    // logical endpoint
	Request uint16 // bRequest<<8 | bmRequestType
	Value   uint16 // wValue
	Index   uint16 // wIndex
	Data    []byte // len(Data) = wLength
}

// Handle registers the handler for the control requests adressed to the logical
// endpoint number le. All handlers must be registered before enabling the
// device.
func (d *Device) Handle(le int, request uint16, handler func(cr *ControlRequest) int) {
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
			crst >>= uint(n)
			cr.LE += n
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
func (d *Device) Config() uint8 {
	return uint8(d.config.Load())
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

// Prime primes the he hardware endpoint with the list of transfer descriptors
// specified by the first and last pointers. It reports whether the endpoint was
// succesfully primed.
//
// To successfully prime an endpoint the device must be in the configured state
// and the selected configuration number must equal cn. Prime alwyas fails in
// any other device state (powered, attach, reset, default FS/HS).
//
// Prime can be used concurently by multiple goroutines also with the same
// endpoint.
//
// The last descriptor in the tdl must have a note set to provide a way for the
// ISR to inform about the end of transfer (see DTD.SetNote). Setting notes for
// the preceding DTDs in the list is optional and depends on the logical
// structure of the transfer.
func (d *Device) Prime(he int, first, last *DTD, cn int) (primed bool) {
	if uint(he-2) >= uint(len(d.dtcm.qhs)-2) {
		panic("bad he")
	}
	if first == nil {
		panic("first == nil")
	}
	if last == nil {
		panic("last == nil")
	}
	if cn == 0 {
		panic("cn == 0")
	}

	last.next = dtdEnd
	if d.config.Load() != uint32(cn) {
		return false
	}

	qh := &d.dtcm.qhs[he]
	qh.mu.Lock()

	// Exclusive appending, still executed concurently with the removing ISR
	// (also look at the comments written there).

	var status uint32
	mask := uint32(1) << (he & 1 * 16) << (he >> 1)
	u := d.u
	tail := qh.tail
	td := (*DTD)(unsafe.Pointer(tail))

	if tail == 0 {
		// The list is empty.
		goto fastPrime
	}
	if !atomic.CompareAndSwapUintptr(&qh.tail, tail, last.uintptr()) {
		// The list has just been emptied by ISR.
		goto fastPrime
	}
	// The list seems to be non-empty. Let's try append tdl to its end.
	if next := td.next; next == 0 || !atomic.CompareAndSwapUintptr(&td.next, next, first.uintptr()) {
		// The ISR marked the list as empty but didn't finished its work because
		// we managed to CAS qh.tail successfully. For this reason it didn't
		// handled td.note so we do it here.
		if td.note != nil {
			td.note.Wakeup()
		}
		goto fastPrime1
	}

	// We appended tdl successfully to the non-empty list so the full prime
	// algorithm is required.

	if d.config.Load() != uint32(cn) {
		qh.mu.Unlock()
		goto reset
	}

	// Check if the endpoint has just been (re)primed.
	if u.ENDPTPRIME.LoadBits(mask) != 0 {
		qh.mu.Unlock()
		return true
	}

	d.pmu.Lock() // prevent other goroutines from overtaking us for priming
	qh.mu.Unlock()

	// Check the endpoint status. Prime if not active.
	for {
		u.USBCMD.SetBits(usb.ATDTW)
		status = u.ENDPTSTAT.Load()
		if u.USBCMD.LoadBits(usb.ATDTW) != 0 {
			break
		}
	}
	u.USBCMD.ClearBits(usb.ATDTW)
	if status&mask == 0 {
		qh.next = first.uintptr()
		qh.token = 0
		mmio.MB()
		u.ENDPTPRIME.SetBits(mask)
	}

	d.pmu.Unlock()
	return true

	// If the endpoint is inactive we can use a simple prime algorithm.
fastPrime:
	qh.tail = last.uintptr()
fastPrime1:

	if d.config.Load() != uint32(cn) {
		qh.mu.Unlock()
		goto reset
	}

	qh.token = 0
	qh.next = first.uintptr()
	atomic.StoreUintptr(&qh.head, first.uintptr())
	u.ENDPTPRIME.SetBits(mask)
	qh.mu.Unlock()
	return true

reset:
	d.u.USBCMD.ClearBits(usb.RS)
	time.Sleep(10 * time.Millisecond)
	d.u.USBCMD.SetBits(usb.RS)
	return false
}
