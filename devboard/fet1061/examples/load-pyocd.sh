#!/bin/sh

name=$(basename $(pwd))

egtool hex -inc ../mbr.img:0x60000000
pyocd load $name.hex
rm $name.hex
