// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

// VERID
const (
	FEATURE   uint32 = 0xFFFF << 0 //+ Module Identification Number
	FEATURE_4 uint32 = 0x04 << 0   //  Standard feature set supporting a 32-bit shift register.
	MINOR     uint32 = 0xFF << 16  //+ Minor Version Number
	MAJOR     uint32 = 0xFF << 24  //+ Major Version Number
)

const (
	FEATUREn = 0
	MINORn   = 16
	MAJORn   = 24
)

// PARAM
const (
	TXFIFO uint32 = 0xFF << 0  //+ Transmit FIFO Size
	RXFIFO uint32 = 0xFF << 8  //+ Receive FIFO Size
	PCSNUM uint32 = 0xFF << 16 //+ PCS Number
)

const (
	TXFIFOn = 0
	RXFIFOn = 8
	PCSNUMn = 16
)

type CR uint32

const (
	MEN   CR = 0x01 << 0 //+ Module Enable
	RST   CR = 0x01 << 1 //+ Software Reset
	DOZEN CR = 0x01 << 2 //+ Doze mode enable
	DBGEN CR = 0x01 << 3 //+ Debug Enable
	RTF   CR = 0x01 << 8 //+ Reset Transmit FIFO
	RRF   CR = 0x01 << 9 //+ Reset Receive FIFO
)

const (
	MENn   = 0
	RSTn   = 1
	DOZENn = 2
	DBGENn = 3
	RTFn   = 8
	RRFn   = 9
)

type SR uint32

const (
	TDF SR = 0x01 << 0  //+ Transmit Data Flag
	RDF SR = 0x01 << 1  //+ Receive Data Flag
	WCF SR = 0x01 << 8  //+ Word Complete Flag
	FCF SR = 0x01 << 9  //+ Frame Complete Flag
	TCF SR = 0x01 << 10 //+ Transfer Complete Flag
	TEF SR = 0x01 << 11 //+ Transmit Error Flag
	REF SR = 0x01 << 12 //+ Receive Error Flag
	DMF SR = 0x01 << 13 //+ Data Match Flag
	MBF SR = 0x01 << 24 //+ Module Busy Flag
)

const (
	TDFn = 0
	RDFn = 1
	WCFn = 8
	FCFn = 9
	TCFn = 10
	TEFn = 11
	REFn = 12
	DMFn = 13
	MBFn = 24
)

type IER uint32

const (
	TDIE IER = 0x01 << 0  //+ Transmit Data Interrupt Enable
	RDIE IER = 0x01 << 1  //+ Receive Data Interrupt Enable
	WCIE IER = 0x01 << 8  //+ Word Complete Interrupt Enable
	FCIE IER = 0x01 << 9  //+ Frame Complete Interrupt Enable
	TCIE IER = 0x01 << 10 //+ Transfer Complete Interrupt Enable
	TEIE IER = 0x01 << 11 //+ Transmit Error Interrupt Enable
	REIE IER = 0x01 << 12 //+ Receive Error Interrupt Enable
	DMIE IER = 0x01 << 13 //+ Data Match Interrupt Enable
)

const (
	TDIEn = 0
	RDIEn = 1
	WCIEn = 8
	FCIEn = 9
	TCIEn = 10
	TEIEn = 11
	REIEn = 12
	DMIEn = 13
)

type DER uint32

const (
	TDDE DER = 0x01 << 0 //+ Transmit Data DMA Enable
	RDDE DER = 0x01 << 1 //+ Receive Data DMA Enable
)

const (
	TDDEn = 0
	RDDEn = 1
)

type CFGR0 uint32

const (
	HREN    CFGR0 = 0x01 << 0 //+ Host Request Enable
	HRPOL   CFGR0 = 0x01 << 1 //+ Host Request Polarity
	HRSEL   CFGR0 = 0x01 << 2 //+ Host Request Select
	CIRFIFO CFGR0 = 0x01 << 8 //+ Circular FIFO Enable
	RDMO    CFGR0 = 0x01 << 9 //+ Receive Data Match Only
)

const (
	HRENn    = 0
	HRPOLn   = 1
	HRSELn   = 2
	CIRFIFOn = 8
	RDMOn    = 9
)

type CFGR1 uint32

const (
	MASTER  CFGR1 = 0x01 << 0  //+ Master Mode
	SAMPLE  CFGR1 = 0x01 << 1  //+ Sample Point
	AUTOPCS CFGR1 = 0x01 << 2  //+ Automatic PCS
	NOSTALL CFGR1 = 0x01 << 3  //+ No Stall
	PCSPOL  CFGR1 = 0x0F << 8  //+ Peripheral Chip Select Polarity
	PCS0H   CFGR1 = 0x01 << 8  //  PCS0 pin is active high
	PCS1H   CFGR1 = 0x02 << 8  //  PCS1 pin is active high
	PCS2H   CFGR1 = 0x04 << 8  //  PCS2 pin is active high
	PCS3H   CFGR1 = 0x08 << 8  //  PCS3 pin is active high
	MATCFG  CFGR1 = 0x07 << 16 //+ Match Configuration
	MATDIS  CFGR1 = 0x00 << 16 //  Match is disabled
	MAT0    CFGR1 = 0x02 << 16 //  Match if data[0]==MATCH0 || data[0]==MATCH1
	MATX    CFGR1 = 0x03 << 16 //  Match if data[x]==MATCH0 || data[x]==MATCH1
	MAT02   CFGR1 = 0x04 << 16 //  Match if data[0:2] == {MATCH0, MATCH1}
	MATX2   CFGR1 = 0x05 << 16 //  Match if data[x:x+2] == {MATCH0, MATCH1}
	MAT0M   CFGR1 = 0x06 << 16 //  Match if data[0]&MATCH1 == MATCH0&MATCH1
	MATXM   CFGR1 = 0x07 << 16 //  Match if data[x]&MATCH1 == MATCH0&MATCH1
	PINCFG  CFGR1 = 0x03 << 24 //+ Pin Configuration
	FD      CFGR1 = 0x00 << 24 //  SIN=Rx, SOUT=Tx in full-duplex mode
	HDSIN   CFGR1 = 0x01 << 24 //  1-bit half-duplex on SIN
	HDSOUT  CFGR1 = 0x02 << 24 //  1-bit half-duplex on SOUT
	FDSWAP  CFGR1 = 0x03 << 24 //  SIN=Tx, SOUT=Rx in full-duplex mode
	OUTCFG  CFGR1 = 0x01 << 26 //+ Output Config
	PCSDATA CFGR1 = 0x01 << 27 //+ Use PCS[3:2] as DATA[3:2] for 4-bit half-duplex mode
)

const (
	MASTERn  = 0
	SAMPLEn  = 1
	AUTOPCSn = 2
	NOSTALLn = 3
	PCSPOLn  = 8
	MATCFGn  = 16
	PINCFGn  = 24
	OUTCFGn  = 26
	PCSCFGn  = 27
)

type CCR uint32

const (
	SCKDIV CCR = 0xFF << 0  //+ SCK Divider
	DBT    CCR = 0xFF << 8  //+ Delay Between Transfers
	PCSSCK CCR = 0xFF << 16 //+ PCS-to-SCK Delay
	SCKPCS CCR = 0xFF << 24 //+ SCK-to-PCS Delay
)

const (
	SCKDIVn = 0
	DBTn    = 8
	PCSSCKn = 16
	SCKPCSn = 24
)

type FCR uint32

const (
	TXWATER FCR = 0x0F << 0  //+ Transmit FIFO Watermark
	RXWATER FCR = 0x0F << 16 //+ Receive FIFO Watermark
)

const (
	TXWATERn = 0
	RXWATERn = 16
)

type FSR uint32

const (
	TXCOUNT FSR = 0x1F << 0  //+ Transmit FIFO Count
	RXCOUNT FSR = 0x1F << 16 //+ Receive FIFO Count
)

const (
	TXCOUNTn = 0
	RXCOUNTn = 16
)

type TCR uint32

const (
	FRAMESZ   TCR = 0xFFF << 0 //+ Frame Size
	WIDTH     TCR = 0x03 << 16 //+ Transfer Width
	WIDTH0    TCR = 0x00 << 16 //  1 bit transfer
	WIDTH1    TCR = 0x01 << 16 //  2 bit transfer
	WIDTH2    TCR = 0x02 << 16 //  4 bit transfer
	TXMSK     TCR = 0x01 << 18 //+ Transmit Data Mask
	RXMSK     TCR = 0x01 << 19 //+ Receive Data Mask
	CONTC     TCR = 0x01 << 20 //+ Continuing Command
	CONT      TCR = 0x01 << 21 //+ Continuous Transfer
	BYSW      TCR = 0x01 << 22 //+ Byte Swap
	LSBF      TCR = 0x01 << 23 //+ LSB First
	TPCS       TCR = 0x03 << 24 //+ Peripheral Chip Select
	TPCS0      TCR = 0x00 << 24 //  Transfer using LPSPI_PCS[0]
	TPCS1      TCR = 0x01 << 24 //  Transfer using LPSPI_PCS[1]
	TPCS2      TCR = 0x02 << 24 //  Transfer using LPSPI_PCS[2]
	TPCS3      TCR = 0x03 << 24 //  Transfer using LPSPI_PCS[3]
	PRESCALE  TCR = 0x07 << 27 //+ Prescaler Value
	PREDIV1   TCR = 0x00 << 27 //  Divide by 1
	PREDIV2   TCR = 0x01 << 27 //  Divide by 2
	PREDIV4   TCR = 0x02 << 27 //  Divide by 4
	PREDIV8   TCR = 0x03 << 27 //  Divide by 8
	PREDIV16  TCR = 0x04 << 27 //  Divide by 16
	PREDIV32  TCR = 0x05 << 27 //  Divide by 32
	PREDIV64  TCR = 0x06 << 27 //  Divide by 64
	PREDIV128 TCR = 0x07 << 27 //  Divide by 128
	CPHA      TCR = 0x01 << 30 //+ Clock Phase
	CPHA0     TCR = 0x00 << 30 //  Capture data on the leading and change on the following edge of SCK
	CPHA1     TCR = 0x01 << 30 //  Change data on the leading and capture on the following edge of SCK
	CPOL      TCR = 0x01 << 31 //+ Clock Polarity
	CPOL0     TCR = 0x00 << 31 //  The inactive state value of SCK is low
	CPOL1     TCR = 0x01 << 31 //  The inactive state value of SCK is high
)

const (
	FRAMESZn  = 0
	WIDTHn    = 16
	TXMSKn    = 18
	RXMSKn    = 19
	CONTCn    = 20
	CONTn     = 21
	BYSWn     = 22
	LSBFn     = 23
	PCSn      = 24
	PRESCALEn = 27
	CPHAn     = 30
	CPOLn     = 31
)

type RSR uint32

const (
	SOF     RSR = 0x01 << 0 //+ Start Of Frame
	RXEMPTY RSR = 0x01 << 1 //+ RX FIFO Empty
)

const (
	SOFn     = 0
	RXEMPTYn = 1
)
