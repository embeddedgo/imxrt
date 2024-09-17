// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"github.com/embeddedgo/device/bus/i2cbus"
)

// Name returns the driver name. The default name is the name of the underlying
// peripheral (e.g. "LPI2C1") but can be changed using SetName.
func (d *Master) Name() string {
	return d.name
}

// SetName allows to change the default driver name.
func (d *Master) SetName(s string) {
	d.name = s
}

// SetID sets the Master ID. The three least significant bits of this ID are
// used for arbitration between competing masters while switching to the high
// speed mode.
func (d *Master) SetID(id uint8) {
	d.id = id
}

// ID rteturns the Master ID. See SetID for more information.
func (d *Master) ID() uint8 {
	return d.id
}

type conn struct {
	d      *Master
	rstart [4]int16
	wstart [3]int16
	n      int8
	open   bool
	wr     bool
}

func (d *Master) NewConn(a i2cbus.Addr) i2cbus.Conn {
	c := &conn{d: d}
	start := Start
	if a&i2cbus.HS != 0 {
		start = StartHS
		cmd := StartNACK | 0x08 | int16(d.id&3)
		c.wstart[c.n] = cmd
		c.rstart[c.n] = cmd
		c.n++
	}
	if a&i2cbus.A10 == 0 {
		cmd := start | int16(a&0x7f)<<1
		c.wstart[c.n] = cmd
		c.rstart[c.n] = cmd | 1
		c.n++
	} else {
		cmd := start | 0xf0 | int16(a&0x300)>>7
		c.wstart[c.n] = cmd
		c.rstart[c.n] = cmd | 1
		c.n++
		cmd = Send | int16(a&0xff)
		c.wstart[c.n] = cmd
		c.rstart[c.n] = cmd
		c.n++
	}
	return c
}

func connErr(c *conn) (err error) {
	d := c.d
	err = d.Err(true)
	if err != nil {
		err = &i2cbus.MasterError{d.name, err}
	}
	return
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
		c.d.WriteCmd(c.wstart[i:c.n]...)
	}
}

func (c *conn) Write(p []byte) (n int, err error) {
	startWrite(c)
	c.d.Write(p)
	err = connErr(c)
	if err == nil {
		n = len(p)
	}
	return
}

func (c *conn) WriteByte(b byte) error {
	startWrite(c)
	c.d.WriteCmd(Send | int16(b))
	return connErr(c)
}

func startRead(c *conn, arg byte) {
	open := c.open
	if !open {
		c.open = true
		c.d.Lock()
	}
	c.wr = false
	i := 0
	if open && c.rstart[0]>>8 == StartNACK>>8 {
		i = 1
	}
	c.rstart[c.n] = Recv | int16(arg)
	c.d.WriteCmd(c.rstart[i : c.n+1]...)
}

func (c *conn) Read(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		return
	}
	if n > 256 {
		n = 256
	}
	startRead(c, byte(n-1))
	c.d.Read(p)
	err = connErr(c)
	if err != nil {
		n = 0
	}
	return
}

func (c *conn) ReadByte() (b byte, err error) {
	startRead(c, 0)
	b = c.d.ReadByte()
	err = connErr(c)
	return
}

func (c *conn) Close() error {
	if !c.open {
		panic("already closed")
	}
	c.d.WriteCmd(Stop)
	c.d.Flush()
	err := connErr(c)
	c.d.Unlock()
	c.open = false
	c.wr = false
	return err
}
