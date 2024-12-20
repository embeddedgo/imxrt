// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package usbphy provides access to the registers of the USBPHY peripheral.
//
// Instances:
//
//	USBPHY1  USBPHY1_BASE  -  USB_PHY1*  USBPHY Register Reference Index
//	USBPHY2  USBPHY2_BASE  -  USB_PHY2*  USBPHY Register Reference Index
//
// Registers:
//
//	0x000 32  PWD                 USB PHY Power-Down Register
//	0x004 32  PWD_SET(PWD)        USB PHY Power-Down Register
//	0x008 32  PWD_CLR(PWD)        USB PHY Power-Down Register
//	0x00C 32  PWD_TOG(PWD)        USB PHY Power-Down Register
//	0x010 32  TX                  USB PHY Transmitter Control Register
//	0x014 32  TX_SET(TX)          USB PHY Transmitter Control Register
//	0x018 32  TX_CLR(TX)          USB PHY Transmitter Control Register
//	0x01C 32  TX_TOG(TX)          USB PHY Transmitter Control Register
//	0x020 32  RX                  USB PHY Receiver Control Register
//	0x024 32  RX_SET(RX)          USB PHY Receiver Control Register
//	0x028 32  RX_CLR(RX)          USB PHY Receiver Control Register
//	0x02C 32  RX_TOG(RX)          USB PHY Receiver Control Register
//	0x030 32  CTRL                USB PHY General Control Register
//	0x034 32  CTRL_SET(CTRL)      USB PHY General Control Register
//	0x038 32  CTRL_CLR(CTRL)      USB PHY General Control Register
//	0x03C 32  CTRL_TOG(CTRL)      USB PHY General Control Register
//	0x040 32  STATUS              USB PHY Status Register
//	0x050 32  DEBUG               USB PHY Debug Register
//	0x054 32  DEBUG_SET(DEBUG)    USB PHY Debug Register
//	0x058 32  DEBUG_CLR(DEBUG)    USB PHY Debug Register
//	0x05C 32  DEBUG_TOG(DEBUG)    USB PHY Debug Register
//	0x060 32  DEBUG0_STATUS       UTMI Debug Status Register 0
//	0x070 32  DEBUG1              UTMI Debug Status Register 1
//	0x074 32  DEBUG1_SET(DEBUG1)  UTMI Debug Status Register 1
//	0x078 32  DEBUG1_CLR(DEBUG1)  UTMI Debug Status Register 1
//	0x07C 32  DEBUG1_TOG(DEBUG1)  UTMI Debug Status Register 1
//	0x080 32  VERSION             UTMI RTL Version
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package usbphy

const (
	_          PWD = 0x3FF << 0  //+ Reserved.
	TXPWDFS    PWD = 0x01 << 10  //+ 0 = Normal operation
	TXPWDIBIAS PWD = 0x01 << 11  //+ 0 = Normal operation
	TXPWDV2I   PWD = 0x01 << 12  //+ 0 = Normal operation
	_          PWD = 0x0F << 13  //+ Reserved.
	RXPWDENV   PWD = 0x01 << 17  //+ 0 = Normal operation
	RXPWD1PT1  PWD = 0x01 << 18  //+ 0 = Normal operation
	RXPWDDIFF  PWD = 0x01 << 19  //+ 0 = Normal operation
	RXPWDRX    PWD = 0x01 << 20  //+ 0 = Normal operation
	_          PWD = 0x7FF << 21 //+ Reserved.
)

const (
	TXPWDFSn    = 10
	TXPWDIBIASn = 11
	TXPWDV2In   = 12
	RXPWDENVn   = 17
	RXPWD1PT1n  = 18
	RXPWDDIFFn  = 19
	RXPWDRXn    = 20
)

const (
	D_CAL              TX = 0x0F << 0  //+ Resistor Trimming Code: 0000 = 0.16% 0111 = Nominal 1111 = +25%
	_                  TX = 0x0F << 4  //+ Reserved. Note: This bit should remain clear.
	TXCAL45DN          TX = 0x0F << 8  //+ Decode to select a 45-Ohm resistance to the USB_DN output pin
	_                  TX = 0x0F << 12 //+ Reserved. Note: This bit should remain clear.
	TXCAL45DP          TX = 0x0F << 16 //+ Decode to select a 45-Ohm resistance to the USB_DP output pin
	_                  TX = 0x3F << 20 //+ Reserved.
	USBPHY_TX_EDGECTRL TX = 0x07 << 26 //+ Controls the edge-rate of the current sensing transistors used in HS transmit
	_                  TX = 0x07 << 29 //+ Reserved.
)

const (
	D_CALn              = 0
	TXCAL45DNn          = 8
	TXCAL45DPn          = 16
	USBPHY_TX_EDGECTRLn = 26
)

const (
	ENVADJ    RX = 0x07 << 0   //+ The ENVADJ field adjusts the trip point for the envelope detector
	_         RX = 0x01 << 3   //+ Reserved.
	DISCONADJ RX = 0x07 << 4   //+ The DISCONADJ field adjusts the trip point for the disconnect detector: 000 = Trip-Level Voltage is 0
	_         RX = 0x7FFF << 7 //+ Reserved.
	RXDBYPASS RX = 0x01 << 22  //+ 0 = Normal operation
	_         RX = 0x1FF << 23 //+ Reserved.
)

const (
	ENVADJn    = 0
	DISCONADJn = 4
	RXDBYPASSn = 22
)

const (
	ENOTG_ID_CHG_IRQ     CTRL = 0x01 << 0  //+ Enable OTG_ID_CHG_IRQ.
	ENHOSTDISCONDETECT   CTRL = 0x01 << 1  //+ For host mode, enables high-speed disconnect detector
	ENIRQHOSTDISCON      CTRL = 0x01 << 2  //+ Enables interrupt for detection of disconnection to Device when in high-speed host mode
	HOSTDISCONDETECT_IRQ CTRL = 0x01 << 3  //+ Indicates that the device has disconnected in high-speed mode
	ENDEVPLUGINDETECT    CTRL = 0x01 << 4  //+ For device mode, enables 200-KOhm pullups for detecting connectivity to the host.
	DEVPLUGIN_POLARITY   CTRL = 0x01 << 5  //+ For device mode, if this bit is cleared to 0, then it trips the interrupt if the device is plugged in
	OTG_ID_CHG_IRQ       CTRL = 0x01 << 6  //+ OTG ID change interrupt. Indicates the value of ID pin changed.
	ENOTGIDDETECT        CTRL = 0x01 << 7  //+ Enables circuit to detect resistance of MiniAB ID pin.
	RESUMEIRQSTICKY      CTRL = 0x01 << 8  //+ Set to 1 will make RESUME_IRQ bit a sticky bit until software clear it
	ENIRQRESUMEDETECT    CTRL = 0x01 << 9  //+ Enables interrupt for detection of a non-J state on the USB line
	RESUME_IRQ           CTRL = 0x01 << 10 //+ Indicates that the host is sending a wake-up after suspend
	ENIRQDEVPLUGIN       CTRL = 0x01 << 11 //+ Enables interrupt for the detection of connectivity to the USB line.
	DEVPLUGIN_IRQ        CTRL = 0x01 << 12 //+ Indicates that the device is connected
	DATA_ON_LRADC        CTRL = 0x01 << 13 //+ Enables the LRADC to monitor USB_DP and USB_DM. This is for use in non-USB modes only.
	ENUTMILEVEL2         CTRL = 0x01 << 14 //+ Enables UTMI+ Level2. This should be enabled if needs to support LS device
	ENUTMILEVEL3         CTRL = 0x01 << 15 //+ Enables UTMI+ Level3
	ENIRQWAKEUP          CTRL = 0x01 << 16 //+ Enables interrupt for the wakeup events.
	WAKEUP_IRQ           CTRL = 0x01 << 17 //+ Indicates that there is a wakeup event
	ENAUTO_PWRON_PLL     CTRL = 0x01 << 18 //+ Enables the feature to auto-enable the POWER bit of HW_CLKCTRL_PLLxCTRL0 if there is wakeup event if USB is suspended
	ENAUTOCLR_CLKGATE    CTRL = 0x01 << 19 //+ Enables the feature to auto-clear the CLKGATE bit if there is wakeup event while USB is suspended
	ENAUTOCLR_PHY_PWD    CTRL = 0x01 << 20 //+ Enables the feature to auto-clear the PWD register bits in USBPHYx_PWD if there is wakeup event while USB is suspended
	ENDPDMCHG_WKUP       CTRL = 0x01 << 21 //+ Enables the feature to wakeup USB if DP/DM is toggled when USB is suspended
	ENIDCHG_WKUP         CTRL = 0x01 << 22 //+ Enables the feature to wakeup USB if ID is toggled when USB is suspended.
	ENVBUSCHG_WKUP       CTRL = 0x01 << 23 //+ Enables the feature to wakeup USB if VBUS is toggled when USB is suspended.
	FSDLL_RST_EN         CTRL = 0x01 << 24 //+ Enables the feature to reset the FSDLL lock detection logic at the end of each TX packet.
	_                    CTRL = 0x03 << 25 //+ Reserved.
	OTG_ID_VALUE         CTRL = 0x01 << 27 //+ Almost same as OTGID_STATUS in USBPHYx_STATUS Register
	HOST_FORCE_LS_SE0    CTRL = 0x01 << 28 //+ Forces the next FS packet that is transmitted to have a EOP with LS timing
	UTMI_SUSPENDM        CTRL = 0x01 << 29 //+ Used by the PHY to indicate a powered-down state
	CLKGATE              CTRL = 0x01 << 30 //+ Gate UTMI Clocks
	SFTRST               CTRL = 0x01 << 31 //+ Writing a 1 to this bit will soft-reset the USBPHYx_PWD, USBPHYx_TX, USBPHYx_RX, and USBPHYx_CTRL registers
)

const (
	ENOTG_ID_CHG_IRQn     = 0
	ENHOSTDISCONDETECTn   = 1
	ENIRQHOSTDISCONn      = 2
	HOSTDISCONDETECT_IRQn = 3
	ENDEVPLUGINDETECTn    = 4
	DEVPLUGIN_POLARITYn   = 5
	OTG_ID_CHG_IRQn       = 6
	ENOTGIDDETECTn        = 7
	RESUMEIRQSTICKYn      = 8
	ENIRQRESUMEDETECTn    = 9
	RESUME_IRQn           = 10
	ENIRQDEVPLUGINn       = 11
	DEVPLUGIN_IRQn        = 12
	DATA_ON_LRADCn        = 13
	ENUTMILEVEL2n         = 14
	ENUTMILEVEL3n         = 15
	ENIRQWAKEUPn          = 16
	WAKEUP_IRQn           = 17
	ENAUTO_PWRON_PLLn     = 18
	ENAUTOCLR_CLKGATEn    = 19
	ENAUTOCLR_PHY_PWDn    = 20
	ENDPDMCHG_WKUPn       = 21
	ENIDCHG_WKUPn         = 22
	ENVBUSCHG_WKUPn       = 23
	FSDLL_RST_ENn         = 24
	OTG_ID_VALUEn         = 27
	HOST_FORCE_LS_SE0n    = 28
	UTMI_SUSPENDMn        = 29
	CLKGATEn              = 30
	SFTRSTn               = 31
)

const (
	_                       STATUS = 0x07 << 0      //+ Reserved.
	HOSTDISCONDETECT_STATUS STATUS = 0x01 << 3      //+ Indicates that the device has disconnected while in high-speed host mode.
	_                       STATUS = 0x03 << 4      //+ Reserved.
	DEVPLUGIN_STATUS        STATUS = 0x01 << 6      //+ Indicates that the device has been connected on the USB_DP and USB_DM lines.
	_                       STATUS = 0x01 << 7      //+ Reserved.
	OTGID_STATUS            STATUS = 0x01 << 8      //+ Indicates the results of ID pin on MiniAB plug
	_                       STATUS = 0x01 << 9      //+ Reserved.
	RESUME_STATUS           STATUS = 0x01 << 10     //+ Indicates that the host is sending a wake-up after suspend and has triggered an interrupt.
	_                       STATUS = 0x1FFFFF << 11 //+ Reserved.
)

const (
	HOSTDISCONDETECT_STATUSn = 3
	DEVPLUGIN_STATUSn        = 6
	OTGID_STATUSn            = 8
	RESUME_STATUSn           = 10
)

const (
	OTGIDPIOLOCK         DEBUG = 0x01 << 0  //+ Once OTG ID from USBPHYx_STATUS_OTGID_STATUS, use this to hold the value
	DEBUG_INTERFACE_HOLD DEBUG = 0x01 << 1  //+ Use holding registers to assist in timing for external UTMI interface.
	HSTPULLDOWN          DEBUG = 0x03 << 2  //+ Set bit 3 to 1 to pull down 15-KOhm on USB_DP line
	ENHSTPULLDOWN        DEBUG = 0x03 << 4  //+ Set bit 5 to 1 to override the control of the USB_DP 15-KOhm pulldown
	_                    DEBUG = 0x03 << 6  //+ Reserved.
	TX2RXCOUNT           DEBUG = 0x0F << 8  //+ Delay in between the end of transmit to the beginning of receive
	ENTX2RXCOUNT         DEBUG = 0x01 << 12 //+ Set this bit to allow a countdown to transition in between TX and RX.
	_                    DEBUG = 0x07 << 13 //+ Reserved.
	SQUELCHRESETCOUNT    DEBUG = 0x1F << 16 //+ Delay in between the detection of squelch to the reset of high-speed RX.
	_                    DEBUG = 0x07 << 21 //+ Reserved.
	ENSQUELCHRESET       DEBUG = 0x01 << 24 //+ Set bit to allow squelch to reset high-speed receive.
	SQUELCHRESETLENGTH   DEBUG = 0x0F << 25 //+ Duration of RESET in terms of the number of 480-MHz cycles.
	HOST_RESUME_DEBUG    DEBUG = 0x01 << 29 //+ Choose to trigger the host resume SE0 with HOST_FORCE_LS_SE0 = 0 or UTMI_SUSPEND = 1.
	GATECLK              DEBUG = 0x01 << 30 //+ Gate Test Clocks
	_                    DEBUG = 0x01 << 31 //+ Reserved.
)

const (
	OTGIDPIOLOCKn         = 0
	DEBUG_INTERFACE_HOLDn = 1
	HSTPULLDOWNn          = 2
	ENHSTPULLDOWNn        = 4
	TX2RXCOUNTn           = 8
	ENTX2RXCOUNTn         = 12
	SQUELCHRESETCOUNTn    = 16
	ENSQUELCHRESETn       = 24
	SQUELCHRESETLENGTHn   = 25
	HOST_RESUME_DEBUGn    = 29
	GATECLKn              = 30
)

const (
	LOOP_BACK_FAIL_COUNT    DEBUG0_STATUS = 0xFFFF << 0 //+ Running count of the failed pseudo-random generator loopback
	UTMI_RXERROR_FAIL_COUNT DEBUG0_STATUS = 0x3FF << 16 //+ Running count of the UTMI_RXERROR.
	SQUELCH_COUNT           DEBUG0_STATUS = 0x3F << 26  //+ Running count of the squelch reset instead of normal end for HS RX.
)

const (
	LOOP_BACK_FAIL_COUNTn    = 0
	UTMI_RXERROR_FAIL_COUNTn = 16
	SQUELCH_COUNTn           = 26
)

const (
	_           DEBUG1 = 0x1FFF << 0   //+ Reserved. Note: This bit should remain clear.
	ENTAILADJVD DEBUG1 = 0x03 << 13    //+ Delay increment of the rise of squelch: 00 = Delay is nominal 01 = Delay is +20% 10 = Delay is -20% 11 = Delay is -40%
	_           DEBUG1 = 0x1FFFF << 15 //+ Reserved.
)

const (
	ENTAILADJVDn = 13
)

const (
	STEP  VERSION = 0xFFFF << 0 //+ Fixed read-only value reflecting the stepping of the RTL version.
	MINOR VERSION = 0xFF << 16  //+ Fixed read-only value reflecting the MINOR field of the RTL version.
	MAJOR VERSION = 0xFF << 24  //+ Fixed read-only value reflecting the MAJOR field of the RTL version.
)

const (
	STEPn  = 0
	MINORn = 16
	MAJORn = 24
)
