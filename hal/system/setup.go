// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build imxrt1060

package system

import (
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/wdog"
)

// FlexSPI root clock is set to 396 MHz / SEMC_PODF / podf
func Setup528_FlexSPI(podf int) {
	runtime.LockOSThread()
	privLevel, _ := rtos.SetPrivLevel(0)

	// The clock configuration left by bootloader deviates significantly from
	// the default configuration you can see in IMXRT1060RM_rev3 fig. 14-2.
	//
	// CCMA.PLL_ARM = 0x80002042: pll1=24MHz*66/2=792MHz
	// CCMA.PLL_SYS = 0x80002001: pll2=24MHz*22=528MHz
	// CCMA.PFD_528 = 0x18131818: pll2pfd3=pll2/24,pll2pfd2=pll2/19,
	//                            pll2pfd1=pll2/24,pll2pfd0=pll2/24
	// CCMA.PLL_USB1= 0x80003000: pll3=24*20=480MHz
	// CCMA.PFD_480 = 0x0F1A230D: pll3pfd3=pll3/15,pll3pfd2=pll3/26,
	//                            pll3pfd1=pll3/35,pll3pfd0=pll3/13
	//
	// CCM.CACRR = 0x00000001: pll1_=pll1/2
	// CCM.CBCMR = 0xF5AE8104: prePeriph<-pll1_
	// CCM.CBCDR = 0x000A8200: periph<-prePeriph, ahb=periph/1=396MHz
	// CCM.CSCMR1= 0x66130001: flexSPI_<-pll3pfd0, flexSPI=flexSPI_/5=7.4MHz

	CCMA := ccm_analog.CCM_ANALOG()
	CCM := ccm.CCM()

	// Use PLL_SYS 528 MHz as ARM Core clock, disable PLL_ARM.
	CCM.CBCMR.StoreBits(ccm.PRE_PERIPH_CLK_SEL, ccm.PRE_PERIPH_CLK_SEL_0)
	CCMA.PLL_ARM_SET.Store(ccm_analog.PLL_ARM_POWERDOWN)

	// Configure the remaining clocks in a way that resembles the default
	// configuration in IMXRT1060RM_rev3 figure 14-2.

	gateAll := ccm_analog.PFD0_CLKGATE | ccm_analog.PFD1_CLKGATE |
		ccm_analog.PFD2_CLKGATE | ccm_analog.PFD3_CLKGATE

	// Restore PFD_528 default dividers
	CCMA.PFD_528_SET.Store(gateAll)
	CCMA.PFD_528.Store(0 |
		27<<ccm_analog.PFD0_FRACn | // 528 MHz * 18 / 27 = 352 MHz
		16<<ccm_analog.PFD1_FRACn | // 528 MHz * 18 / 16 = 594 MHz
		24<<ccm_analog.PFD2_FRACn | // 528 MHz * 18 / 24 = 396 MHz
		32<<ccm_analog.PFD3_FRACn, //  528 MHz * 18 / 32 = 297 MHz
	)

	// Clock FlexSPI from AXI/SEMC clock: PFD_528_PFD2 / 3 / 2 = 66 MHz
	CCM.CSCMR1.StoreBits(
		ccm.FLEXSPI_CLK_SEL|ccm.FLEXSPI_PODF,
		ccm.FLEXSPI_CLK_SEL_0|ccm.CSCMR1(podf-1),
	)

	// After moving FlexSPI to PFD_528 we can restore the PFD_480 defaults.
	CCMA.PFD_480_SET.Store(gateAll)
	CCMA.PFD_480.Store(0 |
		12<<ccm_analog.PFD0_FRACn | // 480 MHz * 18 / 12 = 720 MHz
		13<<ccm_analog.PFD1_FRACn | // 480 MHz * 18 / 13 = 665 MHz
		17<<ccm_analog.PFD2_FRACn | // 480 MHz * 18 / 17 = 508 MHz
		19<<ccm_analog.PFD3_FRACn, //  480 MHz * 18 / 19 = 455 MHz
	)

	// Use OSC_CLK = 24 MHz clock source for GPT and PIT timers.
	CCM.CSCMR1.StoreBits(
		ccm.PERCLK_PODF|ccm.PERCLK_CLK_SEL,
		ccm.PERCLK_PODF_1|ccm.PERCLK_CLK_SEL,
	)
	// Use OSC_CLK = 24 MHz clock source for UART.
	CCM.CSCDR1.StoreBits(
		ccm.UART_CLK_PODF|ccm.UART_CLK_SEL,
		ccm.UART_CLK_PODF_1|ccm.UART_CLK_SEL,
	)

	// Set REFTOP_SELFBIASOFF after analog bandgap stabilized for best noise
	// performance of analog blocks.
	CCMA.MISC0_SET.Store(ccm_analog.MISC0_REFTOP_SELFBIASOFF)

	// The Watchdog Power Down Counter must be disabled within 16 seconds of
	// reset deassertion.
	wdog.WDOG1().WMCR.ClearBits(wdog.PDE)
	wdog.WDOG2().WMCR.ClearBits(wdog.PDE)

	rtos.SetPrivLevel(privLevel)
	runtime.UnlockOSThread()
}
