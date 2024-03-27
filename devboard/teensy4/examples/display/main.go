// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Display draws on the connected display.
package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/examples"

	"github.com/embeddedgo/imxrt/dci/tftdci"
	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpspi"
	"github.com/embeddedgo/imxrt/hal/lpspi/lpspi4dma"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func main() {
	// IO pins
	csn := pins.P10  // B0_00
	mosi := pins.P11 // B0_02
	miso := pins.P12 // B0_01
	sck := pins.P13  // B0_03 // shared with the onboard LED
	dc := pins.P14   // AD_B1_02
	rst := pins.P15  // AD_B1_03 // optional, connect to 3V (exception SSD1306)

	// GPIO output for the display reset signal (optional, exception SSD1306).
	reset := gpio.UsePin(rst, false)
	reset.Port().EnableClock(true)
	reset.SetDirOut(true)
	reset.Clear()           // set reset initial steate low
	rst.Setup(iomux.Drive2) // set rst as output, low state resets the display
	time.Sleep(time.Millisecond)
	reset.Set()

	// Setup LPSPI driver
	spi := lpspi4dma.Master() // lpspi4.Master() is better for small displays
	spi.UsePin(miso, lpspi.SDI)
	spi.UsePin(mosi, lpspi.SDO)
	spi.UsePin(csn, lpspi.PCS0)
	spi.UsePin(sck, lpspi.SCK) // disabled leds.User
	spi.Setup(33.25e6)

	//dp := displays.Adafruit_0i96_128x64_OLED_SSD1306()
	//dp := displays.Adafruit_1i5_128x128_OLED_SSD1351()
	//dp := displays.Adafruit_1i54_240x240_IPS_ST7789()
	dp := displays.Adafruit_2i8_240x320_TFT_ILI9341()
	//dp := displays.ERTFTM_1i54_240x240_IPS_ST7789()
	//dp := displays.MSP4022_4i0_320x480_TFT_ILI9486()
	//dp := displays.Waveshare_1i5_128x128_OLED_SSD1351()

	// Most of the displays accept significant overclocking.
	//dp.MaxReadClk *= 2
	//dp.MaxWriteClk *= 2

	dci := tftdci.NewLPSPI(spi, dc, lpspi.CPOL0|lpspi.CPHA0, dp.MaxReadClk, dp.MaxWriteClk)

	fmt.Println("*** Start ***")

	disp := dp.New(dci)
	for {
		examples.RotateDisplay(disp)
		examples.DrawText(disp)
		examples.GraphicsTest(disp)
	}
}
