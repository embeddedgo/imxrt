// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Work in progress..
package main

import (
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/gpio"
	"github.com/embeddedgo/imxrt/p/iomuxc"
	"github.com/embeddedgo/imxrt/p/iomuxc_gpr"
)

func main() {
	runtime.LockOSThread()
	privLevel, _ := rtos.SetPrivLevel(0)

	CCMA := ccm_analog.CCM_ANALOG()

	// Set REFTOP_SELFBIASOFF after analog bandgap stabilized for best noise
	// performance of analog blocks.
	CCMA.MISC0_SET.Store(ccm_analog.MISC0_REFTOP_SELFBIASOFF)

	// Setup PLL2
	CCMA.PFD_528_SET.Store(ccm_analog.PFD0_CLKGATE | ccm_analog.PFD1_CLKGATE |
		ccm_analog.PFD2_CLKGATE | ccm_analog.PFD3_CLKGATE)
	CCMA.PFD_528.Store(0x2018101B) // 352,594,396,297 MHz

	// Setup PLL3
	CCMA.PFD_480_SET.Store(ccm_analog.PFD0_CLKGATE | ccm_analog.PFD1_CLKGATE |
		ccm_analog.PFD2_CLKGATE | ccm_analog.PFD3_CLKGATE)
	CCMA.PFD_480.Store(0x13110D0C) // PFD0,1,2,3: 720,664,508,454 MHz

	CCM := ccm.CCM()

	// Configure clocks
	CCM.CSCMR1.StoreBits(
		ccm.PERCLK_PODF|ccm.PERCLK_CLK_SEL,
		ccm.PERCLK_PODF_1|ccm.PERCLK_CLK_SEL,
	)
	CCM.CSCDR1.StoreBits(
		ccm.UART_CLK_PODF|ccm.UART_CLK_SEL,
		ccm.UART_CLK_PODF_1|ccm.UART_CLK_SEL,
	)

	IOMUXC_GPR := iomuxc_gpr.IOMUXC_GPR()

	// GPIO_MUX1 ALT5 selects GPIO6 (fast) instead of GPIO1 (slow) for all bits.
	IOMUXC_GPR.GPR26.Store(0xFFFFFFFF)

	IOMUXC := iomuxc.IOMUXC()

	// Connect pad AD_B0_09 to GPIO 1 or 6 (ALT5 iomux mode)
	IOMUXC.SW_MUX_CTL_PAD_GPIO_AD_B0_09.Store(iomuxc.ALT5)

	// Configure pad AD_B0_09: hysteresis:off, 100KΩ pull-down, pull/keeper:off,
	// open-drain:off, speed:low (50 MHz), drive-strength:(150/7)Ω, sr:slow
	IOMUXC.SW_PAD_CTL_PAD_GPIO_AD_B0_09.Store(iomuxc.DSE_7_R0_7)

	rtos.SetPrivLevel(privLevel)

	GPIO6 := gpio.GPIO6()

	// Set GPIO6 bit 9-th to the output mode.
	GPIO6.GDIR.SetBit(9)
	for {
		for i := 0; i < 2e5; i++ {
			GPIO6.DR_CLEAR.Store(1 << 9)
		}
		for i := 0; i < 4e6; i++ {
			GPIO6.DR_SET.Store(1 << 9)
		}
	}
}
