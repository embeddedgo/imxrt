// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uartcon

import (
	"embedded/rtos"
	"os"
	"syscall"

	"github.com/embeddedgo/fs/termfs"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var uart *lpuart.Driver

func write(_ int, p []byte) int {
	n, _ := uart.Write(p)
	return n
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// Setup setpus an LPUART peripheral to work as system console.
func Setup(d *lpuart.Driver, rx, tx iomux.Pin, conf lpuart.Config, baudrate int, name string) {
	// Setup and enable the LPUART driver.
	d.UsePin(rx, lpuart.RXD)
	d.UsePin(tx, lpuart.TXD)
	d.Setup(conf, baudrate)
	d.EnableTx()
	d.EnableRx(64)

	// Set a system writer for print, println, panic, etc.
	uart = d
	rtos.SetSystemWriter(write)

	// Setup a serial console (standard input and output).
	con := termfs.New(name, d, d)
	con.SetCharMap(termfs.InCRLF | termfs.OutLFCRLF)
	con.SetEcho(true)
	con.SetLineMode(true, 256)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	checkErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	checkErr(err)
	os.Stderr = os.Stdout
}

// SetupLight setpus an LPUART to work as light system console.
// It usese termfs.LightFS instead of termfs.FS.
func SetupLight(d *lpuart.Driver, rx, tx iomux.Pin, conf lpuart.Config, baudrate int, name string) {
	// Setup and enable the LPUART driver.
	d.UsePin(rx, lpuart.RXD)
	d.UsePin(tx, lpuart.TXD)
	d.Setup(conf, baudrate)
	d.EnableTx()
	d.EnableRx(64)

	// Set a system writer for print, println, panic, etc.
	uart = d
	rtos.SetSystemWriter(write)

	// Setup a serial console (standard input and output).
	con := termfs.NewLight(name, d, d)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	checkErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	checkErr(err)
	os.Stderr = os.Stdout
}
