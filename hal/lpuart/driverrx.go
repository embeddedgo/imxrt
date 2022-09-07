// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/rtos"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/internal"
)

// EnableRx enables receiving data into internal ring buffer of size bufLen
// characters. The minimum size of the buffer is 2 characters. If DMA is used
// the buffer size is limited to 3276 characters.
func (d *Driver) EnableRx(bufLen int) {
	if d.rxbuf != nil {
		panic("enabled before")
	}
	if bufLen < 2 {
		panic("lpuart: bufLen < 2")
	}
	if rxdma := d.rxdma; rxdma.IsValid() {
		if bufLen > 32767 {
			bufLen = 32767
		}
		d.rxbuf = dma.Alloc[DATA](bufLen)
		ptr, size := unsafe.Pointer(&d.rxbuf[0]), len(d.rxbuf)*2
		rtos.CacheMaint(rtos.DCacheCleanInval, ptr, size)
		tcd := dma.TCD{
			SADDR:       unsafe.Pointer(d.p.DATA.Addr()),
			ATTR:        dma.S16b | dma.D16b,
			ML_NBYTES:   2,
			DADDR:       ptr,
			DOFF:        2,
			ELINK_CITER: int16(len(d.rxbuf)),
			DLAST_SGA:   int32(-size),
			CSR:         dma.INTMAJOR,
			ELINK_BITER: int16(len(d.rxbuf)),
		}
		rxdma.WriteTCD(&tcd)
		rxdma.SetMux((dma.LPUART1_RX + dma.Mux(num(d.p))*2) | dma.En)
		rxdma.EnableReq()
	} else {
		if bufLen > 1<<nshift {
			bufLen = 1 << nshift
		}
		d.rxbuf = make([]DATA, bufLen)
	}
	internal.AtomicStoreBits(&d.p.CTRL, RE|RIE, RE|RIE)
}

// DisableRx disables receiver.
func (d *Driver) DisableRx() {
	p := d.p
	internal.AtomicStoreBits(&p.CTRL, RE|RIE, 0)
	for p.CTRL.LoadBits(RE) != 0 {
		// wait for receiver to finish receiving the last character
	}
	for p.DATA.Load()&RXEMPT == 0 {
		// empty the FIFO to ensure the Rx ISR will no longer use the buffer
	}
	d.rxbuf = nil
}

const (
	nshift = 28
	imask  = uint32(0xffff_ffff) >> (32 - nshift)
	nmask  = int(uint32(0xffff_ffff) >> nshift)
)

func rxISR(d *Driver) {
	wake := atomic.CompareAndSwapUint32(&d.rxwake, 1, 0)
	if rxdma := d.rxdma; rxdma.IsValid() {
		if wake {
			// Complete the DMA disabling sequence (see IMXRT1060RM_rev3 6.4.8).
			// Skip RDMAE check because it should be 0 if Rx IRQ is taken.
			for rxdma.IsReq() {
			}
			// If rxbuf is still empty make the first character available to the
			// waiting goroutine before enabling DMA. Required because there is
			// no guarantee that the DMA will read anything on time.
			if getNextwDMA(d) == d.nextr {
				d.rxfirst = d.p.DATA.Load()
			}
			d.p.BAUD.StoreBits(RDMAE, RDMAE) // enable DMA and gate IRQ
			d.rxready.Wakeup()
		}
		return
	}
	if wake {
		// Use d.rxfirst also in non-DMA mode (works well, preserves symmetry).
		d.rxfirst = d.p.DATA.Load()
		d.rxready.Wakeup()
	}
	// Empty the FIFO quickly. Goal is speed so d.nextw is updated at the end
	// and there is no overflow checking (same as in DMA mode).
	iw := int(d.nextw & imask)
	nw := int(d.nextw >> nshift)
	rxbuf := d.rxbuf
	dr := &d.p.DATA
	for {
		data := dr.Load()
		if data&RXEMPT != 0 {
			break
		}
		rxbuf[iw] = data
		if iw++; iw == len(rxbuf) {
			iw = 0
			nw++
		}
	}
	atomic.StoreUint32(&d.nextw, uint32(nw<<nshift|iw))
}

func waitRxData(d *Driver) (int, error) {
	nextr := d.nextr
	nextw := atomic.LoadUint32(&d.nextw)
	if nextw != nextr {
		goto dataInBuffer
	}
	if stat := &d.p.STAT; stat.Load()&OR != 0 {
		stat.Store(OR)
		return 0, Error(OR)
	}
	d.rxready.Clear()
	atomic.StoreUint32(&d.rxwake, 1)
	nextw = atomic.LoadUint32(&d.nextw)
	if nextw != nextr {
		if !atomic.CompareAndSwapUint32(&d.rxwake, 1, 0) {
			d.rxready.Sleep(-1) // wait for the upcoming wake-up
		}
		goto dataInBuffer
	}
	if !d.rxready.Sleep(d.rxtimeout) {
		if atomic.CompareAndSwapUint32(&d.rxwake, 1, 0) {
			return 0, ErrTimeout
		}
		d.rxready.Sleep(-1) // wait for the upcoming wake-up
	}
	nextw = atomic.LoadUint32(&d.nextw)
dataInBuffer:
	iw, nw := int(nextw&imask), int(nextw>>nshift)
	ir, nr := int(nextr&imask), int(nextr>>nshift)
	if n := (nw - nr) & nmask; n == 0 && iw < ir || n == 1 && iw > ir || n > 1 {
		// Discard buffered data. Does it make sense to salvage something?
		d.nextr = nextw
		return 0, ErrBufOverflow
	}
	return iw, nil
}

func getNextwDMA(d *Driver) uint32 {
	var citer int
	rxdma := d.rxdma
	tcd := rxdma.TCD()
	irq := rxdma.IsInt()
	for {
		citer = int(tcd.ELINK_CITER.Load())
		irq1 := rxdma.IsInt()
		if irq == irq1 {
			break
		}
		irq = irq1
	}
	if irq {
		// Use INTMAJOR to detect CITER wrapping.
		rxdma.ClearInt()
		d.rxdman++ // approximate the non-DMA nw
	}
	return uint32(len(d.rxbuf)-citer) | uint32(d.rxdman)<<nshift
}

func disableIRQenableDMAifnoISR(d *Driver) (noisr bool) {
	internal.AtomicStoreBits(&d.p.CTRL, RIE, 0) // ensure valid noisr below
	noisr = atomic.CompareAndSwapUint32(&d.rxwake, 1, 0)
	if noisr {
		d.p.BAUD.StoreBits(RDMAE, RDMAE) // enable DMA and gate IRQ
	}
	internal.AtomicStoreBits(&d.p.CTRL, RIE, RIE) // reanable gated IRQ
	return
}

func waitRxDataDMA(d *Driver, m int) (int, error) {
	nextr := d.nextr
	nextw := getNextwDMA(d)
	if nextw != nextr {
		goto dataInBuffer
	}
	if stat := &d.p.STAT; stat.Load()&OR != 0 {
		stat.Store(OR)
		return 0, Error(OR)
	}
	d.rxwake = 1
	d.rxready.Clear()
	d.p.BAUD.StoreBits(RDMAE, 0) // disable DMA and ungate IRQ
	nextw = getNextwDMA(d)
	if nextw != nextr {
		if !disableIRQenableDMAifnoISR(d) {
			d.rxready.Sleep(-1) // wait for the upcoming wake-up
		}
		goto dataInBuffer
	}
	if !d.rxready.Sleep(d.rxtimeout) {
		if disableIRQenableDMAifnoISR(d) {
			return 0, ErrTimeout
		}
		d.rxready.Sleep(-1) // wait for the upcoming wake-up
	}
	nextw = getNextwDMA(d)
dataInBuffer:
	iw, nw := int(nextw&imask), int(nextw>>nshift)
	ir, nr := int(nextr&imask), int(nextr>>nshift)
	n := (nw - nr) & nmask
	if n == 0 && iw < ir || n == 1 && iw > ir || n > 1 {
		// Discard buffered data. Does it make sense to salvage something?
		d.nextr = nextw
		d.nextw = nextw // repurpose d.nextw for cache maintanence pointer
		return 0, ErrBufOverflow
	}
	n = n*len(buf) + (iw - ir) // number of buffered words
	if m > n {
		m = n
	}
	if d.nextw != nextw {
		d.nextw = nextw
		ptr, size := unsafe.Pointer(&d.rxbuf[0]), len(d.rxbuf)*2
		rtos.CacheMaint(rtos.DCacheInval, ptr, size)

	}
	return iw, nil
}

const dataErrMask = FRETSC | PARITYE | NOISY

func dataError(w DATA) error {
	return Error(w&FRETSC)<<(FEn-FRETSCn) | Error(w&PARITYE)<<(PFn-PARITYEn) |
		Error(w&NOISY)<<(NFn-NOISYn)
}

// ReadByte implements the io.ByteReader interface.
func (d *Driver) ReadByte() (byte, error) {
	var err error
	if d.rxdma.IsValid() {
		_, err = waitRxDataDMA(d, 1)
	} else {
		_, err = waitRxData(d)
	}
	if err != nil {
		return 0, err
	}
	w := d.rxfirst
	if w&RXEMPT == 0 {
		d.rxfirst = RXEMPT
	} else {
		ir, nr := int(d.nextr&imask), int(d.nextr>>nshift)
		w = d.rxbuf[ir]
		if ir++; ir == len(d.rxbuf) {
			ir = 0
			nr++
		}
		d.nextr = uint32(nr<<nshift | ir)
	}
	if w&dataErrMask != 0 {
		err = dataError(w)
	}
	return byte(w), err
}

// Read implements the io.Reader interface.
func (d *Driver) Read(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}
	var (
		iw  int
		err error
	)
	if d.rxdma.IsValid() {
		iw, err = waitRxDataDMA(d, len(buf))
	} else {
		iw, err = waitRxData(d)
	}
	if err != nil {
		return 0, err
	}
	n := 0
	if w := d.rxfirst; w&RXEMPT == 0 {
		d.rxfirst = RXEMPT
		buf[n] = byte(w)
		n++
		if w&dataErrMask != 0 {
			err = dataError(w)
			return n, err
		}
	}
	ir, nr := int(d.nextr&imask), int(d.nextr>>nshift)
	for ir != iw && n != len(buf) {
		w := d.rxbuf[ir]
		buf[n] = byte(w)
		n++
		if ir++; ir == len(d.rxbuf) {
			ir = 0
			nr++
		}
		if w&dataErrMask != 0 {
			err = dataError(w)
			break
		}
	}
	d.nextr = uint32(nr<<nshift | ir)
	return n, err
}
