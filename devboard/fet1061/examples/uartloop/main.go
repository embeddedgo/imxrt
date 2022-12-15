// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Uartloop can be used to test UART connection.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
	"github.com/embeddedgo/imxrt/p/src"
)

func main() {
	if src.SRC().SRSR.Load() == 3 {
		for i := 0; i < 60; i++ {
			leds.User.Toggle()
			time.Sleep(time.Second / 8)
		}
	}

	// Used IO pins
	rx := pins.P13
	tx := pins.P12

	// Setup LPUART driver
	u := lpuart2.Driver()
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(rx, lpuart.RXD)
	u.UsePin(tx, lpuart.TXD)

	// Enable both directions
	u.EnableRx(64) // use a 64-character ring buffer
	u.EnableTx()

	u.WriteString("0123456789")

	buf := make([]byte, 80)
	for {
		n, err := u.Read(buf)
		if err != nil {
			for i := 0; i < 4; i++ {
				leds.User.Toggle()
				time.Sleep(time.Second)
			}
			continue
		}
		u.WriteString(string(buf[:n]))
		leds.User.Toggle()
	}
}
