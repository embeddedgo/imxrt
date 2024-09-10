// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"embedded/rtos"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

// A Master is a driver for the LPI2C peripheral to perform a master access to
// an I2C bus. It provides two interfaces.
//
// A low-level one for direct interraction with the Data / Command FIFOs of the
// underlying LPI2C peripheral.
//
// Example (error checking omitted for clarity):
//
//	d.WriteCmd(
//		lpi2c.Start|eepromAddr<<1|wr,
//		lpi2c.Send|int16(memAddr),
//		lpi2c.Start|eepromAddr<<1|rd,
//		lpi2c.Recv|int16(len(buf) - 1),
//		lpi2c.Stop,
//	)
//	d.ReadData(buf)
//
// Write methods in the low-level interface are asynchronous, that is, they may
// return before the end of writting all commands/data to the internal FIFO.
// Therefore you cannot modify the data/command buffer pass to last write until
// calling the Flush method. The returned errors are also asynchronous (i.e.
// event the ReadData method may return an error caused by prievious write).
//
// The second interface is a connection oriented one that implements
// io.ReadWriteCloser.
//
// Example (error checking omitted for clarity):
//
//	conn := d.NewConn(eepromAddr)
//	conn.WriteByte(memAaddr)
//	conn.Read(buf)
//	conn.Close()
//
// Both interfaces may be used concurently by multiple goroutines but in such
// a case users of the low-level interface must gain an exclusive access to the
// driver using the embedded mutex.
type Master struct {
	sync.Mutex // use with the low-level interface to share the driver

	p *Periph

	wcmds *int16
	wdata *byte
	wi    int32 // ISR cannot alter the above pointers so it alters wi instead
	wn    int32
	wdone rtos.Note

	rdata *byte
	ri    int32 // ISR cannot alter the above pointer so it alters ri instead
	rn    int32
	rdone rtos.Note

	wdma dma.Channel
	rdma dma.Channel
}

// NewMaster returns a new master-mode driver for p. If valid DMA channels are
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, rdma, wdma dma.Channel) *Master {
	return &Master{p: p, rdma: rdma, wdma: wdma}
}

// Periph returns the underlying LPSPI peripheral.
func (d *Master) Periph() *Periph {
	return d.p
}

// Timing constants.
//
// sclClk = clk / ((CLKHI + CLKLO + 2 + sclLatency) << divN)
//
// sclLatency = roundDown((2 + FILTSCL) >> divN)
const (
	clk  = 60_000_000 // peripheral clock (PLL_USB1 / 8)
	div2 = 1          // divide the 60 MHz clock by 2 (30 MHz)
	div4 = 2          // divide the 60 MHz clock by 4 (15 MHz)
	div8 = 3          // divide the 60 MHz clock by 8 (7.5 MHz)

	// Values copied from Table 47-5. LPI2C Example Timing Configurations.
	fahs = 2<<30 | 17<<SETHOLDn | 40<<CLKLOn | 31<<CLKHIn | 8<<DATAVDn
	plhs = 2<<30 | 7<<SETHOLDn | 15<<CLKLOn | 11<<CLKHIn | 2<<DATAVDn
	hs   = 4<<SETHOLDn | 4<<CLKLOn | 2<<CLKHIn | 1<<DATAVDn

	// The above values divided by 2 with small corrections to work with div4.
	fa = 2<<30 | 9<<SETHOLDn | 20<<CLKLOn | 16<<CLKHIn | 4<<DATAVDn
	pl = 2<<30 | 4<<SETHOLDn | 8<<CLKLOn | 5<<CLKHIn | 1<<DATAVDn

	// Values to obtain the minimal possible sclClk for any div.
	sl = 15<<30 | 31<<SETHOLDn | 63<<CLKLOn | 63<<CLKHIn | 15<<DATAVDn

	timingSlow   = div8<<6 | hs<<34 | sl
	timingStd    = div4<<6 | hs<<34 | sl
	timingFast   = div4<<6 | hs<<34 | fa
	timingPlus   = div4<<6 | hs<<34 | pl
	timingFastHS = div2<<6 | hs<<34 | fahs
	timingPlusHS = div2<<6 | hs<<34 | plhs

	stuckBusTimeout = 40 // ms (TI "I2C Stuck Bus: Prevention and Workarounds")
)

// Speed encodes the timing configuration that determines the maximum
// communication speed (the actual speed depends also on the SCL rise time).
type Speed uint64

const (
	Slow50k    Speed = timingSlow   //  ≤58 kb/s (slow)     and 0.83 Mb/s HS
	Std100k    Speed = timingStd    // ≤114 kb/s (standard) and 1.65 Mb/s HS
	Fast400k   Speed = timingFast   // ≤400 kb/s (Fast)     and 1.65 Mb/s HS
	FastPlus1M Speed = timingPlus   //   ≤1 Mb/s (Fast+)    and 1.65 Mb/s HS
	FastHS     Speed = timingFastHS // ≤400 kb/s (Fast)     and 3.33 Mb/s HS
	FastPlusHS Speed = timingPlusHS //   ≤1 Mb/s (Fast+)    and 3.33 Mb/s HS
)

func (d *Master) Setup(sp Speed) {
	p := d.p
	p.EnableClock(true)
	p.MCR.Store(MRST)
	p.MCR.Store(0)
	p.MCCR0.Store(MCCR(sp) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR1.Store(MCCR(sp>>34) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	pre := MCFGR1(sp) >> 6 & 3 // max. supported MPRESCALE is 3
	p.MCFGR1.Store(MAUTOSTOP | pre<<MPRESCALEn)
	gf := MCFGR2(sp>>30) & 0xf // the used encoding supports MFILT <= 15
	bi := (MCFGR2(sp)>>CLKLOn&63 + MCFGR2(sp)>>SETHOLDn&63 + 2) * 2
	p.MCFGR2.Store(gf<<MFILTSDAn | gf<<MFILTSCLn | bi<<MBUSIDLEn)
	p.MCFGR3.Store(clk * stuckBusTimeout / 1000 / 256 >> pre << PINLOWn)
	p.MCR.Store(MEN)
	// We use the default MFCR=0. It may lead to short discontinuities in the
	// data flow but we value more the reduced IRQ and DRQ frequency.
}

type Error struct {
	SR MSR
}

func (e *Error) Error() string {
	var a [4]string
	es := a[:0:4]
	if e.SR&MNDF != 0 {
		es = append(es, "NACK")
	}
	if e.SR&MALF != 0 {
		es = append(es, "Arbitr")
	}
	if e.SR&MFEF != 0 {
		es = append(es, "FIFO")
	}
	if e.SR&MPLTF != 0 {
		es = append(es, "PinLow")
	}
	return "lpi2c: " + strings.Join(es, ",")
}

// WriteCmd starts writing commands into the Tx FIFO in the background using
// interrupts and/or DMA. WriteCmd is no-op if len(cmd) == 0 so in such case
// it doesn't wait for the end of previous write.
func (d *Master) WriteCmd(cmd ...int16) {
	if len(cmd) == 0 {
		return
	}
	if d.wdma.IsValid() && len(cmd)*2 >= 2*dma.CacheLineSize {
		ptr := unsafe.Pointer(&cmd[0])
		p, q := dma.AlignOffsets(ptr, uintptr(len(cmd)*2))
		dmaStart := int(p / 2)
		dmaEnd := int(q / 2)
		dmaPtr := unsafe.Add(ptr, p)
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			masterWrite(d, ptr, dmaStart, 1)
		}
		masterWriteDMA(d, dmaPtr, dmaN, 1)
		if dmaEnd == len(cmd) {
			return
		}
		cmd = cmd[dmaEnd:]
	}
	masterWrite(d, unsafe.Pointer(&cmd[0]), len(cmd), 1)
}

func (d *Master) Err(clear bool) error {
	p := d.p
	if sr := p.MSR.Load(); sr&errFlags != 0 {
		if clear {
			p.MCR.SetBits(MRRF | MRTF) // clear FIFOs
			p.MSR.Store(sr & errFlags) // clear the error flags
		}
		return &Error{sr} // return all flags for the better context
	}
	return nil
}

const (
	txFIFOLen = 4
	rxFIFOLen = 4

	errFlags = MNDF | MALF | MFEF | MPLTF
)

func masterWrite(d *Master, ptr unsafe.Pointer, n int, lsz uint) {
	if d.wn != 0 {
		// Wait for the ISR to end the previously sheduled transfer.
		d.wdone.Sleep(-1)
		d.wdone.Clear()
		d.wcmds = nil
		d.wdata = nil
		d.wn = 0
	}
	p := d.p
	if p.MSR.Load()&errFlags != 0 {
		return
	}
	// Avoid interrupts if there is a free space in the FIFO.
	i := 0
	if lsz == 0 {
		data := unsafe.Slice((*byte)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOLen {
			if i >= len(data) {
				return
			}
			p.MTDR.Store(int16(data[i]))
			i++
		}
		d.wdata = &data[i]
	} else {
		cmds := unsafe.Slice((*int16)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOLen {
			if i >= len(cmds) {
				return
			}
			p.MTDR.Store(cmds[i])
			i++
		}
		d.wcmds = &cmds[i]
	}
	// The remaining data/commands will be writtend to the FIFO by the ISR.
	d.wi = 0
	atomic.StoreInt32(&d.wn, int32(n-i))
	p.MIER.SetBits(MTDF | errFlags)
}

func masterWriteDMA(d *Master, ptr unsafe.Pointer, n int, lsz uint) {

}

//go:nosplit
//go:nowritebarrierrec
func (d *Master) ISR() {
	p := d.p
	sr := p.MSR.Load()
	if sr&errFlags != 0 {
		// Goroutnies first set d.txn/d.rxn and next set MIER so the ISR must
		// clear MIER before checking d.txn/d.rxn to avoid goroutine stall.
		p.MIER.Store(0)
		// Wake up goroutines. In the meantime they may reenable interrupts so
		// set d.wn/d.rn to -1 as mark that the Wakeup was already called.
		if atomic.LoadInt32(&d.wn) > 0 {
			d.wn = -1
			d.wdone.Wakeup()
		}
		if atomic.LoadInt32(&d.rn) > 0 {
			d.rn = -1
			d.rdone.Wakeup()
		}
		return
	}
	if sr&MTDF != 0 {
		// Because MFCR.TXWATER=0 (see Setup) the FIFO is now empty.
		i, n := d.wi, d.wn
		m := min(i+txFIFOLen, n)
		if d.wdata != nil {
			for _, b := range unsafe.Slice(d.wdata, n)[i:m] {
				p.MTDR.Store(int16(b))
			}
		} else {
			for _, cmd := range unsafe.Slice(d.wcmds, n)[i:m] {
				p.MTDR.Store(cmd)
			}
		}
		d.wi = m
		if m == n {
			// No more commands or data to write.
			p.MIER.Store(0) // disable all interrupts and fix it in a moment
			if atomic.LoadInt32(&d.rn) > 0 {
				// There is a pending read, reenable read interrupts
				p.MIER.Store(MRDF | errFlags)
			}
			d.wdone.Wakeup()
		}
	}
}
