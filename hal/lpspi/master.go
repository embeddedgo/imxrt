// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

type Master struct {
	p      *Periph
	rxdma  dma.Channel
	txdma  dma.Channel
	rxdone rtos.Note
	txdone rtos.Note
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

const (
	clkRoot = 133e6 // 480e6 * 18 / 13 / 5,  TODO: calculate from CCM
	fifoLen = 16    // TODO: calculate from PARAM
)

// Setup enables the SPI clock, resets the peripheral and sets its basic
// configuration and the base SCK clock frequency. The base SPI clock frequency
// is set to baseFreq rounded down to 133 MHz divided by the number from 2 to
// 257. Use WriteCmd to fine tune the configuration and set the SCK prescaler
// to obtain the desired SPI clock frequency (30 MHz max).
func (m *Master) Setup(conf CFGR1, baseFreqHz int) {
	p := m.p
	p.EnableClock(true)
	p.Reset()
	p.CFGR1.Store(conf | MASTER)
	switch {
	case baseFreqHz > clkRoot/2:
		baseFreqHz = clkRoot / 2
	case baseFreqHz <= 0:
		baseFreqHz = 1
	}
	//sckdiv := clkRoot/baseFreqHz - 2 // natural way but rounds sckdiv down
	sckdiv := clkRoot/(baseFreqHz+1) - 1
	if sckdiv > 255 {
		sckdiv = 255
	}
	p.CCR.Store(CCR(sckdiv))
}

func (m *Master) ISR() {
	// ....
	//m.done.Wakeup()
}

// WriteCmd writes a command to the transmit FIFO. You can encode the frame size
// in cmd directly using the FRAMESZ field or specify it using the frameSize
// parameter (FRAMESZ = frameSize-1). The frame size is specified as a numer of
// bits. The minimum supported frame size is 8 bits and maximum is 4096 bits. If
// frameSize <= 32 it also specifies the word size. If frameSize > 32 then the
// word size is 32 except the last one wchich is equal to frameSize % 32 and
// must be >= 2 (e.g. frameSize = 33 is not supported). Be careful to use the
// correct WriteRead* function according to the configured word size.
func (m *Master) WriteCmd(cmd TCR, frameSize int) {
	m.p.TCR.Store(cmd | TCR(frameSize-1)&FRAMESZ)
}

func (m *Master) WriteWord(word uint32) {
	p := m.p
	for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
	}
	p.TDR.Store(word)
}

func (m *Master) ReadWord() uint32 {
	p := m.p
	for p.FSR.LoadBits(RXCOUNT) == 0 {
	}
	return p.RDR.Load()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// writeRead speed is crucial to achive fast bitrates (up to 30 MHz) so it uses
// unsafe pointers instead of slices to speedup things (smaller code size,
// no bound checking, only one increment operation in the loop).
func writeRead(p *Periph, out, in unsafe.Pointer, n int) {
	nr, nw := n, n
	nf := fifoLen // how many words can be written to TDR to don't overflow RDR
	for nw+nr != 0 {
		if nw != 0 {
			m := nf - int(p.FSR.LoadBits(TXCOUNT)>>TXCOUNTn)
			if m <= 0 {
				goto read
			}
			if m > nw {
				m = nw
			}
			nw -= m
			nf -= m
			for end := unsafe.Add(out, m); out != end; out = unsafe.Add(out, 1) {
				p.TDR.Store(uint32(*(*byte)(out)))
			}
		}
	read:
		if nr != 0 {
			m := int(p.FSR.LoadBits(RXCOUNT) >> RXCOUNTn)
			if m > nr {
				m = nr
			}
			nr -= m
			nf += m
			for end := unsafe.Add(in, m); in != end; in = unsafe.Add(in, 1) {
				*(*byte)(in) = byte(p.RDR.Load())
			}
		}
	}
	return
}

// WriteRead writes n = min(len(out), len(in)) words to the transmit FIFO and
// at the same time it reads the same number of words from the receive FIFO. The
// written words are zero-extended bytes from the out slice. The least
// significant bytes from the read words are saved in the in slice.
func (m *Master) WriteRead(out, in []byte) (n int) {
	n = min(len(out), len(in))
	if n <= dma.CacheLineSize*2 || !m.rxdma.IsValid() || !m.txdma.IsValid() {
		writeRead(
			m.p,
			unsafe.Pointer(unsafe.SliceData(out)),
			unsafe.Pointer(unsafe.SliceData(in)),
			n,
		)
		return
	}

	return
}

// WriteStringRead works like WriteRead.
func (m *Master) WriteStringRead(out string, in []byte) int {
	return m.WriteRead(unsafe.Slice(unsafe.StringData(out), len(out)), in)
}
