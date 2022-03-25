package main

import (
	"embedded/mmio"
	"unsafe"
)

func mmio32(addr uintptr) *mmio.U32 {
	return (*mmio.U32)(unsafe.Pointer(addr))
}

const (
	IOMUXC_GPR_ADDR uintptr = 0x400AC000
	IOMUXC_ADDR     uintptr = 0x401F8000
	GPIO6_ADDR      uintptr = 0x42000000
)

type GPIO struct {
	DR        mmio.U32
	GDIR      mmio.U32
	PSR       mmio.U32
	ICR1      mmio.U32
	ICR2      mmio.U32
	IMR       mmio.U32
	ISR       mmio.U32
	EDGE_SEL  mmio.U32
	_         [25]mmio.U32
	DR_SET    mmio.U32
	DR_CLEAR  mmio.U32
	DR_TOGGLE mmio.U32
}

var (
	IOMUXC_SW_MUX_CTL_PAD_GPIO_AD_B0_09 = mmio32(IOMUXC_ADDR + 0x0E0)
	IOMUXC_SW_PAD_CTL_PAD_GPIO_AD_B0_09 = mmio32(IOMUXC_ADDR + 0x2D0)
	IOMUXC_GPR_GPR26                    = mmio32(IOMUXC_GPR_ADDR + 0x068)
	GPIO6                               = (*GPIO)(unsafe.Pointer(GPIO6_ADDR))
)

func main() {
	// Configure GPIO AD_B0_09 (PAD F14) for output
	IOMUXC_SW_MUX_CTL_PAD_GPIO_AD_B0_09.Store(5)
	IOMUXC_SW_PAD_CTL_PAD_GPIO_AD_B0_09.Store(7 << 3)
	IOMUXC_GPR_GPR26.Store(0xFFFFFFFF)
	GPIO6.GDIR.SetBit(9)

	for {
		for i := 0; i < 1e6; i++ {
			GPIO6.DR_CLEAR.Store(1 << 9)
		}
		for i := 0; i < 4e6; i++ {
			GPIO6.DR_SET.Store(1 << 9)
		}
	}
}
