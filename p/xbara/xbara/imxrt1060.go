// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build imxrt1060

// Package xbara provides access to the registers of the XBARA peripheral.
//
// Instances:
//
//	XBARA1  XBARA1_BASE  -  -  Crossbar Switch
//
// Registers:
//
//	0x000 16  SEL0   Crossbar A Select Register 0
//	0x002 16  SEL1   Crossbar A Select Register 1
//	0x004 16  SEL2   Crossbar A Select Register 2
//	0x006 16  SEL3   Crossbar A Select Register 3
//	0x008 16  SEL4   Crossbar A Select Register 4
//	0x00A 16  SEL5   Crossbar A Select Register 5
//	0x00C 16  SEL6   Crossbar A Select Register 6
//	0x00E 16  SEL7   Crossbar A Select Register 7
//	0x010 16  SEL8   Crossbar A Select Register 8
//	0x012 16  SEL9   Crossbar A Select Register 9
//	0x014 16  SEL10  Crossbar A Select Register 10
//	0x016 16  SEL11  Crossbar A Select Register 11
//	0x018 16  SEL12  Crossbar A Select Register 12
//	0x01A 16  SEL13  Crossbar A Select Register 13
//	0x01C 16  SEL14  Crossbar A Select Register 14
//	0x01E 16  SEL15  Crossbar A Select Register 15
//	0x020 16  SEL16  Crossbar A Select Register 16
//	0x022 16  SEL17  Crossbar A Select Register 17
//	0x024 16  SEL18  Crossbar A Select Register 18
//	0x026 16  SEL19  Crossbar A Select Register 19
//	0x028 16  SEL20  Crossbar A Select Register 20
//	0x02A 16  SEL21  Crossbar A Select Register 21
//	0x02C 16  SEL22  Crossbar A Select Register 22
//	0x02E 16  SEL23  Crossbar A Select Register 23
//	0x030 16  SEL24  Crossbar A Select Register 24
//	0x032 16  SEL25  Crossbar A Select Register 25
//	0x034 16  SEL26  Crossbar A Select Register 26
//	0x036 16  SEL27  Crossbar A Select Register 27
//	0x038 16  SEL28  Crossbar A Select Register 28
//	0x03A 16  SEL29  Crossbar A Select Register 29
//	0x03C 16  SEL30  Crossbar A Select Register 30
//	0x03E 16  SEL31  Crossbar A Select Register 31
//	0x040 16  SEL32  Crossbar A Select Register 32
//	0x042 16  SEL33  Crossbar A Select Register 33
//	0x044 16  SEL34  Crossbar A Select Register 34
//	0x046 16  SEL35  Crossbar A Select Register 35
//	0x048 16  SEL36  Crossbar A Select Register 36
//	0x04A 16  SEL37  Crossbar A Select Register 37
//	0x04C 16  SEL38  Crossbar A Select Register 38
//	0x04E 16  SEL39  Crossbar A Select Register 39
//	0x050 16  SEL40  Crossbar A Select Register 40
//	0x052 16  SEL41  Crossbar A Select Register 41
//	0x054 16  SEL42  Crossbar A Select Register 42
//	0x056 16  SEL43  Crossbar A Select Register 43
//	0x058 16  SEL44  Crossbar A Select Register 44
//	0x05A 16  SEL45  Crossbar A Select Register 45
//	0x05C 16  SEL46  Crossbar A Select Register 46
//	0x05E 16  SEL47  Crossbar A Select Register 47
//	0x060 16  SEL48  Crossbar A Select Register 48
//	0x062 16  SEL49  Crossbar A Select Register 49
//	0x064 16  SEL50  Crossbar A Select Register 50
//	0x066 16  SEL51  Crossbar A Select Register 51
//	0x068 16  SEL52  Crossbar A Select Register 52
//	0x06A 16  SEL53  Crossbar A Select Register 53
//	0x06C 16  SEL54  Crossbar A Select Register 54
//	0x06E 16  SEL55  Crossbar A Select Register 55
//	0x070 16  SEL56  Crossbar A Select Register 56
//	0x072 16  SEL57  Crossbar A Select Register 57
//	0x074 16  SEL58  Crossbar A Select Register 58
//	0x076 16  SEL59  Crossbar A Select Register 59
//	0x078 16  SEL60  Crossbar A Select Register 60
//	0x07A 16  SEL61  Crossbar A Select Register 61
//	0x07C 16  SEL62  Crossbar A Select Register 62
//	0x07E 16  SEL63  Crossbar A Select Register 63
//	0x080 16  SEL64  Crossbar A Select Register 64
//	0x082 16  SEL65  Crossbar A Select Register 65
//	0x084 16  CTRL0  Crossbar A Control Register 0
//	0x086 16  CTRL1  Crossbar A Control Register 1
//
// Import:
//
//	github.com/embeddedgo/imxrt/p/mmap
package xbara

const (
	SEL0 SEL0 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT0 (refer to Functional Description section for input/output assignment)
	SEL1 SEL0 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT1 (refer to Functional Description section for input/output assignment)
)

const (
	SEL0n = 0
	SEL1n = 8
)

const (
	SEL2 SEL1 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT2 (refer to Functional Description section for input/output assignment)
	SEL3 SEL1 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT3 (refer to Functional Description section for input/output assignment)
)

const (
	SEL2n = 0
	SEL3n = 8
)

const (
	SEL4 SEL2 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT4 (refer to Functional Description section for input/output assignment)
	SEL5 SEL2 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT5 (refer to Functional Description section for input/output assignment)
)

const (
	SEL4n = 0
	SEL5n = 8
)

const (
	SEL6 SEL3 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT6 (refer to Functional Description section for input/output assignment)
	SEL7 SEL3 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT7 (refer to Functional Description section for input/output assignment)
)

const (
	SEL6n = 0
	SEL7n = 8
)

const (
	SEL8 SEL4 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT8 (refer to Functional Description section for input/output assignment)
	SEL9 SEL4 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT9 (refer to Functional Description section for input/output assignment)
)

const (
	SEL8n = 0
	SEL9n = 8
)

const (
	SEL10 SEL5 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT10 (refer to Functional Description section for input/output assignment)
	SEL11 SEL5 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT11 (refer to Functional Description section for input/output assignment)
)

const (
	SEL10n = 0
	SEL11n = 8
)

const (
	SEL12 SEL6 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT12 (refer to Functional Description section for input/output assignment)
	SEL13 SEL6 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT13 (refer to Functional Description section for input/output assignment)
)

const (
	SEL12n = 0
	SEL13n = 8
)

const (
	SEL14 SEL7 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT14 (refer to Functional Description section for input/output assignment)
	SEL15 SEL7 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT15 (refer to Functional Description section for input/output assignment)
)

const (
	SEL14n = 0
	SEL15n = 8
)

const (
	SEL16 SEL8 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT16 (refer to Functional Description section for input/output assignment)
	SEL17 SEL8 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT17 (refer to Functional Description section for input/output assignment)
)

const (
	SEL16n = 0
	SEL17n = 8
)

const (
	SEL18 SEL9 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT18 (refer to Functional Description section for input/output assignment)
	SEL19 SEL9 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT19 (refer to Functional Description section for input/output assignment)
)

const (
	SEL18n = 0
	SEL19n = 8
)

const (
	SEL20 SEL10 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT20 (refer to Functional Description section for input/output assignment)
	SEL21 SEL10 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT21 (refer to Functional Description section for input/output assignment)
)

const (
	SEL20n = 0
	SEL21n = 8
)

const (
	SEL22 SEL11 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT22 (refer to Functional Description section for input/output assignment)
	SEL23 SEL11 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT23 (refer to Functional Description section for input/output assignment)
)

const (
	SEL22n = 0
	SEL23n = 8
)

const (
	SEL24 SEL12 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT24 (refer to Functional Description section for input/output assignment)
	SEL25 SEL12 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT25 (refer to Functional Description section for input/output assignment)
)

const (
	SEL24n = 0
	SEL25n = 8
)

const (
	SEL26 SEL13 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT26 (refer to Functional Description section for input/output assignment)
	SEL27 SEL13 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT27 (refer to Functional Description section for input/output assignment)
)

const (
	SEL26n = 0
	SEL27n = 8
)

const (
	SEL28 SEL14 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT28 (refer to Functional Description section for input/output assignment)
	SEL29 SEL14 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT29 (refer to Functional Description section for input/output assignment)
)

const (
	SEL28n = 0
	SEL29n = 8
)

const (
	SEL30 SEL15 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT30 (refer to Functional Description section for input/output assignment)
	SEL31 SEL15 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT31 (refer to Functional Description section for input/output assignment)
)

const (
	SEL30n = 0
	SEL31n = 8
)

const (
	SEL32 SEL16 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT32 (refer to Functional Description section for input/output assignment)
	SEL33 SEL16 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT33 (refer to Functional Description section for input/output assignment)
)

const (
	SEL32n = 0
	SEL33n = 8
)

const (
	SEL34 SEL17 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT34 (refer to Functional Description section for input/output assignment)
	SEL35 SEL17 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT35 (refer to Functional Description section for input/output assignment)
)

const (
	SEL34n = 0
	SEL35n = 8
)

const (
	SEL36 SEL18 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT36 (refer to Functional Description section for input/output assignment)
	SEL37 SEL18 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT37 (refer to Functional Description section for input/output assignment)
)

const (
	SEL36n = 0
	SEL37n = 8
)

const (
	SEL38 SEL19 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT38 (refer to Functional Description section for input/output assignment)
	SEL39 SEL19 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT39 (refer to Functional Description section for input/output assignment)
)

const (
	SEL38n = 0
	SEL39n = 8
)

const (
	SEL40 SEL20 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT40 (refer to Functional Description section for input/output assignment)
	SEL41 SEL20 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT41 (refer to Functional Description section for input/output assignment)
)

const (
	SEL40n = 0
	SEL41n = 8
)

const (
	SEL42 SEL21 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT42 (refer to Functional Description section for input/output assignment)
	SEL43 SEL21 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT43 (refer to Functional Description section for input/output assignment)
)

const (
	SEL42n = 0
	SEL43n = 8
)

const (
	SEL44 SEL22 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT44 (refer to Functional Description section for input/output assignment)
	SEL45 SEL22 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT45 (refer to Functional Description section for input/output assignment)
)

const (
	SEL44n = 0
	SEL45n = 8
)

const (
	SEL46 SEL23 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT46 (refer to Functional Description section for input/output assignment)
	SEL47 SEL23 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT47 (refer to Functional Description section for input/output assignment)
)

const (
	SEL46n = 0
	SEL47n = 8
)

const (
	SEL48 SEL24 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT48 (refer to Functional Description section for input/output assignment)
	SEL49 SEL24 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT49 (refer to Functional Description section for input/output assignment)
)

const (
	SEL48n = 0
	SEL49n = 8
)

const (
	SEL50 SEL25 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT50 (refer to Functional Description section for input/output assignment)
	SEL51 SEL25 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT51 (refer to Functional Description section for input/output assignment)
)

const (
	SEL50n = 0
	SEL51n = 8
)

const (
	SEL52 SEL26 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT52 (refer to Functional Description section for input/output assignment)
	SEL53 SEL26 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT53 (refer to Functional Description section for input/output assignment)
)

const (
	SEL52n = 0
	SEL53n = 8
)

const (
	SEL54 SEL27 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT54 (refer to Functional Description section for input/output assignment)
	SEL55 SEL27 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT55 (refer to Functional Description section for input/output assignment)
)

const (
	SEL54n = 0
	SEL55n = 8
)

const (
	SEL56 SEL28 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT56 (refer to Functional Description section for input/output assignment)
	SEL57 SEL28 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT57 (refer to Functional Description section for input/output assignment)
)

const (
	SEL56n = 0
	SEL57n = 8
)

const (
	SEL58 SEL29 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT58 (refer to Functional Description section for input/output assignment)
	SEL59 SEL29 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT59 (refer to Functional Description section for input/output assignment)
)

const (
	SEL58n = 0
	SEL59n = 8
)

const (
	SEL60 SEL30 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT60 (refer to Functional Description section for input/output assignment)
	SEL61 SEL30 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT61 (refer to Functional Description section for input/output assignment)
)

const (
	SEL60n = 0
	SEL61n = 8
)

const (
	SEL62 SEL31 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT62 (refer to Functional Description section for input/output assignment)
	SEL63 SEL31 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT63 (refer to Functional Description section for input/output assignment)
)

const (
	SEL62n = 0
	SEL63n = 8
)

const (
	SEL64 SEL32 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT64 (refer to Functional Description section for input/output assignment)
	SEL65 SEL32 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT65 (refer to Functional Description section for input/output assignment)
)

const (
	SEL64n = 0
	SEL65n = 8
)

const (
	SEL66 SEL33 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT66 (refer to Functional Description section for input/output assignment)
	SEL67 SEL33 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT67 (refer to Functional Description section for input/output assignment)
)

const (
	SEL66n = 0
	SEL67n = 8
)

const (
	SEL68 SEL34 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT68 (refer to Functional Description section for input/output assignment)
	SEL69 SEL34 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT69 (refer to Functional Description section for input/output assignment)
)

const (
	SEL68n = 0
	SEL69n = 8
)

const (
	SEL70 SEL35 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT70 (refer to Functional Description section for input/output assignment)
	SEL71 SEL35 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT71 (refer to Functional Description section for input/output assignment)
)

const (
	SEL70n = 0
	SEL71n = 8
)

const (
	SEL72 SEL36 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT72 (refer to Functional Description section for input/output assignment)
	SEL73 SEL36 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT73 (refer to Functional Description section for input/output assignment)
)

const (
	SEL72n = 0
	SEL73n = 8
)

const (
	SEL74 SEL37 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT74 (refer to Functional Description section for input/output assignment)
	SEL75 SEL37 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT75 (refer to Functional Description section for input/output assignment)
)

const (
	SEL74n = 0
	SEL75n = 8
)

const (
	SEL76 SEL38 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT76 (refer to Functional Description section for input/output assignment)
	SEL77 SEL38 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT77 (refer to Functional Description section for input/output assignment)
)

const (
	SEL76n = 0
	SEL77n = 8
)

const (
	SEL78 SEL39 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT78 (refer to Functional Description section for input/output assignment)
	SEL79 SEL39 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT79 (refer to Functional Description section for input/output assignment)
)

const (
	SEL78n = 0
	SEL79n = 8
)

const (
	SEL80 SEL40 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT80 (refer to Functional Description section for input/output assignment)
	SEL81 SEL40 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT81 (refer to Functional Description section for input/output assignment)
)

const (
	SEL80n = 0
	SEL81n = 8
)

const (
	SEL82 SEL41 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT82 (refer to Functional Description section for input/output assignment)
	SEL83 SEL41 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT83 (refer to Functional Description section for input/output assignment)
)

const (
	SEL82n = 0
	SEL83n = 8
)

const (
	SEL84 SEL42 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT84 (refer to Functional Description section for input/output assignment)
	SEL85 SEL42 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT85 (refer to Functional Description section for input/output assignment)
)

const (
	SEL84n = 0
	SEL85n = 8
)

const (
	SEL86 SEL43 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT86 (refer to Functional Description section for input/output assignment)
	SEL87 SEL43 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT87 (refer to Functional Description section for input/output assignment)
)

const (
	SEL86n = 0
	SEL87n = 8
)

const (
	SEL88 SEL44 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT88 (refer to Functional Description section for input/output assignment)
	SEL89 SEL44 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT89 (refer to Functional Description section for input/output assignment)
)

const (
	SEL88n = 0
	SEL89n = 8
)

const (
	SEL90 SEL45 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT90 (refer to Functional Description section for input/output assignment)
	SEL91 SEL45 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT91 (refer to Functional Description section for input/output assignment)
)

const (
	SEL90n = 0
	SEL91n = 8
)

const (
	SEL92 SEL46 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT92 (refer to Functional Description section for input/output assignment)
	SEL93 SEL46 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT93 (refer to Functional Description section for input/output assignment)
)

const (
	SEL92n = 0
	SEL93n = 8
)

const (
	SEL94 SEL47 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT94 (refer to Functional Description section for input/output assignment)
	SEL95 SEL47 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT95 (refer to Functional Description section for input/output assignment)
)

const (
	SEL94n = 0
	SEL95n = 8
)

const (
	SEL96 SEL48 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT96 (refer to Functional Description section for input/output assignment)
	SEL97 SEL48 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT97 (refer to Functional Description section for input/output assignment)
)

const (
	SEL96n = 0
	SEL97n = 8
)

const (
	SEL98 SEL49 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT98 (refer to Functional Description section for input/output assignment)
	SEL99 SEL49 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT99 (refer to Functional Description section for input/output assignment)
)

const (
	SEL98n = 0
	SEL99n = 8
)

const (
	SEL100 SEL50 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT100 (refer to Functional Description section for input/output assignment)
	SEL101 SEL50 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT101 (refer to Functional Description section for input/output assignment)
)

const (
	SEL100n = 0
	SEL101n = 8
)

const (
	SEL102 SEL51 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT102 (refer to Functional Description section for input/output assignment)
	SEL103 SEL51 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT103 (refer to Functional Description section for input/output assignment)
)

const (
	SEL102n = 0
	SEL103n = 8
)

const (
	SEL104 SEL52 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT104 (refer to Functional Description section for input/output assignment)
	SEL105 SEL52 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT105 (refer to Functional Description section for input/output assignment)
)

const (
	SEL104n = 0
	SEL105n = 8
)

const (
	SEL106 SEL53 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT106 (refer to Functional Description section for input/output assignment)
	SEL107 SEL53 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT107 (refer to Functional Description section for input/output assignment)
)

const (
	SEL106n = 0
	SEL107n = 8
)

const (
	SEL108 SEL54 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT108 (refer to Functional Description section for input/output assignment)
	SEL109 SEL54 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT109 (refer to Functional Description section for input/output assignment)
)

const (
	SEL108n = 0
	SEL109n = 8
)

const (
	SEL110 SEL55 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT110 (refer to Functional Description section for input/output assignment)
	SEL111 SEL55 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT111 (refer to Functional Description section for input/output assignment)
)

const (
	SEL110n = 0
	SEL111n = 8
)

const (
	SEL112 SEL56 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT112 (refer to Functional Description section for input/output assignment)
	SEL113 SEL56 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT113 (refer to Functional Description section for input/output assignment)
)

const (
	SEL112n = 0
	SEL113n = 8
)

const (
	SEL114 SEL57 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT114 (refer to Functional Description section for input/output assignment)
	SEL115 SEL57 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT115 (refer to Functional Description section for input/output assignment)
)

const (
	SEL114n = 0
	SEL115n = 8
)

const (
	SEL116 SEL58 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT116 (refer to Functional Description section for input/output assignment)
	SEL117 SEL58 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT117 (refer to Functional Description section for input/output assignment)
)

const (
	SEL116n = 0
	SEL117n = 8
)

const (
	SEL118 SEL59 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT118 (refer to Functional Description section for input/output assignment)
	SEL119 SEL59 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT119 (refer to Functional Description section for input/output assignment)
)

const (
	SEL118n = 0
	SEL119n = 8
)

const (
	SEL120 SEL60 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT120 (refer to Functional Description section for input/output assignment)
	SEL121 SEL60 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT121 (refer to Functional Description section for input/output assignment)
)

const (
	SEL120n = 0
	SEL121n = 8
)

const (
	SEL122 SEL61 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT122 (refer to Functional Description section for input/output assignment)
	SEL123 SEL61 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT123 (refer to Functional Description section for input/output assignment)
)

const (
	SEL122n = 0
	SEL123n = 8
)

const (
	SEL124 SEL62 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT124 (refer to Functional Description section for input/output assignment)
	SEL125 SEL62 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT125 (refer to Functional Description section for input/output assignment)
)

const (
	SEL124n = 0
	SEL125n = 8
)

const (
	SEL126 SEL63 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT126 (refer to Functional Description section for input/output assignment)
	SEL127 SEL63 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT127 (refer to Functional Description section for input/output assignment)
)

const (
	SEL126n = 0
	SEL127n = 8
)

const (
	SEL128 SEL64 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT128 (refer to Functional Description section for input/output assignment)
	SEL129 SEL64 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT129 (refer to Functional Description section for input/output assignment)
)

const (
	SEL128n = 0
	SEL129n = 8
)

const (
	SEL130 SEL65 = 0x7F << 0 //+ Input (XBARA_INn) to be muxed to XBARA_OUT130 (refer to Functional Description section for input/output assignment)
	SEL131 SEL65 = 0x7F << 8 //+ Input (XBARA_INn) to be muxed to XBARA_OUT131 (refer to Functional Description section for input/output assignment)
)

const (
	SEL130n = 0
	SEL131n = 8
)

const (
	DEN0    CTRL0 = 0x01 << 0  //+ DMA Enable for XBAR_OUT0
	IEN0    CTRL0 = 0x01 << 1  //+ Interrupt Enable for XBAR_OUT0
	EDGE0   CTRL0 = 0x03 << 2  //+ Active edge for edge detection on XBAR_OUT0
	EDGE0_0 CTRL0 = 0x00 << 2  //  STS0 never asserts
	EDGE0_1 CTRL0 = 0x01 << 2  //  STS0 asserts on rising edges of XBAR_OUT0
	EDGE0_2 CTRL0 = 0x02 << 2  //  STS0 asserts on falling edges of XBAR_OUT0
	EDGE0_3 CTRL0 = 0x03 << 2  //  STS0 asserts on rising and falling edges of XBAR_OUT0
	STS0    CTRL0 = 0x01 << 4  //+ Edge detection status for XBAR_OUT0
	DEN1    CTRL0 = 0x01 << 8  //+ DMA Enable for XBAR_OUT1
	IEN1    CTRL0 = 0x01 << 9  //+ Interrupt Enable for XBAR_OUT1
	EDGE1   CTRL0 = 0x03 << 10 //+ Active edge for edge detection on XBAR_OUT1
	EDGE1_0 CTRL0 = 0x00 << 10 //  STS1 never asserts
	EDGE1_1 CTRL0 = 0x01 << 10 //  STS1 asserts on rising edges of XBAR_OUT1
	EDGE1_2 CTRL0 = 0x02 << 10 //  STS1 asserts on falling edges of XBAR_OUT1
	EDGE1_3 CTRL0 = 0x03 << 10 //  STS1 asserts on rising and falling edges of XBAR_OUT1
	STS1    CTRL0 = 0x01 << 12 //+ Edge detection status for XBAR_OUT1
)

const (
	DEN0n  = 0
	IEN0n  = 1
	EDGE0n = 2
	STS0n  = 4
	DEN1n  = 8
	IEN1n  = 9
	EDGE1n = 10
	STS1n  = 12
)

const (
	DEN2    CTRL1 = 0x01 << 0  //+ DMA Enable for XBAR_OUT2
	IEN2    CTRL1 = 0x01 << 1  //+ Interrupt Enable for XBAR_OUT2
	EDGE2   CTRL1 = 0x03 << 2  //+ Active edge for edge detection on XBAR_OUT2
	EDGE2_0 CTRL1 = 0x00 << 2  //  STS2 never asserts
	EDGE2_1 CTRL1 = 0x01 << 2  //  STS2 asserts on rising edges of XBAR_OUT2
	EDGE2_2 CTRL1 = 0x02 << 2  //  STS2 asserts on falling edges of XBAR_OUT2
	EDGE2_3 CTRL1 = 0x03 << 2  //  STS2 asserts on rising and falling edges of XBAR_OUT2
	STS2    CTRL1 = 0x01 << 4  //+ Edge detection status for XBAR_OUT2
	DEN3    CTRL1 = 0x01 << 8  //+ DMA Enable for XBAR_OUT3
	IEN3    CTRL1 = 0x01 << 9  //+ Interrupt Enable for XBAR_OUT3
	EDGE3   CTRL1 = 0x03 << 10 //+ Active edge for edge detection on XBAR_OUT3
	EDGE3_0 CTRL1 = 0x00 << 10 //  STS3 never asserts
	EDGE3_1 CTRL1 = 0x01 << 10 //  STS3 asserts on rising edges of XBAR_OUT3
	EDGE3_2 CTRL1 = 0x02 << 10 //  STS3 asserts on falling edges of XBAR_OUT3
	EDGE3_3 CTRL1 = 0x03 << 10 //  STS3 asserts on rising and falling edges of XBAR_OUT3
	STS3    CTRL1 = 0x01 << 12 //+ Edge detection status for XBAR_OUT3
)

const (
	DEN2n  = 0
	IEN2n  = 1
	EDGE2n = 2
	STS2n  = 4
	DEN3n  = 8
	IEN3n  = 9
	EDGE3n = 10
	STS3n  = 12
)
