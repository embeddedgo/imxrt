// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type ImageVectorTable struct {
	Tag      uint8
	LenHi    uint8
	LenLo    uint8
	Version  uint8
	Entry    uint32
	_        uint32
	DCD      uint32
	BootData uint32
	Self     uint32
	CSF      uint32
	_        uint32
}

const ivtSize = 32

type BootData struct {
	Start  uint32
	Length uint32
	Plugin uint32
}

const bootDataSize = 12

const (
	baseAddr      = 0x60000000
	ivtAddr       = baseAddr + 0x1000  // 0x60001000
	bootDataAddr  = baseAddr + 0x1020  // 0x60001020
	pluginAddr    = baseAddr + 0x1200  // 0x60001200
	stage2IVTAddr = baseAddr + 0x1300  // 0x60001300
	mbrEndAddr    = baseAddr + mbrSize // 0x60002000
)

var regularIVT = &ImageVectorTable{
	Tag:      0xd1,
	LenLo:    ivtSize,
	Version:  0x41,
	Entry:    mbrEndAddr, // for ARMv7-M the address of interrupt vector table
	BootData: bootDataAddr,
	Self:     ivtAddr,
}

var pluginIVT = &ImageVectorTable{
	Tag:      0xd1,
	LenLo:    ivtSize,
	Version:  0x41,
	Entry:    pluginAddr | 1,
	BootData: bootDataAddr,
	Self:     ivtAddr,
}

var stage2IVT = &ImageVectorTable{
	Tag:     0xd1,
	LenLo:   ivtSize,
	Version: 0x41,
	Entry:   mbrEndAddr,
	Self:    stage2IVTAddr,
}

var bootData = &BootData{
	Start: baseAddr, // Flex SPI start address
}
