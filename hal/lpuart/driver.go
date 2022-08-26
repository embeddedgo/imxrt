// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/rtos"
	"strings"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
)

type Error uint32

const (
	EPARITY  = Error(PF)
	EFRAMING = Error(FE)
	ENOISE   = Error(NF)
	EOVERRUN = Error(OR)
)

func (e Error) Error() string {
	var (
		a [4]string
		n int
	)
	if e&EPARITY != 0 {
		a[n] = "parity"
		n++
	}
	if e&EFRAMING != 0 {
		a[n] = "framing"
		n++
	}
	if e&ENOISE != 0 {
		a[n] = "noise"
		n++
	}
	if e&EOVERRUN != 0 {
		a[n] = "overrun"
		n++
	}
	return "lpurat: " + strings.Join(a[:n], ",")
}

type DriverError uint8

const (
	// ErrBufOverflow is returned if one or more received bytes has been dropped
	// because of the lack of free space in the driver's receive buffer.
	ErrBufOverflow DriverError = iota + 1

	// ErrTimeout is returned if timeout occured. It means that the read/write
	// operation has been interrupted. In case of write you can not determine
	// the exact number of bytes sent to the remote party.
	ErrTimeout
)

// Error implements error interface.
func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "lpuart: buffer overflow"
	case ErrTimeout:
		return "lpuart: timeout"
	}
	return ""
}

type Driver struct {
	p *Periph

	// Rx fields
	rxtimeout time.Duration
	rxdma     dma.Channel
	rxready   rtos.Note
	rxbuf     []DATA // Rx ring buffer
	nextr     uint32
	nextw     uint32
	rxwake    uint32
	overflow  uint32

	// Tx fields
	txtimeout time.Duration
	txdma     dma.Channel
	txdata    string
	txn       int
	txmax     int
	txdone    rtos.Note
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph, rxdma, txdma dma.Channel) *Driver {
	return &Driver{
		p:         p,
		rxtimeout: -1,
		rxdma:     rxdma,
		txtimeout: -1,
		txdma:     txdma,
	}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

type Config uint32

const (
	Word7b  = Config(M7)
	Word8b  = Config(0)
	Word9b  = Config(M)
	Word10b = Config(M10)
	ParEven = Config(PE)
	ParOdd  = Config(PE | PT)

	Stop2b = Config(SBNS)
)

// SetConfig configures LPUART.
func (d *Driver) SetConfig(conf Config) {
	const (
		baudMask = M10
		ctrlMask = M7 | M | PE | PT
		_        = -(uint(baudMask) & uint(ctrlMask)) // check colliding bits
	)
	d.p.BAUD.StoreBits(baudMask, BAUD(conf))
	d.p.CTRL.StoreBits(ctrlMask, CTRL(conf)|DOZEEN)
}

// Setup enables clock source, resets, and configures the LPUART peripheral. You
// still have to enable Tx and/or Rx before use it.
func (d *Driver) Setup(conf Config, baudrate int) {
	d.p.EnableClock(true)
	d.p.GLOBAL.Store(RST)
	d.p.GLOBAL.Store(0)
	d.p.SetBaudrate(baudrate)
	d.p.WATER.Store(0)
	d.p.FIFO.Store(RXFE | TXFE)
	d.SetConfig(conf)
	d.txmax = 1 << (d.p.PARAM.LoadBits(TXFIFO) >> TXFIFOn)
}

func (d *Driver) ISR() {
	ctrl := d.p.CTRL.Load()
	stat := d.p.STAT.Load()
	if ctrl&RIE != 0 && stat&RDRF != 0 {
		if d.rxdma.IsValid() {
			// TODO:
		} else {
			readNoDMA(d)
		}
	}
	if ctrl&TIE != 0 && stat&TDRE != 0 {
		if d.rxdma.IsValid() {
			// TODO:
		} else {
			writeNoDMA(d)
		}
	}
}
