// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"os"
	"time"

	"github.com/embeddedgo/device/adc/ads111x"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func logI2CErr(d *lpi2c.Master) bool {
	if err := d.Err(true); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return true
	}
	return false
}

var d *lpi2c.Master

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	p := lpi2c.LPI2C(1)
	d = lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	d.Setup(lpi2c.Fast400k)
	d.UsePin(scl, lpi2c.SCL)
	d.UsePin(sda, lpi2c.SDA)
	irq.LPI2C1.Enable(rtos.IntPrioLow, 0)

	const (
		addr = 0b1001000 << 1 // address when the ADDR pin is connected to GND
		wr   = 0              // write transaction
		rd   = 1              // read transaction
	)

	var buf [2]byte

again:
	time.Sleep(3 * time.Second)
	for i := 0; ; i++ {
		cfg := ads111x.OS | ads111x.AIN0_AIN1 | ads111x.FS2048 | ads111x.SINGLESHOT | ads111x.R8
		d.WriteCmd(
			lpi2c.StartNACK|0b0000_1000, // switch to High Speed mode
			lpi2c.StartHS|addr|wr,
			lpi2c.Send|ads111x.RegCfg,
			lpi2c.Send|int16(cfg>>8),
			lpi2c.Send|int16(cfg&0xff),
		)
		if logI2CErr(d) {
			goto again
		}
		for {
			d.WriteCmd(
				lpi2c.StartHS|addr|rd,
				lpi2c.Recv|int16(len(buf)-1),
			)
			d.Read(buf[:])
			if logI2CErr(d) {
				goto again
			}
			cfg = uint16(buf[0])<<8 | uint16(buf[1])
			if cfg&ads111x.OS != 0 {
				break
			}
		}
		d.WriteCmd(
			lpi2c.StartHS|addr|wr,
			lpi2c.Send|ads111x.RegV,
			lpi2c.StartHS|addr|rd,
			lpi2c.Recv|int16(len(buf)-1),
			lpi2c.Stop,
		)
		d.Read(buf[:])
		if logI2CErr(d) {
			goto again
		}
		scale := 6.144
		if shift := cfg & ads111x.PGA >> ads111x.PGAn; shift != 0 {
			scale = 4.096 / float64(uint(1)<<(shift-1))
		}
		v := int16(buf[0])<<8 | int16(buf[1])
		fmt.Printf("%d: %.6f V\n", i, float64(v)*scale/0x8000)
		time.Sleep(time.Second)
	}

}

//go:interrupthandler
func LPI2C1_Handler() {
	d.ISR()
}
