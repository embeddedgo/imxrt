// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal/periph"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

type Signal int8

// Do not reorder PCS0, SCK, SIN, SOUT constants. The order reflects the
// sequence of select registers and is used by periph.AltFunc function.

const (
	PCS0 Signal = iota // chip select 0
	PCS1               // chip select 1 / host request
	PCS2               // chip select 2 / data 2
	PCS3               // chip select 3 / data 3
	SCK                // clock
	SDI                // MISO / MOSI / data 1
	SDO                // MOSI / MISO / data 0
)

// Pins return IO pins that can be used for singal sig.
func (p *Periph) Pins(sig Signal) []iomux.Pin {
	return periph.Pins(pins[:], alts[:], num(p)*7+int(sig))
}

// UsePin is a helper function that can be used to configure IO pins as required
// by LPUART peripheral. Only certain pins can be used (see datasheet). UsePin
// returns true on succes or false if it isn't possible to use a pin as a sig.
// See also Periph.Pins.
func (d *Master) UsePin(pin iomux.Pin, sig Signal) bool {
	af, sel, daisy := periph.AltFunc(pins[:], alts[:], num(d.p)*7+int(sig), pin)
	if af < 0 {
		return false
	}
	var cfg iomux.Config
	if sig != SDI {
		// TODO: support half duplex mode
		cfg = iomux.Drive2 // 75Ω @ 3.3V, 130Ω @ 1.8V
	}
	pin.SetAltFunc(af)
	pin.Setup(cfg)
	if sel >= 0 {
		iosel := (*[16]mmio.R32[int32])(unsafe.Pointer(daisyBase))
		iosel[sel].Store(int32(daisy))
	}
	return true
}

const daisyBase uintptr = 0x401F_84EC

var pins = [...]iomux.Pin{
	// LPSPI1
	/* PCS0 */ iomux.SD_B0_01, iomux.EMC_30,
	/* PCS1 */ iomux.EMC_31,
	/* PCS2 */ iomux.EMC_40,
	/* PCS3 */ iomux.EMC_41,
	/* SCK  */ iomux.EMC_27, iomux.SD_B0_00,
	/* SDI  */ iomux.EMC_29, iomux.SD_B0_03,
	/* SDO  */ iomux.EMC_28, iomux.SD_B0_02,

	// LPSPI2
	/* PCS0 */ iomux.SD_B1_06, iomux.EMC_01,
	/* PCS1 */ iomux.EMC_14,
	/* PCS2 */ iomux.SD_B1_10,
	/* PCS3 */ iomux.SD_B1_11,
	/* SCK  */ iomux.SD_B1_07, iomux.EMC_00,
	/* SDI  */ iomux.SD_B1_09, iomux.EMC_03,
	/* SDO  */ iomux.SD_B1_08, iomux.EMC_02,

	// LPSPI3
	/* PCS0 */ iomux.AD_B0_03, iomux.AD_B1_12,
	/* PCS1 */ iomux.AD_B0_04,
	/* PCS2 */ iomux.AD_B0_05,
	/* PCS3 */ iomux.AD_B0_06,
	/* SCK  */ iomux.AD_B0_00, iomux.AD_B1_15,
	/* SDI  */ iomux.AD_B0_02, iomux.AD_B1_13,
	/* SDO  */ iomux.AD_B0_01, iomux.AD_B1_14,

	// LPSPI4
	/* PCS0 */ iomux.B0_00, iomux.B1_04,
	/* PCS1 */ iomux.B1_03,
	/* PCS2 */ iomux.B1_02,
	/* PCS3 */ iomux.B1_11,
	/* SCK  */ iomux.B0_03, iomux.B1_07,
	/* SDI  */ iomux.B0_01, iomux.B1_05,
	/* SDO  */ iomux.B0_02, iomux.B1_06,
}

const (
	// no select register (one level of I/O muxing)
	_1 = 0x10 // only one pin for the signal exists
	_2 = 0x20 // two alternative pins for the signal exist

	// select register exists (two levels of I/O muxing)
	s1 = -1<<7 + _1
	s2 = -1<<7 + _2
)

var alts = [...]iomux.AltFunc{
	// LPSPI1
	s2 + iomux.ALT4, iomux.ALT3,
	_1 + iomux.ALT3,
	_1 + iomux.ALT2,
	_1 + iomux.ALT2,
	s2 + iomux.ALT3, iomux.ALT4,
	s2 + iomux.ALT3, iomux.ALT4,
	s2 + iomux.ALT3, iomux.ALT4,

	// LPSPI2
	s2 + iomux.ALT4, iomux.ALT2,
	_1 + iomux.ALT4,
	_1 + iomux.ALT4,
	_1 + iomux.ALT4,
	s2 + iomux.ALT4, iomux.ALT2,
	s2 + iomux.ALT4, iomux.ALT2,
	s2 + iomux.ALT4, iomux.ALT2,

	// LPSPI3
	s2 + iomux.ALT7, iomux.ALT2,
	_1 + iomux.ALT7,
	_1 + iomux.ALT7,
	_1 + iomux.ALT7,
	s2 + iomux.ALT7, iomux.ALT2,
	s2 + iomux.ALT7, iomux.ALT2,
	s2 + iomux.ALT7, iomux.ALT2,

	// LPSPI4
	s2 + iomux.ALT3, iomux.ALT1,
	_1 + iomux.ALT2,
	_1 + iomux.ALT2,
	_1 + iomux.ALT6,
	s2 + iomux.ALT3, iomux.ALT1,
	s2 + iomux.ALT3, iomux.ALT1,
	s2 + iomux.ALT3, iomux.ALT1,
}
