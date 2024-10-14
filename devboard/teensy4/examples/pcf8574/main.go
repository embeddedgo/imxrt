// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Pcf8574 writes consecutive numbers to the remote I/O expander chip (PCF8574)
// using I2C protocol.
//
// The easiest way to try this example is to use a PCF8574 based module intended
// for LCD displays and one or more LEDs. Low-voltage LEDs like red ones
// require a current limiting resistor of the order 150-200 Ω. High voltage LEDs
// like the white ones may work without any resistor.
//
// Connect your LEDs between pin 1 (closest to the I2C connector, 3.3V) and pins
// 4, 5, 6 (PCF8574 P0, P1, P2 outputs). Polarity matters. Pin 1 should be
// connected to the anodes of all LEDS. The easiest way to do it is to use a
// breadboard. Next connect the module pins GND, VCC, SDA, SCL to the Teensy
// pins G, 3V, 18, 19. After programming your Teensy with this example the LEDs
// should start blinking with different frequencies.
//
// As the LEDs are connected between 3,3V and P0, P1, P2 writing the
// corresponding bit zero turns the LED on, writting one turns it off. Because
// of its quasi-bidirectional I/O the PCF8574 can't source enough current to
// stable light a LED connected between Px pin and GND.
package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/lpi2c"
	"github.com/embeddedgo/imxrt/hal/lpi2c/lpi2c1"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

const (
	prefix = 0b0100 // address prefix (0x4)
	a2a1a0 = 0b111  // address pins
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

	for i := 0; ; i++ {
		err := c.WriteByte(byte(i))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second / 4)
	}
}
