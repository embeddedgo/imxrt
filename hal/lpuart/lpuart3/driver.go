// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpuart3

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var driver *lpuart.Driver

// Driver returns ready to use driver for UART.
func Driver() *lpuart.Driver {
	if driver == nil {
		driver = lpuart.NewDriver(lpuart.LPUART(3), dma.Channel{}, dma.Channel{})
		irq.LPUART3.Enable(rtos.IntPrioLow, 0)
	}
	return driver
}

//go:interrupthandler
func _LPUART3_Handler() { driver.ISR() }

//go:linkname _LPUART3_Handler IRQ22_Handler
