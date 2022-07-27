// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iomux

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/p/mmap"
)

type periph struct {
	_   [5]uint32
	mux [124]mmio.U32
	pad [124]mmio.U32
}

func pr() *periph {
	return (*periph)(unsafe.Pointer(mmap.IOMUXC_BASE))
}
