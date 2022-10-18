// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Onewire uses SkipROM addressing to configure and run all DS18x2x temperature
// sensors on the 1-Wire bus conected to the LPUART2 interface. Next it searchs
// for individual sensors and prints the mesured temperatures on the serial
// terimnal connected to the LPUART1 interface.
package main

import (
	"fmt"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
)

func main() {
	// IO pins
	owTx := pins.P14
	owRx := pins.P15
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	con := lpuart1.Driver()
	con.Setup(lpuart.Word8b, 115200)
	con.UsePin(conRx, lpuart.RXD)
	con.UsePin(conTx, lpuart.TXD)
	con.EnableRx(64)
	con.EnableTx()

	// 1-Wire driver
	ow := lpuart2.Driver()
	ow.Setup(lpuart.Word8b, 115200)
	//ow.Periph().CTRL.SetBits(lpuart.LOOPS)
	ow.UsePin(owTx, lpuart.TXD)
	ow.UsePin(owRx, lpuart.RXD)
	//owTx.Setup(iomux.Drive2 | iomux.OpenDrain | iomux.PK | iomux.Pull | iomux.Up22K)
	//owRx.Setup(iomux.Pull | iomux.PK| iomux.Up22K)
	ow.EnableRx(64)
	ow.EnableTx()

	//owdci.SetupLPUART(ow)

	//owRx.Setup(iomux.Drive2)
	//bit := gpio.UsePin(owRx, false)
	//bit.Port().EnableClock(true)
	//bit.SetDirOut(true)

	c := byte('a')
	for {
		fmt.Fprintf(con, "gen: %c -> ow\r\n", c)
		ow.WriteByte(c)
		b, err := ow.ReadByte()
		fmt.Fprintf(con, "read: ow -> %c %v\r\n", b, err)
		if c++; c > 'z' {
			c = 'a'
		}
		//time.Sleep(time.Second)
	}
}
