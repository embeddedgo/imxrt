// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpi2c

import (
	"embedded/rtos"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

// A Master is a driver for the LPI2C peripheral. It provides two kinds of
// interfaces to communicate with slave devices on the I2C bus.
//
// The first interface is a low-level one. It provides a set of methods to
// directly interract with the Data / Command FIFOs of the underlying LPI2C
// peripheral.
//
// Example:
//
//	const wr, rd = 0, 1
//	d.WriteCmds([]int16{
//		lpi2c.Start|eepromAddr<<1|wr,
//		lpi2c.Send|int16(memAddr),
//		lpi2c.Start|eepromAddr<<1|rd,
//		lpi2c.Recv|int16(len(buf) - 1),
//		lpi2c.Stop,
//	})
//	d.ReadBytes(buf)
//	if err := d.Err(true); err != nil {
//
// Write methods in the low-level interface are asynchronous, that is, they may
// return before all commands/data will be written to the FIFO. Therefore you
// must not modify the data/command buffer passed to the last write method until
// the return of the Flush method or another write method.
//
// The read/write methods doesn't return errors. There is an Err method that
// allow to check and reset the LPI2C error flags at a convenient time. Even if
// you call Err after every method call the returned error is still asynchronous
// due to the asynchronous nature of the write methods and the delayed execution
// of commands by the LPI2C peripheral itself. You can use Wait to synchronise
// things but it seems that only the MSDF flag (wait for the Stop Condition) can
// be used to synchronize errors.
//
// The second interface is a connection oriented one that implements the
// i2cbus.Conn interface.
//
// Example:
//
//	c := d.NewConn(eepromAddr)
//	c.WriteByte(memAaddr)
//	c.Read(buf)
//	err := c.Close()
//	if err != nil {
//
// Both interfaces may be used concurently by multiple goroutines but in such
// a case users of the low-level interface must gain an exclusive access to the
// driver using the embedded mutex and wait for the Stop Condition before
// unlocking the Master.
type Master struct {
	sync.Mutex

	name string
	p    *Periph

	wbuf int16
	id   uint8

	cmd   bool
	wdata unsafe.Pointer
	wi    int32 // ISR cannot alter the above pointers so it alters wi instead
	wn    int32
	wdone rtos.Note

	rdata *byte
	ri    int32 // ISR cannot alter the above pointer so it alters ri instead
	rn    int32
	rdone rtos.Note

	dma dma.Channel
}

// NewMaster returns a new master-mode driver for p. If valid DMA channel is
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, dma dma.Channel) *Master {
	return &Master{
		name: "LPI2C" + string(rune('0'+num(p))),
		p:    p,
		dma:  dma,
	}
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

	stuckBusTimeout = 40 // ms, see "I2C Stuck Bus: Prevention and Workarounds"
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

// All dma.Mux slot constants are less than 128 so we can use the string
// conversion to group them in a constant array.
const dmaSlots = "" +
	string(rune(dma.LPI2C1)) +
	string(rune(dma.LPI2C2)) +
	string(rune(dma.LPI2C3)) +
	string(rune(dma.LPI2C4))

// Setup resets and configures the underlying LPI2C pripheral to operate in the
// master mode with the given speed.
func (d *Master) Setup(sp Speed) {
	p := d.p
	p.EnableClock(true)
	p.MCR.Store(MRST)
	p.MCR.Store(0)
	p.MCCR0.Store(MCCR(sp) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	p.MCCR1.Store(MCCR(sp>>34) & (DATAVD | SETHOLD | CLKHI | CLKLO))
	pre := MCFGR1(sp) >> 6 & 3 // max. supported MPRESCALE is 3
	p.MCFGR1.Store(pre << MPRESCALEn)
	gf := MCFGR2(sp>>30) & 0xf // the used encoding supports MFILT <= 15
	bi := (MCFGR2(sp)>>CLKLOn&63 + MCFGR2(sp)>>SETHOLDn&63 + 2) * 2
	p.MCFGR2.Store(gf<<MFILTSDAn | gf<<MFILTSCLn | bi<<MBUSIDLEn)

	// Don't use Pin Low Timeout because it detects the low state of the SCK pin
	// held low by the LPI2C state machine after Start command when it waits for
	// more data/command in the Tx FIFO or a free space in the Rx FIFO.
	//p.MCFGR3.Store(clk * stuckBusTimeout / 1000 / 256 >> pre << PINLOWn)

	if dc := d.dma; dc.IsValid() {
		dc.DisableReq()
		dc.DisableErrInt()
		dc.ClearInt()
		dc.SetMux(dma.Mux(dmaSlots[num(d.p)]) | dma.En)
	}
	p.MCR.Store(MEN)
}

const (
	rxFIFOCap = 4
	txFIFOCap = 4
)

// Flush waits until all commands/data passed to the driver have been consumed
// (in other words, it makes the previous write operation synchronous). You must
// call Flush or write new to enusre the Master stops referencing previously
// written data (to reuse memory or make it available for garbage collection).
// Return from Flush doesn't mean that all data were sent on the bus (there may
// be even full Tx FIFO not handled yet, see Wait).
func (d *Master) Flush() {
	if d.wdata != nil {
		d.wdone.Sleep(-1)
		d.wdone.Clear()
		d.wdata = nil
	}
}

func masterWrite(d *Master, ptr unsafe.Pointer, n int, cmd bool) {
	p := d.p
	if p.MSR.LoadBits(MasterErrFlags) != 0 {
		return
	}
	// To speed things up we try to fill the FIFO in thread mode. As thread code
	// may be interrupted at any time we check TXCOUNT every iteration instead
	// of write as fast as possible the txFIFOCap-TXCOUNT commands/bytes.
	i := 0
	if !cmd {
		data := unsafe.Slice((*byte)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOCap {
			p.MTDR.Store(int16(data[i]))
			if i++; i == len(data) {
				return
			}
		}
	} else {
		cmds := unsafe.Slice((*int16)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOCap {
			p.MTDR.Store(cmds[i])
			if i++; i == len(cmds) {
				return
			}
		}
	}
	// The remaining data/commands will be writtend to the FIFO by the ISR.
	d.cmd = cmd
	d.wdata = ptr
	d.wi = int32(i)
	atomic.StoreInt32(&d.wn, int32(n))
	p.MIER.Store(MTDF | MasterErrFlags) // race with ISR (clear)
}

const dmaMaxMajorIter = 1<<dma.ELINKn - 1 // = 32767

func masterWriteDMA(d *Master, ptr unsafe.Pointer, n int) {
	p := d.p
	if p.MSR.LoadBits(MasterErrFlags) != 0 {
		return
	}
	rtos.CacheMaint(rtos.DCacheFlush, ptr, n)
	const dmaChunk = 4 // eqals 1 x S32b and 4 x D8b, <=txFIFOLen
	tcd := dma.TCD{
		SADDR:       ptr,
		SOFF:        dmaChunk,
		ATTR:        dma.S32b | dma.D8b,
		ML_NBYTES:   dmaChunk,
		DADDR:       unsafe.Pointer(p.MTDR.Addr()),
		ELINK_CITER: dmaMaxMajorIter,
		ELINK_BITER: dmaMaxMajorIter,
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	p.MDER.Store(TDDE) // clears RDDE
	dma := d.dma
	dma.WriteTCD(&tcd)
	tcdio := dma.TCD()
	n /= dmaChunk
	for {
		m := n
		if m > dmaMaxMajorIter {
			m = dmaMaxMajorIter
		}
		n -= m
		if m != dmaMaxMajorIter {
			tcdio.ELINK_CITER.Store(int16(m))
			tcdio.ELINK_BITER.Store(int16(m))
		}
		d.wdata = ptr                // prevent premature GC and make Flush working
		atomic.StoreInt32(&d.wn, -1) // DMA write in progress
		dma.EnableReq()              // accept DMA requests from Tx FIFO
		p.MIER.Store(MasterErrFlags) // handle I2C errors
		if n == 0 {
			break // we don't have to wait for the end of write
		}
		d.Flush() // wait until the major loop complete or error
		if p.MSR.LoadBits(MasterErrFlags) != 0 {
			break
		}
	}
}

// WriteCmd works like WriteCmds but writes only one command word into the Tx
// FIFO.
func (d *Master) WriteCmd(cmd int16) {
	d.Flush()
	d.wbuf = cmd
	masterWrite(d, unsafe.Pointer(&d.wbuf), 1, true)
}

// WriteCmds starts writing commands into the Tx FIFO in the background using
// interrupts and/or DMA. WriteCmd is no-op if len(cmds) == 0.
//
// The LPI2C concept of the combined command and data FIFO greatly simplifies
// use of the I2C master. Thanks to this concept an I2C transaction or even
// multiple transactions can be prepared in advance as an array of commands and
// data, including receive transactions if the amount of data is known.
//
// There is, however, a certain weakness of the LPI2C peripheral when it comes
// to receiving data of an unknown quantity or if the data should be slowly
// received in chunks of size less than the Rx FIFO. Such transfers may require
// issuing repeat start conditions after each chunk to avoid MSR[FEF] error.
// This is because the LPI2C periperal transmits NACK at the end of the Recv
// command if the Rx FIFO isn't full and there is no next Recv or Discard
// command in the Tx FIFO. Two or more consecutive Recv commands in the list
// passed to WriteCmds may also cause the FIFO error because there is no
// guarantee that they will all get into the Tx FIFO on time (the first Recv
// command may be executed and finished by the implicit Stop condition and in
// such the second late Recv command causes MFEF because the Start command is
// required first).
func (d *Master) WriteCmds(cmds []int16) {
	if len(cmds) == 0 {
		return
	}
	d.Flush()
	masterWrite(d, unsafe.Pointer(unsafe.SliceData(cmds)), len(cmds), true)
}

// WriteBytes is like WriteCmds but writes only Send commands with the provided
// data.
func (d *Master) WriteBytes(p []byte) {
	if len(p) == 0 {
		return
	}
	if d.dma.IsValid() && len(p) >= 2*dma.MemAlign {
		ptr := unsafe.Pointer(&p[0])
		ds, de := dma.AlignOffsets(ptr, uintptr(len(p)))
		dmaStart := int(ds)
		dmaEnd := int(de)
		dmaPtr := unsafe.Add(ptr, ds)
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			masterWrite(d, ptr, dmaStart, false)
			d.Flush()
		}
		masterWriteDMA(d, dmaPtr, dmaN)
		if dmaEnd == len(p) {
			return
		}
		p = p[dmaEnd:]
	}
	d.Flush()
	masterWrite(d, unsafe.Pointer(unsafe.SliceData(p)), len(p), false)
}

// WriteStr is like WriteBytes but writes bytes from string instead of slice.
func (d *Master) WriteStr(s string) {
	if len(s) != 0 {
		d.WriteBytes(unsafe.Slice(unsafe.StringData(s), len(s)))
	}
}

func masterRead(d *Master, ptr *byte, n int) {
	p := d.p
	if p.MSR.LoadBits(MasterErrFlags) != 0 {
		return
	}
	// To speed things up we try to empty the FIFO in thread mode.
	i := 0
	data := unsafe.Slice((*byte)(ptr), n)
	for {
		v := p.MRDR.Load()
		if v&RXEMPTY != 0 {
			break
		}
		data[i] = byte(v)
		if i++; i == len(data) {
			return
		}
	}
	// The remaining data will be read by the ISR.
	d.rdata = ptr
	d.ri = int32(i)
	p.MFCR.Store(MFCR(min(n-i, rxFIFOCap)-1) << RXWATERn)
	atomic.StoreInt32(&d.rn, int32(n))
	flags := MRDF | MasterErrFlags
	if d.wn > 0 /* the above atomic.Store allows use of non-atomic load */ {
		flags |= MTDF
	}
	p.MIER.Store(flags) // race with ISR (clear)
	d.rdone.Sleep(-1)
	d.rdone.Clear()
	d.rdata = nil
}

func masterReadDMA(d *Master, ptr unsafe.Pointer, n int) {
	p := d.p
	if p.MSR.LoadBits(MasterErrFlags) != 0 {
		return
	}
	rtos.CacheMaint(rtos.DCacheFlushInval, ptr, n)
	const dmaChunk = 4 // equals 4 x S8b and 1 x D32b, <=rxFIFOLen
	tcd := dma.TCD{
		SADDR:       unsafe.Pointer(d.p.MRDR.Addr()),
		ATTR:        dma.S8b | dma.D32b,
		ML_NBYTES:   dmaChunk,
		DADDR:       ptr,
		DOFF:        dmaChunk,
		ELINK_CITER: dmaMaxMajorIter,
		ELINK_BITER: dmaMaxMajorIter,
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	p.MFCR.Store((dmaChunk - 1) << RXWATERn)
	if atomic.LoadInt32(&d.wn) == -1 {
		d.Flush() // wait for the end of DMA write
	}
	p.MDER.Store(RDDE) // clears TDDE
	dma := d.dma
	dma.WriteTCD(&tcd)
	tcdio := dma.TCD()
	n /= dmaChunk
	for {
		m := n
		if m > dmaMaxMajorIter {
			m = dmaMaxMajorIter
		}
		n -= m
		if m != dmaMaxMajorIter {
			tcdio.ELINK_CITER.Store(int16(m))
			tcdio.ELINK_BITER.Store(int16(m))
		}
		atomic.StoreInt32(&d.rn, -1) // DMA read in progress
		dma.EnableReq()              // accept DMA requests from Rx FIFO
		flags := MasterErrFlags
		if d.wn > 0 /* the above atomic.Store allows use of non-atomic load */ {
			flags |= MTDF
		}
		p.MIER.Store(flags) // handle I2C errors
		d.rdone.Sleep(-1)   // wait until the major loop complete or error
		d.rdone.Clear()
		if n == 0 || p.MSR.LoadBits(MasterErrFlags) != 0 {
			break
		}
	}
}

// ReadBytes reads len(p) data bytes from Rx FIFO. The read data is valid if Err
// returns nil.
func (d *Master) ReadBytes(p []byte) {
	if len(p) == 0 {
		return
	}
	if d.dma.IsValid() && len(p) >= 2*dma.MemAlign {
		ptr := &p[0]
		ds, de := dma.AlignOffsets(unsafe.Pointer(ptr), uintptr(len(p)))
		dmaStart := int(ds)
		dmaEnd := int(de)
		dmaPtr := &p[dmaStart]
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			masterRead(d, ptr, dmaStart)
		}
		masterReadDMA(d, unsafe.Pointer(dmaPtr), dmaN)
		if dmaEnd == len(p) {
			return
		}
		p = p[dmaEnd:]
	}
	masterRead(d, &p[0], len(p))
}

// ReadByte works like ReadBytes but reads only one byte from the Rx FIFO.
func (d *Master) ReadByte() (b byte) {
	masterRead(d, &b, 1)
	return
}

// Status returns the current status of the LPSPI Master. It's intended do to be
// used together with the Clear and Wait methods to check which of the event or
// state flags we were waiting for are actually set.
func (d *Master) Status() MSR {
	return d.p.MSR.Load()
}

const waitFlags = MEPF | MSDF | MDMF | MTDF

// Clear allows to clear the MEPF, MSDF, MDMF event flags. It is intended to be
// used together with the Wait method to wait for events signaled by these
// flags.
func (d *Master) Clear(flags MSR) {
	d.p.MSR.Store(flags & waitFlags)
}

// Wait waits for an the MEPF, MSDF, MDMF event flags and the MTDF state flag or
// an error. The MTDF flag allows to wait for an empty Tx FIFO. In most cases
// you should clear the event flags you want to wait for.
func (d *Master) Wait(flags MSR) {
	flags &= MEPF | MSDF | MDMF | MTDF
	if flags == 0 {
		return
	}
	flags |= MasterErrFlags
	p := d.p
	if p.MSR.LoadBits(flags) != 0 {
		return
	}
	atomic.StoreInt32(&d.rn, int32(flags|1<<31)) // wait: d.rn<0 && d.rn!=-1
	if flags&MTDF == 0 && d.wn > 0 {
		flags |= MTDF
	}
	p.MIER.Store(flags) // race with ISR (clear)
	d.rdone.Sleep(-1)
	d.rdone.Clear()
}

// Err returns the content of the MSR register wrapped into the MasterError type
// if any error flag (see MasterErrFlags) is set. Othewrise it returns nil.
// If clear is true Err clears the Tx FIFO and the error flags in the MSR
// register and if the LPI2C peripheral is in the busy state (MSR[MBF] is set)
// it also releases the bus by writing the Stop command into Tx FIFO.
func (d *Master) Err(clear bool) error {
	p := d.p
	status := p.MSR.Load()
	if e := status & MasterErrFlags; e != 0 {
		if clear {
			// Clear error flags (also clear Tx FIFO to ensure no new errors)
			p.MCR.SetBits(MRTF)
			p.MCR.ClearBits(MEN)
			for p.MSR.LoadBits(MBF) != 0 {
				runtime.Gosched()
			}
			p.MCR.SetBits(MEN)
			p.MSR.Store(e)
		}
		return &MasterError{d.name, status} // all flags for the better context
	}
	return nil
}

// ISR is the interrupt handler for the I2C peripheral used by Master.
//
//go:nosplit
//go:nowritebarrierrec
func (d *Master) ISR() {
	p := d.p

	// Disable interrupts and reenable them later if needed. It reaces with the
	// thread code. If the clearing INTR_MASK here happens before the setting
	// it in the thread code this ISR may run again. We clear d.wn, d.rn before
	// wake-up the thread code so such ISR reentry isn't harmful.
	p.MIER.Store(0) // disable all IRQs and fix it later
	sr := p.MSR.Load()

	if sr&MasterErrFlags != 0 {
		// Tx/Rx FIFOs are kept empty until TX_ABRT IRQ is cleared
		if wn := atomic.LoadInt32(&d.wn); wn != 0 {
			if wn == -1 {
				d.dma.DisableReq()
			}
			d.wn = 0
			d.wdone.Wakeup()
		}
		if rn := atomic.LoadInt32(&d.rn); rn != 0 {
			if rn == -1 {
				d.dma.DisableReq()
			}
			d.rn = 0
			d.rdone.Wakeup()
		}
		return
	}

	var enable MSR

	// Read or wait part.
	done := false
	if n := atomic.LoadInt32(&d.rn); n > 0 {
		// Read
		flags := MRDF | MasterErrFlags
		data := unsafe.Slice(d.rdata, n)
		i := d.ri
		for int(i) < len(data) {
			v := p.MRDR.Load()
			if v&RXEMPTY != 0 {
				break
			}
			data[i] = byte(v)
			i++
		}
		if i != d.ri {
			d.ri = i
			n -= i
			if n == 0 {
				flags = 0
				done = true
			} else {
				if n < rxFIFOCap {
					// Reduce MFCR[RXWATER] to the size of the last data chunk.
					p.MFCR.Store(MFCR(n-1) << RXWATERn)
				}
			}
		}
		enable |= flags
	} else if n < -1 {
		// Wait
		if flags := MSR(n) & waitFlags; flags&sr != 0 {
			done = true
		} else {
			enable |= flags // already contain MasterErrFlags
		}
	}
	if done {
		d.rn = 0
		d.rdone.Wakeup()
	}

	// Write part. May work concurently with the thread read code.
	if n := atomic.LoadInt32(&d.wn); n > 0 {
		flags := MTDF | MasterErrFlags
		if fw := txFIFOCap - p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn; fw != 0 {
			i := d.wi
			m := min(n, int32(fw)+i)
			if !d.cmd {
				for _, b := range unsafe.Slice((*byte)(d.wdata), m)[i:] {
					p.MTDR.Store(int16(b))
				}
			} else {
				for _, cmd := range unsafe.Slice((*int16)(d.wdata), m)[i:] {
					p.MTDR.Store(cmd)
				}
			}
			d.wi = m
			if m == n {
				// Done.
				flags = 0
				d.wn = 0
				d.wdone.Wakeup()
			}
		}
		enable |= flags
	}

	// The situation is clear if enable=0 because we cleared the whole MIER at
	// entry and next we checked d.wn and d.rn. Thread code does this in reverse
	// order so we are sure that there is no any new work for ISR with
	// interrupts disabled.
	if enable != 0 {
		// MIER must be set. There is no problem if the read part set the enable
		// because in this case we are sure that the thread read code waits for
		// this ISR and there is no thread write code doing anything. The
		// problem is if only the write part set the enable. In this case the
		// thread read code may have set MRDF in the meantime and we don't wont
		// disable it here.

		// First store enable as is.
		p.MIER.Store(enable)

		if enable&MRDF == 0 {
			// Then fix MIER if the read work was scheduled in the meantime.
			if n := atomic.LoadInt32(&d.rn); n != 0 {
				if n > 0 {
					enable |= MRDF
				} else {
					enable |= MSR(-n)
				}
				p.MIER.Store(enable)
			}
		}
	}
}

// DMAISR is a DMA interrupt handler for the DMA channel used by Master.
//
//go:nosplit
//go:nowritebarrierrec
func (d *Master) DMAISR() {
	d.dma.ClearInt()
	if atomic.LoadInt32(&d.wn) == -1 {
		d.wn = 0
		d.wdone.Wakeup()
	} else {
		d.rn = 0
		d.rdone.Wakeup()
	}
}
