// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dma/dmairq"
	"github.com/embeddedgo/imxrt/hal/lpi2c"
)

func NewMasterDMA(p *lpi2c.Periph) *lpi2c.Master {
	d := dma.DMA(0)
	d.EnableClock(true)
	dc := d.AllocChannel(false)
	m := lpi2c.NewMaster(p, dc)
	dmairq.SetISR(dc, m.DMAISR)
	return m
}
