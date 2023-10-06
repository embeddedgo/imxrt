// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package usb provides a drivers for the i.MX RT built-in USB
// controllers.
//
// # Device Controller Driver (DCD)
//
// DCD is is primarily used by some the higher-level drivers (e.g. CDC ACM
// serial driver) but can be also used directly as in the following example:
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
//		d.SetNote(&done)
//		buf := dma.MakeSlice[byte](512, 512)
//
//	usbNotReady:
//		usbd.WaitConfig(config)
//
//		for {
//			rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf[0]), len(buf))
//			td.SetupTransfer(unsafe.Pointer(&buf[0]), len(buf))
//			done.Clear()
//
//			if !usbd.Prime(rxEndpoint, td, td, config) {
//				goto usbNotReady
//			}
//			done.Sleep(-1)
//
//			n, stat := rxtd.Status()
//			switch {
//			case stat == 0:
//				n = len(buf) - n
//				handleRxData(buf[:n])
//			case stat&usb.Active != 0:
//				goto usbNotReady
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
//
package usb
