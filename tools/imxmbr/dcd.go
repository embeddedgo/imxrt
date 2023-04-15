// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

// The DCD below should configure the whole FlexRAM as OCRAM but it doesn't
// work. Debuging it shows that there is undocumented restriction to modify
// FLEXRAM_BANK_CFG. It's probably because some of the I.MX chips use the
// FlexRAM to run the bootloader code.

const (
	gpr16 = 0x400A_C040
	gpr17 = 0x400A_C044
)

const (
	dcdNOP   = 0xc0 << 24
	dcdWrite = 0xcc << 24
	dcdClear = dcdWrite | 2<<3
	dcdSet   = dcdWrite | 3<<3
	dcdCheck = 0xcf << 24
)

const dcdLen = 7 // must be in sync with the number of 32 bit words in dcd

var dcd = [dcdLen]uint32{
	0xd2<<24 | dcdLen*4<<8 | 0x41,             // DCD header
	dcdWrite | 3*4<<8 | 4, gpr17, 0x5555_5555, // FLEXRAM_BANK_CFG
	dcdSet | 3*4<<8 | 4, gpr16, 1 << 2, // FLEXRAM_BANK_CFG_SEL
}
