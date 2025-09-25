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
	"math/rand"
	"time"

	"github.com/embeddedgo/device/bus/i2cbus"
	"github.com/embeddedgo/imxrt/hal/lpi2c"
	"github.com/embeddedgo/imxrt/hal/lpi2c/lpi2c1"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

const (
	prefix   = 0b1010 // address prefix (0xa)
	a2a1a0   = 0b000  // address pins
	pageSize = 32
)

func randomData(p []byte) {
	for i := range p {
		p[i] = byte(rand.Int31()&63 + ';')
	}
}

func main() {
	// Used IO pins
	const (
		// UART
		conRx = pins.P23
		conTx = pins.P24

		// I2C
		sda = pins.P8 // AD_B1_01
		scl = pins.P9 // AD_B1_00
	)

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	// LPI2C driver
	master := lpi2c1.Master()
	master.Setup(lpi2c.Std100k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)

	c := master.NewConn(prefix<<3 | a2a1a0)

	var out, in [pageSize]byte

loop:
	for page := 0; ; page++ {
		a := page * pageSize
		addr := []byte{byte(a >> 8), byte(a)}

		n := rand.Intn(pageSize) + 1
		randomData(out[:n])

		fmt.Printf("Wr %2d B page %d: %s ", n, page, out[:n])
		c.Write(addr)
		c.Write(out[:n])
		err := c.Close()
		if err != nil {
			fmt.Println("\nWr error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Wait for the end of write
		for {
			c.Write(addr)
			err = c.Close()
			if err == nil {
				break
			}
			if !errors.Is(err, i2cbus.ErrACK) {
				fmt.Print("\nwait error: ", err, "\n")
				time.Sleep(2 * time.Second)
				continue loop
			}
			fmt.Print(".")
		}
		fmt.Println(" done")

		c.Read(in[:n])
		err = c.Close()
		if err != nil {
			fmt.Println("Rd error:", err)
			time.Sleep(2 * time.Second)
		} else if string(in[:n]) != string(out[:n]) {
			fmt.Printf("Rd %2d B BAD! %d: %s\n\n", n, page, in[:n])
			time.Sleep(2 * time.Second)
		} else {
			fmt.Print("Rd OK\n")
		}
		for i := range in {
			in[i] = ':'
		}
	}
}
