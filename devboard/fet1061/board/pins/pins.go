// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pins

import (
	_ "github.com/embeddedgo/imxrt/devboard/fet1061/board/system"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

const (
	P6   = iomux.AD_B1_07 // K10
	P7   = iomux.AD_B1_06 // J12
	P8   = iomux.AD_B1_01 // K11
	P9   = iomux.AD_B1_00 // J11
	P11  = iomux.AD_B0_00 // M14
	P12  = iomux.AD_B1_02 // L11
	P13  = iomux.AD_B1_03 // M12
	P14  = iomux.AD_B0_15 // L10
	P15  = iomux.AD_B0_14 // H14
	P16  = iomux.B1_12    // D13
	P18  = iomux.B1_13    // D14
	P19  = iomux.B1_15    // B14
	P22  = iomux.B1_14    // C14
	P23  = iomux.AD_B0_13 // L14
	P24  = iomux.AD_B0_12 // K14
	P28  = iomux.AD_B0_03 // G11
	P29  = iomux.AD_B0_02 // M11
	P30  = iomux.AD_B0_11 // G10
	P31  = iomux.AD_B0_09 // F14
	P32  = iomux.AD_B0_07 // F12
	P33  = iomux.AD_B0_06 // E14
	P34  = iomux.AD_B0_08 // F13
	P35  = iomux.AD_B0_10 // G13
	P37  = iomux.B1_08    // A12
	P38  = iomux.B1_07    // B12
	P39  = iomux.B1_09    // A13
	P40  = iomux.B1_11    // C13
	P41  = iomux.B1_06    // C12
	P42  = iomux.B1_04    // E12
	P43  = iomux.B1_05    // D12
	P44  = iomux.B1_10    // B13
	P45  = iomux.EMC_41   // C7
	P46  = iomux.EMC_40   // A7
	P47  = iomux.EMC_39   // B7
	P48  = iomux.AD_B0_04 // F11
	P49  = iomux.AD_B0_05 // G14
	P52  = iomux.B0_10    // D9
	P53  = iomux.B0_11    // A10
	P54  = iomux.B0_13    // D10
	P55  = iomux.B0_12    // C10
	P56  = iomux.B0_14    // E10
	P57  = iomux.B1_00    // A11
	P58  = iomux.B1_03    // D11
	P59  = iomux.B1_01    // B11
	P60  = iomux.B1_02    // C11
	P61  = iomux.B0_15    // E11
	P62  = iomux.B0_01    // E7
	P63  = iomux.B0_00    // D7
	P64  = iomux.B0_02    // E8
	P68  = iomux.SD_B0_03 // K1
	P69  = iomux.SD_B0_02 // J1
	P70  = iomux.SD_B0_01 // J3
	P71  = iomux.SD_B0_00 // J4
	P72  = iomux.SD_B0_05 // J2
	P73  = iomux.SD_B0_04 // H2
	P77  = iomux.SD_B1_04 // P2
	P78  = iomux.SD_B1_02 // M3
	P79  = iomux.SD_B1_03 // M4
	P80  = iomux.SD_B1_00 // L5
	P81  = iomux.SD_B1_01 // M5
	P89  = iomux.AD_B0_01 // H10
	P91  = iomux.AD_B1_13 // H11
	P92  = iomux.AD_B1_14 // G12
	P93  = iomux.AD_B1_12 // H12
	P94  = iomux.AD_B1_15 // J14
	P95  = iomux.AD_B1_11 // J13
	P96  = iomux.AD_B1_04 // L12
	P97  = iomux.AD_B1_10 // L13
	P98  = iomux.AD_B1_09 // M13
	P99  = iomux.AD_B1_05 // K12
	P100 = iomux.AD_B1_08 // H13

	LED5 = P31
)
