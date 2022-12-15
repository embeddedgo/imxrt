// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Onewire uses SkipROM addressing to configure and run all DS18x2x temperature
// sensors on the 1-Wire bus conected to the LPUART2 interface. Next it searchs
// for individual sensors and prints the mesured temperatures on the serial
// terimnal connected to the LPUART1 interface.
package main

import (
	"time"

	"github.com/embeddedgo/imxrt/p/src"

	"github.com/embeddedgo/imxrt/dci/owdci"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/leds"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
	"github.com/embeddedgo/onewire"
)

func main() {
	// IO pins
	owTx := pins.P12
	owRx := pins.P13

	if src.SRC().SRSR.Load() == 3 {
		for {
			leds.User.Toggle()
			time.Sleep(time.Second / 4)
		}
	}

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
		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.WriteScratchpad(127, -128, onewire.T12bit)) {
			continue start
		}

		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.ConvertT()) {
			continue start
		}

		for {
			time.Sleep(50 * time.Millisecond)
			b, err := owm.ReadBit()
			if printErr(err) {
				continue start
			}
			if b != 0 {
				break
			}
		}
		for _, typ := range dtypes {
			s := onewire.NewSearch(typ, false)
			for owm.SearchNext(s) {
				d := s.Dev()
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
				leds.User.Toggle()
				time.Sleep(time.Second)
				leds.User.Toggle()
			}
			if printErr(s.Err()) {
				continue start
			}
		}
	}
}

func printErr(err error) bool {
	if err == nil {
		return false
	}
	for i := 0; i < 4; i++ {
		leds.User.Toggle()
		time.Sleep(time.Second / 8)
	}
	return true
}
