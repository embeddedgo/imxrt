// Copyright 2024 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tftdci

import "github.com/embeddedgo/imxrt/hal/lpi2c"

// LPI2C is an implementation of the display/tft.DCI interface that uses an
// LPI2C peripheral to communicate with the display in what is known as 4-line
// serial mode.
//
// Limitations
//
// Written with the SSD1306 OLED controller in mind and the pix/driver/fbdrv
// driver so it doesn't support ReadBytes, WriteWords, WriteWordN. The Cmd and
// the Write* methods are implemented as single I2C transactions: I2C Start,
// SSD1306 control byte, command/data bytes, I2C Stop (the Co bit in the
// Control Byte is cleared).
type LPI2C struct {
	m    *lpi2c.Master
	addr uint8
}

// NewLPI2C returns new LPI2C based implementation of tftdrv.DCI. User must
// provide a configured LPI2C master driver and the slave address.
func NewLPI2C(m *lpi2c.Master, addr uint8) *LPI2C {
	return &LPI2C{m, addr << 1}
}

func (dci *LPI2C) Driver() *lpi2c.Master { return dci.m }
func (dci *LPI2C) Err(clear bool) error  { return dci.m.Err(clear) }

func start(m *lpi2c.Master, addr uint8) {
	m.Lock()
	m.WriteCmd(lpi2c.Start | int16(addr))
}

func stop(m *lpi2c.Master) {
	m.WriteCmd(lpi2c.Stop)
	m.Unlock()
}

const (
	singlCmd  = 0b1000_0000
	multiCmd  = 0b0000_0000
	singlData = 0b1100_0000
	multiData = 0b0100_0000
)

func (dci *LPI2C) Cmd(p []byte, dataMode int) {
	m := dci.m
	start(m, dci.addr)

	m.WriteCmd(lpi2c.Send | multiCmd)
	m.Write(p)

	stop(m)
}

func (dci *LPI2C) End() {
}

func (dci *LPI2C) WriteBytes(p []uint8) {
	m := dci.m
	start(m, dci.addr)

	m.WriteCmd(lpi2c.Send | multiData)
	m.Write(p)

	stop(m)
}

func (dci *LPI2C) WriteString(s string) {
	m := dci.m
	start(m, dci.addr)

	m.WriteCmd(lpi2c.Send | multiData)
	m.WriteString(s)

	stop(m)
}

func (dci *LPI2C) WriteByteN(b byte, n int) {
	m := dci.m
	start(m, dci.addr)

	m.WriteCmd(lpi2c.Send | multiData)
	for n != 0 {
		m.WriteCmd(lpi2c.Send | int16(b))
		n--
	}

	stop(m)
}

func (dci *LPI2C) WriteWords(p []uint16) {
	// unsupported
}

func (dci *LPI2C) WriteWordN(w uint16, n int) {
	// unsupported
}

func (dci *LPI2C) ReadBytes(p []byte) {
	// unsupported
}
