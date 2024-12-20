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
	"github.com/embeddedgo/imxrt/hal/lpspi/lpspi3dma"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// Serial console pins
	conRx := pins.P23
	conTx := pins.P24

	// Display pins
	rst := pins.P12  // AD_B1_02 // optional, connect to 3V (exception SSD1306)
	miso := pins.P91 // AD_B1_13
	mosi := pins.P92 // AD_B1_14
	csn := pins.P93  // AD_B1_12
	sck := pins.P94  // AD_B1_15
	dc := pins.P95   // AD_B1_11

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	// GPIO output for the display reset signal (optional, exception SSD1306).
	reset := gpio.UsePin(rst, false)
	reset.Port().EnableClock(true)
	reset.SetDirOut(true)
	reset.Clear()           // set reset initial steate low
	rst.Setup(iomux.Drive2) // set rst as output, low state resets the display
	time.Sleep(time.Millisecond)
	reset.Set()

	// Setup LPSPI driver
	spi := lpspi3dma.Master() // lpspi3.Master() is better for small displays
	spi.UsePin(miso, lpspi.SDI)
	spi.UsePin(mosi, lpspi.SDO)
	spi.UsePin(csn, lpspi.PCS0)
	spi.UsePin(sck, lpspi.SCK)
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
