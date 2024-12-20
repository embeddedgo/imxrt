// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package ewm provides access to the registers of the EWM peripheral.
//
// Instances:
//
//	EWM  EWM_BASE  -  EWM*
//
// Registers:
//
//	0x000  8  CTRL          Control Register
//	0x001  8  SERV          Service Register
//	0x002  8  CMPL          Compare Low Register
//	0x003  8  CMPH          Compare High Register
//	0x004  8  CLKCTRL       Clock Control Register
//	0x005  8  CLKPRESCALER  Clock Prescaler Register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package ewm

const (
	EWMEN CTRL = 0x01 << 0 //+ EWM enable.
	ASSIN CTRL = 0x01 << 1 //+ EWM_in's Assertion State Select.
	INEN  CTRL = 0x01 << 2 //+ Input Enable.
	INTEN CTRL = 0x01 << 3 //+ Interrupt Enable.
)

const (
	EWMENn = 0
	ASSINn = 1
	INENn  = 2
	INTENn = 3
)

const (
	SERVICE SERV = 0xFF << 0 //+ SERVICE
)

const (
	SERVICEn = 0
)

const (
	COMPAREL CMPL = 0xFF << 0 //+ COMPAREL
)

const (
	COMPARELn = 0
)

const (
	COMPAREH CMPH = 0xFF << 0 //+ COMPAREH
)

const (
	COMPAREHn = 0
)

const (
	CLKSEL CLKCTRL = 0x03 << 0 //+ CLKSEL
)

const (
	CLKSELn = 0
)

const (
	CLK_DIV CLKPRESCALER = 0xFF << 0 //+ CLK_DIV
)

const (
	CLK_DIVn = 0
)
