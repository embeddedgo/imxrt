#!/bin/sh

# USB
#interface='-u'

# UART using /dev/ttyACM0
interface='-p /dev/ttyACM0'

sdphost $interface -- write-file 0x20000000 ../ivt_flashloader.img
sleep 1
sdphost $interface -- jump-address 0x20000400
sleep 1
blhost  $interface -- get-property 1
blhost  $interface -- fill-memory 0x2000 4 0xC0000006
blhost  $interface -- configure-memory 9 0x2000
blhost  $interface -- flash-image $1 erase 9
