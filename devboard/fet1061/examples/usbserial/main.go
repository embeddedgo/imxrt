// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/ocotp"
	"github.com/embeddedgo/imxrt/p/pmu"
	"github.com/embeddedgo/imxrt/p/usb"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// IO pins
	conRx := pins.P23
	conTx := pins.P24

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	leds.User.SetOn()
	time.Sleep(20 * time.Millisecond) // wait at least 20ms before starting USB

	fmt.Println("Start!")
	initUSB()

	for {
		leds.User.Toggle()
		time.Sleep(time.Second)
	}
}

func initUSB() {

	CCM := ccm.CCM()
	CCMA := ccm_analog.CCM_ANALOG()
	PMU := pmu.PMU()
	OCOTP := ocotp.OCOTP()

	// Ungate all necessary clocks.
	CCM.CCGR2.SetBits(ccm.CG2_6)              // OCOTP
	CCM.CCGR6.SetBits(ccm.CG6_0 | ccm.CG6_11) // USB (usboh3) | CCMA (anadig)

	u := usb.USB1()

	fmt.Printf("PLL_ARM: %032b\n", CCMA.PLL_ARM.Load())
	fmt.Printf("PLL_USB1: %032b\n", CCMA.PLL_USB1.Load())
	fmt.Printf("PLL_USB2: %032b\n", CCMA.PLL_USB2.Load())
	fmt.Printf("MAC: %08x %08x %08x\n", OCOTP.MAC[0], OCOTP.MAC[1], OCOTP.MAC[2])

	fmt.Printf("USB ID: %08x\n", u.ID.Load())

	// Enable internal 3V0 regulator
	const (
		out3v000 = 15 << pmu.OUTPUT_TRGn
		boo0v150 = 6 << pmu.BO_OFFSETn
	)
	PMU.REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)
}
