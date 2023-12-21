// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espnet is an ESP-AT based TCP echo server. See also ../espat that uses the
// espat package directly and has much lower memory requirements. See also the
// same example written for FET1061 and STM32 development boards.
package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
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
	espTx := pins.P14
	espRx := pins.P15

	// ESP-AT
	u := lpuart2.Driver()
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(espRx, lpuart.RXD)
	u.UsePin(espTx, lpuart.TXD)
	u.EnableRx(512)
	u.EnableTx()

	fmt.Print("Initializing ESP-AT module... ")
	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	fatalErr(espnet.SetPasvRecv(dev, true))
	fmt.Println("OK")

	fmt.Println("Waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		fmt.Println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}
	txt, err := dev.CmdStr("+CIPSTA?")
	fatalErr(err)
	fmt.Println(strings.ReplaceAll(txt, "+CIPSTA:", ""))

	ls, err := espnet.ListenDev(dev, "tcp", ":1111")
	fatalErr(err)

	fmt.Println("Listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c net.Conn) {
	var buf [64]byte
	fmt.Println("Connected:", c.RemoteAddr().String())
	_, err := io.WriteString(c, "Echo Server\n\n")
	if logErr(err) {
		return
	}
	for {
		n, err := c.Read(buf[:])
		if err == io.EOF {
			break
		}
		if logErr(err) {
			return
		}
		_, err = c.Write(buf[:n])
		if logErr(err) {
			return
		}
	}
	c.Close()
	fmt.Println("Closed:   ", c.RemoteAddr().String())
}
