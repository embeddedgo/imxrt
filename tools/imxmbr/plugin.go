// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Some Thumb1 instructions that can be useful in the plugin code.

func ADD(u8, Rdn int) uint16 {
	return uint16(0b00110<<11 | Rdn&7<<8 | u8&0xff)
}
func B(offset int) uint16 {
	return uint16(0b11100<<11 | (offset-2)&0b11111111111)
}
func BX(Rm int) uint16 {
	return uint16(0b010001110<<7 | Rm&15<<3)
}
func BKPT(u8 int) uint16 {
	return uint16(0b10111110<<8 | u8&0xff)
}
func LDRPC(ipos, dpos, Rt int) uint16 {
	return uint16(0b01001<<11 | Rt&7<<8 | (dpos-ipos&^1-2)/2&0xff)
}
func LDRW(u5, Rn, Rt int) uint16 {
	return uint16(0b01101<<11 | u5&31<<6 | Rn&7<<3 | Rt&7)
}
func MOV(u8, Rd int) uint16 {
	return uint16(0b00100<<11 | Rd&7<<8 | u8&0xff)
}
func ORR(Rm, Rdn int) uint16 {
	return uint16(0b0100001100<<6 | Rm&7<<3 | Rdn&7)
}
func STRW(Rt, u5, Rn int) uint16 {
	return uint16(0b01100<<11 | u5&31<<6 | Rn&7<<3 | Rt&7)
}
func SUB(u8, Rdn int) uint16 {
	return uint16(0b00111<<11 | Rdn&7<<8 | u8&0xff)
}

const (
	NOP = 0b1011_1111_0000_0000

	R0 = 0
	R1 = 1
	R2 = 2
	R3 = 3
	LR = 14
	PC = 15
)

// Plugin code to configure the whole FlexRAM as OCRAM.
//
// By default FlexRAM is configured by fuses 0x6D0[19:16] = 0b0000 which means
// 128KB 128KB 256KB {O O O O D D I I I I D D O O O O}. If FLEXRAM_BANK_CFG_SEL
// is set FLEXRAM_BANK_CFG is used instead of fuses.
//
// See also the comment in dcd.go about the possible DCD way that unfortunately
// doesn't work.

var plugin = []uint16{
	0:  NOP, // change to B(0) to stop here for debuging
	1:  LDRPC(1, 16, R3),
	2:  STRW(R3, 0, R0),
	3:  LDRPC(3, 18, R3),
	4:  STRW(R3, 0, R1),
	5:  MOV(0, R3),
	6:  STRW(R3, 0, R2),
	7:  LDRPC(7, 20, R0), // load GPR16 address into R0
	8:  LDRPC(8, 22, R3), // load FLEXRAM_BANK_CFG into R3
	9:  STRW(R3, 1, R0),  // store FLEXRAM_BANK_CFG into GPR17
	10: LDRW(0, R0, R3),  // load GPR16 into R3
	11: MOV(1<<2, R2),    // FLEXRAM_BANK_CFG_SEL
	12: ORR(R2, R3),      // set FLEXRAM_BANK_CFG_SEL in GPR16
	13: STRW(R3, 0, R0),  // store updated GPR16
	14: MOV(1, R0),       // success
	15: BX(LR),           // return to bootloader

	// Data
	16: 0x1300, 0x6000, // stage2IVTAddr
	18: 0x0000, 0x0000, // image size, set in main function
	20: 0xC040, 0x400A, // GPR16 addr 0x400A_0xC040
	22: 0x0000, 0x0000, // FLEXRAM_BANK_CFG, set in main function
}

var (
	pluginImageSize  = plugin[18:20]
	pluginFlexRAMCfg = plugin[22:24]
)
