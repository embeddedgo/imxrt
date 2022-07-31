#!/bin/sh

name=$(basename $(pwd))

pyocd load $name+mbr.bin
