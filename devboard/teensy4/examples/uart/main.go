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
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var (
	u    *lpuart.Driver
	note rtos.Note
)

func main() {
	// Use pins
	tx := pins.P24
	rx := pins.P25

	// Configure pins
	tx.Setup(iomux.Drive2)
	tx.SetAltFunc(iomux.ALT2)
	rx.Setup(0)
	rx.SetAltFunc(iomux.ALT2)

	// Setup LPUART driver
	u = lpuart.NewDriver(lpuart.LPUART(1), dma.Channel{}, dma.Channel{})
	u.Setup(lpuart.Word8b, 115200)
	irq.LPUART1.Enable(rtos.IntPrioLow, 0)

	// Enable both directions
	u.EnableRx(64) // use 64 byte ring buffer
	u.EnableTx()

	// Print received data showing reading chunks
	buf := make([]byte, 32)
	for {
		n, err := u.Read(buf)
		for err != nil {
			fmt.Fprintf(u, "error: %v\r\n", err)
			continue
		}
		fmt.Fprintf(u, "%d: %s\r\n", n, buf[:n])
	}
}

//go:interrupthandler
func LPUART1_Handler() {
	leds.User.Toggle() // toggle onboard LED to see interrupts
	u.ISR()
}
