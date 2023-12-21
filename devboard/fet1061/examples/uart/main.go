// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Uart demonstrates how to use LPUART peripheral.
package main

import (
	"fmt"

	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// Used IO pins
	rx := pins.P23
	tx := pins.P24

	// Setup LPUART driver
	u := lpuart1.Driver()
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(rx, lpuart.RXD)
	u.UsePin(tx, lpuart.TXD)

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
