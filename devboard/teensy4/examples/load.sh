#!/bin/sh

name=$(basename $(pwd))

teensy_loader_cli --mcu=TEENSY41 -v $name+mbr.hex || cat <<EOT
teensy_loader_cli may fail on the first attempt. Check all connections and try again.

See also:

https://forum.pjrc.com/threads/67232-Teensy-4-What-triggers-Halfkay-to-do-a-full-flash-erase
https://forum.pjrc.com/threads/69236-TeensyLoader-CLI-issues

EOT