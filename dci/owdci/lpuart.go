// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package owdci provides implementation of onewire.DCI interface.
package owdci

import (
	"time"

	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/onewire"
)

// LPUART wraps an lpuart.Driver to implements the onewire.DCI.
type LPUART lpuart.Driver

func lpuartDrv(dci *LPUART) *lpuart.Driver { return (*lpuart.Driver)(dci) }

// SetupLPUART configures d to be used as onewire.DCI.
func SetupLPUART(d *lpuart.Driver) *LPUART {
	d.Setup(lpuart.Word8b, 115200)
	d.EnableRx(64)
	d.EnableTx()
	return (*LPUART)(d)
}

func (dci *LPUART) Reset() error {
	d := lpuartDrv(dci)
	p := d.Periph()
	p.SetBaudrate(9600)
	err := d.WriteByte(0xf0)
	if err != nil {
		return err
	}
	d.SetReadTimeout(time.Second)
	r, err := d.ReadByte()
	if err != nil {
		return err
	}
	if r == 0xf0 {
		return onewire.ErrNoResponse
	}
	p.SetBaudrate(115200)
	return nil
}

func ignoreNoise(err error) error {
	e, ok := err.(lpuart.Error)
	if !ok {
		return err
	}
	if e != lpuart.ENOISE {
		return e &^ lpuart.ENOISE
	}
	return nil
}

func sendRecvSlot(d *lpuart.Driver, slot byte) (byte, error) {
	if err := d.WriteByte(slot); err != nil {
		d.DiscardRx()
		return 0, err
	}
	d.SetReadTimeout(time.Second)
	b, err := d.ReadByte()
	if err != nil {
		if err = ignoreNoise(err); err != nil {
			d.DiscardRx()
		}
	}
	return b, err
}

func sendRecv(d *lpuart.Driver, slots *[8]byte) error {
	if _, err := d.Write(slots[:]); err != nil {
		d.DiscardRx()
		return err
	}
	d.SetReadTimeout(time.Second)
	for n := 0; n < len(slots); {
		m, err := d.Read(slots[n:])
		n += m
		if err != nil {
			if err = ignoreNoise(err); err != nil {
				d.DiscardRx()
				return err
			}
		}
	}
	return nil
}

func (dci *LPUART) ReadBit() (int, error) {
	slot, err := sendRecvSlot(lpuartDrv(dci), 0xff)
	if err != nil {
		return 0, err
	}
	return int(slot) & 1, nil
}

func (dci *LPUART) WriteBit(bit int) error {
	var b byte
	if bit&1 != 0 {
		b = 0xff
	}
	d := lpuartDrv(dci)
	slot, err := sendRecvSlot(d, b)
	if err != nil {
		return err
	}
	if slot != b {
		d.DiscardRx()
		return onewire.ErrBusFault
	}
	return nil
}

func (dci *LPUART) ReadByte() (byte, error) {
	slots := [8]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	err := sendRecv(lpuartDrv(dci), &slots)
	var v int
	for i, slot := range slots {
		v += int(slot&1) << uint(i)
	}
	return byte(v), err
}

func (dci *LPUART) WriteByte(b byte) error {
	var slots [8]byte
	v := int(b)
	for i := range slots {
		if v&1 != 0 {
			slots[i] = 0xff
		}
		v >>= 1
	}
	d := lpuartDrv(dci)
	if err := sendRecv(d, &slots); err != nil {
		return err
	}
	v = int(b)
	for i, slot := range slots {
		r := v & (1 << uint(i))
		if r != 0 {
			r = 0xff
		}
		if int(slot) != r {
			d.DiscardRx()
			return onewire.ErrBusFault
		}
	}
	return nil
}
