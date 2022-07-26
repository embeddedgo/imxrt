// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Work in progress..
package main

import (
	"embedded/arch/cortexm/systim"
	"embedded/rtos"
	"runtime"
	"time"

	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/gpio"
	"github.com/embeddedgo/imxrt/p/iomuxc"
	"github.com/embeddedgo/imxrt/p/iomuxc_gpr"
)

func main() {
	system.Setup528_FlexSPI(2)

	runtime.LockOSThread()
	privLevel, _ := rtos.SetPrivLevel(0)

	// Use SYSTICK as a system timer
	//
	// In case of i.MX RT the default response for WFE/WFI instruction is to
	// enter the Wait mode. In this mode the whole Cortex-M7 core including
	// NVIC is frozen. The system can be awakened by the GPC Interrupt
	// Controller but the SYSTICK interrupt is CM7 internal signal so it isn't
	// routed to GPC.
	//
	// All this means that SYSTICK is almost useless as a system timer, but
	// we'll use it anyway, for educational purposes, preventing entering Wait
	// mode after WFE/WFI.
	ccm.CCM().CLPCR.StoreBits(ccm.LPM, ccm.LPM_RUN)
	systim.Setup(2e6, 100e3, true)
	rtos.SetSystemTimer(systim.Nanotime, nil)

	// GPIO_MUX1 ALT5 selects GPIO6 (fast) instead of GPIO1 (slow) for bit 9.
	iomuxc_gpr.IOMUXC_GPR().GPR26.Store(1 << 9)

	IOMUXC := iomuxc.IOMUXC()

	// Connect pad AD_B0_09 to GPIO 1 or 6 (ALT5 iomux mode)
	IOMUXC.SW_MUX_CTL_PAD_GPIO_AD_B0_09.Store(iomuxc.ALT5)

	// Configure pad AD_B0_09: hysteresis:off, 100KΩ pull-down, pull/keeper:off,
	// open-drain:off, speed:low (50 MHz), drive-strength:(150/7)Ω, sr:slow
	IOMUXC.SW_PAD_CTL_PAD_GPIO_AD_B0_09.Store(iomuxc.DSE_7_R0_7)

	rtos.SetPrivLevel(privLevel)
	runtime.UnlockOSThread()

	GPIO6 := gpio.GPIO6()

	// Set GPIO6 bit 9-th to the output mode.
	GPIO6.GDIR.SetBit(9)
	for {
		GPIO6.DR_CLEAR.Store(1 << 9)
		time.Sleep(100 * time.Millisecond)
		GPIO6.DR_SET.Store(1 << 9)
		time.Sleep(900 * time.Millisecond)
	}
}
