// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"runtime"
	"time"
	"unsafe"
)

// EnableTx enables Tx part of the LPUART peripheral and setups Tx DMA channel.
func (d *Driver) EnableTx() {
	d.p.CTRL.SetBits(TE)
}

// DisableTx disables Tx part of the LPUART peripheral.
func (d *Driver) DisableTx() {
	for d.p.STAT.LoadBits(TC) == 0 {
		runtime.Gosched()
	}
	d.p.CTRL.ClearBits(TE)
}

func isrTxNoDMA(d *Driver) {
	txn := d.txn
	stop := txn + d.txmax
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
		d.p.CTRL.ClearBits(TIE)
		d.txdone.Wakeup()
	}
}

func (d *Driver) WriteString(s string) (n int, err error) {
	switch len(s) {
	case 0:
		return
	case 1:
		return 1, d.WriteByte(s[0])
	}
	d.txdata = s
	d.txn = 0
	d.txdone.Clear()
	d.p.CTRL.SetBits(TIE)
	if d.txdone.Sleep(d.txtimeout) {
		d.p.CTRL.ClearBits(TIE)
		err = ErrTimeout
	}
	return len(s), err
}

func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

func (d *Driver) WriteByte(b byte) error {
	var start time.Time
	for int(d.p.WATER.LoadBits(TXCOUNT)>>TXCOUNTn) == d.txmax {
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
