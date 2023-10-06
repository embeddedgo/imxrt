// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/rtos"
	"sync"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dtcm"
)

const (
	dqhIOS        = 1 << 15 // interrupt on setup
	dqhDisableZLT = 1 << 29 // zero length termination

	dqhMaxPktLenShift = 16
)

const dtdEnd uintptr = 1

type dQH struct {
	config  uint32
	current uintptr // *DTD (used by controller)
	next    uintptr // *DTD (used by controller/driver)
	token   uint32
	page    [5]uintptr
	head    uintptr // *DTD (used by driver)
	setup   [2]uint32
	new     uintptr // *DTD (used by driver)
	tail    uintptr // *DTD (used by driver)
	mu      sync.Mutex
}

func (qh *dQH) setConfig(maxPktLen int, flags uint32) {
	qh.config = uint32(maxPktLen)<<dqhMaxPktLenShift | flags
	qh.current = 0
	qh.next = dtdEnd
}

// DTD status
const (
	TransErr   = 1 << 3
	DataBufErr = 1 << 5
	Halted     = 1 << 6
	Active     = 1 << 7

	tokMultO = 3 << 10
	tokIOC   = 1 << 15
)

// A DTD is a Device Transfer Descriptor. It MUST BE allocated in the
// non-cacheable memory and 32 byte aligned. The NewDTD and MakeSliceDTD
// functions meet these requirements.
//
// The DTD is used for priming the USB controller endpoints always in the form
// of DTD list. The next field of the last DTD on the list used for priming may
// by changed implicitly so you cannot use nil as an end-of-list mark.
type DTD struct {
	next  uintptr // not a *DTD to avoid write barriers
	token uint32
	page  [5]uintptr
	note  *rtos.Note
}

/*
func (td *DTD) Print() {
	mmio.MB()
	fmt.Printf(
		" %p: next=%#x len=%3d ioc=%d mult=%d stat=0b%08b %#x\n",
		td, td.next.Load(), td.token>>16&0x7fff,
		td.token>>15&1, td.token>>10&3, td.token&0xff, td.page,
	)
}
*/

// NewDTD returns new DTD allocated in the non-cacheable memory. Use carefully
// because currently there is no way to release memory allocated this way.
func NewDTD() *DTD {
	return dtcm.New[DTD](32)
}

// MakeSliceDTD returns new slice of DTD structs allocated in non-cacheable
// memory.  Use carefully because currently there is no way to release memory
// allocated this way.
func MakeSliceDTD(len, cap int) []DTD {
	return dtcm.MakeSlice[DTD](32, len, cap)
}

func (td *DTD) uintptr() uintptr {
	return uintptr(unsafe.Pointer(td))
}

// Status returns a transfer status. If td was used to prime an USB controller
// endpoint the returned value is only valid after waking up from the note.Sleep
// (see SetNote method) that signals the end of transfer to which this td
// belongs to.
//
// N contains the number of bytes in the buffer that remain untransfered.
//
// After a successful transaction, the status byte should be zero. If not zero,
// the Active bit means unfinished transfer due to the Bus Reset and the
// meanings of the remaining bits (Halted, DataBufErr, TransErr) can be found
// in the documentation of the Endpoint Transfer Descriptor (dTD).
func (td *DTD) Status() (n int, status uint8) {
	return int(td.token >> 16 & 0x7fff), uint8(td.token & (Active | Halted | DataBufErr | TransErr))
}

// SetNext sets the td.next field to next.
func (td *DTD) SetNext(next *DTD) {
	td.next = uintptr(unsafe.Pointer(next))
}

// Next returns the content of td.next.
func (td *DTD) Next() *DTD {
	return (*DTD)(unsafe.Pointer(td.next))
}

// SetupTransfer configures td to use the buffer specified by ptr and size for a
// data transfer. As the maximum transfer length that can be handled by single
// DTD is limited it returns how much of the buffer will be used. The limit
// depends on the buffer alignment in memory and can be any number from 16 KiB
// to 20 KiB. The remaining part of the buffer can be transfered using a next
// DTD in the list or assigned to the same DTD next time. In most cases the
// bufer requires a cache maintanance (see also dma.New, dma.MakeSlice,
// rtos.CacheMaint) and  must be keep referenced until the end of transfer to
// avoid GC. The unsafe.Pointer type is there to remind you of both of these
// inconveniences.
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

// SetNote sets the Interrupt On Complete bit (IOC) in the td.token field and
// a note that will be used by an interrupt handler to communicate the
// completion of a transfer. As the Go GC may have no access to the td.note
// field you must keep a reference to the note somewhere. For the same reason
// the DTD type does not provide a method to obtain the note set.
//
// Set note to nil to clear IOC.
//
// After waking up from the note.Sleep the Status method should be used to check
// status of the transfer. Check all transfer descriptors in the list that may
// be signaled by this note.
func (td *DTD) SetNote(note *rtos.Note) {
	td.note = note
	if note != nil {
		td.token |= tokIOC
	} else {
		td.token &^= tokIOC
	}
}
