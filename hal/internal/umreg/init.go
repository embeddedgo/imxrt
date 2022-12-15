// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package umreg containt init function that enables user mode code to access
// peripheral registers.
//
// TODO: There must be a clearer way to do this.
package umreg

import (
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/imxrt/hal/internal/aipstz"
)

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
