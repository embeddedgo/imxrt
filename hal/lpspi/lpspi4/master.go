// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi4

import (
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpspi"
)

var master *lpspi.Master

func Master() *lpspi.Master {
	if master == nil {
		master = lpspi.NewMaster(lpspi.LPSPI(4), dma.Channel{}, dma.Channel{})
	}
	return master
}
