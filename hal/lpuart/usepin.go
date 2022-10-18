// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal/periph"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

type Signal int8

// Do not reorder CTSn, RXD, TXD constants. The order  reflects the sequence of
// select registers and is used by periph.AltFunc function.

const (
	CTSn Signal = iota
	RXD
	TXD
	RTSn
)

// Pins return IO pins that can be used for singal sig.
func (p *Periph) Pins(sig Signal) []iomux.Pin {
	return periph.Pins(pins[:], alts[:], num(p)*4+int(sig))
}

// UsePin is a helper function that can be used to configure IO pins as required
// by LPUART peripheral. Only certain pins can be used (see datasheet). UsePin
// returns true on succes or false if it isn't possible to use a pin as a sig.
// See also Periph.Pins.
func (d *Driver) UsePin(pin iomux.Pin, sig Signal) bool {
	af, s, daisy := periph.AltFunc(pins[:], alts[:], num(d.p)*4+int(sig), pin)
	if af < 0 {
		return false
	}
	var cfg iomux.Config
	if sig >= TXD {
		cfg = iomux.Drive2 // 75Ω @ 3.3V, 130Ω @ 1.8V
	}
	pin.Setup(cfg)
	pin.SetAltFunc(af)
	if s >= 0 {
		iosel := (*[15]mmio.R32[int32])(unsafe.Pointer(uintptr(0x401F852C)))
		iosel[s].Store(int32(daisy))
	}
	return true
}

const daisyBase = 0x401F_852C

var pins = [...]iomux.Pin{
	// LPUART1
	/* CTS */ iomux.AD_B0_14,
	/* RXD */ iomux.AD_B0_13,
	/* TXD */ iomux.AD_B0_12,
	/* RTS */ iomux.AD_B0_15,

	// LPUART2
	/* CTS */ iomux.AD_B1_00,
	/* RXD */ iomux.SD_B1_10, iomux.AD_B1_03,
	/* TXD */ iomux.SD_B1_11, iomux.AD_B1_02,
	/* RTS */ iomux.AD_B1_01,

	// LPUART3
	/* CTS */ iomux.EMC_15, iomux.AD_B1_04,
	/* RXD */ iomux.AD_B1_07, iomux.EMC_14, iomux.B0_09,
	/* TXD */ iomux.AD_B1_06, iomux.EMC_13, iomux.B0_08,
	/* RTS */ iomux.AD_B1_05, iomux.EMC_16,

	// LPUART4
	/* CTS */ iomux.EMC_17,
	/* RXD */ iomux.SD_B1_01, iomux.EMC_20, iomux.B1_01,
	/* TXD */ iomux.SD_B1_00, iomux.EMC_19, iomux.B1_00,
	/* RTS */ iomux.EMC_18,

	// LPUART5
	/* CTS */ iomux.EMC_28,
	/* RXD */ iomux.EMC_24, iomux.B1_13,
	/* TXD */ iomux.EMC_23, iomux.B1_12,
	/* RTS */ iomux.EMC_27,

	// LPUART6
	/* CTS */ iomux.EMC_30,
	/* RXD */ iomux.EMC_26, iomux.AD_B0_03,
	/* TXD */ iomux.EMC_25, iomux.AD_B0_02,
	/* RTS */ iomux.EMC_29,

	// LPUART7
	/* CTS */ iomux.SD_B1_06,
	/* RXD */ iomux.SD_B1_09, iomux.EMC_32,
	/* TXD */ iomux.SD_B1_08, iomux.EMC_31,
	/* RTS */ iomux.SD_B1_07,

	// LPUART8
	/* CTS */ iomux.SD_B0_02,
	/* RXD */ iomux.SD_B0_05, iomux.AD_B1_11, iomux.EMC_39,
	/* TXD */ iomux.SD_B0_04, iomux.AD_B1_10, iomux.EMC_38,
	/* RTS */ iomux.SD_B0_03,
}

const (
	// no select register (one level of I/O muxing)
	_1 = 0x10
	_2 = 0x20
	_3 = 0x30

	// select register exists (two levels of I/O muxing)
	s1 = -1<<7 + _1
	s2 = -1<<7 + _2
	s3 = -1<<7 + _3
)

var alts = [...]iomux.AltFunc{
	// LPUART1
	_1 + iomux.ALT2,
	_1 + iomux.ALT2,
	_1 + iomux.ALT2,
	_1 + iomux.ALT2,

	// LPUART2
	_1 + iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	_1 + iomux.ALT2,

	// LPUART3
	s2 + iomux.ALT2, iomux.ALT2,
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT3,
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT3,
	_2 + iomux.ALT2, iomux.ALT2,

	// LPUART4
	_1 + iomux.ALT2,
	s3 + iomux.ALT4, iomux.ALT2, iomux.ALT2,
	s3 + iomux.ALT4, iomux.ALT2, iomux.ALT2,
	_1 + iomux.ALT2,

	// LPUART5
	_1 + iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT1,
	s2 + iomux.ALT2, iomux.ALT1,
	_1 + iomux.ALT2,

	// LPUART6
	_1 + iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	_1 + iomux.ALT2,

	// LPUART7
	_1 + iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	s2 + iomux.ALT2, iomux.ALT2,
	_1 + iomux.ALT2,

	// LPUART8
	_1 + iomux.ALT2,
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT2,
	s3 + iomux.ALT2, iomux.ALT2, iomux.ALT2,
	_1 + iomux.ALT2,
}
