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
	master.Setup(lpi2c.Std100k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)
	irq.LPI2C1.Enable(rtos.IntPrioLow, 0)

	//c := d.NewConn(prefix | e2e1e0)

	var buf [32]byte

loop:
	for page := 0; ; page++ {
		time.Sleep(2 * time.Second)
		a := page * 32
		mah := byte(a >> 8) // memory address high byte
		mal := byte(a)      // memory address low byte

		// Write string
		fmt.Println(page, "write:")
		master.WriteCmd(
			lpi2c.Start|slaveAddr<<1|wr,
			lpi2c.Send|int16(mah),
			lpi2c.Send|int16(mal),
		)
		s := fmt.Sprintf("> page %#x <", page)
		master.WriteString(s)
		master.WriteCmd(lpi2c.Stop)
		if err := master.Err(true); err != nil {
			fmt.Println("write error:", err)
			continue
		}
		for {
			master.WriteCmd(lpi2c.StartNACK | slaveAddr<<1 | wr)
			if err := master.Err(true); err != nil {
				e := err.(*lpi2c.MasterError)
				if e.SR&lpi2c.MasterErrFlags != lpi2c.MNDF {
					fmt.Println("wait error:", err)
					continue loop
				}
				break
			}
		}

		// Read string
		master.WriteCmd(
			lpi2c.Start|slaveAddr<<1|wr,
			lpi2c.Send|int16(mah),
			lpi2c.Send|int16(mal),
			lpi2c.Start|slaveAddr<<1|rd,
			lpi2c.Recv|int16(len(s)-1),
			lpi2c.Stop,
		)
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
