// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Dma shows how to use the DMA controller to perform the memory to memory
// transfer.
package main

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
)

func main() {
	// Number of words to copy.
	const n = 1e4 // must be <= 32767 because of Example 2.

	// We use dma.Alloc instead of the builtin make function because we need
	// cache-aligned buffers. In case of ordinary (non cache-aligned) buffers
	// the beginning and end of the buffers may require special treatment.
	src := dma.MakeSlice[uint32](n, n)
	dst := dma.MakeSlice[uint32](n, n)

	// Initialize the source memory with known pattern.
	for i := range src {
		src[i] = uint32(i<<16 | i)
	}

	// Unsafe pointers remember us that using DMA may be unsafe.
	srcAddr := unsafe.Pointer(&src[0])
	dstAddr := unsafe.Pointer(&dst[0])

	// Make sure all the values we wrote down in src are in place.
	rtos.CacheMaint(rtos.DCacheFlush, srcAddr, n*4)

	// Ungate the DMA clock.
	d := dma.DMA(0)
	d.EnableClock(true)

	// Allocate a free DMA channel. The hal/dma package, when importerted,
	// changes the default fixed arbitration mode to round-robin. AllocChannel
	// cannot be used with the fixed arbitration mode because in this mode,
	// all channels, even unused, must have unique priorities.
	c := d.AllocChannel(false)

	// Example 1. Transfer all data in the minor loop.
	//
	// Pros: Simple and fast.
	//
	// Cons: May increase the overall DMA latency/jitter because the minor loop
	// cannot be preempted in the round robin mode. Fixed arbitration isn't a
	// 100% solution because nested preemption isn't supported.

	// Flush and invalidate anything in the cache related to the dst buffer.
	rtos.CacheMaint(rtos.DCacheFlushInval, dstAddr, n*4)

	// Prepare a Transfer Control Descriptor. As the CRS[START] bit is set, the
	// transfer will start immediately after we write the prepared TCD to the
	// DMA local memory.
	tcd := dma.TCD{
		SADDR:       srcAddr,             // source address
		SOFF:        4,                   // added to SADDR after each read
		ATTR:        dma.S32b | dma.D32b, // src and dst data transfer sizes
		ML_NBYTES:   n * 4,               // number of bytes for minor loop
		SLAST:       -n * 4,              // added to SADDR when CITER reaches zero
		DADDR:       dstAddr,             // destination address
		DOFF:        4,                   // added to DADDR after each write
		ELINK_CITER: 1,                   // number of itreations in major loop
		DLAST_SGA:   -n * 4,              // added to DADDR when CITER reaches zero
		CSR:         dma.START,           // start the transfer immediately
		ELINK_BITER: 1,                   // reloaded to ELINK_CITER when CITER reaches zero
	}
	c.WriteTCD(&tcd)

	waitAndCheck(c, dst)

	// Example 2. Transfer data using major and minor loops.
	//
	// Pros: The DMA engine can handle multiple active channels at the same time
	// while it is possible to ensure ML_NBYTES uninterrupted transfer for each
	// of them.
	//
	// Cons: More complex and slower than transfering all data in minor loop.

	// Clear destination buffer.
	for i := range dst {
		dst[i] = 0
	}
	rtos.CacheMaint(rtos.DCacheFlushInval, dstAddr, n*4)

	// Modify TCD to use the major loop to the extreme.
	tcd.ML_NBYTES = 4   // extremely short (one-iteration) minor loop
	tcd.ELINK_CITER = n // number of iterations in major loop, must be <= 32767
	tcd.ELINK_BITER = n // must be the same as ELINK_CITER
	tcd.CSR = dma.DREQ  // stop at the end of major loop, required because of AE
	c.WriteTCD(&tcd)

	// CSR[START] starts the main loop for one iteration only. We use DMAMUX AE
	// (always enabled) request and CSR[DREQ] to run the minor loop to the end.
	c.SetMux(dma.En | dma.AE) // assert DMA request permanently
	c.EnableReq()             // accept requests

	waitAndCheck(c, dst)

	// Blink slow if all went well.
	for {
		leds.User.Toggle()
		time.Sleep(time.Second / 2)
	}
}

// WaitAndCheck waits for the end of transfer and checks the content of the dst
// buffer. It blinks fast endlessly if there is no proper pattern in dst.
func waitAndCheck(c dma.Channel, dst []uint32) {
	for c.TCD().CSR.LoadBits(dma.DONE) == 0 {
	}
	for i, w := range dst {
		for w != uint32(i<<16|i) {
			leds.User.Toggle()
			time.Sleep(time.Second / 8)
		}
	}
}
