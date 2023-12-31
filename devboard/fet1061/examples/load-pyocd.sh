#!/bin/sh

name=$(basename $(pwd))

pyocd load $name.hex
