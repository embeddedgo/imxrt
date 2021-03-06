// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package ocotp provides access to the registers of the OCOTP peripheral.
//
// Instances:
//  OCOTP  OCOTP_BASE  -  -
// Registers:
//  0x000 32  CTRL            OTP Controller Control Register
//  0x004 32  CTRL_SET        OTP Controller Control Register
//  0x008 32  CTRL_CLR        OTP Controller Control Register
//  0x00C 32  CTRL_TOG        OTP Controller Control Register
//  0x010 32  TIMING          OTP Controller Timing Register
//  0x020 32  DATA            OTP Controller Write Data Register
//  0x030 32  READ_CTRL       OTP Controller Write Data Register
//  0x040 32  READ_FUSE_DATA  OTP Controller Read Data Register
//  0x050 32  SW_STICKY       Sticky bit Register
//  0x060 32  SCS             Software Controllable Signals Register
//  0x064 32  SCS_SET         Software Controllable Signals Register
//  0x068 32  SCS_CLR         Software Controllable Signals Register
//  0x06C 32  SCS_TOG         Software Controllable Signals Register
//  0x070 32  CRC_ADDR        OTP Controller CRC test address
//  0x080 32  CRC_VALUE       OTP Controller CRC Value Register
//  0x090 32  VERSION         OTP Controller Version Register
//  0x100 32  TIMING2         OTP Controller Timing Register
//  0x400 32  LOCK            Value of OTP Bank0 Word0 (Lock controls)
//  0x410 32  CFG0            Value of OTP Bank0 Word1 (Configuration and Manufacturing Info.)
//  0x420 32  CFG1            Value of OTP Bank0 Word2 (Configuration and Manufacturing Info.)
//  0x430 32  CFG2            Value of OTP Bank0 Word3 (Configuration and Manufacturing Info.)
//  0x440 32  CFG3            Value of OTP Bank0 Word4 (Configuration and Manufacturing Info.)
//  0x450 32  CFG4            Value of OTP Bank0 Word5 (Configuration and Manufacturing Info.)
//  0x460 32  CFG5            Value of OTP Bank0 Word6 (Configuration and Manufacturing Info.)
//  0x470 32  CFG6            Value of OTP Bank0 Word7 (Configuration and Manufacturing Info.)
//  0x480 32  MEM0            Value of OTP Bank1 Word0 (Memory Related Info.)
//  0x490 32  MEM1            Value of OTP Bank1 Word1 (Memory Related Info.)
//  0x4A0 32  MEM2            Value of OTP Bank1 Word2 (Memory Related Info.)
//  0x4B0 32  MEM3            Value of OTP Bank1 Word3 (Memory Related Info.)
//  0x4C0 32  MEM4            Value of OTP Bank1 Word4 (Memory Related Info.)
//  0x4D0 32  ANA0            Value of OTP Bank1 Word5 (Memory Related Info.)
//  0x4E0 32  ANA1            Value of OTP Bank1 Word6 (General Purpose Customer Defined Info.)
//  0x4F0 32  ANA2            Value of OTP Bank1 Word7 (General Purpose Customer Defined Info.)
//  0x500 32  OTPMK0          Value of OTP Bank2 Word0 (OTPMK Key)
//  0x510 32  OTPMK1          Value of OTP Bank2 Word1 (OTPMK Key)
//  0x520 32  OTPMK2          Value of OTP Bank2 Word2 (OTPMK Key)
//  0x530 32  OTPMK3          Value of OTP Bank2 Word3 (OTPMK Key)
//  0x540 32  OTPMK4          Value of OTP Bank2 Word4 (OTPMK Key)
//  0x550 32  OTPMK5          Value of OTP Bank2 Word5 (OTPMK Key)
//  0x560 32  OTPMK6          Value of OTP Bank2 Word6 (OTPMK Key)
//  0x570 32  OTPMK7          Value of OTP Bank2 Word7 (OTPMK Key)
//  0x580 32  SRK0            Shadow Register for OTP Bank3 Word0 (SRK Hash)
//  0x590 32  SRK1            Shadow Register for OTP Bank3 Word1 (SRK Hash)
//  0x5A0 32  SRK2            Shadow Register for OTP Bank3 Word2 (SRK Hash)
//  0x5B0 32  SRK3            Shadow Register for OTP Bank3 Word3 (SRK Hash)
//  0x5C0 32  SRK4            Shadow Register for OTP Bank3 Word4 (SRK Hash)
//  0x5D0 32  SRK5            Shadow Register for OTP Bank3 Word5 (SRK Hash)
//  0x5E0 32  SRK6            Shadow Register for OTP Bank3 Word6 (SRK Hash)
//  0x5F0 32  SRK7            Shadow Register for OTP Bank3 Word7 (SRK Hash)
//  0x600 32  SJC_RESP0       Value of OTP Bank4 Word0 (Secure JTAG Response Field)
//  0x610 32  SJC_RESP1       Value of OTP Bank4 Word1 (Secure JTAG Response Field)
//  0x620 32  MAC0            Value of OTP Bank4 Word2 (MAC Address)
//  0x630 32  MAC1            Value of OTP Bank4 Word3 (MAC Address)
//  0x640 32  MAC2            Value of OTP Bank4 Word4 (MAC2 Address)
//  0x650 32  OTPMK_CRC32     Value of OTP Bank4 Word5 (CRC Key)
//  0x660 32  GP1             Value of OTP Bank4 Word6 (General Purpose Customer Defined Info)
//  0x670 32  GP2             Value of OTP Bank4 Word7 (General Purpose Customer Defined Info)
//  0x680 32  SW_GP1          Value of OTP Bank5 Word0 (SW GP1)
//  0x690 32  SW_GP20         Value of OTP Bank5 Word1 (SW GP2)
//  0x6A0 32  SW_GP21         Value of OTP Bank5 Word2 (SW GP2)
//  0x6B0 32  SW_GP22         Value of OTP Bank5 Word3 (SW GP2)
//  0x6C0 32  SW_GP23         Value of OTP Bank5 Word4 (SW GP2)
//  0x6D0 32  MISC_CONF0      Value of OTP Bank5 Word5 (Misc Conf)
//  0x6E0 32  MISC_CONF1      Value of OTP Bank5 Word6 (Misc Conf)
//  0x6F0 32  SRK_REVOKE      Value of OTP Bank5 Word7 (SRK Revoke)
//  0x800 32  ROM_PATCH0      Value of OTP Bank6 Word0 (ROM Patch)
//  0x810 32  ROM_PATCH1      Value of OTP Bank6 Word1 (ROM Patch)
//  0x820 32  ROM_PATCH2      Value of OTP Bank6 Word2 (ROM Patch)
//  0x830 32  ROM_PATCH3      Value of OTP Bank6 Word3 (ROM Patch)
//  0x840 32  ROM_PATCH4      Value of OTP Bank6 Word4 (ROM Patch)
//  0x850 32  ROM_PATCH5      Value of OTP Bank6 Word5 (ROM Patch)
//  0x860 32  ROM_PATCH6      Value of OTP Bank6 Word6 (ROM Patch)
//  0x870 32  ROM_PATCH7      Value of OTP Bank6 Word7 (ROM Patch)
//  0x880 32  GP30            Value of OTP Bank7 Word0 (GP3)
//  0x890 32  GP31            Value of OTP Bank7 Word1 (GP3)
//  0x8A0 32  GP32            Value of OTP Bank7 Word2 (GP3)
//  0x8B0 32  GP33            Value of OTP Bank7 Word3 (GP3)
//  0x8C0 32  GP40            Value of OTP Bank7 Word4 (GP4)
//  0x8D0 32  GP41            Value of OTP Bank7 Word5 (GP4)
//  0x8E0 32  GP42            Value of OTP Bank7 Word6 (GP4)
//  0x8F0 32  GP43            Value of OTP Bank7 Word7 (GP4)
// Import:
//  github.com/embeddedgo/imxrt/p/mmap
package ocotp

const (
	ADDR           CTRL = 0x3F << 0    //+ ADDR
	RSVD0          CTRL = 0x03 << 6    //+ RSVD0
	BUSY           CTRL = 0x01 << 8    //+ BUSY
	ERROR          CTRL = 0x01 << 9    //+ ERROR
	RELOAD_SHADOWS CTRL = 0x01 << 10   //+ RELOAD_SHADOWS
	CRC_TEST       CTRL = 0x01 << 11   //+ CRC_TEST
	CRC_FAIL       CTRL = 0x01 << 12   //+ CRC_FAIL
	RSVD1          CTRL = 0x07 << 13   //+ RSVD1
	WR_UNLOCK      CTRL = 0xFFFF << 16 //+ WR_UNLOCK
	KEY            CTRL = 0x3E77 << 16 //  Key needed to unlock HW_OCOTP_DATA register.
)

const (
	ADDRn           = 0
	RSVD0n          = 6
	BUSYn           = 8
	ERRORn          = 9
	RELOAD_SHADOWSn = 10
	CRC_TESTn       = 11
	CRC_FAILn       = 12
	RSVD1n          = 13
	WR_UNLOCKn      = 16
)

const (
	ADDR           CTRL_SET = 0x3F << 0    //+ ADDR
	RSVD0          CTRL_SET = 0x03 << 6    //+ RSVD0
	BUSY           CTRL_SET = 0x01 << 8    //+ BUSY
	ERROR          CTRL_SET = 0x01 << 9    //+ ERROR
	RELOAD_SHADOWS CTRL_SET = 0x01 << 10   //+ RELOAD_SHADOWS
	CRC_TEST       CTRL_SET = 0x01 << 11   //+ CRC_TEST
	CRC_FAIL       CTRL_SET = 0x01 << 12   //+ CRC_FAIL
	RSVD1          CTRL_SET = 0x07 << 13   //+ RSVD1
	WR_UNLOCK      CTRL_SET = 0xFFFF << 16 //+ WR_UNLOCK
)

const (
	ADDRn           = 0
	RSVD0n          = 6
	BUSYn           = 8
	ERRORn          = 9
	RELOAD_SHADOWSn = 10
	CRC_TESTn       = 11
	CRC_FAILn       = 12
	RSVD1n          = 13
	WR_UNLOCKn      = 16
)

const (
	ADDR           CTRL_CLR = 0x3F << 0    //+ ADDR
	RSVD0          CTRL_CLR = 0x03 << 6    //+ RSVD0
	BUSY           CTRL_CLR = 0x01 << 8    //+ BUSY
	ERROR          CTRL_CLR = 0x01 << 9    //+ ERROR
	RELOAD_SHADOWS CTRL_CLR = 0x01 << 10   //+ RELOAD_SHADOWS
	CRC_TEST       CTRL_CLR = 0x01 << 11   //+ CRC_TEST
	CRC_FAIL       CTRL_CLR = 0x01 << 12   //+ CRC_FAIL
	RSVD1          CTRL_CLR = 0x07 << 13   //+ RSVD1
	WR_UNLOCK      CTRL_CLR = 0xFFFF << 16 //+ WR_UNLOCK
)

const (
	ADDRn           = 0
	RSVD0n          = 6
	BUSYn           = 8
	ERRORn          = 9
	RELOAD_SHADOWSn = 10
	CRC_TESTn       = 11
	CRC_FAILn       = 12
	RSVD1n          = 13
	WR_UNLOCKn      = 16
)

const (
	ADDR           CTRL_TOG = 0x3F << 0    //+ ADDR
	RSVD0          CTRL_TOG = 0x03 << 6    //+ RSVD0
	BUSY           CTRL_TOG = 0x01 << 8    //+ BUSY
	ERROR          CTRL_TOG = 0x01 << 9    //+ ERROR
	RELOAD_SHADOWS CTRL_TOG = 0x01 << 10   //+ RELOAD_SHADOWS
	CRC_TEST       CTRL_TOG = 0x01 << 11   //+ CRC_TEST
	CRC_FAIL       CTRL_TOG = 0x01 << 12   //+ CRC_FAIL
	RSVD1          CTRL_TOG = 0x07 << 13   //+ RSVD1
	WR_UNLOCK      CTRL_TOG = 0xFFFF << 16 //+ WR_UNLOCK
)

const (
	ADDRn           = 0
	RSVD0n          = 6
	BUSYn           = 8
	ERRORn          = 9
	RELOAD_SHADOWSn = 10
	CRC_TESTn       = 11
	CRC_FAILn       = 12
	RSVD1n          = 13
	WR_UNLOCKn      = 16
)

const (
	STROBE_PROG TIMING = 0xFFF << 0 //+ STROBE_PROG
	RELAX       TIMING = 0x0F << 12 //+ RELAX
	STROBE_READ TIMING = 0x3F << 16 //+ STROBE_READ
	WAIT        TIMING = 0x3F << 22 //+ WAIT
	RSRVD0      TIMING = 0x0F << 28 //+ RSRVD0
)

const (
	STROBE_PROGn = 0
	RELAXn       = 12
	STROBE_READn = 16
	WAITn        = 22
	RSRVD0n      = 28
)

const (
	DATA DATA = 0xFFFFFFFF << 0 //+ DATA
)

const (
	DATAn = 0
)

const (
	READ_FUSE READ_CTRL = 0x01 << 0       //+ READ_FUSE
	RSVD0     READ_CTRL = 0x7FFFFFFF << 1 //+ RSVD0
)

const (
	READ_FUSEn = 0
	RSVD0n     = 1
)

const (
	DATA READ_FUSE_DATA = 0xFFFFFFFF << 0 //+ DATA
)

const (
	DATAn = 0
)

const (
	BLOCK_DTCP_KEY     SW_STICKY = 0x01 << 0      //+ BLOCK_DTCP_KEY
	SRK_REVOKE_LOCK    SW_STICKY = 0x01 << 1      //+ SRK_REVOKE_LOCK
	FIELD_RETURN_LOCK  SW_STICKY = 0x01 << 2      //+ FIELD_RETURN_LOCK
	BLOCK_ROM_PART     SW_STICKY = 0x01 << 3      //+ BLOCK_ROM_PART
	JTAG_BLOCK_RELEASE SW_STICKY = 0x01 << 4      //+ JTAG_BLOCK_RELEASE
	RSVD0              SW_STICKY = 0x7FFFFFF << 5 //+ RSVD0
)

const (
	BLOCK_DTCP_KEYn     = 0
	SRK_REVOKE_LOCKn    = 1
	FIELD_RETURN_LOCKn  = 2
	BLOCK_ROM_PARTn     = 3
	JTAG_BLOCK_RELEASEn = 4
	RSVD0n              = 5
)

const (
	HAB_JDE SCS = 0x01 << 0       //+ HAB_JDE
	SPARE   SCS = 0x3FFFFFFF << 1 //+ SPARE
	LOCK    SCS = 0x01 << 31      //+ LOCK
)

const (
	HAB_JDEn = 0
	SPAREn   = 1
	LOCKn    = 31
)

const (
	HAB_JDE SCS_SET = 0x01 << 0       //+ HAB_JDE
	SPARE   SCS_SET = 0x3FFFFFFF << 1 //+ SPARE
	LOCK    SCS_SET = 0x01 << 31      //+ LOCK
)

const (
	HAB_JDEn = 0
	SPAREn   = 1
	LOCKn    = 31
)

const (
	HAB_JDE SCS_CLR = 0x01 << 0       //+ HAB_JDE
	SPARE   SCS_CLR = 0x3FFFFFFF << 1 //+ SPARE
	LOCK    SCS_CLR = 0x01 << 31      //+ LOCK
)

const (
	HAB_JDEn = 0
	SPAREn   = 1
	LOCKn    = 31
)

const (
	HAB_JDE SCS_TOG = 0x01 << 0       //+ HAB_JDE
	SPARE   SCS_TOG = 0x3FFFFFFF << 1 //+ SPARE
	LOCK    SCS_TOG = 0x01 << 31      //+ LOCK
)

const (
	HAB_JDEn = 0
	SPAREn   = 1
	LOCKn    = 31
)

const (
	DATA_START_ADDR CRC_ADDR = 0xFF << 0  //+ DATA_START_ADDR
	DATA_END_ADDR   CRC_ADDR = 0xFF << 8  //+ DATA_END_ADDR
	CRC_ADDR        CRC_ADDR = 0xFF << 16 //+ CRC_ADDR
	OTPMK_CRC       CRC_ADDR = 0x01 << 24 //+ OTPMK_CRC
	RSVD0           CRC_ADDR = 0x7F << 25 //+ RSVD0
)

const (
	DATA_START_ADDRn = 0
	DATA_END_ADDRn   = 8
	CRC_ADDRn        = 16
	OTPMK_CRCn       = 24
	RSVD0n           = 25
)

const (
	DATA CRC_VALUE = 0xFFFFFFFF << 0 //+ DATA
)

const (
	DATAn = 0
)

const (
	STEP  VERSION = 0xFFFF << 0 //+ STEP
	MINOR VERSION = 0xFF << 16  //+ MINOR
	MAJOR VERSION = 0xFF << 24  //+ MAJOR
)

const (
	STEPn  = 0
	MINORn = 16
	MAJORn = 24
)

const (
	RELAX_PROG TIMING2 = 0xFFF << 0  //+ RELAX_PROG
	RSRVD0     TIMING2 = 0x0F << 12  //+ RSRVD0
	RELAX_READ TIMING2 = 0x3F << 16  //+ RELAX_READ
	RSRVD1     TIMING2 = 0x3FF << 22 //+ RSRVD0
)

const (
	RELAX_PROGn = 0
	RSRVD0n     = 12
	RELAX_READn = 16
	RSRVD1n     = 22
)

const (
	TESTER       LOCK = 0x03 << 0  //+ TESTER
	BOOT_CFG     LOCK = 0x03 << 2  //+ BOOT_CFG
	MEM_TRIM     LOCK = 0x03 << 4  //+ MEM_TRIM
	SJC_RESP     LOCK = 0x01 << 6  //+ SJC_RESP
	GP4_RLOCK    LOCK = 0x01 << 7  //+ GP4_RLOCK
	MAC_ADDR     LOCK = 0x03 << 8  //+ MAC_ADDR
	GP1          LOCK = 0x03 << 10 //+ GP1
	GP2          LOCK = 0x03 << 12 //+ GP2
	ROM_PATCH    LOCK = 0x01 << 15 //+ ROM_PATCH
	SW_GP1       LOCK = 0x01 << 16 //+ SW_GP1
	OTPMK        LOCK = 0x01 << 17 //+ OTPMK
	ANALOG       LOCK = 0x03 << 18 //+ ANALOG
	OTPMK_CRC    LOCK = 0x01 << 20 //+ OTPMK_CRC
	SW_GP2_LOCK  LOCK = 0x01 << 21 //+ SW_GP2_LOCK
	MISC_CONF    LOCK = 0x01 << 22 //+ MISC_CONF
	SW_GP2_RLOCK LOCK = 0x01 << 23 //+ SW_GP2_RLOCK
	GP4          LOCK = 0x03 << 24 //+ GP4
	GP3          LOCK = 0x03 << 26 //+ GP3
	FIELD_RETURN LOCK = 0x0F << 28 //+ FIELD_RETURN
)

const (
	TESTERn       = 0
	BOOT_CFGn     = 2
	MEM_TRIMn     = 4
	SJC_RESPn     = 6
	GP4_RLOCKn    = 7
	MAC_ADDRn     = 8
	GP1n          = 10
	GP2n          = 12
	ROM_PATCHn    = 15
	SW_GP1n       = 16
	OTPMKn        = 17
	ANALOGn       = 18
	OTPMK_CRCn    = 20
	SW_GP2_LOCKn  = 21
	MISC_CONFn    = 22
	SW_GP2_RLOCKn = 23
	GP4n          = 24
	GP3n          = 26
	FIELD_RETURNn = 28
)

const (
	BITS CFG0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG3 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG4 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG5 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS CFG6 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MEM0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MEM1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MEM2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MEM3 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MEM4 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ANA0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ANA1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ANA2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK3 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK4 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK5 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK6 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK7 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK3 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK4 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK5 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK6 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK7 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SJC_RESP0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SJC_RESP1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MAC0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MAC1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MAC2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS OTPMK_CRC32 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SW_GP1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SW_GP20 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SW_GP21 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SW_GP22 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SW_GP23 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MISC_CONF0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS MISC_CONF1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS SRK_REVOKE = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH0 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH1 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH2 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH3 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH4 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH5 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH6 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS ROM_PATCH7 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP30 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP31 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP32 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP33 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP40 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP41 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP42 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)

const (
	BITS GP43 = 0xFFFFFFFF << 0 //+ BITS
)

const (
	BITSn = 0
)
