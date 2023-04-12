// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Blinky flashes the on-board LED.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/p/iomuxc_gpr"
)

func main() {
	delay := 200 * time.Millisecond
	gprs := iomuxc_gpr.IOMUXC_GPR()
	if gprs.GPR17.Load() != 0 {
		delay = 950 * time.Millisecond
	}
	for {
		leds.User.SetOn()
		time.Sleep(50 * time.Millisecond)
		leds.User.SetOff()
		time.Sleep(delay)
	}
}
