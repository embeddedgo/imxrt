// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Eeprom writes and reads the memory of the 24C64/128/256 I2C EEPROM. The
// difference to the less dense 24C0x EEPROMs is the use of 16 bit memory
// address instead of 8 bit one.
package main

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/embeddedgo/device/bus/i2cbus"
	"github.com/embeddedgo/imxrt/hal/lpi2c"
	"github.com/embeddedgo/imxrt/hal/lpi2c/lpi2c1"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

const (
	prefix = 0b1010 // address prefix (0xa)
	a2a1a0 = 0b000  // address pins
)

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	master := lpi2c1.Master()
	master.Setup(lpi2c.Std100k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)

	c := master.NewConn(prefix<<3 | a2a1a0)

	var buf [64]byte

loop:
	for page := 0; ; page++ {
		time.Sleep(2 * time.Second)
		a := page * 32
		addr := []byte{byte(a >> 8), byte(a)}

		s := fmt.Sprintf(">> %#x <<", page)

		fmt.Printf("\nWrite page %d: %s ", page, s)
		c.Write(addr) // replace with c.WriteByte(addr) for 24C0x EEPROMs
		io.WriteString(c, s)
		err := c.Close()
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}

		// Wait for the end of write
		for {
			c.Write(nil)
			err := c.Close()
			if err == nil {
				break
			}
			if !errors.Is(err, i2cbus.ErrACK) {
				fmt.Println("wait error:", err)
				continue loop
			}
			fmt.Print(".")
		}
		fmt.Println(" done")

		c.Write(addr) // replace with c.WriteByte(addr) for 24C0x EEPROMs
		n, _ := c.Read(buf[:len(s)])
		err = c.Close()
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		fmt.Printf("Read %d bytes from page %d: %s\n", n, page, string(buf[:n]))
	}

}
