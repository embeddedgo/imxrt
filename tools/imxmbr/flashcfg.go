// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// IMXRT1060RM_rev3.pdf, Table 9-15. FlexSPI Configuration block
type FlexSPIConfigBlock struct {
	Tag                    uint32
	Version                uint32
	_                      uint32
	ReadSampleClkSrc       uint8
	CSHoldTime             uint8
	CSSetupTime            uint8
	ColumnAdressWidth      uint8
	DeviceModeCfgEnable    uint8
	_                      uint8
	WaitTimeCfgCommands    uint16
	DeviceModeSeq          uint32
	DeviceModeArg          uint32
	ConfigCmdEnable        uint8
	_                      [3]uint8
	ConfigCmdSeqs          [3]uint32
	_                      uint32
	CfgCmdArgs             [3]uint32
	_                      uint32
	ControllerMiscOption   uint32
	DeviceType             uint8
	SFlashPadType          uint8
	SerialClkFreq          uint8
	LUTCustomSeqEnable     uint8
	_                      [2]uint32
	SFlashA1Size           uint32
	SFlashA2Size           uint32
	SFlashB1Size           uint32
	SFlashB2Size           uint32
	CSPadSettingOverride   uint32
	SClkPadSettingOverride uint32
	DataPadSettingOverride uint32
	DQSPadSettingOverride  uint32
	TimeoutInMs            uint32
	CommandInterval        uint32
	DataValidTime          uint32
	BusyOffset             uint16
	BusyBitPolarity        uint16
	LookupTable            [64]uint32
	LUTCustomSeq           [12]uint32
	_                      [4]uint32
}

// IMXRT1060RM_rev3.pdf, Serial NOR configuration block
type SerialNORConfigBlock struct {
	MemCfg               FlexSPIConfigBlock
	PageSize             uint32
	SectorSize           uint32
	IPCmdSerialClkFreq   uint8
	IsUniformBlockSize   uint8
	_                    [2]uint8
	SerialNorType        uint8
	NeedExitNoCmdMode    uint8
	HalfClkForNonReadCmd uint8
	NeedRestoreNoCmdMode uint8
	BlockSize            uint32
	_                    [44]uint8
}

const KiB = 1 << 10

const flashConfigSize = 512

var flashConfig = &SerialNORConfigBlock{
	MemCfg: FlexSPIConfigBlock{
		Tag:              0x42464346,
		Version:          0x56010000,
		ReadSampleClkSrc: 1, // loopback from DQS pad
		CSHoldTime:       3, //
		CSSetupTime:      3, //
		DeviceType:       1, // Serial NOR
		SFlashPadType:    4, // Quad pads
		SerialClkFreq:    8, // 133 MHz
		LookupTable: [64]uint32{
			0:  0x0A1804EB,
			1:  0x26043206,
			4:  0x24040405,
			12: 0x00000406,
			20: 0x08180420,
			32: 0x081804D8,
			36: 0x08180402,
			37: 0x00002004,
			44: 0x00000460,
		},
	},
	PageSize:           256,  // bytes
	SectorSize:         4096, // bytes
	IPCmdSerialClkFreq: 1,    // 30 MHz
	BlockSize:          64 * KiB,
}
