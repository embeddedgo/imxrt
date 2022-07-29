// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
)

// Pin represents one phisical pin (specific pin in specific port).
type Pin struct {
	h uintptr // [31:5] port address, [4:0] pin number
}

// IsValid reports whether p represents a valid pin.
func (p Pin) IsValid() bool {
	return p.h&^0x1F != 0
}

// Port returns the port where the pin is located.
func (p Pin) Port() *Port {
	return (*Port)(unsafe.Pointer(p.h &^ 0x1F))
}

// Num returns the pin number in the port.
func (p Pin) Num() int {
	return int(p.h & 0x1F)
}

// Mask returns bitmask that represents the pin.
func (p Pin) Mask() Pins {
	return Pin0 << uint(p.Num())
}

func (p Pin) DirOut() bool {
	return p.Port().DirOut.LoadPins(p.Mask()) != 0
}

// SetDirOut sets pin direction to output if out or input otherwise.
func (p Pin) SetDirOut(out bool) {
	p.Port().DirOut.StorePins(p.Mask(), Pins(-internal.BoolToInt(out)))
}

// Samples the value of the pin (also in output mode if AltFunc.SION set).
func (p Pin) Load() int {
	return int(p.Port().Sample.Load()) >> uint(p.Num()) & 1
}

// LoadOut returns the output value of the pin if in output mode.
func (p Pin) LoadOut() int {
	return int(p.Port().DR.Load()) >> uint(p.Num()) & 1
}

// Set sets output value of the pin to 1 in one atomic operation.
func (p Pin) Set() {
	p.Port().SetDR.Store(Pin0 << uint(p.Num()))
}

// Clear sets output value of the pin to 0 in one atomic operation.
func (p Pin) Clear() {
	p.Port().ClearDR.Store(Pin0 << uint(p.Num()))
}

// Toggle toggles output value of the pin in one atomic operation.
func (p Pin) Toggle() {
	p.Port().ToggleDR.Store(Pin0 << uint(p.Num()))
}

// Store sets output value of the pin to the least significant bit of val.
func (p Pin) Store(val int) {
	port := p.Port()
	mask := Pin0 << uint(p.Num())
	if val&1 != 0 {
		port.SetDR.Store(mask)
	} else {
		port.ClearDR.Store(mask)
	}
}
