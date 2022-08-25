// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/rtos"
	"strings"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/internal"
)

type Driver struct {
	timeoutRx time.Duration
	timeoutTx time.Duration
	p         *Periph
	rxDMA     dma.Channel
	txDMA     dma.Channel
	rxReady   rtos.Note
	txDone    rtos.Note
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph, rxdma, txdma dma.Channel) *Driver {
	return &Driver{
		timeoutRx: -1,
		timeoutTx: -1,
		p:         p,
		rxDMA:     rxdma,
		txDMA:     txdma,
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
	d.p.CTRL.StoreBits(ctrlMask, CTRL(conf))
}

// Setup enables clock source, resets, and configures the LPUART peripheral. You
// still have to enable Tx and/or Rx before use it.
func (d *Driver) Setup(conf Config, baudrate int) {
	d.p.EnableClock(true)
	d.p.GLOBAL.Store(RST)
	d.p.GLOBAL.Store(0)
	d.p.SetBaudrate(baudrate)
	d.p.WATER.Store(1 << TXWATERn)
	d.p.FIFO.Store(RXFE | TXFE)
	d.SetConfig(conf)
}

func (d *Driver) EnableRx() {
	// TODO: DMA, ...
	d.p.CTRL.SetBits(RE | DOZEEN)
}

func (d *Driver) DisableRx() {
	// TODO: ...
	d.p.CTRL.ClearBits(RE | DOZEEN)
}

// EnableTx enables Tx part of the LPUART peripheral and setups Tx DMA channel.
func (d *Driver) EnableTx() {
	d.p.CTRL.SetBits(TE | DOZEEN)
	// TODO: DMA
}

// DisableTx disables Tx part of the USART peripheral.
func (d *Driver) DisableTx() {
	// TODO: wait for transfer complete (empty FIFO or maybe TC)
	d.p.CTRL.ClearBits(TE | DOZEEN)
}

type Error struct {
	Rx STAT
}

func (e Error) Error() string {
	var (
		a [4]string
		n int
	)
	if e.Rx&PF != 0 {
		a[n] = "parity"
		n++
	}
	if e.Rx&FE != 0 {
		a[n] = "framing"
		n++
	}
	if e.Rx&NF != 0 {
		a[n] = "noise"
		n++
	}
	if e.Rx&OR != 0 {
		a[n] = "overrun"
		n++
	}
	return "LPUART Rx " + strings.Join(a[:n], ",")
}

func (d *Driver) Read(buf []byte) (n int, err error) {
	if len(buf) == 0 {
		return
	}
	stat := d.p.STAT.Load()
	if stat&(RDRF|OR) == 0 {
		d.rxReady.Clear()
		if d.rxReady.Sleep(d.timeoutRx) {
			// todo
		}
		stat = d.p.STAT.Load()
	}
	if !d.rxDMA.IsValid() {
		for n < len(buf) {
			data := d.p.DATA.Load()
			if data&(RXEMPT|FRETSC|PARITYE|NOISY) != 0 {
				var e STAT
				if data&(FRETSC|PARITYE|NOISY) != 0 {
					e = STAT(data&FRETSC)<<(FEn-FRETSCn) |
						STAT(data&PARITYE)<<(PFn-PARITYEn) |
						STAT(data&NOISY)<<(NFn-NOISYn)
				}
				if data&RXEMPT != 0 && stat&OR != 0 {
					// report and clear OR only if no more data in FIFO
					e |= OR
					d.p.STAT.Store(OR)
				}
				return
			}
			buf[n] = byte(data)
			n++
		}
		return
	}
	return
}

func (d *Driver) ISR() {
	ctrl := d.p.CTRL.Load()
	stat := d.p.STAT.Load()
	if ctrl&(RIE|ORIE) != 0 && stat&(RDRF|OR) != 0 {
		internal.AtomicStoreBits(&d.p.CTRL, RIE|ORIE, 0)
		d.rxReady.Wakeup()
	}
	if ctrl&TIE != 0 && stat&TDRE != 0 {
		internal.AtomicStoreBits(&d.p.CTRL, TIE, 0)
		d.txDone.Wakeup()
	}
}
