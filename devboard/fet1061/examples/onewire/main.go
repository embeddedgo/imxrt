// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Onewire uses SkipROM (broadcast) addressing to configure and run all DS18x2x
// temperature sensors on the 1-Wire bus conected to the LPUART2 interface. Next
// it waits for all sensors to finish the conversion. After that is searchs for
// individual sensors on the bus and communicates with them using conventional
// MatchROM (unicast) adressing mode, reads and and prints the mesured
// temperatures on the serial terimnal connected to the LPUART1 interface.
//
// Starting conversion on many sensors at the same time requires much power so
// it may not work in parasite power bus configuration.
package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/dci/owdci"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/onewire"

	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
)

func main() {
	// IO pins
	owTx := pins.P12
	owRx := pins.P13
	conRx := pins.P23
	conTx := pins.P24

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	// 1-Wire driver
	ow := lpuart2.Driver()
	ow.UsePin(owTx, lpuart.TXD)
	ow.UsePin(owRx, lpuart.RXD)
	// Override UsePin settings
	pullUp22k := iomux.PK | iomux.Pull | iomux.Up22k
	owRx.Setup(pullUp22k)
	owTx.Setup(iomux.Drive2 | iomux.OpenDrain | pullUp22k)

	owm := onewire.Master{owdci.SetupLPUART(ow)}

	dtypes := []onewire.Type{onewire.DS18S20, onewire.DS18B20, onewire.DS1822}

start:
	for {
		fmt.Print("\nConfigure all DS18x20, DS1822 to 10bit resolution: ")
		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.WriteScratchpad(127, -128, onewire.T12bit)) {
			continue start
		}
		fmt.Println("OK.")

		fmt.Print("Sending ConvertT command (SkipROM addressing): ")
		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.ConvertT()) {
			continue start
		}
		fmt.Println("OK.")

		fmt.Print("Waiting until all devices finish the conversion: ")
		for {
			time.Sleep(50 * time.Millisecond)
			b, err := owm.ReadBit()
			if printErr(err) {
				continue start
			}
			fmt.Print(". ")
			if b != 0 {
				fmt.Println("OK.")
				break
			}
		}
		fmt.Print("Searching for temperature sensors: ")
		for _, typ := range dtypes {
			s := onewire.NewSearch(typ, false)
			for owm.SearchNext(s) {
				d := s.Dev()
				fmt.Printf("\n %v: ", d)
				if printErr(owm.MatchROM(d)) {
					continue start
				}
				s, err := owm.ReadScratchpad()
				if printErr(err) {
					continue start
				}
				t, err := s.Temp(typ)
				if printErr(err) {
					continue start
				}
				_ = t
				fmt.Printf("%6.2f °C", t)
			}
			if printErr(s.Err()) {
				continue start
			}
		}
		fmt.Print("\nDone.\n\n")
		time.Sleep(2 * time.Second)
	}
}

func printErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Printf("Error: %v\n", err)
	time.Sleep(2 * time.Second)
	return true
}
