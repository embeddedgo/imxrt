// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dma/dmairq"
	"github.com/embeddedgo/imxrt/hal/lpspi"
)

func NewMasterDMA(p *lpspi.Periph) *lpspi.Master {
	d := dma.DMA(0)
	d.EnableClock(true)
	rxdma := d.AllocChannel(false)
	txdma := d.AllocChannel(false)
	spi := lpspi.NewMaster(p, rxdma, txdma)
	dmairq.SetISR(rxdma, spi.RxDMAISR)
	dmairq.SetISR(txdma, spi.TxDMAISR)
	return spi
}
