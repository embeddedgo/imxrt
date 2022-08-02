// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pins

import (
	_ "github.com/embeddedgo/imxrt/devboard/teensy4/board/init"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

// Teensy 4.0 and 4.1 common pin names
const (
	P0  = iomux.AD_B0_03
	P1  = iomux.AD_B0_02
	P2  = iomux.EMC_04
	P3  = iomux.EMC_05
	P4  = iomux.EMC_06
	P5  = iomux.EMC_08
	P6  = iomux.B0_10
	P7  = iomux.B1_01
	P8  = iomux.B1_00
	P9  = iomux.B0_11
	P10 = iomux.B0_00
	P11 = iomux.B0_02
	P12 = iomux.B0_01
	P13 = iomux.B0_03
	P14 = iomux.AD_B1_02
	P15 = iomux.AD_B1_03
	P16 = iomux.AD_B1_07
	P17 = iomux.AD_B1_06
	P18 = iomux.AD_B1_01
	P19 = iomux.AD_B1_00
	P20 = iomux.AD_B1_10
	P21 = iomux.AD_B1_11
	P22 = iomux.AD_B1_08
	P23 = iomux.AD_B1_09
	P24 = iomux.AD_B0_12
	P25 = iomux.AD_B0_13
	P26 = iomux.AD_B1_14
	P27 = iomux.AD_B1_15
	P28 = iomux.EMC_32
	P29 = iomux.EMC_31
	P30 = iomux.EMC_37
	P31 = iomux.EMC_36
	P32 = iomux.B0_12
	P33 = iomux.EMC_07

	LED = P13
)

// Teensy 4.0 specific pin names
const (
	P34_0 = iomux.SD_B0_03
	P35_0 = iomux.SD_B0_02
	P36_0 = iomux.SD_B0_01
	P37_0 = iomux.SD_B0_00
	P38_0 = iomux.SD_B0_04
	P39_0 = iomux.SD_B0_05
	P40_0 = iomux.B0_04
	P41_0 = iomux.B0_05
	P42_0 = iomux.B0_06
	P43_0 = iomux.B0_07
	P44_0 = iomux.B0_08
	P45_0 = iomux.B0_09
)

// Teensy 4.1 specific pin names
const (
	P34_1 = iomux.B1_13
	P35_1 = iomux.B1_12
	P36_1 = iomux.B1_02
	P37_1 = iomux.B1_03
	P38_1 = iomux.AD_B1_12
	P39_1 = iomux.AD_B1_13
	P40_1 = iomux.AD_B1_04
	P41_1 = iomux.AD_B1_05
	P42_1 = iomux.SD_B0_03
	P43_1 = iomux.SD_B0_02
	P44_1 = iomux.SD_B0_01
	P45_1 = iomux.SD_B0_00
	P46_1 = iomux.SD_B0_05
	P47_1 = iomux.SD_B0_04
	P48_1 = iomux.EMC_24
	P49_1 = iomux.EMC_27
	P50_1 = iomux.EMC_28
	P51_1 = iomux.EMC_22
	P52_1 = iomux.EMC_26
	P53_1 = iomux.EMC_25
	P54_1 = iomux.EMC_29
)
