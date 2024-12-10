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

// Lock disables write access to the IOMUX controller. It is typically used to
// prevent accidental changes to the electrical properties of the pins, as such
// a change can sometimes be destructive.
//
// Use carefully!
//
// Locked IOMUX prevents configuring pins. Try to keep the configuration of all
// pins and peripherals in one place. Remember that UsePin functions in HAL
// access IOMUX to set the correct alternate function for pin.
//
// Write access to the locked IOMUX generates the BusFault exception which can
// be very dificult to debug so do not lock IOMUX during development. Lock it
// at the end of the testing stage and in the finished product.
//
// Locked IOMUX may interfere with some debugger and programmer software and
// hardware.
func Lock() {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	internal.ExclusiveStoreBits(&aipstz.P(2).OPACR[3], aipstz.WP<<4, aipstz.WP<<4)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
}

// Unlock enables write access to the IOMUX controller.
func Unlock() {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	internal.ExclusiveStoreBits(&aipstz.P(2).OPACR[3], aipstz.WP<<4, 0)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
}
