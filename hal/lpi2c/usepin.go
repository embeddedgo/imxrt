// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal/periph"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

type Signal int8

// Do not reorder SCL, SDA constants. The order reflects the
// sequence of daisy select registers and is used by periph.AltFunc function.

const (
	SCL  Signal = iota // clock
	SDA                // data
	SCLS               // secondary clock (only LPI2C1)
	SDAS               // secondary data (only LPI2C1)
	HREQ               // host request
)

// Pins return IO pins that can be used for singal sig.
func (p *Periph) Pins(sig Signal) []iomux.Pin {
	return periph.Pins(pins[:], alts[:], num(p)*5+int(sig))
}

// UsePin is a helper function that can be used to configure IO pins as required
// by LPI2C peripheral. It configures the pin as open-drain output with the
// internal 22 kΩ pull-up resistor enadbled so you can avoid external pull-up
// resistors for low-capacitance bus and low-speed transfers.
//
// Only certain pins can be used forLPI2C peripheral (see datasheet). UsePin
// returns true on succes or false if it isn't possible to use a pin as a sig.
func (d *Master) UsePin(pin iomux.Pin, sig Signal) bool {
	n := num(d.p)
	af, sel, daisy := periph.AltFunc(pins[:], alts[:], n*5+int(sig), pin)
	if af < 0 {
		return false
	}
	pin.SetAltFunc(af | iomux.SION)
	if sel >= 0 {
		iosel := (*[16]mmio.R32[int32])(unsafe.Pointer(daisyBase))
		iosel[sel].Store(int32(daisy))
	}
	pin.Setup(iomux.Drive2 | iomux.OpenDrain | iomux.Pull | iomux.Up22k | iomux.Hys)
	return true
}

const daisyBase uintptr = 0x401F_84CC

var pins = [...]iomux.Pin{
	// LPI2C1
	/* SCL  */ iomux.SD_B1_04, iomux.AD_B1_00,
	/* SDA  */ iomux.SD_B1_05, iomux.AD_B1_01,
	/* SCLS */ iomux.AD_B0_00,
	/* SDAS */ iomux.AD_B0_01,
	/* HREQ */ iomux.AD_B0_02,

	// LPI2C2
	/* SCL  */ iomux.SD_B1_11, iomux.B0_04,
	/* SDA  */ iomux.SD_B1_10, iomux.B0_05,
	/* SCLS */ -1,
	/* SDAS */ -1,
	/* HREQ */ -1,

	// LPI2C3
	/* SCL  */ iomux.EMC_22, iomux.SD_B0_00, iomux.AD_B1_07,
	/* SDA  */ iomux.EMC_21, iomux.SD_B0_01, iomux.AD_B1_06,
	/* SCLS */ -1,
	/* SDAS */ -1,
	/* HREQ */ -1,

	// LPI2C4
	/* SCL  */ iomux.EMC_12, iomux.AD_B0_12,
	/* SDA  */ iomux.EMC_11, iomux.AD_B0_13,
	/* SCLS */ -1,
	/* SDAS */ -1,
	/* HREQ */ -1,
}

const (
	// no select register (one level of I/O muxing)
	_0 = 0x00 // no pin for signal exists
	_1 = 0x10 // only one pin for the signal exists
	_2 = 0x20 // two alternative pins for the signal exist
	_3 = 0x30 // three alternative pins for the signal exist

	// select register exists (two levels of I/O muxing)
	s1 = -1<<7 + _1
	s2 = -1<<7 + _2
	s3 = -1<<7 + _3
)

var alts = [...]iomux.AltFunc{
	// LPI2C1
	s2 + iomux.ALT2, iomux.ALT3,
	s2 + iomux.ALT2, iomux.ALT3,
	_1 + iomux.ALT4,
	_1 + iomux.ALT4,
	_1 + iomux.ALT6,

	// LPI2C2
	s2 + iomux.ALT3, iomux.ALT2,
	s2 + iomux.ALT3, iomux.ALT2,
	_0,
	_0,
	_0,

	// LPI2C3
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT1,
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT1,
	_0,
	_0,
	_0,

	// LPI2C4
	s2 + iomux.ALT2, iomux.ALT0,
	s2 + iomux.ALT2, iomux.ALT0,
	_0,
	_0,
	_0,
}
