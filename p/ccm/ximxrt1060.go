// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build imxrt1060

package ccm

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/p/mmap"
)

type Periph struct {
	CCR    mmio.R32[CCR]
	_      uint32
	CSR    mmio.R32[CSR]
	CCSR   mmio.R32[CCSR]
	CACRR  mmio.R32[CACRR]
	CBCDR  mmio.R32[CBCDR]
	CBCMR  mmio.R32[CBCMR]
	CSCMR1 mmio.R32[CSCMR1]
	CSCMR2 mmio.R32[CSCMR2]
	CSCDR1 mmio.R32[CSCDR1]
	CS1CDR mmio.R32[CS1CDR]
	CS2CDR mmio.R32[CS2CDR]
	CDCDR  mmio.R32[CDCDR]
	_      uint32
	CSCDR2 mmio.R32[CSCDR2]
	CSCDR3 mmio.R32[CSCDR3]
	_      [2]uint32
	CDHIPR mmio.R32[CDHIPR]
	_      [2]uint32
	CLPCR  mmio.R32[CLPCR]
	CISR   mmio.R32[CIR]
	CIMR   mmio.R32[CIR]
	CCOSR  mmio.R32[CCOSR]
	CGPR   mmio.R32[CGPR]
	CCGR0  mmio.R32[CCGR0]
	CCGR1  mmio.R32[CCGR1]
	CCGR2  mmio.R32[CCGR2]
	CCGR3  mmio.R32[CCGR3]
	CCGR4  mmio.R32[CCGR4]
	CCGR5  mmio.R32[CCGR5]
	CCGR6  mmio.R32[CCGR6]
	CCGR7  mmio.R32[CCGR7]
	CMEOR  mmio.R32[CMEOR]
}

func CCM() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.CCM_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type CCR uint32

func OSCNT_(p *Periph) mmio.RM32[CCR]            { return mmio.RM32[CCR]{&p.CCR, OSCNT} }
func COSC_EN_(p *Periph) mmio.RM32[CCR]          { return mmio.RM32[CCR]{&p.CCR, COSC_EN} }
func REG_BYPASS_COUNT_(p *Periph) mmio.RM32[CCR] { return mmio.RM32[CCR]{&p.CCR, REG_BYPASS_COUNT} }
func RBC_EN_(p *Periph) mmio.RM32[CCR]           { return mmio.RM32[CCR]{&p.CCR, RBC_EN} }

type CSR uint32

func REF_EN_B_(p *Periph) mmio.RM32[CSR]    { return mmio.RM32[CSR]{&p.CSR, REF_EN_B} }
func CAMP2_READY_(p *Periph) mmio.RM32[CSR] { return mmio.RM32[CSR]{&p.CSR, CAMP2_READY} }
func COSC_READY_(p *Periph) mmio.RM32[CSR]  { return mmio.RM32[CSR]{&p.CSR, COSC_READY} }

type CCSR uint32

func PLL3_SW_CLK_SEL_(p *Periph) mmio.RM32[CCSR] { return mmio.RM32[CCSR]{&p.CCSR, PLL3_SW_CLK_SEL} }

type CACRR uint32

func ARM_PODF_(p *Periph) mmio.RM32[CACRR] { return mmio.RM32[CACRR]{&p.CACRR, ARM_PODF} }

type CBCDR uint32

func SEMC_CLK_SEL_(p *Periph) mmio.RM32[CBCDR] { return mmio.RM32[CBCDR]{&p.CBCDR, SEMC_CLK_SEL} }
func SEMC_ALT_CLK_SEL_(p *Periph) mmio.RM32[CBCDR] {
	return mmio.RM32[CBCDR]{&p.CBCDR, SEMC_ALT_CLK_SEL}
}
func IPG_PODF_(p *Periph) mmio.RM32[CBCDR]       { return mmio.RM32[CBCDR]{&p.CBCDR, IPG_PODF} }
func AHB_PODF_(p *Periph) mmio.RM32[CBCDR]       { return mmio.RM32[CBCDR]{&p.CBCDR, AHB_PODF} }
func SEMC_PODF_(p *Periph) mmio.RM32[CBCDR]      { return mmio.RM32[CBCDR]{&p.CBCDR, SEMC_PODF} }
func PERIPH_CLK_SEL_(p *Periph) mmio.RM32[CBCDR] { return mmio.RM32[CBCDR]{&p.CBCDR, PERIPH_CLK_SEL} }
func PERIPH_CLK2_PODF_(p *Periph) mmio.RM32[CBCDR] {
	return mmio.RM32[CBCDR]{&p.CBCDR, PERIPH_CLK2_PODF}
}

type CBCMR uint32

func LPSPI_CLK_SEL_(p *Periph) mmio.RM32[CBCMR] { return mmio.RM32[CBCMR]{&p.CBCMR, LPSPI_CLK_SEL} }
func FLEXSPI2_CLK_SEL_(p *Periph) mmio.RM32[CBCMR] {
	return mmio.RM32[CBCMR]{&p.CBCMR, FLEXSPI2_CLK_SEL}
}
func PERIPH_CLK2_SEL_(p *Periph) mmio.RM32[CBCMR] { return mmio.RM32[CBCMR]{&p.CBCMR, PERIPH_CLK2_SEL} }
func TRACE_CLK_SEL_(p *Periph) mmio.RM32[CBCMR]   { return mmio.RM32[CBCMR]{&p.CBCMR, TRACE_CLK_SEL} }
func PRE_PERIPH_CLK_SEL_(p *Periph) mmio.RM32[CBCMR] {
	return mmio.RM32[CBCMR]{&p.CBCMR, PRE_PERIPH_CLK_SEL}
}
func LCDIF_PODF_(p *Periph) mmio.RM32[CBCMR]    { return mmio.RM32[CBCMR]{&p.CBCMR, LCDIF_PODF} }
func LPSPI_PODF_(p *Periph) mmio.RM32[CBCMR]    { return mmio.RM32[CBCMR]{&p.CBCMR, LPSPI_PODF} }
func FLEXSPI2_PODF_(p *Periph) mmio.RM32[CBCMR] { return mmio.RM32[CBCMR]{&p.CBCMR, FLEXSPI2_PODF} }

type CSCMR1 uint32

func PERCLK_PODF_(p *Periph) mmio.RM32[CSCMR1] { return mmio.RM32[CSCMR1]{&p.CSCMR1, PERCLK_PODF} }
func PERCLK_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] {
	return mmio.RM32[CSCMR1]{&p.CSCMR1, PERCLK_CLK_SEL}
}
func SAI1_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] { return mmio.RM32[CSCMR1]{&p.CSCMR1, SAI1_CLK_SEL} }
func SAI2_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] { return mmio.RM32[CSCMR1]{&p.CSCMR1, SAI2_CLK_SEL} }
func SAI3_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] { return mmio.RM32[CSCMR1]{&p.CSCMR1, SAI3_CLK_SEL} }
func USDHC1_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] {
	return mmio.RM32[CSCMR1]{&p.CSCMR1, USDHC1_CLK_SEL}
}
func USDHC2_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] {
	return mmio.RM32[CSCMR1]{&p.CSCMR1, USDHC2_CLK_SEL}
}
func FLEXSPI_PODF_(p *Periph) mmio.RM32[CSCMR1] { return mmio.RM32[CSCMR1]{&p.CSCMR1, FLEXSPI_PODF} }
func FLEXSPI_CLK_SEL_(p *Periph) mmio.RM32[CSCMR1] {
	return mmio.RM32[CSCMR1]{&p.CSCMR1, FLEXSPI_CLK_SEL}
}

type CSCMR2 uint32

func CAN_CLK_PODF_(p *Periph) mmio.RM32[CSCMR2] { return mmio.RM32[CSCMR2]{&p.CSCMR2, CAN_CLK_PODF} }
func CAN_CLK_SEL_(p *Periph) mmio.RM32[CSCMR2]  { return mmio.RM32[CSCMR2]{&p.CSCMR2, CAN_CLK_SEL} }
func FLEXIO2_CLK_SEL_(p *Periph) mmio.RM32[CSCMR2] {
	return mmio.RM32[CSCMR2]{&p.CSCMR2, FLEXIO2_CLK_SEL}
}

type CSCDR1 uint32

func UART_CLK_PODF_(p *Periph) mmio.RM32[CSCDR1] { return mmio.RM32[CSCDR1]{&p.CSCDR1, UART_CLK_PODF} }
func UART_CLK_SEL_(p *Periph) mmio.RM32[CSCDR1]  { return mmio.RM32[CSCDR1]{&p.CSCDR1, UART_CLK_SEL} }
func USDHC1_PODF_(p *Periph) mmio.RM32[CSCDR1]   { return mmio.RM32[CSCDR1]{&p.CSCDR1, USDHC1_PODF} }
func USDHC2_PODF_(p *Periph) mmio.RM32[CSCDR1]   { return mmio.RM32[CSCDR1]{&p.CSCDR1, USDHC2_PODF} }
func TRACE_PODF_(p *Periph) mmio.RM32[CSCDR1]    { return mmio.RM32[CSCDR1]{&p.CSCDR1, TRACE_PODF} }

type CS1CDR uint32

func SAI1_CLK_PODF_(p *Periph) mmio.RM32[CS1CDR] { return mmio.RM32[CS1CDR]{&p.CS1CDR, SAI1_CLK_PODF} }
func SAI1_CLK_PRED_(p *Periph) mmio.RM32[CS1CDR] { return mmio.RM32[CS1CDR]{&p.CS1CDR, SAI1_CLK_PRED} }
func FLEXIO2_CLK_PRED_(p *Periph) mmio.RM32[CS1CDR] {
	return mmio.RM32[CS1CDR]{&p.CS1CDR, FLEXIO2_CLK_PRED}
}
func SAI3_CLK_PODF_(p *Periph) mmio.RM32[CS1CDR] { return mmio.RM32[CS1CDR]{&p.CS1CDR, SAI3_CLK_PODF} }
func SAI3_CLK_PRED_(p *Periph) mmio.RM32[CS1CDR] { return mmio.RM32[CS1CDR]{&p.CS1CDR, SAI3_CLK_PRED} }
func FLEXIO2_CLK_PODF_(p *Periph) mmio.RM32[CS1CDR] {
	return mmio.RM32[CS1CDR]{&p.CS1CDR, FLEXIO2_CLK_PODF}
}

type CS2CDR uint32

func SAI2_CLK_PODF_(p *Periph) mmio.RM32[CS2CDR] { return mmio.RM32[CS2CDR]{&p.CS2CDR, SAI2_CLK_PODF} }
func SAI2_CLK_PRED_(p *Periph) mmio.RM32[CS2CDR] { return mmio.RM32[CS2CDR]{&p.CS2CDR, SAI2_CLK_PRED} }

type CDCDR uint32

func FLEXIO1_CLK_SEL_(p *Periph) mmio.RM32[CDCDR] { return mmio.RM32[CDCDR]{&p.CDCDR, FLEXIO1_CLK_SEL} }
func FLEXIO1_CLK_PODF_(p *Periph) mmio.RM32[CDCDR] {
	return mmio.RM32[CDCDR]{&p.CDCDR, FLEXIO1_CLK_PODF}
}
func FLEXIO1_CLK_PRED_(p *Periph) mmio.RM32[CDCDR] {
	return mmio.RM32[CDCDR]{&p.CDCDR, FLEXIO1_CLK_PRED}
}
func SPDIF0_CLK_SEL_(p *Periph) mmio.RM32[CDCDR]  { return mmio.RM32[CDCDR]{&p.CDCDR, SPDIF0_CLK_SEL} }
func SPDIF0_CLK_PODF_(p *Periph) mmio.RM32[CDCDR] { return mmio.RM32[CDCDR]{&p.CDCDR, SPDIF0_CLK_PODF} }
func SPDIF0_CLK_PRED_(p *Periph) mmio.RM32[CDCDR] { return mmio.RM32[CDCDR]{&p.CDCDR, SPDIF0_CLK_PRED} }

type CSCDR2 uint32

func LCDIF_PRED_(p *Periph) mmio.RM32[CSCDR2] { return mmio.RM32[CSCDR2]{&p.CSCDR2, LCDIF_PRED} }
func LCDIF_PRE_CLK_SEL_(p *Periph) mmio.RM32[CSCDR2] {
	return mmio.RM32[CSCDR2]{&p.CSCDR2, LCDIF_PRE_CLK_SEL}
}
func LPI2C_CLK_SEL_(p *Periph) mmio.RM32[CSCDR2] { return mmio.RM32[CSCDR2]{&p.CSCDR2, LPI2C_CLK_SEL} }
func LPI2C_CLK_PODF_(p *Periph) mmio.RM32[CSCDR2] {
	return mmio.RM32[CSCDR2]{&p.CSCDR2, LPI2C_CLK_PODF}
}

type CSCDR3 uint32

func CSI_CLK_SEL_(p *Periph) mmio.RM32[CSCDR3] { return mmio.RM32[CSCDR3]{&p.CSCDR3, CSI_CLK_SEL} }
func CSI_PODF_(p *Periph) mmio.RM32[CSCDR3]    { return mmio.RM32[CSCDR3]{&p.CSCDR3, CSI_PODF} }

type CDHIPR uint32

func SEMC_PODF_BUSY_(p *Periph) mmio.RM32[CDHIPR] {
	return mmio.RM32[CDHIPR]{&p.CDHIPR, SEMC_PODF_BUSY}
}
func AHB_PODF_BUSY_(p *Periph) mmio.RM32[CDHIPR] { return mmio.RM32[CDHIPR]{&p.CDHIPR, AHB_PODF_BUSY} }
func PERIPH2_CLK_SEL_BUSY_(p *Periph) mmio.RM32[CDHIPR] {
	return mmio.RM32[CDHIPR]{&p.CDHIPR, PERIPH2_CLK_SEL_BUSY}
}
func PERIPH_CLK_SEL_BUSY_(p *Periph) mmio.RM32[CDHIPR] {
	return mmio.RM32[CDHIPR]{&p.CDHIPR, PERIPH_CLK_SEL_BUSY}
}
func ARM_PODF_BUSY_(p *Periph) mmio.RM32[CDHIPR] { return mmio.RM32[CDHIPR]{&p.CDHIPR, ARM_PODF_BUSY} }

type CLPCR uint32

func LPM_(p *Periph) mmio.RM32[CLPCR] { return mmio.RM32[CLPCR]{&p.CLPCR, LPM} }
func ARM_CLK_DIS_ON_LPM_(p *Periph) mmio.RM32[CLPCR] {
	return mmio.RM32[CLPCR]{&p.CLPCR, ARM_CLK_DIS_ON_LPM}
}
func SBYOS_(p *Periph) mmio.RM32[CLPCR]          { return mmio.RM32[CLPCR]{&p.CLPCR, SBYOS} }
func DIS_REF_OSC_(p *Periph) mmio.RM32[CLPCR]    { return mmio.RM32[CLPCR]{&p.CLPCR, DIS_REF_OSC} }
func VSTBY_(p *Periph) mmio.RM32[CLPCR]          { return mmio.RM32[CLPCR]{&p.CLPCR, VSTBY} }
func STBY_COUNT_(p *Periph) mmio.RM32[CLPCR]     { return mmio.RM32[CLPCR]{&p.CLPCR, STBY_COUNT} }
func COSC_PWRDOWN_(p *Periph) mmio.RM32[CLPCR]   { return mmio.RM32[CLPCR]{&p.CLPCR, COSC_PWRDOWN} }
func BYPASS_LPM_HS1_(p *Periph) mmio.RM32[CLPCR] { return mmio.RM32[CLPCR]{&p.CLPCR, BYPASS_LPM_HS1} }
func BYPASS_LPM_HS0_(p *Periph) mmio.RM32[CLPCR] { return mmio.RM32[CLPCR]{&p.CLPCR, BYPASS_LPM_HS0} }
func MASK_CORE0_WFI_(p *Periph) mmio.RM32[CLPCR] { return mmio.RM32[CLPCR]{&p.CLPCR, MASK_CORE0_WFI} }
func MASK_SCU_IDLE_(p *Periph) mmio.RM32[CLPCR]  { return mmio.RM32[CLPCR]{&p.CLPCR, MASK_SCU_IDLE} }
func MASK_L2CC_IDLE_(p *Periph) mmio.RM32[CLPCR] { return mmio.RM32[CLPCR]{&p.CLPCR, MASK_L2CC_IDLE} }

type CIR uint32

type CCOSR uint32

func CLKO1_SEL_(p *Periph) mmio.RM32[CCOSR]   { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO1_SEL} }
func CLKO1_DIV_(p *Periph) mmio.RM32[CCOSR]   { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO1_DIV} }
func CLKO1_EN_(p *Periph) mmio.RM32[CCOSR]    { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO1_EN} }
func CLK_OUT_SEL_(p *Periph) mmio.RM32[CCOSR] { return mmio.RM32[CCOSR]{&p.CCOSR, CLK_OUT_SEL} }
func CLKO2_SEL_(p *Periph) mmio.RM32[CCOSR]   { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO2_SEL} }
func CLKO2_DIV_(p *Periph) mmio.RM32[CCOSR]   { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO2_DIV} }
func CLKO2_EN_(p *Periph) mmio.RM32[CCOSR]    { return mmio.RM32[CCOSR]{&p.CCOSR, CLKO2_EN} }

type CGPR uint32

func PMIC_DELAY_SCALER_(p *Periph) mmio.RM32[CGPR] {
	return mmio.RM32[CGPR]{&p.CGPR, PMIC_DELAY_SCALER}
}
func EFUSE_PROG_SUPPLY_GATE_(p *Periph) mmio.RM32[CGPR] {
	return mmio.RM32[CGPR]{&p.CGPR, EFUSE_PROG_SUPPLY_GATE}
}
func SYS_MEM_DS_CTRL_(p *Periph) mmio.RM32[CGPR] { return mmio.RM32[CGPR]{&p.CGPR, SYS_MEM_DS_CTRL} }
func FPL_(p *Periph) mmio.RM32[CGPR]             { return mmio.RM32[CGPR]{&p.CGPR, FPL} }
func INT_MEM_CLK_LPM_(p *Periph) mmio.RM32[CGPR] { return mmio.RM32[CGPR]{&p.CGPR, INT_MEM_CLK_LPM} }

type CCGR0 uint32

func CG0_0_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_0} }
func CG0_1_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_1} }
func CG0_2_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_2} }
func CG0_3_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_3} }
func CG0_4_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_4} }
func CG0_5_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_5} }
func CG0_6_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_6} }
func CG0_7_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_7} }
func CG0_8_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_8} }
func CG0_9_(p *Periph) mmio.RM32[CCGR0]  { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_9} }
func CG0_10_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_10} }
func CG0_11_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_11} }
func CG0_12_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_12} }
func CG0_13_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_13} }
func CG0_14_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_14} }
func CG0_15_(p *Periph) mmio.RM32[CCGR0] { return mmio.RM32[CCGR0]{&p.CCGR0, CG0_15} }

type CCGR1 uint32

func CG1_0_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_0} }
func CG1_1_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_1} }
func CG1_2_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_2} }
func CG1_3_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_3} }
func CG1_4_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_4} }
func CG1_5_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_5} }
func CG1_6_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_6} }
func CG1_7_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_7} }
func CG1_8_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_8} }
func CG1_9_(p *Periph) mmio.RM32[CCGR1]  { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_9} }
func CG1_10_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_10} }
func CG1_11_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_11} }
func CG1_12_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_12} }
func CG1_13_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_13} }
func CG1_14_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_14} }
func CG1_15_(p *Periph) mmio.RM32[CCGR1] { return mmio.RM32[CCGR1]{&p.CCGR1, CG1_15} }

type CCGR2 uint32

func CG2_0_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_0} }
func CG2_1_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_1} }
func CG2_2_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_2} }
func CG2_3_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_3} }
func CG2_4_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_4} }
func CG2_5_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_5} }
func CG2_6_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_6} }
func CG2_7_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_7} }
func CG2_8_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_8} }
func CG2_9_(p *Periph) mmio.RM32[CCGR2]  { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_9} }
func CG2_10_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_10} }
func CG2_11_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_11} }
func CG2_12_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_12} }
func CG2_13_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_13} }
func CG2_14_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_14} }
func CG2_15_(p *Periph) mmio.RM32[CCGR2] { return mmio.RM32[CCGR2]{&p.CCGR2, CG2_15} }

type CCGR3 uint32

func CG3_0_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_0} }
func CG3_1_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_1} }
func CG3_2_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_2} }
func CG3_3_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_3} }
func CG3_4_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_4} }
func CG3_5_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_5} }
func CG3_6_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_6} }
func CG3_7_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_7} }
func CG3_8_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_8} }
func CG3_9_(p *Periph) mmio.RM32[CCGR3]  { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_9} }
func CG3_10_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_10} }
func CG3_11_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_11} }
func CG3_12_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_12} }
func CG3_13_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_13} }
func CG3_14_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_14} }
func CG3_15_(p *Periph) mmio.RM32[CCGR3] { return mmio.RM32[CCGR3]{&p.CCGR3, CG3_15} }

type CCGR4 uint32

func CG4_0_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_0} }
func CG4_1_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_1} }
func CG4_2_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_2} }
func CG4_3_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_3} }
func CG4_4_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_4} }
func CG4_5_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_5} }
func CG4_6_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_6} }
func CG4_7_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_7} }
func CG4_8_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_8} }
func CG4_9_(p *Periph) mmio.RM32[CCGR4]  { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_9} }
func CG4_10_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_10} }
func CG4_11_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_11} }
func CG4_12_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_12} }
func CG4_13_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_13} }
func CG4_14_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_14} }
func CG4_15_(p *Periph) mmio.RM32[CCGR4] { return mmio.RM32[CCGR4]{&p.CCGR4, CG4_15} }

type CCGR5 uint32

func CG5_0_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_0} }
func CG5_1_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_1} }
func CG5_2_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_2} }
func CG5_3_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_3} }
func CG5_4_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_4} }
func CG5_5_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_5} }
func CG5_6_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_6} }
func CG5_7_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_7} }
func CG5_8_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_8} }
func CG5_9_(p *Periph) mmio.RM32[CCGR5]  { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_9} }
func CG5_10_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_10} }
func CG5_11_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_11} }
func CG5_12_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_12} }
func CG5_13_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_13} }
func CG5_14_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_14} }
func CG5_15_(p *Periph) mmio.RM32[CCGR5] { return mmio.RM32[CCGR5]{&p.CCGR5, CG5_15} }

type CCGR6 uint32

func CG6_0_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_0} }
func CG6_1_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_1} }
func CG6_2_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_2} }
func CG6_3_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_3} }
func CG6_4_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_4} }
func CG6_5_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_5} }
func CG6_6_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_6} }
func CG6_7_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_7} }
func CG6_8_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_8} }
func CG6_9_(p *Periph) mmio.RM32[CCGR6]  { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_9} }
func CG6_10_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_10} }
func CG6_11_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_11} }
func CG6_12_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_12} }
func CG6_13_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_13} }
func CG6_14_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_14} }
func CG6_15_(p *Periph) mmio.RM32[CCGR6] { return mmio.RM32[CCGR6]{&p.CCGR6, CG6_15} }

type CCGR7 uint32

func CG7_0_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_0} }
func CG7_1_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_1} }
func CG7_2_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_2} }
func CG7_3_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_3} }
func CG7_4_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_4} }
func CG7_5_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_5} }
func CG7_6_(p *Periph) mmio.RM32[CCGR7] { return mmio.RM32[CCGR7]{&p.CCGR7, CG7_6} }

type CMEOR uint32

func MOD_EN_OV_GPT_(p *Periph) mmio.RM32[CMEOR]  { return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_GPT} }
func MOD_EN_OV_PIT_(p *Periph) mmio.RM32[CMEOR]  { return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_PIT} }
func MOD_EN_USDHC_(p *Periph) mmio.RM32[CMEOR]   { return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_USDHC} }
func MOD_EN_OV_TRNG_(p *Periph) mmio.RM32[CMEOR] { return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_TRNG} }
func MOD_EN_OV_CANFD_CPI_(p *Periph) mmio.RM32[CMEOR] {
	return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_CANFD_CPI}
}
func MOD_EN_OV_CAN2_CPI_(p *Periph) mmio.RM32[CMEOR] {
	return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_CAN2_CPI}
}
func MOD_EN_OV_CAN1_CPI_(p *Periph) mmio.RM32[CMEOR] {
	return mmio.RM32[CMEOR]{&p.CMEOR, MOD_EN_OV_CAN1_CPI}
}
