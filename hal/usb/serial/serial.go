// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package serial

import (
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/usb"
)

type Serial struct {
	d   *usb.Device
	ihe int
	ohe int
	tda *[6]usb.DTD
	buf *[128]byte
}

func New(d *usb.Device, ihe, ohe int) *Serial {
	return &Serial{
		d:   d,
		ihe: ihe,
		ohe: ohe,
		tda: (*[6]usb.DTD)(usb.MakeSliceDTD(6, 6)),
		buf: dma.New[[128]byte](),
	}
}

func (se *Serial) Read(p []byte) (int, error) {
	return 0, nil
}
