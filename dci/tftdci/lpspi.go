// Copyright 2024 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tftdci

import (
	"math/bits"

	"github.com/embeddedgo/imxrt/hal/gpio"
	"github.com/embeddedgo/imxrt/hal/iomux"
	"github.com/embeddedgo/imxrt/hal/lpspi"
)

// LPSPI is an implementation of the display/tft.DCI interface that uses an
// LPSPI peripheral to communicate with the display in what is known as 4-line
// serial mode.
type LPSPI struct {
	spi     *lpspi.Master
	dc      gpio.Bit
	wcmd    lpspi.TCR
	rcmd    lpspi.TCR
	started bool
}

// NewLPSPI returns new LPSPI based implementation of tftdrv.DCI. User must
// provide a configured LPSPI master driver, a DC pin, the required SPI mode
// (CPOL,CPHA) amd the maximum write, read clock speeds according to the display
// controller specification. Note that the maximum speed may be limited by th LPSPI peripheral, the bus topology or
// the specific display design.
func NewLPSPI(drv *lpspi.Master, dc iomux.Pin, mode lpspi.TCR, wclkHz, rclkHz int) *LPSPI {
	dc.Setup(iomux.Drive7)
	dcio := gpio.UsePin(dc, false)
	dcio.Port().EnableClock(true)
	dcio.Clear()
	dcio.SetDirOut(true)
	mode &= lpspi.WIDTH | lpspi.BYSW | lpspi.LSBF | lpspi.TPCS | lpspi.CPHA | lpspi.CPOL
	f := drv.BaseFreqHz()
	x := (f+wclkHz-1)/wclkHz - 1
	if x < 0 {
		x = 0
	}
	prew := lpspi.TCR(bits.Len(uint(x))) << lpspi.PRESCALEn
	x = (f+wclkHz-1)/rclkHz - 1
	if x < 0 {
		x = 0
	}
	prer := lpspi.TCR(bits.Len(uint(x))) << lpspi.PRESCALEn
	return &LPSPI{
		spi:  drv,
		dc:   dcio,
		wcmd: mode | lpspi.CONT | lpspi.RXMSK | prew,
		rcmd: mode | lpspi.CONT | lpspi.TXMSK | prer,
	}
}

func (dci *LPSPI) Driver() *lpspi.Master { return dci.spi }
func (dci *LPSPI) Err(clear bool) error  { return nil }

func (dci *LPSPI) Cmd(p []byte) {
	dci.dc.Clear()
	dci.spi.WriteCmd(dci.wcmd, 8)
	dci.spi.Write(p)
	fsr := &dci.spi.Periph().FSR
	for fsr.LoadBits(lpspi.TXCOUNT) != 0 {
	}
	dci.dc.Set()
}

func (dci *LPSPI) End() {
	dci.spi.WriteCmd(0, 8)
}

func (dci *LPSPI) WriteBytes(p []uint8) {
	dci.spi.WriteCmd(dci.wcmd, 8)
	dci.spi.Write(p)
}

func (dci *LPSPI) WriteString(s string) {
	dci.spi.WriteCmd(dci.wcmd, 8)
	dci.spi.WriteString(s)
}

func (dci *LPSPI) WriteByteN(b byte, n int) {
	dci.spi.WriteCmd(dci.wcmd, 8)
	for n != 0 {
		dci.spi.WriteWord(uint32(b))
	}
}

func (dci *LPSPI) WriteWords(p []uint16) {
	dci.spi.WriteCmd(dci.wcmd, 16)
	dci.spi.Write16(p)
}

func (dci *LPSPI) WriteWordN(w uint16, n int) {
	dci.spi.WriteCmd(dci.wcmd, 16)
	for n != 0 {
		dci.spi.WriteWord(uint32(w))
	}
}

func (dci *LPSPI) ReadBytes(p []byte) {
	for i := range p {
		dci.spi.WriteCmd(dci.rcmd, 8)
		p[i] = byte(dci.spi.ReadWord())
	}
}
