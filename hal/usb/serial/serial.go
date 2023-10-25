// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package serial provides a simple USB CDC ACM serial driver. The main goal
// is simplicity, small code size and robustness but not speed.
//
// The USB protocol is packet-oriented. All this package does is simulate a
// stream-oriented device using a packet-oriented protocol. If all you want to
// do is sending/receiving packets of data over a CDC ACM serial its better and
// quite easy to use the CDC ACM data endponts with the usb package directly.
// Such approach will give you minimal overhead and the maximum possible speed.
package serial

import (
	"embedded/rtos"
	"fmt"
	"math/bits"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dtcm"
	"github.com/embeddedgo/imxrt/hal/usb"

	pusb "github.com/embeddedgo/imxrt/p/usb"
)

// A Serial is a simple CDC ACM driver.
type Serial struct {
	d          *usb.Device
	config     uint8
	txe        uint8
	rxe        uint8
	rn, ri     int
	tda        *[3]usb.DTD  // tda[0]  is for Read,   tda[1:] is for Write
	donea      [3]rtos.Note // donea[0] is for Read,   donea[1:] is for Write
	buf        []byte       // buf[:len] is for Read, buf[len:cap] is for Write
	lineCoding [7]byte
}

func log(s *Serial) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println()
		u := pusb.USB1()
		fmt.Printf(
			"prime: %#x status: %#x\n",
			u.ENDPTPRIME.Load(), u.ENDPTSTAT.Load(),
		)
		s.d.Print(s.rxe)
		fmt.Println("0:")
		s.tda[0].Print()
		fmt.Println()

		s.d.Print(s.txe)
		fmt.Println("1:")
		s.tda[1].Print()
		fmt.Println("2:")
		s.tda[2].Print()
	}
}

func (s *Serial) setLineCoding(cr *usb.ControlRequest) int {
	d := cr.Data
	if len(d) < 7 {
		return 0
	}
	copy(s.lineCoding[:], d)
	fmt.Printf("cdcACMSetLineCoding:")
	baud := uint(d[0]) + uint(d[1])<<8 + uint(d[2])<<16 + uint(d[3])<<24
	stop := float32(d[4]+2) / 2
	pi := d[5]
	if pi > 5 {
		pi = 5
	}
	pi *= 4
	parity := "noneodd evenmarkspacunkn"[pi : pi+4]
	data := d[6]
	fmt.Printf(" -interface: %d\r\n", cr.Index)
	fmt.Printf(" -baudrate:  %d\r\n", baud)
	fmt.Printf(" -data bits: %d\r\n", data)
	fmt.Printf(" -stop bits: %.1f\r\n", stop)
	fmt.Printf(" -parity:    %s\r\n", parity)
	return 0
}

func (s *Serial) getLineCoding(cr *usb.ControlRequest) int {
	fmt.Printf("cdcACMGetLineCoding")
	return copy(cr.Data, s.lineCoding[:])
}

func (s *Serial) setControlLineState(cr *usb.ControlRequest) int {
	fmt.Printf("cdcACMSetControlLineState:\r\n")
	fmt.Printf(" -interface: %d\r\n", cr.Index)
	fmt.Printf(" -DTE:       %d\r\n", cr.Value&1)
	fmt.Printf(" -RTS:       %d\r\n", cr.Value>>1&1)
	return 0
}

// New... rxe (host out), txe (host in).
// MaxPkt must be power of two and equal or multiple of the maximum packet size
// declared in the OUT endpoint descriptor used by this driver as Rx endpoint.
func New(d *usb.Device, rxe, txe int8, maxPkt, config int) *Serial {
	if bits.OnesCount(uint(maxPkt)) != 1 {
		panic("serial: maxPkt must be power of two")
	}
	s := &Serial{
		d:      d,
		config: uint8(config),
		txe:    usb.HE(txe, usb.IN),
		rxe:    usb.HE(rxe, usb.OUT),
		tda:    (*[3]usb.DTD)(usb.MakeSliceDTD(3, 3)),
		buf:    dtcm.MakeSlice[byte](1, maxPkt, maxPkt+dma.CacheLineSize),
	}
	s.tda[0].SetNote(&s.donea[0])
	s.tda[1].SetNote(&s.donea[1])
	s.tda[2].SetNote(&s.donea[2])
	d.Handle(0, 0x2021, s.setLineCoding)
	d.Handle(0, 0x20a1, s.getLineCoding)
	d.Handle(0, 0x2221, s.setControlLineState)

	go log(s)
	return s
}

// dd if=/dev/zero of=/dev/ttyACM0 bs=2048 status=progress

// Read implements io.Reader interface.
func (s *Serial) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	buf := s.buf
	if s.rn != s.ri {
		n = copy(p, buf[s.ri:s.rn])
		s.ri += n
		return
	}
	if uintptr(unsafe.Pointer(&p[0]))&(dma.CacheLineSize-1) == 0 {
		if m := len(p) &^ (len(s.buf) - 1); m != 0 {
			// p is cache-aligned and can hold at least one packet.
			buf = p[:m] // so use it directly as the receive buffer
			rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf[0]), m)
		}
	}
	td, done := &s.tda[0], &s.donea[0]
	n = td.SetupTransfer(unsafe.Pointer(&buf[0]), len(buf))
	done.Clear()
	var (
		m      int
		status uint8
	)
	if !s.d.Prime(s.rxe, td, td, int(s.config)) {
		goto error
	}
	done.Sleep(-1)
	m, status = td.Status()
	if status != 0 {
		goto error
	}
	n -= m
	if &buf[0] != &p[0] {
		s.rn = n
		n = copy(p, buf[:n])
		s.ri = n
	}
	return
error:
	return n, &usb.Error{s.d.Controller(), "serial", s.rxe, status}
}

// Write implements io.Writer interface.
func (s *Serial) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	tds, dones := s.tda[1:], s.donea[1:]
	dtcm := s.buf[len(s.buf):cap(s.buf)]
	m := len(p)
	if m >= dma.CacheLineSize {
		const align = dma.CacheLineSize - 1
		a := int(dma.CacheLineSize-uintptr(unsafe.Pointer(&p[0]))) & align
		if a != 0 {
			m = a
		}
		if n := (len(p) - a) &^ align; n != 0 {
			rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(&p[a]), n)
		}
	}
	// The loop below is a bit convoluted because the next transfer is primed
	// before waiting for the previous one to complete.
	var status uint8
	for i := 0; ; i++ {
		if m != 0 {
			buf := p[n:]
			if m < dma.CacheLineSize {
				copy(dtcm, buf[:m])
				buf = dtcm
			}
			td, done := &tds[i&1], &dones[i&1]
			m = td.SetupTransfer(unsafe.Pointer(&buf[0]), m)
			done.Clear()
			if !s.d.Prime(s.txe, td, td, int(s.config)) {
				goto error
			}
		}
		if i != 0 {
			td, done := &tds[(i-1)&1], &dones[(i-1)&1]
			done.Sleep(-1)
			_, status = td.Status()
			if status != 0 {
				goto error
			}
		}
		if m == 0 {
			return
		}
		n += m
		m = len(p) - n
	}
error:
	return n, &usb.Error{s.d.Controller(), "serial", s.txe, status}
}
