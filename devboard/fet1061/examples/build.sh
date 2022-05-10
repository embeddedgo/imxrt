#!/bin/sh

GOTARGET=imxrt1060

# FLEX_RAM: 512 KiB can assigned to:
# - FLEX_ITCM  at 0x00000000
# - FLEX_DTCM  at 0x20000000
# - FLEX_OCRAM at 0x20280000
# in 32 KiB chunks.
# OCRAM: 512 KiB at 0x20200000

GOTEXT=0x60002000
GOMEM=0x20000000:512K

ISRNAMES=no

name=$(basename $(pwd))

. ../../../../../scripts/build.sh $@ && objcopy -O binary $name.elf $name.tmp && cat ../ivt.bin $name.tmp >$name.bin && rm $name.tmp
