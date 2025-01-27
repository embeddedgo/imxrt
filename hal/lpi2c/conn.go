// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"unsafe"

	"github.com/embeddedgo/device/bus/i2cbus"
)

// Name implements the i2cbus.Master interface. The default name is the name of
// the underlying peripheral (e.g. "LPI2C1") but can be changed using SetName.
func (d *Master) Name() string {
	return d.name
}

// SetName allows to change the default master name (see Name).
func (d *Master) SetName(s string) {
	d.name = s
}

// SetID sets the Master ID. Its three least significant bits are used for
// arbitration between competing masters while switching to the High Speed mode.
func (d *Master) SetID(id uint8) {
	d.id = id
}

// ID rteturns the Master ID. See SetID for more information.
func (d *Master) ID() uint8 {
	return d.id
}

type conn struct {
	d      *Master
	a      i2cbus.Addr
	wstart [3]int16
	rstart [6]int16
	wn     int8
	rn     int8
	open   bool
	wr     bool
}

// NewConn implements the i2cbus.Master interface.
func (d *Master) NewConn(a i2cbus.Addr) i2cbus.Conn {
	c := &conn{d: d, a: a}
	start := Start
	i := 0
	if a&i2cbus.HS != 0 {
		// High Speed mode
		cmd := StartNACK | 0x08 | int16(d.id&3)
		c.wstart[i] = cmd
		c.rstart[i] = cmd
		i++
		start = StartHS
	}
	if a&i2cbus.A10 == 0 {
		// 7b address
		cmd := start | int16(a&0x7f)<<1
		c.wstart[i] = cmd
		c.rstart[i] = cmd | 1
		i++
		c.wn = int8(i)
		c.rn = int8(i)
	} else {
		// 10b address
		cmd0 := start | 0xf0 | int16(a&0x300)>>7
		c.wstart[i] = cmd0
		c.rstart[i] = cmd0
		i++
		cmd1 := Send | int16(a&0xff)
		c.wstart[i] = cmd1
		c.rstart[i] = cmd1
		i++
		c.wn = int8(i)
		c.rstart[i] = cmd0 | 1
		i++
		c.rn = int8(i)
	}
	return c
}

// Addr implements the i2cbus.Conn interface.
func (c *conn) Addr() i2cbus.Addr {
	return c.a
}

// Master implements the i2cbus.Conn interface.
func (c *conn) Master() i2cbus.Master {
	return c.d
}

func startWrite(c *conn) {
	open := c.open
	if !open {
		c.open = true
		c.d.Lock()
	}
	if !c.wr {
		c.wr = true
		i := 0
		if open && c.wstart[0]>>8 == StartNACK>>8 {
			i = 1
		}
		c.d.WriteCmds(c.wstart[i:c.wn])
	}
}

// Write implements the i2cbus.Conn interface and the io.Writer interface.
func (c *conn) Write(p []byte) (n int, err error) {
	startWrite(c)
	if len(p) != 0 {
		c.d.Write(p)
		c.d.Flush() // ensure p isn't used after return
	}
	err = connErr(c)
	if err == nil {
		n = len(p)
	}
	return
}

// WriteString implements the io.StringWriter interface.
func (c *conn) WriteString(s string) (n int, err error) {
	return c.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

// WriteByte implements the i2cbus.Conn interface and the io.ByteWriter
// interface.
func (c *conn) WriteByte(b byte) error {
	startWrite(c)
	c.d.WriteCmd(Send | int16(b))
	return connErr(c)
}

func startRead(c *conn, m int) {
	open := c.open
	if !open {
		c.open = true
		c.d.Lock()
	}
	c.wr = false
	i := 0
	if open && c.rstart[0]>>8 == StartNACK>>8 {
		i = 1 // already in the High Speed mode
	}
	n := c.rn
	if m != 0 {
		c.rstart[n] = Recv | int16(m-1)
		n++
	}
	c.d.WriteCmds(c.rstart[i:n])
}

// Read implements the i2cbus.Conn interface and the io.Reader interface.
func (c *conn) Read(p []byte) (n int, err error) {
	n = len(p)
	if n > 256 {
		n = 256
	}
	startRead(c, n)
	c.d.Read(p)
	err = connErr(c)
	if err != nil {
		n = 0
	}
	return
}

// ReadByte implements the i2cbus.Conn interface and the io.ByteReader
// interface.
func (c *conn) ReadByte() (b byte, err error) {
	startRead(c, 1)
	b = c.d.ReadByte()
	err = connErr(c)
	return
}

// Close implements the i2cbus.Conn interface and the io.Closer interface.
func (c *conn) Close() error {
	if !c.open {
		return nil // already closed
	}
	d := c.d
	d.Clear(MSDF)
	d.WriteCmd(Stop)
	d.Wait(MSDF)
	err := connErr(c)
	if err == nil {
		d.Unlock()
		c.open = false
		c.wr = false
	}
	return err
}

func connErr(c *conn) (err error) {
	d := c.d
	err = d.Err(true)
	if err != nil {
		if err.(*MasterError).Status&MasterErrFlags == MNDF {
			err = i2cbus.ErrACK
		}
		err = &i2cbus.MasterError{Name: d.name, Err: err}
		c.d.Unlock()
		c.open = false
		c.wr = false
	}
	return
}
