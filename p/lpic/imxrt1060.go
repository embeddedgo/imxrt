// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package lpic provides access to the registers of the LPI2C peripheral.
//
// Instances:
//
//	LPI2C1  LPI2C1_BASE  -  LPI2C1*  LPI2C
//	LPI2C2  LPI2C2_BASE  -  LPI2C2*  LPI2C
//	LPI2C3  LPI2C3_BASE  -  LPI2C3*  LPI2C
//	LPI2C4  LPI2C4_BASE  -  LPI2C4*  LPI2C
//
// Registers:
//
//	0x000 32  VERID   Version ID Register
//	0x004 32  PARAM   Parameter Register
//	0x010 32  MCR     Master Control Register
//	0x014 32  MSR     Master Status Register
//	0x018 32  MIER    Master Interrupt Enable Register
//	0x01C 32  MDER    Master DMA Enable Register
//	0x020 32  MCFGR0  Master Configuration Register 0
//	0x024 32  MCFGR1  Master Configuration Register 1
//	0x028 32  MCFGR2  Master Configuration Register 2
//	0x02C 32  MCFGR3  Master Configuration Register 3
//	0x040 32  MDMR    Master Data Match Register
//	0x048 32  MCCR0   Master Clock Configuration Register 0
//	0x050 32  MCCR1   Master Clock Configuration Register 1
//	0x058 32  MFCR    Master FIFO Control Register
//	0x05C 32  MFSR    Master FIFO Status Register
//	0x060 32  MTDR    Master Transmit Data Register
//	0x070 32  MRDR    Master Receive Data Register
//	0x110 32  SCR     Slave Control Register
//	0x114 32  SSR     Slave Status Register
//	0x118 32  SIER    Slave Interrupt Enable Register
//	0x11C 32  SDER    Slave DMA Enable Register
//	0x124 32  SCFGR1  Slave Configuration Register 1
//	0x128 32  SCFGR2  Slave Configuration Register 2
//	0x140 32  SAMR    Slave Address Match Register
//	0x150 32  SASR    Slave Address Status Register
//	0x154 32  STAR    Slave Transmit ACK Register
//	0x160 32  STDR    Slave Transmit Data Register
//	0x170 32  SRDR    Slave Receive Data Register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package lpic

const (
	FEATURE   VERID = 0xFFFF << 0 //+ Feature Specification Number
	FEATURE_2 VERID = 0x02 << 0   //  Master only, with standard feature set
	FEATURE_3 VERID = 0x03 << 0   //  Master and slave, with standard feature set
	MINOR     VERID = 0xFF << 16  //+ Minor Version Number
	MAJOR     VERID = 0xFF << 24  //+ Major Version Number
)

const (
	FEATUREn = 0
	MINORn   = 16
	MAJORn   = 24
)

const (
	MTXFIFO PARAM = 0x0F << 0 //+ Master Transmit FIFO Size
	MRXFIFO PARAM = 0x0F << 8 //+ Master Receive FIFO Size
)

const (
	MTXFIFOn = 0
	MRXFIFOn = 8
)

const (
	MEN   MCR = 0x01 << 0 //+ Master Enable
	RST   MCR = 0x01 << 1 //+ Software Reset
	DOZEN MCR = 0x01 << 2 //+ Doze mode enable
	DBGEN MCR = 0x01 << 3 //+ Debug Enable
	RTF   MCR = 0x01 << 8 //+ Reset Transmit FIFO
	RRF   MCR = 0x01 << 9 //+ Reset Receive FIFO
)

const (
	MENn   = 0
	RSTn   = 1
	DOZENn = 2
	DBGENn = 3
	RTFn   = 8
	RRFn   = 9
)

const (
	TDF  MSR = 0x01 << 0  //+ Transmit Data Flag
	RDF  MSR = 0x01 << 1  //+ Receive Data Flag
	EPF  MSR = 0x01 << 8  //+ End Packet Flag
	SDF  MSR = 0x01 << 9  //+ STOP Detect Flag
	NDF  MSR = 0x01 << 10 //+ NACK Detect Flag
	ALF  MSR = 0x01 << 11 //+ Arbitration Lost Flag
	FEF  MSR = 0x01 << 12 //+ FIFO Error Flag
	PLTF MSR = 0x01 << 13 //+ Pin Low Timeout Flag
	DMF  MSR = 0x01 << 14 //+ Data Match Flag
	MBF  MSR = 0x01 << 24 //+ Master Busy Flag
	BBF  MSR = 0x01 << 25 //+ Bus Busy Flag
)

const (
	TDFn  = 0
	RDFn  = 1
	EPFn  = 8
	SDFn  = 9
	NDFn  = 10
	ALFn  = 11
	FEFn  = 12
	PLTFn = 13
	DMFn  = 14
	MBFn  = 24
	BBFn  = 25
)

const (
	TDIE  MIER = 0x01 << 0  //+ Transmit Data Interrupt Enable
	RDIE  MIER = 0x01 << 1  //+ Receive Data Interrupt Enable
	EPIE  MIER = 0x01 << 8  //+ End Packet Interrupt Enable
	SDIE  MIER = 0x01 << 9  //+ STOP Detect Interrupt Enable
	NDIE  MIER = 0x01 << 10 //+ NACK Detect Interrupt Enable
	ALIE  MIER = 0x01 << 11 //+ Arbitration Lost Interrupt Enable
	FEIE  MIER = 0x01 << 12 //+ FIFO Error Interrupt Enable
	PLTIE MIER = 0x01 << 13 //+ Pin Low Timeout Interrupt Enable
	DMIE  MIER = 0x01 << 14 //+ Data Match Interrupt Enable
)

const (
	TDIEn  = 0
	RDIEn  = 1
	EPIEn  = 8
	SDIEn  = 9
	NDIEn  = 10
	ALIEn  = 11
	FEIEn  = 12
	PLTIEn = 13
	DMIEn  = 14
)

const (
	TDDE MDER = 0x01 << 0 //+ Transmit Data DMA Enable
	RDDE MDER = 0x01 << 1 //+ Receive Data DMA Enable
)

const (
	TDDEn = 0
	RDDEn = 1
)

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

const (
	PRESCALE   MCFGR1 = 0x07 << 0  //+ Prescaler
	PRESCALE_0 MCFGR1 = 0x00 << 0  //  Divide by 1
	PRESCALE_1 MCFGR1 = 0x01 << 0  //  Divide by 2
	PRESCALE_2 MCFGR1 = 0x02 << 0  //  Divide by 4
	PRESCALE_3 MCFGR1 = 0x03 << 0  //  Divide by 8
	PRESCALE_4 MCFGR1 = 0x04 << 0  //  Divide by 16
	PRESCALE_5 MCFGR1 = 0x05 << 0  //  Divide by 32
	PRESCALE_6 MCFGR1 = 0x06 << 0  //  Divide by 64
	PRESCALE_7 MCFGR1 = 0x07 << 0  //  Divide by 128
	AUTOSTOP   MCFGR1 = 0x01 << 8  //+ Automatic STOP Generation
	IGNACK     MCFGR1 = 0x01 << 9  //+ IGNACK
	TIMECFG    MCFGR1 = 0x01 << 10 //+ Timeout Configuration
	MATCFG     MCFGR1 = 0x07 << 16 //+ Match Configuration
	MATCFG_0   MCFGR1 = 0x00 << 16 //  Match is disabled
	MATCFG_2   MCFGR1 = 0x02 << 16 //  Match is enabled (1st data word equals MATCH0 OR MATCH1)
	MATCFG_3   MCFGR1 = 0x03 << 16 //  Match is enabled (any data word equals MATCH0 OR MATCH1)
	MATCFG_4   MCFGR1 = 0x04 << 16 //  Match is enabled (1st data word equals MATCH0 AND 2nd data word equals MATCH1)
	MATCFG_5   MCFGR1 = 0x05 << 16 //  Match is enabled (any data word equals MATCH0 AND next data word equals MATCH1)
	MATCFG_6   MCFGR1 = 0x06 << 16 //  Match is enabled (1st data word AND MATCH1 equals MATCH0 AND MATCH1)
	MATCFG_7   MCFGR1 = 0x07 << 16 //  Match is enabled (any data word AND MATCH1 equals MATCH0 AND MATCH1)
	PINCFG     MCFGR1 = 0x07 << 24 //+ Pin Configuration
	PINCFG_0   MCFGR1 = 0x00 << 24 //  2-pin open drain mode
	PINCFG_1   MCFGR1 = 0x01 << 24 //  2-pin output only mode (ultra-fast mode)
	PINCFG_2   MCFGR1 = 0x02 << 24 //  2-pin push-pull mode
	PINCFG_3   MCFGR1 = 0x03 << 24 //  4-pin push-pull mode
	PINCFG_4   MCFGR1 = 0x04 << 24 //  2-pin open drain mode with separate LPI2C slave
	PINCFG_5   MCFGR1 = 0x05 << 24 //  2-pin output only mode (ultra-fast mode) with separate LPI2C slave
	PINCFG_6   MCFGR1 = 0x06 << 24 //  2-pin push-pull mode with separate LPI2C slave
	PINCFG_7   MCFGR1 = 0x07 << 24 //  4-pin push-pull mode (inverted outputs)
)

const (
	PRESCALEn = 0
	AUTOSTOPn = 8
	IGNACKn   = 9
	TIMECFGn  = 10
	MATCFGn   = 16
	PINCFGn   = 24
)

const (
	BUSIDLE MCFGR2 = 0xFFF << 0 //+ Bus Idle Timeout
	FILTSCL MCFGR2 = 0x0F << 16 //+ Glitch Filter SCL
	FILTSDA MCFGR2 = 0x0F << 24 //+ Glitch Filter SDA
)

const (
	BUSIDLEn = 0
	FILTSCLn = 16
	FILTSDAn = 24
)

const (
	PINLOW MCFGR3 = 0xFFF << 8 //+ Pin Low Timeout
)

const (
	PINLOWn = 8
)

const (
	MATCH0 MDMR = 0xFF << 0  //+ Match 0 Value
	MATCH1 MDMR = 0xFF << 16 //+ Match 1 Value
)

const (
	MATCH0n = 0
	MATCH1n = 16
)

const (
	CLKLO   MCCR0 = 0x3F << 0  //+ Clock Low Period
	CLKHI   MCCR0 = 0x3F << 8  //+ Clock High Period
	SETHOLD MCCR0 = 0x3F << 16 //+ Setup Hold Delay
	DATAVD  MCCR0 = 0x3F << 24 //+ Data Valid Delay
)

const (
	CLKLOn   = 0
	CLKHIn   = 8
	SETHOLDn = 16
	DATAVDn  = 24
)

const (
	CLKLO   MCCR1 = 0x3F << 0  //+ Clock Low Period
	CLKHI   MCCR1 = 0x3F << 8  //+ Clock High Period
	SETHOLD MCCR1 = 0x3F << 16 //+ Setup Hold Delay
	DATAVD  MCCR1 = 0x3F << 24 //+ Data Valid Delay
)

const (
	CLKLOn   = 0
	CLKHIn   = 8
	SETHOLDn = 16
	DATAVDn  = 24
)

const (
	TXWATER MFCR = 0x03 << 0  //+ Transmit FIFO Watermark
	RXWATER MFCR = 0x03 << 16 //+ Receive FIFO Watermark
)

const (
	TXWATERn = 0
	RXWATERn = 16
)

const (
	TXCOUNT MFSR = 0x07 << 0  //+ Transmit FIFO Count
	RXCOUNT MFSR = 0x07 << 16 //+ Receive FIFO Count
)

const (
	TXCOUNTn = 0
	RXCOUNTn = 16
)

const (
	DATA  MTDR = 0xFF << 0 //+ Transmit Data
	CMD   MTDR = 0x07 << 8 //+ Command Data
	CMD_0 MTDR = 0x00 << 8 //  Transmit DATA[7:0]
	CMD_1 MTDR = 0x01 << 8 //  Receive (DATA[7:0] + 1) bytes
	CMD_2 MTDR = 0x02 << 8 //  Generate STOP condition
	CMD_3 MTDR = 0x03 << 8 //  Receive and discard (DATA[7:0] + 1) bytes
	CMD_4 MTDR = 0x04 << 8 //  Generate (repeated) START and transmit address in DATA[7:0]
	CMD_5 MTDR = 0x05 << 8 //  Generate (repeated) START and transmit address in DATA[7:0]. This transfer expects a NACK to be returned.
	CMD_6 MTDR = 0x06 << 8 //  Generate (repeated) START and transmit address in DATA[7:0] using high speed mode
	CMD_7 MTDR = 0x07 << 8 //  Generate (repeated) START and transmit address in DATA[7:0] using high speed mode. This transfer expects a NACK to be returned.
)

const (
	DATAn = 0
	CMDn  = 8
)

const (
	DATA    MRDR = 0xFF << 0  //+ Receive Data
	RXEMPTY MRDR = 0x01 << 14 //+ RX Empty
)

const (
	DATAn    = 0
	RXEMPTYn = 14
)

const (
	SEN    SCR = 0x01 << 0 //+ Slave Enable
	RST    SCR = 0x01 << 1 //+ Software Reset
	FILTEN SCR = 0x01 << 4 //+ Filter Enable
	FILTDZ SCR = 0x01 << 5 //+ Filter Doze Enable
	RTF    SCR = 0x01 << 8 //+ Reset Transmit FIFO
	RRF    SCR = 0x01 << 9 //+ Reset Receive FIFO
)

const (
	SENn    = 0
	RSTn    = 1
	FILTENn = 4
	FILTDZn = 5
	RTFn    = 8
	RRFn    = 9
)

const (
	TDF  SSR = 0x01 << 0  //+ Transmit Data Flag
	RDF  SSR = 0x01 << 1  //+ Receive Data Flag
	AVF  SSR = 0x01 << 2  //+ Address Valid Flag
	TAF  SSR = 0x01 << 3  //+ Transmit ACK Flag
	RSF  SSR = 0x01 << 8  //+ Repeated Start Flag
	SDF  SSR = 0x01 << 9  //+ STOP Detect Flag
	BEF  SSR = 0x01 << 10 //+ Bit Error Flag
	FEF  SSR = 0x01 << 11 //+ FIFO Error Flag
	AM0F SSR = 0x01 << 12 //+ Address Match 0 Flag
	AM1F SSR = 0x01 << 13 //+ Address Match 1 Flag
	GCF  SSR = 0x01 << 14 //+ General Call Flag
	SARF SSR = 0x01 << 15 //+ SMBus Alert Response Flag
	SBF  SSR = 0x01 << 24 //+ Slave Busy Flag
	BBF  SSR = 0x01 << 25 //+ Bus Busy Flag
)

const (
	TDFn  = 0
	RDFn  = 1
	AVFn  = 2
	TAFn  = 3
	RSFn  = 8
	SDFn  = 9
	BEFn  = 10
	FEFn  = 11
	AM0Fn = 12
	AM1Fn = 13
	GCFn  = 14
	SARFn = 15
	SBFn  = 24
	BBFn  = 25
)

const (
	TDIE  SIER = 0x01 << 0  //+ Transmit Data Interrupt Enable
	RDIE  SIER = 0x01 << 1  //+ Receive Data Interrupt Enable
	AVIE  SIER = 0x01 << 2  //+ Address Valid Interrupt Enable
	TAIE  SIER = 0x01 << 3  //+ Transmit ACK Interrupt Enable
	RSIE  SIER = 0x01 << 8  //+ Repeated Start Interrupt Enable
	SDIE  SIER = 0x01 << 9  //+ STOP Detect Interrupt Enable
	BEIE  SIER = 0x01 << 10 //+ Bit Error Interrupt Enable
	FEIE  SIER = 0x01 << 11 //+ FIFO Error Interrupt Enable
	AM0IE SIER = 0x01 << 12 //+ Address Match 0 Interrupt Enable
	AM1F  SIER = 0x01 << 13 //+ Address Match 1 Interrupt Enable
	GCIE  SIER = 0x01 << 14 //+ General Call Interrupt Enable
	SARIE SIER = 0x01 << 15 //+ SMBus Alert Response Interrupt Enable
)

const (
	TDIEn  = 0
	RDIEn  = 1
	AVIEn  = 2
	TAIEn  = 3
	RSIEn  = 8
	SDIEn  = 9
	BEIEn  = 10
	FEIEn  = 11
	AM0IEn = 12
	AM1Fn  = 13
	GCIEn  = 14
	SARIEn = 15
)

const (
	TDDE SDER = 0x01 << 0 //+ Transmit Data DMA Enable
	RDDE SDER = 0x01 << 1 //+ Receive Data DMA Enable
	AVDE SDER = 0x01 << 2 //+ Address Valid DMA Enable
)

const (
	TDDEn = 0
	RDDEn = 1
	AVDEn = 2
)

const (
	ADRSTALL  SCFGR1 = 0x01 << 0  //+ Address SCL Stall
	RXSTALL   SCFGR1 = 0x01 << 1  //+ RX SCL Stall
	TXDSTALL  SCFGR1 = 0x01 << 2  //+ TX Data SCL Stall
	ACKSTALL  SCFGR1 = 0x01 << 3  //+ ACK SCL Stall
	GCEN      SCFGR1 = 0x01 << 8  //+ General Call Enable
	SAEN      SCFGR1 = 0x01 << 9  //+ SMBus Alert Enable
	TXCFG     SCFGR1 = 0x01 << 10 //+ Transmit Flag Configuration
	RXCFG     SCFGR1 = 0x01 << 11 //+ Receive Data Configuration
	IGNACK    SCFGR1 = 0x01 << 12 //+ Ignore NACK
	HSMEN     SCFGR1 = 0x01 << 13 //+ High Speed Mode Enable
	ADDRCFG   SCFGR1 = 0x07 << 16 //+ Address Configuration
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
	ADRSTALLn = 0
	RXSTALLn  = 1
	TXDSTALLn = 2
	ACKSTALLn = 3
	GCENn     = 8
	SAENn     = 9
	TXCFGn    = 10
	RXCFGn    = 11
	IGNACKn   = 12
	HSMENn    = 13
	ADDRCFGn  = 16
)

const (
	CLKHOLD SCFGR2 = 0x0F << 0  //+ Clock Hold Time
	DATAVD  SCFGR2 = 0x3F << 8  //+ Data Valid Delay
	FILTSCL SCFGR2 = 0x0F << 16 //+ Glitch Filter SCL
	FILTSDA SCFGR2 = 0x0F << 24 //+ Glitch Filter SDA
)

const (
	CLKHOLDn = 0
	DATAVDn  = 8
	FILTSCLn = 16
	FILTSDAn = 24
)

const (
	ADDR0 SAMR = 0x3FF << 1  //+ Address 0 Value
	ADDR1 SAMR = 0x3FF << 17 //+ Address 1 Value
)

const (
	ADDR0n = 1
	ADDR1n = 17
)

const (
	RADDR SASR = 0x7FF << 0 //+ Received Address
	ANV   SASR = 0x01 << 14 //+ Address Not Valid
)

const (
	RADDRn = 0
	ANVn   = 14
)

const (
	TXNACK STAR = 0x01 << 0 //+ Transmit NACK
)

const (
	TXNACKn = 0
)

const (
	DATA STDR = 0xFF << 0 //+ Transmit Data
)

const (
	DATAn = 0
)

const (
	DATA    SRDR = 0xFF << 0  //+ Receive Data
	RXEMPTY SRDR = 0x01 << 14 //+ RX Empty
	SOF     SRDR = 0x01 << 15 //+ Start Of Frame
)

const (
	DATAn    = 0
	RXEMPTYn = 14
	SOFn     = 15
)
