// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package semc provides access to the registers of the SEMC peripheral.
//
// Instances:
//
//	SEMC  SEMC_BASE  -  SEMC*
//
// Registers:
//
//	0x000 32  MCR       Module Control Register
//	0x004 32  IOCR      IO Mux Control Register
//	0x008 32  BMCR0     Master Bus (AXI) Control Register 0
//	0x00C 32  BMCR1     Master Bus (AXI) Control Register 1
//	0x010 32  BR0       Base Register 0 (For SDRAM CS0 device)
//	0x014 32  BR1       Base Register 1 (For SDRAM CS1 device)
//	0x018 32  BR2       Base Register 2 (For SDRAM CS2 device)
//	0x01C 32  BR3       Base Register 3 (For SDRAM CS3 device)
//	0x020 32  BR4       Base Register 4 (For NAND device)
//	0x024 32  BR5       Base Register 5 (For NOR device)
//	0x028 32  BR6       Base Register 6 (For PSRAM device)
//	0x02C 32  BR7       Base Register 7 (For DBI-B (MIPI Display Bus Interface Type B) device)
//	0x030 32  BR8       Base Register 8 (For NAND device)
//	0x034 32  DLLCR     DLL Control Register
//	0x038 32  INTEN     Interrupt Enable Register
//	0x03C 32  INTR      Interrupt Enable Register
//	0x040 32  SDRAMCR0  SDRAM control register 0
//	0x044 32  SDRAMCR1  SDRAM control register 1
//	0x048 32  SDRAMCR2  SDRAM control register 2
//	0x04C 32  SDRAMCR3  SDRAM control register 3
//	0x050 32  NANDCR0   NAND control register 0
//	0x054 32  NANDCR1   NAND control register 1
//	0x058 32  NANDCR2   NAND control register 2
//	0x05C 32  NANDCR3   NAND control register 3
//	0x060 32  NORCR0    NOR control register 0
//	0x064 32  NORCR1    NOR control register 1
//	0x068 32  NORCR2    NOR control register 2
//	0x06C 32  NORCR3    NOR control register 3
//	0x070 32  SRAMCR0   SRAM control register 0
//	0x074 32  SRAMCR1   SRAM control register 1
//	0x078 32  SRAMCR2   SRAM control register 2
//	0x07C 32  SRAMCR3   SRAM control register 3
//	0x080 32  DBICR0    DBI-B control register 0
//	0x084 32  DBICR1    DBI-B control register 1
//	0x090 32  IPCR0     IP Command control register 0
//	0x094 32  IPCR1     IP Command control register 1
//	0x098 32  IPCR2     IP Command control register 2
//	0x09C 32  IPCMD     IP Command register
//	0x0A0 32  IPTXDAT   TX DATA register (for IP Command)
//	0x0B0 32  IPRXDAT   RX DATA register (for IP Command)
//	0x0C0 32  STS0      Status register 0
//	0x0C4 32  STS1      Status register 1
//	0x0C8 32  STS2      Status register 2
//	0x0CC 32  STS3      Status register 3
//	0x0D0 32  STS4      Status register 4
//	0x0D4 32  STS5      Status register 5
//	0x0D8 32  STS6      Status register 6
//	0x0DC 32  STS7      Status register 7
//	0x0E0 32  STS8      Status register 8
//	0x0E4 32  STS9      Status register 9
//	0x0E8 32  STS10     Status register 10
//	0x0EC 32  STS11     Status register 11
//	0x0F0 32  STS12     Status register 12
//	0x0F4 32  STS13     Status register 13
//	0x0F8 32  STS14     Status register 14
//	0x0FC 32  STS15     Status register 15
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package semc

const (
	SWRST  MCR = 0x01 << 0  //+ Software Reset
	MDIS   MCR = 0x01 << 1  //+ Module Disable
	DQSMD  MCR = 0x01 << 2  //+ DQS (read strobe) mode
	WPOL0  MCR = 0x01 << 6  //+ WAIT/RDY# polarity for NOR/PSRAM
	WPOL1  MCR = 0x01 << 7  //+ WAIT/RDY# polarity for NAND
	DQSSEL MCR = 0x01 << 10 //+ Select DQS source when DQSMD and DLLSEL both set.
	DLLSEL MCR = 0x01 << 11 //+ Select DLL delay chain clock input.
	CTO    MCR = 0xFF << 16 //+ Command Execution timeout cycles
	BTO    MCR = 0x1F << 24 //+ Bus timeout cycles
	BTO_0  MCR = 0x00 << 24 //  255*1
	BTO_1  MCR = 0x01 << 24 //  255*2 - 255*2^30
	BTO_2  MCR = 0x02 << 24 //  255*2 - 255*2^30
	BTO_3  MCR = 0x03 << 24 //  255*2 - 255*2^30
	BTO_4  MCR = 0x04 << 24 //  255*2 - 255*2^30
	BTO_5  MCR = 0x05 << 24 //  255*2 - 255*2^30
	BTO_6  MCR = 0x06 << 24 //  255*2 - 255*2^30
	BTO_7  MCR = 0x07 << 24 //  255*2 - 255*2^30
	BTO_8  MCR = 0x08 << 24 //  255*2 - 255*2^30
	BTO_9  MCR = 0x09 << 24 //  255*2 - 255*2^30
	BTO_31 MCR = 0x1F << 24 //  255*2^31
)

const (
	SWRSTn  = 0
	MDISn   = 1
	DQSMDn  = 2
	WPOL0n  = 6
	WPOL1n  = 7
	DQSSELn = 10
	DLLSELn = 11
	CTOn    = 16
	BTOn    = 24
)

const (
	MUX_A8     IOCR = 0x07 << 0  //+ SEMC_A8 output selection
	MUX_A8_0   IOCR = 0x00 << 0  //  SDRAM Address bit (A8)
	MUX_A8_1   IOCR = 0x01 << 0  //  NAND CE#
	MUX_A8_2   IOCR = 0x02 << 0  //  NOR CE#
	MUX_A8_3   IOCR = 0x03 << 0  //  PSRAM CE#
	MUX_A8_4   IOCR = 0x04 << 0  //  DBI CSX
	MUX_A8_5   IOCR = 0x05 << 0  //  SDRAM Address bit (A8)
	MUX_A8_6   IOCR = 0x06 << 0  //  SDRAM Address bit (A8)
	MUX_A8_7   IOCR = 0x07 << 0  //  SDRAM Address bit (A8)
	MUX_CSX0   IOCR = 0x07 << 3  //+ SEMC_CSX0 output selection
	MUX_CSX0_0 IOCR = 0x00 << 3  //  NOR/PSRAM Address bit 24 (A24)
	MUX_CSX0_1 IOCR = 0x01 << 3  //  SDRAM CS1
	MUX_CSX0_2 IOCR = 0x02 << 3  //  SDRAM CS2
	MUX_CSX0_3 IOCR = 0x03 << 3  //  SDRAM CS3
	MUX_CSX0_4 IOCR = 0x04 << 3  //  NAND CE#
	MUX_CSX0_5 IOCR = 0x05 << 3  //  NOR CE#
	MUX_CSX0_6 IOCR = 0x06 << 3  //  PSRAM CE#
	MUX_CSX0_7 IOCR = 0x07 << 3  //  DBI CSX
	MUX_CSX1   IOCR = 0x07 << 6  //+ SEMC_CSX1 output selection
	MUX_CSX1_0 IOCR = 0x00 << 6  //  NOR/PSRAM Address bit 25 (A25)
	MUX_CSX1_1 IOCR = 0x01 << 6  //  SDRAM CS1
	MUX_CSX1_2 IOCR = 0x02 << 6  //  SDRAM CS2
	MUX_CSX1_3 IOCR = 0x03 << 6  //  SDRAM CS3
	MUX_CSX1_4 IOCR = 0x04 << 6  //  NAND CE#
	MUX_CSX1_5 IOCR = 0x05 << 6  //  NOR CE#
	MUX_CSX1_6 IOCR = 0x06 << 6  //  PSRAM CE#
	MUX_CSX1_7 IOCR = 0x07 << 6  //  DBI CSX
	MUX_CSX2   IOCR = 0x07 << 9  //+ SEMC_CSX2 output selection
	MUX_CSX2_0 IOCR = 0x00 << 9  //  NOR/PSRAM Address bit 26 (A26)
	MUX_CSX2_1 IOCR = 0x01 << 9  //  SDRAM CS1
	MUX_CSX2_2 IOCR = 0x02 << 9  //  SDRAM CS2
	MUX_CSX2_3 IOCR = 0x03 << 9  //  SDRAM CS3
	MUX_CSX2_4 IOCR = 0x04 << 9  //  NAND CE#
	MUX_CSX2_5 IOCR = 0x05 << 9  //  NOR CE#
	MUX_CSX2_6 IOCR = 0x06 << 9  //  PSRAM CE#
	MUX_CSX2_7 IOCR = 0x07 << 9  //  DBI CSX
	MUX_CSX3   IOCR = 0x07 << 12 //+ SEMC_CSX3 output selection
	MUX_CSX3_0 IOCR = 0x00 << 12 //  NOR/PSRAM Address bit 27 (A27)
	MUX_CSX3_1 IOCR = 0x01 << 12 //  SDRAM CS1
	MUX_CSX3_2 IOCR = 0x02 << 12 //  SDRAM CS2
	MUX_CSX3_3 IOCR = 0x03 << 12 //  SDRAM CS3
	MUX_CSX3_4 IOCR = 0x04 << 12 //  NAND CE#
	MUX_CSX3_5 IOCR = 0x05 << 12 //  NOR CE#
	MUX_CSX3_6 IOCR = 0x06 << 12 //  PSRAM CE#
	MUX_CSX3_7 IOCR = 0x07 << 12 //  DBI CSX
	MUX_RDY    IOCR = 0x07 << 15 //+ SEMC_RDY function selection
	MUX_RDY_0  IOCR = 0x00 << 15 //  NAND Ready/Wait# input
	MUX_RDY_1  IOCR = 0x01 << 15 //  SDRAM CS1
	MUX_RDY_2  IOCR = 0x02 << 15 //  SDRAM CS2
	MUX_RDY_3  IOCR = 0x03 << 15 //  SDRAM CS3
	MUX_RDY_4  IOCR = 0x04 << 15 //  NOR CE#
	MUX_RDY_5  IOCR = 0x05 << 15 //  PSRAM CE#
	MUX_RDY_6  IOCR = 0x06 << 15 //  DBI CSX
	MUX_RDY_7  IOCR = 0x07 << 15 //  NOR/PSRAM Address bit 27
	MUX_CLKX0  IOCR = 0x01 << 24 //+ SEMC_CLKX0 function selection
	MUX_CLKX1  IOCR = 0x01 << 25 //+ SEMC_CLKX1 function selection
)

const (
	MUX_A8n    = 0
	MUX_CSX0n  = 3
	MUX_CSX1n  = 6
	MUX_CSX2n  = 9
	MUX_CSX3n  = 12
	MUX_RDYn   = 15
	MUX_CLKX0n = 24
	MUX_CLKX1n = 25
)

const (
	WQOS BMCR0 = 0x0F << 0  //+ Weight of QoS
	WAGE BMCR0 = 0x0F << 4  //+ Weight of Aging
	WSH  BMCR0 = 0xFF << 8  //+ Weight of Slave Hit (no read/write switch)
	WRWS BMCR0 = 0xFF << 16 //+ Weight of Slave Hit (Read/Write switch)
)

const (
	WQOSn = 0
	WAGEn = 4
	WSHn  = 8
	WRWSn = 16
)

const (
	WQOS BMCR1 = 0x0F << 0  //+ Weight of QoS
	WAGE BMCR1 = 0x0F << 4  //+ Weight of Aging
	WPH  BMCR1 = 0xFF << 8  //+ Weight of Page Hit
	WRWS BMCR1 = 0xFF << 16 //+ Weight of Read/Write switch
	WBR  BMCR1 = 0xFF << 24 //+ Weight of Bank Rotation
)

const (
	WQOSn = 0
	WAGEn = 4
	WPHn  = 8
	WRWSn = 16
	WBRn  = 24
)

const (
	VLD   BR0 = 0x01 << 0     //+ Valid
	MS    BR0 = 0x1F << 1     //+ Memory size
	MS_0  BR0 = 0x00 << 1     //  4KB
	MS_1  BR0 = 0x01 << 1     //  8KB
	MS_2  BR0 = 0x02 << 1     //  16KB
	MS_3  BR0 = 0x03 << 1     //  32KB
	MS_4  BR0 = 0x04 << 1     //  64KB
	MS_5  BR0 = 0x05 << 1     //  128KB
	MS_6  BR0 = 0x06 << 1     //  256KB
	MS_7  BR0 = 0x07 << 1     //  512KB
	MS_8  BR0 = 0x08 << 1     //  1MB
	MS_9  BR0 = 0x09 << 1     //  2MB
	MS_10 BR0 = 0x0A << 1     //  4MB
	MS_11 BR0 = 0x0B << 1     //  8MB
	MS_12 BR0 = 0x0C << 1     //  16MB
	MS_13 BR0 = 0x0D << 1     //  32MB
	MS_14 BR0 = 0x0E << 1     //  64MB
	MS_15 BR0 = 0x0F << 1     //  128MB
	MS_16 BR0 = 0x10 << 1     //  256MB
	MS_17 BR0 = 0x11 << 1     //  512MB
	MS_18 BR0 = 0x12 << 1     //  1GB
	MS_19 BR0 = 0x13 << 1     //  2GB
	MS_20 BR0 = 0x14 << 1     //  4GB
	MS_21 BR0 = 0x15 << 1     //  4GB
	MS_22 BR0 = 0x16 << 1     //  4GB
	MS_23 BR0 = 0x17 << 1     //  4GB
	MS_24 BR0 = 0x18 << 1     //  4GB
	MS_25 BR0 = 0x19 << 1     //  4GB
	MS_26 BR0 = 0x1A << 1     //  4GB
	MS_27 BR0 = 0x1B << 1     //  4GB
	MS_28 BR0 = 0x1C << 1     //  4GB
	MS_29 BR0 = 0x1D << 1     //  4GB
	MS_30 BR0 = 0x1E << 1     //  4GB
	MS_31 BR0 = 0x1F << 1     //  4GB
	BA    BR0 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR1 = 0x01 << 0     //+ Valid
	MS    BR1 = 0x1F << 1     //+ Memory size
	MS_0  BR1 = 0x00 << 1     //  4KB
	MS_1  BR1 = 0x01 << 1     //  8KB
	MS_2  BR1 = 0x02 << 1     //  16KB
	MS_3  BR1 = 0x03 << 1     //  32KB
	MS_4  BR1 = 0x04 << 1     //  64KB
	MS_5  BR1 = 0x05 << 1     //  128KB
	MS_6  BR1 = 0x06 << 1     //  256KB
	MS_7  BR1 = 0x07 << 1     //  512KB
	MS_8  BR1 = 0x08 << 1     //  1MB
	MS_9  BR1 = 0x09 << 1     //  2MB
	MS_10 BR1 = 0x0A << 1     //  4MB
	MS_11 BR1 = 0x0B << 1     //  8MB
	MS_12 BR1 = 0x0C << 1     //  16MB
	MS_13 BR1 = 0x0D << 1     //  32MB
	MS_14 BR1 = 0x0E << 1     //  64MB
	MS_15 BR1 = 0x0F << 1     //  128MB
	MS_16 BR1 = 0x10 << 1     //  256MB
	MS_17 BR1 = 0x11 << 1     //  512MB
	MS_18 BR1 = 0x12 << 1     //  1GB
	MS_19 BR1 = 0x13 << 1     //  2GB
	MS_20 BR1 = 0x14 << 1     //  4GB
	MS_21 BR1 = 0x15 << 1     //  4GB
	MS_22 BR1 = 0x16 << 1     //  4GB
	MS_23 BR1 = 0x17 << 1     //  4GB
	MS_24 BR1 = 0x18 << 1     //  4GB
	MS_25 BR1 = 0x19 << 1     //  4GB
	MS_26 BR1 = 0x1A << 1     //  4GB
	MS_27 BR1 = 0x1B << 1     //  4GB
	MS_28 BR1 = 0x1C << 1     //  4GB
	MS_29 BR1 = 0x1D << 1     //  4GB
	MS_30 BR1 = 0x1E << 1     //  4GB
	MS_31 BR1 = 0x1F << 1     //  4GB
	BA    BR1 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR2 = 0x01 << 0     //+ Valid
	MS    BR2 = 0x1F << 1     //+ Memory size
	MS_0  BR2 = 0x00 << 1     //  4KB
	MS_1  BR2 = 0x01 << 1     //  8KB
	MS_2  BR2 = 0x02 << 1     //  16KB
	MS_3  BR2 = 0x03 << 1     //  32KB
	MS_4  BR2 = 0x04 << 1     //  64KB
	MS_5  BR2 = 0x05 << 1     //  128KB
	MS_6  BR2 = 0x06 << 1     //  256KB
	MS_7  BR2 = 0x07 << 1     //  512KB
	MS_8  BR2 = 0x08 << 1     //  1MB
	MS_9  BR2 = 0x09 << 1     //  2MB
	MS_10 BR2 = 0x0A << 1     //  4MB
	MS_11 BR2 = 0x0B << 1     //  8MB
	MS_12 BR2 = 0x0C << 1     //  16MB
	MS_13 BR2 = 0x0D << 1     //  32MB
	MS_14 BR2 = 0x0E << 1     //  64MB
	MS_15 BR2 = 0x0F << 1     //  128MB
	MS_16 BR2 = 0x10 << 1     //  256MB
	MS_17 BR2 = 0x11 << 1     //  512MB
	MS_18 BR2 = 0x12 << 1     //  1GB
	MS_19 BR2 = 0x13 << 1     //  2GB
	MS_20 BR2 = 0x14 << 1     //  4GB
	MS_21 BR2 = 0x15 << 1     //  4GB
	MS_22 BR2 = 0x16 << 1     //  4GB
	MS_23 BR2 = 0x17 << 1     //  4GB
	MS_24 BR2 = 0x18 << 1     //  4GB
	MS_25 BR2 = 0x19 << 1     //  4GB
	MS_26 BR2 = 0x1A << 1     //  4GB
	MS_27 BR2 = 0x1B << 1     //  4GB
	MS_28 BR2 = 0x1C << 1     //  4GB
	MS_29 BR2 = 0x1D << 1     //  4GB
	MS_30 BR2 = 0x1E << 1     //  4GB
	MS_31 BR2 = 0x1F << 1     //  4GB
	BA    BR2 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR3 = 0x01 << 0     //+ Valid
	MS    BR3 = 0x1F << 1     //+ Memory size
	MS_0  BR3 = 0x00 << 1     //  4KB
	MS_1  BR3 = 0x01 << 1     //  8KB
	MS_2  BR3 = 0x02 << 1     //  16KB
	MS_3  BR3 = 0x03 << 1     //  32KB
	MS_4  BR3 = 0x04 << 1     //  64KB
	MS_5  BR3 = 0x05 << 1     //  128KB
	MS_6  BR3 = 0x06 << 1     //  256KB
	MS_7  BR3 = 0x07 << 1     //  512KB
	MS_8  BR3 = 0x08 << 1     //  1MB
	MS_9  BR3 = 0x09 << 1     //  2MB
	MS_10 BR3 = 0x0A << 1     //  4MB
	MS_11 BR3 = 0x0B << 1     //  8MB
	MS_12 BR3 = 0x0C << 1     //  16MB
	MS_13 BR3 = 0x0D << 1     //  32MB
	MS_14 BR3 = 0x0E << 1     //  64MB
	MS_15 BR3 = 0x0F << 1     //  128MB
	MS_16 BR3 = 0x10 << 1     //  256MB
	MS_17 BR3 = 0x11 << 1     //  512MB
	MS_18 BR3 = 0x12 << 1     //  1GB
	MS_19 BR3 = 0x13 << 1     //  2GB
	MS_20 BR3 = 0x14 << 1     //  4GB
	MS_21 BR3 = 0x15 << 1     //  4GB
	MS_22 BR3 = 0x16 << 1     //  4GB
	MS_23 BR3 = 0x17 << 1     //  4GB
	MS_24 BR3 = 0x18 << 1     //  4GB
	MS_25 BR3 = 0x19 << 1     //  4GB
	MS_26 BR3 = 0x1A << 1     //  4GB
	MS_27 BR3 = 0x1B << 1     //  4GB
	MS_28 BR3 = 0x1C << 1     //  4GB
	MS_29 BR3 = 0x1D << 1     //  4GB
	MS_30 BR3 = 0x1E << 1     //  4GB
	MS_31 BR3 = 0x1F << 1     //  4GB
	BA    BR3 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR4 = 0x01 << 0     //+ Valid
	MS    BR4 = 0x1F << 1     //+ Memory size
	MS_0  BR4 = 0x00 << 1     //  4KB
	MS_1  BR4 = 0x01 << 1     //  8KB
	MS_2  BR4 = 0x02 << 1     //  16KB
	MS_3  BR4 = 0x03 << 1     //  32KB
	MS_4  BR4 = 0x04 << 1     //  64KB
	MS_5  BR4 = 0x05 << 1     //  128KB
	MS_6  BR4 = 0x06 << 1     //  256KB
	MS_7  BR4 = 0x07 << 1     //  512KB
	MS_8  BR4 = 0x08 << 1     //  1MB
	MS_9  BR4 = 0x09 << 1     //  2MB
	MS_10 BR4 = 0x0A << 1     //  4MB
	MS_11 BR4 = 0x0B << 1     //  8MB
	MS_12 BR4 = 0x0C << 1     //  16MB
	MS_13 BR4 = 0x0D << 1     //  32MB
	MS_14 BR4 = 0x0E << 1     //  64MB
	MS_15 BR4 = 0x0F << 1     //  128MB
	MS_16 BR4 = 0x10 << 1     //  256MB
	MS_17 BR4 = 0x11 << 1     //  512MB
	MS_18 BR4 = 0x12 << 1     //  1GB
	MS_19 BR4 = 0x13 << 1     //  2GB
	MS_20 BR4 = 0x14 << 1     //  4GB
	MS_21 BR4 = 0x15 << 1     //  4GB
	MS_22 BR4 = 0x16 << 1     //  4GB
	MS_23 BR4 = 0x17 << 1     //  4GB
	MS_24 BR4 = 0x18 << 1     //  4GB
	MS_25 BR4 = 0x19 << 1     //  4GB
	MS_26 BR4 = 0x1A << 1     //  4GB
	MS_27 BR4 = 0x1B << 1     //  4GB
	MS_28 BR4 = 0x1C << 1     //  4GB
	MS_29 BR4 = 0x1D << 1     //  4GB
	MS_30 BR4 = 0x1E << 1     //  4GB
	MS_31 BR4 = 0x1F << 1     //  4GB
	BA    BR4 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR5 = 0x01 << 0     //+ Valid
	MS    BR5 = 0x1F << 1     //+ Memory size
	MS_0  BR5 = 0x00 << 1     //  4KB
	MS_1  BR5 = 0x01 << 1     //  8KB
	MS_2  BR5 = 0x02 << 1     //  16KB
	MS_3  BR5 = 0x03 << 1     //  32KB
	MS_4  BR5 = 0x04 << 1     //  64KB
	MS_5  BR5 = 0x05 << 1     //  128KB
	MS_6  BR5 = 0x06 << 1     //  256KB
	MS_7  BR5 = 0x07 << 1     //  512KB
	MS_8  BR5 = 0x08 << 1     //  1MB
	MS_9  BR5 = 0x09 << 1     //  2MB
	MS_10 BR5 = 0x0A << 1     //  4MB
	MS_11 BR5 = 0x0B << 1     //  8MB
	MS_12 BR5 = 0x0C << 1     //  16MB
	MS_13 BR5 = 0x0D << 1     //  32MB
	MS_14 BR5 = 0x0E << 1     //  64MB
	MS_15 BR5 = 0x0F << 1     //  128MB
	MS_16 BR5 = 0x10 << 1     //  256MB
	MS_17 BR5 = 0x11 << 1     //  512MB
	MS_18 BR5 = 0x12 << 1     //  1GB
	MS_19 BR5 = 0x13 << 1     //  2GB
	MS_20 BR5 = 0x14 << 1     //  4GB
	MS_21 BR5 = 0x15 << 1     //  4GB
	MS_22 BR5 = 0x16 << 1     //  4GB
	MS_23 BR5 = 0x17 << 1     //  4GB
	MS_24 BR5 = 0x18 << 1     //  4GB
	MS_25 BR5 = 0x19 << 1     //  4GB
	MS_26 BR5 = 0x1A << 1     //  4GB
	MS_27 BR5 = 0x1B << 1     //  4GB
	MS_28 BR5 = 0x1C << 1     //  4GB
	MS_29 BR5 = 0x1D << 1     //  4GB
	MS_30 BR5 = 0x1E << 1     //  4GB
	MS_31 BR5 = 0x1F << 1     //  4GB
	BA    BR5 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR6 = 0x01 << 0     //+ Valid
	MS    BR6 = 0x1F << 1     //+ Memory size
	MS_0  BR6 = 0x00 << 1     //  4KB
	MS_1  BR6 = 0x01 << 1     //  8KB
	MS_2  BR6 = 0x02 << 1     //  16KB
	MS_3  BR6 = 0x03 << 1     //  32KB
	MS_4  BR6 = 0x04 << 1     //  64KB
	MS_5  BR6 = 0x05 << 1     //  128KB
	MS_6  BR6 = 0x06 << 1     //  256KB
	MS_7  BR6 = 0x07 << 1     //  512KB
	MS_8  BR6 = 0x08 << 1     //  1MB
	MS_9  BR6 = 0x09 << 1     //  2MB
	MS_10 BR6 = 0x0A << 1     //  4MB
	MS_11 BR6 = 0x0B << 1     //  8MB
	MS_12 BR6 = 0x0C << 1     //  16MB
	MS_13 BR6 = 0x0D << 1     //  32MB
	MS_14 BR6 = 0x0E << 1     //  64MB
	MS_15 BR6 = 0x0F << 1     //  128MB
	MS_16 BR6 = 0x10 << 1     //  256MB
	MS_17 BR6 = 0x11 << 1     //  512MB
	MS_18 BR6 = 0x12 << 1     //  1GB
	MS_19 BR6 = 0x13 << 1     //  2GB
	MS_20 BR6 = 0x14 << 1     //  4GB
	MS_21 BR6 = 0x15 << 1     //  4GB
	MS_22 BR6 = 0x16 << 1     //  4GB
	MS_23 BR6 = 0x17 << 1     //  4GB
	MS_24 BR6 = 0x18 << 1     //  4GB
	MS_25 BR6 = 0x19 << 1     //  4GB
	MS_26 BR6 = 0x1A << 1     //  4GB
	MS_27 BR6 = 0x1B << 1     //  4GB
	MS_28 BR6 = 0x1C << 1     //  4GB
	MS_29 BR6 = 0x1D << 1     //  4GB
	MS_30 BR6 = 0x1E << 1     //  4GB
	MS_31 BR6 = 0x1F << 1     //  4GB
	BA    BR6 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR7 = 0x01 << 0     //+ Valid
	MS    BR7 = 0x1F << 1     //+ Memory size
	MS_0  BR7 = 0x00 << 1     //  4KB
	MS_1  BR7 = 0x01 << 1     //  8KB
	MS_2  BR7 = 0x02 << 1     //  16KB
	MS_3  BR7 = 0x03 << 1     //  32KB
	MS_4  BR7 = 0x04 << 1     //  64KB
	MS_5  BR7 = 0x05 << 1     //  128KB
	MS_6  BR7 = 0x06 << 1     //  256KB
	MS_7  BR7 = 0x07 << 1     //  512KB
	MS_8  BR7 = 0x08 << 1     //  1MB
	MS_9  BR7 = 0x09 << 1     //  2MB
	MS_10 BR7 = 0x0A << 1     //  4MB
	MS_11 BR7 = 0x0B << 1     //  8MB
	MS_12 BR7 = 0x0C << 1     //  16MB
	MS_13 BR7 = 0x0D << 1     //  32MB
	MS_14 BR7 = 0x0E << 1     //  64MB
	MS_15 BR7 = 0x0F << 1     //  128MB
	MS_16 BR7 = 0x10 << 1     //  256MB
	MS_17 BR7 = 0x11 << 1     //  512MB
	MS_18 BR7 = 0x12 << 1     //  1GB
	MS_19 BR7 = 0x13 << 1     //  2GB
	MS_20 BR7 = 0x14 << 1     //  4GB
	MS_21 BR7 = 0x15 << 1     //  4GB
	MS_22 BR7 = 0x16 << 1     //  4GB
	MS_23 BR7 = 0x17 << 1     //  4GB
	MS_24 BR7 = 0x18 << 1     //  4GB
	MS_25 BR7 = 0x19 << 1     //  4GB
	MS_26 BR7 = 0x1A << 1     //  4GB
	MS_27 BR7 = 0x1B << 1     //  4GB
	MS_28 BR7 = 0x1C << 1     //  4GB
	MS_29 BR7 = 0x1D << 1     //  4GB
	MS_30 BR7 = 0x1E << 1     //  4GB
	MS_31 BR7 = 0x1F << 1     //  4GB
	BA    BR7 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	VLD   BR8 = 0x01 << 0     //+ Valid
	MS    BR8 = 0x1F << 1     //+ Memory size
	MS_0  BR8 = 0x00 << 1     //  4KB
	MS_1  BR8 = 0x01 << 1     //  8KB
	MS_2  BR8 = 0x02 << 1     //  16KB
	MS_3  BR8 = 0x03 << 1     //  32KB
	MS_4  BR8 = 0x04 << 1     //  64KB
	MS_5  BR8 = 0x05 << 1     //  128KB
	MS_6  BR8 = 0x06 << 1     //  256KB
	MS_7  BR8 = 0x07 << 1     //  512KB
	MS_8  BR8 = 0x08 << 1     //  1MB
	MS_9  BR8 = 0x09 << 1     //  2MB
	MS_10 BR8 = 0x0A << 1     //  4MB
	MS_11 BR8 = 0x0B << 1     //  8MB
	MS_12 BR8 = 0x0C << 1     //  16MB
	MS_13 BR8 = 0x0D << 1     //  32MB
	MS_14 BR8 = 0x0E << 1     //  64MB
	MS_15 BR8 = 0x0F << 1     //  128MB
	MS_16 BR8 = 0x10 << 1     //  256MB
	MS_17 BR8 = 0x11 << 1     //  512MB
	MS_18 BR8 = 0x12 << 1     //  1GB
	MS_19 BR8 = 0x13 << 1     //  2GB
	MS_20 BR8 = 0x14 << 1     //  4GB
	MS_21 BR8 = 0x15 << 1     //  4GB
	MS_22 BR8 = 0x16 << 1     //  4GB
	MS_23 BR8 = 0x17 << 1     //  4GB
	MS_24 BR8 = 0x18 << 1     //  4GB
	MS_25 BR8 = 0x19 << 1     //  4GB
	MS_26 BR8 = 0x1A << 1     //  4GB
	MS_27 BR8 = 0x1B << 1     //  4GB
	MS_28 BR8 = 0x1C << 1     //  4GB
	MS_29 BR8 = 0x1D << 1     //  4GB
	MS_30 BR8 = 0x1E << 1     //  4GB
	MS_31 BR8 = 0x1F << 1     //  4GB
	BA    BR8 = 0xFFFFF << 12 //+ Base Address
)

const (
	VLDn = 0
	MSn  = 1
	BAn  = 12
)

const (
	DLLEN        DLLCR = 0x01 << 0 //+ DLL calibration enable.
	DLLRESET     DLLCR = 0x01 << 1 //+ Software could force a reset on DLL by setting this field to 0x1. This will cause the DLL to lose lock and re-calibrate to detect an ref_clock half period phase shift. The reset action is edge triggered, so software need to clear this bit after set this bit (no delay limitation).
	SLVDLYTARGET DLLCR = 0x0F << 3 //+ The delay target for slave delay line is: ((SLVDLYTARGET+1) * 1/32 * clock cycle of reference clock (ipgclock).
	OVRDEN       DLLCR = 0x01 << 8 //+ Slave clock delay line delay cell number selection override enable.
	OVRDVAL      DLLCR = 0x3F << 9 //+ Slave clock delay line delay cell number selection override value.
)

const (
	DLLENn        = 0
	DLLRESETn     = 1
	SLVDLYTARGETn = 3
	OVRDENn       = 8
	OVRDVALn      = 9
)

const (
	IPCMDDONEEN INTEN = 0x01 << 0 //+ IP command done interrupt enable
	IPCMDERREN  INTEN = 0x01 << 1 //+ IP command error interrupt enable
	AXICMDERREN INTEN = 0x01 << 2 //+ AXI command error interrupt enable
	AXIBUSERREN INTEN = 0x01 << 3 //+ AXI bus error interrupt enable
	NDPAGEENDEN INTEN = 0x01 << 4 //+ This bit enable/disable the NDPAGEEND interrupt generation.
	NDNOPENDEN  INTEN = 0x01 << 5 //+ This bit enable/disable the NDNOPEND interrupt generation.
)

const (
	IPCMDDONEENn = 0
	IPCMDERRENn  = 1
	AXICMDERRENn = 2
	AXIBUSERRENn = 3
	NDPAGEENDENn = 4
	NDNOPENDENn  = 5
)

const (
	IPCMDDONE INTR = 0x01 << 0 //+ IP command normal done interrupt
	IPCMDERR  INTR = 0x01 << 1 //+ IP command error done interrupt
	AXICMDERR INTR = 0x01 << 2 //+ AXI command error interrupt
	AXIBUSERR INTR = 0x01 << 3 //+ AXI bus error interrupt
	NDPAGEEND INTR = 0x01 << 4 //+ This interrupt is generated when the last address of one page in NAND device is written by AXI command
	NDNOPEND  INTR = 0x01 << 5 //+ This interrupt is generated when all pending AXI write command to NAND is finished on NAND interface.
)

const (
	IPCMDDONEn = 0
	IPCMDERRn  = 1
	AXICMDERRn = 2
	AXIBUSERRn = 3
	NDPAGEENDn = 4
	NDNOPENDn  = 5
)

const (
	PS    SDRAMCR0 = 0x01 << 0  //+ Port Size
	BL    SDRAMCR0 = 0x07 << 4  //+ Burst Length
	BL_0  SDRAMCR0 = 0x00 << 4  //  1
	BL_1  SDRAMCR0 = 0x01 << 4  //  2
	BL_2  SDRAMCR0 = 0x02 << 4  //  4
	BL_3  SDRAMCR0 = 0x03 << 4  //  8
	BL_4  SDRAMCR0 = 0x04 << 4  //  8
	BL_5  SDRAMCR0 = 0x05 << 4  //  8
	BL_6  SDRAMCR0 = 0x06 << 4  //  8
	BL_7  SDRAMCR0 = 0x07 << 4  //  8
	COL8  SDRAMCR0 = 0x01 << 7  //+ Column 8 selection bit
	COL   SDRAMCR0 = 0x03 << 8  //+ Column address bit number
	COL_0 SDRAMCR0 = 0x00 << 8  //  12 bit
	COL_1 SDRAMCR0 = 0x01 << 8  //  11 bit
	COL_2 SDRAMCR0 = 0x02 << 8  //  10 bit
	COL_3 SDRAMCR0 = 0x03 << 8  //  9 bit
	CL    SDRAMCR0 = 0x03 << 10 //+ CAS Latency
	CL_0  SDRAMCR0 = 0x00 << 10 //  1
	CL_1  SDRAMCR0 = 0x01 << 10 //  1
	CL_2  SDRAMCR0 = 0x02 << 10 //  2
	CL_3  SDRAMCR0 = 0x03 << 10 //  3
	BANK2 SDRAMCR0 = 0x01 << 14 //+ 2 Bank selection bit
)

const (
	PSn    = 0
	BLn    = 4
	COL8n  = 7
	COLn   = 8
	CLn    = 10
	BANK2n = 14
)

const (
	PRE2ACT SDRAMCR1 = 0x0F << 0  //+ PRECHARGE to ACT/Refresh wait time
	ACT2RW  SDRAMCR1 = 0x0F << 4  //+ ACT to Read/Write wait time
	RFRC    SDRAMCR1 = 0x1F << 8  //+ Refresh recovery time
	WRC     SDRAMCR1 = 0x07 << 13 //+ Write recovery time
	CKEOFF  SDRAMCR1 = 0x0F << 16 //+ CKE OFF minimum time
	ACT2PRE SDRAMCR1 = 0x0F << 20 //+ ACT to Precharge minimum time
)

const (
	PRE2ACTn = 0
	ACT2RWn  = 4
	RFRCn    = 8
	WRCn     = 13
	CKEOFFn  = 16
	ACT2PREn = 20
)

const (
	SRRC    SDRAMCR2 = 0xFF << 0  //+ Self Refresh Recovery time
	REF2REF SDRAMCR2 = 0xFF << 8  //+ Refresh to Refresh wait time
	ACT2ACT SDRAMCR2 = 0xFF << 16 //+ ACT to ACT wait time
	ITO     SDRAMCR2 = 0xFF << 24 //+ SDRAM Idle timeout
	ITO_0   SDRAMCR2 = 0x00 << 24 //  IDLE timeout period is 256*Prescale period.
	ITO_1   SDRAMCR2 = 0x01 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_2   SDRAMCR2 = 0x02 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_3   SDRAMCR2 = 0x03 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_4   SDRAMCR2 = 0x04 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_5   SDRAMCR2 = 0x05 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_6   SDRAMCR2 = 0x06 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_7   SDRAMCR2 = 0x07 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_8   SDRAMCR2 = 0x08 << 24 //  IDLE timeout period is ITO*Prescale period.
	ITO_9   SDRAMCR2 = 0x09 << 24 //  IDLE timeout period is ITO*Prescale period.
)

const (
	SRRCn    = 0
	REF2REFn = 8
	ACT2ACTn = 16
	ITOn     = 24
)

const (
	REN        SDRAMCR3 = 0x01 << 0  //+ Refresh enable
	REBL       SDRAMCR3 = 0x07 << 1  //+ Refresh burst length
	REBL_0     SDRAMCR3 = 0x00 << 1  //  1
	REBL_1     SDRAMCR3 = 0x01 << 1  //  2
	REBL_2     SDRAMCR3 = 0x02 << 1  //  3
	REBL_3     SDRAMCR3 = 0x03 << 1  //  4
	REBL_4     SDRAMCR3 = 0x04 << 1  //  5
	REBL_5     SDRAMCR3 = 0x05 << 1  //  6
	REBL_6     SDRAMCR3 = 0x06 << 1  //  7
	REBL_7     SDRAMCR3 = 0x07 << 1  //  8
	PRESCALE   SDRAMCR3 = 0xFF << 8  //+ Prescaler timer period
	PRESCALE_0 SDRAMCR3 = 0x00 << 8  //  256*16 cycle
	PRESCALE_1 SDRAMCR3 = 0x01 << 8  //  PRESCALE*16 cycle
	PRESCALE_2 SDRAMCR3 = 0x02 << 8  //  PRESCALE*16 cycle
	PRESCALE_3 SDRAMCR3 = 0x03 << 8  //  PRESCALE*16 cycle
	PRESCALE_4 SDRAMCR3 = 0x04 << 8  //  PRESCALE*16 cycle
	PRESCALE_5 SDRAMCR3 = 0x05 << 8  //  PRESCALE*16 cycle
	PRESCALE_6 SDRAMCR3 = 0x06 << 8  //  PRESCALE*16 cycle
	PRESCALE_7 SDRAMCR3 = 0x07 << 8  //  PRESCALE*16 cycle
	PRESCALE_8 SDRAMCR3 = 0x08 << 8  //  PRESCALE*16 cycle
	PRESCALE_9 SDRAMCR3 = 0x09 << 8  //  PRESCALE*16 cycle
	RT         SDRAMCR3 = 0xFF << 16 //+ Refresh timer period
	RT_0       SDRAMCR3 = 0x00 << 16 //  256*Prescaler period
	RT_1       SDRAMCR3 = 0x01 << 16 //  RT*Prescaler period
	RT_2       SDRAMCR3 = 0x02 << 16 //  RT*Prescaler period
	RT_3       SDRAMCR3 = 0x03 << 16 //  RT*Prescaler period
	RT_4       SDRAMCR3 = 0x04 << 16 //  RT*Prescaler period
	RT_5       SDRAMCR3 = 0x05 << 16 //  RT*Prescaler period
	RT_6       SDRAMCR3 = 0x06 << 16 //  RT*Prescaler period
	RT_7       SDRAMCR3 = 0x07 << 16 //  RT*Prescaler period
	RT_8       SDRAMCR3 = 0x08 << 16 //  RT*Prescaler period
	RT_9       SDRAMCR3 = 0x09 << 16 //  RT*Prescaler period
	UT         SDRAMCR3 = 0xFF << 24 //+ Refresh urgent threshold
	UT_0       SDRAMCR3 = 0x00 << 24 //  256*Prescaler period
	UT_1       SDRAMCR3 = 0x01 << 24 //  UT*Prescaler period
	UT_2       SDRAMCR3 = 0x02 << 24 //  UT*Prescaler period
	UT_3       SDRAMCR3 = 0x03 << 24 //  UT*Prescaler period
	UT_4       SDRAMCR3 = 0x04 << 24 //  UT*Prescaler period
	UT_5       SDRAMCR3 = 0x05 << 24 //  UT*Prescaler period
	UT_6       SDRAMCR3 = 0x06 << 24 //  UT*Prescaler period
	UT_7       SDRAMCR3 = 0x07 << 24 //  UT*Prescaler period
	UT_8       SDRAMCR3 = 0x08 << 24 //  UT*Prescaler period
	UT_9       SDRAMCR3 = 0x09 << 24 //  UT*Prescaler period
)

const (
	RENn      = 0
	REBLn     = 1
	PRESCALEn = 8
	RTn       = 16
	UTn       = 24
)

const (
	PS     NANDCR0 = 0x01 << 0 //+ Port Size
	SYNCEN NANDCR0 = 0x01 << 1 //+ Select NAND controller mode.
	BL     NANDCR0 = 0x07 << 4 //+ Burst Length
	BL_0   NANDCR0 = 0x00 << 4 //  1
	BL_1   NANDCR0 = 0x01 << 4 //  2
	BL_2   NANDCR0 = 0x02 << 4 //  4
	BL_3   NANDCR0 = 0x03 << 4 //  8
	BL_4   NANDCR0 = 0x04 << 4 //  16
	BL_5   NANDCR0 = 0x05 << 4 //  32
	BL_6   NANDCR0 = 0x06 << 4 //  64
	BL_7   NANDCR0 = 0x07 << 4 //  64
	EDO    NANDCR0 = 0x01 << 7 //+ EDO mode enabled
	COL    NANDCR0 = 0x07 << 8 //+ Column address bit number
	COL_0  NANDCR0 = 0x00 << 8 //  16
	COL_1  NANDCR0 = 0x01 << 8 //  15
	COL_2  NANDCR0 = 0x02 << 8 //  14
	COL_3  NANDCR0 = 0x03 << 8 //  13
	COL_4  NANDCR0 = 0x04 << 8 //  12
	COL_5  NANDCR0 = 0x05 << 8 //  11
	COL_6  NANDCR0 = 0x06 << 8 //  10
	COL_7  NANDCR0 = 0x07 << 8 //  9
)

const (
	PSn     = 0
	SYNCENn = 1
	BLn     = 4
	EDOn    = 7
	COLn    = 8
)

const (
	CES   NANDCR1 = 0x0F << 0  //+ CE setup time
	CEH   NANDCR1 = 0x0F << 4  //+ CE hold time
	WEL   NANDCR1 = 0x0F << 8  //+ WE# LOW time
	WEH   NANDCR1 = 0x0F << 12 //+ WE# HIGH time
	REL   NANDCR1 = 0x0F << 16 //+ RE# LOW time
	REH   NANDCR1 = 0x0F << 20 //+ RE# HIGH time
	TA    NANDCR1 = 0x0F << 24 //+ Turnaround time
	CEITV NANDCR1 = 0x0F << 28 //+ CE# interval time
)

const (
	CESn   = 0
	CEHn   = 4
	WELn   = 8
	WEHn   = 12
	RELn   = 16
	REHn   = 20
	TAn    = 24
	CEITVn = 28
)

const (
	TWHR NANDCR2 = 0x3F << 0  //+ WE# HIGH to RE# LOW wait time
	TRHW NANDCR2 = 0x3F << 6  //+ RE# HIGH to WE# LOW wait time
	TADL NANDCR2 = 0x3F << 12 //+ ALE to WRITE Data start wait time
	TRR  NANDCR2 = 0x3F << 18 //+ Ready to RE# LOW min wait time
	TWB  NANDCR2 = 0x3F << 24 //+ WE# HIGH to busy wait time
)

const (
	TWHRn = 0
	TRHWn = 6
	TADLn = 12
	TRRn  = 18
	TWBn  = 24
)

const (
	NDOPT1 NANDCR3 = 0x01 << 0  //+ NAND option bit 1
	NDOPT2 NANDCR3 = 0x01 << 1  //+ NAND option bit 2
	NDOPT3 NANDCR3 = 0x01 << 2  //+ NAND option bit 3
	CLE    NANDCR3 = 0x01 << 3  //+ NAND CLE Option
	RDS    NANDCR3 = 0x0F << 16 //+ Read Data Setup cycle time.
	RDH    NANDCR3 = 0x0F << 20 //+ Read Data Hold cycle time.
	WDS    NANDCR3 = 0x0F << 24 //+ Write Data Setup cycle time.
	WDH    NANDCR3 = 0x0F << 28 //+ Write Data Hold cycle time.
)

const (
	NDOPT1n = 0
	NDOPT2n = 1
	NDOPT3n = 2
	CLEn    = 3
	RDSn    = 16
	RDHn    = 20
	WDSn    = 24
	WDHn    = 28
)

const (
	PS     NORCR0 = 0x01 << 0  //+ Port Size
	SYNCEN NORCR0 = 0x01 << 1  //+ Select NOR controller mode.
	BL     NORCR0 = 0x07 << 4  //+ Burst Length
	BL_0   NORCR0 = 0x00 << 4  //  1
	BL_1   NORCR0 = 0x01 << 4  //  2
	BL_2   NORCR0 = 0x02 << 4  //  4
	BL_3   NORCR0 = 0x03 << 4  //  8
	BL_4   NORCR0 = 0x04 << 4  //  16
	BL_5   NORCR0 = 0x05 << 4  //  32
	BL_6   NORCR0 = 0x06 << 4  //  64
	BL_7   NORCR0 = 0x07 << 4  //  64
	AM     NORCR0 = 0x03 << 8  //+ Address Mode
	AM_0   NORCR0 = 0x00 << 8  //  Address/Data MUX mode
	AM_1   NORCR0 = 0x01 << 8  //  Advanced Address/Data MUX mode
	AM_2   NORCR0 = 0x02 << 8  //  Address/Data non-MUX mode
	AM_3   NORCR0 = 0x03 << 8  //  Address/Data non-MUX mode
	ADVP   NORCR0 = 0x01 << 10 //+ ADV# polarity
	ADVH   NORCR0 = 0x01 << 11 //+ ADV# level control during address hold state
	COL    NORCR0 = 0x0F << 12 //+ Column Address bit width
	COL_0  NORCR0 = 0x00 << 12 //  12 Bits
	COL_1  NORCR0 = 0x01 << 12 //  11 Bits
	COL_2  NORCR0 = 0x02 << 12 //  10 Bits
	COL_3  NORCR0 = 0x03 << 12 //  9 Bits
	COL_4  NORCR0 = 0x04 << 12 //  8 Bits
	COL_5  NORCR0 = 0x05 << 12 //  7 Bits
	COL_6  NORCR0 = 0x06 << 12 //  6 Bits
	COL_7  NORCR0 = 0x07 << 12 //  5 Bits
	COL_8  NORCR0 = 0x08 << 12 //  4 Bits
	COL_9  NORCR0 = 0x09 << 12 //  3 Bits
	COL_10 NORCR0 = 0x0A << 12 //  2 Bits
	COL_11 NORCR0 = 0x0B << 12 //  12 Bits
	COL_12 NORCR0 = 0x0C << 12 //  12 Bits
	COL_13 NORCR0 = 0x0D << 12 //  12 Bits
	COL_14 NORCR0 = 0x0E << 12 //  12 Bits
	COL_15 NORCR0 = 0x0F << 12 //  12 Bits
)

const (
	PSn     = 0
	SYNCENn = 1
	BLn     = 4
	AMn     = 8
	ADVPn   = 10
	ADVHn   = 11
	COLn    = 12
)

const (
	CES NORCR1 = 0x0F << 0  //+ CE setup time cycle
	CEH NORCR1 = 0x0F << 4  //+ CE hold min time (CEH+1) cycle
	AS  NORCR1 = 0x0F << 8  //+ Address setup time
	AH  NORCR1 = 0x0F << 12 //+ Address hold time
	WEL NORCR1 = 0x0F << 16 //+ WE LOW time (WEL+1) cycle
	WEH NORCR1 = 0x0F << 20 //+ WE HIGH time (WEH+1) cycle
	REL NORCR1 = 0x0F << 24 //+ RE LOW time (REL+1) cycle
	REH NORCR1 = 0x0F << 28 //+ RE HIGH time (REH+1) cycle
)

const (
	CESn = 0
	CEHn = 4
	ASn  = 8
	AHn  = 12
	WELn = 16
	WEHn = 20
	RELn = 24
	REHn = 28
)

const (
	TA    NORCR2 = 0x0F << 8  //+ Turnaround time cycle
	AWDH  NORCR2 = 0x0F << 12 //+ Address to write data hold time cycle
	LC    NORCR2 = 0x0F << 16 //+ Latency count
	RD    NORCR2 = 0x0F << 20 //+ Read cycle time
	CEITV NORCR2 = 0x0F << 24 //+ CE# interval min time
	RDH   NORCR2 = 0x0F << 28 //+ Read cycle hold time
)

const (
	TAn    = 8
	AWDHn  = 12
	LCn    = 16
	RDn    = 20
	CEITVn = 24
	RDHn   = 28
)

const (
	ASSR NORCR3 = 0x0F << 0 //+ Address setup time for synchronous read
	AHSR NORCR3 = 0x0F << 4 //+ Address hold time for synchronous read
)

const (
	ASSRn = 0
	AHSRn = 4
)

const (
	PS     SRAMCR0 = 0x01 << 0  //+ Port Size
	SYNCEN SRAMCR0 = 0x01 << 1  //+ Select SRAM controller mode.
	BL     SRAMCR0 = 0x07 << 4  //+ Burst Length
	BL_0   SRAMCR0 = 0x00 << 4  //  1
	BL_1   SRAMCR0 = 0x01 << 4  //  2
	BL_2   SRAMCR0 = 0x02 << 4  //  4
	BL_3   SRAMCR0 = 0x03 << 4  //  8
	BL_4   SRAMCR0 = 0x04 << 4  //  16
	BL_5   SRAMCR0 = 0x05 << 4  //  32
	BL_6   SRAMCR0 = 0x06 << 4  //  64
	BL_7   SRAMCR0 = 0x07 << 4  //  64
	AM     SRAMCR0 = 0x03 << 8  //+ Address Mode
	AM_0   SRAMCR0 = 0x00 << 8  //  Address/Data MUX mode
	AM_1   SRAMCR0 = 0x01 << 8  //  Advanced Address/Data MUX mode
	AM_2   SRAMCR0 = 0x02 << 8  //  Address/Data non-MUX mode
	AM_3   SRAMCR0 = 0x03 << 8  //  Address/Data non-MUX mode
	ADVP   SRAMCR0 = 0x01 << 10 //+ ADV# polarity
	ADVH   SRAMCR0 = 0x01 << 11 //+ ADV# level control during address hold state
	COL    SRAMCR0 = 0x0F << 12 //+ Column Address bit width
	COL_0  SRAMCR0 = 0x00 << 12 //  12 Bits
	COL_1  SRAMCR0 = 0x01 << 12 //  11 Bits
	COL_2  SRAMCR0 = 0x02 << 12 //  10 Bits
	COL_3  SRAMCR0 = 0x03 << 12 //  9 Bits
	COL_4  SRAMCR0 = 0x04 << 12 //  8 Bits
	COL_5  SRAMCR0 = 0x05 << 12 //  7 Bits
	COL_6  SRAMCR0 = 0x06 << 12 //  6 Bits
	COL_7  SRAMCR0 = 0x07 << 12 //  5 Bits
	COL_8  SRAMCR0 = 0x08 << 12 //  4 Bits
	COL_9  SRAMCR0 = 0x09 << 12 //  3 Bits
	COL_10 SRAMCR0 = 0x0A << 12 //  2 Bits
	COL_11 SRAMCR0 = 0x0B << 12 //  12 Bits
	COL_12 SRAMCR0 = 0x0C << 12 //  12 Bits
	COL_13 SRAMCR0 = 0x0D << 12 //  12 Bits
	COL_14 SRAMCR0 = 0x0E << 12 //  12 Bits
	COL_15 SRAMCR0 = 0x0F << 12 //  12 Bits
)

const (
	PSn     = 0
	SYNCENn = 1
	BLn     = 4
	AMn     = 8
	ADVPn   = 10
	ADVHn   = 11
	COLn    = 12
)

const (
	CES SRAMCR1 = 0x0F << 0  //+ CE setup time cycle
	CEH SRAMCR1 = 0x0F << 4  //+ CE hold min time
	AS  SRAMCR1 = 0x0F << 8  //+ Address setup time
	AH  SRAMCR1 = 0x0F << 12 //+ Address hold time
	WEL SRAMCR1 = 0x0F << 16 //+ WE LOW time (WEL+1) cycle
	WEH SRAMCR1 = 0x0F << 20 //+ WE HIGH time (WEH+1) cycle
	REL SRAMCR1 = 0x0F << 24 //+ RE LOW time (REL+1) cycle
	REH SRAMCR1 = 0x0F << 28 //+ RE HIGH time (REH+1) cycle
)

const (
	CESn = 0
	CEHn = 4
	ASn  = 8
	AHn  = 12
	WELn = 16
	WEHn = 20
	RELn = 24
	REHn = 28
)

const (
	WDS   SRAMCR2 = 0x0F << 0  //+ Write Data setup time (WDS+1) cycle
	WDH   SRAMCR2 = 0x0F << 4  //+ Write Data hold time WDH cycle
	TA    SRAMCR2 = 0x0F << 8  //+ Turnaround time cycle
	AWDH  SRAMCR2 = 0x0F << 12 //+ Address to write data hold time cycle
	LC    SRAMCR2 = 0x0F << 16 //+ Latency count
	RD    SRAMCR2 = 0x0F << 20 //+ Read cycle time
	CEITV SRAMCR2 = 0x0F << 24 //+ CE# interval min time
	RDH   SRAMCR2 = 0x0F << 28 //+ Read cycle hold time
)

const (
	WDSn   = 0
	WDHn   = 4
	TAn    = 8
	AWDHn  = 12
	LCn    = 16
	RDn    = 20
	CEITVn = 24
	RDHn   = 28
)

const (
	PS     DBICR0 = 0x01 << 0  //+ Port Size
	BL     DBICR0 = 0x07 << 4  //+ Burst Length
	BL_0   DBICR0 = 0x00 << 4  //  1
	BL_1   DBICR0 = 0x01 << 4  //  2
	BL_2   DBICR0 = 0x02 << 4  //  4
	BL_3   DBICR0 = 0x03 << 4  //  8
	BL_4   DBICR0 = 0x04 << 4  //  16
	BL_5   DBICR0 = 0x05 << 4  //  32
	BL_6   DBICR0 = 0x06 << 4  //  64
	BL_7   DBICR0 = 0x07 << 4  //  64
	COL    DBICR0 = 0x0F << 12 //+ Column Address bit width
	COL_0  DBICR0 = 0x00 << 12 //  12 Bits
	COL_1  DBICR0 = 0x01 << 12 //  11 Bits
	COL_2  DBICR0 = 0x02 << 12 //  10 Bits
	COL_3  DBICR0 = 0x03 << 12 //  9 Bits
	COL_4  DBICR0 = 0x04 << 12 //  8 Bits
	COL_5  DBICR0 = 0x05 << 12 //  7 Bits
	COL_6  DBICR0 = 0x06 << 12 //  6 Bits
	COL_7  DBICR0 = 0x07 << 12 //  5 Bits
	COL_8  DBICR0 = 0x08 << 12 //  4 Bits
	COL_9  DBICR0 = 0x09 << 12 //  3 Bits
	COL_10 DBICR0 = 0x0A << 12 //  2 Bits
	COL_11 DBICR0 = 0x0B << 12 //  12 Bits
	COL_12 DBICR0 = 0x0C << 12 //  12 Bits
	COL_13 DBICR0 = 0x0D << 12 //  12 Bits
	COL_14 DBICR0 = 0x0E << 12 //  12 Bits
	COL_15 DBICR0 = 0x0F << 12 //  12 Bits
)

const (
	PSn  = 0
	BLn  = 4
	COLn = 12
)

const (
	CES   DBICR1 = 0x0F << 0  //+ CSX Setup Time
	CEH   DBICR1 = 0x0F << 4  //+ CSX Hold Time
	WEL   DBICR1 = 0x0F << 8  //+ WRX Low Time
	WEH   DBICR1 = 0x0F << 12 //+ WRX High Time
	REL   DBICR1 = 0x3F << 16 //+ RDX Low Time
	REH   DBICR1 = 0x3F << 22 //+ RDX High Time
	CEITV DBICR1 = 0x0F << 28 //+ CSX interval min time
)

const (
	CESn   = 0
	CEHn   = 4
	WELn   = 8
	WEHn   = 12
	RELn   = 16
	REHn   = 22
	CEITVn = 28
)

const (
	SA IPCR0 = 0xFFFFFFFF << 0 //+ Slave address
)

const (
	SAn = 0
)

const (
	DATSZ         IPCR1 = 0x07 << 0 //+ Data Size in Byte
	DATSZ_0       IPCR1 = 0x00 << 0 //  4
	DATSZ_1       IPCR1 = 0x01 << 0 //  1
	DATSZ_2       IPCR1 = 0x02 << 0 //  2
	DATSZ_3       IPCR1 = 0x03 << 0 //  3
	DATSZ_4       IPCR1 = 0x04 << 0 //  4
	DATSZ_5       IPCR1 = 0x05 << 0 //  4
	DATSZ_6       IPCR1 = 0x06 << 0 //  4
	DATSZ_7       IPCR1 = 0x07 << 0 //  4
	NAND_EXT_ADDR IPCR1 = 0xFF << 8 //+ NAND Extended Address
)

const (
	DATSZn         = 0
	NAND_EXT_ADDRn = 8
)

const (
	BM0 IPCR2 = 0x01 << 0 //+ Byte Mask for Byte 0 (IPTXD bit 7:0)
	BM1 IPCR2 = 0x01 << 1 //+ Byte Mask for Byte 1 (IPTXD bit 15:8)
	BM2 IPCR2 = 0x01 << 2 //+ Byte Mask for Byte 2 (IPTXD bit 23:16)
	BM3 IPCR2 = 0x01 << 3 //+ Byte Mask for Byte 3 (IPTXD bit 31:24)
)

const (
	BM0n = 0
	BM1n = 1
	BM2n = 2
	BM3n = 3
)

const (
	CMD IPCMD = 0xFFFF << 0  //+ SDRAM Commands: 0x8: READ 0x9: WRITE 0xA: MODESET 0xB: ACTIVE 0xC: AUTO REFRESH 0xD: SELF REFRESH 0xE: PRECHARGE 0xF: PRECHARGE ALL Others: RSVD SELF REFRESH will be sent to all SDRAM devices because they shared same SEMC_CLK pin
	KEY IPCMD = 0xFFFF << 16 //+ This field should be written with 0xA55A when trigging an IP command for all device types
)

const (
	CMDn = 0
	KEYn = 16
)

const (
	DAT IPTXDAT = 0xFFFFFFFF << 0 //+ no description available
)

const (
	DATn = 0
)

const (
	DAT IPRXDAT = 0xFFFFFFFF << 0 //+ no description available
)

const (
	DATn = 0
)

const (
	IDLE  STS0 = 0x01 << 0 //+ Indicating whether SEMC is in IDLE state.
	NARDY STS0 = 0x01 << 1 //+ Indicating NAND device Ready/WAIT# pin level.
)

const (
	IDLEn  = 0
	NARDYn = 1
)

const (
	NDWRPEND STS2 = 0x01 << 3 //+ This field indicating whether there is pending AXI command (write) to NAND device.
)

const (
	NDWRPENDn = 3
)

const (
	NDADDR STS12 = 0xFFFFFFFF << 0 //+ This field indicating the last write address (AXI command) to NAND device (without base address in SEMC_BR4).
)

const (
	NDADDRn = 0
)

const (
	SLVLOCK STS13 = 0x01 << 0 //+ Sample clock slave delay line locked.
	REFLOCK STS13 = 0x01 << 1 //+ Sample clock reference delay line locked.
	SLVSEL  STS13 = 0x3F << 2 //+ Sample clock slave delay line delay cell number selection .
	REFSEL  STS13 = 0x3F << 8 //+ Sample clock reference delay line delay cell number selection.
)

const (
	SLVLOCKn = 0
	REFLOCKn = 1
	SLVSELn  = 2
	REFSELn  = 8
)
