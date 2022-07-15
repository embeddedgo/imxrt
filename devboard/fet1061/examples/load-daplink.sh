#!/bin/sh

MNT=/mnt/sdb

name=$(basename $(pwd))

mount $MNT
cp $name+mbr.bin $MNT
sync
if [ -f $MNT/FAIL.TXT ]; then
	cat $MNT/FAIL.TXT
	umount $MNT
	exit 1
fi
umount $MNT
exit 0
