// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usbcon

import (
	"embedded/rtos"
	"os"
	"syscall"

	"github.com/embeddedgo/fs/termfs"
	"github.com/embeddedgo/imxrt/hal/usb/usbserial"
)

var us *usbserial.Driver

func write(_ int, p []byte) int {
	n, _ := us.Write(p)
	return n
}

func panicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Setup setpus an USB serial to work as the system console.
func Setup(d *usbserial.Driver, name string) {
	d.SetWriteSink(true)
	d.SetAutoFlush(true)

	// Set a system writer for print, println, panic, etc.
	us = d
	rtos.SetSystemWriter(write)

	// Setup a serial console (standard input and output).
	con := termfs.New(name, d, d)
	con.SetCharMap(termfs.InCRLF | termfs.OutLFCRLF)
	con.SetEcho(true)
	con.SetLineMode(true, 256)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	panicErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	panicErr(err)
	os.Stderr = os.Stdout
}

// SetupLight setpus an USB serial to work as the system console.
// It usese termfs.LightFS instead of termfs.FS.
func SetupLight(d *usbserial.Driver, name string) {
	d.SetWriteSink(true)
	d.SetAutoFlush(true)

	// Set a system writer for print, println, panic, etc.
	us = d
	rtos.SetSystemWriter(write)

	// Setup a serial console (standard input and output).
	con := termfs.NewLight(name, d, d)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	panicErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	panicErr(err)
	os.Stderr = os.Stdout
}
