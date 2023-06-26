#!/bin/sh

FLASH_SIZE=4096 # KiB
FLEXRAM_BANK_CFG=0x5555_5555 # whole 512 KiB FlexRAM as OCRAM

imxmbr -flash $FLASH_SIZE -flexram $FLEXRAM_BANK_CFG mbr.img
