// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi2dma

import (
	"github.com/embeddedgo/imxrt/hal/lpspi"
	"github.com/embeddedgo/imxrt/hal/lpspi/internal"
)

var master *lpspi.Master

func Master() *lpspi.Master {
	if master == nil {
		master = internal.NewMasterDMA(lpspi.LPSPI(2))
	}
	return master
}
