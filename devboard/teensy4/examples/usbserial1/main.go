// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/hal/usb"

	"github.com/embeddedgo/imxrt/p/pmu"
	pusb "github.com/embeddedgo/imxrt/p/usb"
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
	usbd.Init(rtos.IntPrioLow, descriptors, true)
	usbd.Enable()

	fmt.Println("USB enabled. Waiting 2s...")

	time.Sleep(2 * time.Second)

	fmt.Println("Go!")

	/*const txt = "Hello, Wolrd!"
	  txtd := usb.NewDTD()
	  txtd.SetupTransfer(unsafe.Pointer(unsafe.StringData(txt)), len(txt))
	  usbd.Print(4*2 + 1)
	  txtd.Print()
	  usbd.Prime(4*2+1, txtd)*/

	var note rtos.Note
	buf0 := dma.MakeSlice[byte](512, 512)
	rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf0[0]), len(buf0))
	buf1 := dma.MakeSlice[byte](512, 512)
	rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf1[0]), len(buf1))
	rxtd0 := usb.NewDTD()
	rxtd0.SetNote(&note)
	rxtd0.SetupTransfer(unsafe.Pointer(&buf0[0]), len(buf0))
	rxtd1 := usb.NewDTD()
	rxtd1.SetNote(&note)
	rxtd1.SetupTransfer(unsafe.Pointer(&buf1[0]), len(buf1))
	rxtd0.SetNext(rxtd1)
	rxtd0.Print()
	rxtd1.Print()
	usbd.Print(3 * 2)
	usbd.Prime(3*2, rxtd0)

	pu := pusb.USB1()
	for {
		fmt.Printf(
			"eprime: %#x estat: %#x ecomplt: %#x nak: %#x\n",
			pu.ENDPTPRIME.Load(), pu.ENDPTSTAT.Load(), pu.ENDPTCOMPLETE.Load(),
			pu.ENDPTNAK.Load(),
		)
		//usbd.Print(4*2 + 1)
		//txtd.Print()
		usbd.Print(3 * 2)
		rxtd0.Print()
		rxtd1.Print()
		time.Sleep(5 * time.Second)
	}
}

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
