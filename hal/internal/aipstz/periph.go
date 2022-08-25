// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build imxrt1060

package aipstz

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/p/mmap"
)

const (
	TP = 1 << 0
	WP = 1 << 1
	SP = 1 << 2
	BW = 1 << 3
)

type Periph struct {
	MPR   mmio.R32[uint32]
	_     [15]uint32
	OPACR [5]mmio.R32[uint32]
}

func P(i int) *Periph {
	step := mmap.AIPSTZ2_BASE - mmap.AIPSTZ1_BASE
	return (*Periph)(unsafe.Pointer(mmap.AIPSTZ1_BASE + uintptr(i-1)*step))
}
