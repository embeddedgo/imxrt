// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"github.com/embeddedgo/imxrt/hal/dma"
)

type Master struct {
	p     *Periph
	rxdma dma.Channel
	txdma dma.Channel
}

// NewMaster returns a new master-mode driver for p. If valid DMA channels are
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, rxdma, txdma dma.Channel) *Master {
	return &Master{p: p, rxdma: rxdma, txdma: txdma}
}

// Periph returns the underlying LPSPI peripheral.
func (d *Master) Periph() *Periph {
	return d.p
}

// See Table 47-5. LPI2C Example Timing Configurations
const (
	clk  = 60_000_000 // peripheral clock (PLL_USB1 / 8)
	pre8 = 3
	sl   = 0xf<<30 | 0x20<<SETHOLDn | 0x3f<<CLKLOn | 0x3f<<CLKHIn | 0x10<<DATAVDn
	pre4 = 2 // divides by 4 the 60 MHz clock
	st   = 0xf<<30 | 0x20<<SETHOLDn | 0x3f<<CLKLOn | 0x3f<<CLKHIn | 0x10<<DATAVDn
	fa   = 0x1<<30 | 0x08<<SETHOLDn | 0x14<<CLKLOn | 0x10<<CLKHIn | 0x04<<DATAVDn
	pl   = 0x1<<30 | 0x03<<SETHOLDn | 0x08<<CLKLOn | 0x05<<CLKHIn | 0x01<<DATAVDn
	pre2 = 1 // divides by 2 the 60 MHz clock
	fahs = 0x2<<30 | 0x11<<SETHOLDn | 0x28<<CLKLOn | 0x1f<<CLKHIn | 0x08<<DATAVDn
	plhs = 0x2<<30 | 0x07<<SETHOLDn | 0x0f<<CLKLOn | 0x0b<<CLKHIn | 0x02<<DATAVDn
	hs   = 0x0<<30 | 0x04<<SETHOLDn | 0x04<<CLKLOn | 0x02<<CLKHIn | 0x01<<DATAVDn

	setupSlow   = pre8<<6 | hs<<34 | sl
	setupStd    = pre4<<6 | hs<<34 | st
	setupFast   = pre4<<6 | hs<<34 | fa
	setupPlus   = pre4<<6 | hs<<34 | pl
	setupFastHS = pre2<<6 | hs<<34 | fahs
	setupPlusHS = pre2<<6 | hs<<34 | plhs
)

// Mode of operation (argument for Master.Setup method).
const (
	Slow       uint64 = setupSlow
	Std        uint64 = setupStd    // 116 kb/s (Std)   and 1.7 Mb/s HS
	Fast       uint64 = setupFast   // 400 kb/s (Fast)  and 1.7 Mb/s HS
	FastPlus   uint64 = setupPlus   //   1 Mb/s (Fast+) and 1.7 Mb/s HS
	FastHS     uint64 = setupFastHS // 400 kb/s (Fast)  and 3.3 Mb/s HS
	FastPlusHS uint64 = setupPlusHS //   1 Mb/s (Fast+) and 3.3 Mb/s HS
)

func (d *Master) Setup(mode uint64) {
	p := d.p
	p.EnableClock(true)
	p.MCR.Store(MRST)
	p.MCR.Store(0)
	p.MCCR0.Store(MCCR(mode) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR1.Store(MCCR(mode>>34) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	pre := uint(mode) >> 6 & 3 // max. supported MPRESCALE is 3
	p.MCFGR1.Store(MCFGR1(pre) << MPRESCALEn)
	gf := MCFGR2(mode>>30) & 0xf // max. supported MFILT is 15
	p.MCFGR2.Store(gf<<MFILTSDAn | gf<<MFILTSCLn)
	const timeout = clk * 15 / 1000 // number of peripheral clock cycles that equals 15 ms
	p.MCFGR3.Store(timeout / 256 >> pre << PINLOWn)
	p.MFCR.Store(3 << TXWATERn)
	p.MCR.Store(MEN)
}
