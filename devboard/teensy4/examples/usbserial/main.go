// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/imxrt/p/ccm"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
)

func main() {
		leds.User.SetOn()
		time.Sleep(20 * time.Millisecond) // wait at least 20ms before starting USB
		initUSB()
}


func initUSB() {
	CCM := ccm.CCM()

	CCMA.PFD_480

	// Ungate USB clocks.
	CCM.CCGR6.Set(ccm.CG6_0)
}