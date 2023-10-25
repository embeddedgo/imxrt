// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/hal/system/timer/systick"
)

func init() {
	/*
		// Reconfigure the internal USB regulator.
		const (
			out3v000 = 15 << pmu.OUTPUT_TRGn
			boo0v150 = 6 << pmu.BO_OFFSETn
		)
		pmu.PMU().REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)
	*/

	system.Setup528_FlexSPI()
	systick.Setup(2e6)
}
