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
	wcd     lpspi.TCR
	rc      lpspi.TCR
	rd      lpspi.TCR
	readBug bool
}

func presc(base, clk int) lpspi.TCR {
	x := (base+clk-1)/clk - 1
	if x < 0 {
		x = 0
	}
	x = bits.Len(uint(x))
	if x > 7 {
		x = 7
	}
	return lpspi.TCR(x) << lpspi.PRESCALEn
}

// NewLPSPI returns new LPSPI based implementation of tftdrv.DCI. User must
// provide a configured LPSPI master driver, a DC pin, the required SPI mode
// (CPOL,CPHA) amd the maximum write, read clock speeds according to the display
// controller specification. Note that the maximum speed may be limited by th LPSPI peripheral, the bus topology or
// the specific display design.
func NewLPSPI(drv *lpspi.Master, dc iomux.Pin, mode lpspi.TCR, rclkHz, wclkHz int) *LPSPI {
	dc.Setup(iomux.Drive7)
	dcio := gpio.UsePin(dc, false)
	dcio.Port().EnableClock(true)
	dcio.Clear()
	dcio.SetDirOut(true)
	mode &= lpspi.WIDTH | lpspi.BYSW | lpspi.LSBF | lpspi.TPCS | lpspi.CPHA | lpspi.CPOL
	f := drv.BaseFreqHz()
	wpre := presc(f, wclkHz)
	rpre := presc(f, rclkHz)
	return &LPSPI{
		spi: drv,
		dc:  dcio,
		wcd: mode | lpspi.CONT | lpspi.CONTC | lpspi.RXMSK | wpre,
		rc:  mode | lpspi.CONT | lpspi.CONTC | lpspi.RXMSK | rpre,
		rd:  mode | lpspi.CONT | lpspi.CONTC | rpre,
	}
}

func (dci *LPSPI) Driver() *lpspi.Master { return dci.spi }
func (dci *LPSPI) Err(clear bool) error  { return nil }

func (dci *LPSPI) Cmd(p []byte, dataMode int) {
	// LPSPI doesn't allow to change the SCK prescaler when CS is asserted but
	// CS must be asserted for the whole read transaction. Write transactions
	// aren't so restrictive but we handle them the same way.
	cmd := dci.wcd
	if dataMode == 2 {
		cmd = dci.rc
	}
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(cmd&^lpspi.CONTC, 8)
	// We must break the abstracion provided by the SPI driver because of the
	// asynchronous DC signal that must be synchronized in some way.
	sp := spi.Periph()
	for sp.SR.LoadBits(lpspi.MBF) != 0 {
		// Waiting for the previous transfer to complete before asserting DC.
	}
	dci.dc.Clear() // assert DC to select the command mode
	// Probably all real commands are 1 byte in size but we support long
	// commands or multiple commands in p.
	for _, b := range p {
		sp.SR.Store(lpspi.WCF)  // reset the Word Complete Flag
		sp.TDR.Store(uint32(b)) // send the command byte
		for sp.SR.LoadBits(lpspi.WCF) == 0 {
			// DC must be low until the last bit of the last byte so we have to
			// wait for the Word Complete Flag after each byte sent.
		}
	}
	dci.dc.Set() // deassert DC at the end of command
}

func (dci *LPSPI) End() {
	spi := dci.spi
	spi.Disable() // the simplest way to deasert CS when TCR.CONT is set
	if dci.readBug {
		dci.readBug = false
		spi.ReadWord()
	}
}

func (dci *LPSPI) WriteBytes(p []uint8) {
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(dci.wcd, 8)
	spi.Write(p)
}

func (dci *LPSPI) WriteString(s string) {
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(dci.wcd, 8)
	spi.WriteString(s)
}

func (dci *LPSPI) WriteByteN(b byte, n int) {
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(dci.wcd, 8)
	for ; n != 0; n-- {
		spi.WriteWord(uint32(b))
	}
}

func (dci *LPSPI) WriteWords(p []uint16) {
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(dci.wcd, 16)
	spi.Write16(p)
}

func (dci *LPSPI) WriteWordN(w uint16, n int) {
	spi := dci.spi
	spi.Enable()
	spi.WriteCmd(dci.wcd, 16)
	for ; n != 0; n-- {
		spi.WriteWord(uint32(w))
	}
}

func (dci *LPSPI) ReadBytes(p []byte) {
	spi := dci.spi
	spi.Enable()
	if !dci.readBug {
		// Handle the read bug (see description of lpspi.Master.WriteCmd).
		// This workaround causes to read an excess byte from the display but
		// it seems to not be a problem for all known displays.
		dci.readBug = true
		spi.WriteCmd(dci.rd, 8)
		spi.WriteWord(0)
	}
	for i := range p {
		spi.WriteWord(0)
		p[i] = byte(spi.ReadWord())
	}
}
