// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dma/dmairq"
	"github.com/embeddedgo/imxrt/hal/lpspi"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// Used IO pins
	conRx := pins.P23
	conTx := pins.P24
	miso := pins.P91 // AD_B1_13
	mosi := pins.P92 // AD_B1_14
	csn := pins.P93  // AD_B1_12
	sck := pins.P94  // AD_B1_15

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
	dmairq.SetISR(rxdma, spi.RxDMAISR)
	spi.Setup(lpspi.FD, 19e6)
	spi.Enable()

	fmt.Println("*** Start ***")

	// CPOL0,CPHA=0,19MHz/2=9.5MHz,PCS0,MSBF,1bit
	spi.WriteCmd(lpspi.PREDIV2, 8)

	in := make([]byte, 1e5)
	out := make([]byte, len(in))
	for i := range in {
		in[i] = byte(i)
	}
	for {
		t0 := time.Now()
		spi.WriteRead(out, in)
		t1 := time.Now()
		for i := range out {
			for out[i] != in[i] {
				fmt.Printf(
					"out[%d] != in[%d] (%d != %d)\n",
					i, i, out[i], in[i],
				)
			}
			out[i] = 0
		}
		fmt.Printf(
			"%.0f kB/s\n",
			float64(len(out))*float64(time.Second/1000)/float64(t1.Sub(t0)),
		)
	}
}
