// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package usbserial provides a simple USB CDC ACM serial driver. The main goal
// is simplicity, small code size and robustness but not speed.
//
// The USB protocol is packet-oriented. This package simulates a stream-oriented
// device using a packet-oriented protocol. If all you want to do is sending
// and/or receiving packets of data over an USB its better and quite easy to
// use the CDC ACM data endponts with the usb package directly. Such approach
// will use the native packet/transaction oriented interface with the minimal
// overhead and maximum possible speed.
package usbserial

import (
	"embedded/rtos"
	"math/bits"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dtcm"
	"github.com/embeddedgo/imxrt/hal/usb"
)

// A Serial is a simple CDC ACM driver.
type Driver struct {
	d          *usb.Device
	interf     uint8
	txe, rxe   uint8
	wn         uint8
	rn, ri     int
	tda        *[3]usb.DTD  // tda[:2]   is for Write, tda[2] is for Read,
	donea      [3]rtos.Note // donea[:2] is for Write, donea[2] is for Read
	buf        []byte       // buf[:len] is for Read, buf[len:cap] is for Write
	lineCoding [7]byte
	autoFlush  bool
	writeSink  bool
	dtr        atomic.Int32
}

/*
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
*/

// NewDriver... rxe (host out), txe (host in).
// MaxPkt must be power of two and equal or multiple of the maximum packet size
// declared in the OUT endpoint descriptor used by this driver as Rx endpoint.
func NewDriver(d *usb.Device, interf uint8, rxe, txe int8, maxPkt int) *Driver {
	if bits.OnesCount(uint(maxPkt)) != 1 {
		panic("serial: maxPkt must be power of two")
	}
	s := &Driver{
		d:      d,
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

// Read implements io.Reader interface.
func (s *Driver) Read(p []byte) (n int, err error) {
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
	if !s.d.Prime(s.rxe, td, td) {
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

// SetAutoFlush enables/disables the AutoFlush mode. If AutoFlush is enabled,
// Write calls Flush before exit.
func (s *Driver) SetAutoFlush(af bool) {
	s.autoFlush = af
}

// SetWriteSink the WriteSink mode. Enabled WriteSink mode ensures that writes
// will not block if the USB serial device is not open for reading on the host
// side.
func (s *Driver) SetWriteSink(ws bool) {
	s.writeSink = ws
}

func writeDrop(s *Driver) bool {
	// BUG: naive approach, the Write/Flush can hang on corner cases
	return s.writeSink && (s.d.Config() == 0 || s.dtr.Load() == 0)
}

// Write implements io.Writer interface.
func (s *Driver) Write(p []byte) (n int, err error) {
	return s.WriteString(*(*string)(unsafe.Pointer(&p)))
}

// WriteString implements io.StringWriter interface.
func (s *Driver) WriteString(p string) (n int, err error) {
	if len(p) == 0 {
		return
	}
	if writeDrop(s) {
		return len(p), nil
	}
	dtcm := s.buf[len(s.buf):cap(s.buf)]
	nh := len(p) // unaligned head bytes, send through dtcm buffer
	nm := 0      // middle bytes, send directly from p, require rtos.DCacheFlush
	nt := 0      // unaligned tail bytes, send through dtcm buffer
	if nh > len(dtcm) {
		const align = dma.CacheLineSize - 1
		a := uintptr(unsafe.Pointer(unsafe.StringData(p)))
		nh = int(dma.CacheLineSize-a) & align
		nm = len(p) - nh
		nt = nm & align
		nm -= nt
		if nm != 0 {
			rtos.CacheMaint(rtos.DCacheFlush, unsafe.Pointer(unsafe.StringData(p[nh:])), nm)
		}
	}
	var (
		status uint8
		buf    unsafe.Pointer
		m      int
	)
	// The first (unaligned) part of p or the whole p if len(p) <= len(dtcm)
	wn := int(s.wn)
	if nh != 0 {
		m = nh
		goto useDTCM
	}
next:
	// The middle (aligned) part of p.
	if nm != 0 {
		buf = unsafe.Pointer(unsafe.StringData(p[n:]))
		m = nm
		nm = 0
		goto loop
	}
	// The last (unaligned) part of p.
	if nt != 0 {
		m = nt
		nt = 0
		goto useDTCM
	}
	// Done.
	s.wn = uint8(2 | wn&1)
	if s.autoFlush {
		err = s.Flush()
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
		if !s.d.Prime(s.txe, td, td) {
			goto error
		}
		if wn != 0 {
			td, done = &s.tda[(wn-1)&1], &s.donea[(wn-1)&1]
			if writeDrop(s) {
				return len(p), nil
			}
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
func (s *Driver) Flush() error {
	if writeDrop(s) {
		return nil
	}
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

var interfaces = make(map[uint8]*Driver)

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
	/*
		fmt.Printf("cdcACMGetLineCoding")
	*/
	return copy(cr.Data, s.lineCoding[:])
}

func setControlLineState(cr *usb.ControlRequest) int {
	s := interfaces[uint8(cr.Index)]
	if s == nil {
		return 0
	}
	s.dtr.Store(int32(cr.Value) & 1)
	/*
		fmt.Printf("cdcACMSetControlLineState:\r\n")
		fmt.Printf(" -interface: %d\r\n", cr.Index)
		fmt.Printf(" -DTR:       %d\r\n", cr.Value&1)
		fmt.Printf(" -RTS:       %d\r\n", cr.Value>>1&1)
	*/
	return 0
}
