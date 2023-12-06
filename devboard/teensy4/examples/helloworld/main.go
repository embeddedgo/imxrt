// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helloworld endlessly prints a line of text on the system console.
package main

import (
	"fmt"
	"time"

	_ "github.com/embeddedgo/imxrt/devboard/teensy4/board/system"
)

func main() {
	for {
		fmt.Println("Hello, World!")
		time.Sleep(time.Second)
	}
}
