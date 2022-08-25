// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iomux

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/imxrt/p/mmap"
)

func GPR(i int) *mmio.R32[uint32] {
	return &(*[34]mmio.R32[uint32])(unsafe.Pointer(mmap.IOMUXC_GPR_BASE))[i]
}

// GPR2 bits
const (
	AXBS_L_AHBXL_HIGH_PRIORITY uint32 = 0x01 << 0  //+ AXBS_L AHBXL master has higher priority.Do not set both DMA and AHBXL to high priority.
	AXBS_L_DMA_HIGH_PRIORITY   uint32 = 0x01 << 1  //+ AXBS_L DMA master has higher priority.Do not set both DMA and AHBXL to high priority.
	AXBS_L_FORCE_ROUND_ROBIN   uint32 = 0x01 << 2  //+ Force Round Robin in AXBS_L
	AXBS_P_M0_HIGH_PRIORITY    uint32 = 0x01 << 3  //+ AXBS_P M0 master has higher priority.Do not set both M1 and M0 to high priority.
	AXBS_P_M1_HIGH_PRIORITY    uint32 = 0x01 << 4  //+ AXBS_P M1 master has higher priority.Do not set both M1 and M0 to high priority.
	AXBS_P_FORCE_ROUND_ROBIN   uint32 = 0x01 << 5  //+ Force Round Robin in AXBS_P. This bit can override master M0 M1 high priority configuration.
	CANFD_FILTER_BYPASS        uint32 = 0x01 << 6  //+ Disable CANFD filter
	L2_MEM_EN_POWERSAVING      uint32 = 0x01 << 12 //+ enable power saving features on L2 memory
	RAM_AUTO_CLK_GATING_EN     uint32 = 0x01 << 13 //+ Automatically gate off RAM clock when RAM is not accessed.
	L2_MEM_DEEPSLEEP           uint32 = 0x01 << 14 //+ control how memory enter Deep Sleep mode (shutdown periphery power, but maintain memory contents, outputs of memory are pulled low)
	MQS_CLK_DIV                uint32 = 0xFF << 16 //+ Divider ratio control for mclk from hmclk. mclk frequency = 1/(n+1) * hmclk frequency.
	DIVIDE_1                   uint32 = 0x00 << 16 //  mclk frequency = 1/1 * hmclk frequency
	DIVIDE_2                   uint32 = 0x01 << 16 //  mclk frequency = 1/2 * hmclk frequency
	DIVIDE_3                   uint32 = 0x02 << 16 //  mclk frequency = 1/3 * hmclk frequency
	DIVIDE_4                   uint32 = 0x03 << 16 //  mclk frequency = 1/4 * hmclk frequency
	DIVIDE_5                   uint32 = 0x04 << 16 //  mclk frequency = 1/5 * hmclk frequency
	DIVIDE_6                   uint32 = 0x05 << 16 //  mclk frequency = 1/6 * hmclk frequency
	DIVIDE_7                   uint32 = 0x06 << 16 //  mclk frequency = 1/7 * hmclk frequency
	DIVIDE_8                   uint32 = 0x07 << 16 //  mclk frequency = 1/8 * hmclk frequency
	DIVIDE_9                   uint32 = 0x08 << 16 //  mclk frequency = 1/9 * hmclk frequency
	DIVIDE_10                  uint32 = 0x09 << 16 //  mclk frequency = 1/10 * hmclk frequency
	DIVIDE_11                  uint32 = 0x0A << 16 //  mclk frequency = 1/11 * hmclk frequency
	DIVIDE_12                  uint32 = 0x0B << 16 //  mclk frequency = 1/12 * hmclk frequency
	DIVIDE_13                  uint32 = 0x0C << 16 //  mclk frequency = 1/13 * hmclk frequency
	DIVIDE_14                  uint32 = 0x0D << 16 //  mclk frequency = 1/14 * hmclk frequency
	DIVIDE_15                  uint32 = 0x0E << 16 //  mclk frequency = 1/15 * hmclk frequency
	DIVIDE_16                  uint32 = 0x0F << 16 //  mclk frequency = 1/16 * hmclk frequency
	DIVIDE_17                  uint32 = 0x10 << 16 //  mclk frequency = 1/17 * hmclk frequency
	DIVIDE_18                  uint32 = 0x11 << 16 //  mclk frequency = 1/18 * hmclk frequency
	DIVIDE_19                  uint32 = 0x12 << 16 //  mclk frequency = 1/19 * hmclk frequency
	DIVIDE_20                  uint32 = 0x13 << 16 //  mclk frequency = 1/20 * hmclk frequency
	DIVIDE_21                  uint32 = 0x14 << 16 //  mclk frequency = 1/21 * hmclk frequency
	DIVIDE_22                  uint32 = 0x15 << 16 //  mclk frequency = 1/22 * hmclk frequency
	DIVIDE_23                  uint32 = 0x16 << 16 //  mclk frequency = 1/23 * hmclk frequency
	DIVIDE_24                  uint32 = 0x17 << 16 //  mclk frequency = 1/24 * hmclk frequency
	DIVIDE_25                  uint32 = 0x18 << 16 //  mclk frequency = 1/25 * hmclk frequency
	DIVIDE_26                  uint32 = 0x19 << 16 //  mclk frequency = 1/26 * hmclk frequency
	DIVIDE_27                  uint32 = 0x1A << 16 //  mclk frequency = 1/27 * hmclk frequency
	DIVIDE_28                  uint32 = 0x1B << 16 //  mclk frequency = 1/28 * hmclk frequency
	DIVIDE_29                  uint32 = 0x1C << 16 //  mclk frequency = 1/29 * hmclk frequency
	DIVIDE_30                  uint32 = 0x1D << 16 //  mclk frequency = 1/30 * hmclk frequency
	DIVIDE_31                  uint32 = 0x1E << 16 //  mclk frequency = 1/31 * hmclk frequency
	DIVIDE_32                  uint32 = 0x1F << 16 //  mclk frequency = 1/32 * hmclk frequency
	DIVIDE_33                  uint32 = 0x20 << 16 //  mclk frequency = 1/33 * hmclk frequency
	DIVIDE_34                  uint32 = 0x21 << 16 //  mclk frequency = 1/34 * hmclk frequency
	DIVIDE_35                  uint32 = 0x22 << 16 //  mclk frequency = 1/35 * hmclk frequency
	DIVIDE_36                  uint32 = 0x23 << 16 //  mclk frequency = 1/36 * hmclk frequency
	DIVIDE_37                  uint32 = 0x24 << 16 //  mclk frequency = 1/37 * hmclk frequency
	DIVIDE_38                  uint32 = 0x25 << 16 //  mclk frequency = 1/38 * hmclk frequency
	DIVIDE_39                  uint32 = 0x26 << 16 //  mclk frequency = 1/39 * hmclk frequency
	DIVIDE_40                  uint32 = 0x27 << 16 //  mclk frequency = 1/40 * hmclk frequency
	DIVIDE_41                  uint32 = 0x28 << 16 //  mclk frequency = 1/41 * hmclk frequency
	DIVIDE_42                  uint32 = 0x29 << 16 //  mclk frequency = 1/42 * hmclk frequency
	DIVIDE_43                  uint32 = 0x2A << 16 //  mclk frequency = 1/43 * hmclk frequency
	DIVIDE_44                  uint32 = 0x2B << 16 //  mclk frequency = 1/44 * hmclk frequency
	DIVIDE_45                  uint32 = 0x2C << 16 //  mclk frequency = 1/45 * hmclk frequency
	DIVIDE_46                  uint32 = 0x2D << 16 //  mclk frequency = 1/46 * hmclk frequency
	DIVIDE_47                  uint32 = 0x2E << 16 //  mclk frequency = 1/47 * hmclk frequency
	DIVIDE_48                  uint32 = 0x2F << 16 //  mclk frequency = 1/48 * hmclk frequency
	DIVIDE_49                  uint32 = 0x30 << 16 //  mclk frequency = 1/49 * hmclk frequency
	DIVIDE_50                  uint32 = 0x31 << 16 //  mclk frequency = 1/50 * hmclk frequency
	DIVIDE_51                  uint32 = 0x32 << 16 //  mclk frequency = 1/51 * hmclk frequency
	DIVIDE_52                  uint32 = 0x33 << 16 //  mclk frequency = 1/52 * hmclk frequency
	DIVIDE_53                  uint32 = 0x34 << 16 //  mclk frequency = 1/53 * hmclk frequency
	DIVIDE_54                  uint32 = 0x35 << 16 //  mclk frequency = 1/54 * hmclk frequency
	DIVIDE_55                  uint32 = 0x36 << 16 //  mclk frequency = 1/55 * hmclk frequency
	DIVIDE_56                  uint32 = 0x37 << 16 //  mclk frequency = 1/56 * hmclk frequency
	DIVIDE_57                  uint32 = 0x38 << 16 //  mclk frequency = 1/57 * hmclk frequency
	DIVIDE_58                  uint32 = 0x39 << 16 //  mclk frequency = 1/58 * hmclk frequency
	DIVIDE_59                  uint32 = 0x3A << 16 //  mclk frequency = 1/59 * hmclk frequency
	DIVIDE_60                  uint32 = 0x3B << 16 //  mclk frequency = 1/60 * hmclk frequency
	DIVIDE_61                  uint32 = 0x3C << 16 //  mclk frequency = 1/61 * hmclk frequency
	DIVIDE_62                  uint32 = 0x3D << 16 //  mclk frequency = 1/62 * hmclk frequency
	DIVIDE_63                  uint32 = 0x3E << 16 //  mclk frequency = 1/63 * hmclk frequency
	DIVIDE_64                  uint32 = 0x3F << 16 //  mclk frequency = 1/64 * hmclk frequency
	DIVIDE_65                  uint32 = 0x40 << 16 //  mclk frequency = 1/65 * hmclk frequency
	DIVIDE_66                  uint32 = 0x41 << 16 //  mclk frequency = 1/66 * hmclk frequency
	DIVIDE_67                  uint32 = 0x42 << 16 //  mclk frequency = 1/67 * hmclk frequency
	DIVIDE_68                  uint32 = 0x43 << 16 //  mclk frequency = 1/68 * hmclk frequency
	DIVIDE_69                  uint32 = 0x44 << 16 //  mclk frequency = 1/69 * hmclk frequency
	DIVIDE_70                  uint32 = 0x45 << 16 //  mclk frequency = 1/70 * hmclk frequency
	DIVIDE_71                  uint32 = 0x46 << 16 //  mclk frequency = 1/71 * hmclk frequency
	DIVIDE_72                  uint32 = 0x47 << 16 //  mclk frequency = 1/72 * hmclk frequency
	DIVIDE_73                  uint32 = 0x48 << 16 //  mclk frequency = 1/73 * hmclk frequency
	DIVIDE_74                  uint32 = 0x49 << 16 //  mclk frequency = 1/74 * hmclk frequency
	DIVIDE_75                  uint32 = 0x4A << 16 //  mclk frequency = 1/75 * hmclk frequency
	DIVIDE_76                  uint32 = 0x4B << 16 //  mclk frequency = 1/76 * hmclk frequency
	DIVIDE_77                  uint32 = 0x4C << 16 //  mclk frequency = 1/77 * hmclk frequency
	DIVIDE_78                  uint32 = 0x4D << 16 //  mclk frequency = 1/78 * hmclk frequency
	DIVIDE_79                  uint32 = 0x4E << 16 //  mclk frequency = 1/79 * hmclk frequency
	DIVIDE_80                  uint32 = 0x4F << 16 //  mclk frequency = 1/80 * hmclk frequency
	DIVIDE_81                  uint32 = 0x50 << 16 //  mclk frequency = 1/81 * hmclk frequency
	DIVIDE_82                  uint32 = 0x51 << 16 //  mclk frequency = 1/82 * hmclk frequency
	DIVIDE_83                  uint32 = 0x52 << 16 //  mclk frequency = 1/83 * hmclk frequency
	DIVIDE_84                  uint32 = 0x53 << 16 //  mclk frequency = 1/84 * hmclk frequency
	DIVIDE_85                  uint32 = 0x54 << 16 //  mclk frequency = 1/85 * hmclk frequency
	DIVIDE_86                  uint32 = 0x55 << 16 //  mclk frequency = 1/86 * hmclk frequency
	DIVIDE_87                  uint32 = 0x56 << 16 //  mclk frequency = 1/87 * hmclk frequency
	DIVIDE_88                  uint32 = 0x57 << 16 //  mclk frequency = 1/88 * hmclk frequency
	DIVIDE_89                  uint32 = 0x58 << 16 //  mclk frequency = 1/89 * hmclk frequency
	DIVIDE_90                  uint32 = 0x59 << 16 //  mclk frequency = 1/90 * hmclk frequency
	DIVIDE_91                  uint32 = 0x5A << 16 //  mclk frequency = 1/91 * hmclk frequency
	DIVIDE_92                  uint32 = 0x5B << 16 //  mclk frequency = 1/92 * hmclk frequency
	DIVIDE_93                  uint32 = 0x5C << 16 //  mclk frequency = 1/93 * hmclk frequency
	DIVIDE_94                  uint32 = 0x5D << 16 //  mclk frequency = 1/94 * hmclk frequency
	DIVIDE_95                  uint32 = 0x5E << 16 //  mclk frequency = 1/95 * hmclk frequency
	DIVIDE_96                  uint32 = 0x5F << 16 //  mclk frequency = 1/96 * hmclk frequency
	DIVIDE_97                  uint32 = 0x60 << 16 //  mclk frequency = 1/97 * hmclk frequency
	DIVIDE_98                  uint32 = 0x61 << 16 //  mclk frequency = 1/98 * hmclk frequency
	DIVIDE_99                  uint32 = 0x62 << 16 //  mclk frequency = 1/99 * hmclk frequency
	DIVIDE_100                 uint32 = 0x63 << 16 //  mclk frequency = 1/100 * hmclk frequency
	DIVIDE_101                 uint32 = 0x64 << 16 //  mclk frequency = 1/101 * hmclk frequency
	DIVIDE_102                 uint32 = 0x65 << 16 //  mclk frequency = 1/102 * hmclk frequency
	DIVIDE_103                 uint32 = 0x66 << 16 //  mclk frequency = 1/103 * hmclk frequency
	DIVIDE_104                 uint32 = 0x67 << 16 //  mclk frequency = 1/104 * hmclk frequency
	DIVIDE_105                 uint32 = 0x68 << 16 //  mclk frequency = 1/105 * hmclk frequency
	DIVIDE_106                 uint32 = 0x69 << 16 //  mclk frequency = 1/106 * hmclk frequency
	DIVIDE_107                 uint32 = 0x6A << 16 //  mclk frequency = 1/107 * hmclk frequency
	DIVIDE_108                 uint32 = 0x6B << 16 //  mclk frequency = 1/108 * hmclk frequency
	DIVIDE_109                 uint32 = 0x6C << 16 //  mclk frequency = 1/109 * hmclk frequency
	DIVIDE_110                 uint32 = 0x6D << 16 //  mclk frequency = 1/110 * hmclk frequency
	DIVIDE_111                 uint32 = 0x6E << 16 //  mclk frequency = 1/111 * hmclk frequency
	DIVIDE_112                 uint32 = 0x6F << 16 //  mclk frequency = 1/112 * hmclk frequency
	DIVIDE_113                 uint32 = 0x70 << 16 //  mclk frequency = 1/113 * hmclk frequency
	DIVIDE_114                 uint32 = 0x71 << 16 //  mclk frequency = 1/114 * hmclk frequency
	DIVIDE_115                 uint32 = 0x72 << 16 //  mclk frequency = 1/115 * hmclk frequency
	DIVIDE_116                 uint32 = 0x73 << 16 //  mclk frequency = 1/116 * hmclk frequency
	DIVIDE_117                 uint32 = 0x74 << 16 //  mclk frequency = 1/117 * hmclk frequency
	DIVIDE_118                 uint32 = 0x75 << 16 //  mclk frequency = 1/118 * hmclk frequency
	DIVIDE_119                 uint32 = 0x76 << 16 //  mclk frequency = 1/119 * hmclk frequency
	DIVIDE_120                 uint32 = 0x77 << 16 //  mclk frequency = 1/120 * hmclk frequency
	DIVIDE_121                 uint32 = 0x78 << 16 //  mclk frequency = 1/121 * hmclk frequency
	DIVIDE_122                 uint32 = 0x79 << 16 //  mclk frequency = 1/122 * hmclk frequency
	DIVIDE_123                 uint32 = 0x7A << 16 //  mclk frequency = 1/123 * hmclk frequency
	DIVIDE_124                 uint32 = 0x7B << 16 //  mclk frequency = 1/124 * hmclk frequency
	DIVIDE_125                 uint32 = 0x7C << 16 //  mclk frequency = 1/125 * hmclk frequency
	DIVIDE_126                 uint32 = 0x7D << 16 //  mclk frequency = 1/126 * hmclk frequency
	DIVIDE_127                 uint32 = 0x7E << 16 //  mclk frequency = 1/127 * hmclk frequency
	DIVIDE_128                 uint32 = 0x7F << 16 //  mclk frequency = 1/128 * hmclk frequency
	DIVIDE_129                 uint32 = 0x80 << 16 //  mclk frequency = 1/129 * hmclk frequency
	DIVIDE_130                 uint32 = 0x81 << 16 //  mclk frequency = 1/130 * hmclk frequency
	DIVIDE_131                 uint32 = 0x82 << 16 //  mclk frequency = 1/131 * hmclk frequency
	DIVIDE_132                 uint32 = 0x83 << 16 //  mclk frequency = 1/132 * hmclk frequency
	DIVIDE_133                 uint32 = 0x84 << 16 //  mclk frequency = 1/133 * hmclk frequency
	DIVIDE_134                 uint32 = 0x85 << 16 //  mclk frequency = 1/134 * hmclk frequency
	DIVIDE_135                 uint32 = 0x86 << 16 //  mclk frequency = 1/135 * hmclk frequency
	DIVIDE_136                 uint32 = 0x87 << 16 //  mclk frequency = 1/136 * hmclk frequency
	DIVIDE_137                 uint32 = 0x88 << 16 //  mclk frequency = 1/137 * hmclk frequency
	DIVIDE_138                 uint32 = 0x89 << 16 //  mclk frequency = 1/138 * hmclk frequency
	DIVIDE_139                 uint32 = 0x8A << 16 //  mclk frequency = 1/139 * hmclk frequency
	DIVIDE_140                 uint32 = 0x8B << 16 //  mclk frequency = 1/140 * hmclk frequency
	DIVIDE_141                 uint32 = 0x8C << 16 //  mclk frequency = 1/141 * hmclk frequency
	DIVIDE_142                 uint32 = 0x8D << 16 //  mclk frequency = 1/142 * hmclk frequency
	DIVIDE_143                 uint32 = 0x8E << 16 //  mclk frequency = 1/143 * hmclk frequency
	DIVIDE_144                 uint32 = 0x8F << 16 //  mclk frequency = 1/144 * hmclk frequency
	DIVIDE_145                 uint32 = 0x90 << 16 //  mclk frequency = 1/145 * hmclk frequency
	DIVIDE_146                 uint32 = 0x91 << 16 //  mclk frequency = 1/146 * hmclk frequency
	DIVIDE_147                 uint32 = 0x92 << 16 //  mclk frequency = 1/147 * hmclk frequency
	DIVIDE_148                 uint32 = 0x93 << 16 //  mclk frequency = 1/148 * hmclk frequency
	DIVIDE_149                 uint32 = 0x94 << 16 //  mclk frequency = 1/149 * hmclk frequency
	DIVIDE_150                 uint32 = 0x95 << 16 //  mclk frequency = 1/150 * hmclk frequency
	DIVIDE_151                 uint32 = 0x96 << 16 //  mclk frequency = 1/151 * hmclk frequency
	DIVIDE_152                 uint32 = 0x97 << 16 //  mclk frequency = 1/152 * hmclk frequency
	DIVIDE_153                 uint32 = 0x98 << 16 //  mclk frequency = 1/153 * hmclk frequency
	DIVIDE_154                 uint32 = 0x99 << 16 //  mclk frequency = 1/154 * hmclk frequency
	DIVIDE_155                 uint32 = 0x9A << 16 //  mclk frequency = 1/155 * hmclk frequency
	DIVIDE_156                 uint32 = 0x9B << 16 //  mclk frequency = 1/156 * hmclk frequency
	DIVIDE_157                 uint32 = 0x9C << 16 //  mclk frequency = 1/157 * hmclk frequency
	DIVIDE_158                 uint32 = 0x9D << 16 //  mclk frequency = 1/158 * hmclk frequency
	DIVIDE_159                 uint32 = 0x9E << 16 //  mclk frequency = 1/159 * hmclk frequency
	DIVIDE_160                 uint32 = 0x9F << 16 //  mclk frequency = 1/160 * hmclk frequency
	DIVIDE_161                 uint32 = 0xA0 << 16 //  mclk frequency = 1/161 * hmclk frequency
	DIVIDE_162                 uint32 = 0xA1 << 16 //  mclk frequency = 1/162 * hmclk frequency
	DIVIDE_163                 uint32 = 0xA2 << 16 //  mclk frequency = 1/163 * hmclk frequency
	DIVIDE_164                 uint32 = 0xA3 << 16 //  mclk frequency = 1/164 * hmclk frequency
	DIVIDE_165                 uint32 = 0xA4 << 16 //  mclk frequency = 1/165 * hmclk frequency
	DIVIDE_166                 uint32 = 0xA5 << 16 //  mclk frequency = 1/166 * hmclk frequency
	DIVIDE_167                 uint32 = 0xA6 << 16 //  mclk frequency = 1/167 * hmclk frequency
	DIVIDE_168                 uint32 = 0xA7 << 16 //  mclk frequency = 1/168 * hmclk frequency
	DIVIDE_169                 uint32 = 0xA8 << 16 //  mclk frequency = 1/169 * hmclk frequency
	DIVIDE_170                 uint32 = 0xA9 << 16 //  mclk frequency = 1/170 * hmclk frequency
	DIVIDE_171                 uint32 = 0xAA << 16 //  mclk frequency = 1/171 * hmclk frequency
	DIVIDE_172                 uint32 = 0xAB << 16 //  mclk frequency = 1/172 * hmclk frequency
	DIVIDE_173                 uint32 = 0xAC << 16 //  mclk frequency = 1/173 * hmclk frequency
	DIVIDE_174                 uint32 = 0xAD << 16 //  mclk frequency = 1/174 * hmclk frequency
	DIVIDE_175                 uint32 = 0xAE << 16 //  mclk frequency = 1/175 * hmclk frequency
	DIVIDE_176                 uint32 = 0xAF << 16 //  mclk frequency = 1/176 * hmclk frequency
	DIVIDE_177                 uint32 = 0xB0 << 16 //  mclk frequency = 1/177 * hmclk frequency
	DIVIDE_178                 uint32 = 0xB1 << 16 //  mclk frequency = 1/178 * hmclk frequency
	DIVIDE_179                 uint32 = 0xB2 << 16 //  mclk frequency = 1/179 * hmclk frequency
	DIVIDE_180                 uint32 = 0xB3 << 16 //  mclk frequency = 1/180 * hmclk frequency
	DIVIDE_181                 uint32 = 0xB4 << 16 //  mclk frequency = 1/181 * hmclk frequency
	DIVIDE_182                 uint32 = 0xB5 << 16 //  mclk frequency = 1/182 * hmclk frequency
	DIVIDE_183                 uint32 = 0xB6 << 16 //  mclk frequency = 1/183 * hmclk frequency
	DIVIDE_184                 uint32 = 0xB7 << 16 //  mclk frequency = 1/184 * hmclk frequency
	DIVIDE_185                 uint32 = 0xB8 << 16 //  mclk frequency = 1/185 * hmclk frequency
	DIVIDE_186                 uint32 = 0xB9 << 16 //  mclk frequency = 1/186 * hmclk frequency
	DIVIDE_187                 uint32 = 0xBA << 16 //  mclk frequency = 1/187 * hmclk frequency
	DIVIDE_188                 uint32 = 0xBB << 16 //  mclk frequency = 1/188 * hmclk frequency
	DIVIDE_189                 uint32 = 0xBC << 16 //  mclk frequency = 1/189 * hmclk frequency
	DIVIDE_190                 uint32 = 0xBD << 16 //  mclk frequency = 1/190 * hmclk frequency
	DIVIDE_191                 uint32 = 0xBE << 16 //  mclk frequency = 1/191 * hmclk frequency
	DIVIDE_192                 uint32 = 0xBF << 16 //  mclk frequency = 1/192 * hmclk frequency
	DIVIDE_193                 uint32 = 0xC0 << 16 //  mclk frequency = 1/193 * hmclk frequency
	DIVIDE_194                 uint32 = 0xC1 << 16 //  mclk frequency = 1/194 * hmclk frequency
	DIVIDE_195                 uint32 = 0xC2 << 16 //  mclk frequency = 1/195 * hmclk frequency
	DIVIDE_196                 uint32 = 0xC3 << 16 //  mclk frequency = 1/196 * hmclk frequency
	DIVIDE_197                 uint32 = 0xC4 << 16 //  mclk frequency = 1/197 * hmclk frequency
	DIVIDE_198                 uint32 = 0xC5 << 16 //  mclk frequency = 1/198 * hmclk frequency
	DIVIDE_199                 uint32 = 0xC6 << 16 //  mclk frequency = 1/199 * hmclk frequency
	DIVIDE_200                 uint32 = 0xC7 << 16 //  mclk frequency = 1/200 * hmclk frequency
	DIVIDE_201                 uint32 = 0xC8 << 16 //  mclk frequency = 1/201 * hmclk frequency
	DIVIDE_202                 uint32 = 0xC9 << 16 //  mclk frequency = 1/202 * hmclk frequency
	DIVIDE_203                 uint32 = 0xCA << 16 //  mclk frequency = 1/203 * hmclk frequency
	DIVIDE_204                 uint32 = 0xCB << 16 //  mclk frequency = 1/204 * hmclk frequency
	DIVIDE_205                 uint32 = 0xCC << 16 //  mclk frequency = 1/205 * hmclk frequency
	DIVIDE_206                 uint32 = 0xCD << 16 //  mclk frequency = 1/206 * hmclk frequency
	DIVIDE_207                 uint32 = 0xCE << 16 //  mclk frequency = 1/207 * hmclk frequency
	DIVIDE_208                 uint32 = 0xCF << 16 //  mclk frequency = 1/208 * hmclk frequency
	DIVIDE_209                 uint32 = 0xD0 << 16 //  mclk frequency = 1/209 * hmclk frequency
	DIVIDE_210                 uint32 = 0xD1 << 16 //  mclk frequency = 1/210 * hmclk frequency
	DIVIDE_211                 uint32 = 0xD2 << 16 //  mclk frequency = 1/211 * hmclk frequency
	DIVIDE_212                 uint32 = 0xD3 << 16 //  mclk frequency = 1/212 * hmclk frequency
	DIVIDE_213                 uint32 = 0xD4 << 16 //  mclk frequency = 1/213 * hmclk frequency
	DIVIDE_214                 uint32 = 0xD5 << 16 //  mclk frequency = 1/214 * hmclk frequency
	DIVIDE_215                 uint32 = 0xD6 << 16 //  mclk frequency = 1/215 * hmclk frequency
	DIVIDE_216                 uint32 = 0xD7 << 16 //  mclk frequency = 1/216 * hmclk frequency
	DIVIDE_217                 uint32 = 0xD8 << 16 //  mclk frequency = 1/217 * hmclk frequency
	DIVIDE_218                 uint32 = 0xD9 << 16 //  mclk frequency = 1/218 * hmclk frequency
	DIVIDE_219                 uint32 = 0xDA << 16 //  mclk frequency = 1/219 * hmclk frequency
	DIVIDE_220                 uint32 = 0xDB << 16 //  mclk frequency = 1/220 * hmclk frequency
	DIVIDE_221                 uint32 = 0xDC << 16 //  mclk frequency = 1/221 * hmclk frequency
	DIVIDE_222                 uint32 = 0xDD << 16 //  mclk frequency = 1/222 * hmclk frequency
	DIVIDE_223                 uint32 = 0xDE << 16 //  mclk frequency = 1/223 * hmclk frequency
	DIVIDE_224                 uint32 = 0xDF << 16 //  mclk frequency = 1/224 * hmclk frequency
	DIVIDE_225                 uint32 = 0xE0 << 16 //  mclk frequency = 1/225 * hmclk frequency
	DIVIDE_226                 uint32 = 0xE1 << 16 //  mclk frequency = 1/226 * hmclk frequency
	DIVIDE_227                 uint32 = 0xE2 << 16 //  mclk frequency = 1/227 * hmclk frequency
	DIVIDE_228                 uint32 = 0xE3 << 16 //  mclk frequency = 1/228 * hmclk frequency
	DIVIDE_229                 uint32 = 0xE4 << 16 //  mclk frequency = 1/229 * hmclk frequency
	DIVIDE_230                 uint32 = 0xE5 << 16 //  mclk frequency = 1/230 * hmclk frequency
	DIVIDE_231                 uint32 = 0xE6 << 16 //  mclk frequency = 1/231 * hmclk frequency
	DIVIDE_232                 uint32 = 0xE7 << 16 //  mclk frequency = 1/232 * hmclk frequency
	DIVIDE_233                 uint32 = 0xE8 << 16 //  mclk frequency = 1/233 * hmclk frequency
	DIVIDE_234                 uint32 = 0xE9 << 16 //  mclk frequency = 1/234 * hmclk frequency
	DIVIDE_235                 uint32 = 0xEA << 16 //  mclk frequency = 1/235 * hmclk frequency
	DIVIDE_236                 uint32 = 0xEB << 16 //  mclk frequency = 1/236 * hmclk frequency
	DIVIDE_237                 uint32 = 0xEC << 16 //  mclk frequency = 1/237 * hmclk frequency
	DIVIDE_238                 uint32 = 0xED << 16 //  mclk frequency = 1/238 * hmclk frequency
	DIVIDE_239                 uint32 = 0xEE << 16 //  mclk frequency = 1/239 * hmclk frequency
	DIVIDE_240                 uint32 = 0xEF << 16 //  mclk frequency = 1/240 * hmclk frequency
	DIVIDE_241                 uint32 = 0xF0 << 16 //  mclk frequency = 1/241 * hmclk frequency
	DIVIDE_242                 uint32 = 0xF1 << 16 //  mclk frequency = 1/242 * hmclk frequency
	DIVIDE_243                 uint32 = 0xF2 << 16 //  mclk frequency = 1/243 * hmclk frequency
	DIVIDE_244                 uint32 = 0xF3 << 16 //  mclk frequency = 1/244 * hmclk frequency
	DIVIDE_245                 uint32 = 0xF4 << 16 //  mclk frequency = 1/245 * hmclk frequency
	DIVIDE_246                 uint32 = 0xF5 << 16 //  mclk frequency = 1/246 * hmclk frequency
	DIVIDE_247                 uint32 = 0xF6 << 16 //  mclk frequency = 1/247 * hmclk frequency
	DIVIDE_248                 uint32 = 0xF7 << 16 //  mclk frequency = 1/248 * hmclk frequency
	DIVIDE_249                 uint32 = 0xF8 << 16 //  mclk frequency = 1/249 * hmclk frequency
	DIVIDE_250                 uint32 = 0xF9 << 16 //  mclk frequency = 1/250 * hmclk frequency
	DIVIDE_251                 uint32 = 0xFA << 16 //  mclk frequency = 1/251 * hmclk frequency
	DIVIDE_252                 uint32 = 0xFB << 16 //  mclk frequency = 1/252 * hmclk frequency
	DIVIDE_253                 uint32 = 0xFC << 16 //  mclk frequency = 1/253 * hmclk frequency
	DIVIDE_254                 uint32 = 0xFD << 16 //  mclk frequency = 1/254 * hmclk frequency
	DIVIDE_255                 uint32 = 0xFE << 16 //  mclk frequency = 1/255 * hmclk frequency
	DIVIDE_256                 uint32 = 0xFF << 16 //  mclk frequency = 1/256 * hmclk frequency
	MQS_SW_RST                 uint32 = 0x01 << 24 //+ MQS software reset
	MQS_EN                     uint32 = 0x01 << 25 //+ MQS enable.
	MQS_OVERSAMPLE             uint32 = 0x01 << 26 //+ Used to control the PWM oversampling rate compared with mclk.
	QTIMER1_TMR_CNTS_FREEZE    uint32 = 0x01 << 28 //+ QTIMER1 timer counter freeze
	QTIMER2_TMR_CNTS_FREEZE    uint32 = 0x01 << 29 //+ QTIMER2 timer counter freeze
	QTIMER3_TMR_CNTS_FREEZE    uint32 = 0x01 << 30 //+ QTIMER3 timer counter freeze
	QTIMER4_TMR_CNTS_FREEZE    uint32 = 0x01 << 31 //+ QTIMER4 timer counter freeze
)

const (
	AXBS_L_AHBXL_HIGH_PRIORITYn = 0
	AXBS_L_DMA_HIGH_PRIORITYn   = 1
	AXBS_L_FORCE_ROUND_ROBINn   = 2
	AXBS_P_M0_HIGH_PRIORITYn    = 3
	AXBS_P_M1_HIGH_PRIORITYn    = 4
	AXBS_P_FORCE_ROUND_ROBINn   = 5
	CANFD_FILTER_BYPASSn        = 6
	L2_MEM_EN_POWERSAVINGn      = 12
	RAM_AUTO_CLK_GATING_ENn     = 13
	L2_MEM_DEEPSLEEPn           = 14
	MQS_CLK_DIVn                = 16
	MQS_SW_RSTn                 = 24
	MQS_ENn                     = 25
	MQS_OVERSAMPLEn             = 26
	QTIMER1_TMR_CNTS_FREEZEn    = 28
	QTIMER2_TMR_CNTS_FREEZEn    = 29
	QTIMER3_TMR_CNTS_FREEZEn    = 30
	QTIMER4_TMR_CNTS_FREEZEn    = 31
)

// GPR3 bits
const (
	OCRAM_CTL       uint32 = 0x0F << 0  //+ OCRAM_CTL[3] - write address pipeline control bit
	DCP_KEY_SEL     uint32 = 0x01 << 4  //+ Select 128-bit dcp key from 256-bit key from snvs/ocotp
	OCRAM2_CTL      uint32 = 0x0F << 8  //+ OCRAM2_CTL[3] - write address pipeline control bit
	AXBS_L_HALT_REQ uint32 = 0x01 << 15 //+ Request to halt axbs_l
	OCRAM_STATUS    uint32 = 0x0F << 16 //+ This field shows the OCRAM pipeline settings status, controlled by OCRAM_CTL bits respectively
	OCRAM2_STATUS   uint32 = 0x0F << 24 //+ This field shows the OCRAM2 pipeline settings status, controlled by OCRAM2_CTL bits respectively
	AXBS_L_HALTED   uint32 = 0x01 << 31 //+ This bit shows the status of axbs_l
)

const (
	OCRAM_CTLn       = 0
	DCP_KEY_SELn     = 4
	OCRAM2_CTLn      = 8
	AXBS_L_HALT_REQn = 15
	OCRAM_STATUSn    = 16
	OCRAM2_STATUSn   = 24
	AXBS_L_HALTEDn   = 31
)

// GPR4 bits
const (
	EDMA_STOP_REQ     uint32 = 0x01 << 0  //+ EDMA stop request.
	CAN1_STOP_REQ     uint32 = 0x01 << 1  //+ CAN1 stop request.
	CAN2_STOP_REQ     uint32 = 0x01 << 2  //+ CAN2 stop request.
	TRNG_STOP_REQ     uint32 = 0x01 << 3  //+ TRNG stop request.
	ENET_STOP_REQ     uint32 = 0x01 << 4  //+ ENET stop request.
	SAI1_STOP_REQ     uint32 = 0x01 << 5  //+ SAI1 stop request.
	SAI2_STOP_REQ     uint32 = 0x01 << 6  //+ SAI2 stop request.
	SAI3_STOP_REQ     uint32 = 0x01 << 7  //+ SAI3 stop request.
	ENET2_STOP_REQ    uint32 = 0x01 << 8  //+ ENET2 stop request.
	SEMC_STOP_REQ     uint32 = 0x01 << 9  //+ SEMC stop request.
	PIT_STOP_REQ      uint32 = 0x01 << 10 //+ PIT stop request.
	FLEXSPI_STOP_REQ  uint32 = 0x01 << 11 //+ FlexSPI stop request.
	FLEXIO1_STOP_REQ  uint32 = 0x01 << 12 //+ FlexIO1 stop request.
	FLEXIO2_STOP_REQ  uint32 = 0x01 << 13 //+ FlexIO2 stop request.
	FLEXIO3_STOP_REQ  uint32 = 0x01 << 14 //+ On-platform flexio3 stop request.
	FLEXSPI2_STOP_REQ uint32 = 0x01 << 15 //+ FlexSPI2 stop request.
	EDMA_STOP_ACK     uint32 = 0x01 << 16 //+ EDMA stop acknowledge. This is a status (read-only) bit
	CAN1_STOP_ACK     uint32 = 0x01 << 17 //+ CAN1 stop acknowledge.
	CAN2_STOP_ACK     uint32 = 0x01 << 18 //+ CAN2 stop acknowledge.
	TRNG_STOP_ACK     uint32 = 0x01 << 19 //+ TRNG stop acknowledge
	ENET_STOP_ACK     uint32 = 0x01 << 20 //+ ENET stop acknowledge.
	SAI1_STOP_ACK     uint32 = 0x01 << 21 //+ SAI1 stop acknowledge
	SAI2_STOP_ACK     uint32 = 0x01 << 22 //+ SAI2 stop acknowledge
	SAI3_STOP_ACK     uint32 = 0x01 << 23 //+ SAI3 stop acknowledge
	ENET2_STOP_ACK    uint32 = 0x01 << 24 //+ ENET2 stop acknowledge.
	SEMC_STOP_ACK     uint32 = 0x01 << 25 //+ SEMC stop acknowledge
	PIT_STOP_ACK      uint32 = 0x01 << 26 //+ PIT stop acknowledge
	FLEXSPI_STOP_ACK  uint32 = 0x01 << 27 //+ FLEXSPI stop acknowledge
	FLEXIO1_STOP_ACK  uint32 = 0x01 << 28 //+ FLEXIO1 stop acknowledge
	FLEXIO2_STOP_ACK  uint32 = 0x01 << 29 //+ FLEXIO2 stop acknowledge
	FLEXIO3_STOP_ACK  uint32 = 0x01 << 30 //+ On-platform FLEXIO3 stop acknowledge
	FLEXSPI2_STOP_ACK uint32 = 0x01 << 31 //+ FLEXSPI2 stop acknowledge
)

const (
	EDMA_STOP_REQn     = 0
	CAN1_STOP_REQn     = 1
	CAN2_STOP_REQn     = 2
	TRNG_STOP_REQn     = 3
	ENET_STOP_REQn     = 4
	SAI1_STOP_REQn     = 5
	SAI2_STOP_REQn     = 6
	SAI3_STOP_REQn     = 7
	ENET2_STOP_REQn    = 8
	SEMC_STOP_REQn     = 9
	PIT_STOP_REQn      = 10
	FLEXSPI_STOP_REQn  = 11
	FLEXIO1_STOP_REQn  = 12
	FLEXIO2_STOP_REQn  = 13
	FLEXIO3_STOP_REQn  = 14
	FLEXSPI2_STOP_REQn = 15
	EDMA_STOP_ACKn     = 16
	CAN1_STOP_ACKn     = 17
	CAN2_STOP_ACKn     = 18
	TRNG_STOP_ACKn     = 19
	ENET_STOP_ACKn     = 20
	SAI1_STOP_ACKn     = 21
	SAI2_STOP_ACKn     = 22
	SAI3_STOP_ACKn     = 23
	ENET2_STOP_ACKn    = 24
	SEMC_STOP_ACKn     = 25
	PIT_STOP_ACKn      = 26
	FLEXSPI_STOP_ACKn  = 27
	FLEXIO1_STOP_ACKn  = 28
	FLEXIO2_STOP_ACKn  = 29
	FLEXIO3_STOP_ACKn  = 30
	FLEXSPI2_STOP_ACKn = 31
)

// GPR5 bits
const (
	WDOG1_MASK         uint32 = 0x01 << 6  //+ WDOG1 Timeout Mask
	WDOG2_MASK         uint32 = 0x01 << 7  //+ WDOG2 Timeout Mask
	GPT2_CAPIN1_SEL    uint32 = 0x01 << 23 //+ GPT2 input capture channel 1 source select
	GPT2_CAPIN2_SEL    uint32 = 0x01 << 24 //+ GPT2 input capture channel 2 source select
	ENET_EVENT3IN_SEL  uint32 = 0x01 << 25 //+ ENET input timer event3 source select
	ENET2_EVENT3IN_SEL uint32 = 0x01 << 26 //+ ENET2 input timer event3 source select
	VREF_1M_CLK_GPT1   uint32 = 0x01 << 28 //+ GPT1 1 MHz clock source select
	VREF_1M_CLK_GPT2   uint32 = 0x01 << 29 //+ GPT2 1 MHz clock source select
)

const (
	WDOG1_MASKn         = 6
	WDOG2_MASKn         = 7
	GPT2_CAPIN1_SELn    = 23
	GPT2_CAPIN2_SELn    = 24
	ENET_EVENT3IN_SELn  = 25
	ENET2_EVENT3IN_SELn = 26
	VREF_1M_CLK_GPT1n   = 28
	VREF_1M_CLK_GPT2n   = 29
)

// GPR6 bits
const (
	QTIMER1_TRM0_INPUT_SEL uint32 = 0x01 << 0  //+ QTIMER1 TMR0 input select
	QTIMER1_TRM1_INPUT_SEL uint32 = 0x01 << 1  //+ QTIMER1 TMR1 input select
	QTIMER1_TRM2_INPUT_SEL uint32 = 0x01 << 2  //+ QTIMER1 TMR2 input select
	QTIMER1_TRM3_INPUT_SEL uint32 = 0x01 << 3  //+ QTIMER1 TMR3 input select
	QTIMER2_TRM0_INPUT_SEL uint32 = 0x01 << 4  //+ QTIMER2 TMR0 input select
	QTIMER2_TRM1_INPUT_SEL uint32 = 0x01 << 5  //+ QTIMER2 TMR1 input select
	QTIMER2_TRM2_INPUT_SEL uint32 = 0x01 << 6  //+ QTIMER2 TMR2 input select
	QTIMER2_TRM3_INPUT_SEL uint32 = 0x01 << 7  //+ QTIMER2 TMR3 input select
	QTIMER3_TRM0_INPUT_SEL uint32 = 0x01 << 8  //+ QTIMER3 TMR0 input select
	QTIMER3_TRM1_INPUT_SEL uint32 = 0x01 << 9  //+ QTIMER3 TMR1 input select
	QTIMER3_TRM2_INPUT_SEL uint32 = 0x01 << 10 //+ QTIMER3 TMR2 input select
	QTIMER3_TRM3_INPUT_SEL uint32 = 0x01 << 11 //+ QTIMER3 TMR3 input select
	QTIMER4_TRM0_INPUT_SEL uint32 = 0x01 << 12 //+ QTIMER4 TMR0 input select
	QTIMER4_TRM1_INPUT_SEL uint32 = 0x01 << 13 //+ QTIMER4 TMR1 input select
	QTIMER4_TRM2_INPUT_SEL uint32 = 0x01 << 14 //+ QTIMER4 TMR2 input select
	QTIMER4_TRM3_INPUT_SEL uint32 = 0x01 << 15 //+ QTIMER4 TMR3 input select
	IOMUXC_XBAR_DIR_SEL_4  uint32 = 0x01 << 16 //+ IOMUXC XBAR_INOUT4 function direction select
	IOMUXC_XBAR_DIR_SEL_5  uint32 = 0x01 << 17 //+ IOMUXC XBAR_INOUT5 function direction select
	IOMUXC_XBAR_DIR_SEL_6  uint32 = 0x01 << 18 //+ IOMUXC XBAR_INOUT6 function direction select
	IOMUXC_XBAR_DIR_SEL_7  uint32 = 0x01 << 19 //+ IOMUXC XBAR_INOUT7 function direction select
	IOMUXC_XBAR_DIR_SEL_8  uint32 = 0x01 << 20 //+ IOMUXC XBAR_INOUT8 function direction select
	IOMUXC_XBAR_DIR_SEL_9  uint32 = 0x01 << 21 //+ IOMUXC XBAR_INOUT9 function direction select
	IOMUXC_XBAR_DIR_SEL_10 uint32 = 0x01 << 22 //+ IOMUXC XBAR_INOUT10 function direction select
	IOMUXC_XBAR_DIR_SEL_11 uint32 = 0x01 << 23 //+ IOMUXC XBAR_INOUT11 function direction select
	IOMUXC_XBAR_DIR_SEL_12 uint32 = 0x01 << 24 //+ IOMUXC XBAR_INOUT12 function direction select
	IOMUXC_XBAR_DIR_SEL_13 uint32 = 0x01 << 25 //+ IOMUXC XBAR_INOUT13 function direction select
	IOMUXC_XBAR_DIR_SEL_14 uint32 = 0x01 << 26 //+ IOMUXC XBAR_INOUT14 function direction select
	IOMUXC_XBAR_DIR_SEL_15 uint32 = 0x01 << 27 //+ IOMUXC XBAR_INOUT15 function direction select
	IOMUXC_XBAR_DIR_SEL_16 uint32 = 0x01 << 28 //+ IOMUXC XBAR_INOUT16 function direction select
	IOMUXC_XBAR_DIR_SEL_17 uint32 = 0x01 << 29 //+ IOMUXC XBAR_INOUT17 function direction select
	IOMUXC_XBAR_DIR_SEL_18 uint32 = 0x01 << 30 //+ IOMUXC XBAR_INOUT18 function direction select
	IOMUXC_XBAR_DIR_SEL_19 uint32 = 0x01 << 31 //+ IOMUXC XBAR_INOUT19 function direction select
)

const (
	QTIMER1_TRM0_INPUT_SELn = 0
	QTIMER1_TRM1_INPUT_SELn = 1
	QTIMER1_TRM2_INPUT_SELn = 2
	QTIMER1_TRM3_INPUT_SELn = 3
	QTIMER2_TRM0_INPUT_SELn = 4
	QTIMER2_TRM1_INPUT_SELn = 5
	QTIMER2_TRM2_INPUT_SELn = 6
	QTIMER2_TRM3_INPUT_SELn = 7
	QTIMER3_TRM0_INPUT_SELn = 8
	QTIMER3_TRM1_INPUT_SELn = 9
	QTIMER3_TRM2_INPUT_SELn = 10
	QTIMER3_TRM3_INPUT_SELn = 11
	QTIMER4_TRM0_INPUT_SELn = 12
	QTIMER4_TRM1_INPUT_SELn = 13
	QTIMER4_TRM2_INPUT_SELn = 14
	QTIMER4_TRM3_INPUT_SELn = 15
	IOMUXC_XBAR_DIR_SEL_4n  = 16
	IOMUXC_XBAR_DIR_SEL_5n  = 17
	IOMUXC_XBAR_DIR_SEL_6n  = 18
	IOMUXC_XBAR_DIR_SEL_7n  = 19
	IOMUXC_XBAR_DIR_SEL_8n  = 20
	IOMUXC_XBAR_DIR_SEL_9n  = 21
	IOMUXC_XBAR_DIR_SEL_10n = 22
	IOMUXC_XBAR_DIR_SEL_11n = 23
	IOMUXC_XBAR_DIR_SEL_12n = 24
	IOMUXC_XBAR_DIR_SEL_13n = 25
	IOMUXC_XBAR_DIR_SEL_14n = 26
	IOMUXC_XBAR_DIR_SEL_15n = 27
	IOMUXC_XBAR_DIR_SEL_16n = 28
	IOMUXC_XBAR_DIR_SEL_17n = 29
	IOMUXC_XBAR_DIR_SEL_18n = 30
	IOMUXC_XBAR_DIR_SEL_19n = 31
)

// GPR7 bits
const (
	LPI2C1_STOP_REQ  uint32 = 0x01 << 0  //+ LPI2C1 stop request
	LPI2C2_STOP_REQ  uint32 = 0x01 << 1  //+ LPI2C2 stop request
	LPI2C3_STOP_REQ  uint32 = 0x01 << 2  //+ LPI2C3 stop request
	LPI2C4_STOP_REQ  uint32 = 0x01 << 3  //+ LPI2C4 stop request
	LPSPI1_STOP_REQ  uint32 = 0x01 << 4  //+ LPSPI1 stop request
	LPSPI2_STOP_REQ  uint32 = 0x01 << 5  //+ LPSPI2 stop request
	LPSPI3_STOP_REQ  uint32 = 0x01 << 6  //+ LPSPI3 stop request
	LPSPI4_STOP_REQ  uint32 = 0x01 << 7  //+ LPSPI4 stop request
	LPUART1_STOP_REQ uint32 = 0x01 << 8  //+ LPUART1 stop request
	LPUART2_STOP_REQ uint32 = 0x01 << 9  //+ LPUART1 stop request
	LPUART3_STOP_REQ uint32 = 0x01 << 10 //+ LPUART3 stop request
	LPUART4_STOP_REQ uint32 = 0x01 << 11 //+ LPUART4 stop request
	LPUART5_STOP_REQ uint32 = 0x01 << 12 //+ LPUART5 stop request
	LPUART6_STOP_REQ uint32 = 0x01 << 13 //+ LPUART6 stop request
	LPUART7_STOP_REQ uint32 = 0x01 << 14 //+ LPUART7 stop request
	LPUART8_STOP_REQ uint32 = 0x01 << 15 //+ LPUART8 stop request
	LPI2C1_STOP_ACK  uint32 = 0x01 << 16 //+ LPI2C1 stop acknowledge
	LPI2C2_STOP_ACK  uint32 = 0x01 << 17 //+ LPI2C2 stop acknowledge
	LPI2C3_STOP_ACK  uint32 = 0x01 << 18 //+ LPI2C3 stop acknowledge
	LPI2C4_STOP_ACK  uint32 = 0x01 << 19 //+ LPI2C4 stop acknowledge
	LPSPI1_STOP_ACK  uint32 = 0x01 << 20 //+ LPSPI1 stop acknowledge
	LPSPI2_STOP_ACK  uint32 = 0x01 << 21 //+ LPSPI2 stop acknowledge
	LPSPI3_STOP_ACK  uint32 = 0x01 << 22 //+ LPSPI3 stop acknowledge
	LPSPI4_STOP_ACK  uint32 = 0x01 << 23 //+ LPSPI4 stop acknowledge
	LPUART1_STOP_ACK uint32 = 0x01 << 24 //+ LPUART1 stop acknowledge
	LPUART2_STOP_ACK uint32 = 0x01 << 25 //+ LPUART1 stop acknowledge
	LPUART3_STOP_ACK uint32 = 0x01 << 26 //+ LPUART3 stop acknowledge
	LPUART4_STOP_ACK uint32 = 0x01 << 27 //+ LPUART4 stop acknowledge
	LPUART5_STOP_ACK uint32 = 0x01 << 28 //+ LPUART5 stop acknowledge
	LPUART6_STOP_ACK uint32 = 0x01 << 29 //+ LPUART6 stop acknowledge
	LPUART7_STOP_ACK uint32 = 0x01 << 30 //+ LPUART7 stop acknowledge
	LPUART8_STOP_ACK uint32 = 0x01 << 31 //+ LPUART8 stop acknowledge
)

const (
	LPI2C1_STOP_REQn  = 0
	LPI2C2_STOP_REQn  = 1
	LPI2C3_STOP_REQn  = 2
	LPI2C4_STOP_REQn  = 3
	LPSPI1_STOP_REQn  = 4
	LPSPI2_STOP_REQn  = 5
	LPSPI3_STOP_REQn  = 6
	LPSPI4_STOP_REQn  = 7
	LPUART1_STOP_REQn = 8
	LPUART2_STOP_REQn = 9
	LPUART3_STOP_REQn = 10
	LPUART4_STOP_REQn = 11
	LPUART5_STOP_REQn = 12
	LPUART6_STOP_REQn = 13
	LPUART7_STOP_REQn = 14
	LPUART8_STOP_REQn = 15
	LPI2C1_STOP_ACKn  = 16
	LPI2C2_STOP_ACKn  = 17
	LPI2C3_STOP_ACKn  = 18
	LPI2C4_STOP_ACKn  = 19
	LPSPI1_STOP_ACKn  = 20
	LPSPI2_STOP_ACKn  = 21
	LPSPI3_STOP_ACKn  = 22
	LPSPI4_STOP_ACKn  = 23
	LPUART1_STOP_ACKn = 24
	LPUART2_STOP_ACKn = 25
	LPUART3_STOP_ACKn = 26
	LPUART4_STOP_ACKn = 27
	LPUART5_STOP_ACKn = 28
	LPUART6_STOP_ACKn = 29
	LPUART7_STOP_ACKn = 30
	LPUART8_STOP_ACKn = 31
)

// GPR8 bits
const (
	LPI2C1_IPG_STOP_MODE  uint32 = 0x01 << 0  //+ LPI2C1 stop mode selection, cannot change when ipg_stop is asserted.
	LPI2C1_IPG_DOZE       uint32 = 0x01 << 1  //+ LPI2C1 ipg_doze mode
	LPI2C2_IPG_STOP_MODE  uint32 = 0x01 << 2  //+ LPI2C2 stop mode selection, cannot change when ipg_stop is asserted.
	LPI2C2_IPG_DOZE       uint32 = 0x01 << 3  //+ LPI2C2 ipg_doze mode
	LPI2C3_IPG_STOP_MODE  uint32 = 0x01 << 4  //+ LPI2C3 stop mode selection, cannot change when ipg_stop is asserted.
	LPI2C3_IPG_DOZE       uint32 = 0x01 << 5  //+ LPI2C3 ipg_doze mode
	LPI2C4_IPG_STOP_MODE  uint32 = 0x01 << 6  //+ LPI2C4 stop mode selection, cannot change when ipg_stop is asserted.
	LPI2C4_IPG_DOZE       uint32 = 0x01 << 7  //+ LPI2C4 ipg_doze mode
	LPSPI1_IPG_STOP_MODE  uint32 = 0x01 << 8  //+ LPSPI1 stop mode selection, cannot change when ipg_stop is asserted.
	LPSPI1_IPG_DOZE       uint32 = 0x01 << 9  //+ LPSPI1 ipg_doze mode
	LPSPI2_IPG_STOP_MODE  uint32 = 0x01 << 10 //+ LPSPI2 stop mode selection, cannot change when ipg_stop is asserted.
	LPSPI2_IPG_DOZE       uint32 = 0x01 << 11 //+ LPSPI2 ipg_doze mode
	LPSPI3_IPG_STOP_MODE  uint32 = 0x01 << 12 //+ LPSPI3 stop mode selection, cannot change when ipg_stop is asserted.
	LPSPI3_IPG_DOZE       uint32 = 0x01 << 13 //+ LPSPI3 ipg_doze mode
	LPSPI4_IPG_STOP_MODE  uint32 = 0x01 << 14 //+ LPSPI4 stop mode selection, cannot change when ipg_stop is asserted.
	LPSPI4_IPG_DOZE       uint32 = 0x01 << 15 //+ LPSPI4 ipg_doze mode
	LPUART1_IPG_STOP_MODE uint32 = 0x01 << 16 //+ LPUART1 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART1_IPG_DOZE      uint32 = 0x01 << 17 //+ LPUART1 ipg_doze mode
	LPUART2_IPG_STOP_MODE uint32 = 0x01 << 18 //+ LPUART2 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART2_IPG_DOZE      uint32 = 0x01 << 19 //+ LPUART2 ipg_doze mode
	LPUART3_IPG_STOP_MODE uint32 = 0x01 << 20 //+ LPUART3 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART3_IPG_DOZE      uint32 = 0x01 << 21 //+ LPUART3 ipg_doze mode
	LPUART4_IPG_STOP_MODE uint32 = 0x01 << 22 //+ LPUART4 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART4_IPG_DOZE      uint32 = 0x01 << 23 //+ LPUART4 ipg_doze mode
	LPUART5_IPG_STOP_MODE uint32 = 0x01 << 24 //+ LPUART5 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART5_IPG_DOZE      uint32 = 0x01 << 25 //+ LPUART5 ipg_doze mode
	LPUART6_IPG_STOP_MODE uint32 = 0x01 << 26 //+ LPUART6 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART6_IPG_DOZE      uint32 = 0x01 << 27 //+ LPUART6 ipg_doze mode
	LPUART7_IPG_STOP_MODE uint32 = 0x01 << 28 //+ LPUART7 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART7_IPG_DOZE      uint32 = 0x01 << 29 //+ LPUART7 ipg_doze mode
	LPUART8_IPG_STOP_MODE uint32 = 0x01 << 30 //+ LPUART8 stop mode selection, cannot change when ipg_stop is asserted.
	LPUART8_IPG_DOZE      uint32 = 0x01 << 31 //+ LPUART8 ipg_doze mode
)

const (
	LPI2C1_IPG_STOP_MODEn  = 0
	LPI2C1_IPG_DOZEn       = 1
	LPI2C2_IPG_STOP_MODEn  = 2
	LPI2C2_IPG_DOZEn       = 3
	LPI2C3_IPG_STOP_MODEn  = 4
	LPI2C3_IPG_DOZEn       = 5
	LPI2C4_IPG_STOP_MODEn  = 6
	LPI2C4_IPG_DOZEn       = 7
	LPSPI1_IPG_STOP_MODEn  = 8
	LPSPI1_IPG_DOZEn       = 9
	LPSPI2_IPG_STOP_MODEn  = 10
	LPSPI2_IPG_DOZEn       = 11
	LPSPI3_IPG_STOP_MODEn  = 12
	LPSPI3_IPG_DOZEn       = 13
	LPSPI4_IPG_STOP_MODEn  = 14
	LPSPI4_IPG_DOZEn       = 15
	LPUART1_IPG_STOP_MODEn = 16
	LPUART1_IPG_DOZEn      = 17
	LPUART2_IPG_STOP_MODEn = 18
	LPUART2_IPG_DOZEn      = 19
	LPUART3_IPG_STOP_MODEn = 20
	LPUART3_IPG_DOZEn      = 21
	LPUART4_IPG_STOP_MODEn = 22
	LPUART4_IPG_DOZEn      = 23
	LPUART5_IPG_STOP_MODEn = 24
	LPUART5_IPG_DOZEn      = 25
	LPUART6_IPG_STOP_MODEn = 26
	LPUART6_IPG_DOZEn      = 27
	LPUART7_IPG_STOP_MODEn = 28
	LPUART7_IPG_DOZEn      = 29
	LPUART8_IPG_STOP_MODEn = 30
	LPUART8_IPG_DOZEn      = 31
)

// GPR10 bits
const (
	NIDEN                       uint32 = 0x01 << 0  //+ ARM non-secure (non-invasive) debug enable
	DBG_EN                      uint32 = 0x01 << 1  //+ ARM invasive debug enable
	SEC_ERR_RESP                uint32 = 0x01 << 2  //+ Security error response enable for all security gaskets (on both AHB and AXI buses)
	DCPKEY_OCOTP_OR_KEYMUX      uint32 = 0x01 << 4  //+ DCP Key selection bit.
	OCRAM_TZ_EN                 uint32 = 0x01 << 8  //+ OCRAM TrustZone (TZ) enable.
	OCRAM_TZ_ADDR               uint32 = 0x7F << 9  //+ OCRAM TrustZone (TZ) start address
	LOCK_NIDEN                  uint32 = 0x01 << 16 //+ Lock NIDEN field for changes
	LOCK_DBG_EN                 uint32 = 0x01 << 17 //+ Lock DBG_EN field for changes
	LOCK_SEC_ERR_RESP           uint32 = 0x01 << 18 //+ Lock SEC_ERR_RESP field for changes
	LOCK_DCPKEY_OCOTP_OR_KEYMUX uint32 = 0x01 << 20 //+ Lock DCP Key OCOTP/Key MUX selection bit
	LOCK_OCRAM_TZ_EN            uint32 = 0x01 << 24 //+ Lock OCRAM_TZ_EN field for changes
	LOCK_OCRAM_TZ_ADDR          uint32 = 0x7F << 25 //+ Lock OCRAM_TZ_ADDR field for changes
	LOCK_OCRAM_TZ_ADDR_0        uint32 = 0x00 << 25 //  Field is not locked
	LOCK_OCRAM_TZ_ADDR_1        uint32 = 0x01 << 25 //  Field is locked (read access only)
)

const (
	NIDENn                       = 0
	DBG_ENn                      = 1
	SEC_ERR_RESPn                = 2
	DCPKEY_OCOTP_OR_KEYMUXn      = 4
	OCRAM_TZ_ENn                 = 8
	OCRAM_TZ_ADDRn               = 9
	LOCK_NIDENn                  = 16
	LOCK_DBG_ENn                 = 17
	LOCK_SEC_ERR_RESPn           = 18
	LOCK_DCPKEY_OCOTP_OR_KEYMUXn = 20
	LOCK_OCRAM_TZ_ENn            = 24
	LOCK_OCRAM_TZ_ADDRn          = 25
)

// GPR11 bits
const (
	M7_APC_AC_R0_CTRL   uint32 = 0x03 << 0 //+ Access control of memory region-0
	M7_APC_AC_R0_CTRL_0 uint32 = 0x00 << 0 //  No access protection
	M7_APC_AC_R0_CTRL_1 uint32 = 0x01 << 0 //  M7 debug protection enabled
	M7_APC_AC_R0_CTRL_2 uint32 = 0x02 << 0 //  FlexSPI access protection
	M7_APC_AC_R0_CTRL_3 uint32 = 0x03 << 0 //  Both M7 debug and FlexSPI access are protected
	M7_APC_AC_R1_CTRL   uint32 = 0x03 << 2 //+ Access control of memory region-1
	M7_APC_AC_R1_CTRL_0 uint32 = 0x00 << 2 //  No access protection
	M7_APC_AC_R1_CTRL_1 uint32 = 0x01 << 2 //  M7 debug protection enabled
	M7_APC_AC_R1_CTRL_2 uint32 = 0x02 << 2 //  FlexSPI access protection
	M7_APC_AC_R1_CTRL_3 uint32 = 0x03 << 2 //  Both M7 debug and FlexSPI access are protected
	M7_APC_AC_R2_CTRL   uint32 = 0x03 << 4 //+ Access control of memory region-2
	M7_APC_AC_R2_CTRL_0 uint32 = 0x00 << 4 //  No access protection
	M7_APC_AC_R2_CTRL_1 uint32 = 0x01 << 4 //  M7 debug protection enabled
	M7_APC_AC_R2_CTRL_2 uint32 = 0x02 << 4 //  FlexSPI access protection
	M7_APC_AC_R2_CTRL_3 uint32 = 0x03 << 4 //  Both M7 debug and FlexSPI access are protected
	M7_APC_AC_R3_CTRL   uint32 = 0x03 << 6 //+ Access control of memory region-3
	M7_APC_AC_R3_CTRL_0 uint32 = 0x00 << 6 //  No access protection
	M7_APC_AC_R3_CTRL_1 uint32 = 0x01 << 6 //  M7 debug protection enabled
	M7_APC_AC_R3_CTRL_2 uint32 = 0x02 << 6 //  FlexSPI access protection
	M7_APC_AC_R3_CTRL_3 uint32 = 0x03 << 6 //  Both M7 debug and FlexSPI access are protected
	BEE_DE_RX_EN        uint32 = 0x0F << 8 //+ BEE data decryption of memory region-n (n = 3 to 0)
)

const (
	M7_APC_AC_R0_CTRLn = 0
	M7_APC_AC_R1_CTRLn = 2
	M7_APC_AC_R2_CTRLn = 4
	M7_APC_AC_R3_CTRLn = 6
	BEE_DE_RX_ENn      = 8
)

// GPR12 bits
const (
	FLEXIO1_IPG_STOP_MODE uint32 = 0x01 << 0 //+ FlexIO1 stop mode selection. Cannot change when ipg_stop is asserted.
	FLEXIO1_IPG_DOZE      uint32 = 0x01 << 1 //+ FLEXIO1 ipg_doze mode
	FLEXIO2_IPG_STOP_MODE uint32 = 0x01 << 2 //+ FlexIO2 stop mode selection. Cannot change when ipg_stop is asserted.
	FLEXIO2_IPG_DOZE      uint32 = 0x01 << 3 //+ FLEXIO2 ipg_doze mode
	ACMP_IPG_STOP_MODE    uint32 = 0x01 << 4 //+ ACMP stop mode selection. Cannot change when ipg_stop is asserted.
	FLEXIO3_IPG_STOP_MODE uint32 = 0x01 << 5 //+ FlexIO3 stop mode selection. Cannot change when ipg_stop is asserted.
	FLEXIO3_IPG_DOZE      uint32 = 0x01 << 6 //+ FLEXIO3 ipg_doze mode
)

const (
	FLEXIO1_IPG_STOP_MODEn = 0
	FLEXIO1_IPG_DOZEn      = 1
	FLEXIO2_IPG_STOP_MODEn = 2
	FLEXIO2_IPG_DOZEn      = 3
	ACMP_IPG_STOP_MODEn    = 4
	FLEXIO3_IPG_STOP_MODEn = 5
	FLEXIO3_IPG_DOZEn      = 6
)

// GPR13 bits
const (
	ARCACHE_USDHC  uint32 = 0x01 << 0  //+ uSDHC block cacheable attribute value of AXI read transactions
	AWCACHE_USDHC  uint32 = 0x01 << 1  //+ uSDHC block cacheable attribute value of AXI write transactions
	CANFD_STOP_REQ uint32 = 0x01 << 4  //+ CANFD stop request.
	CACHE_ENET     uint32 = 0x01 << 7  //+ ENET block cacheable attribute value of AXI transactions
	CACHE_USB      uint32 = 0x01 << 13 //+ USB block cacheable attribute value of AXI transactions
	CANFD_STOP_ACK uint32 = 0x01 << 20 //+ CANFD stop acknowledge.
)

const (
	ARCACHE_USDHCn  = 0
	AWCACHE_USDHCn  = 1
	CANFD_STOP_REQn = 4
	CACHE_ENETn     = 7
	CACHE_USBn      = 13
	CANFD_STOP_ACKn = 20
)

// GPR14 bits
const (
	ACMP1_CMP_IGEN_TRIM_DN uint32 = 0x01 << 0  //+ reduces ACMP1 internal bias current by 30%
	ACMP2_CMP_IGEN_TRIM_DN uint32 = 0x01 << 1  //+ reduces ACMP2 internal bias current by 30%
	ACMP3_CMP_IGEN_TRIM_DN uint32 = 0x01 << 2  //+ reduces ACMP3 internal bias current by 30%
	ACMP4_CMP_IGEN_TRIM_DN uint32 = 0x01 << 3  //+ reduces ACMP4 internal bias current by 30%
	ACMP1_CMP_IGEN_TRIM_UP uint32 = 0x01 << 4  //+ increases ACMP1 internal bias current by 30%
	ACMP2_CMP_IGEN_TRIM_UP uint32 = 0x01 << 5  //+ increases ACMP2 internal bias current by 30%
	ACMP3_CMP_IGEN_TRIM_UP uint32 = 0x01 << 6  //+ increases ACMP3 internal bias current by 30%
	ACMP4_CMP_IGEN_TRIM_UP uint32 = 0x01 << 7  //+ increases ACMP4 internal bias current by 30%
	ACMP1_SAMPLE_SYNC_EN   uint32 = 0x01 << 8  //+ ACMP1 sample_lv source select
	ACMP2_SAMPLE_SYNC_EN   uint32 = 0x01 << 9  //+ ACMP2 sample_lv source select
	ACMP3_SAMPLE_SYNC_EN   uint32 = 0x01 << 10 //+ ACMP3 sample_lv source select
	ACMP4_SAMPLE_SYNC_EN   uint32 = 0x01 << 11 //+ ACMP4 sample_lv source select
	CM7_CFGITCMSZ          uint32 = 0x0F << 16 //+ ITCM total size configuration
	CM7_CFGITCMSZ_0        uint32 = 0x00 << 16 //  0 KB (No ITCM)
	CM7_CFGITCMSZ_3        uint32 = 0x03 << 16 //  4 KB
	CM7_CFGITCMSZ_4        uint32 = 0x04 << 16 //  8 KB
	CM7_CFGITCMSZ_5        uint32 = 0x05 << 16 //  16 KB
	CM7_CFGITCMSZ_6        uint32 = 0x06 << 16 //  32 KB
	CM7_CFGITCMSZ_7        uint32 = 0x07 << 16 //  64 KB
	CM7_CFGITCMSZ_8        uint32 = 0x08 << 16 //  128 KB
	CM7_CFGITCMSZ_9        uint32 = 0x09 << 16 //  256 KB
	CM7_CFGITCMSZ_10       uint32 = 0x0A << 16 //  512 KB
	CM7_CFGDTCMSZ          uint32 = 0x0F << 20 //+ DTCM total size configuration
	CM7_CFGDTCMSZ_0        uint32 = 0x00 << 20 //  0 KB (No DTCM)
	CM7_CFGDTCMSZ_3        uint32 = 0x03 << 20 //  4 KB
	CM7_CFGDTCMSZ_4        uint32 = 0x04 << 20 //  8 KB
	CM7_CFGDTCMSZ_5        uint32 = 0x05 << 20 //  16 KB
	CM7_CFGDTCMSZ_6        uint32 = 0x06 << 20 //  32 KB
	CM7_CFGDTCMSZ_7        uint32 = 0x07 << 20 //  64 KB
	CM7_CFGDTCMSZ_8        uint32 = 0x08 << 20 //  128 KB
	CM7_CFGDTCMSZ_9        uint32 = 0x09 << 20 //  256 KB
	CM7_CFGDTCMSZ_10       uint32 = 0x0A << 20 //  512 KB
)

const (
	ACMP1_CMP_IGEN_TRIM_DNn = 0
	ACMP2_CMP_IGEN_TRIM_DNn = 1
	ACMP3_CMP_IGEN_TRIM_DNn = 2
	ACMP4_CMP_IGEN_TRIM_DNn = 3
	ACMP1_CMP_IGEN_TRIM_UPn = 4
	ACMP2_CMP_IGEN_TRIM_UPn = 5
	ACMP3_CMP_IGEN_TRIM_UPn = 6
	ACMP4_CMP_IGEN_TRIM_UPn = 7
	ACMP1_SAMPLE_SYNC_ENn   = 8
	ACMP2_SAMPLE_SYNC_ENn   = 9
	ACMP3_SAMPLE_SYNC_ENn   = 10
	ACMP4_SAMPLE_SYNC_ENn   = 11
	CM7_CFGITCMSZn          = 16
	CM7_CFGDTCMSZn          = 20
)

// GPR16 bits
const (
	INIT_ITCM_EN         uint32 = 0x01 << 0 //+ ITCM enable initialization out of reset
	INIT_DTCM_EN         uint32 = 0x01 << 1 //+ DTCM enable initialization out of reset
	FLEXRAM_BANK_CFG_SEL uint32 = 0x01 << 2 //+ FlexRAM bank config source select
)

const (
	INIT_ITCM_ENn         = 0
	INIT_DTCM_ENn         = 1
	FLEXRAM_BANK_CFG_SELn = 2
)

// GPR17: FlexRAM bank config value

// GPR18 bits
const (
	LOCK_M7_APC_AC_R0_BOT uint32 = 0x01 << 0       //+ lock M7_APC_AC_R0_BOT field for changes
	M7_APC_AC_R0_BOT      uint32 = 0x1FFFFFFF << 3 //+ APC end address of memory region-0
)

const (
	LOCK_M7_APC_AC_R0_BOTn = 0
	M7_APC_AC_R0_BOTn      = 3
)

// GPR19 bits
const (
	LOCK_M7_APC_AC_R0_TOP uint32 = 0x01 << 0       //+ lock M7_APC_AC_R0_TOP field for changes
	M7_APC_AC_R0_TOP      uint32 = 0x1FFFFFFF << 3 //+ APC start address of memory region-0
)

const (
	LOCK_M7_APC_AC_R0_TOPn = 0
	M7_APC_AC_R0_TOPn      = 3
)

// GPR20 bits
const (
	LOCK_M7_APC_AC_R1_BOT uint32 = 0x01 << 0       //+ lock M7_APC_AC_R1_BOT field for changes
	M7_APC_AC_R1_BOT      uint32 = 0x1FFFFFFF << 3 //+ APC end address of memory region-1
)

const (
	LOCK_M7_APC_AC_R1_BOTn = 0
	M7_APC_AC_R1_BOTn      = 3
)

// GPR21 bits
const (
	LOCK_M7_APC_AC_R1_TOP uint32 = 0x01 << 0       //+ lock M7_APC_AC_R1_TOP field for changes
	M7_APC_AC_R1_TOP      uint32 = 0x1FFFFFFF << 3 //+ APC start address of memory region-1
)

const (
	LOCK_M7_APC_AC_R1_TOPn = 0
	M7_APC_AC_R1_TOPn      = 3
)

// GPR22 bits
const (
	LOCK_M7_APC_AC_R2_BOT uint32 = 0x01 << 0       //+ lock M7_APC_AC_R2_BOT field for changes
	M7_APC_AC_R2_BOT      uint32 = 0x1FFFFFFF << 3 //+ APC end address of memory region-2
)

const (
	LOCK_M7_APC_AC_R2_BOTn = 0
	M7_APC_AC_R2_BOTn      = 3
)

// GPR23 bits
const (
	LOCK_M7_APC_AC_R2_TOP uint32 = 0x01 << 0       //+ lock M7_APC_AC_R2_TOP field for changes
	M7_APC_AC_R2_TOP      uint32 = 0x1FFFFFFF << 3 //+ APC start address of memory region-2
)

const (
	LOCK_M7_APC_AC_R2_TOPn = 0
	M7_APC_AC_R2_TOPn      = 3
)

// GPR24 bits
const (
	LOCK_M7_APC_AC_R3_BOT uint32 = 0x01 << 0       //+ lock M7_APC_AC_R3_BOT field for changes
	M7_APC_AC_R3_BOT      uint32 = 0x1FFFFFFF << 3 //+ APC end address of memory region-3
)

const (
	LOCK_M7_APC_AC_R3_BOTn = 0
	M7_APC_AC_R3_BOTn      = 3
)

// GPR25 bits
const (
	LOCK_M7_APC_AC_R3_TOP uint32 = 0x01 << 0       //+ lock M7_APC_AC_R3_TOP field for changes
	M7_APC_AC_R3_TOP      uint32 = 0x1FFFFFFF << 3 //+ APC start address of memory region-3
)

const (
	LOCK_M7_APC_AC_R3_TOPn = 0
	M7_APC_AC_R3_TOPn      = 3
)

// GPR26: GPIO1 and GPIO6 share same IO MUX function, GPIO_MUX1 selects one GPIO function.

// GPR27: GPIO2 and GPIO7 share same IO MUX function, GPIO_MUX2 selects one GPIO function.

// GPR28: GPIO3 and GPIO8 share same IO MUX function, GPIO_MUX3 selects one GPIO function.

// GPR29: GPIO4 and GPIO9 share same IO MUX function, GPIO_MUX4 selects one GPIO function.

// GPR30: Start address of flexspi1 and flexspi2

// GPR31: End address of flexspi1 and flexspi2

// GPR32: Offset address of flexspi1 and flexspi2

// GPR33 bits
const (
	OCRAM2_TZ_EN          uint32 = 0x01 << 0  //+ OCRAM2 TrustZone (TZ) enable.
	OCRAM2_TZ_ADDR        uint32 = 0x7F << 1  //+ OCRAM2 TrustZone (TZ) start address
	LOCK_OCRAM2_TZ_EN     uint32 = 0x01 << 16 //+ Lock OCRAM2_TZ_EN field for changes
	LOCK_OCRAM2_TZ_ADDR   uint32 = 0x7F << 17 //+ Lock OCRAM2_TZ_ADDR field for changes
	LOCK_OCRAM2_TZ_ADDR_0 uint32 = 0x00 << 17 //  Field is not locked
	LOCK_OCRAM2_TZ_ADDR_1 uint32 = 0x01 << 17 //  Field is locked (read access only)
)

const (
	OCRAM2_TZ_ENn        = 0
	OCRAM2_TZ_ADDRn      = 1
	LOCK_OCRAM2_TZ_ENn   = 16
	LOCK_OCRAM2_TZ_ADDRn = 17
)

// GPR34 bits
const (
	SIP_TEST_MUX_BOOT_PIN_SEL uint32 = 0xFF << 0 //+ Boot Pin select in SIP_TEST_MUX
	SIP_TEST_MUX_QSPI_SIP_EN  uint32 = 0x01 << 8 //+ Enable SIP_TEST_MUX
)

const (
	SIP_TEST_MUX_BOOT_PIN_SELn = 0
	SIP_TEST_MUX_QSPI_SIP_ENn  = 8
)
