// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	_ "github.com/embeddedgo/imxrt/devboard/fet1061/board/system"
	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

var User LED // LED5 (blue)

type LED struct{ bit gpio.Bit }

func (d LED) SetOn()        { d.bit.Clear() }
func (d LED) SetOff()       { d.bit.Set() }
func (d LED) Toggle()       { d.bit.Toggle() }
func (d LED) Set(on int)    { d.bit.Store(on) }
func (d LED) Get() int      { return d.bit.Load() }
func (d LED) Pin() gpio.Bit { return d.bit }

func init() {
	User.bit = gpio.UsePin(iomux.AD_B0_09, false)
	User.bit.Port().EnableClock(true) // lp=true to don't interfere with other users of this port that enabled it earlier and require clock in low-power mode
	User.bit.SetDirOut(true)
	User.SetOff()
	iomux.AD_B0_09.Setup(iomux.Drive7)
}
