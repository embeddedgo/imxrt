// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to flash the onboard LED without using HAL or any
// other packages outside the Embedded Go standard library.
package main

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"unsafe"
)

const (
	IOMUXC_ADDR uintptr = 0x401F8000
	GPIO1_ADDR  uintptr = 0x401B8000
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

func main() {
	// By default, supervisor privilege level is required to access most
	// peripherals. This can be changed using in APISTZ registers.
	runtime.LockOSThread()
	rtos.SetPrivLevel(0)

	// Configure pad AD_B0_09: hysteresis:off, 100KΩ pull-down, pull/keeper:off,
	// open-drain:off, speed:low (50 MHz), drive-strength:(150/7)Ω, sr:slow
	PAD_CTL_AD_B0_09 := (*mmio.U32)(unsafe.Pointer(IOMUXC_ADDR + 0x2D0))
	PAD_CTL_AD_B0_09.Store(7 << 3)

	// By default the AD_B0_09 pad is used as JTAG_TDI (ALT0 mux mode).
	// Connect it to the GPIO1 bit 9 (ALT5 mux mode).
	MUX_CTL_AD_B0_09 := (*mmio.U32)(unsafe.Pointer(IOMUXC_ADDR + 0x0E0))
	MUX_CTL_AD_B0_09.Store(5)

	// Configure GPIO1 bit 9 as output.
	GPIO1 := (*GPIO)(unsafe.Pointer(GPIO1_ADDR))
	GPIO1.GDIR.SetBit(9)

	// Blinking in a loop. As the system timer isn't initialized the required
	// delays are implemented by writting GPIO registers multiple times.
	for {
		for i := 0; i < 1e6; i++ {
			GPIO1.DR_CLEAR.Store(1 << 9)
		}
		for i := 0; i < 2e7; i++ {
			GPIO1.DR_SET.Store(1 << 9)
		}
	}
}
