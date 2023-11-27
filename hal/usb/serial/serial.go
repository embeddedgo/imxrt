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
	interf     uint8
	txe, rxe   uint8
	wn         uint8
	rn, ri     int
	tda        *[3]usb.DTD  // tda[:2]   is for Write, tda[2] is for Read,
	donea      [3]rtos.Note // donea[:2] is for Write, donea[2] is for Read
	buf        []byte       // buf[:len] is for Read, buf[len:cap] is for Write
	lineCoding [7]byte
	autoFlush  bool
	rxto       time.Duration
	txto       time.Duration
}

func (s *Serial) SetRxTimeout(to time.Duration) {
	s.rxto = to
}

func (s *Serial) SetTxTimeout(to time.Duration) {
	s.txto = to
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
		fmt.Println("2:")
		s.tda[2].Print()
		fmt.Println()

		s.d.Print(s.txe)
		fmt.Println("0:")
		s.tda[0].Print()
		fmt.Println("1:")
		s.tda[1].Print()
	}
}

// New... rxe (host out), txe (host in).
// MaxPkt must be power of two and equal or multiple of the maximum packet size
// declared in the OUT endpoint descriptor used by this driver as Rx endpoint.
func New(d *usb.Device, interf uint8, rxe, txe int8, maxPkt, config int) *Serial {
	if bits.OnesCount(uint(maxPkt)) != 1 {
		panic("serial: maxPkt must be power of two")
	}
	s := &Serial{
		d:      d,
		config: uint8(config),
		interf: interf,
		txe:    usb.HE(txe, usb.IN),
		rxe:    usb.HE(rxe, usb.OUT),
		tda:    (*[3]usb.DTD)(usb.MakeSliceDTD(3, 3)),
		buf:    dtcm.MakeSlice[byte](1, maxPkt, maxPkt+2*dma.CacheLineSize),
	}
	s.tda[0].SetNote(&s.donea[0])
	s.tda[1].SetNote(&s.donea[1])
	s.tda[2].SetNote(&s.donea[2])
	interfaces[interf] = s
	d.Handle(0, 0x2021, setLineCoding)
	d.Handle(0, 0x20a1, getLineCoding)
	d.Handle(0, 0x2221, setControlLineState)
	//go log(s)
	return s
}

var A, M, N int

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
		if m := len(p) &^ (len(buf) - 1); m != 0 {
			// p is cache-aligned and can hold at least one packet.
			buf = p[:m] // so use it directly as the receive buffer
			rtos.CacheMaint(rtos.DCacheInval, unsafe.Pointer(&buf[0]), m)
		}
	}
	td, done := &s.tda[2], &s.donea[2]
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

// SetAutoFlush enables/disables the AutoFlush mode. If AutoFlush is enabled
// Write calls Flush before exit.
func (s *Serial) SetAutoFlush(af bool) {
	s.autoFlush = af
}

// Write implements io.Writer interface.
func (s *Serial) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	dtcm := s.buf[len(s.buf):cap(s.buf)]
	nh := len(p) // unaligned head bytes, send through dtcm buffer
	nm := 0      // middle bytes, send directly from p, require rtos.DCacheFlush
	nt := 0      // unaligned tail bytes, send through dtcm buffer
	if nh > len(dtcm) {
		const align = dma.CacheLineSize - 1
		nh = int(dma.CacheLineSize-uintptr(unsafe.Pointer(&p[0]))) & align
		nm = len(p) - nh
		nt = nm & align
		nm -= nt
		if nm != 0 {
			rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(&p[nh]), nm)
		}
	}
	var (
		status uint8
		buf    unsafe.Pointer
		m      int
	)
	wn := int(s.wn)
	if nh != 0 {
		m = nh
		goto useDTCM
	}
next:
	if nm != 0 {
		buf = unsafe.Pointer(&p[n])
		m = nm
		nm = 0
		goto loop
	}
	if nt != 0 {
		m = nt
		nt = 0
		goto useDTCM
	}
	if s.autoFlush {
		err = s.Flush()
	} else {
		s.wn = uint8(2 | wn&1)
	}
	return
useDTCM:
	copy(dtcm, p[n:n+m])
	buf = unsafe.Pointer(&dtcm[0])
loop:
	for {
		td, done := &s.tda[wn&1], &s.donea[wn&1]
		k := td.SetupTransfer(buf, m)
		done.Clear()
		if !s.d.Prime(s.txe, td, td, int(s.config)) {
			goto error
		}
		if wn != 0 {
			td, done = &s.tda[(wn-1)&1], &s.donea[(wn-1)&1]
			done.Sleep(-1)
			_, status = td.Status()
			if status != 0 {
				goto error
			}
		}
		wn++
		n += k
		m -= k
		if m == 0 {
			goto next
		}
	}
error:
	s.wn = 0
	return n, &usb.Error{s.d.Controller(), "serial", s.txe, status}
}

// Flush ensures that the last data written were sent to the USB host.
func (s *Serial) Flush() error {
	if s.wn == 0 {
		return nil
	}
	wn := (s.wn - 1) & 1
	s.wn = 0
	td, done := &s.tda[wn], &s.donea[wn]
	done.Sleep(-1)
	if _, status := td.Status(); status != 0 {
		return &usb.Error{s.d.Controller(), "serial", s.txe, status}
	}
	return nil
}

var interfaces = make(map[uint8]*Serial)

func setLineCoding(cr *usb.ControlRequest) int {
	s := interfaces[uint8(cr.Index)]
	if s == nil {
		return 0
	}
	d := cr.Data
	if len(d) < 7 {
		return 0
	}
	copy(s.lineCoding[:], d)
	/*
		fmt.Printf("cdcACMSetLineCoding:\r\n")
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
	*/
	return 0
}

func getLineCoding(cr *usb.ControlRequest) int {
	s := interfaces[uint8(cr.Index)]
	if s == nil {
		return 0
	}
	//fmt.Printf("cdcACMGetLineCoding")
	return copy(cr.Data, s.lineCoding[:])
}

func setControlLineState(cr *usb.ControlRequest) int {
	s := interfaces[uint8(cr.Index)]
	if s == nil {
		return 0
	}
	/*
		fmt.Printf("cdcACMSetControlLineState:\r\n")
		fmt.Printf(" -interface: %d\r\n", cr.Index)
		fmt.Printf(" -DTE:       %d\r\n", cr.Value&1)
		fmt.Printf(" -RTS:       %d\r\n", cr.Value>>1&1)
	*/
	return 0
}
