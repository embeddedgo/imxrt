// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/p/mmap"
)

// Pins is a bitmask which represents the pins of GPIO port.
type Pins uint32

const (
	Pin0 Pins = 1 << iota
	Pin1
	Pin2
	Pin3
	Pin4
	Pin5
	Pin6
	Pin7
	Pin8
	Pin9
	Pin10
	Pin11
	Pin12
	Pin13
	Pin14
	Pin15
	Pin16
	Pin17
	Pin18
	Pin19
	Pin20
	Pin21
	Pin22
	Pin23
	Pin24
	Pin25
	Pin26
	Pin27
	Pin28
	Pin29
	Pin30
	Pin31
)

type PinReg struct{ U32 mmio.U32 }

func (r *PinReg) Load() Pins              { return Pins(r.U32.Load()) }
func (r *PinReg) Store(pins Pins)         { r.U32.Store(uint32(pins)) }
func (r *PinReg) LoadPins(mask Pins) Pins { return Pins(r.U32.Load()) & mask }

func (r *PinReg) StorePins(mask, pins Pins) {
	internal.AtomicStoreBits(&r.U32, uint32(mask), uint32(pins))
}

type Port struct {
	DR       PinReg // Data register. Its bits are reflected on the output pins.
	DirOut   PinReg // Sets pins to output mode.
	Sample   PinReg // Samples input pins (also output pins if AltFunc.SION set)
	icr      [2]mmio.U32
	IntEna   PinReg // Enables pins as an interrupt source.
	Pending  PinReg // Interrupt pending register. Write to clear.
	EdgeSel  PinReg // Selects the edge that generates interrupts.
	_        [25]uint32
	SetDR    PinReg // Use to set bits in data register.
	ClearDR  PinReg // Use to clear bits in data register.
	ToggleDR PinReg // Use to toggle bits in data register.
}

// P returns n-th GPIO port
func P(n int) *Port {
	var addr uintptr
	switch n {
	case 1:
		addr = mmap.GPIO1_BASE
	case 2:
		addr = mmap.GPIO2_BASE
	case 3:
		addr = mmap.GPIO3_BASE
	case 4:
		addr = mmap.GPIO4_BASE
	case 5:
		addr = mmap.GPIO5_BASE
	case 6:
		addr = mmap.GPIO6_BASE
	case 7:
		addr = mmap.GPIO7_BASE
	case 8:
		addr = mmap.GPIO8_BASE
	case 9:
		addr = mmap.GPIO9_BASE
	default:
		panic("bad GPIO port number")
	}
	return (*Port)(unsafe.Pointer(addr))
}

func (p *Port) Pin(n int) Pin {
	if uint(n) > 31 {
		panic("bad GPIO pin number")
	}
	addr := uintptr(unsafe.Pointer(p))
	return Pin{addr | uintptr(n)}
}
