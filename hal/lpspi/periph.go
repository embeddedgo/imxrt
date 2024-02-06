// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/hal/internal/ccm"
	"github.com/embeddedgo/imxrt/p/mmap"
)

type Periph struct {
	VERID mmio.R32[uint32]
	PARAM mmio.R32[uint32]
	_     [2]uint32
	CR    mmio.R32[CR]
	SR    mmio.R32[SR]
	IER   mmio.R32[IER]
	DER   mmio.R32[DER]
	CFGR0 mmio.R32[CFGR0]
	CFGR1 mmio.R32[CFGR1]
	_     [2]uint32
	DMR0  mmio.R32[uint32]
	DMR1  mmio.R32[uint32]
	_     [2]uint32
	CCR   mmio.R32[CCR]
	_     [5]uint32
	FCR   mmio.R32[FCR]
	FSR   mmio.R32[FSR]
	TCR   mmio.R32[TCR]
	TDR   mmio.R32[uint32]
	_     [2]uint32
	RSR   mmio.R32[RSR]
	RDR   mmio.R32[uint32]
}

func LPSPI(n int) *Periph {
	if n--; uint(n) > 3 {
		panic("wrong LPSPI number")
	}
	const base = mmap.LPSPI1_BASE
	const step = mmap.LPSPI2_BASE - mmap.LPSPI1_BASE
	return (*Periph)(unsafe.Pointer(base + uintptr(n)*step))
}

func num(p *Periph) int {
	const step = mmap.LPSPI2_BASE - mmap.LPSPI1_BASE
	return int((uintptr(unsafe.Pointer(p)) - mmap.LPSPI1_BASE) / step)
}

// EnableClock enables the clock for the LPSPI peripheral.
// lp determines whether the clock remains on in low power WAIT mode.
func (p *Periph) EnableClock(lp bool) {
	if n := num(p); uint(n) < 4 {
		ccm.CCGR(1).SetCG(n, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
	}
}

// DisableClock disables the clock for the LPSPI peripheral.
func (p *Periph) DisableClock() {
	if n := num(p); uint(n) < 4 {
		ccm.CCGR(1).SetCG(n, 0)
	}
}

func (p *Periph) Reset() {
       p.CR.Store(RRF | RTF | RST)
       p.CR.Store(0)
}
