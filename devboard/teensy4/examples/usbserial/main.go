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
	"github.com/embeddedgo/imxrt/hal/usb/usbserial"
)

const hello = "\nHello World from i.MX RT!\nabcdefghijklmnoprstuvwxyz\nABCDEFGHIJKLMNOPRSTUVWXYZ\n!@#$%^&*(){}_-+=><\n\n"

var usbd *usb.Device

func main() {
	// IO pins
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	fmt.Println("Start!")

	const (
		interf = 0
		in     = 2 // input endpoint (host perspective), device Tx
		out    = 2 // output endopint (host prespective), device Rx
		maxPkt = 512
	)

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	se := usbserial.NewDriver(usbd, interf, out, in, maxPkt)
	se.SetWriteSink(true)
	se.SetAutoFlush(true)
	usbd.Enable()

	for i := 0; ; i++ {
		_, err := se.WriteString(hello) // USB DMA directly from Flash
		if err != nil {
			fmt.Println("WriteString", err)
			time.Sleep(time.Second)
		}
		_, err = fmt.Fprintln(se, len(hello), i)
		if err != nil {
			fmt.Println("Fprintln:", err)
			time.Sleep(time.Second)
		}
	}
}

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
