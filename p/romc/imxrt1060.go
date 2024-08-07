// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package romc provides access to the registers of the ROMC peripheral.
//
// Instances:
//
//	ROMC  ROMC_BASE  -  -
//
// Registers:
//
//	0x0D4 32  ROMPATCH[8]   ROMC Data Registers
//	0x0F4 32  ROMPATCHCNTL  ROMC Control Register
//	0x0F8 32  ROMPATCHENH   ROMC Enable Register High
//	0x0FC 32  ROMPATCHENL   ROMC Enable Register Low
//	0x100 32  ROMPATCH[16]  ROMC Address Registers
//	0x208 32  ROMPATCHSR    ROMC Status Register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package romc

const (
	DATAX ROMPATCH = 0xFFFFFFFF << 0 //+ Data Fix Registers - Stores the data used for 1-word data fix operations
)

const (
	DATAXn = 0
)

const (
	DATAFIX   ROMPATCHCNTL = 0xFF << 0  //+ Data Fix Enable - Controls the use of the first 8 address comparators for 1-word data fix or for code patch routine
	DATAFIX_0 ROMPATCHCNTL = 0x00 << 0  //  Address comparator triggers a opcode patch
	DATAFIX_1 ROMPATCHCNTL = 0x01 << 0  //  Address comparator triggers a data fix
	DIS       ROMPATCHCNTL = 0x01 << 29 //+ ROMC Disable -- This bit, when set, disables all ROMC operations
)

const (
	DATAFIXn = 0
	DISn     = 29
)

const (
	ENABLE   ROMPATCHENL = 0xFFFF << 0 //+ Enable Address Comparator - This bit enables the corresponding address comparator to trigger an event
	ENABLE_0 ROMPATCHENL = 0x00 << 0   //  Address comparator disabled
	ENABLE_1 ROMPATCHENL = 0x01 << 0   //  Address comparator enabled, ROMC will trigger a opcode patch or data fix event upon matching of the associated address
)

const (
	ENABLEn = 0
)

const (
	THUMBX ROMPATCH = 0x01 << 0     //+ THUMB Comparator Select - Indicates that this address will trigger a THUMB opcode patch or an Arm opcode patch
	ADDRX  ROMPATCH = 0x3FFFFF << 1 //+ Address Comparator Registers - Indicates the memory address to be watched
)

const (
	THUMBXn = 0
	ADDRXn  = 1
)

const (
	SOURCE    ROMPATCHSR = 0x3F << 0  //+ ROMC Source Number - Binary encoding of the number of the address comparator which has an address match in the most recent patch event on ROMC AHB
	SOURCE_0  ROMPATCHSR = 0x00 << 0  //  Address Comparator 0 matched
	SOURCE_1  ROMPATCHSR = 0x01 << 0  //  Address Comparator 1 matched
	SOURCE_15 ROMPATCHSR = 0x0F << 0  //  Address Comparator 15 matched
	SW        ROMPATCHSR = 0x01 << 17 //+ ROMC AHB Multiple Address Comparator matches Indicator - Indicates that multiple address comparator matches occurred
)

const (
	SOURCEn = 0
	SWn     = 17
)
