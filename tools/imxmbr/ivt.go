package main

type ImageVectorTable struct {
	Tag      uint8
	LenHi    uint8
	LenLo    uint8
	Version  uint8
	Entry    uint32
	_        uint32
	DCD      uint32
	BootData uint32
	Self     uint32
	CSF      uint32
	_        uint32
}

type BootData struct {
	Start  uint32
	Length uint32
	Plugin uint32
}

var ivt = &ImageVectorTable{
	Tag:      0xd1,
	LenLo:    32, // sizeof(ivt)
	Version:  0x43,
	Entry:    0x60002000, // the beggining of program image (ARMv7-M vector table)
	BootData: 0x60001020,
	Self:     0x60001000, // ivt address on SPI NOR Flash
}

var bootData = &BootData{
	Start: 0x60000000, // Flex SPI start address
}

const dcdLen = 2

const (
	dcdNOP   = 0xc0 << 24
	dcdWrite = 0xcc << 24
	dcdCheck = 0xcf << 24
)

const (
	dcdClear = dcdWrite | 2<<3
	dcdSet   = dcdWrite | 3<<3
)

const (
	gpr16 = 0x400A_C040
	gpr17 = 0x400A_C044
)

var dcd = [dcdLen]uint32{
	0xd2<<24 | dcdLen<<8 | 0x41,                     // TAG
	//dcdWrite | (1+2)<<10 | 4, gpr16, 1 << 2, 1 << 2, // FLEXRAM_BANK_CFG_SEL
	//dcdWrite | (1+1)<<10 | 4, gpr17, 0x5555_5555, 0xffff_ffff, // FLEXRAM_BANK_CFG
}

// By default FlexRAM is configured by fuses 0x6D0[19:16] = 0b0000 which means
// 128KB 128KB 256KB {O O O O D D I I I I D D O O O O}. If FLEXRAM_BANK_CFG_SEL
// is set FLEXRAM_BANK_CFG is used instead of fuses.
