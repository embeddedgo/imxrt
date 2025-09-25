// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"strings"

	"github.com/embeddedgo/device/bus/i2cbus"
)

const MasterErrFlags = MNDF | MALF | MFEF | MPLTF

// MasterError contains value of the Master Status Register with one or more
// error flags set.
type MasterError struct {
	Bus    string
	Status MSR // value of the Master Status Register as read by Master.Err
}

func (e *MasterError) Is(target error) bool {
	return target == i2cbus.ErrACK && e.Status&MNDF != 0
}

func (e *MasterError) Error() string {
	var a [4]string
	es := a[:0:4]
	if e.Status&MNDF != 0 {
		es = append(es, "ACK")
	}
	if e.Status&MALF != 0 {
		es = append(es, "Arbitr")
	}
	if e.Status&MFEF != 0 {
		es = append(es, "FIFO")
	}
	if e.Status&MPLTF != 0 {
		es = append(es, "PinLow")
	}
	return "lpi2c: " + e.Bus + ": " + strings.Join(es, ",")
}
