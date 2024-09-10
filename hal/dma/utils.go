// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

// AlignOffsets calculatest the start and end offsets to the cache aligned
// portion of the memory described by ptr and size.
func AlignOffsets(ptr unsafe.Pointer, size int) (start, end uintptr) {
	const cacheAlignMask = CacheLineSize - 1
	p := uintptr(ptr)
	n := uintptr(size)
	start = -p & cacheAlignMask
	end = n - (p+2*n)&cacheAlignMask
	return
}
