// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/rtos"
	"runtime"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

// A Master is a driver to the LPSPI peripheral used in master mode.
type Master struct {
	p      *Periph
	rxdma  dma.Channel
	txdma  dma.Channel
	done   rtos.Note
	sckdiv uint16
	slow   bool
}

// NewMaster returns a new master-mode driver for p. If valid DMA channels are
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, rxdma, txdma dma.Channel) *Master {
	return &Master{p: p, rxdma: rxdma, txdma: txdma}
}

// Periph returns the underlying LPSPI peripheral.
func (d *Master) Periph() *Periph {
	return d.p
}

// All dma.Mux slot constants are less than 128 so we can use the string
// conversion to group them in a constant array.
const (
	txDMASlots = "" +
		string(rune(dma.LPSPI1_TX)) +
		string(rune(dma.LPSPI2_TX)) +
		string(rune(dma.LPSPI3_TX)) +
		string(rune(dma.LPSPI4_TX))
	rxDMASlots = "" +
		string(rune(dma.LPSPI1_RX)) +
		string(rune(dma.LPSPI2_RX)) +
		string(rune(dma.LPSPI3_RX)) +
		string(rune(dma.LPSPI4_RX))
)

// Enable enables LPSPI peripheral.
func (d *Master) Enable() {
	d.p.CR.Store(DBGEN | MEN)
}

// Disable disables LPSPI peripheral.
func (d *Master) Disable() {
	fsr := &d.p.FSR
	for fsr.LoadBits(TXCOUNT) != 0 {
	}
	d.p.CR.Store(0)
}

const (
	clkRoot   = 133e6 // 480e6 * 18 / 13 / 5,  TODO: calculate from CCM?
	fifoLen   = 16    // TODO: calculate from PARAM?
	dmaBurst  = fifoLen * 3 / 4
	slowPresc = clkRoot / 100e3
)

// Setup enables the SPI clock, resets the peripheral and sets the base SCK
// clock frequency to baseFreqHz rounded down to 133/n MHz, where n is an
// integer number from 2 to 257. The LPSPI controller is configured as master
// (CFGR1=MASTER). Other configuration registers have their default values. For
// custom configuration use the Periph method to access configuration registers.
// Different slave devices on the bust may require different SPI mode (CPOL,
// CPHA) and clock speed therefore, these types of settings are configured per
// transaction (see WriteCmd). The resulting SPI clock frequency should not
// exceed 30 MHz (33 MHz seems to work as well and there are reports that even
// 60 MHz is achievable).
func (d *Master) Setup(baseFreqHz int) {
	p := d.p
	p.EnableClock(true)
	p.Reset()
	p.CFGR1.Store(MASTER)
	switch {
	case baseFreqHz > clkRoot/2:
		baseFreqHz = clkRoot / 2
	case baseFreqHz <= 0:
		baseFreqHz = 1
	}
	sckdiv := (clkRoot+baseFreqHz-1)/baseFreqHz - 2
	if sckdiv > 255 {
		sckdiv = 255
	}
	d.sckdiv = uint16(sckdiv + 2)
	p.CCR.Store(CCR(sckdiv))
	p.FCR.Store((dmaBurst-1)<<RXWATERn | (fifoLen-dmaBurst)<<TXWATERn)
	if txdma := d.txdma; txdma.IsValid() {
		txdma.DisableReq()
		txdma.DisableErrInt()
		txdma.ClearInt()
		txdma.SetMux(dma.Mux(txDMASlots[num(d.p)]) | dma.En)
		p.DER.SetBits(TDDE)
	}
	if rxdma := d.rxdma; rxdma.IsValid() {
		rxdma.DisableReq()
		rxdma.DisableErrInt()
		rxdma.ClearInt()
		rxdma.SetMux(dma.Mux(rxDMASlots[num(d.p)]) | dma.En)
		p.DER.SetBits(RDDE)
	}
}

// BaseFreqHz returns the base frequency configured by the Setup method.
func (d *Master) BaseFreqHz() int {
	div := int(d.sckdiv)
	return (clkRoot + div - 1) / div
}

// RxDMAISR should be configured as an Rx DMA interrupt handler if DMA is used
// for read or write-read transactions (Write*, WriteRead* methods).
//
//go:nosplit
func (d *Master) RxDMAISR() {
	d.rxdma.ClearInt()
	d.done.Wakeup()
}

// TxDMAISR should be configured as a Tx DMA interrupt handler if DMA is used
// for write-only transactions (Write* methods, excluding WriteRead*).
//
//go:nosplit
func (d *Master) TxDMAISR() {
	d.txdma.ClearInt()
	d.done.Wakeup()
}

// WriteCmd writes a command to the transmit FIFO. These commands should be used
// to select a slave device by asserting PCSx pin, configure clock prescaler,
// signal polarity, phase and other things in different way for every
// transaction. You can encode the frame size in cmd directly using the FRAMESZ
// field or specify it using the frameSize parameter (FRAMESZ = frameSize-1).
// The frame size is specified as a numer of bits. The minimum supported frame
// size is 8 bits and maximum is 4096 bits. If frameSize <= 32 it also specifies
// the word size. If frameSize > 32 then the word size is 32 except the last one
// which is equal to frameSize % 32 and must be >= 2 (e.g. frameSize = 33 is
// not supported).
//
// # LPSPI BUGS
//
// The LPSPI peripheral has two bugs not mentioned in the errata that reveal in
// the master mode when the TCR.CONT bit (continuous transfer) is set:
//
// 1. In bidirectional mode, when you write n words to the TDR, you can read
// only n-1 words from RDR. The last word can be read after CONT is cleared or
// the peripheral is disabled in the CR register. It seems the received word is
// stored somewhere before it enters the receive FIFO so the writes to the TDR
// and reads from the RDR are out of sync for one word.
//
// 2. In Rx-only mode, the LPSPI starts reading data just after the command with
// the CONT and TXMSK bits is written to TCR. If you do not read from the RDR
// then 17 words are read from BUS (16 into FIFO and 1 elsewhere) and the REF
// error flag is set. If you next will read 1 word from the FIFO the LPSPI will
// read next 2 words from the BUS (one of them is lost).
func (d *Master) WriteCmd(cmd TCR, frameSize int) {
	p, slow := d.p, d.slow
	for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
		if slow {
			runtime.Gosched()
		}
	}
	p.TCR.Store(cmd | TCR(frameSize-1)&FRAMESZ)

	// Calculate speed to decide whether use runtime.Gosched when busy waiting.
	presc := cmd & PRESCALE >> PRESCALEn
	width := cmd & WIDTH >> WIDTHn
	d.slow = uint16(d.sckdiv)<<(presc+2-width) >= slowPresc<<2
}

// WriteWord writes a 32-bit data word to the transmit FIFO, waiting for a free
// FIFO slot if not available.
func (d *Master) WriteWord(word uint32) {
	p, slow := d.p, d.slow
	for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
		if slow {
			runtime.Gosched()
		}
	}
	p.TDR.Store(word)
}

// ReadWord reads a 32-bit data word from the receive FIFO, waiting for data if
// not available.
func (d *Master) ReadWord() uint32 {
	p, slow := d.p, d.slow
	for p.FSR.LoadBits(RXCOUNT) == 0 {
		if slow {
			runtime.Gosched()
		}
	}
	return p.RDR.Load()
}

type dataWord interface {
	~uint8 | ~uint16 | ~uint32
}

// WriteRead writes and reads n bytes from/to out/in. It requires the empty
// recieve FIFO, including any potential new words caused by pending transaction
// because of the previous write. The transmit FIFO may contain any number of
// commands (that don't cause the new data to receive) but no data. WriteRead
// speed is crucial to achive fast bitrates (up to 30 MHz) so we use unsafe
// pointers instead of slices to speed things up (smaller code size, no bound
// checking, only one increment operation in the loop).
func writeReadCPU[T dataWord](d *Master, po, pi unsafe.Pointer, n int) {
	p, slow := d.p, d.slow
	sz := int(unsafe.Sizeof(T(0)))
	nr, nw := n, n
	nf := fifoLen // how many words can be written to TDR to don't overflow RDR
	for nw+nr != 0 {
		var mr, mw int
		if nw != 0 {
			mw = nf - int(p.FSR.LoadBits(TXCOUNT)>>TXCOUNTn)
			if mw <= 0 {
				goto read
			}
			if mw > nw {
				mw = nw
			}
			nw -= mw
			nf -= mw
			for end := unsafe.Add(po, mw*sz); po != end; po = unsafe.Add(po, sz) {
				p.TDR.Store(uint32(*(*T)(po)))
			}
		}
	read:
		if nr != 0 {
			mr = int(p.FSR.LoadBits(RXCOUNT) >> RXCOUNTn)
			if mr > nr {
				mr = nr
			}
			nr -= mr
			nf += mr
			for end := unsafe.Add(pi, mr*sz); pi != end; pi = unsafe.Add(pi, sz) {
				*(*T)(pi) = T(p.RDR.Load())
			}
		}
		if mw+mr == 0 && slow {
			runtime.Gosched()
		}
	}
	return
}

func writeCPU[T dataWord](d *Master, po unsafe.Pointer, n int) (end unsafe.Pointer) {
	p, slow := d.p, d.slow
	sz := int(unsafe.Sizeof(T(0)))
	for end = unsafe.Add(po, n*sz); po != end; po = unsafe.Add(po, sz) {
		for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
			if slow {
				runtime.Gosched()
			}
		}
		p.TDR.Store(uint32(*(*T)(po)))
	}
	return
}

func readCPU[T dataWord](d *Master, pi unsafe.Pointer, n int) (end unsafe.Pointer) {
	p, slow := d.p, d.slow
	sz := int(unsafe.Sizeof(T(0)))
	for end = unsafe.Add(pi, n*sz); pi != end; pi = unsafe.Add(pi, sz) {
		for p.FSR.LoadBits(RXCOUNT) == 0 {
			if slow {
				runtime.Gosched()
			}
		}
		*(*T)(pi) = T(p.RDR.Load())
	}
	return
}

// BidirSizes calculates the transfer sizes to be performed by CPU (ho, hi, to,
// ti) and DMA (dn). It ensures that the middle dn*dmaBurst words are 32-bit
// aligned and the cache maintenance operations performed for the dn*dmaBurst
// words in the middle of the po, pi doesn't affect the memory outside these
// buffers. The following inequalities are true: ho >= hi, ti >= to.
func bidirSizes(po, pi unsafe.Pointer, n int, lsz uint) (ho, hi, dn, to, ti int) {
	const alignMask = dma.MemAlign - 1
	ho = int(-uintptr(po)) & alignMask
	hi = int(-uintptr(pi)) & alignMask
	lenBytes := n << lsz
	burstBytes := dmaBurst << lsz
	to = int(uintptr(po)+uintptr(lenBytes)) & alignMask
	ti = int(uintptr(pi)+uintptr(lenBytes)) & alignMask
	if a := hi - ho; a > 0 {
		// We cannot read more than what was written.
		ho += (a + 3) &^ 3
	} else if a = -a - (fifoLen<<lsz - burstBytes); a > 0 {
		// We can write more than will be immediately read but not that much.
		hi += (a + 3) &^ 3
	}
	dn = (lenBytes - max(ho+to, hi+ti)) / burstBytes
	to = lenBytes - ho - dn*burstBytes
	ti = lenBytes - hi - dn*burstBytes
	// Convert to words
	ho >>= lsz
	hi >>= lsz
	to >>= lsz
	ti >>= lsz
	return
}

func writeReadDMA(d *Master, po, pi unsafe.Pointer, n int, lsz uint) (npo, npi unsafe.Pointer) {
	burstBytes := dmaBurst << lsz
	rtos.CacheMaint(rtos.DCacheFlush, po, n*burstBytes)
	rtos.CacheMaint(rtos.DCacheFlushInval, pi, n*burstBytes)

	// Now perform the bidirectional DMA transfer using two DMA channels. The
	// whole transfer is synhronized by Rx channel only.

	const maxMajorIter = 1<<dma.LINKCHn - 1 // only 511 because of ELINK Rx->Tx
	rxdma, txdma := d.rxdma, d.txdma

	// Configure Tx DMA channel.
	tcd := dma.TCD{
		SADDR:       po,
		SOFF:        4,
		ATTR:        dma.S32b | dma.ATTR(lsz)<<dma.DSIZEn,
		ML_NBYTES:   uint32(burstBytes),
		DADDR:       unsafe.Pointer(d.p.TDR.Addr()),
		ELINK_CITER: maxMajorIter,
		ELINK_BITER: maxMajorIter,
	}
	txdma.WriteTCD(&tcd)

	// Configure Rx DMA channel. It uses ELINK to start the Tx channel minor
	// loop only after it finishes its own minor loop so the space in the Rx
	// FIFO is guaranteed.
	tcd.SADDR = unsafe.Pointer(d.p.RDR.Addr())
	tcd.SOFF = 0
	tcd.ATTR = dma.ATTR(lsz)<<dma.SSIZEn | dma.D32b
	tcd.DADDR = pi
	tcd.DOFF = 4
	tcd.ELINK_CITER = dma.ELINK | int16(txdma.Num()<<dma.LINKCHn) | maxMajorIter
	tcd.ELINK_BITER = tcd.ELINK_CITER
	tcd.CSR = dma.DREQ | dma.INTMAJOR
	rxdma.WriteTCD(&tcd)

	npo = unsafe.Add(po, n*burstBytes)
	npi = unsafe.Add(pi, n*burstBytes)

	rxtcd, txtcd := rxdma.TCD(), txdma.TCD()
	for n != 0 {
		m := n
		if m > maxMajorIter {
			m = maxMajorIter
		}
		n -= m

		if m != maxMajorIter {
			txtcd.ELINK_CITER.Store(int16(m))
			txtcd.ELINK_BITER.Store(int16(m))
		}
		txtcd.CSR.Store(dma.START) // run minor loop immediately

		if m != maxMajorIter {
			linkIter := dma.ELINK | int16(txdma.Num()<<dma.LINKCHn) | int16(m)
			rxtcd.ELINK_CITER.Store(linkIter)
			rxtcd.ELINK_BITER.Store(linkIter)
		}
		d.done.Clear()
		rxdma.EnableReq() // accept DMA requests from Rx FIFO
		d.done.Sleep(-1)  // wait until the Rx DMA major loop complete
	}

	return
}

func writeRead[T dataWord](d *Master, out, in []T) (n int) {
	n = min(len(out), len(in))
	if n == 0 {
		return
	}
	po := unsafe.Pointer(unsafe.SliceData(out))
	pi := unsafe.Pointer(unsafe.SliceData(in))
	sz := int(unsafe.Sizeof(T(0)))
	lsz := uint(sz >> 1) // log2(sz) for 1, 2, 4

	// Use DMA only for long transfers. Short ones are handled by CPU.
	if n <= 3*dma.MemAlign/sz || !d.rxdma.IsValid() || !d.txdma.IsValid() {
		writeReadCPU[T](d, po, pi, n)
		return
	}

	ho, hi, dn, to, ti := bidirSizes(po, pi, n, lsz)

	// Use CPU to transfer the buffers heads to align them for DMA.
	writeReadCPU[T](d, po, pi, hi)
	po = unsafe.Add(po, hi*sz)
	pi = unsafe.Add(pi, hi*sz)
	po = writeCPU[T](d, po, ho-hi)

	// Used DMA to write and read the aligned middle.
	po, pi = writeReadDMA(d, po, pi, dn, lsz)

	// Use CPU to handle the unaligned tails.
	pi = readCPU[T](d, pi, ti-to)
	writeReadCPU[T](d, po, pi, to)

	return
}

// WriteRead writes n = min(len(out), len(in)) bytes to the transmit FIFO,
// zero-extending any byte to the full 32-bit FIFO word. At the same time it
// reads the same number of bytes from the receive FIFO, using only the low
// significant bytes from the available 32-bit FIFO words.
func (d *Master) WriteRead(out, in []byte) (n int) {
	return writeRead(d, out, in)
}

// WriteStringRead works like WriteRead but the output bytes are taken from the
// string.
func (d *Master) WriteStringRead(out string, in []byte) int {
	return writeRead(d, unsafe.Slice(unsafe.StringData(out), len(out)), in)
}

// WriteRead16 works like WriteRead but with 16-bit words instead of bytes.
func (d *Master) WriteRead16(out, in []uint16) (n int) {
	return writeRead(d, out, in)
}

// WriteRead32 works like WriteRead but with 32-bit words instead of bytes.
func (d *Master) WriteRead32(out, in []uint32) (n int) {
	return writeRead(d, out, in)
}

func unidirSizes(ptr unsafe.Pointer, n int, lsz uint) (hn, dn, tn int) {
	const alignMask = dma.MemAlign - 1
	hn = int(-uintptr(ptr)) & alignMask
	lenBytes := n << lsz
	burstBytes := dmaBurst << lsz
	tn = int(uintptr(ptr)+uintptr(lenBytes)) & alignMask
	dn = (lenBytes - hn - tn) / burstBytes
	tn = lenBytes - hn - dn*burstBytes
	// Convert to words
	hn >>= lsz
	tn >>= lsz
	return
}

func writeDMA(d *Master, p unsafe.Pointer, n int, lsz uint) (np unsafe.Pointer) {
	burstBytes := dmaBurst << lsz
	rtos.CacheMaint(rtos.DCacheFlush, p, n*burstBytes)

	// Now write the aligned middle of the buffer using DMA.

	const maxMajorIter = 1<<dma.ELINKn - 1 // = 32767
	txdma := d.txdma

	// Configure Tx DMA channel.
	tcd := dma.TCD{
		SADDR:       p,
		SOFF:        4,
		ATTR:        dma.S32b | dma.ATTR(lsz)<<dma.DSIZEn,
		ML_NBYTES:   uint32(burstBytes),
		DADDR:       unsafe.Pointer(d.p.TDR.Addr()),
		ELINK_CITER: maxMajorIter,
		ELINK_BITER: maxMajorIter,
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	txdma.WriteTCD(&tcd)

	np = unsafe.Add(p, n*burstBytes)

	tcdio := txdma.TCD()
	for n != 0 {
		m := n
		if m > maxMajorIter {
			m = maxMajorIter
		}
		n -= m
		if m != maxMajorIter {
			tcdio.ELINK_CITER.Store(int16(m))
			tcdio.ELINK_BITER.Store(int16(m))
		}
		d.done.Clear()
		txdma.EnableReq() // accept DMA requests from Tx FIFO
		d.done.Sleep(-1)  // wait until the major loop complete
	}

	return
}

func write[T dataWord](d *Master, out []T) {
	if len(out) == 0 {
		return
	}
	po := unsafe.Pointer(unsafe.SliceData(out))
	sz := int(unsafe.Sizeof(T(0)))
	lsz := uint(sz >> 1) // log2(sz) for 1, 2, 4

	// Use DMA only for long transfers. Short ones are handled by CPU.
	if len(out) <= 3*dma.MemAlign/sz || !d.txdma.IsValid() {
		writeCPU[T](d, po, len(out))
		return
	}

	hn, dn, tn := unidirSizes(po, len(out), lsz)

	// Use CPU to write the buffer head to align it for DMA.
	po = writeCPU[T](d, po, hn)

	// Used DMA to write the aligned middle.
	po = writeDMA(d, po, dn, lsz)

	// Use CPU to handle the unaligned tail.
	writeCPU[T](d, po, tn)
}

// Write implements the io.Writer interface. It works like Write32 but for 8-bit
// words.
func (d *Master) Write(p []byte) (int, error) {
	write(d, p)
	return len(p), nil
}

// WriteString implemets the io.StringWriter interface. See Write for more
// information.
func (d *Master) WriteString(s string) (int, error) {
	write(d, unsafe.Slice(unsafe.StringData(s), len(s)))
	return len(s), nil
}

// Write16 works like Write32 but for 16-bit words.
func (d *Master) Write16(p []uint16) {
	write(d, p)
}

// Write32 is designed for unidirectional mode of operation, e.g. the TCR.RXMSK
// bit was set by the last command. It may also be used for bidirectional
// transfers, provided len(p) is less than the free space in the receive FIFO
// (not recommended, use WriteRead32 instead).
// always returns len(p), nil.
func (d *Master) Write32(p []uint32) {
	write(d, p)
}

func readDMA(d *Master, p unsafe.Pointer, n int, lsz uint) (np unsafe.Pointer) {
	burstBytes := dmaBurst << lsz
	rtos.CacheMaint(rtos.DCacheFlushInval, p, n*burstBytes)

	// Now read into the aligned middle of the buffer using DMA.

	const maxMajorIter = 1<<dma.ELINKn - 1 // = 32767
	rxdma := d.rxdma

	// Configure Rx DMA channel.
	tcd := dma.TCD{
		SADDR:       unsafe.Pointer(d.p.RDR.Addr()),
		ATTR:        dma.ATTR(lsz)<<dma.SSIZEn | dma.D32b,
		ML_NBYTES:   uint32(burstBytes),
		DADDR:       p,
		ELINK_CITER: maxMajorIter,
		ELINK_BITER: maxMajorIter,
		CSR:         dma.DREQ | dma.INTMAJOR,
	}
	p = unsafe.Add(p, n*burstBytes)
	rxdma.WriteTCD(&tcd)

	tcdio := rxdma.TCD()
	for n != 0 {
		m := n
		if m > maxMajorIter {
			m = maxMajorIter
		}
		n -= m
		if m != maxMajorIter {
			tcdio.ELINK_CITER.Store(int16(m))
			tcdio.ELINK_BITER.Store(int16(m))
		}
		d.done.Clear()
		rxdma.EnableReq() // accept DMA requests from Rx FIFO
		d.done.Sleep(-1)  // wait until the major loop complete
	}

	return
}

func read[T dataWord](d *Master, in []T) {
	if len(in) == 0 {
		return
	}
	pi := unsafe.Pointer(unsafe.SliceData(in))
	sz := int(unsafe.Sizeof(T(0)))
	lsz := uint(sz >> 1) // log2(sz) for 1, 2, 4

	// Use DMA only for long transfers. Short ones are handled by CPU.
	if len(in) <= 3*dma.MemAlign/sz || !d.rxdma.IsValid() {
		readCPU[T](d, pi, len(in))
		return
	}

	hn, dn, tn := unidirSizes(pi, len(in), lsz)

	// Use CPU to read the buffer head to align it for DMA.
	pi = readCPU[T](d, pi, hn)

	// Used DMA to read the aligned middle.
	pi = readDMA(d, pi, dn, lsz)

	// Use CPU to handle the unaligned tail.
	readCPU[T](d, pi, tn)
}

// Read implements the io.Reader interface. It works like Read32 but for 8-bit
// words instead of bytes.
//
// BUG: Typical usage scenarios of this function require 8-bit frame size with
// the TXMSK and CONT bits set but there is a hardware bug that makes this
// configuration unusable (see WriteCmd for more information).
func (d *Master) Read(p []byte) (int, error) {
	read(d, p)
	return len(p), nil
}

// Read16 works like Read32 but for 16-bit words instead of bytes.
//
// BUG: Typical usage scenarios of this function require 16-bit frame size with
// the TXMSK and CONT bits set but there is a hardware bug that makes this
// configuration unusable (see WriteCmd for more information).
func (d *Master) Read16(p []uint16) {
	read(d, p)
}

// Read32 is designed for the unidirectional mode of operation, e.g. the TXMSK
// bit and the proper frame size were set by the last command. It  may also be
// used for bidirectional transfers provided there are at least len(p) words
// available in the recevie FIFO (not recommended, use WriteRead32 instead).
//
// BUG: There are known hardware bugs related to Rx-only mode (see WriteCmd for
// more information). In contrast to Read and Read16 this function can be used
// with TXMSK set provided the frame size is set to 32*len(p) and the CONT bit
// is cleared.
func (d *Master) Read32(p []uint32) {
	read(d, p)
}
