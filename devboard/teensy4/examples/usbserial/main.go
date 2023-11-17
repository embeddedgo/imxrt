// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"io"
	"time"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/hal/usb"
	"github.com/embeddedgo/imxrt/hal/usb/serial"
)

var usbd *usb.Device

func cdcSetLineCoding(cr *usb.ControlRequest) int {
	fmt.Printf("cdcSetLineCoding: %+v\r\n", *cr)
	fmt.Printf(" -interface: %d\r\n", cr.Index)
	d := cr.Data
	if len(d) < 7 {
		fmt.Println(" len(d) =", len(d))
		return 0
	}
	fmt.Printf(" -baudrate:  %d\r\n", d[0]+d[1]<<8+d[2]<<16+d[3]<<24)
	fmt.Printf(" -stop bits: %.1f\r\n", float32(d[4]+2)/2)
	par := d[5]
	if par > 5 {
		par = 5
	}
	par *= 4
	fmt.Printf(" -parity:    %s\r\n", "noneodd evenmarkspacunkn"[par:par+4])
	fmt.Printf(" -data bits: %d\r\n", d[6])
	return 0
}

func cdcSetControlLineState(cr *usb.ControlRequest) int {
	fmt.Printf("cdcSetControlLineState: %+v\r\n", *cr)
	return 0
}

var dot = []byte{'.'}

func main() {
	// IO pins
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	fmt.Println("Start!")

	const (
		config = 1
		in     = 2 // input endpoint, host perspective, device Tx
		out    = 2 // output endopint, host prespective, device Rx
		maxPkt = 512
	)

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	se := serial.New(usbd, out, in, maxPkt, config)
	usbd.Enable()

	time.Sleep(5 * time.Second)
	fmt.Println("Go!")

	go send(se, config)
	recv(se, config)
}

func send(w io.Writer, config int) {
	buf := make([]byte, 4096+7)[7:]
	for i := range buf {
		buf[i] = byte(i)
	}

usbNotReady:
	fmt.Println("\nsend: Waiting for USB...")
	usbd.WaitConfig(config)
	fmt.Println("\nsend: USB is ready. Sending!")

	for o := 0; ; o += 256 {
		_, err := w.Write(buf[:o&(len(buf)-1)])
		if err != nil {
			if e, ok := err.(*usb.Error); ok && e.NotReady() {
				goto usbNotReady
			}
			fmt.Println("\nsend:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		//os.Stdout.Write(sendChr)
	}
}

func recv(r io.Reader, config int) {
	buf := make([]byte, 512)
	var cnt byte

usbNotReady:
	fmt.Println("\nrecv: Waiting for USB...")
	usbd.WaitConfig(config)
	fmt.Println("\nrecv: USB is ready. Receiving!")

	for {
		n, err := r.Read(buf)
		if err != nil {
			if e, ok := err.(*usb.Error); ok && e.NotReady() {
				goto usbNotReady
			}
			fmt.Println("\nrecv:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		//os.Stdout.Write(recvChr)
		for i, b := range buf[:n] {
			buf[i] = 0
			if b != cnt {
				fmt.Printf("\nrecv: buf[%d]=%d != %d\n", i, b, cnt)
				time.Sleep(5 * time.Second)
				cnt = 0
				break
			}
			cnt++
		}
	}
}

var (
	recvChr = []byte{'r'}
	sendChr = []byte{'s'}
)

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
