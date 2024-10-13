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

// A Master is a driver for the LPI2C peripheral. It provides two kinds of
// interfaces to communicate with slave devices on the I2C bus.
//
// The first interface is a low-level one. It provides a set of methods to
// directly interract with the Data / Command FIFOs of the underlying LPI2C
// peripheral.
//
// Example:
//
//	d.WriteCmds(
//		lpi2c.Start|eepromAddr<<1|wr,
//		lpi2c.Send|int16(memAddr),
//		lpi2c.Start|eepromAddr<<1|rd,
//		lpi2c.Recv|int16(len(buf) - 1),
//		lpi2c.Stop,
//	)
//	d.Read(buf)
//	if err := d.Err(true); err != nil {
//
// Write methods in the low-level interface are asynchronous, that is, they may
// return before all commands/data will be written to the FIFO. Therefore you
// must not modify the data/command buffer pass to the last write method until
// the return of the Flush method or another write method.
//
// The read/write methods doesn't return errors. Instead, they check the status
// of the LPI2C peripheral before starting doing anything and return in case of
// error. There is an Err method that allow to check and reset the LPI2C error
// flags at a convenient time. Even if you call Err after every method call the
// returned error is still asynchronous due to the asynchronous nature of the
// write methods and the delayed execution of commands by the LPI2C peripheral
// itself. You can use Wait to synchronise things but all status flags other
// than MSDF (Stop Condition) seems to be inherently asynchronous too.
//
// The second interface is a connection oriented one that implements the
// i2cbus.Conn interface.
//
// Example:
//
//	conn := d.NewConn(eepromAddr)
//	conn.WriteByte(memAaddr)
//	conn.Read(buf)
//	err := conn.Close()
//	if err != nil {
//
// Both interfaces may be used concurently by multiple goroutines but in such
// a case users of the low-level interface must gain an exclusive access to the
// driver using the embedded mutex.
type Master struct {
	sync.Mutex // use with the low-level interface to share the driver

	name string
	p    *Periph
	id   uint8

	rbuf byte
	wbuf int16

	wcmds *int16
	wdata *byte
	wi    int32 // ISR cannot alter the above pointers so it alters wi instead
	wn    int32
	wdone rtos.Note

	rdata *byte
	ri    int32 // ISR cannot alter the above pointer so it alters ri instead
	rn    int32
	rdone rtos.Note

	dma   dma.Channel
	ddone rtos.Note
}

// NewMaster returns a new master-mode driver for p. If valid DMA channel is
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, dma dma.Channel) *Master {
	return &Master{
		name: string([]byte{'L', 'P', 'I', '2', 'C', '1' + byte(num(p))}),
		p:    p, dma: dma,
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
	string(dma.LPI2C1) +
	string(dma.LPI2C2) +
	string(dma.LPI2C3) +
	string(dma.LPI2C4)

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

// MasterError contains value of the Master Status Register with one or more
// error flags set.
type MasterError struct {
	Status MSR // value of the Master Status Register as read by Master.Err
}

func (e *MasterError) Error() string {
	var a [4]string
	es := a[:0:4]
	if e.Status&MNDF != 0 {
		es = append(es, "NACK")
	}
	if e.Status&MALF != 0 {
		es = append(es, "Arbitr")
	}
	if e.Status&MFEF != 0 {
		es = append(es, "FIFO")
	}
	if e.Status&MPLTF != 0 {
		es = append(es, "PinLow")
	}
	return "lpi2c master: " + strings.Join(es, ",")
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
			p.MCR.SetBits(MRTF) // clear Tx FIFOs
			p.MSR.Store(e)      // clear the error flags
			if p.MSR.LoadBits(MBF) != 0 {
				p.MTDR.Store(Stop) // release the bus
			}
		}
		return &MasterError{status} // all flags for the better context
	}
	return nil
}

// WriteCmds starts writing commands into the Tx FIFO in the background using
// interrupts and/or DMA. WriteCmd is no-op if len(cmds) == 0.
//
// The LPI2C concept of a combined command and data FIFO greatly simplifies use
// of the I2C protocol. Thanks to this concept an I2C transaction or even
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
// guarantee that they will all get into the Tx FIFO on time.
func (d *Master) WriteCmds(cmds []int16) {
	if len(cmds) == 0 {
		return
	}
	// Can't use DMA for commands because the DMA request/channel is shared
	// between Tx and Rx so we must wait for the end of Tx DMA before starting
	// Rx DMA. At the same time the Tx transfer may contain Recv command which
	// may not end before the subsequent read operation will complete.
	masterWrite(d, unsafe.Pointer(&cmds[0]), len(cmds), true)
}

// Write is like WriteCmds but writes only Send commands with the provided data.
func (d *Master) Write(p []byte) {
	if len(p) == 0 {
		return
	}
	if d.dma.IsValid() && len(p) >= 2*dma.CacheLineSize {
		ptr := unsafe.Pointer(&p[0])
		ds, de := dma.AlignOffsets(ptr, uintptr(len(p)))
		dmaStart := int(ds)
		dmaEnd := int(de)
		dmaPtr := unsafe.Add(ptr, ds)
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			masterWrite(d, ptr, dmaStart, false)
		}
		masterWriteDMA(d, dmaPtr, dmaN)
		if dmaEnd == len(p) {
			return
		}
		p = p[dmaEnd:]
	}
	masterWrite(d, unsafe.Pointer(&p[0]), len(p), false)
}

// WriteString is like Write but writes bytes from string instead of slice.
func (d *Master) WriteString(s string) {
	d.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

// WriteCmd works like WriteCmds but writes only one command word into the Tx
// FIFO. It's lighter than WriteCmds([]int16{cmd}) or Write([]byte{b}) mainly
// because there is no slice allocation required but its code is also much
// simpler.
func (d *Master) WriteCmd(cmd int16) {
	//masterWrite(d, unsafe.Pointer(&cmd), 1, true)
	//return

	if d.wn != 0 {
		masterWaitWrite(d)
	}
	p := d.p
	if p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn != txFIFOLen {
		p.MTDR.Store(cmd)
		return
	}
	d.wbuf = cmd
	d.wcmds = &d.wbuf
	d.wi = 0
	atomic.StoreInt32(&d.wn, 1)
	// The ISR may already finish here so the next line may reenable IRQs.
	p.MIER.Store(MTDF | MasterErrFlags)
}

const (
	txFIFOLen = 4
	rxFIFOLen = 4
)

const MasterErrFlags = MNDF | MALF | MFEF | MPLTF

// Wait until the ISR will end the previously scheduled transfer.
func masterWaitWrite(d *Master) {
	// Wait for the ISR to end the previously scheduled transfer.
	d.wdone.Sleep(-1)
	d.wdone.Clear()
	d.wcmds = nil
	d.wdata = nil
	d.wn = 0
}

func masterWrite(d *Master, ptr unsafe.Pointer, n int, cmd bool) {
	if d.wn != 0 {
		masterWaitWrite(d)
	}
	p := d.p
	// To speed things up, first try to write directly into the FIFO.
	i := 0
	if !cmd {
		data := unsafe.Slice((*byte)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOLen {
			p.MTDR.Store(int16(data[i]))
			if i++; i == len(data) {
				return
			}
		}
		d.wdata = &data[i]
	} else {
		cmds := unsafe.Slice((*int16)(ptr), n)
		for p.MFSR.LoadBits(TXCOUNT)>>TXCOUNTn < txFIFOLen {
			p.MTDR.Store(cmds[i])
			if i++; i == len(cmds) {
				return
			}
		}
		d.wcmds = &cmds[i]
	}
	// The remaining data/commands will be writtend to the FIFO by the ISR.
	d.wi = 0
	atomic.StoreInt32(&d.wn, int32(n-i))
	// The ISR may already finish here so the next line may reenable IRQs.
	p.MIER.Store(MTDF | MasterErrFlags)
}

const dmaMaxMajorIter = 1<<dma.ELINKn - 1 // = 32767

func masterWriteDMA(d *Master, ptr unsafe.Pointer, n int) {
	rtos.CacheMaint(rtos.DCacheFlush, ptr, n)
	if d.wn != 0 {
		masterWaitWrite(d)
	}
	tcd := dma.TCD{
		SADDR:       ptr,
		SOFF:        txFIFOLen,
		ATTR:        dma.S32b | dma.D8b,
		ML_NBYTES:   uint32(txFIFOLen),
		DADDR:       unsafe.Pointer(d.p.MTDR.Addr()),
		ELINK_CITER: dmaMaxMajorIter,
		ELINK_BITER: dmaMaxMajorIter,
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	d.p.MDER.Store(TDDE)
	dma := d.dma
	dma.WriteTCD(&tcd)
	tcdio := dma.TCD()
	n /= txFIFOLen
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
		dma.EnableReq()   // accept DMA requests from Tx FIFO
		d.ddone.Sleep(-1) // wait until the major loop complete
		d.ddone.Clear()
		if n == 0 {
			break
		}
	}
	d.p.MDER.Store(0)
}

// Read reads len(p) data bytes from Rx FIFO. The read data is valid if Err
// returns nil.
func (d *Master) Read(p []byte) {
	if len(p) == 0 {
		return
	}
	if d.dma.IsValid() && len(p) >= 3*dma.CacheLineSize {
		ptr := &p[0]
		ds, de := dma.AlignOffsets(unsafe.Pointer(ptr), uintptr(len(p)))
		dmaStart := int(ds)
		dmaEnd := int(de)
		dmaPtr := &p[dmaStart]
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			masterRead(d, ptr, dmaStart)
		}
		masterReadDMA(d, dmaPtr, dmaN)
		if dmaEnd == len(p) {
			return
		}
		p = p[dmaEnd:]
	}
	masterRead(d, &p[0], len(p))
}

// ReadByte works like Read but reads only one byte from the Rx FIFO. It's
// lighter than Read(oneByteSlice) because the way Read works causes that its
// argument definitely escapes so may require allocation, in worst case on every
// call. The ReadByte code is also much simpler than the Read code.
func (d *Master) ReadByte() byte {
	p := d.p
	v := p.MRDR.Load()
	if v&RXEMPTY != 0 {
		return byte(v)
	}
	d.rdata = &d.rbuf
	d.ri = 0
	p.MFCR.Store(0)
	flags := MRDF | MasterErrFlags
	atomic.StoreInt32(&d.rn, 1)
	if d.wn > 0 /* can avoid atomic.Load because of the above atomic.Store */ {
		flags |= MTDF
	}
	// The ISR may already finish here so the next line may reenable IRQs.
	p.MIER.Store(flags)
	d.rdone.Sleep(-1)
	d.rdone.Clear()
	return d.rbuf
}

func masterRead(d *Master, ptr *byte, n int) {
	p := d.p
	if p.MSR.Load()&MasterErrFlags != 0 {
		return
	}
	// Avoid interrupts if there is data in the FIFO.
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
	n -= i
	// The remaining data/commands will be read by the ISR.
	d.rdata = &data[i]
	d.ri = 0
	p.MFCR.Store(MFCR(min(n, rxFIFOLen)-1) << RXWATERn)
	flags := MRDF | MasterErrFlags
	atomic.StoreInt32(&d.rn, int32(n))
	if d.wn > 0 /* can avoid atomic.Load because of the above atomic.Store */ {
		flags |= MTDF
	}
	// The ISR may already finish here so the next line may reenable IRQs.
	p.MIER.Store(flags)
	d.rdone.Sleep(-1)
	d.rdone.Clear()
	d.rdata = nil
}

func masterReadDMA(d *Master, ptr *byte, n int) {
	// TODO:
}

// Flush waits for the last command passed to the last WriteCmd call or last
// data byte passed to the last Write/WriteString call to be written to the Tx
// FIFO. Return from Flush doesn't mean the written commands/data were or even
// will be executed/sent.
func (d *Master) Flush() {
	if d.wn != 0 {
		masterWaitWrite(d)
	}
}

// Status returns the current status of the LPSPI Master. It's intended do to
// be used with together with the Clear and Wait methods to check which of the
// events we were waiting for actually took place.
//
// You won't read this in the RM:
//
// In caes of repeated START the detection of NACK causes setting of both MSDF
// and MEPF flags. Usually MNDF flag is set before MSDF, MEPF but sometimes
// it happens in the reverse order. After NACK the SDA stays high, the SCL
// stays low which probably causes that the MBF and MBBF flags are set. The
// only way to clear MBF,MBBF is to write the Stop command into MTDR or reset
// the peripheral (disabling and reenabling it doesn't work). After the Stop
// command SCL is momentarily pulled low to allow releasing SDA and next SCL
// what corresponds to the Stop Condition on the bus. SM says that MBBF reflets
// the bus state. But it's not cleare how MBF and MBBF relate to each other.
func (d *Master) Status() MSR {
	return d.p.MSR.Load()
}

// Clear allows to clear the MEPF, MSDF, MDMF in the MSR register. It is
// intended to be used together with the Wait method to wait for events
// signaled by these flags.
func (d *Master) Clear(flags MSR) {
	d.p.MSR.Store(flags & (MEPF | MSDF | MDMF))
}

// Wait waits for an event described by the MEPF, MSDF, MDMF, MTDF flags or an
// error.  The MTDF flag allows to wait for an empty Tx FIFO. In most cases you
// should clear the flag you want to wait for.
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
	atomic.StoreInt32(&d.rn, -int32(flags))
	if flags&MTDF == 0 && d.wn > 0 /* no atomic.Load because of the above atomic.Store */ {
		flags |= MTDF
	}
	// The ISR may already finish here so the next line may reenable IRQs.
	p.MIER.Store(flags)
	d.rdone.Sleep(-1)
	d.rdone.Clear()
}

//go:nosplit
//go:nowritebarrierrec
func (d *Master) ISR() {
	// The tricky part of this code is the concurrent access of the MIER
	// register by this ISR and read/write/wait functions in thread mode. There
	// isn't clear that the MMIO supports RDEX/STREX instruction so we don't use
	// atomics.
	p := d.p
	p.MIER.Store(0) // disable all IRQs and fix it later
	sr := p.MSR.Load()

	if sr&MasterErrFlags != 0 {
		if atomic.LoadInt32(&d.wn) > 0 {
			d.wn = -1
			d.wdone.Wakeup()
		}
		if atomic.LoadInt32(&d.rn) != 0 {
			d.rn = 0
			d.rdone.Wakeup()
		}
		return
	}

	var ie MSR

	// Write part. May work concurently with masterRead.
	if n := atomic.LoadInt32(&d.wn); n > 0 {
		// Because MFCR[TXWATER]=0 (see Setup) the FIFO is now empty.
		i := d.wi
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
			// Done
			d.wn = -1 // avoid rentry because of possible race on MIER
			d.wdone.Wakeup()
		} else {
			ie = MTDF | MasterErrFlags
		}
	}

	// Read or wait part.
	done := false
	if n := atomic.LoadInt32(&d.rn); n > 0 {
		// Read
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
		d.ri = i
		if n := len(data) - int(i); n == 0 {
			done = true
		} else {
			ie |= MRDF | MasterErrFlags
			if n < rxFIFOLen {
				// Reduce MFCR[RXWATER] to the size of the last chunk of data.
				p.MFCR.Store(MFCR(n-1) << RXWATERn)
			}
		}
	} else if n < 0 {
		// Wait
		if flags := MSR(-n); flags&sr != 0 {
			done = true
		} else {
			ie |= flags // already contain MasterErrFlags
		}
	}
	if done {
		d.rn = 0 // avoid rentry because of possible race on MIER
		d.rdone.Wakeup()
	}

	// The situation is clear if ie=0 because we cleared whole MIER at entry and
	// next checked d.wn and d.rn. Thread code does this in reverse order so we
	// are sure that there is no any new work for ISR with interrupts disabled.
	if ie != 0 {
		// MIER must be set. There is no problem if read part set ie because in
		// this case we are sure that the thread read code waits for this ISR
		// and there is no thread write code doing anything. The problem is if
		// only the write part set ie. In this case the thread read code may
		// have set MRDF in the meantime and we don't wont disable it here.

		// First store ie as is.
		p.MIER.Store(ie)

		if ie&^(MTDF|MasterErrFlags) == 0 {
			// Then fix MIER if the read work was scheduled in the meantime.
			if n := atomic.LoadInt32(&d.rn); n != 0 {
				if n > 0 {
					ie |= MRDF
				} else {
					ie |= MSR(-n)
				}
				p.MIER.Store(ie)
			}
		}
	}
}

// DMAISR should be configured as a DMA interrupt handler if DMA is used.
//
//go:nosplit
func (d *Master) DMAISR() {
	d.dma.ClearInt()
	d.ddone.Wakeup()
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
