// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iomux

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/hal/internal/aipstz"
	"github.com/embeddedgo/imxrt/p/mmap"
)

type periph struct {
	_   [5]uint32
	mux [124]mmio.U32
	pad [124]mmio.U32
}

func pr() *periph {
	return (*periph)(unsafe.Pointer(mmap.IOMUXC_BASE))
}

// Lock disables write access to the IOMUX controller. Use carefully because
// locked IOMUX may interfere with some programming software or hardware.
func Lock() {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	internal.AtomicStoreBits(&aipstz.P(2).OPACR[3], aipstz.WP<<4, aipstz.WP<<4)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
}

// Unlock enables write access to the IOMUX controller.
func Unlock() {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	internal.AtomicStoreBits(&aipstz.P(2).OPACR[3], aipstz.WP<<4, 0)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
}
