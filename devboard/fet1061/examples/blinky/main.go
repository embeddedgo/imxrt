// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to flash the onboard LED.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/hal/system/timer/systick"
)

func main() {
	system.Setup528_FlexSPI()
	systick.Setup(2e6)

	ledPin := iomux.AD_B0_09
	ledPin.Setup(iomux.Drive7)

	led := gpio.UsePin(ledPin, true)
	led.SetDirOut(true)

	for {
		led.Clear()
		time.Sleep(50 * time.Millisecond)
		led.Set()
		time.Sleep(950 * time.Millisecond)

		led.Toggle()
		time.Sleep(500 * time.Millisecond)
	}
}
