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

	"github.com/embeddedgo/imxrt/dci/owdci"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
)

func main() {
	// IO pins
	owRxTx := pins.P12
	conRx := pins.P23
	conTx := pins.P24

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
	ow.UsePin(owRxTx, lpuart.TXD) // for AltFunc, pin config overrided below
	owRxTx.Setup(iomux.Drive2 | iomux.OpenDrain | iomux.Pull | iomux.Up22K)

	dci := owdci.SetupLPUART(ow)

	// Print received data showing reading chunks
	buf := make([]byte, 80)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Fprintf(con, "error: %v\r\n", err)
			continue
		}
		fmt.Fprintf(con, "%d: %s\r\n", n, buf[:n])
		for _, b := range buf[:n] {
			dci.WriteByte(b)
		}
	}
}
