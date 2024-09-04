// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/hal/internal/ccm"
	"github.com/embeddedgo/imxrt/p/mmap"
)

type Periph struct {
	VERID  mmio.R32[uint32]
	PARAM  mmio.R32[uint32]
	_      [2]uint32
	MCR    mmio.R32[MCR]
	MSR    mmio.R32[MSR]
	MIER   mmio.R32[MSR]
	MDER   mmio.R32[DER]
	MCFGR0 mmio.R32[MCFGR0]
	MCFGR1 mmio.R32[MCFGR1]
	MCFGR2 mmio.R32[MCFGR2]
	MCFGR3 mmio.R32[MCFGR3]
	_      [4]uint32
	MDMR   mmio.R32[MDMR]
	_      uint32
	MCCR0  mmio.R32[MCCR]
	_      uint32
	MCCR1  mmio.R32[MCCR]
	_      uint32
	MFCR   mmio.R32[MFCR]
	MFSR   mmio.R32[MFSR]
	MTDR   mmio.R32[MTDR]
	_      [3]uint32
	MRDR   mmio.R32[RDR]
	_      [39]uint32
	SCR    mmio.R32[SCR]
	SSR    mmio.R32[SSR]
	SIER   mmio.R32[SSR]
	SDER   mmio.R32[DER]
	_      uint32
	SCFGR1 mmio.R32[SCFGR1]
	SCFGR2 mmio.R32[SCFGR2]
	_      [5]uint32
	SAMR   mmio.R32[SAMR]
	_      [3]uint32
	SASR   mmio.R32[SASR]
	STAR   mmio.R32[STAR]
	_      [2]uint32
	STDR   mmio.R32[uint32]
	_      [3]uint32
	SRDR   mmio.R32[RDR]
}

// LPI2C returns the LPI2Cn peripheral.
func LPI2C(n int) *Periph {
	if n--; uint(n) > 2 {
		panic("wrong LPI2C number")
	}
	const base = mmap.LPI2C1_BASE
	const step = mmap.LPI2C2_BASE - mmap.LPI2C1_BASE
	return (*Periph)(unsafe.Pointer(base + uintptr(n)*step))
}

func num(p *Periph) int {
	const step = mmap.LPI2C2_BASE - mmap.LPI2C1_BASE
	return int((uintptr(unsafe.Pointer(p)) - mmap.LPI2C1_BASE) / step)
}

// EnableClock enables the clock for the LPI2C peripheral.
// lp determines whether the clock remains on in low power WAIT mode.
func (p *Periph) EnableClock(lp bool) {
	if n := num(p); uint(n) <= 2 {
		ccm.CCGR(2).SetCG(3+n, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
	}
}
