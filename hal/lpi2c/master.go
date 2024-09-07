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
	div4 = 2          // divides by 4 the 60 MHz clock
	st   = 9<<30 | 32<<SETHOLDn | 63<<CLKLOn | 63<<CLKHIn | 16<<DATAVDn
	fa   = 2<<30 | 10<<SETHOLDn | 20<<CLKLOn | 16<<CLKHIn | 5<<DATAVDn
	pl   = 2<<30 | 4<<SETHOLDn | 8<<CLKLOn | 5<<CLKHIn | 1<<DATAVDn
	div2 = 1 // divides by 2 the 60 MHz clock
	fahs = 2<<30 | 0x11<<SETHOLDn | 0x28<<CLKLOn | 0x1f<<CLKHIn | 0x08<<DATAVDn
	plhs = 2<<30 | 0x07<<SETHOLDn | 0x0f<<CLKLOn | 0x0b<<CLKHIn | 0x02<<DATAVDn
	hs   = 0<<30 | 0x04<<SETHOLDn | 0x04<<CLKLOn | 0x02<<CLKHIn | 0x01<<DATAVDn

	setupStd    = div4<<6 | hs<<34 | st
	setupFast   = div4<<6 | hs<<34 | fa
	setupPlus   = div4<<6 | hs<<34 | pl
	setupFastHS = div2<<6 | hs<<34 | fahs
	setupPlusHS = div2<<6 | hs<<34 | plhs
)

// Mode of operation (argument for Master.Setup method).
const (
	Std100k    uint64 = setupStd    // ≤115 kb/s (Std)   and 1.7 Mb/s HS
	Fast400k   uint64 = setupFast   // ≤400 kb/s (Fast)  and 1.7 Mb/s HS
	FastPlus1M uint64 = setupPlus   //   ≤1 Mb/s (Fast+) and 1.7 Mb/s HS
	FastHS     uint64 = setupFastHS // ≤400 kb/s (Fast)  and 3.3 Mb/s HS
	FastPlusHS uint64 = setupPlusHS //   ≤1 Mb/s (Fast+) and 3.3 Mb/s HS
)

func (d *Master) Setup(mode uint64) {
	p := d.p
	p.EnableClock(true)
	p.MCR.Store(MRST)
	p.MCR.Store(0)
	p.MCCR0.Store(MCCR(mode) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR1.Store(MCCR(mode>>34) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	pre := MCFGR1(mode) >> 6 & 3 // max. supported MPRESCALE is 3
	p.MCFGR1.Store(pre << MPRESCALEn)
	gf := MCFGR2(mode>>30) & 0xf // max. supported MFILT is 15
	bi := (MCFGR2(mode)>>CLKLOn&63 + MCFGR2(mode)>>SETHOLDn&63 + 2) * 2
	p.MCFGR2.Store(gf<<MFILTSDAn | gf<<MFILTSCLn | bi<<MBUSIDLEn)
	const timeout = clk * 15 / 1000 // number of clock cycles that equals 15 ms
	p.MCFGR3.Store(timeout / 256 >> pre << PINLOWn)
	p.MFCR.Store(3 << TXWATERn)
	p.MCR.Store(MEN)
}
