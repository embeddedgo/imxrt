// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"sync"

	"github.com/embeddedgo/imxrt/hal/dma"
)

// A Master is a driver for the LPI2C peripheral to perform a master access to
// an I2C bus.
type Master struct {
	sync.Mutex // useful to share driver, used by connection interface

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

// Timing constants.
//
// sclClk = clk / ((CLKHI + CLKLO + 2 + sclLatency) << divN)
//
// sclLatency = roundDown((2 + FILTSCL) >> divN)
const (
	clk  = 60_000_000 // peripheral clock (PLL_USB1 / 8)
	div2 = 1          // divide the 60 MHz clock by 2 (30 MHz)
	div4 = 2          // divide the 60 MHz clock by 4 (15 MHz)
	div8 = 3          // divide the 60 MHz clock by 8 (7.5 MHz)

	// Values copied from Table 47-5. LPI2C Example Timing Configurations.
	fahs = 2<<30 | 17<<SETHOLDn | 40<<CLKLOn | 31<<CLKHIn | 8<<DATAVDn
	plhs = 2<<30 | 7<<SETHOLDn | 15<<CLKLOn | 11<<CLKHIn | 2<<DATAVDn
	hs   = 4<<SETHOLDn | 4<<CLKLOn | 2<<CLKHIn | 1<<DATAVDn

	// The above values divided by 2 with small corrections to work with div4.
	fa = 2<<30 | 9<<SETHOLDn | 20<<CLKLOn | 16<<CLKHIn | 4<<DATAVDn
	pl = 2<<30 | 4<<SETHOLDn | 8<<CLKLOn | 5<<CLKHIn | 1<<DATAVDn

	// Values to obtain the minimal possible sclClk for any div.
	sl = 15<<30 | 31<<SETHOLDn | 63<<CLKLOn | 63<<CLKHIn | 15<<DATAVDn

	timingSlow   = div8<<6 | hs<<34 | sl
	timingStd    = div4<<6 | hs<<34 | sl
	timingFast   = div4<<6 | hs<<34 | fa
	timingPlus   = div4<<6 | hs<<34 | pl
	timingFastHS = div2<<6 | hs<<34 | fahs
	timingPlusHS = div2<<6 | hs<<34 | plhs

	stuckBusTimeout = 40 // ms (TI "I2C Stuck Bus: Prevention and Workarounds")
)

// Speed encodes the timing configuration that determines the maximum
// communication speed (the actual speed depends also on the SCL rise time).
type Speed uint64

const (
	Slow50k    Speed = timingSlow   //  ≤58 kb/s (Slow)  and 0.83 Mb/s HS
	Std100k    Speed = timingStd    // ≤114 kb/s (Std)   and 1.65 Mb/s HS
	Fast400k   Speed = timingFast   // ≤400 kb/s (Fast)  and 1.65 Mb/s HS
	FastPlus1M Speed = timingPlus   //   ≤1 Mb/s (Fast+) and 1.65 Mb/s HS
	FastHS     Speed = timingFastHS // ≤400 kb/s (Fast)  and 3.33 Mb/s HS
	FastPlusHS Speed = timingPlusHS //   ≤1 Mb/s (Fast+) and 3.33 Mb/s HS
)

func (d *Master) Setup(sp Speed) {
	p := d.p
	p.EnableClock(true)
	p.MCR.Store(MRST)
	p.MCR.Store(0)
	p.MCCR0.Store(MCCR(sp) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR1.Store(MCCR(sp>>34) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	pre := MCFGR1(sp) >> 6 & 3 // max. supported MPRESCALE is 3
	p.MCFGR1.Store(pre << MPRESCALEn)
	gf := MCFGR2(sp>>30) & 0xf // the used encoding supports MFILT <= 15
	bi := (MCFGR2(sp)>>CLKLOn&63 + MCFGR2(sp)>>SETHOLDn&63 + 2) * 2
	p.MCFGR2.Store(gf<<MFILTSDAn | gf<<MFILTSCLn | bi<<MBUSIDLEn)
	p.MCFGR3.Store(clk * stuckBusTimeout / 1000 / 256 >> pre << PINLOWn)
	p.MFCR.Store(3 << TXWATERn)
	p.MCR.Store(MEN)
}
