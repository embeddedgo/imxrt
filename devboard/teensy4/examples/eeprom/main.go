// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Eeprom writes and read the memory of the 24C64 I2C EEPROM (64 Kb = 8192 B =
// 256 pages * 32 B/page).
package main

import (
	"embedded/rtos"
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpi2c"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

const (
	prefix    = 0b1010 // address prefix
	e2e1e0    = 0      // address pins
	slaveAddr = prefix<<3 | e2e1e0
	wr        = 0 // write transaction
	rd        = 1 // read transaction
)

var master *lpi2c.Master

func main() {
	// Used IO pins
	sda := pins.P18 // AD_B1_01
	scl := pins.P19 // AD_B1_00

	// Setup LPI2C driver
	p := lpi2c.LPI2C(1)
	master = lpi2c.NewMaster(p, dma.Channel{}, dma.Channel{})
	master.Setup(lpi2c.Slow50k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)
	irq.LPI2C1.Enable(rtos.IntPrioLow, 0)

	//c := d.NewConn(prefix | e2e1e0)

	var buf [32]byte

	//	var log [64]uint16

loop:
	for page := 0; ; page++ {
		time.Sleep(2 * time.Second)
		a := page * 32
		mah := byte(a >> 8) // memory address high byte
		mal := byte(a)      // memory address low byte

		// Write string
		fmt.Println(page, "write:")
		master.WriteCmds([]int16{
			lpi2c.Start | slaveAddr<<1 | wr,
			lpi2c.Send | int16(mah),
			lpi2c.Send | int16(mal),
		})
		s := fmt.Sprintf("> str %#x <    .", page)
		master.WriteString(s)
		master.WriteCmd(lpi2c.Stop)
		if err := master.Err(true); err != nil {
			fmt.Println("write error:", err)
			continue
		}

		i := 0
		p := master.Periph()
		for p.MSR.LoadBits(lpi2c.MasterErrFlags|lpi2c.MTDF) == 0 {
		}

		time.Sleep(time.Second)
		master.WriteCmd(lpi2c.Start | slaveAddr<<1 | wr)
		for {
			p.MSR.Store(lpi2c.MEPF)
			master.WriteCmd(lpi2c.Start | slaveAddr<<1 | wr)
			for p.MSR.LoadBits(lpi2c.MasterErrFlags|lpi2c.MEPF) == 0 {
			}
			err := master.Err(true)
			if err == nil {
				break
			}
			sr := err.(*lpi2c.MasterError).SR
			if sr&lpi2c.MasterErrFlags != lpi2c.MNDF {
				fmt.Println("wait error:", err)
				continue loop
			}
			i++
		}
		/*
			fmt.Println("wait log: ", i)
			for _, u16 := range log {
				fmt.Printf("%d: %08b %08b\n", i, byte(u16>>8), byte(u16))
			}
			sr := p.MSR.Load()
			u16 := uint16(sr&0x3fff | sr>>10&0xc000)
			fmt.Printf(
				".: %08b %08b %d\n", byte(u16>>8), byte(u16),
				p.MFSR.LoadBits(lpi2c.TXCOUNT)>>lpi2c.TXCOUNTn,
			)
			if err := master.Err(true); err != nil {
				fmt.Println("wait1 error:", err)
				continue
			}
			time.Sleep(2 * time.Second)
		*/

		// Read string
		master.WriteCmds([]int16{
			lpi2c.Start | slaveAddr<<1 | wr,
			lpi2c.Send | int16(mah),
			lpi2c.Send | int16(mal),
			lpi2c.Start | slaveAddr<<1 | rd,
			lpi2c.Recv | int16(len(s)-1),
			lpi2c.Stop,
		})
		master.Read(buf[:len(s)])
		if err := master.Err(true); err != nil {
			fmt.Println("read error:", err)
			continue
		}
		fmt.Println(page, "read:", string(buf[:len(s)]))
	}

}

//go:interrupthandler
func LPI2C1_Handler() {
	master.ISR()
}
