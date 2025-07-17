#!/bin/sh

GOENV=../go.env go build
egtool hex -inc ../mbr.img:0x60000000