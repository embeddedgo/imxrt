// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/imxrt/devboard/fet1061/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
)

func logErr(err error) bool {
	if err != nil {
		fmt.Println("error:", err.Error())
		return true
	}
	return false
}

func fatalErr(err error) {
	for err != nil {
		fmt.Println("error:", err.Error())
		time.Sleep(time.Second)
	}
}

func main() {
	// IO pins
	espTx := pins.P12
	espRx := pins.P13
	conRx := pins.P23
	conTx := pins.P24

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	// ESP-AT
	u := lpuart2.Driver()
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(espRx, lpuart.RXD)
	u.UsePin(espTx, lpuart.TXD)
	u.EnableRx(256)
	u.EnableTx()

	fmt.Println("\n* Ready *\n\n")

	dev := espat.NewDevice("esp0", u, u)
	dev.Init(false)
	fatalErr(espnet.SetPasvRecv(dev, true))

	/*
		for msg := range dev.Async() {
			fmt.Println(msg)
			if msg == "WIFI GOT IP" {
				break
			}
		}
	*/

	ls, err := espnet.ListenDev(dev, "tcp", 1111)
	fatalErr(err)

	fmt.Println("listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c net.Conn) {
	fmt.Println("connected:", c.RemoteAddr().String())
	defer fmt.Println("closed:   ", c.RemoteAddr().String())

	_, err := io.WriteString(c, "Echo Server\n\n")
	if logErr(err) {
		return
	}
	var buf [64]byte
	for {
		n, err := c.Read(buf[:])
		if err == io.EOF {
			return
		}
		if logErr(err) {
			return
		}
		_, err = c.Write(buf[:n])
		if logErr(err) {
			return
		}
	}
}
