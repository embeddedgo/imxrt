// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/rtos"

	"github.com/embeddedgo/imxrt/hal/dma"
)

type Master struct {
	p     *Periph
	rxdma dma.Channel
	txdma dma.Channel
	done  rtos.Note
}

// NewMaster returns a new master-mode driver for p.
func NewMaster(p *Periph, rxdma, txdma dma.Channel) *Master {
	return &Master{p: p, rxdma: rxdma, txdma: txdma}
}

// Periph returns the underlying LPSPI peripheral.
func (m *Master) Periph() *Periph {
	return m.p
}

// Enable enables LPSPI peripheral.
func (m *Master) Enable() {
	m.p.CR.Store(DBGEN | MEN)
}

// Disable disables LPSPI peripheral.
func (m *Master) Disable() {
	m.p.CR.Store(0)
}

// TODO: calculate from CCM settings
const lpspiClkRoot = 132.923e6 // 480e6 * 18 / 13 / 5

// Setup enables the SPI clock, resets the peripheral and sets its basic
// configuration and the base SCK clock frequency. This functions allows you to
// set the baseFreq up to 66.5 MHz. Use Cmd to fine tune the configuration and
// set the SCK prescaler (don't exceed the maximum supported SCK clock: 30 MHz).
func (m *Master) Setup(conf CFGR1, baseFreqHz int) {
	p := m.p
	p.EnableClock(true)
	p.CR.Store(RRF | RTF | RST)
	p.CR.Store(0)
	p.CFGR1.Store(conf | MASTER)
	switch {
	case baseFreqHz > lpspiClkRoot/2:
		baseFreqHz = lpspiClkRoot / 2
	case baseFreqHz <= 0:
		baseFreqHz = 1
	}
	sckdiv := lpspiClkRoot/baseFreqHz - 2
	if sckdiv > 255 {
		sckdiv = 255
	}
	p.CCR.Store(CCR(sckdiv))
	//n := 1<<p.PARAM.LoadBits(TXFIFO)
	p.FCR.Store(3)
}

func (m *Master) ISR() {
	//d.p.Event(END).DisableIRQ()
	m.done.Wakeup()
}

func (m *Master) WriteRead(out, in []byte) int {
	return 0
}
