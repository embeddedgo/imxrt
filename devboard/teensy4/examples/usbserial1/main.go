// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/dma"
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
	if false {
		pmu.PMU().REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)
	}

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	usbd.Enable()

	var note rtos.Note
	config := uint8(1)
	rxe := 2 * 2
	rxtd := usb.NewDTD()
	rxtd.SetNote(&note)
	txe := 2*2 + 1
	txtd := usb.NewDTD()
	txtd.SetNote(&note)
	buf := dma.MakeSlice[byte](512, 512)

usbNotReady:
	fmt.Println("Waiting for USB...")
	usbd.WaitConfig(config)
	fmt.Println("USB is ready.")

	for {
		var (
			n    int
			stat uint8
		)
		rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf[0]), len(buf))
		rxtd.SetupTransfer(unsafe.Pointer(&buf[0]), len(buf))
		note.Clear()

		if !usbd.Prime(rxe, rxtd, config) {
			goto usbNotReady
		}
		note.Sleep(-1)
		usbd.Clean(rxe)

		n, stat = rxtd.Status()
		if stat != 0 {
			if stat&usb.Active != 0 {
				goto usbNotReady
			}
			fmt.Printf("Rx error: 0b%08b\n", stat)
			time.Sleep(time.Second)
			continue
		}

		n = len(buf) - n
		fmt.Printf("received %d bytes: %s\n", n, buf[:n])

		if strings.TrimSpace(string(buf[:n])) == "reset" {
			fmt.Println("* Reset! *")
			usbd.Disable()
			time.Sleep(time.Second)
			usbd.Enable()
			fmt.Println("* Go! *")
		}

		txtd.SetupTransfer(unsafe.Pointer(&buf[0]), n)
		note.Clear()

		if !usbd.Prime(txe, txtd, config) {
			goto usbNotReady
		}
		note.Sleep(-1)
		usbd.Clean(txe)

		_, stat = txtd.Status()
		if stat != 0 {
			if stat&usb.Active != 0 {
				goto usbNotReady
			}
			fmt.Printf("Tx error: 0b%08b\n", stat)
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("sent %d bytes\n", n)
	}
}

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
