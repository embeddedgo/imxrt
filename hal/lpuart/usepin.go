// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart

import (
	"github.com/embeddedgo/imxrt/hal/internal/periph"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

type Signal uint8

const (
	RTSn Signal = iota
	TXD
	CTSn
	RXD
)

// UsePin is a helper function that can be used to configure IO pins as required
// by USART peripheral. Only certain pins can be used (see datasheet).
func (d *Driver) UsePin(pin iomux.Pin, sig Signal) {
	var cfg iomux.Config
	if sig <= TXD {
		cfg = iomux.Drive2 // 75Ω @ 3V3, R = 130Ω @ 1V8
	}
	pin.Setup(cfg)
	af := periph.AltFunc(pins[:], alts[:], num(d.p)*4+int(sig), pin)
	pin.SetAltFunc(af)
}

var pins = [...]iomux.Pin{
	// LPUART1
	/* RTS */ iomux.AD_B0_15,
	/* TXD */ iomux.AD_B0_12,
	/* CTS */ iomux.AD_B0_14,
	/* RXD */ iomux.AD_B0_13,

	// LPUART2
	/* RTS */ iomux.AD_B1_01,
	/* TXD */ iomux.SD_B1_11, iomux.AD_B1_02,
	/* CTS */ iomux.AD_B1_00,
	/* RXD */ iomux.AD_B1_03, iomux.SD_B1_10,

	// LPUART3
	/* RTS */ iomux.AD_B1_05, iomux.EMC_16,
	/* TXD */ iomux.AD_B1_06, iomux.B0_08, iomux.EMC_13,
	/* CTS */ iomux.EMC_15, iomux.AD_B1_04,
	/* RXD */ iomux.EMC_14, iomux.B0_09, iomux.AD_B1_07,

	// LPUART4
	/* RTS */ iomux.AD_B1_05, iomux.EMC_16,
	/* TXD */ iomux.AD_B1_06, iomux.B0_08, iomux.EMC_13,
	/* CTS */ iomux.EMC_15, iomux.AD_B1_04,
	/* RXD */ iomux.EMC_14, iomux.B0_09, iomux.AD_B1_07,

	// LPUART5
	/* RTS */ iomux.EMC_27,
	/* TXD */ iomux.EMC_23, iomux.B1_12,
	/* CTS */ iomux.EMC_28,
	/* RXD */ iomux.B1_13, iomux.EMC_24,

	// LPUART6
	/* RTS */ iomux.EMC_29,
	/* TXD */ iomux.EMC_25, iomux.AD_B0_02,
	/* CTS */ iomux.EMC_30,
	/* RXD */ iomux.AD_B0_03, iomux.EMC_26,

	// LPUART7
	/* RTS */ iomux.SD_B1_07,
	/* TXD */ iomux.SD_B1_08, iomux.EMC_31,
	/* CTS */ iomux.SD_B1_06,
	/* RXD */ iomux.EMC_32, iomux.SD_B1_09,

	// LPUART8
	/* RTS */ iomux.SD_B0_03,
	/* TXD */ iomux.EMC_38, iomux.AD_B1_10, iomux.SD_B0_04,
	/* CTS */ iomux.SD_B0_02,
	/* RXD */ iomux.SD_B0_05, iomux.AD_B1_11, iomux.EMC_39,
}

var alts = [...]iomux.AltFunc{
	// LPUART1
	0x10 + iomux.ALT2,
	0x10 + iomux.ALT2,
	0x10 + iomux.ALT2,
	0x10 + iomux.ALT2,

	// LPUART2
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,

	// LPUART3
	0x20 + iomux.ALT2, iomux.ALT3,
	0x30 + iomux.ALT2, iomux.ALT3, iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,
	0x30 + iomux.ALT2, iomux.ALT3, iomux.ALT2,

	// LPUART4
	0x20 + iomux.ALT2, iomux.ALT3,
	0x30 + iomux.ALT2, iomux.ALT3, iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,
	0x30 + iomux.ALT2, iomux.ALT3, iomux.ALT2,

	// LPUART5
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT1,
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT1, iomux.ALT2,

	// LPUART6
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,

	// LPUART7
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,
	0x10 + iomux.ALT2,
	0x20 + iomux.ALT2, iomux.ALT2,

	// LPUART8
	0x10 + iomux.ALT2,
	0x30 + iomux.ALT2, iomux.ALT2, iomux.ALT2,
	0x10 + iomux.ALT2,
	0x30 + iomux.ALT2, iomux.ALT2, iomux.ALT2,
}
