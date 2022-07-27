// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iomux

// Pin represents an I/O pin (pad).
type Pin uint8

const (
	EMC_00 Pin = iota
	EMC_01
	EMC_02
	EMC_03
	EMC_04
	EMC_05
	EMC_06
	EMC_07
	EMC_08
	EMC_09
	EMC_10
	EMC_11
	EMC_12
	EMC_13
	EMC_14
	EMC_15
	EMC_16
	EMC_17
	EMC_18
	EMC_19
	EMC_20
	EMC_21
	EMC_22
	EMC_23
	EMC_24
	EMC_25
	EMC_26
	EMC_27
	EMC_28
	EMC_29
	EMC_30
	EMC_31
	EMC_32
	EMC_33
	EMC_34
	EMC_35
	EMC_36
	EMC_37
	EMC_38
	EMC_39
	EMC_40
	EMC_41
	AD_B0_00
	AD_B0_01
	AD_B0_02
	AD_B0_03
	AD_B0_04
	AD_B0_05
	AD_B0_06
	AD_B0_07
	AD_B0_08
	AD_B0_09
	AD_B0_10
	AD_B0_11
	AD_B0_12
	AD_B0_13
	AD_B0_14
	AD_B0_15
	AD_B1_00
	AD_B1_01
	AD_B1_02
	AD_B1_03
	AD_B1_04
	AD_B1_05
	AD_B1_06
	AD_B1_07
	AD_B1_08
	AD_B1_09
	AD_B1_10
	AD_B1_11
	AD_B1_12
	AD_B1_13
	AD_B1_14
	AD_B1_15
	B0_00
	B0_01
	B0_02
	B0_03
	B0_04
	B0_05
	B0_06
	B0_07
	B0_08
	B0_09
	B0_10
	B0_11
	B0_12
	B0_13
	B0_14
	B0_15
	B1_00
	B1_01
	B1_02
	B1_03
	B1_04
	B1_05
	B1_06
	B1_07
	B1_08
	B1_09
	B1_10
	B1_11
	B1_12
	B1_13
	B1_14
	B1_15
	SD_B0_00
	SD_B0_01
	SD_B0_02
	SD_B0_03
	SD_B0_04
	SD_B0_05
	SD_B1_00
	SD_B1_01
	SD_B1_02
	SD_B1_03
	SD_B1_04
	SD_B1_05
	SD_B1_06
	SD_B1_07
	SD_B1_08
	SD_B1_09
	SD_B1_10
	SD_B1_11
)

type Config uint32

const (
	FastSR Config = 1 << 0 // Enable fast slew rate

	Drive  Config = 7 << 3 // Drive strength field
	Drive0 Config = 0 << 3 // Rout = ∞Ω (output driver disabled)
	Drive1 Config = 1 << 3 // Rout = R, R = 150Ω @ 3V3, R = 260Ω @ 1V8
	Drive2 Config = 2 << 3 // Rout = R / 2
	Drive3 Config = 3 << 3 // Rout = R / 3
	Drive4 Config = 4 << 3 // Rout = R / 4
	Drive5 Config = 5 << 3 // Rout = R / 5
	Drive6 Config = 6 << 3 // Rout = R / 6
	Drive7 Config = 7 << 3 // Rout = R / 7

	Speed       Config = 3 << 6 // Speed field
	SpeedLow    Config = 0 << 6 // Speed low (50MHz)
	SpeedMedium Config = 1 << 6 // Speed medium (100MHz)
	SpeedFast   Config = 2 << 6 // Speed fast (150MHz)
	SpeedMax    Config = 3 << 6 // Speed max (200MHz)

	OpenDrain Config = 1 << 11 // Enable open drain mode

	PK Config = 1 << 12 // Enable pull/keep mode

	Pull Config = 1 << 13 // Use pull mode instead of keep mode

	PullSel  Config = 3 << 14 // Select pull direction and strength
	Down100k Config = 0 << 14 // 100KΩ pull-down
	Up47K    Config = 1 << 14 //  47KΩ pull-up
	Up100K   Config = 2 << 14 // 100KΩ pull up
	Up22K    Config = 3 << 14 //  22KΩ pull up

	Hys Config = 1 << 16 //+ Enable hysteresis mode
)

// AltFunc represents a mux mode.
type AltFunc uint8

const (
	ALT   AltFunc = 0xf << 0 // Mux mode select field
	ALT0  AltFunc = 0x0 << 0 // Select ALT0 mux mode
	ALT1  AltFunc = 0x1 << 0 // Select ALT1 mux mode
	ALT2  AltFunc = 0x2 << 0 // Select ALT2 mux mode
	ALT3  AltFunc = 0x3 << 0 // Select ALT3 mux mode
	ALT4  AltFunc = 0x4 << 0 // Select ALT4 mux mode
	ALT5  AltFunc = 0x5 << 0 // Select ALT5 mux mode
	ALT6  AltFunc = 0x6 << 0 // Select ALT6 mux mode
	ALT7  AltFunc = 0x7 << 0 // Select ALT7 mux mode
	ALT8  AltFunc = 0x8 << 0 // Select ALT8 mux mode
	ALT9  AltFunc = 0x9 << 0 // Select ALT9 mux mode
	ALT10 AltFunc = 0xa << 0 // Select ALT10 mux mode

	SION AltFunc = 0x1 << 4 // Software Input On Field
)

func (p Pin) Config() Config {
	return Config(pr().pad[p].Load())
}

// Setup configures pin.
func (p Pin) Setup(cfg Config) {
	pr().pad[p].Store(uint32(cfg))
}

// AltFunc returns a currently set muxmode for pin.
func (p Pin) AltFunc() AltFunc {
	return AltFunc(pr().mux[p].Load())
}

// SetAltFunc sets a mux mode for pin.
func (p Pin) SetAltFunc(af AltFunc) {
	pr().mux[p].Store(uint32(af))
}
