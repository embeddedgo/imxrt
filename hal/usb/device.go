// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/mmio"
	"fmt"
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

// We use DTCM for buffers that require specific alignment and to avoid cache
// maintenance operations for EP0 transactions in ISR.
type dtcmem struct {
	qhs  [leNum * 2]dQH // queue heads, 1024 B, requires 4096 alignment
	dtd  DTD            // control data stage TD, 32 B, requires 32 alignment
	std  DTD            // control status stage TD, 32 B, requires 32 alignment
	data [256]byte      // data buffer, used in the control data stage
}

var dtcmCache [2]*dtcmem // cache the allocated DTCM for both USB controllers

type Device struct {
	u    *usb.Periph
	phy  *usbphy.Periph
	des  map[uint32][]byte
	dtcm *dtcmem
	pmu  sync.Mutex // prime mutex

	configured    atomic.Bool
	configuration uint8
}

func (d *Device) Print(i int) {
	qh := &d.dtcm.qhs[i]
	mmio.MB()
	fmt.Printf(
		"%#x qh[%d]: mult=%d zlt=%d maxpkt=%4d ios=%d current=%#x\n",
		d.u.ENDPTCTRL[i/2].Load(), i, qh.config>>30&3, qh.config>>29&1, qh.config>>16&0x3ff, qh.config>>15&1, qh.current,
	)
}

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
		m.dtd.next.Store(dtdEnd)
		m.std.next.Store(dtdEnd)
		dtcmCache[controller] = m
	}
	for he := range m.qhs {
		m.qhs[he].tlist.Store(dtdEnd)
	}
	d.dtcm = m
	return d
}

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
	d.dtcm.qhs[0].setConfig(64, dqhIOS) // Rx (host OUT)
	d.dtcm.qhs[1].setConfig(64, 0)      // Tx (host IN)
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
}

func (d *Device) Enable() {
	time.Sleep(10 * time.Millisecond) // ensure a long enough disconnect state
	d.u.USBCMD.SetBits(usb.RS)
}

func (d *Device) Disable() {
	d.configured.Store(false)
	d.u.USBCMD.ClearBits(usb.RS)
}

//go:nosplit
func (d *Device) ISR() {
	u := d.u
	status := u.USBSTS.Load()
	u.USBSTS.Store(status)

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
			// Remove completed DTDs from tlists and wakeup the goroutines that
			// waiting for them.
			ec &^= 1<<16 | 1
			for he, ec := 2, ec>>1; ec != 0; he, ec = he+2, ec>>1 {
				if he == 32 {
					he = 3
					ec >>= 1
				}
				if ec&1 == 0 {
					continue
				}
				qh := &d.dtcm.qhs[he]
				next := qh.tlist.Load()
				for next != dtdEnd {
					td := (*DTD)(unsafe.Pointer(next))
					if td.token&Active != 0 {
						break
					}
					next = td.next.Load()
					qh.tlist.Store(next)
					if td.note != nil {
						td.note.Wakeup()
					}
				}
			}
		}
	}

	if status&usb.URI != 0 {
		// 42.5.6.2.1 Bus Reset
		d.configured.Store(false)
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
		// Clean all tlists and wakeup gorutines that waits for the primed
		// transactions..
		for he := 2; he < len(d.dtcm.qhs); he++ {
			next := d.dtcm.qhs[he].tlist.Swap(dtdEnd)
			for next != dtdEnd {
				td := (*DTD)(unsafe.Pointer(next))
				if td.token&Active != 0 && td.note != nil {
					td.note.Wakeup()
				}
				next = td.next.Load()
			}
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
func setupRequest(d *Device, setup [2]uint32) {
	typ := uint8(setup[0] & 0x7f)
	req := uint16(setup[0] >> 7 & 0x1ff)
	val := uint16(setup[0] >> 16)
	idx := uint16(setup[1])
	siz := int(setup[1] >> 16)

	u := d.u
	// Standard device/interface/endpoint requests are handled in ISR directly.
	// Other requests are signaled to the handle goroutines.
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
			print("reqGetDescriptor\r\n")
			desc, ok := d.des[uint32(val)<<16|uint32(idx)]
			if !ok {
				print("unknown descr: ", uint32(val)<<16|uint32(idx), "\r\n")
				return
			}
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
			d.dtcm.data[0] = d.configuration
			d.prime(ep0tx, d.dataTD(1))
			d.prime(ep0rx, d.statusTD())
			return

		case reqSetConfiguration: // enables the device
			print("reqSetConfiguration\r\n")
			maxPkt := 64
			switch u.PORTSC1.LoadBits(usb.PSPD) >> usb.PSPDn {
			case 0:
				print(" full speed\r\n")
			case 1:
				print(" low speed\r\n")
			case 2:
				print(" high speed\r\n")
				maxPkt = 512
			default:
				print(" ??? speed\r\n")
			}

			d.configuration = uint8(val)
			// 42.5.6.3.1 Endpoint Initialization
			// TODO: this must be infered from descriptors
			const (
				CDC_ACM_ENDPOINT = 1
				CDC_RX_ENDPOINT  = 2
				CDC_TX_ENDPOINT  = 2
			)
			d.dtcm.qhs[CDC_ACM_ENDPOINT*2+1].setConfig(16, 0)
			d.dtcm.qhs[CDC_RX_ENDPOINT*2+0].setConfig(maxPkt, 0)
			d.dtcm.qhs[CDC_TX_ENDPOINT*2+1].setConfig(maxPkt, 0)
			mmio.MB()
			u.ENDPTCTRL[CDC_ACM_ENDPOINT].Store(3<<usb.TXTn | usb.TXR | usb.TXE | 2<<usb.RXTn)                                  // interrupt
			u.ENDPTCTRL[CDC_RX_ENDPOINT].Store(2<<usb.RXTn | usb.RXR | usb.RXE | 2<<usb.TXTn | usb.TXR | usb.TXE | 2<<usb.RXTn) // bulk, RX + TX
			d.prime(ep0tx, d.statusTD())

			d.configured.Store(true)
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

// Prime primes the he hardware endpoint with tdl list of transfer descriptors.
// It reports whether the endpoint was succesfully primed. The device must be
// in configured state to priming an endpoint succesfully, i.e. Prime alwyas
// fails in any other device state (powered, attach, reset, default FS/HS).
//
// The last descriptor in the tdl must have a note set to provide a way to
// inform about the end of transfer (see DTD.SetNote). Setting a note for the
// preceding DTDs in the list is optional and depends on the logical structure
// of the transfer.
func (d *Device) Prime(he int, tdl *DTD) (primed bool) {
	if !d.configured.Load() {
		return false
	}
	qh := &d.dtcm.qhs[he]
	// Append tdl to the end of he list.
	for {
		next := &qh.tlist
		for {
			p := next.Load()
			if p == dtdEnd {
				break
			}
			next = &(*DTD)(unsafe.Pointer(p)).next
		}
		if next.CompareAndSwap(dtdEnd, tdl.uintptr()) {
			break
		}
	}
	mask := uint32(1) << (he & 1 * 16) << (he >> 1)
	u := d.u
	if u.ENDPTPRIME.LoadBits(mask) != 0 {
		return true
	}
	d.pmu.Lock()
	var status uint32
	for {
		u.USBCMD.SetBits(usb.ATDTW)
		status = u.ENDPTSTAT.Load()
		if u.USBCMD.LoadBits(usb.ATDTW) != 0 {
			break
		}
	}
	u.USBCMD.ClearBits(usb.ATDTW)
	if status&mask == 0 {
		qh.next = tdl.uintptr()
		qh.token = 0
		mmio.MB()
		u.ENDPTPRIME.SetBits(mask)
	}
	d.pmu.Unlock()
	return true
}
