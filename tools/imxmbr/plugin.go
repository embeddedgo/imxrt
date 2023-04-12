package main

// Some Thumb1 instructions thar can be useful in the plugin code.
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

// Plugin code to configure the whole FlexRAM as OCRAM. See also the comment
// about the possible DCD way that unfortunately doesn't work.
var plugin = []uint16{
	0:  B(0),
	1:  LDRPC(1, 16, R3),
	2:  STRW(R3, 0, R0),
	3:  LDRPC(3, 18, R3),
	4:  STRW(R3, 0, R1),
	5:  MOV(0, R3),
	6:  STRW(R3, 0, R2),
	7:  LDRPC(7, 20, R0), // load GPR16 address into R0
	8:  LDRPC(8, 22, R3), // load 0x5555_5555 into R3
	9:  STRW(R3, 1, R0),  // store 0x5555_5555 into GPR17
	10: LDRW(0, R0, R3),  // load GPR16 into R3
	11: MOV(1<<2, R2),    // FLEXRAM_BANK_CFG_SEL
	12: ORR(R2, R3),      // set FLEXRAM_BANK_CFG_SEL in GPR16
	13: STRW(R3, 0, R0),  // store updated GPR16
	14: MOV(1, R0),       //  success
	15: BX(LR),           // return to bootloader

	// Data
	16: 0x1300, 0x6000, // stage2IVTAddr
	18: 0x0000, 0x0040, // image size
	20: 0xC040, 0x400A, // GPR16 addr 0x400A_0xC040
	22: 0x5555, 0x5555,
}

// The DCD below should configure the whole FlexRAM as OCRAM but it doesn't
// work. Debuging it shows that there is undocumented restriction to modify
// FLEXRAM_BANK_CFG. It's probably because some of the I.MX chips use the
// FlexRAM to run the bootloader code.

const (
	dcdNOP   = 0xc0 << 24
	dcdWrite = 0xcc << 24
	dcdClear = dcdWrite | 2<<3
	dcdSet   = dcdWrite | 3<<3
	dcdCheck = 0xcf << 24
)
