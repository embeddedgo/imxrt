#!/bin/sh

FLASH_SIZE=4096 # KiB
FLEXRAM_BANK_CFG=0x5555_5556 # 480 KiB OCRAM, 32 KiB DTCM

imxmbr -flash $FLASH_SIZE -flexram $FLEXRAM_BANK_CFG mbr.img
