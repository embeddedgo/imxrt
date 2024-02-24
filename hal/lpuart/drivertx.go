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

// All dma.Mux slot constants are less than 128 so we can use the string
// conversion to group them in a constant array.
const txDMASlots = "" +
	string(dma.LPUART1_TX) +
	string(dma.LPUART2_TX) +
	string(dma.LPUART3_TX) +
	string(dma.LPUART4_TX) +
	string(dma.LPUART5_TX) +
	string(dma.LPUART6_TX) +
	string(dma.LPUART7_TX) +
	string(dma.LPUART8_TX)

// EnableTx enables Tx part of the LPUART peripheral.
func (d *Driver) EnableTx() {
	if txdma := d.txdma; txdma.IsValid() {
		txdma.DisableReq()
		txdma.SetMux(dma.Mux(txDMASlots[num(d.p)]) | dma.En)
	}
	internal.AtomicStoreBits(&d.p.CTRL, TE, TE)
}

// DisableTx disables Tx part of the LPUART peripheral.
func (d *Driver) DisableTx() {
	for d.p.STAT.LoadBits(TC) == 0 {
		runtime.Gosched()
	}
	internal.AtomicStoreBits(&d.p.CTRL, TE, 0)
	if c := d.txdma; c.IsValid() {
		c.SetMux(0)
	}
}

//go:nosplit
func txISR(d *Driver) {
	dr := &d.p.DATA
	n := d.txn
	if n < 0 {
		n = -n
	}
	m := n - d.txi
	if m > 1<<d.txlog2max {
		m = 1 << d.txlog2max
	}
	if d.txn >= 0 {
		addr := uintptr(d.txd) + uintptr(d.txi)
		end := addr + uintptr(m)
		for addr < end {
			dr.Store(uint16(*(*byte)(unsafe.Pointer(addr))))
			addr++
		}
	} else {
		addr := uintptr(d.txd) + uintptr(d.txi)*2
		end := addr + uintptr(m)*2
		for addr < end {
			dr.Store((*(*uint16)(unsafe.Pointer(addr))))
			addr++
		}
	}
	if d.txi += m; d.txi == n {
		internal.AtomicStoreBits(&d.p.CTRL, TIE, 0)
		d.txdone.Wakeup()
	}
}

func (d *Driver) TxDMAISR() {
	d.txdma.ClearInt()
	d.txdone.Wakeup()
}

func write(d *Driver, s string, s16 []uint16) (err error) {
	if len(s) != 0 {
		if len(s) == 1 {
			return d.WriteWord16(uint16(s[0]))
		}
		d.txd = *(*unsafe.Pointer)(unsafe.Pointer(&s))
		d.txn = len(s)
	} else {
		if len(s16) == 1 {
			return d.WriteWord16(s16[0])
		}
		d.txd = unsafe.Pointer(&s16[0])
		d.txn = -len(s16)
	}
	d.txdone.Clear()
	internal.AtomicStoreBits(&d.p.CTRL, TIE, TIE)
	if !d.txdone.Sleep(d.txtimeout) {
		internal.AtomicStoreBits(&d.p.CTRL, TIE, 0)
		err = ErrTimeout
	}
	d.txd = nil
	d.txi = 0
	return
}

// writeDMA only works with aligned multiples of 32 bytes
func writeDMA(d *Driver, s string, s16 []uint16) error {
	var (
		ptr  unsafe.Pointer
		n    int
		attr dma.ATTR
	)
	if len(s) != 0 {
		ptr = *(*unsafe.Pointer)(unsafe.Pointer(&s))
		n = len(s)
		attr = dma.S32b | dma.D8b
	} else {
		ptr = unsafe.Pointer(&s16[0])
		n = len(s16) * 2
		attr = dma.S32b | dma.D16b
	}
	rtos.CacheMaint(rtos.DCacheClean, ptr, n)
	n >>= d.txlog2max // number of minor loops (major loop iterations)
	m := n
	if m > 32767 {
		m = 32767
	}
	txdma := d.txdma
	tcd := dma.TCD{
		SADDR:       ptr,
		SOFF:        4,
		ATTR:        attr,
		ML_NBYTES:   1 << d.txlog2max,
		DADDR:       unsafe.Pointer(d.p.DATA.Addr()),
		ELINK_CITER: int16(m),
		CSR:         dma.DREQ | dma.INTMAJOR,
		ELINK_BITER: int16(m),
	}
	txdma.WriteTCD(&tcd)
	tcdio := txdma.TCD()
	for {
		d.txdone.Clear()
		txdma.EnableReq()
		if !d.txdone.Sleep(d.txtimeout) {
			txdma.DisableReq()
			for tcdio.CSR.LoadBits(dma.ACTIVE) != 0 {
				runtime.Gosched()
			}
			return ErrTimeout
		}
		if n -= m; n == 0 {
			break
		}
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
func dmaOffsets(p unsafe.Pointer, size int) (start, end uintptr) {
	const alignMask = dma.CacheLineSize - 1
	ptr := uintptr(p)
	start = -ptr & alignMask
	ptr += uintptr(size)
	end = uintptr(size) - ptr&alignMask
	return
}

// WriteString implements the io.StringWriter interface.
//
//go:nosplit
func (d *Driver) WriteString(s string) (n int, err error) {
	switch {
	case len(s) == 0:
		return
	case rtos.HandlerMode():
		return sysWrite(d, s)
	case len(s) >= 32 && d.txdma.IsValid():
		// DMA can handle only cache-aligned transfers, because of the required
		// cache maintenance operations that must don't overlap accidentally.
		dmaStart, dmaEnd := dmaOffsets(*(*unsafe.Pointer)(unsafe.Pointer(&s)), len(s))
		if dmaStart < dmaEnd {
			if dmaStart != 0 {
				err = write(d, s[:dmaStart], nil)
				if err != nil {
					break
				}
			}
			err = writeDMA(d, s[dmaStart:dmaEnd], nil)
			if err != nil {
				break
			}
			if dmaEnd != uintptr(len(s)) {
				err = write(d, s[dmaEnd:], nil)
			}
			break
		}
		fallthrough
	default:
		err = write(d, s, nil)
	}
	if err == nil {
		n = len(s)
	}
	return
}

// Write implements the io.Writer interface.
//
//go:nosplit
func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

// Write16 works like Write but writes 16-bit words to the DATA register.
func (d *Driver) Write16(s []uint16) (n int, err error) {
	switch {
	case len(s) == 0:
		return
	case len(s) >= 16 && d.txdma.IsValid():
		// DMA can handle only cache-aligned transfers, because of the required
		// cache maintenance operations that must don't overlap accidentally.
		dmaStart, dmaEnd := dmaOffsets(unsafe.Pointer(&s[0]), len(s)*2)
		dmaStart /= 2
		dmaEnd /= 2
		if dmaStart < dmaEnd {
			if dmaStart != 0 {
				err = write(d, "", s[:dmaStart])
				if err != nil {
					break
				}
			}
			err = writeDMA(d, "", s[dmaStart:dmaEnd])
			if err != nil {
				break
			}
			if dmaEnd != uintptr(len(s)) {
				err = write(d, "", s[dmaEnd:])
			}
			break
		}
		fallthrough
	default:
		err = write(d, "", s)
	}
	if err == nil {
		n = len(s)
	}
	return
}

// WriteWord16 works like WriteByte but writes 16-bit word to the DATA register.
func (d *Driver) WriteWord16(w uint16) error {
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
	d.p.DATA.Store(w)
	return nil
}

// WriteByte implements the io.ByteWriter interface.
func (d *Driver) WriteByte(b byte) error {
	return d.WriteWord16(uint16(b))
}

// sysWrite is called in handler mode. It is used by print and println mainly
// to print a stack trace before system halt.
// BUG: multiple cores not supported.
//
//go:nosplit
func sysWrite(d *Driver, s string) (int, error) {
	var dmux dma.Mux
	if c := d.txdma; c.IsValid() {
		dmux = c.Mux()
		c.SetMux(0) // stop request from LPUART
		for c.IsReq() {
			// wait for the end of current request
		}
	}
	p := d.p
	for _, b := range []byte(s) {
		for int(p.WATER.LoadBits(TXCOUNT)>>TXCOUNTn) == 1<<d.txlog2max {
			// busy waiting
		}
		p.DATA.Store(uint16(b))
	}
	if c := d.txdma; c.IsValid() {
		c.SetMux(dmux)
	}
	return len(s), nil
}
