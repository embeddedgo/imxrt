// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package usb provides a high level access to the i.MX RT built-in USB
// controllers.
//
// # Logical vs hardware endpoint numbers.
//
// This package uses two ways to describe an USB endpoint.
//
// The Hardware Endpoint number (he) specifies an unidirectional communication
// channel available in the USB device. The supported direction is encoded in
// the least significant bit of the he (an even he means OUT endpoint, an odd he
// means IN endpoint).
//
// The Logical Endpoint number (le) is what the USB documentation means when
// it uses the phrase endpoint or endpoint number. In this package it is used
// mainly in the context of control endpoints due to their obligatory
// bidirectional nature. In fact, a logical endpoint without specifying its
// direction is not a precise term because it may point to two unrelated
// communication channels.
//
// This package uses uint8 for he and int8 for le. The connection between le and
// he is as follows: le = int8(he >> 1). The HE and LE functions can be used for
// readabble conversions.
//
// # Device Controller Driver (DCD)
//
// DCD is primarily intended to be used by the higher-level drivers (e.g. CDC
// ACM serial driver) but can be also used directly as in the following example:
//
//	var usbd *usb.Device
//
//	func main() {
//		usbd = usb.NewDevice(1)
//		usbd.Init(rtos.IntPrioLow, descriptors, false)
//		usbd.Enable()
//
//		var done rtos.Note
//		td := usb.NewDTD()
//		td.SetNote(&done)
//		buf := dma.MakeSlice[byte](512, 512)
//		config := 1
//
//	usbNotReady:
//		usbd.WaitConfig(config)
//
//		for {
//			bufp := unsafe.Pointer(&buf[0])
//			rtos.CacheMaint(rtos.DCacheInval, bufp, len(buf))
//			td.SetupTransfer(bufp, len(buf))
//			done.Clear()
//
//			if !usbd.Prime(rxhe, td, td, config) {
//				goto usbNotReady
//			}
//			done.Sleep(-1)
//
//			n, stat := td.Status()
//			switch {
//			case stat == 0:
//				n = len(buf) - n
//				handleRxData(buf[:n])
//
//			case stat&usb.Active != 0:
//				goto usbNotReady
//
//			default:
//				handleRxError(stat)
//			}
//		}
//	}
//
//	//go:interrupthandler
//	func USB_OTG1_Handler() {
//		usbd.ISR()
//	}
package usb
