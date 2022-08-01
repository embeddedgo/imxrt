#!/bin/sh

name=$(basename $(pwd))

teensy_loader_cli --mcu=TEENSY41 -v -w $name+mbr.hex