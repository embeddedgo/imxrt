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

const ivtSize = 32

type BootData struct {
	Start  uint32
	Length uint32
	Plugin uint32
}

const bootDataSize = 12

const (
	baseAddr      = 0x60000000
	ivtAddr       = baseAddr + 0x1000  // 0x60001000
	bootDataAddr  = baseAddr + 0x1020  // 0x60001020
	dcdAddr       = baseAddr + 0x1030  // 0x60001030
	pluginAddr    = baseAddr + 0x1200  // 0x60001200
	stage2IVTAddr = baseAddr + 0x1300  // 0x60001300
	mbrEndAddr    = baseAddr + mbrSize // 0x60002000
)

var regularIVT = &ImageVectorTable{
	Tag:     0xd1,
	LenLo:   ivtSize,
	Version: 0x41,
	Entry:   mbrEndAddr, // for ARMv7-M the address of interrupt vector table
	//DCD:      dcdAddr, // DCD cannot be used to set FLEXRAM_BANK_CFG
	BootData: bootDataAddr,
	Self:     ivtAddr,
}

var pluginIVT = &ImageVectorTable{
	Tag:      0xd1,
	LenLo:    ivtSize,
	Version:  0x41,
	Entry:    pluginAddr | 1,
	BootData: bootDataAddr,
	Self:     ivtAddr,
}

var stage2IVT = &ImageVectorTable{
	Tag:     0xd1,
	LenLo:   ivtSize,
	Version: 0x41,
	Entry:   mbrEndAddr,
	Self:    stage2IVTAddr,
}

var bootData = &BootData{
	Start: baseAddr, // Flex SPI start address
}

// By default FlexRAM is configured by fuses 0x6D0[19:16] = 0b0000 which means
// 128KB 128KB 256KB {O O O O D D I I I I D D O O O O}. If FLEXRAM_BANK_CFG_SEL
// is set FLEXRAM_BANK_CFG is used instead of fuses.

const (
	gpr16 = 0x400A_C040
	gpr17 = 0x400A_C044
)

const dcdLen = 7 // must be in sync with the number of 32 bit words in dcd

var dcd = [dcdLen]uint32{
	0xd2<<24 | dcdLen*4<<8 | 0x41,             // DCD header
	dcdWrite | 3*4<<8 | 4, gpr17, 0x5555_5555, // FLEXRAM_BANK_CFG
	dcdSet | 3*4<<8 | 4, gpr16, 1 << 2, // FLEXRAM_BANK_CFG_SEL
}
