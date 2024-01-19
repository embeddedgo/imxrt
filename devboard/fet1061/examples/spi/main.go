// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync/atomic"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpspi"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

var a atomic.Bool

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

	// Setup LPSPI3 driver
	spi := lpspi.NewMaster(lpspi.LPSPI(3), dma.Channel{}, dma.Channel{})
	spi.UsePin(miso, lpspi.SDI)
	spi.UsePin(mosi, lpspi.SDO)
	spi.UsePin(csn, lpspi.PCS0)
	spi.UsePin(sck, lpspi.SCK)

	spi.Setup(lpspi.FD, 66e6)
	spi.Enable()
	p := spi.Periph()

	/*
		p.EnableClock(true)
		p.CR.Store(lpspi.RRF | lpspi.RTF | lpspi.RST) // resert
		p.CR.Store(0)
		p.CFGR1.Store(lpspi.MASTER) // | lpspi.SAMPLE)
		p.CCR.Store(0x0409_0808)
		p.FCR.Store(3)
		p.CR.Store(lpspi.DBGEN | lpspi.MEN)
	*/

	// CPOL0,CPHA=0,66MHz/8=8.25MHz,PCS0,MSBF,1bit
	p.TCR.Store(lpspi.PREDIV8 | (32-1)<<lpspi.FRAMESZn)

	for {
		for p.SR.LoadBits(lpspi.TDF) == 0 {
		}
		p.TDR.Store(0x1234_5678)
		p.SR.Store(lpspi.TDF)
		for p.SR.LoadBits(lpspi.RDF) == 0 {
		}
		u32 := p.RDR.Load()
		p.SR.Store(lpspi.RDF)
		fmt.Printf("%08x\n", u32)
	}
}
