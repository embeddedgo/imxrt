// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Usb works as en Echo Server on the second USB serial port, writting back all
// received data. Additionally, it logs all I/O transaction on the system
// console (first USB serial port). You can talk to it using a terminal
// emulator program like Putty (Windows) or picocom (Linux, Mac):
//
//	# Show console logs:
//	cat /dev/ttyACM0
//
//	# Talk to this program:
//	picocom --imap crcrlf /dev/ttyACM1
//
// Try copy/paste more text to the terminal emulator window to see slightly
// longer than single-byte transactions.
//
// See also the usbserial example which does the same in much simpler way, using
// the usbserial package.
package main

import (
	"embedded/rtos"
	"fmt"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/usb"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/system"
)

// rtos.CacheMaint is almost always required when you transfer data through USB.
// But in fact, this code requires it only once (after allocating buf), because
// in this example the CPU doesn't read/modify the received data. Nevertheless,
// both cache maintenance operations (before sending and before receiving) are
// presented for completeness.

func main() {
	usbd, out, in, _ := system.USBIO()
	rxe := usb.HE(out, usb.OUT)
	txe := usb.HE(in, usb.IN)

	buf := dma.MakeSlice[byte](512, 512)
	done := new(rtos.Note)
	td := usb.NewDTD()
	td.SetNote(done)

usbNotReady:
	usbd.WaitConfig(0)

	for {
		rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf[0]), len(buf))
		td.SetupTransfer(unsafe.Pointer(&buf[0]), len(buf))
		done.Clear()

		if !usbd.Prime(rxe, td, td) {
			goto usbNotReady
		}
		done.Sleep(-1)

		n, stat := td.Status()
		if stat != 0 {
			if stat&usb.Active != 0 {
				goto usbNotReady
			}
			fmt.Printf("read error: 0b%08b\n", stat)
			continue
		}
		n = len(buf) - n
		fmt.Println("usb recv:", n)

		rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(&buf[0]), alignUp(n))
		td.SetupTransfer(unsafe.Pointer(&buf[0]), n)
		done.Clear()

		if !usbd.Prime(txe, td, td) {
			goto usbNotReady
		}
		done.Sleep(-1)

		_, stat = td.Status()
		if stat != 0 {
			if stat&usb.Active != 0 {
				goto usbNotReady
			}
			fmt.Printf("write error: 0b%08b\n", stat)
			continue
		}
		fmt.Println("usb sent:", n)
	}
}

func alignUp(n int) int {
	const align = dma.MemAlign - 1
	return (n + align) &^ align
}
