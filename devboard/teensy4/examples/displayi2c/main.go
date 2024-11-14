// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Displayi2c draws on the connected I2C display.
package main

import (
	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/examples"

	"github.com/embeddedgo/imxrt/dci/tftdci"
	"github.com/embeddedgo/imxrt/hal/lpi2c"
	"github.com/embeddedgo/imxrt/hal/lpi2c/lpi2c1"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func main() {
	// IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	master := lpi2c1.Master()
	master.Setup(lpi2c.Fast400k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)

	dci := tftdci.NewLPI2C(master, 0b0111100)
	disp := displays.Adafruit_0i96_128x64_OLED_SSD1306().New(dci)
	for {
		examples.RotateDisplay(disp)
		examples.DrawText(disp)
		examples.GraphicsTest(disp)
	}
}
