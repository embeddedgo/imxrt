// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Usbserial works as en Echo Server on the second USB serial port, writting
// back all received data. Additionally, it logs all I/O transaction on the
// system console (first USB serial port). You can talk to it using a terminal
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
// See also the usb example which does the same but uses the usb package
// directly.
package main

import (
	"fmt"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/system"
	"github.com/embeddedgo/imxrt/hal/usb"
	"github.com/embeddedgo/imxrt/hal/usb/usbserial"
)

func main() {
	usbd, rx, tx, interf := system.USBIO()
	usbd.Disable()
	se := usbserial.NewDriver(usbd, interf, rx, tx, 512)
	usbd.Enable()

	var buf [512]byte

usbNotReady:
	usbd.WaitConfig(0)

	for {
		n, err := se.Read(buf[:])
		if err != nil {
			if e, ok := err.(*usb.Error); ok && e.NotReady() {
				goto usbNotReady
			}
			fmt.Println("read error:", err)
			continue
		}
		fmt.Println("usbserial recv:", n)

		_, err = se.Write(buf[:n])
		if err != nil {
			if e, ok := err.(*usb.Error); ok && e.NotReady() {
				goto usbNotReady
			}
			fmt.Println("write error:", err)
		}
		fmt.Println("usbserial sent:", n)
	}
}
