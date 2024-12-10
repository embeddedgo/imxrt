// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ccm provides simple interface to the CCGR registrs. As a side effect
// it also access to the peripheral registers in user mode.
package ccm

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/p/mmap"

	_ "github.com/embeddedgo/imxrt/hal/internal/umreg" // user mode peripherals
)

// We assume that every HAL package imports this package for clock management
// so it enables user mode peripherals

const (
	ClkEn int8 = 1 << 0 // enable clock
	ClkLP int8 = 1 << 1 // keep clock enabled in low-power WAIT mode
)

type CCGR_ struct {
	R32 mmio.R32[uint32]
}

func (r *CCGR_) CG(i int) int8 {
	return int8(r.R32.Load() >> uint(i*2) & 3)
}

func (r *CCGR_) SetCG(i int, cg int8) {
	i *= 2
	internal.ExclusiveStoreBits(&r.R32, 3<<uint(i), uint32(cg)<<uint(i))
}

func CCGR(i int) *CCGR_ {
	return &(*[8]CCGR_)(unsafe.Pointer(mmap.CCM_BASE + 0x068))[i]
}
