// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Httpserver runs the Go net/http server using the ESP-AT TCP/IP stack.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart2"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func fatalErr(err error) {
	for err != nil {
		log.Println("error:", err.Error())
		time.Sleep(time.Second)
	}
}

func main() {
	log.Default().SetOutput(os.Stdout)

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

	time.Sleep(5 * time.Second) // to allow to see messages printed below

	log.Print("Initializing ESP-AT module... ")
	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	fatalErr(espnet.SetPasvRecv(dev, true))
	log.Println("OK")

	log.Println("Waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		log.Println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}
	txt, err := dev.CmdStr("+CIPSTA?")
	fatalErr(err)
	log.Println(strings.ReplaceAll(txt, "+CIPSTA:", ""))

	ls, err := espnet.ListenDev(dev, "tcp", ":80")
	fatalErr(err)
	fmt.Println("Listen on:", ls.Addr())
	fatalErr(http.Serve(ls, http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.RequestURI)
	fmt.Fprintln(w, "Go HTTP server")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Method:    ", r.Method)
	fmt.Fprintln(w, "URL:       ", r.URL)
	fmt.Fprintln(w, "Proto:     ", r.Proto)
	fmt.Fprintln(w, "Host:      ", r.Host)
	fmt.Fprintln(w, "RemoteAddr:", r.RemoteAddr)
	fmt.Fprintln(w, "RequestURI:", r.RequestURI)
}
