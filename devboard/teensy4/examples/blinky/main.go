// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Blinky flashes the on-board LED.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
)

func main() {
	for {
		leds.User.SetOn()
		time.Sleep(50 * time.Millisecond)
		leds.User.SetOff()
		time.Sleep(950 * time.Millisecond)
	}
}
