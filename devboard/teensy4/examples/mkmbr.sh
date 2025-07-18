#!/bin/sh

FLASH_SIZE=8192 # KiB, 2048 in case of Teensy 4.0
FLEXRAM_BANK_CFG=0x5555_5556 # 480 KiB OCRAM, 32 KiB DTCM

egtool imxmbr -flash $FLASH_SIZE -flexram $FLEXRAM_BANK_CFG mbr.img
