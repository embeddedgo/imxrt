// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import (
	"embedded/mmio"
	"unsafe"
)

type Channel struct {
	h uintptr
}

type TCD struct {
	SADDR       unsafe.Pointer
	SOFF        int16
	ATTR        ATTR
	ML_NBYTES   int32
	SLAST       int32
	DADDR       unsafe.Pointer
	DOFF        int16
	ELINK_CITER uint16
	DLASTSGA    int32
	CSR         CSR
	ELINK_BITER uint16
}

type ATTR uint16

const (
	DSIZE  ATTR = 0x07 << 0  //+ Destination data transfer size
	D8b    ATTR = 0x00 << 0  //  8-bit
	D16b   ATTR = 0x01 << 0  //  16-bit
	D32b   ATTR = 0x02 << 0  //  32-bit
	D64b   ATTR = 0x03 << 8  //  64-bit
	D4x64b ATTR = 0x05 << 8  //  32-byte burst (4 beats of 64 bits)
	DMOD   ATTR = 0x1F << 3  //+ Destination Address Modulo
	SSIZE  ATTR = 0x07 << 8  //+ Source data transfer size
	S8b    ATTR = 0x00 << 8  //  8-bit
	S16b   ATTR = 0x01 << 8  //  16-bit
	S32b   ATTR = 0x02 << 8  //  32-bit
	S64b   ATTR = 0x03 << 8  //  64-bit
	S4x64b ATTR = 0x05 << 8  //  32-byte burst (4 beats of 64 bits)
	SMOD   ATTR = 0x1F << 11 //+ Source Address Modulo

	DSIZEn = 0
	DMODn  = 3
	SSIZEn = 8
	SMODn  = 11
)

// ML_NBYTES ML fields
const (
	MLOFF int32 = 0x0fffff << 10 //+ Sign-extended offset applied to the source or destination address after the minor loop completes.
	DMLOE int32 = 0x01 << 30     //+ Destination Minor Loop Offset enable
	SMLOE int32 = -0x1 << 31     //+ Source Minor Loop Offset Enable

	MLOFFn = 10
	DMLOEn = 30
	SMLOEn = 31
)

// ELINK_CITER, ELINK_BITER ELINK fields
const (
	LINKCH uint16 = 0x1F << 9  //+ Minor Loop Link Channel Number
	ELINK  uint16 = 0x01 << 15 //+ Enable channel-to-channel linking on minor-loop complete

	LINKCHn = 9
	ELINKn  = 15
)

type CSR uint16

const (
	START       CSR = 0x01 << 0  //+ Channel Start
	INTMAJOR    CSR = 0x01 << 1  //+ Enable an interrupt when major iteration count completes
	INTHALF     CSR = 0x01 << 2  //+ Enable an interrupt when major counter is half complete
	DREQ        CSR = 0x01 << 3  //+ Disable Request
	ESG         CSR = 0x01 << 4  //+ Enable Scatter/Gather Processing
	MAJORELINK  CSR = 0x01 << 5  //+ Enable channel-to-channel linking on major loop complete
	ACTIVE      CSR = 0x01 << 6  //+ Channel Active
	DONE        CSR = 0x01 << 7  //+ Channel Done
	MAJORLINKCH CSR = 0x1F << 8  //+ Major Loop Link Channel Number
	BWC         CSR = 0x03 << 14 //+ Bandwidth Control
	Stall0c     CSR = 0x00 << 14 //  No eDMA engine stalls
	Stall4c     CSR = 0x02 << 14 //  eDMA engine stalls for 4 cycles after each R/W
	Stall8c     CSR = 0x03 << 14 //  eDMA engine stalls for 8 cycles after each R/W

	STARTn       = 0
	INTMAJORn    = 1
	INTHALFn     = 2
	DREQn        = 3
	ESGn         = 4
	MAJORELINKn  = 5
	ACTIVEn      = 6
	DONEn        = 7
	MAJORLINKCHn = 8
	BWCn         = 14
)

func d(c Channel) *Controller { return (*Controller)(unsafe.Pointer(c.h &^ 31)) }
func n(c Channel) uint        { return uint(c.h) & 31 }

func (c Channel) ReqEnabled() bool { return d(c).erq.Load()>>n(c)&1 != 0 }
func (c Channel) EnableReq()       { d(c).serq.Store(uint8(n(c))) }
func (c Channel) DisableReq()      { d(c).cerq.Store(uint8(n(c))) }
func (c Channel) IsReq() bool      { return d(c).hrs.Load()>>n(c)&1 != 0 }

func (c Channel) IsErr() bool         { return d(c).err.Load()>>n(c)&1 != 0 }
func (c Channel) ClearErr()           { d(c).cerr.Store(uint8(n(c))) }
func (c Channel) ErrIntEnabled() bool { return d(c).eei.Load()>>n(c)&1 != 0 }
func (c Channel) EnableErrInt()       { d(c).seei.Store(uint8(n(c))) }
func (c Channel) DisableErrInt()      { d(c).ceei.Store(uint8(n(c))) }

func (c Channel) IsInt() bool { return d(c).int.Load()>>n(c)&1 != 0 }
func (c Channel) ClearInt()   { d(c).cint.Store(uint8(n(c))) }

func (c Channel) ClearDone() { d(c).cdne.Store(uint8(n(c))) }
func (c Channel) Start()     { d(c).ssrt.Store(uint8(n(c))) }

type Prio uint8

const (
	CHPRI  Prio = 0x0F << 0 //+ Arbitration Priority
	GRPPRI Prio = 0x03 << 4 //+ Current Group Priority
	DPA    Prio = 0x01 << 6 //+ Disable Preempt Ability
	ECP    Prio = 0x01 << 7 //+ Enable Channel Preemption

	CHPRIn  = 0
	GRPPRIn = 4
	DPAn    = 6
	ECPn    = 7
)

func (c Channel) Prio() Prio {
	n := n(c)
	return Prio(d(c).dchpri[n&^3|(3-n&3)].Load())
}
func (c Channel) SetPrio(prio Prio) {
	n := n(c)
	d(c).dchpri[n&^3|(3-n&3)].Store(uint8(prio))
}

func (c Channel) ReadTCD(tcd *TCD) {
	tcda := (*[8]uint32)(unsafe.Pointer(tcd))
	tcdio := (*[8]mmio.U32)(unsafe.Pointer(&d(c).tcd[n(c)]))
	tcda[0] = tcdio[0].Load()
	tcda[1] = tcdio[1].Load()
	tcda[2] = tcdio[2].Load()
	tcda[3] = tcdio[3].Load()
	tcda[4] = tcdio[4].Load()
	tcda[5] = tcdio[5].Load()
	tcda[6] = tcdio[6].Load()
	tcda[7] = tcdio[7].Load()
}

func (c Channel) WriteTCD(tcd *TCD) {
	tcda := (*[8]uint32)(unsafe.Pointer(tcd))
	tcdio := (*[8]mmio.U32)(unsafe.Pointer(&d(c).tcd[n(c)]))
	tcdio[0].Store(tcda[0])
	tcdio[1].Store(tcda[1])
	tcdio[2].Store(tcda[2])
	tcdio[3].Store(tcda[3])
	tcdio[4].Store(tcda[4])
	tcdio[5].Store(tcda[5])
	tcdio[6].Store(tcda[6])
	tcdio[7].Store(tcda[7])
}

func (c Channel) LoadCSR() CSR {
	return CSR(d(c).tcd[n(c)].csr.Load())
}
