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

	pad := iomux.AD_B0_09
	pad.Setup(iomux.Drive7)
	pad.SetAltFunc(iomux.GPIO)
	iomux.Lock()

	pin := gpio.P(1).Pin(9) // AD_B0 goes to GPIO port 1 (slow) or 6 (fast)
	pin.ConnectMux()
	pin.SetDirOut(true)

	for {
		pin.Clear()
		time.Sleep(50 * time.Millisecond)
		pin.Set()
		time.Sleep(950 * time.Millisecond)

		//pin.Toggle()
		//time.Sleep(500 * time.Millisecond)
	}
}
