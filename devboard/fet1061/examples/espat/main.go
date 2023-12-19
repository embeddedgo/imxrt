// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espnet is an ESP-AT based TCP echo server. See also ../espat that uses the
// espat package directly and has much lower memory requirements. See also th
// same example written for Teensy 4 and STM32 development boards.
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/embeddedgo/espat"
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

	fmt.Print("Initializing ESP-AT module... ")
	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	_, err := dev.Cmd("+CIPMUX=1")
	fatalErr(err)
	_, err = dev.Cmd("+CIPRECVMODE=1")
	fatalErr(err)
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

	const port = "1111"

	dev.SetServer(true)
	_, err = dev.Cmd("+CIPSERVER=1," + port)
	fatalErr(err)

	fmt.Println("Listen on :" + port)
	for conn := range dev.Server() {
		go handle(conn)
	}
}

var welcome = []byte("Echo Server\n\n")

func handle(c *espat.Conn) {
	fmt.Println("Connected:", c.ID)
	if logErr(send(c, welcome)) {
		return
	}
	var buf [64]byte
	for {
		if _, ok := <-c.Ch; !ok {
			break // connection closed by remote part
		}
		n, err := c.Dev.CmdInt("+CIPRECVDATA=", buf[:], c.ID, len(buf))
		if logErr(err) {
			return
		}
		if logErr(send(c, buf[:n])) {
			return
		}
	}
	fmt.Println("Closed:   ", c.ID)
}

func send(c *espat.Conn, p []byte) error {
	c.Dev.Lock()
	defer c.Dev.Unlock()
	if _, err := c.Dev.UnsafeCmd("+CIPSEND=", c.ID, len(p)); err != nil {
		return err
	}
	if _, err := c.Dev.UnsafeWrite(p); err != nil {
		return err
	}
	_, err := c.Dev.UnsafeCmd("")
	return err
}
