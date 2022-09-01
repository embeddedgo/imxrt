// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import (
	"embedded/mmio"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/internal"
	"github.com/embeddedgo/imxrt/hal/internal/ccm"
	"github.com/embeddedgo/imxrt/p/mmap"
)

type Controller struct {
	CR     mmio.R32[CR]
	es     mmio.U32
	_      uint32
	erq    mmio.U32
	_      uint32
	eei    mmio.U32
	ceei   mmio.U8
	seei   mmio.U8
	cerq   mmio.U8
	serq   mmio.U8
	cdne   mmio.U8
	ssrt   mmio.U8
	cerr   mmio.U8
	cint   mmio.U8
	_      uint32
	int    mmio.U32
	_      uint32
	err    mmio.U32
	_      uint32
	hrs    mmio.U32
	_      [3]uint32
	ears   mmio.U32
	_      [46]uint32
	dchpri [32]mmio.U8 // 3,2,1,0, 7,6,5,4, 11,10,9,8, ...
	_      [952]uint32
	tcd    [32]TCDIO
}

type TCDIO struct {
	SADDR       mmio.P32
	SOFF        mmio.R16[int16]
	ATTR        mmio.R16[ATTR]
	ML_NBYTES   mmio.R32[int32]
	SLAST       mmio.R32[int32]
	DADDR       mmio.P32
	DOFF        mmio.R16[int16]
	ELINK_CITER mmio.R16[uint16]
	DLAST_SGA   mmio.R32[int32]
	CSR         mmio.R16[CSR]
	ELINK_BITER mmio.R16[uint16]
}

func DMA(n int) *Controller {
	if n != 0 {
		panic("wrong DMA number")
	}
	return (*Controller)(unsafe.Pointer(mmap.DMA0_BASE))
}

// EnableClock enables clock for DMA controller.
// lp determines whether the clock remains on in low power WAIT mode.
func (d *Controller) EnableClock(lp bool) {
	ccm.CCGR(5).SetCG(3, ccm.ClkEn|int8(internal.BoolToInt(lp)<<1))
}

// DisableClock disables clock for DMA controller.
func (d *Controller) DisableClock() {
	ccm.CCGR(5).SetCG(3, 0)
}

type CR uint32

const (
	EDBG    CR = 0x01 << 1  //+ Enable Debug
	ERCA    CR = 0x01 << 2  //+ Enable Round Robin Channel Arbitration
	ERGA    CR = 0x01 << 3  //+ Enable Round Robin Group Arbitration
	HOE     CR = 0x01 << 4  //+ Halt On Error
	HALT    CR = 0x01 << 5  //+ Halt DMA Operations
	CLM     CR = 0x01 << 6  //+ Continuous Link Mode
	EMLM    CR = 0x01 << 7  //+ Enable Minor Loop Mapping
	GRP0PRI CR = 0x01 << 8  //+ Channel Group 0 Priority
	GRP1PRI CR = 0x01 << 10 //+ Channel Group 1 Priority
	ECX     CR = 0x01 << 16 //+ Error Cancel Transfer
	CX      CR = 0x01 << 17 //+ Cancel Transfer
	ACT     CR = 0x01 << 31 //+ DMA Active Status

	EDBGn    = 1
	ERCAn    = 2
	ERGAn    = 3
	HOEn     = 4
	HALTn    = 5
	CLMn     = 6
	EMLMn    = 7
	GRP0PRIn = 8
	GRP1PRIn = 10
	ECXn     = 16
	CXn      = 17
	ACTn     = 31
)

type Error uint32

const (
	DBE Error = 0x01 << 0  //+ Destination Bus Error
	SBE Error = 0x01 << 1  //+ Source Bus Error
	SGE Error = 0x01 << 2  //+ Scatter/Gather Configuration Error
	NCE Error = 0x01 << 3  //+ NBYTES/CITER Configuration Error
	DOE Error = 0x01 << 4  //+ Destination Offset Error
	DAE Error = 0x01 << 5  //+ Destination Address Error
	SOE Error = 0x01 << 6  //+ Source Offset Error
	SAE Error = 0x01 << 7  //+ Source Address Error
	CNE Error = 0x1F << 8  //+ Error Channel Number or Canceled Channel Number
	CPE Error = 0x01 << 14 //+ Channel Priority Error
	GPE Error = 0x01 << 15 //+ Group Priority Error
	CXE Error = 0x01 << 16 //+ Transfer Canceled
	VLD Error = 0x01 << 31 //+ VLD

	DBEn = 0
	SBEn = 1
	SGEn = 2
	NCEn = 3
	DOEn = 4
	DAEn = 5
	SOEn = 6
	SAEn = 7
	CNEn = 8
	CPEn = 14
	GPEn = 15
	CXEn = 16
	VLDn = 31
)

func (e Error) Error() string {
	if e == 0 {
		return ""
	}
	return "DMA error"
}

func (d *Controller) Err() Error {
	return Error(d.es.Load())
}

// Channel returns n-th channel of the controller. If you wont to obtain a
// free channel use AllocChannel.
func (d *Controller) Channel(n int) Channel {
	return Channel{uintptr(unsafe.Pointer(d)) | uintptr(n&31)}
}

var chanMask uint32 = 0xffff_ffff

// AllocChannel allocates a free channel in the controller. If pit is true the
// channel must have a periodic triggering capability. AllocChannel returns
// invalid channel if there is no free channel to be allocated.
// Use Channel.Free to free an unused channel.
func (d *Controller) AllocChannel(pit bool) Channel {
	for {
		chs := atomic.LoadUint32(&chanMask)
		n := 31
		if pit {
			// only first 4 channels have PIT capability
			chs &= 15
			n = 3
		}
		if chs == 0 {
			return Channel{}
		}
		mask := uint32(1) << uint(n)
		for chs&mask == 0 {
			mask >>= 1
			n--
		}
		if atomic.CompareAndSwapUint32(&chanMask, chs, chs&^mask) {
			return d.Channel(n)
		}
	}
}
