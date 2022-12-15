// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Blinky flashes the on-board LED.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/p/src"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
)

func main() {
	if src.SRC().SRSR.Load() == 3 {
		for {

			time.Sleep(time.Second / 8)
		}
	}
	for {
		leds.User.SetOn()
		time.Sleep(50 * time.Millisecond)
		leds.User.SetOff()
		time.Sleep(950 * time.Millisecond)
	}
}
