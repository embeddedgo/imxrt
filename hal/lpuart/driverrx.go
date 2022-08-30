// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"sync/atomic"

	"github.com/embeddedgo/imxrt/hal/dma"
)

func (d *Driver) EnableRx(bufLen int) {
	if d.rxbuf != nil {
		panic("enabled before")
	}
	if bufLen < 2 {
		panic("bufLen < 2")
	}
	if d.rxdma.IsValid() {
		d.rxbuf = dma.Alloc[DATA](bufLen)
		// TODO: DMA version
	} else {
		d.rxbuf = make([]DATA, bufLen)
		d.p.CTRL.SetBits(RIE)
	}
	d.p.CTRL.SetBits(RE)
}

func (d *Driver) DisableRx() {
	d.p.CTRL.SetBits(RIE)
	d.p.CTRL.ClearBits(RE)
	for d.p.CTRL.LoadBits(RE) != 0 {
		// wait for receiver to finish receiving the last character
	}
	for d.p.DATA.Load()&RXEMPT == 0 {
		// empty the FIFO to ensure Rx ISR will no longer use the buffer
	}
	d.rxbuf = nil
}

// isrRxNoDMA reads from hardware FIFO until empty even if there is no space in
// d.rxbuf. This simplifies the receiving code and makes it possible to
// distinguish between the EOVERRUN (interrupt handler too slow or interrup
// latency to high) and the ErrBufOverflow (reading goroutine too slow).
func isrRxNoDMA(d *Driver) {
	nextw := d.nextw
	nextr := atomic.LoadUint32(&d.nextr)
	rxbuf := d.rxbuf
	dr := &d.p.DATA
	// We assume the loop below will empty the FIFO quickly...
	for {
		data := dr.Load()
		if data&RXEMPT != 0 {
			break
		}
		rxbuf[nextw] = data
		if nextw++; int(nextw) == len(rxbuf) {
			nextw = 0
		}
		if nextr == nextw {
			nextr = atomic.LoadUint32(&d.nextr)
			if nextr == nextw {
				atomic.StoreUint32(&d.overflow, 1)
				continue
			}
		}
	}
	// ...so we update d.nextw and check wake-up condition outside the loop. It
	// may not be optimal for multi-core systems but is simple and efficient.
	atomic.StoreUint32(&d.nextw, nextw)
	if atomic.LoadUint32(&d.rxwake) != 0 {
		d.rxwake = 0
		d.rxready.Wakeup()
	}
}

// Len returns the number of bytes that are ready to read from Rx buffer.
func (d *Driver) Len() int {
	n := int(atomic.LoadUint32(&d.nextw)) - int(d.nextr)
	if n < 0 {
		n += len(d.rxbuf)
	}
	return n
}

func waitRxData(d *Driver) int {
	nextw := atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		return int(nextw)
	}
	d.rxready.Clear()
	atomic.StoreUint32(&d.rxwake, 1)
	nextw = atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		if atomic.SwapUint32(&d.rxwake, 0) == 0 {
			d.rxready.Sleep(-1) // wait for the upcoming wake up
		}
		return int(nextw)
	}
	if !d.rxready.Sleep(d.rxtimeout) {
		if atomic.SwapUint32(&d.rxwake, 0) != 0 {
			return int(nextw)
		}
		d.rxready.Sleep(-1) // wait for the upcoming wake up
	}
	nextw = atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		return int(nextw)
	}
	panic("empty rxbuf")
}

func markDataRead(d *Driver, nextr int) error {
	if nextr >= len(d.rxbuf) {
		nextr -= len(d.rxbuf)
	}
	atomic.StoreUint32(&d.nextr, uint32(nextr))
	if atomic.LoadUint32(&d.overflow) != 0 {
		atomic.StoreUint32(&d.overflow, 0)
		return ErrBufOverflow
	}
	return nil
}

const dataErrMask = FRETSC | PARITYE | NOISY

func dataError(w DATA) error {
	return Error(w&FRETSC)<<(FEn-FRETSCn) |
		Error(w&PARITYE)<<(PFn-PARITYEn) |
		Error(w&NOISY)<<(NFn-NOISYn)
}

// ReadByte reads one byte and returns error if detected. ReadByte blocks only
// if the internal buffer is empty (d.Len() > 0 ensure non-blocking read).
func (d *Driver) ReadByte() (b byte, err error) {
	nextw := waitRxData(d)
	nextr := int(d.nextr)
	if nextw == nextr {
		return 0, ErrTimeout
	}
	w := d.rxbuf[nextr]
	b = byte(w)
	err = markDataRead(d, nextr+1)
	if err != nil {
		return
	}
	if w&dataErrMask != 0 {
		err = dataError(w)
	}
	return
}

func (d *Driver) Read(buf []byte) (n int, err error) {
	if len(buf) == 0 {
		return
	}
	nextw := waitRxData(d)
	nextr := int(d.nextr)
	if nextw == nextr {
		return 0, ErrTimeout
	}
	var dataErr error
	for {
		w := d.rxbuf[nextr]
		buf[n] = byte(w)
		n++
		if nextr++; nextr == len(d.rxbuf) {
			nextr = 0
		}
		if w&dataErrMask != 0 {
			dataErr = dataError(w)
			break
		}
		if nextr == int(nextw) || n == len(buf) {
			break
		}
	}
	if err = markDataRead(d, nextr); err == nil {
		err = dataErr
	}
	return
}
