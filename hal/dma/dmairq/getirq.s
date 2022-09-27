// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

#include "textflag.h"

// func getIRQ() int
TEXT ·getIRQ(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  IPSR, R0
	AND   $0x1ff, R0
	SUB   $16, R0  // convert from exception number to IRQ number
	MOVW  R0, ret+0(FP)
	RET
