// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func pr[T ~uint32](name string, v T) {
	fmt.Printf(
		"%s %08b_%08b %08b_%08b\n", name,
		v>>24&0xff, v>>16&0xff, v>>8&0xff, v&0xff,
	)
}

func write(p *lpi2c.Periph, cmds ...int16) {
	for _, cmd := range cmds {
		for p.MSR.LoadBits(lpi2c.MTDF) == 0 {
		}
		p.MTDR.Store(cmd)
	}
}

func read(p *lpi2c.Periph, buf []byte) {
	for i := range buf {
		var v lpi2c.RDR
		for {
			v = p.MRDR.Load()
			if v&lpi2c.RXEMPTY == 0 {
				break
			}
		}
		buf[i] = byte(v)
	}
}

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPSPI3 driver
	p := lpi2c.LPI2C(1)
	d := lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	d.UsePin(scl, lpi2c.SCL)
	d.UsePin(sda, lpi2c.SDA)

	d.Setup(lpi2c.FastHS)

	time.Sleep(5 * time.Second)
	pr("MCR:   ", p.MCR.Load())
	pr("MSR:   ", p.MSR.Load())
	pr("MIER   ", p.MIER.Load())
	pr("MDER:  ", p.MDER.Load())
	pr("MCFGR0:", p.MCFGR0.Load())
	pr("MCFGR1:", p.MCFGR1.Load())
	pr("MCFGR2:", p.MCFGR2.Load())
	pr("MCFGR3:", p.MCFGR3.Load())
	pr("MDMR:  ", p.MDMR.Load())
	pr("MCCR0: ", p.MCCR0.Load())
	pr("MCCR1: ", p.MCCR1.Load())
	pr("MFCR:  ", p.MFCR.Load())

	const (
		addr = 0b1001000 << 1 // address when the ADDR pin is connected to GND
		wr   = 0              // write transaction
		rd   = 1              // read transaction
	)

	var buf [2]byte

	for i := 0; ; i++ {
		write(
			p,
			lpi2c.StartNACK|0b0000_1000, // switch to High Speed mode
			lpi2c.StartHS|addr|wr,
			lpi2c.Send|1,
			lpi2c.StartHS|addr|rd,
			lpi2c.Recv|int16(len(buf)-1),
			lpi2c.Stop,
		)
		read(p, buf[:])
		fmt.Printf("%d: %08b %08b\n", i, buf[0], buf[1])

		time.Sleep(time.Second)
	}

}
