// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "unsafe"

func BoolToInt(x bool) int { return int(uint8(*(*uint8)(unsafe.Pointer(&x)))) }
