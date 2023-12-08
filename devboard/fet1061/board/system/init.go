// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/embeddedgo/imxrt/hal/system"
	"github.com/embeddedgo/imxrt/hal/system/timer/systick"
)

func init() {
	system.Setup528_FlexSPI()
	systick.Setup(2e6)
}
