#!/bin/sh

name=$(basename $(pwd))

emgo build $@ && cat ../mbr.img $name.bin >$name+mbr.bin
