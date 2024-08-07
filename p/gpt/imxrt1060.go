// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package gpt provides access to the registers of the GPT peripheral.
//
// Instances:
//
//	GPT1  GPT1_BASE  -  GPT1*
//	GPT2  GPT2_BASE  -  GPT2*
//
// Registers:
//
//	0x000 32  CR    GPT Control Register
//	0x004 32  PR    GPT Prescaler Register
//	0x008 32  SR    GPT Status Register
//	0x00C 32  IR    GPT Interrupt Register
//	0x010 32  OCR1  GPT Output Compare Register 1
//	0x014 32  OCR2  GPT Output Compare Register 2
//	0x018 32  OCR3  GPT Output Compare Register 3
//	0x01C 32  ICR1  GPT Input Capture Register 1
//	0x020 32  ICR2  GPT Input Capture Register 2
//	0x024 32  CNT   GPT Counter Register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package gpt

const (
	EN       CR = 0x01 << 0  //+ GPT Enable
	ENMOD    CR = 0x01 << 1  //+ GPT Enable mode
	DBGEN    CR = 0x01 << 2  //+ GPT debug mode enable
	WAITEN   CR = 0x01 << 3  //+ GPT Wait Mode enable
	DOZEEN   CR = 0x01 << 4  //+ GPT Doze Mode Enable
	STOPEN   CR = 0x01 << 5  //+ GPT Stop Mode enable
	CLKSRC   CR = 0x07 << 6  //+ Clock Source select
	CLKSRC_0 CR = 0x00 << 6  //  No clock
	CLKSRC_1 CR = 0x01 << 6  //  Peripheral Clock (ipg_clk)
	CLKSRC_2 CR = 0x02 << 6  //  High Frequency Reference Clock (ipg_clk_highfreq)
	CLKSRC_3 CR = 0x03 << 6  //  External Clock
	CLKSRC_4 CR = 0x04 << 6  //  Low Frequency Reference Clock (ipg_clk_32k)
	CLKSRC_5 CR = 0x05 << 6  //  Crystal oscillator as Reference Clock (ipg_clk_24M)
	FRR      CR = 0x01 << 9  //+ Free-Run or Restart mode
	EN_24M   CR = 0x01 << 10 //+ Enable 24 MHz clock input from crystal
	SWR      CR = 0x01 << 15 //+ Software reset
	IM1      CR = 0x03 << 16 //+ See IM2
	IM2      CR = 0x03 << 18 //+ IM2 (bits 19-18, Input Capture Channel 2 operating mode) IM1 (bits 17-16, Input Capture Channel 1 operating mode) The IMn bit field determines the transition on the input pin (for Input capture channel n), which will trigger a capture event
	IM2_0    CR = 0x00 << 18 //  capture disabled
	IM2_1    CR = 0x01 << 18 //  capture on rising edge only
	IM2_2    CR = 0x02 << 18 //  capture on falling edge only
	IM2_3    CR = 0x03 << 18 //  capture on both edges
	OM1      CR = 0x07 << 20 //+ See OM3
	OM2      CR = 0x07 << 23 //+ See OM3
	OM3      CR = 0x07 << 26 //+ OM3 (bits 28-26) controls the Output Compare Channel 3 operating mode
	OM3_0    CR = 0x00 << 26 //  Output disconnected. No response on pin.
	OM3_1    CR = 0x01 << 26 //  Toggle output pin
	OM3_2    CR = 0x02 << 26 //  Clear output pin
	OM3_3    CR = 0x03 << 26 //  Set output pin
	OM3_4    CR = 0x04 << 26 //  Generate an active low pulse (that is one input clock wide) on the output pin.
	FO1      CR = 0x01 << 29 //+ See F03
	FO2      CR = 0x01 << 30 //+ See F03
	FO3      CR = 0x01 << 31 //+ FO3 Force Output Compare Channel 3 FO2 Force Output Compare Channel 2 FO1 Force Output Compare Channel 1 The FOn bit causes the pin action programmed for the timer Output Compare n pin (according to the OMn bits in this register)
)

const (
	ENn     = 0
	ENMODn  = 1
	DBGENn  = 2
	WAITENn = 3
	DOZEENn = 4
	STOPENn = 5
	CLKSRCn = 6
	FRRn    = 9
	EN_24Mn = 10
	SWRn    = 15
	IM1n    = 16
	IM2n    = 18
	OM1n    = 20
	OM2n    = 23
	OM3n    = 26
	FO1n    = 29
	FO2n    = 30
	FO3n    = 31
)

const (
	PRESCALER       PR = 0xFFF << 0 //+ Prescaler bits
	PRESCALER_0     PR = 0x00 << 0  //  Divide by 1
	PRESCALER_1     PR = 0x01 << 0  //  Divide by 2
	PRESCALER_4095  PR = 0xFFF << 0 //  Divide by 4096
	PRESCALER24M    PR = 0x0F << 12 //+ Prescaler bits
	PRESCALER24M_0  PR = 0x00 << 12 //  Divide by 1
	PRESCALER24M_1  PR = 0x01 << 12 //  Divide by 2
	PRESCALER24M_15 PR = 0x0F << 12 //  Divide by 16
)

const (
	PRESCALERn    = 0
	PRESCALER24Mn = 12
)

const (
	OF1 SR = 0x01 << 0 //+ See OF3
	OF2 SR = 0x01 << 1 //+ See OF3
	OF3 SR = 0x01 << 2 //+ OF3 Output Compare 3 Flag OF2 Output Compare 2 Flag OF1 Output Compare 1 Flag The OFn bit indicates that a compare event has occurred on Output Compare channel n
	IF1 SR = 0x01 << 3 //+ See IF2
	IF2 SR = 0x01 << 4 //+ IF2 Input capture 2 Flag IF1 Input capture 1 Flag The IFn bit indicates that a capture event has occurred on Input Capture channel n
	ROV SR = 0x01 << 5 //+ Rollover Flag
)

const (
	OF1n = 0
	OF2n = 1
	OF3n = 2
	IF1n = 3
	IF2n = 4
	ROVn = 5
)

const (
	OF1IE IR = 0x01 << 0 //+ See OF3IE
	OF2IE IR = 0x01 << 1 //+ See OF3IE
	OF3IE IR = 0x01 << 2 //+ OF3IE Output Compare 3 Interrupt Enable OF2IE Output Compare 2 Interrupt Enable OF1IE Output Compare 1 Interrupt Enable The OFnIE bit controls the Output Compare Channel n interrupt
	IF1IE IR = 0x01 << 3 //+ See IF2IE
	IF2IE IR = 0x01 << 4 //+ IF2IE Input capture 2 Interrupt Enable IF1IE Input capture 1 Interrupt Enable The IFnIE bit controls the IFnIE Input Capture n Interrupt Enable
	ROVIE IR = 0x01 << 5 //+ Rollover Interrupt Enable. The ROVIE bit controls the Rollover interrupt.
)

const (
	OF1IEn = 0
	OF2IEn = 1
	OF3IEn = 2
	IF1IEn = 3
	IF2IEn = 4
	ROVIEn = 5
)

const (
	COMP OCR1 = 0xFFFFFFFF << 0 //+ Compare Value
)

const (
	COMPn = 0
)

const (
	COMP OCR2 = 0xFFFFFFFF << 0 //+ Compare Value
)

const (
	COMPn = 0
)

const (
	COMP OCR3 = 0xFFFFFFFF << 0 //+ Compare Value
)

const (
	COMPn = 0
)

const (
	CAPT ICR1 = 0xFFFFFFFF << 0 //+ Capture Value
)

const (
	CAPTn = 0
)

const (
	CAPT ICR2 = 0xFFFFFFFF << 0 //+ Capture Value
)

const (
	CAPTn = 0
)

const (
	COUNT CNT = 0xFFFFFFFF << 0 //+ Counter Value. The COUNT bits show the current count value of the GPT counter.
)

const (
	COUNTn = 0
)
