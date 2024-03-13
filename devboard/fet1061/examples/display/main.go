// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Display draws on the connected display.
package main

import (
	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/driver/tftdrv/ili9341"
	"github.com/embeddedgo/display/pix/examples"

	"github.com/embeddedgo/imxrt/dci/tftdci"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dma/dmairq"
	"github.com/embeddedgo/imxrt/hal/lpspi"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// Assign GPIO pins
	conRx := pins.P23
	conTx := pins.P24

	miso := pins.P91 // AD_B1_13
	mosi := pins.P92 // AD_B1_14
	csn := pins.P93  // AD_B1_12
	sck := pins.P94  // AD_B1_15
	dc := pins.P95   // AD_B1_11

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	// Enable DMA controller and allocate two channels for the LPUART driver.
	d := dma.DMA(0)
	d.EnableClock(true)
	rxdma := d.AllocChannel(false)
	txdma := d.AllocChannel(false)

	// Setup LPSPI3 driver
	spi := lpspi.NewMaster(lpspi.LPSPI(3), rxdma, txdma)
	spi.UsePin(miso, lpspi.SDI)
	spi.UsePin(mosi, lpspi.SDO)
	spi.UsePin(csn, lpspi.PCS0)
	spi.UsePin(sck, lpspi.SCK)
	spi.Setup(19e6)
	spi.Enable()

	dmairq.SetISR(rxdma, spi.RxDMAISR)
	dmairq.SetISR(txdma, spi.TxDMAISR)

	writeClk := min(ili9341.MaxSPIWriteClock, 33e6)
	readClk := min(ili9341.MaxSPIReadClock, 33e6)

	dci := tftdci.NewLPSPI(spi, dc, lpspi.CPOL0|lpspi.CPHA0, writeClk, readClk)

	cmd := []byte{1, 2, 3, 4, 5, 6, 7}

	for {
		dci.Cmd(cmd)
		dci.WriteBytes(cmd)
		dci.End()
	}

	return
	// Run

	disp := displays.Adafruit_2i8_240x320_TFT_ILI9341(dci)

	//disp := displays.Adafruit_0i96_128x64_OLED_SSD1306(dci)
	//disp := displays.Adafruit_1i5_128x128_OLED_SSD1351(dci)
	//disp := displays.Adafruit_1i54_240x240_IPS_ST7789(dci)
	//disp := displays.ERTFTM_1i54_240x240_IPS_ST7789(dci)
	//disp := displays.MSP4022_4i0_320x480_TFT_ILI9486(dci)
	//disp := displays.Waveshare_1i5_128x128_OLED_SSD1351(dci)

	for {
		examples.RotateDisplay(disp)
		examples.DrawText(disp)
		examples.GraphicsTest(disp)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
