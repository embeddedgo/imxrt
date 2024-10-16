// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package xtaloscm provides access to the registers of the XTALOSC24M peripheral.
//
// Instances:
//
//	XTALOSC24M  XTALOSC24M_BASE  -  -  XTALOSC24M
//
// Registers:
//
//	0x150 32  MISC0            Miscellaneous Register 0
//	0x154 32  MISC0_SET        Miscellaneous Register 0
//	0x158 32  MISC0_CLR        Miscellaneous Register 0
//	0x15C 32  MISC0_TOG        Miscellaneous Register 0
//	0x270 32  LOWPWR_CTRL      XTAL OSC (LP) Control Register
//	0x274 32  LOWPWR_CTRL_SET  XTAL OSC (LP) Control Register
//	0x278 32  LOWPWR_CTRL_CLR  XTAL OSC (LP) Control Register
//	0x27C 32  LOWPWR_CTRL_TOG  XTAL OSC (LP) Control Register
//	0x2A0 32  OSC_CONFIG0      XTAL OSC Configuration 0 Register
//	0x2A4 32  OSC_CONFIG0_SET  XTAL OSC Configuration 0 Register
//	0x2A8 32  OSC_CONFIG0_CLR  XTAL OSC Configuration 0 Register
//	0x2AC 32  OSC_CONFIG0_TOG  XTAL OSC Configuration 0 Register
//	0x2B0 32  OSC_CONFIG1      XTAL OSC Configuration 1 Register
//	0x2B4 32  OSC_CONFIG1_SET  XTAL OSC Configuration 1 Register
//	0x2B8 32  OSC_CONFIG1_CLR  XTAL OSC Configuration 1 Register
//	0x2BC 32  OSC_CONFIG1_TOG  XTAL OSC Configuration 1 Register
//	0x2C0 32  OSC_CONFIG2      XTAL OSC Configuration 2 Register
//	0x2C4 32  OSC_CONFIG2_SET  XTAL OSC Configuration 2 Register
//	0x2C8 32  OSC_CONFIG2_CLR  XTAL OSC Configuration 2 Register
//	0x2CC 32  OSC_CONFIG2_TOG  XTAL OSC Configuration 2 Register
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package xtaloscm

const (
	REFTOP_PWD         MISC0 = 0x01 << 0  //+ Control bit to power-down the analog bandgap reference circuitry
	REFTOP_SELFBIASOFF MISC0 = 0x01 << 3  //+ Control bit to disable the self-bias circuit in the analog bandgap
	REFTOP_VBGADJ      MISC0 = 0x07 << 4  //+ Not related to oscillator.
	REFTOP_VBGADJ_0    MISC0 = 0x00 << 4  //  Nominal VBG
	REFTOP_VBGADJ_1    MISC0 = 0x01 << 4  //  VBG+0.78%
	REFTOP_VBGADJ_2    MISC0 = 0x02 << 4  //  VBG+1.56%
	REFTOP_VBGADJ_3    MISC0 = 0x03 << 4  //  VBG+2.34%
	REFTOP_VBGADJ_4    MISC0 = 0x04 << 4  //  VBG-0.78%
	REFTOP_VBGADJ_5    MISC0 = 0x05 << 4  //  VBG-1.56%
	REFTOP_VBGADJ_6    MISC0 = 0x06 << 4  //  VBG-2.34%
	REFTOP_VBGADJ_7    MISC0 = 0x07 << 4  //  VBG-3.12%
	REFTOP_VBGUP       MISC0 = 0x01 << 7  //+ Status bit that signals the analog bandgap voltage is up and stable
	STOP_MODE_CONFIG   MISC0 = 0x03 << 10 //+ Configure the analog behavior in stop mode.Not related to oscillator.
	STOP_MODE_CONFIG_0 MISC0 = 0x00 << 10 //  All analog except rtc powered down on stop mode assertion. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_1 MISC0 = 0x01 << 10 //  Certain analog functions such as certain regulators left up. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_2 MISC0 = 0x02 << 10 //  XtalOsc=off, RCOsc=on, Old BG=on, New BG=off.
	STOP_MODE_CONFIG_3 MISC0 = 0x03 << 10 //  XtalOsc=off, RCOsc=on, Old BG=off, New BG=on.
	DISCON_HIGH_SNVS   MISC0 = 0x01 << 12 //+ This bit controls a switch from VDD_HIGH_IN to VDD_SNVS_IN.
	OSC_I              MISC0 = 0x03 << 13 //+ This field determines the bias current in the 24MHz oscillator
	NOMINAL            MISC0 = 0x00 << 13 //  Nominal
	MINUS_12_5_PERCENT MISC0 = 0x01 << 13 //  Decrease current by 12.5%
	MINUS_25_PERCENT   MISC0 = 0x02 << 13 //  Decrease current by 25.0%
	MINUS_37_5_PERCENT MISC0 = 0x03 << 13 //  Decrease current by 37.5%
	OSC_XTALOK         MISC0 = 0x01 << 15 //+ Status bit that signals that the output of the 24-MHz crystal oscillator is stable
	OSC_XTALOK_EN      MISC0 = 0x01 << 16 //+ This bit enables the detector that signals when the 24MHz crystal oscillator is stable.
	CLKGATE_CTRL       MISC0 = 0x01 << 25 //+ This bit allows disabling the clock gate (always ungated) for the xtal 24MHz clock that clocks the digital logic in the analog block
	CLKGATE_DELAY      MISC0 = 0x07 << 26 //+ This field specifies the delay between powering up the XTAL 24MHz clock and releasing the clock to the digital logic inside the analog block
	CLKGATE_DELAY_0    MISC0 = 0x00 << 26 //  0.5ms
	CLKGATE_DELAY_1    MISC0 = 0x01 << 26 //  1.0ms
	CLKGATE_DELAY_2    MISC0 = 0x02 << 26 //  2.0ms
	CLKGATE_DELAY_3    MISC0 = 0x03 << 26 //  3.0ms
	CLKGATE_DELAY_4    MISC0 = 0x04 << 26 //  4.0ms
	CLKGATE_DELAY_5    MISC0 = 0x05 << 26 //  5.0ms
	CLKGATE_DELAY_6    MISC0 = 0x06 << 26 //  6.0ms
	CLKGATE_DELAY_7    MISC0 = 0x07 << 26 //  7.0ms
	RTC_XTAL_SOURCE    MISC0 = 0x01 << 29 //+ This field indicates which chip source is being used for the rtc clock.
	XTAL_24M_PWD       MISC0 = 0x01 << 30 //+ This field powers down the 24M crystal oscillator if set true.
	VID_PLL_PREDIV     MISC0 = 0x01 << 31 //+ Predivider for the source clock of the PLL's. Not related to oscillator.
)

const (
	REFTOP_PWDn         = 0
	REFTOP_SELFBIASOFFn = 3
	REFTOP_VBGADJn      = 4
	REFTOP_VBGUPn       = 7
	STOP_MODE_CONFIGn   = 10
	DISCON_HIGH_SNVSn   = 12
	OSC_In              = 13
	OSC_XTALOKn         = 15
	OSC_XTALOK_ENn      = 16
	CLKGATE_CTRLn       = 25
	CLKGATE_DELAYn      = 26
	RTC_XTAL_SOURCEn    = 29
	XTAL_24M_PWDn       = 30
	VID_PLL_PREDIVn     = 31
)

const (
	REFTOP_PWD         MISC0_SET = 0x01 << 0  //+ Control bit to power-down the analog bandgap reference circuitry
	REFTOP_SELFBIASOFF MISC0_SET = 0x01 << 3  //+ Control bit to disable the self-bias circuit in the analog bandgap
	REFTOP_VBGADJ      MISC0_SET = 0x07 << 4  //+ Not related to oscillator.
	REFTOP_VBGADJ_0    MISC0_SET = 0x00 << 4  //  Nominal VBG
	REFTOP_VBGADJ_1    MISC0_SET = 0x01 << 4  //  VBG+0.78%
	REFTOP_VBGADJ_2    MISC0_SET = 0x02 << 4  //  VBG+1.56%
	REFTOP_VBGADJ_3    MISC0_SET = 0x03 << 4  //  VBG+2.34%
	REFTOP_VBGADJ_4    MISC0_SET = 0x04 << 4  //  VBG-0.78%
	REFTOP_VBGADJ_5    MISC0_SET = 0x05 << 4  //  VBG-1.56%
	REFTOP_VBGADJ_6    MISC0_SET = 0x06 << 4  //  VBG-2.34%
	REFTOP_VBGADJ_7    MISC0_SET = 0x07 << 4  //  VBG-3.12%
	REFTOP_VBGUP       MISC0_SET = 0x01 << 7  //+ Status bit that signals the analog bandgap voltage is up and stable
	STOP_MODE_CONFIG   MISC0_SET = 0x03 << 10 //+ Configure the analog behavior in stop mode.Not related to oscillator.
	STOP_MODE_CONFIG_0 MISC0_SET = 0x00 << 10 //  All analog except rtc powered down on stop mode assertion. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_1 MISC0_SET = 0x01 << 10 //  Certain analog functions such as certain regulators left up. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_2 MISC0_SET = 0x02 << 10 //  XtalOsc=off, RCOsc=on, Old BG=on, New BG=off.
	STOP_MODE_CONFIG_3 MISC0_SET = 0x03 << 10 //  XtalOsc=off, RCOsc=on, Old BG=off, New BG=on.
	DISCON_HIGH_SNVS   MISC0_SET = 0x01 << 12 //+ This bit controls a switch from VDD_HIGH_IN to VDD_SNVS_IN.
	OSC_I              MISC0_SET = 0x03 << 13 //+ This field determines the bias current in the 24MHz oscillator
	NOMINAL            MISC0_SET = 0x00 << 13 //  Nominal
	MINUS_12_5_PERCENT MISC0_SET = 0x01 << 13 //  Decrease current by 12.5%
	MINUS_25_PERCENT   MISC0_SET = 0x02 << 13 //  Decrease current by 25.0%
	MINUS_37_5_PERCENT MISC0_SET = 0x03 << 13 //  Decrease current by 37.5%
	OSC_XTALOK         MISC0_SET = 0x01 << 15 //+ Status bit that signals that the output of the 24-MHz crystal oscillator is stable
	OSC_XTALOK_EN      MISC0_SET = 0x01 << 16 //+ This bit enables the detector that signals when the 24MHz crystal oscillator is stable.
	CLKGATE_CTRL       MISC0_SET = 0x01 << 25 //+ This bit allows disabling the clock gate (always ungated) for the xtal 24MHz clock that clocks the digital logic in the analog block
	CLKGATE_DELAY      MISC0_SET = 0x07 << 26 //+ This field specifies the delay between powering up the XTAL 24MHz clock and releasing the clock to the digital logic inside the analog block
	CLKGATE_DELAY_0    MISC0_SET = 0x00 << 26 //  0.5ms
	CLKGATE_DELAY_1    MISC0_SET = 0x01 << 26 //  1.0ms
	CLKGATE_DELAY_2    MISC0_SET = 0x02 << 26 //  2.0ms
	CLKGATE_DELAY_3    MISC0_SET = 0x03 << 26 //  3.0ms
	CLKGATE_DELAY_4    MISC0_SET = 0x04 << 26 //  4.0ms
	CLKGATE_DELAY_5    MISC0_SET = 0x05 << 26 //  5.0ms
	CLKGATE_DELAY_6    MISC0_SET = 0x06 << 26 //  6.0ms
	CLKGATE_DELAY_7    MISC0_SET = 0x07 << 26 //  7.0ms
	RTC_XTAL_SOURCE    MISC0_SET = 0x01 << 29 //+ This field indicates which chip source is being used for the rtc clock.
	XTAL_24M_PWD       MISC0_SET = 0x01 << 30 //+ This field powers down the 24M crystal oscillator if set true.
	VID_PLL_PREDIV     MISC0_SET = 0x01 << 31 //+ Predivider for the source clock of the PLL's. Not related to oscillator.
)

const (
	REFTOP_PWDn         = 0
	REFTOP_SELFBIASOFFn = 3
	REFTOP_VBGADJn      = 4
	REFTOP_VBGUPn       = 7
	STOP_MODE_CONFIGn   = 10
	DISCON_HIGH_SNVSn   = 12
	OSC_In              = 13
	OSC_XTALOKn         = 15
	OSC_XTALOK_ENn      = 16
	CLKGATE_CTRLn       = 25
	CLKGATE_DELAYn      = 26
	RTC_XTAL_SOURCEn    = 29
	XTAL_24M_PWDn       = 30
	VID_PLL_PREDIVn     = 31
)

const (
	REFTOP_PWD         MISC0_CLR = 0x01 << 0  //+ Control bit to power-down the analog bandgap reference circuitry
	REFTOP_SELFBIASOFF MISC0_CLR = 0x01 << 3  //+ Control bit to disable the self-bias circuit in the analog bandgap
	REFTOP_VBGADJ      MISC0_CLR = 0x07 << 4  //+ Not related to oscillator.
	REFTOP_VBGADJ_0    MISC0_CLR = 0x00 << 4  //  Nominal VBG
	REFTOP_VBGADJ_1    MISC0_CLR = 0x01 << 4  //  VBG+0.78%
	REFTOP_VBGADJ_2    MISC0_CLR = 0x02 << 4  //  VBG+1.56%
	REFTOP_VBGADJ_3    MISC0_CLR = 0x03 << 4  //  VBG+2.34%
	REFTOP_VBGADJ_4    MISC0_CLR = 0x04 << 4  //  VBG-0.78%
	REFTOP_VBGADJ_5    MISC0_CLR = 0x05 << 4  //  VBG-1.56%
	REFTOP_VBGADJ_6    MISC0_CLR = 0x06 << 4  //  VBG-2.34%
	REFTOP_VBGADJ_7    MISC0_CLR = 0x07 << 4  //  VBG-3.12%
	REFTOP_VBGUP       MISC0_CLR = 0x01 << 7  //+ Status bit that signals the analog bandgap voltage is up and stable
	STOP_MODE_CONFIG   MISC0_CLR = 0x03 << 10 //+ Configure the analog behavior in stop mode.Not related to oscillator.
	STOP_MODE_CONFIG_0 MISC0_CLR = 0x00 << 10 //  All analog except rtc powered down on stop mode assertion. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_1 MISC0_CLR = 0x01 << 10 //  Certain analog functions such as certain regulators left up. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_2 MISC0_CLR = 0x02 << 10 //  XtalOsc=off, RCOsc=on, Old BG=on, New BG=off.
	STOP_MODE_CONFIG_3 MISC0_CLR = 0x03 << 10 //  XtalOsc=off, RCOsc=on, Old BG=off, New BG=on.
	DISCON_HIGH_SNVS   MISC0_CLR = 0x01 << 12 //+ This bit controls a switch from VDD_HIGH_IN to VDD_SNVS_IN.
	OSC_I              MISC0_CLR = 0x03 << 13 //+ This field determines the bias current in the 24MHz oscillator
	NOMINAL            MISC0_CLR = 0x00 << 13 //  Nominal
	MINUS_12_5_PERCENT MISC0_CLR = 0x01 << 13 //  Decrease current by 12.5%
	MINUS_25_PERCENT   MISC0_CLR = 0x02 << 13 //  Decrease current by 25.0%
	MINUS_37_5_PERCENT MISC0_CLR = 0x03 << 13 //  Decrease current by 37.5%
	OSC_XTALOK         MISC0_CLR = 0x01 << 15 //+ Status bit that signals that the output of the 24-MHz crystal oscillator is stable
	OSC_XTALOK_EN      MISC0_CLR = 0x01 << 16 //+ This bit enables the detector that signals when the 24MHz crystal oscillator is stable.
	CLKGATE_CTRL       MISC0_CLR = 0x01 << 25 //+ This bit allows disabling the clock gate (always ungated) for the xtal 24MHz clock that clocks the digital logic in the analog block
	CLKGATE_DELAY      MISC0_CLR = 0x07 << 26 //+ This field specifies the delay between powering up the XTAL 24MHz clock and releasing the clock to the digital logic inside the analog block
	CLKGATE_DELAY_0    MISC0_CLR = 0x00 << 26 //  0.5ms
	CLKGATE_DELAY_1    MISC0_CLR = 0x01 << 26 //  1.0ms
	CLKGATE_DELAY_2    MISC0_CLR = 0x02 << 26 //  2.0ms
	CLKGATE_DELAY_3    MISC0_CLR = 0x03 << 26 //  3.0ms
	CLKGATE_DELAY_4    MISC0_CLR = 0x04 << 26 //  4.0ms
	CLKGATE_DELAY_5    MISC0_CLR = 0x05 << 26 //  5.0ms
	CLKGATE_DELAY_6    MISC0_CLR = 0x06 << 26 //  6.0ms
	CLKGATE_DELAY_7    MISC0_CLR = 0x07 << 26 //  7.0ms
	RTC_XTAL_SOURCE    MISC0_CLR = 0x01 << 29 //+ This field indicates which chip source is being used for the rtc clock.
	XTAL_24M_PWD       MISC0_CLR = 0x01 << 30 //+ This field powers down the 24M crystal oscillator if set true.
	VID_PLL_PREDIV     MISC0_CLR = 0x01 << 31 //+ Predivider for the source clock of the PLL's. Not related to oscillator.
)

const (
	REFTOP_PWDn         = 0
	REFTOP_SELFBIASOFFn = 3
	REFTOP_VBGADJn      = 4
	REFTOP_VBGUPn       = 7
	STOP_MODE_CONFIGn   = 10
	DISCON_HIGH_SNVSn   = 12
	OSC_In              = 13
	OSC_XTALOKn         = 15
	OSC_XTALOK_ENn      = 16
	CLKGATE_CTRLn       = 25
	CLKGATE_DELAYn      = 26
	RTC_XTAL_SOURCEn    = 29
	XTAL_24M_PWDn       = 30
	VID_PLL_PREDIVn     = 31
)

const (
	REFTOP_PWD         MISC0_TOG = 0x01 << 0  //+ Control bit to power-down the analog bandgap reference circuitry
	REFTOP_SELFBIASOFF MISC0_TOG = 0x01 << 3  //+ Control bit to disable the self-bias circuit in the analog bandgap
	REFTOP_VBGADJ      MISC0_TOG = 0x07 << 4  //+ Not related to oscillator.
	REFTOP_VBGADJ_0    MISC0_TOG = 0x00 << 4  //  Nominal VBG
	REFTOP_VBGADJ_1    MISC0_TOG = 0x01 << 4  //  VBG+0.78%
	REFTOP_VBGADJ_2    MISC0_TOG = 0x02 << 4  //  VBG+1.56%
	REFTOP_VBGADJ_3    MISC0_TOG = 0x03 << 4  //  VBG+2.34%
	REFTOP_VBGADJ_4    MISC0_TOG = 0x04 << 4  //  VBG-0.78%
	REFTOP_VBGADJ_5    MISC0_TOG = 0x05 << 4  //  VBG-1.56%
	REFTOP_VBGADJ_6    MISC0_TOG = 0x06 << 4  //  VBG-2.34%
	REFTOP_VBGADJ_7    MISC0_TOG = 0x07 << 4  //  VBG-3.12%
	REFTOP_VBGUP       MISC0_TOG = 0x01 << 7  //+ Status bit that signals the analog bandgap voltage is up and stable
	STOP_MODE_CONFIG   MISC0_TOG = 0x03 << 10 //+ Configure the analog behavior in stop mode.Not related to oscillator.
	STOP_MODE_CONFIG_0 MISC0_TOG = 0x00 << 10 //  All analog except rtc powered down on stop mode assertion. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_1 MISC0_TOG = 0x01 << 10 //  Certain analog functions such as certain regulators left up. XtalOsc=on, RCOsc=off;
	STOP_MODE_CONFIG_2 MISC0_TOG = 0x02 << 10 //  XtalOsc=off, RCOsc=on, Old BG=on, New BG=off.
	STOP_MODE_CONFIG_3 MISC0_TOG = 0x03 << 10 //  XtalOsc=off, RCOsc=on, Old BG=off, New BG=on.
	DISCON_HIGH_SNVS   MISC0_TOG = 0x01 << 12 //+ This bit controls a switch from VDD_HIGH_IN to VDD_SNVS_IN.
	OSC_I              MISC0_TOG = 0x03 << 13 //+ This field determines the bias current in the 24MHz oscillator
	NOMINAL            MISC0_TOG = 0x00 << 13 //  Nominal
	MINUS_12_5_PERCENT MISC0_TOG = 0x01 << 13 //  Decrease current by 12.5%
	MINUS_25_PERCENT   MISC0_TOG = 0x02 << 13 //  Decrease current by 25.0%
	MINUS_37_5_PERCENT MISC0_TOG = 0x03 << 13 //  Decrease current by 37.5%
	OSC_XTALOK         MISC0_TOG = 0x01 << 15 //+ Status bit that signals that the output of the 24-MHz crystal oscillator is stable
	OSC_XTALOK_EN      MISC0_TOG = 0x01 << 16 //+ This bit enables the detector that signals when the 24MHz crystal oscillator is stable.
	CLKGATE_CTRL       MISC0_TOG = 0x01 << 25 //+ This bit allows disabling the clock gate (always ungated) for the xtal 24MHz clock that clocks the digital logic in the analog block
	CLKGATE_DELAY      MISC0_TOG = 0x07 << 26 //+ This field specifies the delay between powering up the XTAL 24MHz clock and releasing the clock to the digital logic inside the analog block
	CLKGATE_DELAY_0    MISC0_TOG = 0x00 << 26 //  0.5ms
	CLKGATE_DELAY_1    MISC0_TOG = 0x01 << 26 //  1.0ms
	CLKGATE_DELAY_2    MISC0_TOG = 0x02 << 26 //  2.0ms
	CLKGATE_DELAY_3    MISC0_TOG = 0x03 << 26 //  3.0ms
	CLKGATE_DELAY_4    MISC0_TOG = 0x04 << 26 //  4.0ms
	CLKGATE_DELAY_5    MISC0_TOG = 0x05 << 26 //  5.0ms
	CLKGATE_DELAY_6    MISC0_TOG = 0x06 << 26 //  6.0ms
	CLKGATE_DELAY_7    MISC0_TOG = 0x07 << 26 //  7.0ms
	RTC_XTAL_SOURCE    MISC0_TOG = 0x01 << 29 //+ This field indicates which chip source is being used for the rtc clock.
	XTAL_24M_PWD       MISC0_TOG = 0x01 << 30 //+ This field powers down the 24M crystal oscillator if set true.
	VID_PLL_PREDIV     MISC0_TOG = 0x01 << 31 //+ Predivider for the source clock of the PLL's. Not related to oscillator.
)

const (
	REFTOP_PWDn         = 0
	REFTOP_SELFBIASOFFn = 3
	REFTOP_VBGADJn      = 4
	REFTOP_VBGUPn       = 7
	STOP_MODE_CONFIGn   = 10
	DISCON_HIGH_SNVSn   = 12
	OSC_In              = 13
	OSC_XTALOKn         = 15
	OSC_XTALOK_ENn      = 16
	CLKGATE_CTRLn       = 25
	CLKGATE_DELAYn      = 26
	RTC_XTAL_SOURCEn    = 29
	XTAL_24M_PWDn       = 30
	VID_PLL_PREDIVn     = 31
)

const (
	RC_OSC_EN             LOWPWR_CTRL = 0x01 << 0  //+ RC Osc. enable control.
	OSC_SEL               LOWPWR_CTRL = 0x01 << 4  //+ Select the source for the 24MHz clock.
	LPBG_SEL              LOWPWR_CTRL = 0x01 << 5  //+ Bandgap select. Not related to oscillator.
	LPBG_TEST             LOWPWR_CTRL = 0x01 << 6  //+ Low power bandgap test bit. Not related to oscillator.
	REFTOP_IBIAS_OFF      LOWPWR_CTRL = 0x01 << 7  //+ Low power reftop ibias disable. Not related to oscillator.
	L1_PWRGATE            LOWPWR_CTRL = 0x01 << 8  //+ L1 power gate control. Used as software override. Not related to oscillator.
	L2_PWRGATE            LOWPWR_CTRL = 0x01 << 9  //+ L2 power gate control. Used as software override. Not related to oscillator.
	CPU_PWRGATE           LOWPWR_CTRL = 0x01 << 10 //+ CPU power gate control. Used as software override. Test purpose only Not related to oscillator.
	DISPLAY_PWRGATE       LOWPWR_CTRL = 0x01 << 11 //+ Display logic power gate control. Used as software override. Not related to oscillator.
	RCOSC_CG_OVERRIDE     LOWPWR_CTRL = 0x01 << 13 //+ For debug purposes only
	XTALOSC_PWRUP_DELAY   LOWPWR_CTRL = 0x03 << 14 //+ Specifies the time delay between when the 24MHz xtal is powered up until it is stable and ready to use
	XTALOSC_PWRUP_DELAY_0 LOWPWR_CTRL = 0x00 << 14 //  0.25ms
	XTALOSC_PWRUP_DELAY_1 LOWPWR_CTRL = 0x01 << 14 //  0.5ms
	XTALOSC_PWRUP_DELAY_2 LOWPWR_CTRL = 0x02 << 14 //  1ms
	XTALOSC_PWRUP_DELAY_3 LOWPWR_CTRL = 0x03 << 14 //  2ms
	XTALOSC_PWRUP_STAT    LOWPWR_CTRL = 0x01 << 16 //+ Status of the 24MHz xtal oscillator.
	MIX_PWRGATE           LOWPWR_CTRL = 0x01 << 17 //+ Display power gate control. Used as software mask. Set to zero to force ungated.
	GPU_PWRGATE           LOWPWR_CTRL = 0x01 << 18 //+ GPU power gate control. Used as software mask. Set to zero to force ungated.
)

const (
	RC_OSC_ENn           = 0
	OSC_SELn             = 4
	LPBG_SELn            = 5
	LPBG_TESTn           = 6
	REFTOP_IBIAS_OFFn    = 7
	L1_PWRGATEn          = 8
	L2_PWRGATEn          = 9
	CPU_PWRGATEn         = 10
	DISPLAY_PWRGATEn     = 11
	RCOSC_CG_OVERRIDEn   = 13
	XTALOSC_PWRUP_DELAYn = 14
	XTALOSC_PWRUP_STATn  = 16
	MIX_PWRGATEn         = 17
	GPU_PWRGATEn         = 18
)

const (
	RC_OSC_EN             LOWPWR_CTRL_SET = 0x01 << 0  //+ RC Osc. enable control.
	OSC_SEL               LOWPWR_CTRL_SET = 0x01 << 4  //+ Select the source for the 24MHz clock.
	LPBG_SEL              LOWPWR_CTRL_SET = 0x01 << 5  //+ Bandgap select. Not related to oscillator.
	LPBG_TEST             LOWPWR_CTRL_SET = 0x01 << 6  //+ Low power bandgap test bit. Not related to oscillator.
	REFTOP_IBIAS_OFF      LOWPWR_CTRL_SET = 0x01 << 7  //+ Low power reftop ibias disable. Not related to oscillator.
	L1_PWRGATE            LOWPWR_CTRL_SET = 0x01 << 8  //+ L1 power gate control. Used as software override. Not related to oscillator.
	L2_PWRGATE            LOWPWR_CTRL_SET = 0x01 << 9  //+ L2 power gate control. Used as software override. Not related to oscillator.
	CPU_PWRGATE           LOWPWR_CTRL_SET = 0x01 << 10 //+ CPU power gate control. Used as software override. Test purpose only Not related to oscillator.
	DISPLAY_PWRGATE       LOWPWR_CTRL_SET = 0x01 << 11 //+ Display logic power gate control. Used as software override. Not related to oscillator.
	RCOSC_CG_OVERRIDE     LOWPWR_CTRL_SET = 0x01 << 13 //+ For debug purposes only
	XTALOSC_PWRUP_DELAY   LOWPWR_CTRL_SET = 0x03 << 14 //+ Specifies the time delay between when the 24MHz xtal is powered up until it is stable and ready to use
	XTALOSC_PWRUP_DELAY_0 LOWPWR_CTRL_SET = 0x00 << 14 //  0.25ms
	XTALOSC_PWRUP_DELAY_1 LOWPWR_CTRL_SET = 0x01 << 14 //  0.5ms
	XTALOSC_PWRUP_DELAY_2 LOWPWR_CTRL_SET = 0x02 << 14 //  1ms
	XTALOSC_PWRUP_DELAY_3 LOWPWR_CTRL_SET = 0x03 << 14 //  2ms
	XTALOSC_PWRUP_STAT    LOWPWR_CTRL_SET = 0x01 << 16 //+ Status of the 24MHz xtal oscillator.
	MIX_PWRGATE           LOWPWR_CTRL_SET = 0x01 << 17 //+ Display power gate control. Used as software mask. Set to zero to force ungated.
	GPU_PWRGATE           LOWPWR_CTRL_SET = 0x01 << 18 //+ GPU power gate control. Used as software mask. Set to zero to force ungated.
)

const (
	RC_OSC_ENn           = 0
	OSC_SELn             = 4
	LPBG_SELn            = 5
	LPBG_TESTn           = 6
	REFTOP_IBIAS_OFFn    = 7
	L1_PWRGATEn          = 8
	L2_PWRGATEn          = 9
	CPU_PWRGATEn         = 10
	DISPLAY_PWRGATEn     = 11
	RCOSC_CG_OVERRIDEn   = 13
	XTALOSC_PWRUP_DELAYn = 14
	XTALOSC_PWRUP_STATn  = 16
	MIX_PWRGATEn         = 17
	GPU_PWRGATEn         = 18
)

const (
	RC_OSC_EN             LOWPWR_CTRL_CLR = 0x01 << 0  //+ RC Osc. enable control.
	OSC_SEL               LOWPWR_CTRL_CLR = 0x01 << 4  //+ Select the source for the 24MHz clock.
	LPBG_SEL              LOWPWR_CTRL_CLR = 0x01 << 5  //+ Bandgap select. Not related to oscillator.
	LPBG_TEST             LOWPWR_CTRL_CLR = 0x01 << 6  //+ Low power bandgap test bit. Not related to oscillator.
	REFTOP_IBIAS_OFF      LOWPWR_CTRL_CLR = 0x01 << 7  //+ Low power reftop ibias disable. Not related to oscillator.
	L1_PWRGATE            LOWPWR_CTRL_CLR = 0x01 << 8  //+ L1 power gate control. Used as software override. Not related to oscillator.
	L2_PWRGATE            LOWPWR_CTRL_CLR = 0x01 << 9  //+ L2 power gate control. Used as software override. Not related to oscillator.
	CPU_PWRGATE           LOWPWR_CTRL_CLR = 0x01 << 10 //+ CPU power gate control. Used as software override. Test purpose only Not related to oscillator.
	DISPLAY_PWRGATE       LOWPWR_CTRL_CLR = 0x01 << 11 //+ Display logic power gate control. Used as software override. Not related to oscillator.
	RCOSC_CG_OVERRIDE     LOWPWR_CTRL_CLR = 0x01 << 13 //+ For debug purposes only
	XTALOSC_PWRUP_DELAY   LOWPWR_CTRL_CLR = 0x03 << 14 //+ Specifies the time delay between when the 24MHz xtal is powered up until it is stable and ready to use
	XTALOSC_PWRUP_DELAY_0 LOWPWR_CTRL_CLR = 0x00 << 14 //  0.25ms
	XTALOSC_PWRUP_DELAY_1 LOWPWR_CTRL_CLR = 0x01 << 14 //  0.5ms
	XTALOSC_PWRUP_DELAY_2 LOWPWR_CTRL_CLR = 0x02 << 14 //  1ms
	XTALOSC_PWRUP_DELAY_3 LOWPWR_CTRL_CLR = 0x03 << 14 //  2ms
	XTALOSC_PWRUP_STAT    LOWPWR_CTRL_CLR = 0x01 << 16 //+ Status of the 24MHz xtal oscillator.
	MIX_PWRGATE           LOWPWR_CTRL_CLR = 0x01 << 17 //+ Display power gate control. Used as software mask. Set to zero to force ungated.
	GPU_PWRGATE           LOWPWR_CTRL_CLR = 0x01 << 18 //+ GPU power gate control. Used as software mask. Set to zero to force ungated.
)

const (
	RC_OSC_ENn           = 0
	OSC_SELn             = 4
	LPBG_SELn            = 5
	LPBG_TESTn           = 6
	REFTOP_IBIAS_OFFn    = 7
	L1_PWRGATEn          = 8
	L2_PWRGATEn          = 9
	CPU_PWRGATEn         = 10
	DISPLAY_PWRGATEn     = 11
	RCOSC_CG_OVERRIDEn   = 13
	XTALOSC_PWRUP_DELAYn = 14
	XTALOSC_PWRUP_STATn  = 16
	MIX_PWRGATEn         = 17
	GPU_PWRGATEn         = 18
)

const (
	RC_OSC_EN             LOWPWR_CTRL_TOG = 0x01 << 0  //+ RC Osc. enable control.
	OSC_SEL               LOWPWR_CTRL_TOG = 0x01 << 4  //+ Select the source for the 24MHz clock.
	LPBG_SEL              LOWPWR_CTRL_TOG = 0x01 << 5  //+ Bandgap select. Not related to oscillator.
	LPBG_TEST             LOWPWR_CTRL_TOG = 0x01 << 6  //+ Low power bandgap test bit. Not related to oscillator.
	REFTOP_IBIAS_OFF      LOWPWR_CTRL_TOG = 0x01 << 7  //+ Low power reftop ibias disable. Not related to oscillator.
	L1_PWRGATE            LOWPWR_CTRL_TOG = 0x01 << 8  //+ L1 power gate control. Used as software override. Not related to oscillator.
	L2_PWRGATE            LOWPWR_CTRL_TOG = 0x01 << 9  //+ L2 power gate control. Used as software override. Not related to oscillator.
	CPU_PWRGATE           LOWPWR_CTRL_TOG = 0x01 << 10 //+ CPU power gate control. Used as software override. Test purpose only Not related to oscillator.
	DISPLAY_PWRGATE       LOWPWR_CTRL_TOG = 0x01 << 11 //+ Display logic power gate control. Used as software override. Not related to oscillator.
	RCOSC_CG_OVERRIDE     LOWPWR_CTRL_TOG = 0x01 << 13 //+ For debug purposes only
	XTALOSC_PWRUP_DELAY   LOWPWR_CTRL_TOG = 0x03 << 14 //+ Specifies the time delay between when the 24MHz xtal is powered up until it is stable and ready to use
	XTALOSC_PWRUP_DELAY_0 LOWPWR_CTRL_TOG = 0x00 << 14 //  0.25ms
	XTALOSC_PWRUP_DELAY_1 LOWPWR_CTRL_TOG = 0x01 << 14 //  0.5ms
	XTALOSC_PWRUP_DELAY_2 LOWPWR_CTRL_TOG = 0x02 << 14 //  1ms
	XTALOSC_PWRUP_DELAY_3 LOWPWR_CTRL_TOG = 0x03 << 14 //  2ms
	XTALOSC_PWRUP_STAT    LOWPWR_CTRL_TOG = 0x01 << 16 //+ Status of the 24MHz xtal oscillator.
	MIX_PWRGATE           LOWPWR_CTRL_TOG = 0x01 << 17 //+ Display power gate control. Used as software mask. Set to zero to force ungated.
	GPU_PWRGATE           LOWPWR_CTRL_TOG = 0x01 << 18 //+ GPU power gate control. Used as software mask. Set to zero to force ungated.
)

const (
	RC_OSC_ENn           = 0
	OSC_SELn             = 4
	LPBG_SELn            = 5
	LPBG_TESTn           = 6
	REFTOP_IBIAS_OFFn    = 7
	L1_PWRGATEn          = 8
	L2_PWRGATEn          = 9
	CPU_PWRGATEn         = 10
	DISPLAY_PWRGATEn     = 11
	RCOSC_CG_OVERRIDEn   = 13
	XTALOSC_PWRUP_DELAYn = 14
	XTALOSC_PWRUP_STATn  = 16
	MIX_PWRGATEn         = 17
	GPU_PWRGATEn         = 18
)

const (
	START           OSC_CONFIG0 = 0x01 << 0  //+ Start/stop bit for the RC tuning calculation logic. If stopped the tuning logic is reset.
	ENABLE          OSC_CONFIG0 = 0x01 << 1  //+ Enables the tuning logic to calculate new RC tuning values
	BYPASS          OSC_CONFIG0 = 0x01 << 2  //+ Bypasses any calculated RC tuning value and uses the programmed register value.
	INVERT          OSC_CONFIG0 = 0x01 << 3  //+ Invert the stepping of the calculated RC tuning value.
	RC_OSC_PROG     OSC_CONFIG0 = 0xFF << 4  //+ RC osc. tuning values.
	HYST_PLUS       OSC_CONFIG0 = 0x0F << 12 //+ Positive hysteresis value
	HYST_MINUS      OSC_CONFIG0 = 0x0F << 16 //+ Negative hysteresis value
	RC_OSC_PROG_CUR OSC_CONFIG0 = 0xFF << 24 //+ The current tuning value in use.
)

const (
	STARTn           = 0
	ENABLEn          = 1
	BYPASSn          = 2
	INVERTn          = 3
	RC_OSC_PROGn     = 4
	HYST_PLUSn       = 12
	HYST_MINUSn      = 16
	RC_OSC_PROG_CURn = 24
)

const (
	START           OSC_CONFIG0_SET = 0x01 << 0  //+ Start/stop bit for the RC tuning calculation logic. If stopped the tuning logic is reset.
	ENABLE          OSC_CONFIG0_SET = 0x01 << 1  //+ Enables the tuning logic to calculate new RC tuning values
	BYPASS          OSC_CONFIG0_SET = 0x01 << 2  //+ Bypasses any calculated RC tuning value and uses the programmed register value.
	INVERT          OSC_CONFIG0_SET = 0x01 << 3  //+ Invert the stepping of the calculated RC tuning value.
	RC_OSC_PROG     OSC_CONFIG0_SET = 0xFF << 4  //+ RC osc. tuning values.
	HYST_PLUS       OSC_CONFIG0_SET = 0x0F << 12 //+ Positive hysteresis value
	HYST_MINUS      OSC_CONFIG0_SET = 0x0F << 16 //+ Negative hysteresis value
	RC_OSC_PROG_CUR OSC_CONFIG0_SET = 0xFF << 24 //+ The current tuning value in use.
)

const (
	STARTn           = 0
	ENABLEn          = 1
	BYPASSn          = 2
	INVERTn          = 3
	RC_OSC_PROGn     = 4
	HYST_PLUSn       = 12
	HYST_MINUSn      = 16
	RC_OSC_PROG_CURn = 24
)

const (
	START           OSC_CONFIG0_CLR = 0x01 << 0  //+ Start/stop bit for the RC tuning calculation logic. If stopped the tuning logic is reset.
	ENABLE          OSC_CONFIG0_CLR = 0x01 << 1  //+ Enables the tuning logic to calculate new RC tuning values
	BYPASS          OSC_CONFIG0_CLR = 0x01 << 2  //+ Bypasses any calculated RC tuning value and uses the programmed register value.
	INVERT          OSC_CONFIG0_CLR = 0x01 << 3  //+ Invert the stepping of the calculated RC tuning value.
	RC_OSC_PROG     OSC_CONFIG0_CLR = 0xFF << 4  //+ RC osc. tuning values.
	HYST_PLUS       OSC_CONFIG0_CLR = 0x0F << 12 //+ Positive hysteresis value
	HYST_MINUS      OSC_CONFIG0_CLR = 0x0F << 16 //+ Negative hysteresis value
	RC_OSC_PROG_CUR OSC_CONFIG0_CLR = 0xFF << 24 //+ The current tuning value in use.
)

const (
	STARTn           = 0
	ENABLEn          = 1
	BYPASSn          = 2
	INVERTn          = 3
	RC_OSC_PROGn     = 4
	HYST_PLUSn       = 12
	HYST_MINUSn      = 16
	RC_OSC_PROG_CURn = 24
)

const (
	START           OSC_CONFIG0_TOG = 0x01 << 0  //+ Start/stop bit for the RC tuning calculation logic. If stopped the tuning logic is reset.
	ENABLE          OSC_CONFIG0_TOG = 0x01 << 1  //+ Enables the tuning logic to calculate new RC tuning values
	BYPASS          OSC_CONFIG0_TOG = 0x01 << 2  //+ Bypasses any calculated RC tuning value and uses the programmed register value.
	INVERT          OSC_CONFIG0_TOG = 0x01 << 3  //+ Invert the stepping of the calculated RC tuning value.
	RC_OSC_PROG     OSC_CONFIG0_TOG = 0xFF << 4  //+ RC osc. tuning values.
	HYST_PLUS       OSC_CONFIG0_TOG = 0x0F << 12 //+ Positive hysteresis value
	HYST_MINUS      OSC_CONFIG0_TOG = 0x0F << 16 //+ Negative hysteresis value
	RC_OSC_PROG_CUR OSC_CONFIG0_TOG = 0xFF << 24 //+ The current tuning value in use.
)

const (
	STARTn           = 0
	ENABLEn          = 1
	BYPASSn          = 2
	INVERTn          = 3
	RC_OSC_PROGn     = 4
	HYST_PLUSn       = 12
	HYST_MINUSn      = 16
	RC_OSC_PROG_CURn = 24
)

const (
	COUNT_RC_TRG OSC_CONFIG1 = 0xFFF << 0  //+ The target count used to tune the RC OSC frequency
	COUNT_RC_CUR OSC_CONFIG1 = 0xFFF << 20 //+ The current tuning value in use.
)

const (
	COUNT_RC_TRGn = 0
	COUNT_RC_CURn = 20
)

const (
	COUNT_RC_TRG OSC_CONFIG1_SET = 0xFFF << 0  //+ The target count used to tune the RC OSC frequency
	COUNT_RC_CUR OSC_CONFIG1_SET = 0xFFF << 20 //+ The current tuning value in use.
)

const (
	COUNT_RC_TRGn = 0
	COUNT_RC_CURn = 20
)

const (
	COUNT_RC_TRG OSC_CONFIG1_CLR = 0xFFF << 0  //+ The target count used to tune the RC OSC frequency
	COUNT_RC_CUR OSC_CONFIG1_CLR = 0xFFF << 20 //+ The current tuning value in use.
)

const (
	COUNT_RC_TRGn = 0
	COUNT_RC_CURn = 20
)

const (
	COUNT_RC_TRG OSC_CONFIG1_TOG = 0xFFF << 0  //+ The target count used to tune the RC OSC frequency
	COUNT_RC_CUR OSC_CONFIG1_TOG = 0xFFF << 20 //+ The current tuning value in use.
)

const (
	COUNT_RC_TRGn = 0
	COUNT_RC_CURn = 20
)

const (
	COUNT_1M_TRG  OSC_CONFIG2 = 0xFFF << 0 //+ The target count used to tune the 1MHz clock frequency
	ENABLE_1M     OSC_CONFIG2 = 0x01 << 16 //+ Enable the 1MHz clock output. 0 - disabled; 1 - enabled.
	MUX_1M        OSC_CONFIG2 = 0x01 << 17 //+ Mux the corrected or uncorrected 1MHz clock to the output
	CLK_1M_ERR_FL OSC_CONFIG2 = 0x01 << 31 //+ Flag indicates that the count_1m count wasn't reached within 1 32kHz period
)

const (
	COUNT_1M_TRGn  = 0
	ENABLE_1Mn     = 16
	MUX_1Mn        = 17
	CLK_1M_ERR_FLn = 31
)

const (
	COUNT_1M_TRG  OSC_CONFIG2_SET = 0xFFF << 0 //+ The target count used to tune the 1MHz clock frequency
	ENABLE_1M     OSC_CONFIG2_SET = 0x01 << 16 //+ Enable the 1MHz clock output. 0 - disabled; 1 - enabled.
	MUX_1M        OSC_CONFIG2_SET = 0x01 << 17 //+ Mux the corrected or uncorrected 1MHz clock to the output
	CLK_1M_ERR_FL OSC_CONFIG2_SET = 0x01 << 31 //+ Flag indicates that the count_1m count wasn't reached within 1 32kHz period
)

const (
	COUNT_1M_TRGn  = 0
	ENABLE_1Mn     = 16
	MUX_1Mn        = 17
	CLK_1M_ERR_FLn = 31
)

const (
	COUNT_1M_TRG  OSC_CONFIG2_CLR = 0xFFF << 0 //+ The target count used to tune the 1MHz clock frequency
	ENABLE_1M     OSC_CONFIG2_CLR = 0x01 << 16 //+ Enable the 1MHz clock output. 0 - disabled; 1 - enabled.
	MUX_1M        OSC_CONFIG2_CLR = 0x01 << 17 //+ Mux the corrected or uncorrected 1MHz clock to the output
	CLK_1M_ERR_FL OSC_CONFIG2_CLR = 0x01 << 31 //+ Flag indicates that the count_1m count wasn't reached within 1 32kHz period
)

const (
	COUNT_1M_TRGn  = 0
	ENABLE_1Mn     = 16
	MUX_1Mn        = 17
	CLK_1M_ERR_FLn = 31
)

const (
	COUNT_1M_TRG  OSC_CONFIG2_TOG = 0xFFF << 0 //+ The target count used to tune the 1MHz clock frequency
	ENABLE_1M     OSC_CONFIG2_TOG = 0x01 << 16 //+ Enable the 1MHz clock output. 0 - disabled; 1 - enabled.
	MUX_1M        OSC_CONFIG2_TOG = 0x01 << 17 //+ Mux the corrected or uncorrected 1MHz clock to the output
	CLK_1M_ERR_FL OSC_CONFIG2_TOG = 0x01 << 31 //+ Flag indicates that the count_1m count wasn't reached within 1 32kHz period
)

const (
	COUNT_1M_TRGn  = 0
	ENABLE_1Mn     = 16
	MUX_1Mn        = 17
	CLK_1M_ERR_FLn = 31
)
