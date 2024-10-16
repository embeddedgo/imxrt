// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/mmio"
	"embedded/rtos"
	"math/bits"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/p/usb"
)

// All functions/methods that may run in the interrupt context (Cortex-M handler
// mode) should be placed in this file.
//
// All functions in this file must have the go:nosplit directive.

// ISR handles USB interrupts.
//
//go:nosplit
//go:nowritebarrierrec
func (d *Device) ISR() {
	u := d.u
	status := u.USBSTS.Load()
	u.USBSTS.Store(status)

	//print("ISR ", status, "\r\n")

	if status&usb.UI != 0 {
		// Check for setup request.
		for {
			// 42.5.6.4.2.1 Setup Phase
			ess := uint16(u.ENDPTSETUPSTAT.Load())
			if ess == 0 {
				break
			}
			u.ENDPTSETUPSTAT.Store(uint32(ess)) // clear
			for le := 0; ess != 0; le, ess = le+1, ess>>1 {
				n := bits.TrailingZeros16(ess)
				le, ess = le+n, ess>>uint(n) // skip the zero bits
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
				flush := uint32(0x0001_0001) << uint(le)
				u.ENDPTFLUSH.Store(flush)
				for u.ENDPTFLUSH.LoadBits(flush) != 0 {
				}
				if le == 0 && setup[0]&0x7f <= 2 {
					// Standard device/interface/endpoint requests are handled
					// directly in the ISR.
					n := parseSetup(&d.cr, setup)
					d.cr.Data = d.cr.Data[:n:maxCtrlData] // avoid write barrier
					execContorHandler(d, &d.dtcm.isr, &d.cr, d.controlHandlerISR)
				} else {
					// Other requests are forwarded to the regristered handlers
					// and executed in thread mode.
					d.crsa[le] = setup
					for {
						crst := d.crst
						if atomic.CompareAndSwapUint32(&d.crst, crst, crst|1<<le) {
							break
						}
					}
					d.crno.Wakeup()
				}
			}
		}
		// Handle completed transfers.
		if ec := u.ENDPTCOMPLETE.LoadBits(0xfffe_fffe); ec != 0 {
			u.ENDPTCOMPLETE.Store(ec) // clear
			// Wake up goroutines that wait for the completed transfers. This
			// code runs concurently with the Prime method (also look at the
			// comments written there).
			for i := 0; ec != 0; i, ec = i+1, ec>>1 {
				n := bits.TrailingZeros32(ec)
				i, ec = i+n, ec>>uint(n) // skip the zero bits
				he := i&15<<1 | i>>4
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
			// BUG: Unlikely case, but not handled properly.
		}
		// Wake up goroutines that still waiting for the end of transfer.
		for he := range d.dtcm.qhs {
			qh := &d.dtcm.qhs[he]
			if he > 2 {
				removeAndWakeup(qh, 0)
			}
			qh.next = dtdEnd
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
}

//go:nosplit
func (td *DTD) uintptr() uintptr {
	return uintptr(unsafe.Pointer(td))
}

// SetupTransfer configures td to use the buffer specified by ptr and size for a
// data transfer. As the maximum transfer length that can be handled by single
// DTD is limited it returns how much of the buffer will be used. The limit
// depends on the buffer alignment in memory and can be any number from 16 KiB
// to 20 KiB. The remaining part of the buffer can be transfered using a next
// DTD in the list or assigned to the same DTD next time. In most cases the
// bufer requires a cache maintanance (see also dma.New, dma.MakeSlice,
// rtos.CacheMaint) and must be keep referenced until the end of transfer to
// avoid GC. The unsafe.Pointer type is there to remind you of both of these
// inconveniences.
//
//go:nosplit
func (td *DTD) SetupTransfer(ptr unsafe.Pointer, size int) (n int) {
	if size > 0 {
		addr := uintptr(ptr)
		td.page[0] = addr
		pa := addr&^0x0fff + 0x1000
		td.page[1] = pa
		pa += 0x1000
		td.page[2] = pa
		pa += 0x1000
		td.page[3] = pa
		pa += 0x1000
		td.page[4] = pa
		pa += 0x1000
		n = int(pa - addr)
		if n > size {
			n = size
		}
	}
	td.token = td.token&(tokIOC|tokMultO) | uint32(n<<16) | Active
	return
}

//go:nosplit
func (qh *dQH) setConf(maxPktLen int, flags uint32) {
	qh.config = uint32(maxPktLen)<<dqhMaxPktLenShift | flags
	qh.current = 0
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
		if !atomic.CompareAndSwapUintptr(&td.next, dtdEnd, dtdRm) {
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

// Prime performs simplified prime algorithm. Intended for control endpoints.
// Can be used in ISR.
//
//go:nosplit
func (d *Device) prime(he uint8, td *DTD) {
	mask := uint32(1) << (he & 1 * 16) << (he >> 1)
	u := d.u
	qh := &d.dtcm.qhs[he]
	for qh.next != dtdEnd && u.ENDPTCOMPLETE.LoadBits(mask) == 0 {
		// Wait for the previous transfer to complete.
		mmio.MB()
	}
	u.ENDPTCOMPLETE.Store(mask) // clear
	qh.next = td.uintptr()
	qh.token = 0
	mmio.MB()
	u.ENDPTPRIME.SetBits(mask)
}

//go:nosplit
func badControlRequest(d *Device, cr *ControlRequest) {
	print("Unknown or bad USB control reqeust:\r\n")
	print(" LE:      ", cr.LE, "\r\n")
	print(" Request: ", cr.Request, "\r\n")
	print(" Value:   ", cr.Value, "\r\n")
	print(" Index:   ", cr.Index, "\r\n")
	print(" DataLen: ", len(cr.Data), "\r\n")
	// 42.5.6.3.2 Protocol stall
	d.u.ENDPTCTRL[cr.LE].Store(usb.RXS | usb.TXS)
}

//go:nosplit
func execContorHandler(d *Device, ctds *ctds, cr *ControlRequest, h func(r *ControlRequest) int) {
	he := uint8(cr.LE) * 2
	she := he // status he
	if cr.Request>>7&1 == 0 {
		she++
		// Receive data
		if len(cr.Data) != 0 {
			ctds.dtd.SetupTransfer(unsafe.Pointer(&cr.Data[0]), len(cr.Data))
			d.prime(he, &ctds.dtd)
		}
	}
	n := h(cr)
	if n < 0 {
		badControlRequest(d, cr)
		return
	}
	if he == she {
		// Send data
		var p unsafe.Pointer
		if n != 0 {
			p = unsafe.Pointer(&cr.Data[0])
		}
		ctds.dtd.SetupTransfer(p, n)
		d.prime(he+1, &ctds.dtd)
	}
	// Send/receive status.
	ctds.std.SetupTransfer(nil, 0)
	d.prime(she, &ctds.std)
}

//go:nosplit
func parseSetup(cr *ControlRequest, setup [2]uint32) int {
	cr.Request = uint16(setup[0])
	cr.Value = uint16(setup[0] >> 16)
	cr.Index = uint16(setup[1])
	n := int(setup[1] >> 16)
	if n > maxCtrlData {
		n = maxCtrlData
	}
	return n
}

// Standard requests (contain the direction bit).
const (
	reqGetStatus        = 0x00<<1 | 1
	reqClearFeature     = 0x01<<1 | 0
	reqSetFeature       = 0x03<<1 | 0
	reqSetAdress        = 0x05<<1 | 0
	reqGetDescriptor    = 0x06<<1 | 1
	reqSetDescriptor    = 0x07<<1 | 0
	reqGetConfiguration = 0x08<<1 | 1
	reqSetConfiguration = 0x09<<1 | 0
	reqGetInterface     = 0x0a<<1 | 1
	reqSetInterface     = 0x11<<1 | 0
)

//go:nosplit
func (d *Device) controlHandlerISR(cr *ControlRequest) int {
	typ := uint8(cr.Request & 0x7f) // request type without direction
	req := cr.Request >> 7 & 0x1ff  // request number with direction
	switch typ {
	case 0x00: // Standard Device Request
		switch req {
		case reqGetStatus:
			if len(cr.Data) < 2 {
				break
			}
			cr.Data[0] = 0 // bus powered, DEVICE_REMOTE_WAKEUP unset
			cr.Data[1] = 0
			return 2

		case reqClearFeature, reqSetFeature:
			// Not supported.
			// Used to clear/set DEVICE_REMOTE_WAKEUP, TEST_MODE.

		case reqSetAdress:
			addr := cr.Value & 0x7f
			d.u.DEVADDR_PLISTBASE.Store(1<<24 | uint32(addr)<<25)
			return 0

		case reqGetDescriptor:
			key := uint32(cr.Value)<<16 | uint32(cr.Index)
			desc := d.des[key]
			return copy(cr.Data, desc)

		case reqSetDescriptor:
			// Not supported.

		case reqGetConfiguration:
			if len(cr.Data) < 1 {
				break
			}
			cr.Data[0] = uint8(d.config.Load())
			return 1

		case reqSetConfiguration: // enables the device
			// Deconfigure endpoints.
			u := d.u
			for i := 1; i < leNum; i++ {
				u.ENDPTCTRL[i].Store(0)
			}
			cnf := uint32(cr.Value) & 0xff
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
					if n == 7 && cfd[1] == 5 && uint(cfd[2]&0x0f)-1 < uint(leNum)-1 {
						le := int(cfd[2] & 0x0f)
						dir := cfd[2] >> 7 // 0: Rx (OUT),  1: Tx (IN)
						shift := uint(dir) * 16
						he := le*2 + int(dir)
						typ := cfd[3] & 3
						maxPkt := int(cfd[4]) | int(cfd[5])<<8

						// 42.5.6.3.1 Endpoint Initialization
						flags := dqhDisableZLT & (uint32(dir) - 1)
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
			d.config.Store(cnf)
			if cnf == 0 {
				return 0
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
			return 0
		}
	case 0x01: // Standard Interface Request
		switch req {
		case reqGetStatus:
			if len(cr.Data) < 2 {
				break
			}
			cr.Data[0] = 0
			cr.Data[1] = 0
			return 2
		case reqGetInterface:
			if len(cr.Data) < 1 {
				break
			}
			cr.Data[0] = 0
			return 1
		case reqSetInterface:
			// TODO: support alternate interface settings
		}

	case 0x02: // Standard Endpoint Request
		le := int(cr.Index & 0x7F)
		if le >= leNum {
			return -1
		}
		epctl := &d.u.ENDPTCTRL[le]
		mask := usb.RXS
		if cr.Index&0x80 != 0 {
			mask = usb.TXS
		}
		switch req {
		case reqGetStatus:
			if len(cr.Data) < 2 {
				break
			}
			stall := internal.BoolToInt(epctl.LoadBits(mask) != 0)
			cr.Data[0] = byte(stall)
			cr.Data[1] = 0
			return 2
		case reqClearFeature:
			epctl.ClearBits(mask)
			return 0
		case reqSetFeature:
			epctl.SetBits(mask)
			return 0
		}
	}
	return -1
}

// Prime primes the he hardware endpoint with the list of transfer descriptors
// specified by the first and last pointers. It reports whether the endpoint was
// succesfully primed.
//
// To successfully prime an endpoint the device must be in the configured state.
//
// Prime can be used concurently by multiple goroutines also with the same
// endpoint.
//
// The last descriptor in the tdl must have a note set to provide a way for the
// ISR to inform about the end of transfer (see DTD.SetNote). Setting notes for
// the preceding DTDs in the list is optional and depends on the logical
// structure of the transfer.
//
//go:nosplit
func (d *Device) Prime(he uint8, first, last *DTD) (primed bool) {
	if uint(he-2) >= uint(len(d.dtcm.qhs)-2) {
		panic("bad he")
	}
	if first == nil {
		panic("first == nil")
	}
	if last == nil {
		panic("last == nil")
	}
	cfg := d.config.Load()
	if cfg == 0 {
		return false
	}

	var status uint32
	mask := uint32(1) << (he & 1 * 16) << (he >> 1)
	u := d.u
	qh := &d.dtcm.qhs[he]
	last.next = dtdEnd

	if rtos.HandlerMode() {
		// The code below is here mainly to support the print and println
		// functions used to debug or to print a stack trace in the IRQ handler
		// mode when USB serial is used as the system console.
		// BUG: unrealiable because of the possible concurrent acces to qh and u
		for qh.next&(dtdEnd&dtdRm) == 0 {
			// Wait for the posible active transfer to complete.
			mmio.MB()
		}
		u.ENDPTCOMPLETE.Store(mask) // clear
		qh.next = first.uintptr()
		qh.token = 0
		mmio.MB()
		u.ENDPTPRIME.SetBits(mask)
		for u.ENDPTCOMPLETE.LoadBits(mask) == 0 {
			// Ensure this transfer is completd before return because the system
			// (including USB) may be halted after that.
		}
		return true
	}

	qh.mu.Lock()
	defer qh.mu.Unlock()

	tail := qh.tail
	td := (*DTD)(unsafe.Pointer(tail))

	if tail == 0 || !atomic.CompareAndSwapUintptr(&qh.tail, tail, last.uintptr()) {
		// The list is empty or just been emptied by ISR.
		qh.tail = last.uintptr()
		goto primeEmpty
	}

	// The list seems to be non-empty. Let's try append our dTDs to its end.
	if next := td.next; next == dtdRm || !atomic.CompareAndSwapUintptr(&td.next, next, first.uintptr()) {
		// The ISR marked the list as empty but didn't finished its work because
		// we managed to CAS qh.tail successfully. For this reason it didn't
		// handled td.note so we do it here.
		if td.note != nil {
			td.note.Wakeup()
		}
		goto primeEmpty
	}

	// We appended our dTDs successfully to the non-empty list.

	// Check if the endpoint has just been (re)primed.
	if u.ENDPTPRIME.LoadBits(mask) != 0 {
		//fmt.Printf("pp %#x %#x %#x\n", qh.current, qh.next, qh.token)
		goto end
	}

	// Obtain the endpoint status.
	d.atdtwm.Lock()
	for {
		u.USBCMD.SetBits(usb.ATDTW)
		status = u.ENDPTSTAT.Load()
		if u.USBCMD.LoadBits(usb.ATDTW) != 0 {
			break
		}
	}
	d.atdtwm.Unlock()

	if status&mask != 0 || qh.current == last.uintptr() {
		// The endpoint is active or the controller already finished our dTDs.
		goto end
	}
	goto prime

	//fmt.Printf("## %#x %#x %#x %#x\n", status, qh.current, qh.next, qh.token)
primeEmpty:
	qh.head = first.uintptr()
prime:
	qh.token = 0
	qh.next = first.uintptr()
	mmio.MB()
	u.ENDPTPRIME.SetBits(mask)

end:
	if d.config.Load() == cfg {
		return true
	}
	// We primed the endpoint but in the meantime the active configuration
	// changed. Reset the USB.
	d.u.USBCMD.ClearBits(usb.RS)
	time.Sleep(10 * time.Millisecond)
	d.u.USBCMD.SetBits(usb.RS)
	return false
}
