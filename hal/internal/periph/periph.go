// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package periph

import (
	"github.com/embeddedgo/imxrt/hal/iomux"
)

func AltFunc(pins []iomux.Pin, afs []iomux.AltFunc, psig int, pin iomux.Pin) iomux.AltFunc {
	i := uint(0)
	for ; psig != 0; psig-- {
		i += uint(afs[i]) >> 4
	}
	for m := i + uint(afs[i]>>4); i < m; i++ {
		if pins[i] == pin {
			return afs[i] & 0x0f
		}
	}
	return -1
}

func Pins(pins []iomux.Pin, afs []iomux.AltFunc, psig int) []iomux.Pin {
	i := uint(0)
	for ; psig != 0; psig-- {
		i += uint(afs[i]) >> 4
	}
	return pins[i : i+uint(afs[i]>>4)]
}
