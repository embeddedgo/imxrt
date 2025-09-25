### FET1061 overview

The FET1061 board has much potential and is ideal to be used in an end product but isn't begginer friendly at all. No ready to use pins, small pads difficult to solder, no 3.3V source on board, no boot and reset buttons/jumpers.

### Prerequisites

You need an SWD programmer connected to the SWD pins of the board or use UART (slow).

The fastest flashing speed you can probably achiewe with the High-Speed DAPLink Debug Probe (based on the ATSAM3U4C, $20 on Aliexpress) together with pyOCD (see load-pyocd.sh). This is setup I use but any other SWD programmer and OpenOCD can be used as well. I use pyOCD for flashing and OpenOCD for debugging (OpenOCD supports GDB server on stdin/stdout which is more convenient to script than the TCP server).

### Boot modes

Pads 48 (BOOT_MODE0) and 49 (BOOT_MODE1) set the boot configuration for the chip.

| BOOT_MODE1 | BOOT_MODE0 | Boot Type         |
|------------|------------|-------------------|
| 0          | 0          | boot from fuses   |
| 0          | 1          | serial downloader |
| 1          | 0          | internal boot     |
| 1          | 1          | reserveed         |

When pads 48, 49 are left floating the boot type is *internal boot* (the board pulls down the pad 48 to GDN, pad 49 is pulled up to 3V).

The *internal boot* mode runs the program in flash. It can be theoretically used when flashing but in practice it doesn't work probably because of the WFE instruction used by the running Go programs to sleep the CPU when waiting for an IRQ/event (it should work when busy waiting is used instead of WFE/WFI).

For realiable flashing you should put the MCU to the *boot from fuses* mode or the *serial downloader* mode. For SWD flashing simply short the pad 49 to GND while rebooting the board. You can reboot it by momemtary shorting the pad 20 (POR_B) to GND or simply disconnect and connect power. For UART flashing you should also pull pad 48 to the externally generated 3.3 V (not exposed on the board pads).

### Compiling

Enter the `blinky` directory and use:

```
egtool build
```

or

```
GOENV=../go.env go build
```

to compile the blinky programm. You can do the same with any other example or in your own program directory (copy/create the `go.env` configuration file in there).

As an effect of the compilation you obtain an ELF file.

### SWD flashing

To run your program from the onboard QSPI flash the begginign of the flash must contain some additional information. We provide them in the `mbr.img` file.

Use the below two commands for flashing:

```
egtool hex -inc ../mbr.img:0x60000000
pyocd load blinky.hex
```

The first one create the `blinky.hex` file that contains both MBR and your program code. The second one runs pyOCD to program the onboard QSPI flash with it. See also the `load_pyocd.sh` script.

### UART/USB flashing

You can use the `sdphost` and `blhost` commands together with the `flashloader` binary (they are provided by NXP) to flash the HEX file via UART or USB if you added it to the board. The boot mode must be set to *serial downloader*. UART flashing is slow. See the `load-flashloader.sh` for more information.