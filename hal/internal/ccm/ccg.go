// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ccm

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/p/mmap"
)

const (
	ClkEn int8 = 1 << 0 // enable clock
	ClkLP int8 = 1 << 1 // keep clock enabled in low-power WAIT mode
)

type CCGR_ struct {
	U32 mmio.U32
}

func (r *CCGR_) CG(i int) int8 {
	return int8(r.U32.Load() >> uint(i*2) & 3)
}

func (r *CCGR_) SetCG(i int, cg int8) {
	i *= 2
	internal.AtomicStoreBits(&r.U32, 3<<uint(i), uint32(cg)<<uint(i))
}

func CCGR(i int) *CCGR_ {
	return &(*[8]CCGR_)(unsafe.Pointer(mmap.CCM_BASE + 0x068))[i]
}
