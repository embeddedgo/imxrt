#!/bin/sh

name=$(basename $(pwd))

emgo build $@ && cat ../mbr.img $name.bin >$name+mbr.bin && objcopy --change-addresses 0x60000000 -I binary -O ihex $name+mbr.bin $name+mbr.hex
