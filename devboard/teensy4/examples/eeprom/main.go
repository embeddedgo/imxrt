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

func write(p *lpi2c.Periph, cmds ...lpi2c.MTDR) {
	for _, cmd := range cmds {
		for p.MSR.LoadBits(lpi2c.MTDF) == 0 {
		}
		p.MTDR.Store(cmd)
	}
	fmt.Printf("MFSR: %#x\n", p.MFSR.Load())
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

	time.Sleep(5 * time.Second)

	// Setup LPSPI3 driver
	p := lpi2c.LPI2C(1)
	d := lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	d.UsePin(scl, lpi2c.SCL)
	d.UsePin(sda, lpi2c.SDA)

	d.Setup(lpi2c.Fast400k)

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

	time.Sleep(5 * time.Second)

	fmt.Println("Go!")

	const (
		prefix = 0x5 << 4 // 0b1010 address prefix
		a2a1   = 0 << 2   // address pins
		p0     = 0 << 1   // page
		wr     = 0        // write transaction
		rd     = 1        // read transaction
		addr   = 0xf0     // address in page
	)

	for {
		var buf [16]byte
		write(
			p,
			lpi2c.Start|prefix|a2a1|p0|wr,
			0xf0,
			//lpi2c.Recv|lpi2c.MTDR(len(buf)-1),
			lpi2c.Stop,
		)

		for {
			time.Sleep(2 * time.Second)
			pr("MSR:   ", p.MSR.Load())
			fmt.Printf("MFSR: %#x\n", p.MFSR.Load())
		}

		read(p, buf[:])
		fmt.Printf("%s\n", buf[:])

		time.Sleep(time.Second)
	}

}