// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package systick allows to use ARMv7-M SysTick timer as ticking system timer.
package systick

import (
	"embedded/arch/cortexm/systim"
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/imxrt/p/ccm"
)

// Setup setups the SysTick timer as a system timer. It is used in ticking mode
// running the thread scheduler every periodns nanoseconds. Considr using a
// tickless timer (eg. RTC based) if available.
//
// The sheduler uses WFE instruction to suspend code execution until an event
// occurs. In case of i.MX RT the default response for WFE/WFI instruction is to
// enter the Wait mode. In this mode the whole Cortex-M7 core including NVIC is
// frozen. The system can be awakened by the GPC Interrupt Controller but the
// SYSTICK interrupt is an internal signal of the CM7 core so it isn't routed to
// the GPC.
//
// All this means that SysTick timer is almost useless as a system timer but can
// be used for testing or educational purposes. In order for it to be usable,
// Setup prevents entering Wait mode after WFE/WFI CLPCR.LPM to LPM_RUN.
func Setup(periodns int64) {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	ccm.CCM().CLPCR.StoreBits(ccm.LPM, ccm.LPM_RUN)
	systim.Setup(periodns, 100e3, true)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
	rtos.SetSystemTimer(systim.Nanotime, nil)
}
