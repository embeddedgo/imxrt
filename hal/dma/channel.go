// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import (
	"embedded/mmio"
	"sync/atomic"
	"unsafe"
)

// A Channel represents a DMA+DMAMUX channel together with the corresponding
// location in TCD memory.
type Channel struct {
	h uintptr
}

// A TCD represents a Transfer Control Descriptor
type TCD struct {
	SADDR       unsafe.Pointer
	SOFF        int16
	ATTR        ATTR
	ML_NBYTES   uint32
	SLAST       int32
	DADDR       unsafe.Pointer
	DOFF        int16
	ELINK_CITER int16
	DLAST_SGA   int32
	CSR         CSR
	ELINK_BITER int16
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
	DREQ        CSR = 0x01 << 3  //+ Disable Request at the end of major loop.
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

// Free frees the channel so the Controller.AllocChannel can allocate it next
// time.
func (c Channel) Free() {
	mask := uint32(1) << n(c)
	for {
		chs := atomic.LoadUint32(&chanMask)
		if atomic.CompareAndSwapUint32(&chanMask, chs, chs|mask) {
			break
		}
	}
}

func (c Channel) IsValid() bool    { return c.h != 0 }
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

// A Prio contains channel priority and some additional flags used in
// fixed-priority arbitration mode.
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

// Prio returns the current channel priority.
func (c Channel) Prio() Prio {
	n := n(c)
	return Prio(d(c).dchpri[n&^3|(3-n&3)].Load())
}

// SetPrio sets the channel priority.
func (c Channel) SetPrio(prio Prio) {
	n := n(c)
	d(c).dchpri[n&^3|(3-n&3)].Store(uint8(prio))
}

// ReadTCD reads a transfer controll descriptor from TCD memory.
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

// WriteTCD writes a transfer controll descriptor to the TCD memory. The eDMA
// memory controller probably informs the eDMA engine if the CSR[START] bit is
// set so the engine can start channel immediately.
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

// TCD returns the pointer to the corresponding TCD memory. You can use it to
// alter TCD fields in place.
func (c Channel) TCD() *TCDIO {
	return &d(c).tcd[n(c)]
}

// A Mux represents a configuration of DMAMUX for a DMA channel.
type Mux uint32

const (
	Src Mux = 0x7F << 0  //+ DMA Channel Source (Slot Number)
	AE  Mux = 0x01 << 29 //+ DMA Channel Always Enable
	PIT Mux = 0x01 << 30 //+ DMA Channel Trigger Enable
	En  Mux = 0x01 << 31 //+ DMA Mux Channel Enable

	// Sources (slots)
	FLEXIO1_REQ01     Mux = 0
	FLEXIO2_REQ01     Mux = 1
	LPUART1_TX        Mux = 2
	LPUART1_RX        Mux = 3
	LPUART3_TX        Mux = 4
	LPUART3_RX        Mux = 5
	LPUART5_TX        Mux = 6
	LPUART5_RX        Mux = 7
	LPUART7_TX        Mux = 8
	LPUART7_RX        Mux = 9
	FLEXCAN3          Mux = 11
	CSI               Mux = 12
	LPSPI1_RX         Mux = 13
	LPSPI1_TX         Mux = 14
	LPSPI3_RX         Mux = 15
	LPSPI3_TX         Mux = 16
	LPI2C1            Mux = 17
	LPI2C3            Mux = 18
	SAI1_RX           Mux = 19
	SAI1_TX           Mux = 20
	SAI2_RX           Mux = 21
	SAI2_TX           Mux = 22
	ADC_ETC           Mux = 23
	ADC1              Mux = 24
	ACMP              Mux = 25
	Reserved          Mux = 27
	FLEXSPI_RX        Mux = 28
	FLEXSPI_TX        Mux = 29
	XBAR1_REQ0        Mux = 30
	XBAR1_REQ1        Mux = 31
	FLEXPWM1_CAPT0    Mux = 32
	FLEXPWM1_CAPT1    Mux = 33
	FLEXPWM1_CAPT2    Mux = 34
	FLEXPWM1_CAPT3    Mux = 35
	FLEXPWM1_VAL0     Mux = 36
	FLEXPWM1_VAL1     Mux = 37
	FLEXPWM1_VAL2     Mux = 38
	FLEXPWM1_VAL3     Mux = 39
	FLEXPWM3_CAPT0    Mux = 40
	FLEXPWM3_CAPT1    Mux = 41
	FLEXPWM3_CAPT2    Mux = 42
	FLEXPWM3_CAPT3    Mux = 43
	FLEXPWM3_VAL0     Mux = 44
	FLEXPWM3_VAL1     Mux = 45
	FLEXPWM3_VAL2     Mux = 46
	FLEXPWM3_VAL3     Mux = 47
	QTIMER1_T0_CAPT   Mux = 48
	QTIMER1_T1_CAPT   Mux = 49
	QTIMER1_T2_CAPT   Mux = 50
	QTIMER1_T3_CAPT   Mux = 51
	QTIMER1_T0_CMPLD1 Mux = 52
	QTIMER1_T1_CMPLD2 Mux = 52
	QTIMER1_T1_CMPLD1 Mux = 53
	QTIMER1_T0_CMPLD2 Mux = 53
	QTIMER1_T2_CMPLD1 Mux = 54
	QTIMER1_T3_CMPLD2 Mux = 54
	QTIMER1_T3_CMPLD1 Mux = 55
	QTIMER1_T2_CMPLD2 Mux = 55
	QTIMER3_T0_CAPT   Mux = 56
	QTIMER3_T0_CMPLD1 Mux = 56
	QTIMER3_T1_CMPLD2 Mux = 56
	QTIMER3_T1_CAPT   Mux = 57
	QTIMER3_T1_CMPLD1 Mux = 57
	QTIMER3_T0_CMPLD2 Mux = 57
	QTIMER3_T2_CAPT   Mux = 58
	QTIMER3_T2_CMPLD1 Mux = 58
	QTIMER3_T3_CMPLD2 Mux = 58
	QTIMER3_T3_CAPT   Mux = 59
	QTIMER3_T3_CMPLD1 Mux = 59
	QTIMER3_T2_CMPLD2 Mux = 59
	FLEXSPI2_RX       Mux = 60
	FLEXSPI2_TX       Mux = 61
	FLEXIO1_REQ23     Mux = 64
	FLEXIO2_REQ23     Mux = 65
	LPUART2_TX        Mux = 66
	LPUART2_RX        Mux = 67
	LPUART4_TX        Mux = 68
	LPUART4_RX        Mux = 69
	LPUART6_TX        Mux = 70
	LPUART6_RX        Mux = 71
	LPUART8_TX        Mux = 72
	LPUART8_RX        Mux = 73
	PXP               Mux = 75
	LCDIF             Mux = 76
	LPSPI2_RX         Mux = 77
	LPSPI2_TX         Mux = 78
	LPSPI4_RX         Mux = 79
	LPSPI4_TX         Mux = 80
	LPI2C2            Mux = 81
	LPI2C4            Mux = 82
	SAI3_RX           Mux = 83
	SAI3_TX           Mux = 84
	SPDIF_RX          Mux = 85
	SPDIF_TX          Mux = 86
	ADC2              Mux = 88
	ACMP2             Mux = 89
	ACMP4             Mux = 90
	ENET_T0           Mux = 92
	ENET_T1           Mux = 93
	XBAR1_REQ2        Mux = 94
	XBAR1_REQ3        Mux = 95
	FLEXPWM2_CAPT0    Mux = 96
	FLEXPWM2_CAPT1    Mux = 97
	FLEXPWM2_CAPT2    Mux = 98
	FLEXPWM2_CAPT3    Mux = 99
	FLEXPWM2_VAL0     Mux = 100
	FLEXPWM2_VAL1     Mux = 101
	FLEXPWM2_VAL2     Mux = 102
	FLEXPWM2_VAL3     Mux = 103
	FLEXPWM4_CAPT0    Mux = 104
	FLEXPWM4_CAPT1    Mux = 105
	FLEXPWM4_CAPT2    Mux = 106
	FLEXPWM4_CAPT3    Mux = 107
	FLEXPWM4_VAL0     Mux = 108
	FLEXPWM4_VAL1     Mux = 109
	FLEXPWM4_VAL2     Mux = 110
	FLEXPWM4_VAL3     Mux = 111
	QTIMER2_T0_CAPT   Mux = 112
	QTIMER2_T1_CAPT   Mux = 113
	QTIMER2_T2_CAPT   Mux = 114
	QTIMER2_T3_CAPT   Mux = 115
	QTIMER2_T0_CMPLD1 Mux = 116
	QTIMER2_T1_CMPLD2 Mux = 116
	QTIMER2_T1_CMPLD1 Mux = 117
	QTIMER2_T0_CMPLD2 Mux = 117
	QTIMER2_T2_CMPLD1 Mux = 118
	QTIMER2_T3_CMPLD2 Mux = 118
	QTIMER2_T3_CMPLD1 Mux = 119
	QTIMER2_T2_CMPLD2 Mux = 119
	QTIMER4_T0_CAPT   Mux = 120
	QTIMER4_T0_CMPLD1 Mux = 120
	QTIMER4_T1_CMPLD2 Mux = 120
	QTIMER4_T1_CAPT   Mux = 121
	QTIMER4_T1_CMPLD1 Mux = 121
	QTIMER4_T0_CMPLD2 Mux = 121
	QTIMER4_T2_CAPT   Mux = 122
	QTIMER4_T2_CMPLD1 Mux = 122
	QTIMER4_T3_CMPLD2 Mux = 122
	QTIMER4_T3_CAPT   Mux = 123
	QTIMER4_T3_CMPLD1 Mux = 123
	QTIMER4_T2_CMPLD2 Mux = 123
	ENET2_T0          Mux = 124
	ENET2_T1          Mux = 125
)

func (c Channel) Mux() Mux {
	return Mux(d(c).chcfg[n(c)].Load())
}

func (c Channel) SetMux(mux Mux) {
	d(c).chcfg[n(c)].Store(uint32(mux))
}
