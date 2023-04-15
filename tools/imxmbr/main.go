// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

func fatalErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprint(
		os.Stderr,
		"\nUsage:\n  imxrtmbr [options] MBR_FILE\n\nOptions:\n",
	)
	flag.PrintDefaults()
}

const mbrSize = 8192

func main() {
	var flashSize, imageSize, flexRAMCfg uint

	flag.UintVar(
		&flashSize, "flash", 0,
		"flash size (KiB)",
	)
	flag.UintVar(
		&imageSize, "image", 0,
		"program image size (KiB), 0 means all the remaining flash space",
	)
	flag.UintVar(
		&flexRAMCfg, "flexram", 0,
		"FlexRAM configuration (the value to write to the GPR17)",
	)
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		return
	}

	if flashSize == 0 {
		fmt.Fprintln(os.Stderr, "flash size not set")
		usage()
		os.Exit(1)
	}
	flashSize *= KiB
	imageSize *= KiB
	if imageSize == 0 {
		imageSize = flashSize - mbrSize
	}

	flashConfig.MemCfg.SFlashA1Size = uint32(flashSize)

	f, err := os.Create(flag.Arg(0))
	fatalErr(err)
	w := bufio.NewWriter(f)
	fatalErr(binary.Write(w, binary.LittleEndian, flashConfig))
	for a := baseAddr + flashConfigSize; a < ivtAddr; a++ {
		fatalErr(w.WriteByte(0xff))
	}
	if flexRAMCfg == 0 {
		bootData.Length = uint32(imageSize)
		fatalErr(binary.Write(w, binary.LittleEndian, regularIVT))
		fatalErr(binary.Write(w, binary.LittleEndian, bootData))
		for a := bootDataAddr + bootDataSize; a < mbrEndAddr; a++ {
			fatalErr(w.WriteByte(0xff))
		}
	} else {
		bootData.Length = uint32(pluginAddr + len(plugin)*2 - baseAddr)
		bootData.Plugin = 1
		imageSize -= stage2IVTAddr - baseAddr
		pluginImageSize[0] = uint16(imageSize)
		pluginImageSize[1] = uint16(imageSize >> 16)
		pluginFlexRAMCfg[0] = uint16(flexRAMCfg)
		pluginFlexRAMCfg[1] = uint16(flexRAMCfg >> 16)
		fatalErr(binary.Write(w, binary.LittleEndian, pluginIVT))
		fatalErr(binary.Write(w, binary.LittleEndian, bootData))
		for a := bootDataAddr + bootDataSize; a < pluginAddr; a++ {
			fatalErr(w.WriteByte(0xff))
		}
		fatalErr(binary.Write(w, binary.LittleEndian, plugin))
		for a := pluginAddr + len(plugin)*2; a < stage2IVTAddr; a++ {
			fatalErr(w.WriteByte(0xff))
		}
		fatalErr(binary.Write(w, binary.LittleEndian, stage2IVT))
		for a := stage2IVTAddr + ivtSize; a < mbrEndAddr; a++ {
			fatalErr(w.WriteByte(0xff))
		}
	}
	fatalErr(w.Flush())
	fatalErr(f.Close())
}
