// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	_ "github.com/embeddedgo/imxrt/devboard/teensy4/board/system"
	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
)

var User LED // Orange LED

type LED struct{ bit gpio.Bit }

func (d LED) SetOn()        { d.bit.Set() }
func (d LED) SetOff()       { d.bit.Clear() }
func (d LED) Toggle()       { d.bit.Toggle() }
func (d LED) Set(on int)    { d.bit.Store(on) }
func (d LED) Get() int      { return d.bit.Load() }
func (d LED) Pin() gpio.Bit { return d.bit }

func init() {
	User.bit = gpio.UsePin(iomux.B0_03, true)
	User.bit.SetDirOut(true)
	iomux.B0_03.Setup(iomux.Drive7)
}
