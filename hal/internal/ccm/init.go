// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ccm

import (
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/imxrt/hal/internal/aipstz"
)

// Enable full access to all peripherals in user mode. We assume that every HAL
// package imports this package for clock management so it should be run before
// any other HAL function that access registers, in particular, befor any other
// init function.
func init() {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)

	for i := 1; i < 5; i++ {
		opacr := &aipstz.P(i).OPACR
		for k := 0; k < len(opacr); k++ {
			opacr[k].Store(0)
		}
	}

	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
}
