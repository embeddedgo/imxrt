// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpspi"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func main() {
	// Used IO pins
	csn := pins.P10  // B0_00
	mosi := pins.P11 // B0_02
	miso := pins.P12 // B0_01
	sck := pins.P13  // B0_03

	// Setup LPSPI3 driver
	spi := lpspi.NewMaster(lpspi.LPSPI(3), dma.Channel{}, dma.Channel{})
	spi.UsePin(csn, lpspi.PCS0)
	spi.UsePin(mosi, lpspi.SDO)
	spi.UsePin(miso, lpspi.SDI)
	spi.UsePin(sck, lpspi.SCK)

}
