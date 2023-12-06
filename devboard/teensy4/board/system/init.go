// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package system, when imported, performs the default configuration of the
// Teensy 4.x board.
//
// The board is configured to run with 528 MHz system clock. The USB interface
// is configured to provide two bidirectional communication channels with the
// host visible as two CDC ACM serial ports. The first one is used as the system
// console, the second one remains free for any use (see the USBIO function for
// more information).
//
// All packages in the board directory imports this package implicitly.
package system

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/hal/system/console/usbcon"
	"github.com/embeddedgo/imxrt/hal/system/timer/systick"
	"github.com/embeddedgo/imxrt/hal/usb"
	"github.com/embeddedgo/imxrt/hal/usb/usbserial"
)

var usbd *usb.Device

// USBIO returns the USB device and two logical endpoint numbers configured to
// be used as one bidirectional or two unidirectional communication channels.
// You can use them directly for raw USB communication, or they can be passed
// (along with statInt) to the usbserial.NewDriver function to obtain an
// emulated serial port.
//
// On the host side, this bidirectional communication interface is visible as
// the second CDC ACM serial port. Nevertheless, thanks to libraries like libusb
// it can be used in raw mode to implement any required communication protocol.
func USBIO() (d *usb.Device, outRx, inTx int8, statInt uint8) {
	return usbd, int8(acm1_DataOUT[0] & 15), int8(acm1_DataIN[0] & 15),
		acm1_StatusInt[0]
}

func init() {
	/*
		// Reconfigure the internal USB regulator.
		const (
			out3v000 = 15 << pmu.OUTPUT_TRGn
			boo0v150 = 6 << pmu.BO_OFFSETn
		)
		pmu.PMU().REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)
	*/

	system.Setup528_FlexSPI()
	systick.Setup(2e6)

	// USB console

	interf := acm0_StatusInt[0]
	in := int8(acm0_DataIN[0] & 15)   // input endpoint (host perspective)
	out := int8(acm0_DataOUT[0] & 15) // output endopint (host prespective)
	maxPkt := int(acmDataSize480[0])*256 + int(acmDataSize480[1])

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	se := usbserial.NewDriver(usbd, interf, out, in, maxPkt)
	se.SetWriteSink(true)
	se.SetAutoFlush(true)
	usbd.Enable()

	usbcon.Setup(se, "USBSERIAL")
}

//go:interrupthandler
func _USB_OTG1_Handler() { usbd.ISR() }

//go:linkname _USB_OTG1_Handler IRQ113_Handler
