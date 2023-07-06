// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

const cacheLineSize = 32 // Cortex-M7

const CacheLineSize = cacheLineSize

func alloc(size uintptr) unsafe.Pointer {
	size = (size + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	size += cacheLineSize // extra space for address alignment
	buf := make([]byte, size)
	addr := uintptr(unsafe.Pointer(&buf[0]))
	addr = (addr + (cacheLineSize - 1)) &^ (cacheLineSize - 1)
	return unsafe.Pointer(addr)
}

// New works like new(T) but guarantees that the allocated variable does not
// share the same line in data cache with any other variable.
func New[T any]() (ptr *T) {
	return (*T)(alloc(unsafe.Sizeof(*ptr)))
}

// MakeSlice works like make([]T, len, cap) but guarantees that the returned
// slice does not share the same line in data cache with any other variable.
func MakeSlice[T any](len, cap int) (slice []T) {
	ptr := alloc(unsafe.Sizeof(slice[0]) * uintptr(cap))
	return unsafe.Slice((*T)(ptr), cap)[:len]
}
