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
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/lpuart"
)

func dividers(clk, baud int) (osr, sbr int) {
	lowestE := 1<<31 - 1
	for o := 32; o >= 4; o-- {
		bo := baud * o
		minS := clk / bo
		// check s = minS and s = minS + 1
		for s := minS; ; s++ {
			e := clk - bo*s
			if e < 0 {
				e = -e
			}
			if e < lowestE {
				lowestE = e
				osr = o
				sbr = s
				if e == 0 {
					return
				}
			}
			if s != minS {
				break
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

	ccm.CCM().CCGR5.StoreBits(ccm.CG5_12, 3<<ccm.CG5_12n) // enable in all modes

	const uartClkRoot = 80e6

	osr, sbr := dividers(uartClkRoot, 115200)
	var baud lpuart.BAUD
	if osr < 8 {
		baud = lpuart.BOTHEDGE

	}
	baud |= lpuart.BAUD((osr-1)<<lpuart.OSRn | sbr)

	u := lpuart.LPUART1()
	u.BAUD.Store(baud)
	u.CTRL.Store(lpuart.RE | lpuart.TE | lpuart.DOZEEN)
	//u.FIFO.Store(lpuart.RXFE | lpuart.TXFE)

	for {
		var data lpuart.DATA
		for {
			data = u.DATA.Load()
			if data&^(lpuart.IDLINE|0xff) == 0 {
				break
			}
			u.STAT.Store(lpuart.OR) // clear possible Overrun flag
		}
		u.DATA.Store(data & 0xff)
		leds.User.Toggle()
	}
}
