// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/mmio"
	"embedded/rtos"
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

const leNum = 8

// Hardware endpoints
const (
	ep0rx = 0
	ep0tx = 1
)

type noteNext struct {
	note rtos.Note
	next uintptr // *noteNext
}

func (nn *noteNext) uintptr() uintptr {
	return uintptr(unsafe.Pointer(nn))
}

// We use DTCM for buffers that require specific alignment and to avoid cache
// maintenance operations for EP0 transactions in ISR.
type dtcmem struct {
	qhs  [leNum * 2]dQH // queue heads, 1024 B, requires 4096 alignment
	dtd  DTD            // control data stage TD, 32 B, requires 32 alignment
	std  DTD            // control status stage TD, 32 B, requires 32 alignment
	data [256]byte      // data buffer, used in the control data stage
}

var dtcmCache [2]*dtcmem // cache the allocated DTCM for both USB controllers

// A Device represents an USB Device Controler Driver (DCD).
type Device struct {
	u      *usb.Periph
	phy    *usbphy.Periph
	des    map[uint32][]byte
	dtcm   *dtcmem
	pmu    sync.Mutex // prime mutex
	config atomic.Uint32
	cwl    atomic.Uintptr // *wait
	cwlmu  sync.Mutex
	ctreq  rtos.Note
	cthm   map[uint16]func(r *ControlRequest) int
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
		m.dtd.next = dtdEnd
		m.std.next = dtdEnd
		dtcmCache[controller] = m
	}
	for he := range m.qhs {
		qh := &m.qhs[he]
		qh.head = dtdEnd
	}
	d.dtcm = m
	return d
}

// Handle registers the handler for the endpoint 0 control request with the
// given request (see ControlRequest). All handlers must be registered before
// enabling the device.
func (d *Device) Handle(request uint16, handler func(r *ControlRequest) int) {

}

// Init initializes the USB device controler and the driver itself.
func (d *Device) Init(intPrio int, descriptors map[uint32][]byte, forceFullSpeed bool) {
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

	//go ctrlReqHandler(d)
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

//go:nosplit
func (d *Device) ISR() {
	u := d.u
	status := u.USBSTS.Load()
	u.USBSTS.Store(status)

	//print("ISR ", status, "\r\n")

	if status&usb.UI != 0 {
		// Check for setup reques
		for {
			ess := u.ENDPTSETUPSTAT.Load() & (1<<leNum - 1)
			if ess == 0 {
				break
			}
			u.ENDPTSETUPSTAT.Store(ess) // clear
			for le := uint(0); ess != 0; le, ess = le+1, ess>>1 {
				if ess&1 == 0 {
					continue
				}
				var setup [2]uint32
				for {
					u.USBCMD.SetBits(usb.SUTW)
					setup = d.dtcm.qhs[le*2].setup
					mmio.MB() // ensure setup read before checking SUTW
					if u.USBCMD.LoadBits(usb.SUTW) != 0 {
						break
					}
				}
				u.USBCMD.ClearBits(usb.SUTW)
				flush := uint32(0x0001_0001) << le
				u.ENDPTFLUSH.Store(flush)
				for u.ENDPTFLUSH.LoadBits(flush) != 0 {
				}
				if le == 0 {
					setupRequest(d, setup)
				}
			}
		}

		if ec := u.ENDPTCOMPLETE.Load(); ec != 0 {
			u.ENDPTCOMPLETE.Store(ec) // clear
			// Wake up goroutines that wait for the completed transfers. This
			// code runs concurently with the Prime method (also look at the
			// comments written there).
			ec &^= 1<<16 | 1
			for he, ec := 2, ec>>1; ec != 0; he, ec = he+2, ec>>1 {
				if he == 32 {
					he = 3
					ec >>= 1
				}
				if ec&1 == 0 {
					continue
				}
				removeAndWakeup(&d.dtcm.qhs[he], Active)
			}
		}
	}

	if status&usb.URI != 0 {
		// 42.5.6.2.1 Bus Reset
		d.config.Store(0)
		u.ENDPTSETUPSTAT.Store(u.ENDPTSETUPSTAT.Load())
		u.ENDPTCOMPLETE.Store(u.ENDPTCOMPLETE.Load())
		for u.ENDPTPRIME.Load() != 0 {
		}
		u.ENDPTFLUSH.Store(0xffff_ffff)
		// The above clanup task must be performed before the end of reset.
		if u.PORTSC1.LoadBits(usb.PR) == 0 {
			// Too late. End of reset detected. Hardware reset needed.
			// BUG: Unlikely case, but not handled properly..
		}
		// Wake up goroutines that still waiting for the end of transfer.
		for he := 2; he < len(d.dtcm.qhs); he++ {
			removeAndWakeup(&d.dtcm.qhs[he], 0)
		}
	}

	if status&usb.SRI != 0 {
	}
	if status&usb.TI0 != 0 {
	}
	if status&usb.TI1 != 0 {
	}
	if status&usb.PCI != 0 {
	}
	if status&usb.SLI != 0 {
		// 42.5.6.2.2.1 Suspend. Could be signaled somehow to the application.
	}
	if status&usb.UEI != 0 {
		// BUG: there is no handling of USB errors
	}
}

//go:nosplit
func removeAndWakeup(qh *dQH, active uint32) {
	p := atomic.SwapUintptr(&qh.head, dtdEnd)
	for p != dtdEnd {
		td := (*DTD)(unsafe.Pointer(p))
		if td.token&active != 0 {
			qh.head = td.uintptr()
			break
		}
		p = td.next
		if p != dtdEnd {
			goto wakeup
		}
		// The last item on the list requires special treatment.
		// Let's try to mark td as selected for deletion.
		if !atomic.CompareAndSwapUintptr(&td.next, dtdEnd, 0) {
			// Failed, so td.next points to a newly added DTD list.
			p = td.next
			goto wakeup
		}
		// Marked. Try to clear the reference to td in qh.tail.
		if !atomic.CompareAndSwapUintptr(&qh.tail, td.uintptr(), 0) {
			// Failed, so qh.tail now points to the end of newly
			// added DTD list. Cannot wake up the goroutine waiting
			// for this td because it may still be referenced by the
			// appending goroutine.
			break
		}
	wakeup:
		if td.note != nil {
			td.note.Wakeup()
		}
	}
}

// Standard requests
const (
	reqGetStatus        = 0x0080 >> 7
	reqClearFeature     = 0x0100 >> 7
	reqSetFeature       = 0x0300 >> 7
	reqSetAdress        = 0x0500 >> 7
	reqGetDescriptor    = 0x0680 >> 7
	reqSetDescriptor    = 0x0700 >> 7
	reqGetConfiguration = 0x0880 >> 7
	reqSetConfiguration = 0x0900 >> 7
)

// Class requests
const (
	reqCDCSetLineCoding       = 0x2000 >> 7
	reqCDCSetControlLineState = 0x2200 >> 7
)

//go:nosplit
func (d *Device) statusTD() *DTD {
	d.dtcm.std.SetupTransfer(nil, 0)
	return &d.dtcm.std
}

//go:nosplit
func (d *Device) dataTD(n int) *DTD {
	if n > len(d.dtcm.data) {
		panic("dtcm.data buffer too small")
	}
	d.dtcm.dtd.SetupTransfer(unsafe.Pointer(&d.dtcm.data), n)
	return &d.dtcm.dtd
}

//go:nosplit
func setupRequest(d *Device, setup [2]uint32) {
	typ := uint8(setup[0] & 0x7f)
	req := uint16(setup[0] >> 7 & 0x1ff)
	val := uint16(setup[0] >> 16)
	idx := uint16(setup[1])
	siz := int(setup[1] >> 16)

	u := d.u
	// Standard device/interface/endpoint requests are handled directly in the
	// ISR. Other requests are forwarded to the registered callback functions
	// and executed in thread mode.
	switch typ {
	case 0x00: // Standard Device Request
		print("device: ")
		switch req {
		case reqGetStatus:
			print("reqGetStatus\r\n")
			d.dtcm.data[0] = 0
			d.dtcm.data[1] = 0
			d.prime(ep0tx, d.dataTD(2))
			d.prime(ep0rx, d.statusTD())
			return
		case reqClearFeature:
			print("reqClearFeature\r\n")

		case reqSetFeature:
			print("reqSetFeature\r\n")

		case reqSetAdress:
			print("reqSetAdress\r\n")
			d.prime(ep0tx, d.statusTD())
			addr := val & 0x7f
			u.DEVADDR_PLISTBASE.Store(1<<24 | uint32(addr)<<25)
			return
		case reqGetDescriptor:
			print("reqGetDescriptor ", uint32(val)<<16|uint32(idx), "\r\n")
			desc := d.des[uint32(val)<<16|uint32(idx)]
			n := len(desc)
			if n > siz {
				n = siz
			}
			n = copy(d.dtcm.data[:], desc[:n])
			d.prime(ep0tx, d.dataTD(n))
			d.prime(ep0rx, d.statusTD())
			return

		case reqSetDescriptor:
			print("reqSetDescriptor\r\n")

		case reqGetConfiguration:
			print("reqGetConfiguration\r\n")
			d.dtcm.data[0] = uint8(d.config.Load())
			d.prime(ep0tx, d.dataTD(1))
			d.prime(ep0rx, d.statusTD())
			return

		case reqSetConfiguration: // enables the device
			print("reqSetConfiguration\r\n")
			// Deconfigure endpoints.
			for i := 1; i < leNum; i++ {
				u.ENDPTCTRL[i].Store(0)
			}
			cnf := uint32(val) & 0xff
			if cnf != 0 {
				// Select the appropriate configuration descriptors.
				cfd := d.des[0x0200_0000]
				if u.PORTSC1.LoadBits(usb.PSPD)>>usb.PSPDn < 2 {
					cfd = d.des[0x0700_0000]
				}
				// Configure endpoints according to the endpoint descriptors.
				for len(cfd) > 2 {
					n := int(cfd[0])
					if len(cfd) < n {
						break
					}
					if n == 7 && cfd[1] == 5 && uint(cfd[2]&0x0f)-1 < leNum-1 {
						le := int(cfd[2] & 0x0f)
						dir := cfd[2] >> 7 // 0: Rx (OUT),  1: Tx (IN)
						shift := uint(dir) * 16
						he := le*2 + int(dir)
						typ := cfd[3] & 3
						maxPkt := int(cfd[4]) | int(cfd[5])<<8

						// 42.5.6.3.1 Endpoint Initialization
						flags := dqhDisableZLT & uint32(dir-1)
						d.dtcm.qhs[he].setConf(maxPkt, flags)
						mask := usb.ENDPTCTRL(0xffff) << shift
						other := u.ENDPTCTRL[le].LoadBits(^mask)
						if typ != 0 && other == 0 {
							other = 2 << usb.TXTn >> shift
						}
						cfg := usb.ENDPTCTRL(typ)<<usb.RXTn | usb.RXR | usb.RXE
						mmio.MB()
						u.ENDPTCTRL[le].Store(other | cfg<<shift)
					}
					cfd = cfd[n:]
				}
			}
			d.prime(ep0tx, d.statusTD())
			d.config.Store(cnf)
			if cnf == 0 {
				return
			}
			for {
				p := d.cwl.Load()
				if p == 0 {
					break
				}
				nn := (*noteNext)(unsafe.Pointer(p))
				if d.cwl.CompareAndSwap(p, nn.next) {
					nn.note.Wakeup() // succesfully removed w so send the note
				}
			}
			return
		}
	case 0x01: // Standard Interface Request
		print("interface: ?\r\n")

	case 0x02: // Standard Endpoint Request
		print("endpoint: ")
		le := idx & 0x7F
		print("endpoint ", le, ": ")
		if le > 7 {
			return
		}
		epctl := &u.ENDPTCTRL[le]
		mask := usb.RXS
		if idx&0x80 != 0 {
			mask = usb.TXS
		}
		switch req {
		case reqGetStatus:
			print("reqGetStatus\r\n")
			stall := byte(0)
			if epctl.LoadBits(mask) != 0 {
				stall = 1
			}
			d.dtcm.data[1] = stall
			d.dtcm.data[1] = 0
			d.prime(ep0tx, d.dataTD(2))
			d.prime(ep0rx, d.statusTD())
		case reqClearFeature:
			print("reqClearFeature\r\n")
			epctl.ClearBits(mask)
			d.prime(ep0tx, d.statusTD())
		case reqSetFeature:
			print("reqSetFeature\r\n")
			epctl.SetBits(mask)
			d.prime(ep0tx, d.statusTD())
		}
		return
	case 0x21: // Class Interface Request
		switch req {
		case reqCDCSetLineCoding:
			print("reqCDCSetLineCoding\r\n")
			d.prime(ep0rx, d.dataTD(7))
			d.prime(ep0tx, d.statusTD())
			return

		case reqCDCSetControlLineState:
			print("reqCDCSetControlLineState\r\n")
			// idx contains CDC_STATUS_INTERFACE id to distinguish beetwen
			// multiple CDC interfaces, val contains RTS DTR config bits.
			d.prime(ep0tx, d.statusTD())
			return
		}
	}
	print("unknown: ", setup[0]&0xffff, " typ=", typ, " req=", req, " val=", val, " usb stall\r\n")
	u.ENDPTCTRL[0].Store(usb.RXS | usb.TXS) // 42.5.6.3.2 Protocol stall
}

//go:nosplit
func (d *Device) prime(he int, td *DTD) {
	mask := uint32(1) << (he & 1 * 16) << (he >> 1)
	u := d.u
	for (u.ENDPTPRIME.Load()|u.ENDPTSTAT.Load())&mask != 0 {
		// This is simplified prime algorithm intended to be used in ISR. It can
		// prime inactive endpoint only (no support for dTD lists).
	}
	qh := &d.dtcm.qhs[he]
	qh.next = td.uintptr()
	qh.token = 0
	mmio.MB()
	u.ENDPTPRIME.SetBits(mask)
}

func ctrlReqHandler(d *Device) {

}

// Config returns the configuration number selected during the USB enumeration
// process or zero if the device is not in the configured state.
func (d *Device) Config() uint8 {
	return uint8(d.config.Load())
}

// WaitConfig waits for the selection of the cn configuration number during the
// USB enumeration process or simply for the configured state if cn is zero.
func (d *Device) WaitConfig(cn uint8) {
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

/*
func clean(tlist *uintptr) {
	p := *tlist
	for p != dtdEnd {
		td := (*DTD)(unsafe.Pointer(p))
		if atomic.LoadUint32(&td.token)&tokRemove == 0 {
			break
		}
		p = td.next
		if p == dtdEnd {
			break
		}
	}
	*tlist = p
}
*/

// Prime primes the he hardware endpoint with the list of transfer descriptors
// specified by first and last. It reports whether the endpoint was succesfully
// primed.
//
// To successfully prime an endpoint the device must be in the configured state
// and the selected configuration number must equal cn. Prime alwyas fails in
// any other device state (powered, attach, reset, default FS/HS).
//
// Prime panics if he is invalid or tdl is nil or cn is zero.
//
// The last descriptor in the tdl must have a note set to provide a way for the
// ISR to inform about the end of transfer (see DTD.SetNote). Setting notes for
// the preceding DTDs in the list is optional and depends on the logical
// structure of the transfer.
func (d *Device) Prime(he int, first, last *DTD, cn uint8) (primed bool) {
	if he < 2 || he >= len(d.dtcm.qhs) {
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
	qh.token = 0
	qh.next = first.uintptr()
	atomic.StoreUintptr(&qh.head, first.uintptr())
	u.ENDPTPRIME.SetBits(mask)
	qh.mu.Unlock()
	return true
}