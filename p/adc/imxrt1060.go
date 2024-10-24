// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package adc provides access to the registers of the ADC peripheral.
//
// Instances:
//
//	ADC1  ADC1_BASE  -  ADC1*  Analog-to-Digital Converter
//	ADC2  ADC2_BASE  -  ADC2*  Analog-to-Digital Converter
//
// Registers:
//
//	0x000 32  HC0    Control register for hardware triggers
//	0x004 32  HC[7]  Control register for hardware triggers
//	0x020 32  HS     Status register for HW triggers
//	0x024 32  R0     Data result register for HW triggers
//	0x028 32  R[7]   Data result register for HW triggers
//	0x044 32  CFG    Configuration register
//	0x048 32  GC     General control register
//	0x04C 32  GS     General status register
//	0x050 32  CV     Compare value register
//	0x054 32  OFS    Offset correction value register
//	0x058 32  CAL    Calibration value register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package adc

const (
	ADCH    HC0 = 0x1F << 0 //+ Input Channel Select
	ADCH_16 HC0 = 0x10 << 0 //  External channel selection from ADC_ETC
	ADCH_25 HC0 = 0x19 << 0 //  VREFSH = internal channel, for ADC self-test, hard connected to VRH internally
	ADCH_31 HC0 = 0x1F << 0 //  Conversion Disabled. Hardware Triggers will not initiate any conversion.
	AIEN    HC0 = 0x01 << 7 //+ Conversion Complete Interrupt Enable/Disable Control
)

const (
	ADCHn = 0
	AIENn = 7
)

const (
	ADCH    HC = 0x1F << 0 //+ Input Channel Select
	ADCH_16 HC = 0x10 << 0 //  External channel selection from ADC_ETC
	ADCH_25 HC = 0x19 << 0 //  VREFSH = internal channel, for ADC self-test, hard connected to VRH internally
	ADCH_31 HC = 0x1F << 0 //  Conversion Disabled. Hardware Triggers will not initiate any conversion.
	AIEN    HC = 0x01 << 7 //+ Conversion Complete Interrupt Enable/Disable Control
)

const (
	ADCHn = 0
	AIENn = 7
)

const (
	COCO0 HS = 0x01 << 0 //+ Conversion Complete Flag
)

const (
	COCO0n = 0
)

const (
	CDATA R0 = 0xFFF << 0 //+ Data (result of an ADC conversion)
)

const (
	CDATAn = 0
)

const (
	CDATA R = 0xFFF << 0 //+ Data (result of an ADC conversion)
)

const (
	CDATAn = 0
)

const (
	ADICLK   CFG = 0x03 << 0  //+ Input Clock Select
	ADICLK_0 CFG = 0x00 << 0  //  IPG clock
	ADICLK_1 CFG = 0x01 << 0  //  IPG clock divided by 2
	ADICLK_3 CFG = 0x03 << 0  //  Asynchronous clock (ADACK)
	MODE     CFG = 0x03 << 2  //+ Conversion Mode Selection
	MODE_0   CFG = 0x00 << 2  //  8-bit conversion
	MODE_1   CFG = 0x01 << 2  //  10-bit conversion
	MODE_2   CFG = 0x02 << 2  //  12-bit conversion
	ADLSMP   CFG = 0x01 << 4  //+ Long Sample Time Configuration
	ADIV     CFG = 0x03 << 5  //+ Clock Divide Select
	ADIV_0   CFG = 0x00 << 5  //  Input clock
	ADIV_1   CFG = 0x01 << 5  //  Input clock / 2
	ADIV_2   CFG = 0x02 << 5  //  Input clock / 4
	ADIV_3   CFG = 0x03 << 5  //  Input clock / 8
	ADLPC    CFG = 0x01 << 7  //+ Low-Power Configuration
	ADSTS    CFG = 0x03 << 8  //+ Defines the sample time duration
	ADSTS_0  CFG = 0x00 << 8  //  Sample period (ADC clocks) = 2 if ADLSMP=0b Sample period (ADC clocks) = 12 if ADLSMP=1b
	ADSTS_1  CFG = 0x01 << 8  //  Sample period (ADC clocks) = 4 if ADLSMP=0b Sample period (ADC clocks) = 16 if ADLSMP=1b
	ADSTS_2  CFG = 0x02 << 8  //  Sample period (ADC clocks) = 6 if ADLSMP=0b Sample period (ADC clocks) = 20 if ADLSMP=1b
	ADSTS_3  CFG = 0x03 << 8  //  Sample period (ADC clocks) = 8 if ADLSMP=0b Sample period (ADC clocks) = 24 if ADLSMP=1b
	ADHSC    CFG = 0x01 << 10 //+ High Speed Configuration
	REFSEL   CFG = 0x03 << 11 //+ Voltage Reference Selection
	REFSEL_0 CFG = 0x00 << 11 //  Selects VREFH/VREFL as reference voltage.
	ADTRG    CFG = 0x01 << 13 //+ Conversion Trigger Select
	AVGS     CFG = 0x03 << 14 //+ Hardware Average select
	AVGS_0   CFG = 0x00 << 14 //  4 samples averaged
	AVGS_1   CFG = 0x01 << 14 //  8 samples averaged
	AVGS_2   CFG = 0x02 << 14 //  16 samples averaged
	AVGS_3   CFG = 0x03 << 14 //  32 samples averaged
	OVWREN   CFG = 0x01 << 16 //+ Data Overwrite Enable
)

const (
	ADICLKn = 0
	MODEn   = 2
	ADLSMPn = 4
	ADIVn   = 5
	ADLPCn  = 7
	ADSTSn  = 8
	ADHSCn  = 10
	REFSELn = 11
	ADTRGn  = 13
	AVGSn   = 14
	OVWRENn = 16
)

const (
	ADACKEN GC = 0x01 << 0 //+ Asynchronous clock output enable
	DMAEN   GC = 0x01 << 1 //+ DMA Enable
	ACREN   GC = 0x01 << 2 //+ Compare Function Range Enable
	ACFGT   GC = 0x01 << 3 //+ Compare Function Greater Than Enable
	ACFE    GC = 0x01 << 4 //+ Compare Function Enable
	AVGE    GC = 0x01 << 5 //+ Hardware average enable
	ADCO    GC = 0x01 << 6 //+ Continuous Conversion Enable
	CAL     GC = 0x01 << 7 //+ Calibration
)

const (
	ADACKENn = 0
	DMAENn   = 1
	ACRENn   = 2
	ACFGTn   = 3
	ACFEn    = 4
	AVGEn    = 5
	ADCOn    = 6
	CALn     = 7
)

const (
	ADACT GS = 0x01 << 0 //+ Conversion Active
	CALF  GS = 0x01 << 1 //+ Calibration Failed Flag
	AWKST GS = 0x01 << 2 //+ Asynchronous wakeup interrupt status
)

const (
	ADACTn = 0
	CALFn  = 1
	AWKSTn = 2
)

const (
	CV1 CV = 0xFFF << 0  //+ Compare Value 1
	CV2 CV = 0xFFF << 16 //+ Compare Value 2
)

const (
	CV1n = 0
	CV2n = 16
)

const (
	OFS  OFS = 0xFFF << 0 //+ Offset value
	SIGN OFS = 0x01 << 12 //+ Sign bit
)

const (
	OFSn  = 0
	SIGNn = 12
)

const (
	CAL_CODE CAL = 0x0F << 0 //+ Calibration Result Value
)

const (
	CAL_CODEn = 0
)
