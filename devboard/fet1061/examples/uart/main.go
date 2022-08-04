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

	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/lpuart"

	"github.com/embeddedgo/imxrt/hal/iomux"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
)

func dividers(clk, baud uint) (osr, sbr int) {
	div := int((clk + baud/2) / baud)
	le := 1<<31 - 1
	for o := 32; o >= 4; o-- {
		for s := 1; s <= 8191; s++ {
			e := div - o*s
			if e < 0 {
				e = -e
			}
			if e < le {
				le = e
				osr = o
				sbr = s
				if e == 0 {
					return
				}
			}
		}
	}
	return
}

func main() {
	tx := iomux.AD_B0_12
	rx := iomux.AD_B0_13

	tx.Setup(iomux.Drive2)
	tx.SetAltFunc(iomux.ALT2)
	rx.Setup(0)
	rx.SetAltFunc(iomux.ALT2)

	CCM := ccm.CCM()
	CCM.CSCDR1.StoreBits(ccm.UART_CLK_PODF, 0<<ccm.UART_CLK_PODFn)
	CCM.CCGR5.StoreBits(ccm.CG5_12, 3<<ccm.CG5_12n) // enable in all modes

	// UART_CLK_ROOT = 480e6 / 6 / (UART_CLK_PODF+1) = 80e6

	var baud lpuart.BAUD
	osr, sbr := dividers(80e6, 9600)
	if osr < 8 {
		baud = lpuart.BOTHEDGE
	}
	baud |= lpuart.BAUD((osr-1)<<lpuart.OSRn | sbr)

	u := lpuart.LPUART1()
	u.BAUD.Store(baud)
	u.CTRL.Store(lpuart.RE | lpuart.TE | lpuart.DOZEEN)

	for {
		var data lpuart.DATA
		for {
			data = u.DATA.Load()
			if data&^(lpuart.IDLINE|0x3ff) == 0 {
				break
			}
		}
		u.DATA.Store(data & 0xff)
		leds.User.SetOn()
		time.Sleep(10 * time.Millisecond)
		leds.User.SetOff()
	}
}
