// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/rtos"
	"strings"
	"time"
	"unsafe"

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
		return "lpuart: Rx buffer overflow"
	case ErrTimeout:
		return "lpuart: timeout"
	}
	return ""
}

// A Driver is a driver to the LPUART peripheral. It provides standard io.Reader
// and io.Writer interface that can be used to read/write stream of 8-bit
// characters. It also provides couple of methods to configure and manage the
// underlying LPURAR peripheral for typical applications and to handle stream of
// 9 and 10 bit characters. For more complex scenarios you can directly access
// all LPUART registers.
//
// The receiver, if enabled, continuously writes received data to the internal
// ring buffer which minimizes the risk of data loss. All provided reading
// methods read from this buffer. The detection of a buffer overflow is based on
// a best-effort strategy. You cannot rely on it, nor can you rely on error
// detection provided by hardware. Both provide qualitative, not quantitative,
// information.
//
// The sending and receiving subsystems of the driver are completly independent.
// Each of them can be independently turned on, off, and used by different
// goroutines. Each one can be also configured in DMA or no-DMA mode.
type Driver struct {
	p *Periph

	// Rx fields
	rxtimeout time.Duration
	rxdma     dma.Channel
	rxready   rtos.Note
	rxbuf     []uint16 // Rx ring buffer
	nextr     uint32   // 30 LSBits: index in rxbuf, 2 MSBits: loop count
	nextw     uint32   // 30 LSBits: index in rxbuf, 2 MSBits: loop count
	rxwake    uint32
	rxdman    uint32
	rxfirst   uint16

	// Tx fields
	txtimeout time.Duration
	txdma     dma.Channel
	txd       unsafe.Pointer
	txi       int
	txn       int
	txlog2max uint
	txdone    rtos.Note
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph, rxdma, txdma dma.Channel) *Driver {
	return &Driver{
		p:         p,
		rxtimeout: -1,
		rxdma:     rxdma,
		rxfirst:   RXEMPT,
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
	p := d.p
	p.BAUD.StoreBits(baudMask, BAUD(conf))
	p.CTRL.StoreBits(ctrlMask, CTRL(conf))
}

// Setup enables clock source, resets, and configures the LPUART peripheral. You
// still have to enable Tx and/or Rx before use it.
func (d *Driver) Setup(conf Config, baudrate int) {
	p := d.p
	p.EnableClock(true)
	p.GLOBAL.Store(RST) // reset
	p.GLOBAL.Store(0)
	p.WATER.Store(0)
	p.FIFO.Store(RXFE | TXFE)
	var dmae BAUD
	if rxdma := d.rxdma; rxdma.IsValid() {
		rxdma.DisableReq()
		rxdma.DisableErrInt()
		rxdma.ClearInt()
		dmae = RDMAE
	}
	if txdma := d.txdma; txdma.IsValid() {
		txdma.DisableReq()
		txdma.DisableErrInt()
		txdma.ClearInt()
		dmae |= TDMAE
	}
	if dmae != 0 {
		// Enable DMA requests. Gates IRQ (undocumented?).
		p.BAUD.StoreBits(TDMAE|RDMAE, dmae)
	}
	p.SetBaudrate(baudrate)
	d.SetConfig(conf)
	d.txlog2max = uint(d.p.PARAM.LoadBits(TXFIFO) >> TXFIFOn)
}

// SetReadTimeout sets the read timeout used by Read* functions.
func (d *Driver) SetReadTimeout(timeout time.Duration) {
	d.rxtimeout = timeout
}

// SetWriteTimeout sets the write timeout used by Write* functions.
func (d *Driver) SetWriteTimeout(timeout time.Duration) {
	d.txtimeout = timeout
}

//go:nosplit
//go:nowritebarrierrec
func (d *Driver) ISR() {
	p := d.p
	ctrl := p.CTRL.Load()
	stat := p.STAT.Load()
	if ctrl&RIE != 0 && stat&RDRF != 0 {
		rxISR(d)
	}
	if ctrl&TIE != 0 && stat&TDRE != 0 {
		txISR(d)
	}
}
