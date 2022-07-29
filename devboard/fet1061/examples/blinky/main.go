// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to flash the onboard LED.
package main

import (
	"embedded/rtos"
	"runtime"
	"time"

	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/hal/system/timer/systick"

	"github.com/embeddedgo/imxrt/p/iomuxc_gpr"
)

func main() {
	system.Setup528_FlexSPI()
	systick.Setup(2e6)

	runtime.LockOSThread()
	privLevel, _ := rtos.SetPrivLevel(0)

	// GPIO_MUX1 ALT5 selects GPIO6 (fast) instead of GPIO1 (slow) for bit 9.
	iomuxc_gpr.IOMUXC_GPR().GPR26.Store(1 << 9)

	pad := iomux.AD_B0_09
	pad.Setup(iomux.Drive7)
	pad.SetAltFunc(iomux.ALT5) // use as GPIO6 pin 9

	rtos.SetPrivLevel(privLevel)
	runtime.UnlockOSThread()

	pin := gpio.P(6).Pin(9)
	pin.SetDirOut(true)
	for {
		//pin.Clear()
		//time.Sleep(50 * time.Millisecond)
		//pin.Set()
		//time.Sleep(950 * time.Millisecond)

		pin.Toggle()
		time.Sleep(500 * time.Millisecond)
	}
}
