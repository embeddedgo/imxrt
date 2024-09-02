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
	pre = 1 // for 60 MHz clock (PLL_USB1 / 8)
	fa  = 0x2<<6 | 0x11<<SETHOLDn | 0x28<<CLKLOn | 0x1f<<CLKHIn | 0x08<<DATAVDn
	pl  = 0x2<<6 | 0x07<<SETHOLDn | 0x0f<<CLKLOn | 0x0b<<CLKHIn | 0x01<<DATAVDn
	hs  = 0x0<<6 | 0x04<<SETHOLDn | 0x04<<CLKLOn | 0x02<<CLKHIn | 0x01<<DATAVDn

	setupFast = pre<<60 | hs<<30 | fa
	setupPlus = pre<<60 | hs<<30 | pl
)

// Mode of operation for Master.Setup method.
const (
	Fast     uint64 = setupFast // 400 kb/s Fast and 3.33 Mb/s HS
	FastPlus uint64 = setupPlus //  1 Mb/s Fast+ and 3.33 Mb/s HS
)

func (d *Master) Setup(mode uint64) {
	p := d.p
	p.EnableClock(true)
	p.MCCR[0].Store(MCCR(mode) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR[1].Store(MCCR(mode>>30) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCFGR1.Store(MCFGR1(mode >> 60))
	gf := mode >> 6 & 3 // cfg supports only up to 3 cycles glitch filter
	p.MCFGR2.Store(MCFGR2(gf<<24 | gf<<16))
}
