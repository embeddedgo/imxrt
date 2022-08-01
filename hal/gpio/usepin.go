// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"github.com/embeddedgo/imxrt/hal/iomux"
)

const (
	p1 = 1 << 5
	p2 = 2 << 5
	p3 = 3 << 5
	p4 = 4 << 5
)

var portBits = [...]uint8{
	iomux.EMC_00:   p4 + 0,
	iomux.EMC_01:   p4 + 1,
	iomux.EMC_02:   p4 + 2,
	iomux.EMC_03:   p4 + 3,
	iomux.EMC_04:   p4 + 4,
	iomux.EMC_05:   p4 + 5,
	iomux.EMC_06:   p4 + 6,
	iomux.EMC_07:   p4 + 7,
	iomux.EMC_08:   p4 + 8,
	iomux.EMC_09:   p4 + 9,
	iomux.EMC_10:   p4 + 10,
	iomux.EMC_11:   p4 + 11,
	iomux.EMC_12:   p4 + 12,
	iomux.EMC_13:   p4 + 13,
	iomux.EMC_14:   p4 + 14,
	iomux.EMC_15:   p4 + 15,
	iomux.EMC_16:   p4 + 16,
	iomux.EMC_17:   p4 + 17,
	iomux.EMC_18:   p4 + 18,
	iomux.EMC_19:   p4 + 19,
	iomux.EMC_20:   p4 + 20,
	iomux.EMC_21:   p4 + 21,
	iomux.EMC_22:   p4 + 22,
	iomux.EMC_23:   p4 + 23,
	iomux.EMC_24:   p4 + 24,
	iomux.EMC_25:   p4 + 25,
	iomux.EMC_26:   p4 + 26,
	iomux.EMC_27:   p4 + 27,
	iomux.EMC_28:   p4 + 28,
	iomux.EMC_29:   p4 + 29,
	iomux.EMC_30:   p4 + 30,
	iomux.EMC_31:   p4 + 31,
	iomux.EMC_32:   p3 + 18,
	iomux.EMC_33:   p3 + 19,
	iomux.EMC_34:   p3 + 20,
	iomux.EMC_35:   p3 + 21,
	iomux.EMC_36:   p3 + 22,
	iomux.EMC_37:   p3 + 23,
	iomux.EMC_38:   p3 + 24,
	iomux.EMC_39:   p3 + 25,
	iomux.EMC_40:   p3 + 26,
	iomux.EMC_41:   p3 + 27,
	iomux.AD_B0_00: p1 + 0,
	iomux.AD_B0_01: p1 + 1,
	iomux.AD_B0_02: p1 + 2,
	iomux.AD_B0_03: p1 + 3,
	iomux.AD_B0_04: p1 + 4,
	iomux.AD_B0_05: p1 + 5,
	iomux.AD_B0_06: p1 + 6,
	iomux.AD_B0_07: p1 + 7,
	iomux.AD_B0_08: p1 + 8,
	iomux.AD_B0_09: p1 + 9,
	iomux.AD_B0_10: p1 + 10,
	iomux.AD_B0_11: p1 + 11,
	iomux.AD_B0_12: p1 + 12,
	iomux.AD_B0_13: p1 + 13,
	iomux.AD_B0_14: p1 + 14,
	iomux.AD_B0_15: p1 + 15,
	iomux.AD_B1_00: p1 + 16,
	iomux.AD_B1_01: p1 + 17,
	iomux.AD_B1_02: p1 + 18,
	iomux.AD_B1_03: p1 + 19,
	iomux.AD_B1_04: p1 + 20,
	iomux.AD_B1_05: p1 + 21,
	iomux.AD_B1_06: p1 + 22,
	iomux.AD_B1_07: p1 + 23,
	iomux.AD_B1_08: p1 + 24,
	iomux.AD_B1_09: p1 + 25,
	iomux.AD_B1_10: p1 + 26,
	iomux.AD_B1_11: p1 + 27,
	iomux.AD_B1_12: p1 + 28,
	iomux.AD_B1_13: p1 + 29,
	iomux.AD_B1_14: p1 + 30,
	iomux.AD_B1_15: p1 + 31,
	iomux.B0_00:    p2 + 0,
	iomux.B0_01:    p2 + 1,
	iomux.B0_02:    p2 + 2,
	iomux.B0_03:    p2 + 3,
	iomux.B0_04:    p2 + 4,
	iomux.B0_05:    p2 + 5,
	iomux.B0_06:    p2 + 6,
	iomux.B0_07:    p2 + 7,
	iomux.B0_08:    p2 + 8,
	iomux.B0_09:    p2 + 9,
	iomux.B0_10:    p2 + 10,
	iomux.B0_11:    p2 + 11,
	iomux.B0_12:    p2 + 12,
	iomux.B0_13:    p2 + 13,
	iomux.B0_14:    p2 + 14,
	iomux.B0_15:    p2 + 15,
	iomux.B1_00:    p2 + 16,
	iomux.B1_01:    p2 + 17,
	iomux.B1_02:    p2 + 18,
	iomux.B1_03:    p2 + 19,
	iomux.B1_04:    p2 + 20,
	iomux.B1_05:    p2 + 21,
	iomux.B1_06:    p2 + 22,
	iomux.B1_07:    p2 + 23,
	iomux.B1_08:    p2 + 24,
	iomux.B1_09:    p2 + 25,
	iomux.B1_10:    p2 + 26,
	iomux.B1_11:    p2 + 27,
	iomux.B1_12:    p2 + 28,
	iomux.B1_13:    p2 + 29,
	iomux.B1_14:    p2 + 30,
	iomux.B1_15:    p2 + 31,
	iomux.SD_B0_00: p3 + 12,
	iomux.SD_B0_01: p3 + 13,
	iomux.SD_B0_02: p3 + 14,
	iomux.SD_B0_03: p3 + 15,
	iomux.SD_B0_04: p3 + 16,
	iomux.SD_B0_05: p3 + 17,
	iomux.SD_B1_00: p3 + 0,
	iomux.SD_B1_01: p3 + 1,
	iomux.SD_B1_02: p3 + 2,
	iomux.SD_B1_03: p3 + 3,
	iomux.SD_B1_04: p3 + 4,
	iomux.SD_B1_05: p3 + 5,
	iomux.SD_B1_06: p3 + 6,
	iomux.SD_B1_07: p3 + 7,
	iomux.SD_B1_08: p3 + 8,
	iomux.SD_B1_09: p3 + 9,
	iomux.SD_B1_10: p3 + 10,
	iomux.SD_B1_11: p3 + 11,
}

func UsePin(pin iomux.Pin, fast bool) Bit {
	portBit := int(portBits[pin])
	pn := portBit >> 5
	if fast {
		pn += 5
	}
	bit := P(pn).Bit(portBit & 31)
	bit.ConnectMux()
	pin.SetAltFunc(iomux.GPIO)
	return bit
}
