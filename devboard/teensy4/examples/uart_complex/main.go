// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Uart demonstrates how to use LPUART peripheral.
package main

import (
	"embedded/rtos"
	"fmt"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var u *lpuart.Driver

func main() {
	// Used IO pins
	tx := pins.P24
	rx := pins.P25

	// Setup LPUART driver
	u = lpuart.NewDriver(lpuart.LPUART(1), dma.Channel{}, dma.Channel{})
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(rx, lpuart.RXD)
	u.UsePin(tx, lpuart.TXD)
	irq.LPUART1.Enable(rtos.IntPrioLow, 0)

	// Enable both directions
	u.EnableRx(64) // use a 64-character ring buffer
	u.EnableTx()

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
	leds.User.Toggle() // visualize UART interrupts
	u.ISR()
}
