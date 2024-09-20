// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Eeprom writes and read the memory of the 24C64 I2C EEPROM.
package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

const (
	prefix = 0x1010 << 3 // address prefix
	e2e1e0 = 0           // address pins
)

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	p := lpi2c.LPI2C(1)
	d := lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	d.Setup(lpi2c.Std100k)
	d.UsePin(scl, lpi2c.SCL)
	d.UsePin(sda, lpi2c.SDA)

	c := d.NewConn(prefix | e2e1e0)

	time.Sleep(2 * time.Second)

	for i := 0; ; i++ {
		fmt.Println(i, "write")
		// Write a byte
		c.Write([]byte{byte(i >> 8), byte(i)}) // memory address
		c.WriteByte(byte(i))
		c.Close()
		// Read a byte
		fmt.Println(i, "read")
		for {
			_, err := c.Write([]byte{byte(i >> 8), byte(i)}) // memory address
			if err == nil {
				break
			}
			fmt.Println(err)
		}
		b, _ := c.ReadByte()
		fmt.Println("read byte:", b)

		time.Sleep(time.Second)
	}

}
