// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/hal/usb"

	"github.com/embeddedgo/imxrt/p/pmu"
)

var usbd *usb.Device

func main() {
	// IO pins
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	time.Sleep(20 * time.Millisecond) // wait at least 20ms before starting USB

	fmt.Println("Start!")

	// Enable internal 3V0 regulator
	const (
		out3v000 = 15 << pmu.OUTPUT_TRGn
		boo0v150 = 6 << pmu.BO_OFFSETn
	)
	pmu.PMU().REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	usbd.Enable()

	fmt.Println("USB enabled.")
	for {
		time.Sleep(time.Second)
	}
}

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
