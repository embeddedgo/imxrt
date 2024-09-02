// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPSPI3 driver
	p := lpi2c.LPI2C(1)
	d := lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	d.UsePin(scl, lpi2c.SCL)
	d.UsePin(sda, lpi2c.SDA)
	d.Setup(lpi2c.Fast)

	for {
		fmt.Printf("VERID: %#x\n", p.VERID.Load())
		fmt.Printf("PARAM: %#x\n", p.PARAM.Load())
		time.Sleep(5 * time.Second)
	}
}
