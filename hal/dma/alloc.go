// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

const cacheLineSize = 32 // Cortex-M7

const CacheLineSize = cacheLineSize

// Alloc works like make([]T, n) but guarantees that the returned slice does
// not share the same line in data cache with any other variable.
func Alloc[T any](n int) (cacheAligned []T) {
	size := unsafe.Sizeof(cacheAligned[0]) * uintptr(n)
	size = (size + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	size += cacheLineSize // extra space for address alignment
	buf := make([]byte, size)
	addr := uintptr(unsafe.Pointer(&buf[0]))
	addr = (addr + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	return unsafe.Slice((*T)(unsafe.Pointer(addr)), n)
}

/*
// AllocBytes allocates cache aligned buffer in RAM. It guarantees that the returned
// slice does not share the same line in data cache with any other variable.
func AllocBytes(size uintptr) (cacheAligned []byte) {
	n := (size + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	n += cacheLineSize // extra space for address alignment
	buf := make([]byte, n)
	addr := uintptr(unsafe.Pointer(&buf[0]))
	addr = (addr + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	return (*[size]byte)(unsafe.Pointer(addr))[:]
}
*/
