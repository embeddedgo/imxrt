// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Eeprom writes and reads the memory of the 24C64/128/256 I2C EEPROM. The
// difference to the less dense 24C0x EEPROMs is the use of 16 bit memory
// address instead of 8 bit one.
package main

import (
	"embedded/rtos"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/embeddedgo/device/bus/i2cbus"
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
	master = lpi2c.NewMaster(p, dma.Channel{})
	master.Setup(lpi2c.Std100k)
	master.UsePin(scl, lpi2c.SCL)
	master.UsePin(sda, lpi2c.SDA)
	irq.LPI2C1.Enable(rtos.IntPrioLow, 0)

	c := master.NewConn(slaveAddr)

	var buf [32]byte

loop:
	for page := 0; ; page++ {
		time.Sleep(2 * time.Second)
		a := page * 32
		addr := []byte{byte(a >> 8), byte(a)}

		s := fmt.Sprintf(">> %#x <<", page)

		fmt.Printf("\nWrite page %d: %s\n", page, s)
		c.Write(addr) // replace with c.WriteByte(addr) for 24C0x EEPROMs
		io.WriteString(c, s)
		err := c.Close()
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}

		// Wait for the end of write
		for {
			c.Write(nil)
			err := c.Close()
			if err == nil {
				break
			}
			if !errors.Is(err, i2cbus.ErrACK) {
				fmt.Println("wait error:", err)
				continue loop
			}
			fmt.Print(".")
		}
		fmt.Println(" done")

		c.Write(addr) // replace with c.WriteByte(addr) for 24C0x EEPROMs
		n, _ := c.Read(buf[:len(s)])
		err = c.Close()
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		fmt.Printf("Read %d bytes from page %d: %s\n", n, page, string(buf[:n]))
	}

}

//go:interrupthandler
func LPI2C1_Handler() {
	master.ISR()
}

func pr[T ~uint32](name string, v T) {
	print(name, ": ")
	for i := 32; i != 0; i-- {
		if i&7 == 0 && i != 32 {
			print("_")
		}
		print(v >> (i - 1) & 1)
	}
	print("\r\n")
}
