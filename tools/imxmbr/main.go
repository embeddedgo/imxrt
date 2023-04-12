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
		imgSize = flashSize - 8192
	}

	flashConfig.MemCfg.SFlashA1Size = uint32(flashSize)
	bootData.Length = uint32(imgSize)

	f, err := os.Create("mbr.bin")
	fatalErr(err)
	w := bufio.NewWriter(f)
	fatalErr(binary.Write(w, binary.LittleEndian, flashConfig))
	for i := 512; i < 4096; i++ {
		fatalErr(w.WriteByte(0xff))
	}
	fatalErr(binary.Write(w, binary.LittleEndian, ivt))
	fatalErr(binary.Write(w, binary.LittleEndian, bootData))
	for i := 4140; i < 8192; i++ {
		fatalErr(w.WriteByte(0xff))
	}
	fatalErr(w.Flush())
	fatalErr(f.Close())
}
