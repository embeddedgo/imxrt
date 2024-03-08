// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lpspi

import (
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
)

type dataWord interface {
	~uint8 | ~uint16 | ~uint32
}

type Master struct {
	p     *Periph
	rxdma dma.Channel
	txdma dma.Channel
	done  rtos.Note
}

// NewMaster returns a new master-mode driver for p.
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
		string(dma.LPSPI1_TX) +
		string(dma.LPSPI2_TX) +
		string(dma.LPSPI3_TX) +
		string(dma.LPSPI4_TX)
	rxDMASlots = "" +
		string(dma.LPSPI1_RX) +
		string(dma.LPSPI2_RX) +
		string(dma.LPSPI3_RX) +
		string(dma.LPSPI4_RX)
)

// Enable enables LPSPI peripheral.
func (d *Master) Enable() {
	if rxdma := d.rxdma; rxdma.IsValid() {
		rxdma.DisableReq()
		rxdma.ClearInt()
	}
	if txdma := d.txdma; txdma.IsValid() {
		txdma.DisableReq()
		txdma.ClearInt()
	}
	d.p.CR.Store(DBGEN | MEN)
}

// Disable disables LPSPI peripheral.
func (d *Master) Disable() {
	d.p.CR.Store(0)
}

const (
	clkRoot  = 133e6 // 480e6 * 18 / 13 / 5,  TODO: calculate from CCM?
	fifoLen  = 16    // TODO: calculate from PARAM?
	dmaBurst = fifoLen * 3 / 4
)

// Setup enables the SPI clock, resets the peripheral and sets its basic
// configuration and the base SCK clock frequency. The base SPI clock frequency
// is set to baseFreq rounded down to 133 MHz divided by the number from 2 to
// 257. Use WriteCmd to fine tune the configuration and set the SCK prescaler
// to obtain the desired SPI clock frequency (datasheet says 30 MHz max, 33 MHz
// seems to work as well and there are reports that even 2x overclocking is
// achievable).
func (d *Master) Setup(conf CFGR1, baseFreqHz int) {
	p := d.p
	p.EnableClock(true)
	p.Reset()
	p.CFGR1.Store(conf | MASTER)
	switch {
	case baseFreqHz > clkRoot/2:
		baseFreqHz = clkRoot / 2
	case baseFreqHz <= 0:
		baseFreqHz = 1
	}
	//sckdiv := clkRoot/baseFreqHz - 2 // natural way but rounds sckdiv down
	sckdiv := clkRoot/(baseFreqHz+1) - 1
	if sckdiv > 255 {
		sckdiv = 255
	}
	p.CCR.Store(CCR(sckdiv))
	p.FCR.Store((dmaBurst-1)<<RXWATERn | (fifoLen-dmaBurst)<<TXWATERn)
	if txdma := d.txdma; txdma.IsValid() {
		txdma.DisableReq()
		txdma.DisableErrInt()
		txdma.SetMux(dma.Mux(txDMASlots[num(d.p)]) | dma.En)
		p.DER.SetBits(TDDE)
	}
	if rxdma := d.rxdma; rxdma.IsValid() {
		rxdma.DisableReq()
		rxdma.DisableErrInt()
		rxdma.SetMux(dma.Mux(rxDMASlots[num(d.p)]) | dma.En)
		p.DER.SetBits(RDDE)
	}
}

// RxDMAISR is required in DMA mode if Read* or WriteRead* methods are used.
func (d *Master) RxDMAISR() {
	d.rxdma.ClearInt()
	d.done.Wakeup()
}

// TxDMAISR is required in DMA mode if Write* (excluding WriteRead*) methods are
// used.
func (d *Master) TxDMAISR() {
	d.txdma.ClearInt()
	d.done.Wakeup()
}

// WriteCmd writes a command to the transmit FIFO. You can encode the frame size
// in cmd directly using the FRAMESZ field or specify it using the frameSize
// parameter (FRAMESZ = frameSize-1). The frame size is specified as a numer of
// bits. The minimum supported frame size is 8 bits and maximum is 4096 bits. If
// frameSize <= 32 it also specifies the word size. If frameSize > 32 then the
// word size is 32 except the last one which is equal to frameSize % 32 and
// must be >= 2 (e.g. frameSize = 33 is not supported).
func (d *Master) WriteCmd(cmd TCR, frameSize int) {
	d.p.TCR.Store(cmd | TCR(frameSize-1)&FRAMESZ)
}

func (d *Master) WriteWord(word uint32) {
	p := d.p
	for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
	}
	p.TDR.Store(word)
}

func (d *Master) ReadWord() uint32 {
	p := d.p
	for p.FSR.LoadBits(RXCOUNT) == 0 {
	}
	return p.RDR.Load()
}

func min[T int | uint | uintptr](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func max[T int | uint | uintptr](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// WriteRead writes and reads n bytes from/to out/in. It requires the empty
// recieve FIFO, including any potential new words caused by pending transaction
// because of the previous write. The transmit FIFO may contain any number of
// commands (that don't cause the new data to receive) but no data. WriteRead
// speed is crucial to achive fast bitrates (up to 30 MHz) so we use unsafe
// pointers instead of slices to speed things up (smaller code size, no bound
// checking, only one increment operation in the loop).
func writeRead[T dataWord](p *Periph, out, in unsafe.Pointer, n int) {
	sz := int(unsafe.Sizeof(T(0)))
	nr, nw := n, n
	nf := fifoLen // how many words can be written to TDR to don't overflow RDR
	for nw+nr != 0 {
		if nw != 0 {
			m := nf - int(p.FSR.LoadBits(TXCOUNT)>>TXCOUNTn)
			if m <= 0 {
				goto read
			}
			if m > nw {
				m = nw
			}
			nw -= m
			nf -= m
			for end := unsafe.Add(out, m*sz); out != end; out = unsafe.Add(out, sz) {
				p.TDR.Store(uint32(*(*T)(out)))
			}
		}
	read:
		if nr != 0 {
			m := int(p.FSR.LoadBits(RXCOUNT) >> RXCOUNTn)
			if m > nr {
				m = nr
			}
			nr -= m
			nf += m
			for end := unsafe.Add(in, m*sz); in != end; in = unsafe.Add(in, sz) {
				*(*T)(in) = T(p.RDR.Load())
			}
		}
	}
	return
}

func write[T dataWord](p *Periph, out unsafe.Pointer, n int) (end unsafe.Pointer) {
	sz := int(unsafe.Sizeof(T(0)))
	for end = unsafe.Add(out, n*sz); out != end; out = unsafe.Add(out, sz) {
		for p.FSR.LoadBits(TXCOUNT) == fifoLen<<TXCOUNTn {
		}
		p.TDR.Store(uint32(*(*T)(out)))
	}
	return
}

func read[T dataWord](p *Periph, in unsafe.Pointer, n int) (end unsafe.Pointer) {
	sz := int(unsafe.Sizeof(T(0)))
	for end = unsafe.Add(in, n*sz); in != end; in = unsafe.Add(in, sz) {
		for p.FSR.LoadBits(RXCOUNT) == 0 {
		}
		*(*T)(in) = T(p.RDR.Load())
	}
	return
}

// WriteReadSizes calculates the transfer sizes to be performed by CPU (ho, hi,
// to, ti) and DMA (d). It ensures that the middle dn*dmaBurst words are 32bit
// aligned and the cache maintenance operations performed for the d*dmaBurst
// words in the middle of the po, pi doesn't affect the memory outside these
// buffers. The following inequalities are true: ho >= hi, ti >= to.
func writeReadSizes(po, pi unsafe.Pointer, n int, lsz uint) (ho, hi, dn, to, ti int) {
	const cacheAlignMask = dma.CacheLineSize - 1
	ho = int(-uintptr(po)) & cacheAlignMask
	hi = int(-uintptr(pi)) & cacheAlignMask
	lenBytes := n << lsz
	burstBytes := dmaBurst << lsz
	to = int(uintptr(po)+uintptr(lenBytes)) & cacheAlignMask
	ti = int(uintptr(pi)+uintptr(lenBytes)) & cacheAlignMask
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

func writeReadDMA[T dataWord](d *Master, out, in []T) (n int) {
	n = min(len(out), len(in))
	if n == 0 {
		return
	}
	p := d.p
	po := unsafe.Pointer(unsafe.SliceData(out))
	pi := unsafe.Pointer(unsafe.SliceData(in))
	sz := int(unsafe.Sizeof(T(0)))
	lsz := uint(sz >> 1) // log2(sz) for 1, 2, 4

	// Use DMA only for long transfers. Short ones are handled by CPU.
	if n <= 3*dma.CacheLineSize/sz || !d.rxdma.IsValid() || !d.txdma.IsValid() {
		writeRead[T](p, po, pi, n)
		return
	}

	ho, hi, dn, to, ti := writeReadSizes(po, pi, n, lsz)

	// Use CPU to transfer the buffers heads to align them for DMA.
	writeRead[T](p, po, pi, hi)
	po = unsafe.Add(po, hi*sz)
	pi = unsafe.Add(pi, hi*sz)
	po = write[T](p, po, ho-hi)

	burstBytes := dmaBurst * sz
	rtos.CacheMaint(rtos.DCacheFlush, po, dn*burstBytes)
	rtos.CacheMaint(rtos.DCacheFlushInval, pi, dn*burstBytes)

	// Now perform the bidirectional DMA transfer using two DMA channels. The
	// whole transfer is synhronized by Rx channel only.

	rxdma, txdma := d.rxdma, d.txdma
	const maxMajorIter = 1<<dma.LINKCHn - 1 // only 511 because of ELINK Rx->Tx

	// Configure Tx DMA channel.
	tcd := dma.TCD{
		SADDR:       po,
		SOFF:        4,
		ATTR:        dma.S32b | dma.ATTR(lsz)<<dma.DSIZEn,
		ML_NBYTES:   uint32(burstBytes),
		DADDR:       unsafe.Pointer(p.TDR.Addr()),
		ELINK_CITER: maxMajorIter,
		ELINK_BITER: maxMajorIter,
	}
	txdma.WriteTCD(&tcd)

	// Configure Rx DMA channel. It uses ELINK to start the Tx channel minor
	// loop only after it finishes its own minor loop so the space in the Rx
	// FIFO is guaranteed.
	tcd.SADDR = unsafe.Pointer(p.RDR.Addr())
	tcd.SOFF = 0
	tcd.ATTR = dma.ATTR(lsz)<<dma.SSIZEn | dma.D32b
	tcd.DADDR = pi
	tcd.DOFF = 4
	tcd.ELINK_CITER = dma.ELINK | int16(txdma.Num()<<dma.LINKCHn) | maxMajorIter
	tcd.ELINK_BITER = tcd.ELINK_CITER
	tcd.CSR = dma.DREQ | dma.INTMAJOR
	rxdma.WriteTCD(&tcd)

	po = unsafe.Add(po, dn*burstBytes)
	pi = unsafe.Add(pi, dn*burstBytes)

	rxtcd, txtcd := rxdma.TCD(), txdma.TCD()
	for dn != 0 {
		m := dn
		if m > maxMajorIter {
			m = maxMajorIter
		}
		dn -= m

		if m != maxMajorIter {
			txtcd.ELINK_CITER.Store(int16(m))
			txtcd.ELINK_BITER.Store(int16(m))
		}
		txtcd.CSR.Store(dma.START) // start minor loop immediately

		if m != maxMajorIter {
			linkIter := dma.ELINK | int16(txdma.Num()<<dma.LINKCHn) | int16(m)
			rxtcd.ELINK_CITER.Store(linkIter)
			rxtcd.ELINK_BITER.Store(linkIter)
		}
		d.done.Clear()
		rxdma.EnableReq() // start minor loop if Rx FIFO contains enough data

		// Wait until the Rx DMA major loop complete.
		d.done.Sleep(-1)
	}

	// Use CPU to handle the unaligned tails.
	pi = read[T](p, pi, ti-to)
	writeRead[T](p, po, pi, to)

	return
}

// WriteRead writes n = min(len(out), len(in)) words to the transmit FIFO and
// at the same time it reads the same number of words from the receive FIFO.
// The written words are zero-extended bytes from the out slice. The least
// significant bytes from the read words are saved in the in slice.
func (d *Master) WriteRead(out, in []byte) (n int) {
	return writeReadDMA(d, out, in)
}

// WriteStringRead works like WriteRead.
func (m *Master) WriteStringRead(out string, in []byte) int {
	return writeReadDMA(d, unsafe.Slice(unsafe.StringData(out), len(out)), in)
}
