// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// THIS CODE DOES NOT WORK YET!
//
// WORK IN PROGRESS...
//
// Uart demonstrates how to use LPUART peripheral.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/hal/iomux"

	"github.com/embeddedgo/imxrt/p/lpuart"
)

const (
	uartClkDiv = 1
)

func main() {
	tx := iomux.AD_B0_12
	rx := iomux.AD_B0_13

	tx.Setup(iomux.Drive2)
	tx.SetAltFunc(iomux.ALT2)
	rx.Setup(0)
	rx.SetAltFunc(iomux.ALT2)

	// uartClkHz := 480e6 / 6 / uartClkDiv = 80e6
	// br = uartClkHz / ((OSR+1) * SBR)
	// SBR = 80e6 / (16 * br)

	div := uint32(16 * 9600)
	sbr := uint32(80e6+div/2) / div

	u := lpuart.LPUART1()
	u.BAUD.U32.Store(sbr)
	u.CTRL.Store(lpuart.RE | lpuart.TE)

	for {
		var data uint32
		for {
			data = u.DATA.U32.Load()
			if data&^0xff == 0 {
				break
			}
		}
		u.DATA.U32.Store(data)
		leds.User.SetOn()
		time.Sleep(10 * time.Millisecond)
		leds.User.SetOff()
	}
}
