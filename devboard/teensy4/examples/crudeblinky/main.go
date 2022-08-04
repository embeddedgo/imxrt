// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Crudeblinky flashes the on-board LED without using HAL or any other packages
// outside the Embedded Go standard library.
package main

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"unsafe"
)

const (
	IOMUXC_ADDR uintptr = 0x401F8000
	GPIO2_ADDR  uintptr = 0x401BC000
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
	// By default, access to most peripherals require the supervisor privilege
	// level. This can be changed in APISTZ registers.
	runtime.LockOSThread()
	rtos.SetPrivLevel(0)

	// Configure the LED pin: hysteresis:off, 100KΩ pull-down, pull/keeper:off,
	// open-drain:off, speed:low (50 MHz), drive-strength:(150/7)Ω, sr:slow
	PAD_CTL_B0_03 := (*mmio.U32)(unsafe.Pointer(IOMUXC_ADDR + 0x338))
	PAD_CTL_B0_03.Store(7 << 3)

	// By default the B0_03 pin is connected to the GPIO2 bit 3 (ALT5 mux mode)
	// so we don't need to change anything in MUX_CTL_B0_03.

	// Configure GPIO2 bit 3 as output.
	GPIO2 := (*GPIO)(unsafe.Pointer(GPIO2_ADDR))
	GPIO2.GDIR.SetBit(3) // output mode

	// Blinking in a loop. As the system timer isn't initialized the required
	// delays are implemented by writting GPIO registers multiple times.
	for {
		for i := 0; i < 1e6; i++ {
			GPIO2.DR_SET.Store(1 << 3)
		}
		for i := 0; i < 2e7; i++ {
			GPIO2.DR_CLEAR.Store(1 << 3)
		}
	}
}
