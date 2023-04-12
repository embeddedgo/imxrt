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
		"\nUsage:\n  imxrtmbr [options] MBR_FILE\n\nOptoins:\n",
	)
	flag.PrintDefaults()
}

const mbrSize = 8192

func main() {
	var flashSize, imgSize uint

	flag.UintVar(&flashSize, "flash", 0, "flash size in megabytes (MiB)")
	flag.UintVar(&imgSize, "image", 0, "program image size in bytes (0 means all the remaining flash space)")
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
	flashSize *= MiB
	if imgSize == 0 {
		imgSize = flashSize - mbrSize
	}

	flashConfig.MemCfg.SFlashA1Size = uint32(flashSize)
	bootData.Length = uint32(imgSize)
	bootData.Plugin = 1

	f, err := os.Create(flag.Arg(0))
	fatalErr(err)
	w := bufio.NewWriter(f)
	fatalErr(binary.Write(w, binary.LittleEndian, flashConfig))
	for a := baseAddr + flashConfigSize; a < ivtAddr; a++ {
		fatalErr(w.WriteByte(0xff))
	}
	fatalErr(binary.Write(w, binary.LittleEndian, pluginIVT))
	fatalErr(binary.Write(w, binary.LittleEndian, bootData))
	for a := bootDataAddr + bootDataSize; a < dcdAddr; a++ {
		fatalErr(w.WriteByte(0xff))
	}
	fatalErr(binary.Write(w, binary.BigEndian, dcd))
	for a := dcdAddr + len(dcd)*4; a < pluginAddr; a++ {
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
	fatalErr(w.Flush())
	fatalErr(f.Close())
}
