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
	VERID  mmio.R32[uint32] // Version ID Register
	PARAM  mmio.R32[uint32] // Parameter Register
	_      [2]uint32
	MCR    mmio.R32[MCR]    // Master Control Register
	MSR    mmio.R32[MSR]    // Master Status Register
	MIER   mmio.R32[MSR]    // Master Interrupt Enable Register
	MDER   mmio.R32[DER]    // Master DMA Enable Register
	MCFGR0 mmio.R32[MCFGR0] // Master Configuration Register 0
	MCFGR1 mmio.R32[MCFGR1] // Master Configuration Register 1
	MCFGR2 mmio.R32[MCFGR2] // Master Configuration Register 2
	MCFGR3 mmio.R32[MCFGR3] // Master Configuration Register 3
	_      [4]uint32
	MDMR   mmio.R32[MDMR] // Master Data Match Register
	_      uint32
	MCCR0  mmio.R32[MCCR] // Master Clock Configuration Register 0
	_      uint32
	MCCR1  mmio.R32[MCCR] // Master Clock Configuration Register 1
	_      uint32
	MFCR   mmio.R32[MFCR] // Master FIFO Control Register
	MFSR   mmio.R32[MFSR] // Master FIFO Status Register
	MTDR   mmio.R32[MTDR] // Master Transmit Data Register
	_      [3]uint32
	MRDR   mmio.R32[RDR] // Master Receive Data Register
	_      [39]uint32
	SCR    mmio.R32[SCR] // Slave Control Register
	SSR    mmio.R32[SSR] // Slave Status Register
	SIER   mmio.R32[SSR] // Slave Interrupt Enable Register
	SDER   mmio.R32[DER] // Slave DMA Enable Register
	_      uint32
	SCFGR1 mmio.R32[SCFGR1] // Slave Configuration Register 1
	SCFGR2 mmio.R32[SCFGR2] // Slave Configuration Register 2
	_      [5]uint32
	SAMR   mmio.R32[SAMR] // Slave Address Match Register
	_      [3]uint32
	SASR   mmio.R32[SASR] // Slave Address Status Register
	STAR   mmio.R32[STAR] // Slave Transmit ACK Register
	_      [2]uint32
	STDR   mmio.R32[uint32] // Slave Transmit Data Register
	_      [3]uint32
	SRDR   mmio.R32[RDR] // Slave Receive Data Register
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
