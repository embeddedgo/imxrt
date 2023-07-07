// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/mmio"
	"embedded/rtos"
	"fmt"
	"time"
	"unsafe"

	"github.com/embeddedgo/imxrt/hal/dma"
	"github.com/embeddedgo/imxrt/hal/dtcm"
	"github.com/embeddedgo/imxrt/hal/irq"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/p/ccm"
	"github.com/embeddedgo/imxrt/p/ccm_analog"
	"github.com/embeddedgo/imxrt/p/pmu"
	"github.com/embeddedgo/imxrt/p/usb"
	"github.com/embeddedgo/imxrt/p/usbphy"

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
	u := usb.USB1()
	phy := usbphy.USBPHY1()

	// Ungate all necessary clocks.
	CCM.CCGR6.SetBits(ccm.CG6_0 | ccm.CG6_11) // USB (usboh3) | CCMA (anadig)
	CCMA.PLL_USB1_SET.Store(ccm_analog.PLL_USB_EN_USB_CLKS)

	// Enable internal 3V0 regulator
	const (
		out3v000 = 15 << pmu.OUTPUT_TRGn
		boo0v150 = 6 << pmu.BO_OFFSETn
	)
	PMU.REG_3P0.Store(out3v000 | boo0v150 | pmu.ENABLE_LINREG)

	//u.BURSTSIZE.Store(4<<usb.TXPBURSTn | 4<<usb.RXPBURSTn)

	fmt.Println("USB reset")
	phy.CTRL_SET.Store(usbphy.SFTRST)
	u.USBCMD.SetBits(usb.RST)
	for u.USBCMD.LoadBits(usb.RST) != 0 {
	}
	phy.CTRL_CLR.Store(usbphy.SFTRST)

	time.Sleep(25 * time.Millisecond)

	fmt.Printf("\nUSB.USBMODE:%08x\n", u.USBMODE.Load())
	fmt.Printf("PHY.PWD:    %08x\n", phy.PWD.Load())
	fmt.Printf("PHY.TX:     %08x\n", phy.TX.Load())
	fmt.Printf("PHY.RX:     %08x\n", phy.RX.Load())
	fmt.Printf("PHY.CTRL:   %08x\n", phy.CTRL.Load())

	phy.CTRL_CLR.Store(usbphy.CLKGATE)
	phy.PWD.Store(0)
	u.USBMODE.Store(usb.CM_2 | usb.SLOM)

	dQHList = dtcm.MakeSlice[dQH](4096, 5, 5)
	dQHList[0].config = ep0size<<maxPktLenn | 1<<intOnSetupn
	dQHList[1].config = ep0size << maxPktLenn
	descBuf = dtcm.MakeSlice[byte](32, 256, 256)
	ep0data = dtcm.New[dTD](32)
	ep0ack = dtcm.New[dTD](32)

	u.ASYNC_ENDPTLISTADDR.Store(uint32(uintptr(unsafe.Pointer(&dQHList[0]))))

	irq.USB_OTG1.Enable(rtos.IntPrioLow, 0)

	u.USBINTR.Store(usb.UE | usb.UEE | usb.URE | usb.SLE)

	u.USBCMD.Store(usb.RS)

	for i := 0; ; i++ {
		time.Sleep(time.Second)
	}
}

type dQH struct {
	config  uint32
	current uintptr
	next    uintptr
	_       [7]uint32 // overlay area
	setup   [2]uint32
	_       [4]uint32 // padding to make dQH 64 bytes in size, unused
}

var (
	dQHList []dQH
	descBuf []byte
	ep0data *dTD
	ep0ack  *dTD
)

type dTD struct {
	next  uintptr
	token uint32
	bufp  [5]uintptr
	_     uint32
}

// Endpoint Queue Head (dQH) must be 64 byte aligned in memory.
type DQH struct {
	epcap   uint32
	current uintptr
	dTD
	setup [2]uint32
	note  *rtos.Note
	_     [3]uint32 // increase the struct size to 64 bytes
}

func (qh *DQH) InitRx(pktSize int, zlt bool) {

}

const (
	intOnSetupn  = 15
	maxPktLenn   = 16
	zeroLenTermn = 29
	multn        = 30
)

var (
	ep0notifyMask usb.ENDPTCOMPLETE
	epNnotifyMask usb.ENDPTCOMPLETE
	configuration uint8
	highSpeed     bool
)

//go:interrupthandler
func USB_OTG1_Handler() {
	u := usb.USB1()
	status := u.USBSTS.Load()
	u.USBSTS.Store(status)

	if status&usb.UI != 0 {
		print("UI\r\n")
		for {
			ess := u.ENDPTSETUPSTAT.Load()
			if ess == 0 {
				break
			}
			u.ENDPTSETUPSTAT.Store(ess)
			var setup0, setup1 uint32
			for {
				u.USBCMD.SetBits(usb.SUTW)
				setup0 = dQHList[0].setup[0]
				setup1 = dQHList[0].setup[1]
				if u.USBCMD.LoadBits(usb.SUTW) != 0 {
					break
				}
			}
			u.USBCMD.ClearBits(usb.SUTW)
			u.ENDPTFLUSH.Store(1<<usb.FETBn | 1<<usb.FERBn)
			for u.ENDPTFLUSH.LoadBits(1<<usb.FETBn|1<<usb.FERBn) != 0 {
			}
			ep0notifyMask = 0
			ep0setup(setup0, setup1)
		}
		if ec := u.ENDPTCOMPLETE.Load(); ec != 0 {
			u.ENDPTCOMPLETE.Store(ec)
			if ec&ep0notifyMask != 0 {
				ep0notifyMask = 0
				ep0complete()
			}
			ec &= epNnotifyMask

			if ec != 0 {
				// TODO..
			}
			if ec != 0 {
				// TODO..
			}
		}
	}
	if status&usb.URI != 0 {
		print("URI\r\n")
		u.ENDPTSETUPSTAT.Store(u.ENDPTSETUPSTAT.Load())
		u.ENDPTCOMPLETE.Store(u.ENDPTCOMPLETE.Load())
		for u.ENDPTPRIME.Load() != 0 {
		}
		u.ENDPTFLUSH.Store(0xffff_ffff)
		epNnotifyMask = 0
	}
	if status&usb.TI0 != 0 {
		print("TI0\r\n")
	}
	if status&usb.TI1 != 0 {
		print("TI1\r\n")
	}
	if status&usb.PCI != 0 {
		//print("PCI\r\n")
		if u.PORTSC1.LoadBits(usb.HSP) != 0 {
			//print(" high speed\r\f")
			highSpeed = true
		} else {
			//print(" full speed\r\f")
			highSpeed = false
		}
	}
	if status&usb.SLI != 0 {
		print("SLI\r\n")
	}
	if status&usb.UEI != 0 {
		print("UEI\r\n")
	}
	if u.USBINTR.LoadBits(usb.SRE) != 0 && status&usb.SRI != 0 {
		print("reboot\r\n")
	}
}

func ep0setup(setup0, setup1 uint32) {
	reqType := setup0 & 0xff
	recipient := "reserved"
	switch reqType & 0x1f {
	case 0:
		recipient = "device"
	case 1:
		recipient = "interface"
	case 2:
		recipient = "endpoint"
	case 3:
		recipient = "other"
	}
	typ := "reserved"
	switch reqType >> 5 & 3 {
	case 0:
		typ = "standard"
	case 1:
		typ = "class"
	case 2:
		typ = "vendor"
	}
	dir := "host->dev"
	if reqType>>7&1 != 0 {
		dir = "dev->host"
	}
	print(" recipient=", recipient, " type=", typ, " dir=", dir, "\r\n")

	req := "unknown"
	switch setup0 >> 8 & 0xff {
	case 0:
		req = "GET_STATUS"
	case 1:
		req = "CLEAR_FEATURE"
	case 3:
		req = "SET_FEATURE"
	case 5:
		req = "SET_ADDRESS"
	case 6:
		req = "GET_DESCRIPTOR"
	case 7:
		req = "SET_DESCRIPTOR"
	case 8:
		req = "GET_CONFIGURATION"
	case 9:
		req = "SET_CONFIGURATION"
	}
	print(" req=", req, " val=", setup0>>16, " idx=", setup1&0xffff, " len=", setup1>>16, "\r\n")

	u := usb.USB1()
	switch setup0 & 0xffff {
	case 0x0500: // SET_ADDRESS
		ep0receive(nil, false)
		addr := setup0 >> 16 & 0x7f
		u.DEVADDR_PLISTBASE.Store(1<<24 | addr<<25)
		return
	case 0x0680, 0x0681: // GET_DESCRIPTOR
		vi := setup0&0xffff_0000 | setup1&0x0000_ffff
		d := descriptors[vi]
		if d != nil {
			n := len(d)
			if wLength := int(setup1 >> 16); n > wLength {
				n = wLength
			}
			copy(descBuf, d[:n]) // copy to DTCM
			ep0transmit(descBuf[:n], false)
			return
		}
	case 0x0880: // GET_CONFIGURATION
		descBuf[0] = configuration
		ep0transmit(descBuf[:1], false)
		return
	case 0x0900: // SET_CONFIGURATION
		configuration = uint8(setup0 >> 16)
		u.ENDPTCTRL[2].Store(3<<usb.TXTn | usb.TXR | usb.TXE)
		u.ENDPTCTRL[3].Store(2<<usb.RXTn | usb.RXR | usb.RXE)
		u.ENDPTCTRL[4].Store(2<<usb.TXTn | usb.TXR | usb.TXE)
		usbSerialConfig()
		ep0receive(nil, false)
		return
	default:
		print(" unknown:", setup0&0xffff, "\r\n")
	}
	print("usb stall\r\n")
	u.ENDPTCTRL[0].Store(usb.RXS | usb.TXS) // stall
}

func ep0complete() {
	// ep0setupData
}

func ep0transmit(data []byte, notify bool) {
	u := usb.USB1()
	if len(data) != 0 {
		ep0data.next = 1
		ep0data.token = uint32(len(data)<<16 | 1<<7)
		addr := uintptr(unsafe.Pointer(&data[0]))
		ep0data.bufp[0] = addr
		ep0data.bufp[1] = addr + 4096
		ep0data.bufp[2] = addr + 8192
		ep0data.bufp[3] = addr + 12288
		ep0data.bufp[4] = addr + 16384
		dQHList[1].next = uintptr(unsafe.Pointer(ep0data))
		mmio.MB()
		u.ENDPTPRIME.SetBits(1 << usb.PETBn)
		for u.ENDPTPRIME.Load() != 0 {
		}
	}
	ep0ack.next = 1
	ep0ack.token = 1 << 7
	if notify {
		ep0ack.token |= 1 << 15
	}
	ep0ack.bufp[0] = 0
	dQHList[0].next = uintptr(unsafe.Pointer(ep0ack))
	mmio.MB()
	u.ENDPTCOMPLETE.Store(1<<usb.ERCEn | 1<<usb.ETCEn)
	u.ENDPTPRIME.SetBits(1 << usb.PERBn)
	if notify {
		ep0notifyMask = 1 << usb.PERBn
	} else {
		ep0notifyMask = 0
	}
	for u.ENDPTPRIME.Load() != 0 {
	}
}

func ep0receive(data []byte, notify bool) {
	u := usb.USB1()
	if len(data) != 0 {
		ep0data.next = 1
		ep0data.token = uint32(len(data)<<16 | 1<<7)
		addr := uintptr(unsafe.Pointer(&data[0]))
		ep0data.bufp[0] = addr
		ep0data.bufp[1] = addr + 4096
		ep0data.bufp[2] = addr + 8192
		ep0data.bufp[3] = addr + 12288
		ep0data.bufp[4] = addr + 16384
		dQHList[0].next = uintptr(unsafe.Pointer(ep0data))
		mmio.MB()
		u.ENDPTPRIME.SetBits(1 << usb.PERBn)
		for u.ENDPTPRIME.Load() != 0 {
		}
	}
	ep0ack.next = 1
	ep0ack.token = 1 << 7
	if notify {
		ep0ack.token |= 1 << 15
	}
	ep0ack.bufp[0] = 0
	dQHList[1].next = uintptr(unsafe.Pointer(ep0ack))
	mmio.MB()
	u.ENDPTCOMPLETE.Store(1<<usb.ERCEn | 1<<usb.ETCEn)
	u.ENDPTPRIME.SetBits(1 << usb.PETBn)
	if notify {
		ep0notifyMask = 1 << usb.PETBn
	} else {
		ep0notifyMask = 0
	}
	for u.ENDPTPRIME.Load() != 0 {
	}
}

const (
	usbSerialBufRxOffset = 0
	usbSerialBufTxOffset = usbSerialBufRxOffset + CDC_RX_SIZE_480
	usbSerialBufCAOffset = usbSerialBufTxOffset + CDC_ACM_SIZE
	usbSerialMaxBufLen   = usbSerialBufCAOffset + CDC_ACM_SIZE
)

var usbSerialBuf = dma.MakeSlice[byte](usbSerialMaxBufLen, usbSerialMaxBufLen)

func usbSerialConfig() {
	for i := range usbSerialBuf {
		usbSerialBuf[i] = 0
	}
}
