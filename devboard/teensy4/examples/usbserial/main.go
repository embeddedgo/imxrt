// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/hal/system/dtcm"
	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/pmu"
	"github.com/embeddedgo/imxrt/p/usb"
	"github.com/embeddedgo/imxrt/p/usbphy"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
)

func main() {
	// IO pins
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")

	time.Sleep(20 * time.Millisecond) // wait at least 20ms before starting USB

	fmt.Println("Start!")
	initUSB()

	for {
		time.Sleep(time.Second)
	}
}

func initUSB() {

	CCM := ccm.CCM()
	CCMA := ccm_analog.CCM_ANALOG()
	PMU := pmu.PMU()
	USB := usb.USB1()
	PHY := usbphy.USBPHY1()

	// Ungate all necessary clocks.
	CCM.CCGR6.SetBits(ccm.CG6_0 | ccm.CG6_11) // USB (usboh3) | CCMA (anadig)
	CCMA.PLL_USB1_SET.Store(ccm_analog.PLL_USB_EN_USB_CLKS)

	// Enable internal 3V0 regulator
	const (
		out3v000 = 15 << pmu.OUTPUT_TRGn
		boo0v150 = 6 << pmu.BO_OFFSETn
	)
	PMU.REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)

	//USB.BURSTSIZE.Store(4<<usb.TXPBURSTn | 4<<usb.RXPBURSTn)

	fmt.Printf("PLL_USB1:   %032b\n", CCMA.PLL_USB1.Load())
	fmt.Printf("PFD_480:    %08x\n", CCMA.PFD_480.Load())
	fmt.Printf("PLL_USB2:   %032b\n", CCMA.PLL_USB2.Load())
	fmt.Printf("USB.ID:     %08x\n", USB.ID.Load())
	fmt.Printf("\nUSB.USBMODE:%08x\n", USB.USBMODE.Load())
	fmt.Printf("USB.BURSTSIZE:%08x\n", USB.BURSTSIZE.Load())
	fmt.Printf("PHY.PWD:    %08x\n", PHY.PWD.Load())
	fmt.Printf("PHY.TX:     %08x\n", PHY.TX.Load())
	fmt.Printf("PHY.RX:     %08x\n", PHY.RX.Load())
	fmt.Printf("PHY.CTRL:   %08x\n", PHY.CTRL.Load())

	fmt.Print("\nUSB reset... ")

	PHY.CTRL_SET.Store(usbphy.SFTRST)
	USB.USBCMD.SetBits(usb.RST)

	count := 0
	for USB.USBCMD.LoadBits(usb.RST) != 0 {
		count++
	}
	PHY.CTRL_CLR.Store(usbphy.SFTRST)
	fmt.Println(count)
	time.Sleep(25 * time.Millisecond)

	fmt.Printf("\nUSB.USBMODE:%08x\n", USB.USBMODE.Load())
	fmt.Printf("PHY.PWD:    %08x\n", PHY.PWD.Load())
	fmt.Printf("PHY.TX:     %08x\n", PHY.TX.Load())
	fmt.Printf("PHY.RX:     %08x\n", PHY.RX.Load())
	fmt.Printf("PHY.CTRL:   %08x\n", PHY.CTRL.Load())

	PHY.CTRL_CLR.Store(usbphy.CLKGATE)
	PHY.PWD.Store(0)
	USB.USBMODE.Store(usb.CM_2 | usb.SLOM)

	dQHList := dtcm.MakeSlice[dQH](4096, 4, 4)
	dQHList[0].epcap = 64<<maxPktLenn | 1<<intOnSetupn
	dQHList[1].epcap = 64 << maxPktLenn

	USB.ASYNC_ENDPTLISTADDR.Store(uint32(uintptr(unsafe.Pointer(&dQHList[0]))))

	irq.USB_OTG1.Enable(rtos.IntPrioLow, 0)

	USB.USBINTR.Store(usb.UE | usb.UEE | usb.URE | usb.SLE)

	USB.USBCMD.Store(usb.RS)

	for i := 0; ; i++ {
		print("i=", i, " ", USB.ASYNC_ENDPTLISTADDR.Load(), "\r\n")
		time.Sleep(time.Second)
	}
}

// Endpoint Queue Head (dQH) must be 64 byte aligned in memory.
type dQH struct {
	epcap   uint32
	current uintptr
	next    uintptr
	token   uint32
	bufp0   uintptr
	bufp1   uintptr
	bufp2   uintptr
	bufp3   uintptr
	bufp4   uintptr
	_       uint32
	setup0  uint32
	setup1  uint32
	_       [4]uint32 // increase the struct size to 64 bytes
}

const (
	intOnSetupn  = 15
	maxPktLenn   = 16
	zeroLenTermn = 29
	multn        = 30
)

//go:interrupthandler
func USB_OTG1_Handler() {
	u := usb.USB1()
	status := u.USBSTS.Load()
	u.USBSTS.Store(status)

	leds.User.Toggle()
	print("USB_OTG1_Handler: ", status, "\r\n")

	if status&usb.UI != 0 {
	}
	if status&usb.URI != 0 {

	}
	if status&usb.TI0 != 0 {

	}
	if status&usb.TI1 != 0 {

	}
	if status&usb.PCI != 0 {

	}
	if status&usb.SLI != 0 {

	}
	if status&usb.UEI != 0 {

	}
	if u.USBINTR.LoadBits(usb.SRE) != 0 && status&usb.SRI != 0 {
		// reboot
	}

}
