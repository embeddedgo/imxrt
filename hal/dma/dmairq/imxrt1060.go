// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build imxrt1060

package dmairq

import (
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
)

var handlers [32]uintptr

func setISR(c dma.Channel, isr func()) {
	h := *(*uintptr)(unsafe.Pointer(&isr))
	atomic.StoreUintptr(&handlers[c.Num()], h)
}

func dispatch(cn int) {
	d := dma.DMA(0)
	for {
		if d.Channel(cn).IsInt() {
			if h := atomic.LoadUintptr(&handlers[cn]); h != 0 {
				(*(*func())(unsafe.Pointer(&h)))()
			}
		}
		if cn += 16; cn >= 32 {
			break
		}
	}
}

// TODO: Avoid 16 separate handlers below if we move exception vectors to ITCM
// (take ~1280 bytes of flash/cache). Use one ISR for all 16 IRQs and getIRQ
// from getirq.s to obtain the interrupt number.

//go:interrupthandler
func _DMA0_DMA16_Handler() { dispatch(0) }

//go:interrupthandler
func _DMA1_DMA17_Handler() { dispatch(1) }

//go:interrupthandler
func _DMA2_DMA18_Handler() { dispatch(2) }

//go:interrupthandler
func _DMA3_DMA19_Handler() { dispatch(3) }

//go:interrupthandler
func _DMA4_DMA20_Handler() { dispatch(4) }

//go:interrupthandler
func _DMA5_DMA21_Handler() { dispatch(5) }

//go:interrupthandler
func _DMA6_DMA22_Handler() { dispatch(6) }

//go:interrupthandler
func _DMA7_DMA23_Handler() { dispatch(7) }

//go:interrupthandler
func _DMA8_DMA24_Handler() { dispatch(8) }

//go:interrupthandler
func _DMA9_DMA25_Handler() { dispatch(9) }

//go:interrupthandler
func _DMA10_DMA26_Handler() { dispatch(10) }

//go:interrupthandler
func _DMA11_DMA27_Handler() { dispatch(11) }

//go:interrupthandler
func _DMA12_DMA28_Handler() { dispatch(12) }

//go:interrupthandler
func _DMA13_DMA29_Handler() { dispatch(13) }

//go:interrupthandler
func _DMA14_DMA30_Handler() { dispatch(14) }

//go:interrupthandler
func _DMA15_DMA31_Handler() { dispatch(15) }

func enableIRQs(prio int) {
	for i := irq.DMA0_DMA16; i < irq.DMA_ERROR; i++ {
		i.Enable(prio, 0)
	}
}

//go:linkname _DMA0_DMA16_Handler IRQ0_Handler
//go:linkname _DMA1_DMA17_Handler IRQ1_Handler
//go:linkname _DMA2_DMA18_Handler IRQ2_Handler
//go:linkname _DMA3_DMA19_Handler IRQ3_Handler
//go:linkname _DMA4_DMA20_Handler IRQ4_Handler
//go:linkname _DMA5_DMA21_Handler IRQ5_Handler
//go:linkname _DMA6_DMA22_Handler IRQ6_Handler
//go:linkname _DMA7_DMA23_Handler IRQ7_Handler
//go:linkname _DMA8_DMA24_Handler IRQ8_Handler
//go:linkname _DMA9_DMA25_Handler IRQ9_Handler
//go:linkname _DMA10_DMA26_Handler IRQ10_Handler
//go:linkname _DMA11_DMA27_Handler IRQ11_Handler
//go:linkname _DMA12_DMA28_Handler IRQ12_Handler
//go:linkname _DMA13_DMA29_Handler IRQ13_Handler
//go:linkname _DMA14_DMA30_Handler IRQ14_Handler
//go:linkname _DMA15_DMA31_Handler IRQ15_Handler

// //go:linkname DMA_ERROR_Handler IRQ16_Handler
