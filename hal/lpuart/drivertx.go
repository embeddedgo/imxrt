// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/rtos"
	"runtime"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/internal"
)

// EnableTx enables Tx part of the LPUART peripheral and setups Tx DMA channel.
func (d *Driver) EnableTx() {
	if c := d.txdma; c.IsValid() {
		c.DisableReq()
		c.SetMux((dma.LPUART1_TX + dma.Mux(num(d.p))) | dma.En)
	}
	internal.AtomicStoreBits(&d.p.CTRL, TE, TE)
}

// DisableTx disables Tx part of the LPUART peripheral. While Tx is disabled the
// Tx DMA channel can be used for other things.
func (d *Driver) DisableTx() {
	for d.p.STAT.LoadBits(TC) == 0 {
		runtime.Gosched()
	}
	internal.AtomicStoreBits(&d.p.CTRL, TE, 0)
	if c := d.txdma; c.IsValid() {
		c.SetMux(0)
	}
}

func txISR(d *Driver) {
	txn := d.txn
	stop := txn + 1<<d.txlog2max
	txdata := d.txdata
	if stop > len(txdata) {
		stop = len(txdata)
	}
	dr := &d.p.DATA
	for txn < stop {
		dr.Store(DATA(txdata[txn]))
		txn++
	}
	d.txn = txn
	if txn == len(txdata) {
		internal.AtomicStoreBits(&d.p.CTRL, TIE, 0)
		d.txdone.Wakeup()
	}
}

func (d *Driver) TxDMAISR() {
	ch := d.txdma
	if ch.IsInt() {
		ch.ClearInt()
		d.txdone.Wakeup()
	}
}

func writeString(d *Driver, s string) error {
	if len(s) == 1 {
		return d.WriteByte(s[0])
	}
	d.txdata = s
	d.txn = 0
	d.txdone.Clear()
	internal.AtomicStoreBits(&d.p.CTRL, TIE, TIE)
	if !d.txdone.Sleep(d.txtimeout) {
		internal.AtomicStoreBits(&d.p.CTRL, TIE, 0)
		return ErrTimeout
	}
	return nil
}

func writeStringDMA(d *Driver, s string) error {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&s))
	rtos.CacheMaint(rtos.DCacheClean, ptr, len(s))
	n := len(s) >> d.txlog2max // number of minor loops (major loop iterations)
	m := n
	if m > 32767 {
		m = 32767
	}
	tcd := dma.TCD{
		SADDR:       ptr,
		SOFF:        4,
		ATTR:        dma.S32b | dma.D8b,
		ML_NBYTES:   1 << d.txlog2max,
		DADDR:       unsafe.Pointer(d.p.DATA.Addr()),
		ELINK_CITER: int16(m),
		ELINK_BITER: int16(m),
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	d.txdma.WriteTCD(&tcd)
	for {
		d.txdma.EnableReq()
		if !d.txdone.Sleep(d.txtimeout) {
			// TODO: cancel transfer
			return ErrTimeout
		}
		if n -= m; n == 0 {
			break
		}
		tcdio := d.txdma.TCD()
		m = n
		if m >= 32767 {
			m = 32767
		} else {
			tcdio.ELINK_CITER.Store(int16(m))
			tcdio.ELINK_BITER.Store(int16(m))
		}
	}
	return nil
}

// dmaOffsets calculates a part of the string that is cache aligned
func dmaOffsets(s string) (start, end int) {
	const alignMask = dma.CacheLineSize - 1
	ptr := *(*uintptr)(unsafe.Pointer(&s))
	start = -int(ptr&alignMask) & alignMask
	ptr += uintptr(len(s))
	end = len(s) - int(ptr&alignMask)
	return
}

func (d *Driver) WriteString(s string) (n int, err error) {
	switch {
	case len(s) == 0:
		return
	case len(s) >= 32 && d.txdma.IsValid():
		if dmaStart, dmaEnd := dmaOffsets(s); dmaStart < dmaEnd {
			if dmaStart != 0 {
				err = writeString(d, s[:dmaStart])
				if err != nil {
					break
				}
			}
			err = writeStringDMA(d, s[dmaStart:dmaEnd])
			if err != nil {
				break
			}
			if dmaEnd != len(s) {
				err = writeString(d, s[dmaEnd:])
			}
			break
		}
		fallthrough
	default:
		err = writeString(d, s)
	}
	if err == nil {
		n = len(s)
	}
	return
}

func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

func (d *Driver) WriteByte(b byte) error {
	var start time.Time
	for int(d.p.WATER.LoadBits(TXCOUNT)>>TXCOUNTn) == 1<<d.txlog2max {
		if d.txtimeout >= 0 {
			t := time.Now()
			if start.IsZero() {
				start = t
				continue
			}
			if t.Sub(start) >= d.txtimeout {
				return ErrTimeout
			}
		}
		runtime.Gosched()
	}
	d.p.DATA.Store(DATA(b))
	return nil
}
