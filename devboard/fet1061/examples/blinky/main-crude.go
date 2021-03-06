// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"unsafe"

	"github.com/embeddedgo/imxrt/p/gpio"
)

func mmio32(addr uintptr) *mmio.U32 {
	return (*mmio.U32)(unsafe.Pointer(addr))
}

const (
	IOMUXC_GPR_ADDR uintptr = 0x400AC000
	CCM_ANALOG_ADDR uintptr = 0x400D8000
	CCM_ADDR        uintptr = 0x400FC000
	IOMUXC_ADDR     uintptr = 0x401F8000
	GPIO6_ADDR      uintptr = 0x42000000
)

type GPIO struct {
	DR        mmio.U32
	GDIR      mmio.U32
	PSR       mmio.U32
	ICR1      mmio.U32
	ICR2      mmio.U32
	IMR       mmio.U32
	ISR       mmio.U32
	EDGE_SEL  mmio.U32
	_         [25]mmio.U32
	DR_SET    mmio.U32
	DR_CLEAR  mmio.U32
	DR_TOGGLE mmio.U32
}

var (
	GPIO6 = (*GPIO)(unsafe.Pointer(GPIO6_ADDR))
	stop  = true
)

func main() {
	runtime.LockOSThread()
	privLevel, _ := rtos.SetPrivLevel(0)

	//CCM_ANALOG_PFD_480 := mmio32(CCM_ANALOG_ADDR + 0x0F0)
	//CCM_ANALOG_PFD_480_SET := mmio32(CCM_ANALOG_ADDR + 0x0F0 + 4)
	CCM_ANALOG_PFD_528 := mmio32(CCM_ANALOG_ADDR + 0x100)
	CCM_ANALOG_PFD_528_SET := mmio32(CCM_ANALOG_ADDR + 0x100 + 4)
	PMU_MISC0_SET := mmio32(CCM_ANALOG_ADDR + 0x150 + 4)
	CCM_CSCMR1 := mmio32(CCM_ADDR + 0x01C)
	CCM_CSCDR1 := mmio32(CCM_ADDR + 0x024)
	IOMUXC_SW_MUX_CTL_PAD_GPIO_AD_B0_09 := mmio32(IOMUXC_ADDR + 0x0E0)
	IOMUXC_SW_PAD_CTL_PAD_GPIO_AD_B0_09 := mmio32(IOMUXC_ADDR + 0x2D0)
	IOMUXC_GPR_GPR26 := mmio32(IOMUXC_GPR_ADDR + 0x068)

	// Set REFTOP_SELFBIASOFF after analog bandgap stabilized for best noise
	// performance of analog blocks.
	PMU_MISC0_SET.Store(1 << 3)

	// Setup PLL2
	CCM_ANALOG_PFD_528_SET.Store(0x80808080) // gate PFD0,1,2,3
	CCM_ANALOG_PFD_528.Store(0x2018101B)     // PFD0,1,2,3: 352,594,396,297 MHz

	// Setup PLL3
	//CCM_ANALOG_PFD_480_SET.Store(0x80808080) // gate PFD0,1,2,3
	//CCM_ANALOG_PFD_480.Store(0x13110D0C)     // PFD0,1,2,3: 720,664,508,454 MHz

	// Configure clocks
	CCM_CSCMR1.StoreBits(0x7F, 1<<6) // set PERCLK_CLK = OSC_CLK (24 MHz)
	CCM_CSCDR1.StoreBits(0x7F, 1<<6) // set UART_CLK   = OSC_CLK (24 MHz)

	// GPIO_MUX1 ALT5 selects GPIO6 (fast) instead of GPIO1 (slow) for all bits.
	IOMUXC_GPR_GPR26.Store(0xFFFFFFFF)

	// Connect pad AD_B0_09 to GPIO 1 or 6 (ALT5 iomux mode)
	IOMUXC_SW_MUX_CTL_PAD_GPIO_AD_B0_09.Store(5)

	// Configure pad AD_B0_09: hysteresis:off, 100K?? pull-down, pull/keeper:off,
	// open-drain:off, speed:low (50 MHz), drive-strength:(150/7)??, sr:slow
	IOMUXC_SW_PAD_CTL_PAD_GPIO_AD_B0_09.Store(7 << 3)

	rtos.SetPrivLevel(privLevel)

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
