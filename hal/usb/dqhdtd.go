// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"embedded/rtos"
	"fmt"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

const (
	dqhIOS = 1 << 15 // interrupt on setup
	dqhZLT = 1 << 29 // zero length termination

	dqhMaxPktLenShift = 16
)

const dtdEnd uintptr = 1

type dQH struct {
	config  uint32
	current uintptr // *DTD
	overlay DTD
	setup   [2]uint32
	head    *DTD
	tail    *DTD
	_       [2]uint32 // padding to make dQH 64 bytes in size, unused
}

func (qh *dQH) setConfig(maxPktLen int, flags uint32) {
	qh.config = uint32(maxPktLen)<<dqhMaxPktLenShift | flags
	qh.current = 0
	*(*uintptr)(unsafe.Pointer(&qh.overlay.next)) = dtdEnd
}

// DTD.Token bit fields
const (
	tokActive uint32 = 1 << 7
	tokMultO  uint32 = 3 << 10
	tokIOC    uint32 = 1 << 15
)

type DTD struct {
	next  *DTD
	token uint32
	page  [5]uintptr
	note  unsafe.Pointer // *rtos.Note or some pointer
}

func (td *DTD) Print() {
	rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(td), 32)
	fmt.Printf(
		" %p: next=%p len=%3d ioc=%d mult=%d stat=0b%08b %#x\n",
		td, td.next, td.token>>16&0x7fff,
		td.token>>15&1, td.token>>10&3, td.token&0xff, td.page,
	)
}

func NewDTD() *DTD {
	td := dma.New[DTD]()
	td.next = (*DTD)(unsafe.Pointer(dtdEnd))
	return td
}

func MakeSliceDTD(len, cap int) []DTD {
	tds := dma.MakeSlice[DTD](len, cap)
	for i := range tds {
		tds[i].next = (*DTD)(unsafe.Pointer(dtdEnd))
	}
	return tds
}

// SetNext sets the next field in the DTD to next. Use SetLast to mark this DTD
// as the last one in the list (don't use nil because the controller doesn't
// recognize it as termination mark).
func (td *DTD) SetNext(next *DTD) {
	td.next = next
	rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(td), 32)
}

// SetLast marks this DTD as the last in the DTD list.
func (td *DTD) SetLast() {
	td.next = (*DTD)(unsafe.Pointer(dtdEnd))
	rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(td), 32)
}

// Next returns a pointer to the next DTD on the DTD lists. It returns nil if
// the next field is nil or contains a last mark.
func (td *DTD) Next() *DTD {
	next := td.next
	if uintptr(unsafe.Pointer(next)) == dtdEnd {
		next = nil
	}
	return next
}

func (td *DTD) setNextNoWB(next *DTD) {
	*(*uintptr)(unsafe.Pointer(&td.next)) = uintptr(unsafe.Pointer(next))
}

// SetupTransfer configures d to use the bufer specified by ptr and size for a
// data transfer. As the maximum transfer length that can be handled by single
// DTD is limited it returns how much of the buffer will be used. The limit
// depends on the buffer alignment in memory and can be any number from 16 to
// 20 KiB. The remaining part of the buffer can be transfered using a next DTD
// in the list or assigned to the same DTD next time. In most cases the bufer
// requires a cache maintanance (see also dma.New, dma.MakeSlice,
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
	td.token = td.token&(tokIOC|tokMultO) | uint32(n<<16) | tokActive
	rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(td), 32)
	return
}

func (td *DTD) SetNote(note *rtos.Note) {
	td.note = unsafe.Pointer(note)
	if note != nil {
		td.token |= tokIOC
	} else {
		td.token &^= tokIOC
	}
	rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(td), 32)
}

func (td *DTD) Note() *rtos.Note {
	return (*rtos.Note)(td.note)
}

// SetRef can be used to store a pointer in DTD. SetRef shares the same field
// in DTD that is used by SetNote. It is intended to be used to store a
// reference to the data buffer set by SetBuf in case you do not want to wait
// for the transaction to complete and do not want to reuse the buffer again.
// Storing a reference to the buffer in DTD ensures no GC until the transaction
// completes.
func (td *DTD) SetRef(ptr unsafe.Pointer) {
	td.note = ptr
	td.token &^= tokIOC
	rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(td), 32)
}

func (td *DTD) Ref() unsafe.Pointer {
	if td.token&tokIOC == 0 {
		return td.note
	}
	return nil
}

func (td *DTD) Status() uint32 {
	rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(td), 32)
	return td.token
}
