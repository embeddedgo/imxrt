// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

// VERID
const (
	FEATURE   uint32 = 0xFFFF << 0 //+ Feature Specification Number
	FEATURE_2 uint32 = 0x02 << 0   //  Master only, with standard feature set
	FEATURE_3 uint32 = 0x03 << 0   //  Master and slave, with standard feature set
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
	MTXFIFO uint32 = 0x0F << 0 //+ Master Transmit FIFO Size
	MRXFIFO uint32 = 0x0F << 8 //+ Master Receive FIFO Size
)

const (
	MTXFIFOn = 0
	MRXFIFOn = 8
)

type MCR uint32

const (
	MEN    MCR = 0x01 << 0 //+ Master Enable
	MRST   MCR = 0x01 << 1 //+ Software Reset
	MDOZEN MCR = 0x01 << 2 //+ Doze mode enable
	MDBGEN MCR = 0x01 << 3 //+ Debug Enable
	MRTF   MCR = 0x01 << 8 //+ Reset Transmit FIFO
	MRRF   MCR = 0x01 << 9 //+ Reset Receive FIFO
)

const (
	MENn    = 0
	MRSTn   = 1
	MDOZENn = 2
	MDBGENn = 3
	MRTFn   = 8
	MRRFn   = 9
)

type MSR uint32

const (
	MTDF  MSR = 0x01 << 0  //+ Transmit Data Flag
	MRDF  MSR = 0x01 << 1  //+ Receive Data Flag
	MEPF  MSR = 0x01 << 8  //+ End Packet Flag
	MSDF  MSR = 0x01 << 9  //+ STOP Detect Flag
	MNDF  MSR = 0x01 << 10 //+ NACK Detect Flag
	MALF  MSR = 0x01 << 11 //+ Arbitration Lost Flag
	MFEF  MSR = 0x01 << 12 //+ FIFO Error Flag
	MPLTF MSR = 0x01 << 13 //+ Pin Low Timeout Flag
	MDMF  MSR = 0x01 << 14 //+ Data Match Flag
	MBF   MSR = 0x01 << 24 //+ Master Busy Flag
	MBBF  MSR = 0x01 << 25 //+ Bus Busy Flag
)

const (
	MTDFn  = 0
	MRDFn  = 1
	MEPFn  = 8
	MSDFn  = 9
	MNDFn  = 10
	MALFn  = 11
	MFEFn  = 12
	MPLTFn = 13
	MDMFn  = 14
	MBFn   = 24
	MBBFn  = 25
)

type MCFGR0 uint32

const (
	HREN    MCFGR0 = 0x01 << 0 //+ Host Request Enable
	HRPOL   MCFGR0 = 0x01 << 1 //+ Host Request Polarity
	HRSEL   MCFGR0 = 0x01 << 2 //+ Host Request Select
	CIRFIFO MCFGR0 = 0x01 << 8 //+ Circular FIFO Enable
	RDMO    MCFGR0 = 0x01 << 9 //+ Receive Data Match Only
)

const (
	HRENn    = 0
	HRPOLn   = 1
	HRSELn   = 2
	CIRFIFOn = 8
	RDMOn    = 9
)

type MCFGR1 uint32

const (
	MPRESCALE              MCFGR1 = 0x07 << 0  //+ Prescaler
	Div1                   MCFGR1 = 0x00 << 0  //  Divide by 1
	Div2                   MCFGR1 = 0x01 << 0  //  Divide by 2
	Div4                   MCFGR1 = 0x02 << 0  //  Divide by 4
	Div8                   MCFGR1 = 0x03 << 0  //  Divide by 8
	Div16                  MCFGR1 = 0x04 << 0  //  Divide by 16
	Div32                  MCFGR1 = 0x05 << 0  //  Divide by 32
	Div64                  MCFGR1 = 0x06 << 0  //  Divide by 64
	Div128                 MCFGR1 = 0x07 << 0  //  Divide by 128
	MAUTOSTOP              MCFGR1 = 0x01 << 8  //+ Automatic STOP Generation
	MIGNACK                MCFGR1 = 0x01 << 9  //+ IGNACK
	MTIMECFG               MCFGR1 = 0x01 << 10 //+ Timeout Configuration
	MATCFG                 MCFGR1 = 0x07 << 16 //+ Match Configuration
	Disable                MCFGR1 = 0x00 << 16 //  Match is disabled
	D0eqM0_or_D0eqM1       MCFGR1 = 0x02 << 16 //  Match is enabled (1st data word equals MATCH0 OR MATCH1)
	DXeqM0_or_DXeqM1       MCFGR1 = 0x03 << 16 //  Match is enabled (any data word equals MATCH0 OR MATCH1)
	D0D1_eq_M0M1           MCFGR1 = 0x04 << 16 //  Match is enabled (1st data word equals MATCH0 AND 2nd data word equals MATCH1)
	DXDX1_eq_M0M1          MCFGR1 = 0x05 << 16 //  Match is enabled (any data word equals MATCH0 AND next data word equals MATCH1)
	D0andM0_eq_M0andM1     MCFGR1 = 0x06 << 16 //  Match is enabled (1st data word AND MATCH1 equals MATCH0 AND MATCH1)
	DXandM0_eq_M0andM1     MCFGR1 = 0x07 << 16 //  Match is enabled (any data word AND MATCH1 equals MATCH0 AND MATCH1)
	MPINCFG                MCFGR1 = 0x07 << 24 //+ Pin Configuration
	OpenDrain2pin          MCFGR1 = 0x00 << 24 //  2-pin open drain mode
	OutputOnly2pin         MCFGR1 = 0x01 << 24 //  2-pin output only mode (ultra-fast mode)
	PushPull2pin           MCFGR1 = 0x02 << 24 //  2-pin push-pull mode
	PushPull4pin           MCFGR1 = 0x03 << 24 //  4-pin push-pull mode
	OpenDrain2pinSepSlave  MCFGR1 = 0x04 << 24 //  2-pin open drain mode with separate LPI2C slave
	OutputOnly2pinSepSlave MCFGR1 = 0x05 << 24 //  2-pin output only mode (ultra-fast mode) with separate LPI2C slave
	PushPull2pinSepSlave   MCFGR1 = 0x06 << 24 //  2-pin push-pull mode with separate LPI2C slave
	PushPull4pinInverted   MCFGR1 = 0x07 << 24 //  4-pin push-pull mode (inverted outputs)
)

const (
	MPRESCALEn = 0
	MAUTOSTOPn = 8
	MIGNACKn   = 9
	MTIMECFGn  = 10
	MATCFGn    = 16
	MPINCFGn   = 24
)

type MCFGR2 uint32

const (
	MBUSIDLE MCFGR2 = 0xFFF << 0 //+ Bus Idle Timeout
	MFILTSCL MCFGR2 = 0x0F << 16 //+ Glitch Filter SCL
	MFILTSDA MCFGR2 = 0x0F << 24 //+ Glitch Filter SDA
)

const (
	MBUSIDLEn = 0
	MFILTSCLn = 16
	MFILTSDAn = 24
)

type MCFGR3 uint32

const (
	PINLOW MCFGR3 = 0xFFF << 8 //+ Pin Low Timeout
)

const (
	PINLOWn = 8
)

type MDMR uint32

const (
	MATCH0 MDMR = 0xFF << 0  //+ Match 0 Value
	MATCH1 MDMR = 0xFF << 16 //+ Match 1 Value
)

const (
	MATCH0n = 0
	MATCH1n = 16
)

type MCCR uint32

const (
	CLKLO   MCCR = 0x3F << 0  //+ Clock Low Period
	CLKHI   MCCR = 0x3F << 8  //+ Clock High Period
	SETHOLD MCCR = 0x3F << 16 //+ Setup Hold Delay
	DATAVD  MCCR = 0x3F << 24 //+ Data Valid Delay
)

const (
	CLKLOn   = 0
	CLKHIn   = 8
	SETHOLDn = 16
	DATAVDn  = 24
)

type MFCR uint32

const (
	TXWATER MFCR = 0x03 << 0  //+ Transmit FIFO Watermark
	RXWATER MFCR = 0x03 << 16 //+ Receive FIFO Watermark
)

const (
	TXWATERn = 0
	RXWATERn = 16
)

type MFSR uint32

const (
	TXCOUNT MFSR = 0x07 << 0  //+ Transmit FIFO Count
	RXCOUNT MFSR = 0x07 << 16 //+ Receive FIFO Count
)

const (
	TXCOUNTn = 0
	RXCOUNTn = 16
)

const (
	DATA        int16 = 0xFF << 0 //+ Transmit Data
	CMD         int16 = 0x07 << 8 //+ Command Data
	Send        int16 = 0x00 << 8 //  Transmit DATA[7:0]
	Recv        int16 = 0x01 << 8 //  Receive (DATA[7:0] + 1) bytes
	Stop        int16 = 0x02 << 8 //  Generate STOP condition
	Discard     int16 = 0x03 << 8 //  Receive and discard (DATA[7:0] + 1) bytes
	Start       int16 = 0x04 << 8 //  Generate (repeated) START and transmit address in DATA[7:0]
	StartNACK   int16 = 0x05 << 8 //  Generate (repeated) START and transmit address in DATA[7:0]. This transfer expects a NACK to be returned.
	StartHS     int16 = 0x06 << 8 //  Generate (repeated) START and transmit address in DATA[7:0] using high speed mode
	StartHSNACK int16 = 0x07 << 8 //  Generate (repeated) START and transmit address in DATA[7:0] using high speed mode. This transfer expects a NACK to be returned.
)

const (
	DATAn = 0
	CMDn  = 8
)

type SCR uint32

const (
	SEN     SCR = 0x01 << 0 //+ Slave Enable
	SRST    SCR = 0x01 << 1 //+ Software Reset
	SFILTEN SCR = 0x01 << 4 //+ Filter Enable
	SFILTDZ SCR = 0x01 << 5 //+ Filter Doze Enable
	SRTF    SCR = 0x01 << 8 //+ Reset Transmit FIFO
	SRRF    SCR = 0x01 << 9 //+ Reset Receive FIFO
)

const (
	SENn     = 0
	SRSTn    = 1
	SFILTENn = 4
	SFILTDZn = 5
	SRTFn    = 8
	SRRFn    = 9
)

type SSR uint32

const (
	STDF  SSR = 0x01 << 0  //+ Transmit Data Flag
	SRDF  SSR = 0x01 << 1  //+ Receive Data Flag
	SAVF  SSR = 0x01 << 2  //+ Address Valid Flag
	STAF  SSR = 0x01 << 3  //+ Transmit ACK Flag
	SRSF  SSR = 0x01 << 8  //+ Repeated Start Flag
	SSDF  SSR = 0x01 << 9  //+ STOP Detect Flag
	SBEF  SSR = 0x01 << 10 //+ Bit Error Flag
	SFEF  SSR = 0x01 << 11 //+ FIFO Error Flag
	SAM0F SSR = 0x01 << 12 //+ Address Match 0 Flag
	SAM1F SSR = 0x01 << 13 //+ Address Match 1 Flag
	SGCF  SSR = 0x01 << 14 //+ General Call Flag
	SSARF SSR = 0x01 << 15 //+ SMBus Alert Response Flag
	SBF   SSR = 0x01 << 24 //+ Slave Busy Flag
	SBBF  SSR = 0x01 << 25 //+ Bus Busy Flag
)

const (
	STDFn  = 0
	SRDFn  = 1
	SAVFn  = 2
	STAFn  = 3
	SRSFn  = 8
	SSDFn  = 9
	SBEFn  = 10
	SFEFn  = 11
	SAM0Fn = 12
	SAM1Fn = 13
	SGCFn  = 14
	SSARFn = 15
	SBFn   = 24
	SBBFn  = 25
)

type DER uint32

const (
	TDDE DER = 0x01 << 0 //+ Transmit Data DMA Enable
	RDDE DER = 0x01 << 1 //+ Receive Data DMA Enable
	AVDE DER = 0x01 << 2 //+ Address Valid DMA Enable (slave only)
)

const (
	TDDEn = 0
	RDDEn = 1
	AVDEn = 2
)

type SCFGR1 uint32

const (
	SADRSTALL SCFGR1 = 0x01 << 0  //+ Address SCL Stall
	SRXSTALL  SCFGR1 = 0x01 << 1  //+ RX SCL Stall
	STXDSTALL SCFGR1 = 0x01 << 2  //+ TX Data SCL Stall
	SACKSTALL SCFGR1 = 0x01 << 3  //+ ACK SCL Stall
	SGCEN     SCFGR1 = 0x01 << 8  //+ General Call Enable
	SSAEN     SCFGR1 = 0x01 << 9  //+ SMBus Alert Enable
	STXCFG    SCFGR1 = 0x01 << 10 //+ Transmit Flag Configuration
	SRXCFG    SCFGR1 = 0x01 << 11 //+ Receive Data Configuration
	SIGNACK   SCFGR1 = 0x01 << 12 //+ Ignore NACK
	SHSMEN    SCFGR1 = 0x01 << 13 //+ High Speed Mode Enable
	SADDRCFG  SCFGR1 = 0x07 << 16 //+ Address Configuration
	ADDRCFG_0 SCFGR1 = 0x00 << 16 //  Address match 0 (7-bit)
	ADDRCFG_1 SCFGR1 = 0x01 << 16 //  Address match 0 (10-bit)
	ADDRCFG_2 SCFGR1 = 0x02 << 16 //  Address match 0 (7-bit) or Address match 1 (7-bit)
	ADDRCFG_3 SCFGR1 = 0x03 << 16 //  Address match 0 (10-bit) or Address match 1 (10-bit)
	ADDRCFG_4 SCFGR1 = 0x04 << 16 //  Address match 0 (7-bit) or Address match 1 (10-bit)
	ADDRCFG_5 SCFGR1 = 0x05 << 16 //  Address match 0 (10-bit) or Address match 1 (7-bit)
	ADDRCFG_6 SCFGR1 = 0x06 << 16 //  From Address match 0 (7-bit) to Address match 1 (7-bit)
	ADDRCFG_7 SCFGR1 = 0x07 << 16 //  From Address match 0 (10-bit) to Address match 1 (10-bit)
)

const (
	SADRSTALLn = 0
	SRXSTALLn  = 1
	STXDSTALLn = 2
	SACKSTALLn = 3
	SGCENn     = 8
	SSAENn     = 9
	STXCFGn    = 10
	SRXCFGn    = 11
	SIGNACKn   = 12
	SHSMENn    = 13
	SADDRCFGn  = 16
)

type SCFGR2 uint32

const (
	SCLKHOLD SCFGR2 = 0x0F << 0  //+ Clock Hold Time
	SDATAVD  SCFGR2 = 0x3F << 8  //+ Data Valid Delay
	SFILTSCL SCFGR2 = 0x0F << 16 //+ Glitch Filter SCL
	SFILTSDA SCFGR2 = 0x0F << 24 //+ Glitch Filter SDA
)

const (
	SCLKHOLDn = 0
	SDATAVDn  = 8
	SFILTSCLn = 16
	SFILTSDAn = 24
)

type SAMR uint32

const (
	ADDR0 SAMR = 0x3FF << 1  //+ Address 0 Value
	ADDR1 SAMR = 0x3FF << 17 //+ Address 1 Value
)

const (
	ADDR0n = 1
	ADDR1n = 17
)

type SASR uint32

const (
	RADDR SASR = 0x7FF << 0 //+ Received Address
	ANV   SASR = 0x01 << 14 //+ Address Not Valid
)

const (
	RADDRn = 0
	ANVn   = 14
)

type STAR uint32

const (
	TXNACK STAR = 0x01 << 0 //+ Transmit NACK
)

const (
	TXNACKn = 0
)

type RDR uint32

const (
	RXDATA  RDR = 0xFF << 0  //+ Receive Data
	RXEMPTY RDR = 0x01 << 14 //+ RX Empty
	SOF     RDR = 0x01 << 15 //+ Start Of Frame (slave only)
)

const (
	RXDATAn  = 0
	RXEMPTYn = 14
	SOFn     = 15
)
