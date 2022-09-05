// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Uart-details demonstrates how to use LPUART peripheral using hal/lpuart
// package directly.
package main

import (
	"embedded/rtos"
	"fmt"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var u *lpuart.Driver

func main() {
	// Used IO pins
	rx := pins.P23
	tx := pins.P24

	// Setup DMA controller
	d := dma.DMA(0)
	d.EnableClock(true)
	d.CR.SetBits(dma.ERCA | dma.ERGA | dma.HOE) // round robin, halt on error

	// Setup LPUART driver
	u = lpuart.NewDriver(lpuart.LPUART(1), dma.Channel{}, d.Channel(0))
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(rx, lpuart.RXD)
	u.UsePin(tx, lpuart.TXD)
	irq.LPUART1.Enable(rtos.IntPrioLow, 0)
	irq.DMA0_DMA16.Enable(rtos.IntPrioLow, 0)

	// Enable both directions
	u.EnableRx(64) // use a 64-character ring buffer
	u.EnableTx()

	u.WriteString("abcdefghijklmnoprtsuvwxyz+01234567890-ABCDEFGHIJKLMNOPRSTUVWXYZ\r\n")

	// Print received data showing reading chunks
	buf := make([]byte, 80)
	for {
		n, err := u.Read(buf)
		if err != nil {
			fmt.Fprintf(u, "error: %v\r\n", err)
			continue
		}
		fmt.Fprintf(u, "%d: %s\r\n", n, buf[:n])
	}
}

//go:interrupthandler
func LPUART1_Handler() {
	u.ISR()
}

//go:interrupthandler()
func DMA0_DMA16_Handler() {
	u.TxDMAISR()
	leds.User.Toggle() // visualize DMA interrupts
}
