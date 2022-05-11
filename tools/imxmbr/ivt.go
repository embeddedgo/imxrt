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
	Self:     0x1000, // ivt address on SPI NOR Flash
}

var bootData = &BootData{
	Start: 0x60000000, // Flex SPI start address
}
