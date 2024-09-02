// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package periph

import "github.com/embeddedgo/imxrt/hal/iomux"

// AltFunc returns the configuration for the pin to be used as the signal number
// sig of the peripheral number p where psig = p * numSig + sig. See Pins for
// description of pins and afs.
func AltFunc(pins []iomux.Pin, afs []iomux.AltFunc, psig int, pin iomux.Pin) (af iomux.AltFunc, sel, daisy int) {
	i := 0
	for ; psig != 0; psig-- {
		n := int(afs[i]) >> 4
		if n < 0 {
			n &= 7
			sel++
		}
		i += n
	}
	af = -1
	n := int(afs[i]) >> 4
	if n < 0 {
		n &= 7
	} else {
		sel = -1
	}
	for daisy < n {
		if pins[i+daisy] == pin {
			af = afs[i+daisy] & 0x0f
			break
		}
		daisy++
	}
	return
}

// Pins returns the list of pins that can be used for the peripheral number p
// and the signal number sig where psig = p * numSig + sig. Pins should contain
// the ordered array of pins that can be used by p. Afs contains alternate
// functions for these pins and the encoded structure of the I/O mux.
func Pins(pins []iomux.Pin, afs []iomux.AltFunc, psig int) []iomux.Pin {
	i := uint(0)
	for ; psig != 0; psig-- {
		i += uint(afs[i]) >> 4
	}
	return pins[i : i+uint(afs[i]>>4)]
}
