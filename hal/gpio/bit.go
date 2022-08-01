// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
)

// Bit represents a single bit in a GPIO port.
type Bit struct {
	h uintptr // [31:5] port address, [4:0] bit number
}

// IsValid reports whether b represents a valid bit.
func (b Bit) IsValid() bool {
	return b.h&^31 != 0
}

// Port returns the port where the bit is located.
func (b Bit) Port() *Port {
	return (*Port)(unsafe.Pointer(b.h &^ 31))
}

// Num returns the bit number in the port.
func (b Bit) Num() int {
	return int(b.h & 31)
}

// Mask returns a bitmask that represents the bit.
func (b Bit) Mask() uint32 {
	return 1 << uint(b.Num())
}

func (b Bit) DirOut() bool {
	return b.Port().DirOut.LoadBits(b.Mask()) != 0
}

// SetDirOut sets the bit direction to output if out or input otherwise.
func (b Bit) SetDirOut(out bool) {
	b.Port().DirOut.StoreBits(b.Mask(), -uint32(internal.BoolToInt(out)))
}

// Samples the value of the connected pin (also in output mode if AltFunc.SION
// is set).
func (b Bit) Load() int {
	return int(b.Port().Sample.Load()) >> uint(b.Num()) & 1
}

// LoadOut returns the bit value. In output mode it is the value stored in DR
// register. In input mode it works like Load.
func (b Bit) LoadOut() int {
	return int(b.Port().DR.Load()) >> uint(b.Num()) & 1
}

// Set sets the output value of the bit to 1 in one atomic operation.
func (b Bit) Set() {
	b.Port().SetDR.Store(1 << uint(b.Num()))
}

// Clear sets the output value of the bit to 0 in one atomic operation.
func (b Bit) Clear() {
	b.Port().ClearDR.Store(1 << uint(b.Num()))
}

// Toggle toggles the output value of the bit in one atomic operation.
func (b Bit) Toggle() {
	b.Port().ToggleDR.Store(1 << uint(b.Num()))
}

// Store sets the bit value to the least significant bit of val.
func (b Bit) Store(val int) {
	port := b.Port()
	mask := uint32(1) << uint(b.Num())
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

// IntConf returns the interrupt configuration of bit.
func (b Bit) IntConf() int {
	n := uint(b.Num())
	shift := n * 2 & 15
	return int(b.Port().IntCfg[n>>4].Load()>>shift) & 3
}

// SetIntConf sets the interrupt configuration for bit.
func (b Bit) SetIntConf(cfg int) {
	n := uint(b.Num())
	shift := n * 2 & 15
	internal.AtomicStoreBits(&b.Port().IntCfg[n>>4], 3<<shift, uint32(cfg<<shift))
}

// IntPending reports whether the interrupt coresponding to b is pending.
func (b Bit) IntPending() bool {
	return b.Port().Pending.LoadBits(1<<uint(b.Num())) != 0
}

// ClearPending clears the pending state of the interrupt coresponding to b.
func (b Bit) ClearPending() {
	b.Port().Pending.Store(1 << uint(b.Num()))
}

// ConnectMux works like Port.ConnectMux(b.Mask())
func (b Bit) ConnectMux() {
	b.Port().ConnectMux(1 << uint(b.Num()))
}

// ConnectMux reports wheter the bit is connected to IOMUX.
func (b Bit) MuxConnected() bool {
	return b.Port().MuxConnected()>>uint(b.Num())&1 != 0
}
