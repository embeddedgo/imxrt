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
	"embedded/rtos"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var (
	u    *lpuart.Driver
	note rtos.Note
)

func main() {
	tx := iomux.AD_B0_12
	rx := iomux.AD_B0_13

	tx.Setup(iomux.Drive2)
	tx.SetAltFunc(iomux.ALT2)
	rx.Setup(0)
	rx.SetAltFunc(iomux.ALT2)

	u = lpuart.NewDriver(lpuart.LPUART(1), dma.Channel{}, dma.Channel{})
	u.Setup(lpuart.Word8b, 115200)
	u.Periph().FIFO.Store(lpuart.RXFE | lpuart.TXFE)
	irq.LPUART1.Enable(rtos.IntPrioLow, 0)

	u.EnableRx()
	u.EnableTx()
	for {
		wait(u.Periph(), lpuart.RDRF)
		data := u.Periph().DATA.Load()
		wait(u.Periph(), lpuart.TDRE)
		u.Periph().DATA.Store(data & 0xff)
		leds.User.Toggle()
	}
}

const (
	errMask  = lpuart.PF | lpuart.FE | lpuart.NF | lpuart.OR
	errIntEn = lpuart.CTRL(errMask) << (lpuart.PEIEn - lpuart.PFn)
)

func wait(u *lpuart.Periph, ev lpuart.STAT) (err lpuart.STAT) {
	note.Clear()
	u.CTRL.SetBits(lpuart.CTRL(ev) | errIntEn) // enable interrupts
	note.Sleep(-1)                             // wait for interrupt
	err = u.STAT.Load() & errMask              // check for errors
	u.STAT.Store(err)                          // clear known errors
	return
}

//go:interrupthandler
func LPUART1_Handler() {
	u.Periph().CTRL.ClearBits(lpuart.RIE | lpuart.TIE | errIntEn) // disable interrupts
	note.Wakeup()
}
