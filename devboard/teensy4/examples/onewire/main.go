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
	"io"
	"time"

	"github.com/embeddedgo/imxrt/dci/owdci"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
	"github.com/embeddedgo/onewire"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
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
	ow.UsePin(owTx, lpuart.TXD)
	ow.UsePin(owRx, lpuart.RXD)
	// Override UsePin settings
	pullUp22k := iomux.PK | iomux.Pull | iomux.Up22k
	owRx.Setup(pullUp22k)
	owTx.Setup(iomux.Drive2 | iomux.OpenDrain | pullUp22k)

	owm := onewire.Master{owdci.SetupLPUART(ow)}

	dtypes := []onewire.Type{onewire.DS18S20, onewire.DS18B20, onewire.DS1822}

	// This algorithm configures and starts conversion simultaneously on all
	// temperature sensors on the bus. It is fast but doesn't work in case of
	// if the parasite power bus configuration is used.

start:
	for {
		fmt.Fprint(con, "\r\nConfigure all DS18x20, DS1822 to 10bit resolution: ")
		if printErr(con, owm.SkipROM()) {
			continue start
		}
		if printErr(con, owm.WriteScratchpad(127, -128, onewire.T12bit)) {
			continue start
		}
		printOK(con)

		fmt.Fprint(con, "Sending ConvertT command (SkipROM addressing): ")
		if printErr(con, owm.SkipROM()) {
			continue start
		}
		if printErr(con, owm.ConvertT()) {
			continue start
		}
		printOK(con)

		fmt.Fprint(con, "Waiting until all devices finish the conversion: ")
		for {
			time.Sleep(50 * time.Millisecond)
			b, err := owm.ReadBit()
			if printErr(con, err) {
				continue start
			}
			fmt.Fprint(con, ". ")
			if b != 0 {
				printOK(con)
				break
			}
		}
		fmt.Fprint(con, "Searching for temperature sensors: ")
		for _, typ := range dtypes {
			s := onewire.NewSearch(typ, false)
			for owm.SearchNext(s) {
				d := s.Dev()
				fmt.Fprintf(con, "\r\n %v: ", d)
				if printErr(con, owm.MatchROM(d)) {
					continue start
				}
				s, err := owm.ReadScratchpad()
				if printErr(con, err) {
					continue start
				}
				t, err := s.Temp(typ)
				if printErr(con, err) {
					continue start
				}
				fmt.Fprintf(con, "%6.2f °C", t)
			}
			if printErr(con, s.Err()) {
				continue start
			}
		}
		fmt.Fprint(con, "\r\nDone.\r\n\r\n")
		time.Sleep(2 * time.Second)
	}
}

func printErr(w io.Writer, err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintf(w, "Error: %v\r\n", err)
	time.Sleep(2 * time.Second)
	return true
}

func printOK(w io.Writer) {
	fmt.Fprintf(w, "OK.\r\n")
}
