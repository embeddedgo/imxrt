```
Num Ball Signal            GPIO       Description        Default
--------------------------------------------------------------------------------
1   -    5V_SYS            -          power supply       -
2   -    5V_SYS            -          power supply       -
3   -    GND               -          ground             -
4   -    GND               -          ground             -
5   -    GND               -          ground             -
6   K10  K10_CSI_HSYNC     GPIO1_IO23 GPIO               GPIO1_IO23
7   J12  J12_CSI_VSYNC     GPIO1_IO22 GPIO               GPIO1_IO22
8   K11  K11_I2C1_SDA      GPIO1_IO17 IC data            I2C1_SDA
9   J11  J11_I2C1_SCL      GPIO1_IO16 IC clock           I2C1_SCL
10  -    GND               -          ground             -
11  M14  M14_GPIO1_IO00    GPIO1_IO00 GPIO               GPIO1_IO00
12  L11  L11_UART2_TX      GPIO1_IO18 UART TxD           UART2_TX
13  M12  M12_UART2_RX      GPIO1_IO19 UART RxD           UART2_RX
14  L10  L10_CAN2_RX       GPIO1_IO15 CAN RxD            CAN2_RX
15  H14  H14_CAN2_TX       GPIO1_IO14 CAN TxD            CAN2_TX
16  D13  D13_SD1_CD_B      GPIO2_IO28 SD card detect     SD1_CD_B
17  K7   K7_PMIC_ON_REQ    GPIO5_IO01 PMIC enable power  PMIC_ON_REQ
18  D14  D14_WDOG_B        GPIO2_IO29 WDOG extrenal sig  WDOG_B (*)
19  B14  B14_PWM3          GPIO2_IO31 PWM                PWM3
20  -    POR_BUTTON        -          reset button       -
21  M7   M7_POR_B          -          power on reset     POR_B
22  C14  C14_GPIO2_IO30    GPIO2_IO30 GPIO               GPIO2_IO30
23  L14  L14_UART1_RXD     GPIO1_IO13 UART RxD           UART1_RXD
24  K14  K14_UART1_TXD     GPIO1_IO12 UART TxD           UART1_TXD
25  -    GND               -          ground             -
26  -    GND               -          ground             -
27  -    VDD_COIN_3V3      -          RTC power input    VDD_COIN_3V3
28  G11  G11_GPIO1_IO03    GPIO1_IO03 GPIO               GPIO1_IO03
29  M11  M11_PWM0          GPIO1_IO02 PWM                PWM0
30  G10  G10_JTAG_TRSTB    GPIO1_IO11 JTAG reset input   JTAG_TRSTB
31  F14  F14_JTAG_TDI      GPIO1_IO09 JTAG data input    JTAG_TDI (LED5)
32  F12  F12_JTAG_TCK      GPIO1_IO07 JTAG clock input   JTAG_TCK
33  E14  E14_JTAG_TMS      GPIO1_IO06 JTAG mode select   JTAG_TMS
34  F13  F13_JTAG_MOD      GPIO1_IO08 JTAG common/1149.1 JTAG_MOD
35  G13  G13_JTAG_TDO      GPIO1_IO10 JTAG data output   JTAG_TDO
36  -    GND               -          ground             -
37  A12  A12_ENET_TX_DATA1 GPIO2_IO24 Ethernet TxD 1     ENET_TX_DATA1
38  B12  B12_ENET_TX_DATA0 GPIO2_IO23 Ethernet TxD 0     ENET_TX_DATA0
39  A13  A13_ENET_TX_EN    GPIO2_IO25 Ethernet Tx enable ENET_TX_EN
40  C13  C13_ENET_RX_ER    GPIO2_IO27 Ethernet Rx error  ENET_RX_ER
41  C12  C12_ENET_RX_EN    GPIO2_IO22 Ethernet Rx enable ENET_RX_EN
42  E12  E12_ENET_RX_DATA0 GPIO2_IO20 Ethernet RxD 0     ENET_RX_DATA0
43  D12  D12_ENET_RX_DATA1 GPIO2_IO21 Ethernet RxD 1     ENET_RX_DATA1
44  B13  B13_ENET_TX_CLK   GPIO2_IO26 Ethernet Tx clock  ENET_TX_CLK
45  C7   C7_SEMC_CSX0      GPIO3_IO27 Ethernet SMI data  ENET_MDIO
46  A7   A7_SEMC_RDY       GPIO3_IO26 Ethernet SMI clock ENET_MDC
47  B7   B7_SEMC_DQS       GPIO3_IO25 GPIO               GPIO3_IO25
48  F11  F11_BOOT_MODE0    GPIO1_IO04 BOOT_MODE[0]       BOOT_MODE0 (10K to GND)
49  G14  G14_BOOT_MODE1    GPIO1_IO05 BOOT_MODE[1]       BOOT_MODE1 (10K to 3V3)
50  -    GND               -          ground             -
51  -    GND               -          ground             -
52  D9   D9_LCDIF_DATA06   GPIO2_IO10 boot config 6      BT_CFG6
53  A10  A10_LCDIF_DATA07  GPIO2_IO11 boot config 7      BT_CFG7
54  D10  D10_LCDIF_DATA09  GPIO2_IO13 Ethernet TxD 1     ENET2_TX_DATA1
55  C10  C10_LCDIF_DATA08  GPIO2_IO12 Ethernet TxD 0     ENET2_TX_DATA0
56  E10  E10_LCDIF_DATA10  GPIO2_IO14 Ethernet Tx enable ENET2_TX_EN
57  A11  A11_LCDIF_DATA12  GPIO2_IO16 Ethernet Rx error  ENET2_RX_ER
58  D11  D11_LCDIF_DATA15  GPIO2_IO19 Ethernet Rx enable ENET2_RX_EN
59  B11  B11_LCDIF_DATA13  GPIO2_IO17 Ethernet RxD 0     ENET2_RX_DATA0
60  C11  C11_LCDIF_DATA14  GPIO2_IO18 Ethernet RxD 1     ENET2_RX_DATA1
61  E11  E11_LCDIF_DATA11  GPIO2_IO15 Ethernet Tx clock  ENET2_TX_CLK
62  E7   E7_LCDIF_DE       GPIO2_IO01 Ethernet SMI data  ENET2_MDIO
63  D7   D7_LCDIF_CLK      GPIO2_IO00 Ethernet SMI clock ENET2_MDC
64  E8   E8_LCDIF_HSYNC    GPIO2_IO02 GPIO               GPIO2_IO02
65  L6   L6_WAKEUP         GPIO5_IO00 wake up from SVNS  WAKEUP
66  M6   M6_ONOFF          -          on/off button      ONOFF
67  -    GND               -          ground             -
68  K1   K1_SD1_D1         GPIO3_IO15 uSDHC data 1       SD1_D1
69  J1   J1_SD1_D0         GPIO3_IO14 uSDHC data 0       SD1_D0
70  J3   J3_SD1_CLK        GPIO3_IO13 uSDHC clock        SD1_CLK
71  J4   J4_SD1_CMD        GPIO3_IO12 uSDHC command      SD1_CMD
72  J2   J2_SD1_D3         GPIO3_IO17 uSDHC data 3       SD1_D3
73  H2   H2_SD1_D2         GPIO3_IO16 uSDHC data 2       SD1_D2
74  -    GND               -          ground             -
75  -    GND               -          ground             -
76  -    GND               -          ground             -
77  P2   P2_FlexSPI_CLK_B  GPIO3_IO04 QSPI_B clock       FlexSPI_CLK_B
78  M3   M3_FlexSPI_D1_B   GPIO3_IO02 QSPI_B data 1      FlexSPI_D1_B
79  M4   M4_FlexSPI_D0_B   GPIO3_IO03 QSPI_B data 0      FlexSPI_D0_B
80  L5   L5_FlexSPI_D3_B   GPIO3_IO00 QSPI_B data 3      FlexSPI_D3_B
81  M5   M5_FlexSPI_D2_B   GPIO3_IO01 QSPI_B data 2      FlexSPI_D2_B
82  -    GND               -          ground             -
83  P6   P6_OTG2_VBUS      -          USB power 5V       OTG2_VBUS
84  N7   N7_OTG2_D_N       -          USB data-          OTG2_D_N
85  P7   P7_OTG2_D_P       -          USB data+          OTG2_D_P
86  N6   N6_OTG1_VBUS      -          USB power 5V       OTG1_VBUS
87  M8   M8_OTG1_D_N       -          USB data-          OTG1_D_N
88  L8   L8_OTG1_D_P       -          USB data+          OTG1_D_P
89  H10  H10_OTG1_ID       GPIO1_IO01 USB OTG ID pin     OTG1_ID
82  -    GND               -          ground             -
91  H11  H11_ADC2_02       GPIO1_IO29 ADC TSC Y+         TOUCH_Y+
92  G12  G12_ADC2_03       GPIO1_IO30 ADC TSC X-         TOUCH_X-
93  H12  H12_ADC2_01       GPIO1_IO28 ADC TSC Y-         TOUCH_Y-
94  J14  J14_ADC2_04       GPIO1_IO31 ADC TSC X+         TOUCH_X+
95  J13  J13_CSI_DATA06    GPIO1_IO27 CSI data 6         CSI_DATA06
96  L12  L12_CSI_PIXCLK    GPIO1_IO20 CSI pixel clock    CSI_PIXCLK
97  L13  L13_CSI_DATA07    GPIO1_IO26 CSI data 7         CSI_DATA07
98  M13  M13_CSI_DATA08    GPIO1_IO25 CSI data 8         CSI_DATA08
99  K12  K12_CSI_MCLK      GPIO1_IO21 CSI master clock   CSI_MCLK
100 H13  H13_CSI_DATA09    GPIO1_IO24 CSI data 9         CSI_DATA09
```

(\*) The pins marked with `(*)` in the `Default` column are the pins used by the core board and you should avoid to use them, otherwise it may cause abnormal startup and other problems.
