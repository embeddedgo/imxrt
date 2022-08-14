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

type VERID uint32
type PARAM uint32
type GLOBAL uint32
type PINCFG uint32
type BAUD uint32
type STAT uint32
type CTRL uint32
type DATA uint32
type MATCH uint32
type MODIR uint32
type FIFO uint32
type WATER uint32

type R32[T ~uint32] struct{ U32 mmio.U32 }

func (r *R32[T]) LoadBits(mask T) T   { return T(r.U32.LoadBits(uint32(mask))) }
func (r *R32[T]) StoreBits(mask, b T) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *R32[T]) SetBits(mask T)      { r.U32.SetBits(uint32(mask)) }
func (r *R32[T]) ClearBits(mask T)    { r.U32.ClearBits(uint32(mask)) }
func (r *R32[T]) Load() T             { return T(r.U32.Load()) }
func (r *R32[T]) Store(b T)           { r.U32.Store(uint32(b)) }

type Periph struct {
	VERID  R32[VERID]
	PARAM  R32[PARAM]
	GLOBAL R32[GLOBAL]
	PINCFG R32[PINCFG]
	BAUD   R32[BAUD]
	STAT   R32[STAT]
	CTRL   R32[CTRL]
	DATA   R32[DATA]
	MATCH  R32[MATCH]
	MODIR  R32[MODIR]
	FIFO   R32[FIFO]
	WATER  R32[WATER]
}

func LPUART(n int) *Periph {
	if n--; uint(n) > 7 {
		panic("wrong LPUART number")
	}
	base := mmap.LPUART1_BASE
	step := mmap.LPUART2_BASE - mmap.LPUART1_BASE
	return (*Periph)(unsafe.Pointer(base + uintptr(n)*step))
}

func cg(p *Periph) (*ccm.CCGR_, int) {
	switch uintptr(unsafe.Pointer(p)) {
	case mmap.LPUART1_BASE:
		return ccm.CCGR(5), 12
	case mmap.LPUART2_BASE:
		return ccm.CCGR(0), 14
	case mmap.LPUART3_BASE:
		return ccm.CCGR(0), 6
	case mmap.LPUART4_BASE:
		return ccm.CCGR(1), 12
	case mmap.LPUART5_BASE:
		return ccm.CCGR(3), 1
	case mmap.LPUART6_BASE:
		return ccm.CCGR(3), 3
	case mmap.LPUART7_BASE:
		return ccm.CCGR(5), 13
	case mmap.LPUART8_BASE:
		return ccm.CCGR(6), 7
	}
	return nil, 0
}

// EnableClock enables clock for LPUART peripheral.
// lp determines whether the clock remains on in low power WAIT mode.
func (p *Periph) EnableClock(lp bool) {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
	}
}

// DisableClock disables clock for LPUART peripheral.
func (p *Periph) DisableClock() {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, 0)
	}
}

// Reset resets all internal logic and registers of LPUART peripheral.
func (p *Periph) Reset() {
	p.GLOBAL.Store(2)
	p.GLOBAL.Store(0)
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

/*
// SetBaudrate sets the UART speed [sym/s].
func (p *Periph) SetBaudrate(baud int) {
	const uartClkRoot = 80e6 // TODO: calculate from CCM settings
	osr, sbr := dividers(uartClkRoot, baud)
	var bc lpuart.BAUD
	if osr < 8 {
		bc = lpuart.BOTHEDGE

	}
	bc |= lpuart.BAUD((osr-1)<<lpuart.OSRn | sbr)
	p.BAUD.Store(baud)
}
*/