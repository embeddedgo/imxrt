#!/bin/sh

GOTARGET=imxrt1060

# FLEX_RAM: 512 KiB can assigned to:
# - FLEX_ITCM  at 0x00000000
# - FLEX_DTCM  at 0x20000000
# - FLEX_OCRAM at 0x20280000 (continuation of OCRAM below)
# in 32 KiB chunks.
# OCRAM: 512 KiB at 0x20200000

GOTEXT=0x60002000
GOMEM=0x20200000:512K

ISRNAMES=no

name=$(basename $(pwd))

. $(emgo env GOROOT)/../scripts/build.sh $@ && objcopy -O binary $name.elf $name.bin && cat ../mbr.img $name.bin >$name+mbr.bin
