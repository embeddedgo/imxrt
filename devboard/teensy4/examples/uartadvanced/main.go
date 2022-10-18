// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Uartadvanced demonstrates how to use the LPUART peripheral including all
// details like interrupts and DMA. In practice, using DMA with small
// 64-character Rx buffer, small portions of Tx data and slow 115200 speed is
// rather overkill, mainly because of required cache maintenance operations.
package main

import (
	"embedded/rtos"
	"fmt"
	"time"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dma/dmairq"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
)

var u *lpuart.Driver

func main() {
	// Used IO pins
	tx := pins.P24
	rx := pins.P25

	// Enable DMA controller and allocate two channels for the LPUART driver.
	d := dma.DMA(0)
	d.EnableClock(true)
	rxdma := d.AllocChannel(false)
	txdma := d.AllocChannel(false)

	// Setup LPUART driver
	u = lpuart.NewDriver(lpuart.LPUART(1), rxdma, txdma)
	u.Setup(lpuart.Word8b, 115200)
	u.UsePin(rx, lpuart.RXD)
	u.UsePin(tx, lpuart.TXD)
	irq.LPUART1.Enable(rtos.IntPrioLow, 0)
	dmairq.SetISR(rxdma, u.RxDMAISR)
	dmairq.SetISR(txdma, u.TxDMAISR)

	// Enable both directions
	u.EnableRx(64) // 64-character ring buffer
	u.EnableTx()

	t0 := time.Now()
	u.WriteString(akermanianSteppe)
	t1 := time.Now()
	sps := float64(len(akermanianSteppe)) * float64(time.Second) / float64(t1.Sub(t0))
	bps := sps * 8
	baud := sps * 10 // 1b start + 8b data + 1b stop
	fmt.Fprintf(
		u, "%d bytes @ %.0f b/s (%.0f baud)\r\n",
		len(akermanianSteppe), bps, baud,
	)

	// Print received data showing reading chunks
	buf := make([]byte, 80)
	for {
		n, err := u.Read(buf)
		if err != nil {
			fmt.Fprintf(u, "error: %v\r\n", err)
			continue
		}
		fmt.Fprintf(u, "%d: %s\r\n", n, buf[:n])
		if buf[0] == '*' {
			break
		}
	}

	u.WriteString("\r\n\r\n*** Receive data as 16-bit words ***\r\n\r\n")

	buf16 := make([]uint16, 80)
	for {
		n, err := u.Read16(buf16)
		if err != nil {
			fmt.Fprintf(u, "error: %v\r\n", err)
			continue
		}
		fmt.Fprintf(u, "%d: %x\r\n   ", n, buf16[:n])
		for _, w := range buf16[:n] {
			u.WriteByte(byte(w))
		}
		u.WriteString("\r\n")
	}
}

//go:interrupthandler
func LPUART1_Handler() {
	u.ISR()
	leds.User.Toggle() // visualize LPUART interrupts
}

const akermanianSteppe = "\r\n" +
	"The Akkerman Steppe poem by Adam Mickiewicz translated to\r\n" +
	"English by Leo Yankevich.\r\n" +
	"\r\n" +
	"I launch myself across the dry and open narrows,\r\n" +
	"My carriage plunging into green as if a ketch,\r\n" +
	"Floundering through the meadow flowers in the stretch.\r\n" +
	"I pass an archipelago of coral yarrows.\r\n" +
	"\r\n" +
	"It's dusk now, not a road in sight, nor ancient barrows.\r\n" +
	"I look up at the sky and look for stars to catch.\r\n" +
	"There distant clouds glint-there tomorrow starts to etch;\r\n" +
	"The Dnieper glimmers; Akkerman's lamp shines and harrows.\r\n" +
	"\r\n" +
	"I stand in stillness, hear the migratory cranes,\r\n" +
	"Their necks and wings beyond the reach of preying hawks;\r\n" +
	"Hear where the sooty copper glides across the plains,\r\n" +
	"\r\n" +
	"Where on its underside a viper writhes through stalks.\r\n" +
	"Amid the hush I lean my ears down grassy lanes\r\n" +
	"And listen for a voice from home. Nobody talks.\r\n" +
	"\r\n\r\n"
