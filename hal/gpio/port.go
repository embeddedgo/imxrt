// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/hal/internal/ccm"
	"github.com/embeddedgo/imxrt/hal/internal/iomux"
	"github.com/embeddedgo/imxrt/p/mmap"
)

type Shared struct{ mmio.R32[uint32] }

func (b *Shared) StoreBits(mask, bits uint32) {
	internal.AtomicStoreBits(&b.R32, mask, bits)
}

func (b *Shared) SetBits(mask uint32) {
	internal.AtomicStoreBits(&b.R32, mask, mask)
}

func (b *Shared) ClearBits(mask uint32) {
	internal.AtomicStoreBits(&b.R32, mask, 0)
}

type Port struct {
	DR       Shared           // Data register. Its bits are reflected on the output pins.
	DirOut   Shared           // Sets the connected pins to the output mode.
	Sample   mmio.R32[uint32] // Samples input pins (also output pins if AltFunc.SION set)
	IntCfg   [2]Shared        // Interrupt configuration (2 bits for one input)
	IntEna   Shared           // Enables connected pins as an interrupt source.
	Pending  Shared           // Interrupt pending register. Write to clear.
	EdgeSel  Shared           // Configures the edge detector (subset of IntCfg)
	_        [25]uint32
	SetDR    mmio.R32[uint32] // Use to set bits in data register.
	ClearDR  mmio.R32[uint32] // Use to clear bits in data register.
	ToggleDR mmio.R32[uint32] // Use to toggle bits in data register.
}

var portAddrs = [...]uintptr{
	mmap.GPIO1_BASE,
	mmap.GPIO2_BASE,
	mmap.GPIO3_BASE,
	mmap.GPIO4_BASE,
	mmap.GPIO5_BASE,
	mmap.GPIO6_BASE,
	mmap.GPIO7_BASE,
	mmap.GPIO8_BASE,
	mmap.GPIO9_BASE,
}

// P returns n-th GPIO port. Ports 1 to 5 are slow, ports 6 to 9 are fast.
func P(n int) *Port { return (*Port)(unsafe.Pointer(portAddrs[n-1])) }

// Num returns the GPIO port number.
func (p *Port) Num() int {
	addr := uintptr(unsafe.Pointer(p))
	for i, base := range portAddrs {
		if addr == base {
			return i + 1
		}
	}
	return -1
}

var cgs = [...]uint8{
	1<<4 | 13,
	0<<4 | 15,
	2<<4 | 13,
	3<<4 | 6,
	1<<4 | 15,
}

func cg(p *Port) (*ccm.CCGR_, int) {
	if n := p.Num() - 1; n < len(cgs) {
		cg := cgs[n]
		return ccm.CCGR(int(cg) >> 4), int(cg) & 15
	}
	return nil, 0
}

// EnableClock enables clock for port p.
// lp determines whether the clock remains on in low power WAIT mode.
func (p *Port) EnableClock(lp bool) {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
	}
}

// DisableClock disables clock for port p.
func (p *Port) DisableClock() {
	ccgr, cgn := cg(p)
	if ccgr != nil {
		ccgr.SetCG(cgn, 0)
	}
}

func (p *Port) Bit(n int) Bit {
	if uint(n) > 31 {
		panic("bad GPIO bit number")
	}
	addr := uintptr(unsafe.Pointer(p))
	return Bit{addr | uintptr(n)}
}

// MuxConnected returns bits connected to IOMUX.
func (p *Port) MuxConnected() uint32 {
	mask := uint32(0xffffffff)
	n := p.Num()
	if n == 5 {
		return mask
	}
	if n > 5 {
		mask = 0
		n -= 5
	}
	return iomux.GPR(25+n).Load() ^ mask
}

// ConnectMux connects the bits specified in mask to the IOMUX. All bits in port
// 5 are always connected. In case of the following ports the connection is
// mutually exclusive: 1 and 6, 2 and 7, 3 and 8, 4 and 9. For example, if the
// bit 2 in port 3 has been connected to the IOMUX then at the same time the bit
// 2 in port 8 has been disconnected.
func (p *Port) ConnectMux(mask uint32) {
	n := p.Num()
	if n == 5 {
		return
	}
	v := uint32(0)
	if n > 5 {
		v = 0xffffffff
		n -= 5
	}
	internal.AtomicStoreBits(iomux.GPR(25+n), mask, v)
}
