// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c3dma

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpi2c"
	"github.com/embeddedgo/imxrt/hal/lpi2c/internal"
)

var master *lpi2c.Master

func Master() *lpi2c.Master {
	if master == nil {
		master = internal.NewMasterDMA(lpi2c.LPI2C(3))
		irq.LPI2C3.Enable(rtos.IntPrioLow, 0)
	}
	return master
}

//go:interrupthandler
func _LPI2C3_Handler() { master.ISR() }

//go:linkname _LPI2C3_Handler IRQ30_Handler
