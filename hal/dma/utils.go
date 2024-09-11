// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

// AlignOffsets calculatest the start and end offsets to the cache aligned
// portion of the memory described by ptr and size.
func AlignOffsets(ptr unsafe.Pointer, size uintptr) (start, end uintptr) {
	const cacheAlignMask = CacheLineSize - 1
	p := uintptr(ptr)
	start = -p & cacheAlignMask
	end = size - (p+size)&cacheAlignMask
	return
}
