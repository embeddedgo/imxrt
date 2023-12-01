// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uartcon

import (
	"embedded/rtos"
	"os"
	"syscall"

	"github.com/embeddedgo/fs/termfs"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/usb/usbserial"
)

var se *usbserial.Driver

func write(_ int, p []byte) int {
	n, _ := se.Write(p)
	return n
}

// Setup setpus an USB serial to work as system console.
func Setup(d *usbserial.Driver, name string) {
	d.SetWriteSink(true)
	//d.SetAutoFlush(true)

	// Set a system writer for print, println, panic, etc.
	se = d
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


// SetupLight setpus an USB serial to work as light system console.
// It usese termfs.LightFS instead of termfs.FS.
func SetupLight(d *usbserial.Drive, name string) {
	d.SetWriteSink(true)
	//d.SetAutoFlush(true)

	// Set a system writer for print, println, panic, etc.
	se = d
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
