// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import "fmt"

// An Error may be used by a higher-level driver to inform about unsuccessfull
// transfer with the additional information provided by the device driver or the
// hardware controller.
type Error struct {
	Controller int    // controler number
	Function   string // function name
	HE         uint8  // hardware endpoint number
	Status     uint8  // DTD status field
}

// NotReady reports whether the error was returned because the USB device was
// in the unconfigured state or the selected configuration number does not match
// the required one.
func (e *Error) NotReady() bool {
	return e.Status&^Active == 0
}

func (e *Error) Error() string {
	dir := e.HE&1 ^ 1
	return fmt.Sprintf(
		"USB controler: %d, function: %s: endpoint: %d %s (he:%d): status: %#08b",
		e.Controller, e.Function, e.HE>>1, "inout"[dir*2:dir*3+2], e.HE, e.Status,
	)
}
