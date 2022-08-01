#!/bin/sh

# FLEX_RAM: 512 KiB can assigned to:
# - FLEX_ITCM  at 0x00000000
# - FLEX_DTCM  at 0x20000000
# - FLEX_OCRAM at 0x20280000 (continuation of OCRAM below)
# in 32 KiB chunks.
# OCRAM: 512 KiB at 0x20200000

export GOTARGET=imxrt1060
export GOTEXT=0x60002000
export GOMEM=0x20200000:512K
export GOOUT=bin

name=$(basename $(pwd))

emgo build $@ && cat ../mbr.img $name.bin >$name+mbr.bin && objcopy --change-addresses 0x60000000 -I binary -O ihex $name+mbr.bin $name+mbr.hex
