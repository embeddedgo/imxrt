// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Ads1115 communicates with the ADS1115 analog to digital converter using I2C
// protocol. It uses the single-shot mode of operation to periodically start an
// A/D conversion to read a voltage between the AIN0 and ANI1 pins.
//
// This example presents two possible ways to communicate with an I2C device.
// The first one uses the low-level interface (lowLevelWay function), the second
// one uset the connection interface (connWay function). The low-level interface
// may be more efficient and simpler in some specific cases. The connection
// interface is more portable. Both can be used simultaneously as in this
// example. If used concurently by multiple gorutines, the low-level interface
// require using the Master.Mutex to ensure exclusive access to the Master
// driver (the connection interface does it by default).
//
// This exapmle also presents how to use the I2C High-Sped mode. If you
// encounter errors try to disable it. In case of the connection interface it's
// easy: simply remove the i2cbus.HS flag from the adderess. In case of the
// low-level interface you have to some remove/change Start* commands so it
// requires more knowledge. If you are unsure how to do it simply comment out
// the lowLevelWay function call.
package main

import (
	"embedded/rtos"
	"fmt"
	"os"
	"time"

	"github.com/embeddedgo/device/adc/ads111x"
	"github.com/embeddedgo/device/bus/i2cbus"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

const (
	addr = 0b1001000 // address if the ADDR pin is connected to GND
	cfg  = ads111x.OS | ads111x.AIN0_AIN1 | ads111x.FS2048 | ads111x.SINGLESHOT | ads111x.R8
)

var master *lpi2c.Master

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	p := lpi2c.LPI2C(1)
	master = lpi2c.NewMaster(p, dma.Channel{})
	master.Setup(lpi2c.Std100k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)
	irq.LPI2C1.Enable(rtos.IntPrioLow, 0)

	c := master.NewConn(addr | i2cbus.HS)
	for {
		time.Sleep(time.Second)
		lowLevelWay(master)
		time.Sleep(time.Second)
		connWay(c)
	}
}

func lowLevelWay(d *lpi2c.Master) {
	const (
		wr = 0 // write transaction
		rd = 1 // read transaction
	)
	var buf [2]byte

	// Start a new A/D conversion.
	d.WriteCmds([]int16{
		lpi2c.StartNACK | 0b0000_1000, // switch to I2C High Speed mode
		lpi2c.StartHS | addr<<1 | wr,
		lpi2c.Send | ads111x.RegCfg,
		lpi2c.Send | int16(cfg>>8),
		lpi2c.Send | int16(cfg&0xff),
	})
	if logErr(d.Err(true)) {
		return
	}

	// Wait for the end of conversion.
	for {
		d.WriteCmd(lpi2c.StartHS | addr<<1 | rd)
		d.WriteCmd(lpi2c.Recv | int16(len(buf)-1))
		d.Read(buf[:])
		if logErr(d.Err(true)) {
			return
		}
		if buf[0]&byte(ads111x.OS>>8) != 0 {
			break
		}
	}

	// Read the result of ADC.
	d.WriteCmds([]int16{
		lpi2c.StartHS | addr<<1 | wr,
		lpi2c.Send | ads111x.RegV,
		lpi2c.StartHS | addr<<1 | rd,
		lpi2c.Recv | int16(len(buf)-1),
		lpi2c.Stop,
	})
	d.Read(buf[:])
	if logErr(d.Err(true)) {
		return
	}

	// Convert to volts and print.
	fmt.Printf("lowLevelWay: %.6f V\n", volt(buf))
}

func connWay(c i2cbus.Conn) {
	defer c.Close() // I2C Stop
	var buf [2]byte

	// Start a new A/D conversion.
	_, err := c.Write([]byte{
		ads111x.RegCfg,
		byte(cfg >> 8),
		byte(cfg & 0xff),
	})
	if logErr(err) {
		return
	}

	// Wait for the end of conversion.
	for {
		_, err := c.Read(buf[:])
		if logErr(err) {
			return
		}
		if buf[0]&byte(ads111x.OS>>8) != 0 {
			break
		}
	}

	// Read the result of ADC.
	err = c.WriteByte(ads111x.RegV)
	if logErr(err) {
		return
	}
	_, err = c.Read(buf[:])
	if logErr(err) {
		return
	}

	// Convert to volts and print.
	fmt.Printf("connWay:     %.6f V\n", volt(buf))
}

func volt(buf [2]byte) float64 {
	scale := 6.144
	if shift := cfg & ads111x.PGA >> ads111x.PGAn; shift != 0 {
		scale = 4.096 / float64(uint(1)<<(shift-1))
	}
	return float64(int16(buf[0])<<8|int16(buf[1])) * scale / 0x8000
}

func logErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintln(os.Stderr, err)
	return true
}

//go:interrupthandler
func LPI2C1_Handler() {
	master.ISR()
}
