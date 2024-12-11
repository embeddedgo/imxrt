// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dtcm allows to allocate variables and slices of arbitrary type in
// the DTCM memory. The allocation is permanent (there is no way to free
// allocated memory).
//
// The actual size of DTCM is determined by the plugin code in the mbr.img. See
// tools/imxmbr for more information.
package dtcm

import (
	"math/bits"
	"sync/atomic"
	"unsafe"
)

const (
	base uintptr = 0x2000_0000
	end          = base + 32*1024 // Size of DTCM, must be in sync with mbr.img.
)

var free = base

func init() {
	runtime_memclrNoHeapPointers(unsafe.Pointer(base), end-base)
}

func alloc(align, size uintptr) unsafe.Pointer {
	if bits.OnesCount32(uint32(align)) != 1 {
		panic("bad align")
	}
	align--
	for {
		oldFree := atomic.LoadUintptr(&free)
		ptr := (oldFree + align) &^ align
		if ptr+size > end {
			panic("out of DTCM")
		}
		if atomic.CompareAndSwapUintptr(&free, oldFree, ptr+size) {
			return unsafe.Pointer(ptr)
		}
	}
}

// New works like new(T) but guarantees that the allocated variable has the
// given alignmet. Align must be a power of two.
func New[T any](align uintptr) (ptr *T) {
	return (*T)(alloc(align, unsafe.Sizeof(*ptr)))
}

// MakeSlice works like make([]T, len, cap) but guarantees that the returned
// slice has the given alignmet. Align must be a power of two.
func MakeSlice[T any](align uintptr, len, cap int) (slice []T) {
	ptr := alloc(align, unsafe.Sizeof(slice[0])*uintptr(cap))
	return unsafe.Slice((*T)(ptr), cap)[:len]
}

//go:linkname runtime_memclrNoHeapPointers runtime.memclrNoHeapPointers
func runtime_memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
