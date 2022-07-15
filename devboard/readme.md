## Support for I.MX RT development boards *WORK IN PROGRESS*

### Directory structure

Every board directory contains a set of packages (in *board* subdirectory) that provides the interface to the peripherals available on the board (for now the support is modest: only LEDs and buttons).

The board/init package, when imported, configures the whole system for typical usage. If you use any other package from *board* directory the board/init package is imported implicitly to ensure the board is properly configured.

The *examples* subdirectory as the name suggests includes sample code, but also scripts and configuration that help to build, load and debug.

There is also *doc* subdirectory that contain useful information and other resources about the development board.

### Supported boards

[fet1061](fet1061): Forlinx FET1061–S System On Module: [MIMXRT1061CVL5B](https://www.nxp.com/part/MIMXRT1061CVL5B#/) + 4 MB QSPI NOR Flash [W25Q32JVSIQ](https://www.winbond.com/resource-files/w25q32jv%20spi%20revc%2008302016.pdf), [website](https://www.forlinx.net/product/imx-rt1061-system-on-module-44.html)

![FET1061-S](fet1061/doc/board.jpg)
