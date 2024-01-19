// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

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
	GLOBAL mmio.R32[GLOBAL]
	PINCFG mmio.R32[PINCFG]
	BAUD   mmio.R32[BAUD]
	STAT   mmio.R32[STAT]
	CTRL   mmio.R32[CTRL]
	DATA   mmio.R16[uint16]
	MATCH  mmio.R32[uint32]
	MODIR  mmio.R32[MODIR]
	FIFO   mmio.R32[FIFO]
	WATER  mmio.R32[uint32]
}

func LPUART(n int) *Periph {
	if n--; uint(n) > 7 {
		panic("wrong LPUART number")
	}
	const base = mmap.LPUART1_BASE
	const step = mmap.LPUART2_BASE - mmap.LPUART1_BASE
	return (*Periph)(unsafe.Pointer(base + uintptr(n)*step))
}

func num(p *Periph) int {
	const step = mmap.LPUART2_BASE - mmap.LPUART1_BASE
	return int((uintptr(unsafe.Pointer(p)) - mmap.LPUART1_BASE) / step)
}

var cgs = [...]uint8{
	5<<4 | 12,
	0<<4 | 14,
	0<<4 | 6,
	1<<4 | 12,
	3<<4 | 1,
	3<<4 | 3,
	5<<4 | 13,
	6<<4 | 7,
}

func cg(p *Periph) (*ccm.CCGR_, int) {
	if n := num(p); uint(n) < uint(len(cgs)) {
		cg := cgs[n]
		return ccm.CCGR(int(cg) >> 4), int(cg) & 15
	}
	return nil, 0
}

// EnableClock enables the clock for the LPUART peripheral.
// lp determines whether the clock remains on in low power WAIT mode.
func (p *Periph) EnableClock(lp bool) {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
	}
}

// DisableClock disables the clock for the LPUART peripheral.
func (p *Periph) DisableClock() {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, 0)
	}
}

func dividers(clk, baud int) (osr, sbr int) {
	lowestE := 1<<31 - 1
	for o := 32; o >= 4; o-- {
		bo := baud * o
		minS := clk / bo
		// check s = minS and s = minS + 1
		for s := minS; ; s++ {
			e := clk - bo*s
			if e < 0 {
				e = -e
			}
			if e < lowestE {
				lowestE = e
				osr = o
				sbr = s
				if e == 0 {
					return
				}
			}
			if s != minS {
				break
			}
		}
	}
	return
}

// SetBaudrate sets the UART speed [sym/s].
func (p *Periph) SetBaudrate(baud int) {
	const uartClkRoot = 80e6 // TODO: calculate from CCM settings
	osr, sbr := dividers(uartClkRoot, baud)
	var baudBits BAUD
	if osr < 8 {
		baudBits = BOTHEDGE
	}
	baudBits |= BAUD((osr-1)<<OSRn | sbr)
	p.BAUD.StoreBits(BOTHEDGE|OSR|SBR, baudBits)
}
