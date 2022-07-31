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

// Interrupt configuration constants
const (
	IntLow     = 0 // interrupt is low-level sensitive
	IntHigh    = 1 // interrupt is high-level sensitive
	IntRising  = 2 // interrupt is rising-edge sensitive
	IntFalling = 3 // interrupt is falling-edge sensitive
)

// IntConf returns the interrupt configuration for pin.
func (p Pin) IntConf() int {
	n := uint(p.Num())
	shift := n * 2 & 15
	return int(p.Port().icr[n>>4].Load()>>shift) & 3
}

// SetIntConf sets the interrupt configuration for pin. It supports two more
// options than Port.EdgeSel register.
func (p Pin) SetIntConf(cfg int) {
	n := uint(p.Num())
	shift := n * 2 & 15
	internal.AtomicStoreBits(&p.Port().icr[n>>4], 3<<shift, uint32(cfg<<shift))
}

// IntPending reports wheter the pending status of the pin interrupt.
func (p Pin) IntPending() bool {
	return p.Port().Pending.LoadPins(Pin0<<uint(p.Num())) != 0
}

// ClearPending clears the pending state of the pin interrupt.
func (p Pin) ClearPending() {
	p.Port().Pending.Store(Pin0 << uint(p.Num()))
}

// ConnectMux works like Port.ConnectMux(p.Mask())
func (p Pin) ConnectMux() {
	p.Port().ConnectMux(Pin0 << uint(p.Num()))
}

// ConnectMux reports wheter the pin is connected to IOMUX.
func (p Pin) MuxConnected() bool {
	return p.Port().MuxConnected()>>uint(p.Num())&1 != 0
}
