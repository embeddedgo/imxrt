#!/usr/bin/env python3
# -*- coding: utf-8 -*-

""" i.MX binary loading lib.

This library allows loading and extracting data from a
given binary that should be loaded to i.MX device.

Currently parsing should work for i.MX53 and i.MX 6 devices.

This version works with Python 3.x (tested with 3.4)

"""

import struct

_version = "1.0"

#valid offsets within binary
_imx_ivt_valid_offsets=(0,256,1024,4096)

#valid IVT versions
_imx_ivt_valid_versions=(0x40,0x41,0x42,0x43)

_imx_dcd_valid_tags=(0xD2,)
#ref. manual specifies version 0x41, but SW uses 0x40 sometimes and it works OK
_imx_dcd_valid_versions=(0x40,0x41,0x42,0x43)

class ImxRegInterpreter:
    def __init__(self):
        pass
    def interpret(self, register, value):
        raise NotImplementedError()

class ImxBin:
    def __init__(self,filename):

        self.file = open(filename,'rb')
        self.file_data = self.file.read()

        #find offset of IVT table. Typically it's 0 or 1k
        self.ivt_offset = 0
        ivt_found=False

        while(not ivt_found):
            self.ivt_offset = self.file_data.find(0xD1,self.ivt_offset)
            if(self.ivt_offset == -1):
                raise(Exception("IVT table not found in given file"))
            if((self.ivt_offset in _imx_ivt_valid_offsets)
               and (self.file_data[self.ivt_offset+3]
                    in _imx_ivt_valid_versions)):
                ivt_found = True
            else:
                self.ivt_offset = self.ivt_offset + 1
                if(self.ivt_offset > max(_imx_ivt_valid_offsets)):
                    raise(Exception("IVT table not found in given file"))

        #read out IVT from the image. IVT length field is big endian...
        self.ivt_length, = struct.unpack_from(
            ">xH",
            self.file_data,
            self.ivt_offset)
        
        self.ivt_entry, \
        self.ivt_dcd, \
        self.ivt_boot_data, \
        self.ivt_self, \
        self.ivt_csf = struct.unpack_from(
            "<xxxxIxxxxIIII",
            self.file_data,
            self.ivt_offset)
        
        #read out DCD data as well. Often used, useful for debugging as well.
        self.dcd_length = 0
        self.dcd_data = None

        if(self.ivt_dcd != 0):
            dcd_offset = self.ivt_offset + (self.ivt_dcd - self.ivt_self)
            
            dcd_tag, self.dcd_length, dcd_version = \
                     struct.unpack_from(">BHB",
                                        self.file_data,
                                        dcd_offset)

            if ((dcd_tag not in _imx_dcd_valid_tags) or \
                (dcd_version not in _imx_dcd_valid_versions)):
                raise Exception("DCD table malformed")

            if(self.dcd_length>4):
                self.dcd_data=self.file_data[dcd_offset+4: \
                                             dcd_offset+4+(self.dcd_length-4)]
            
        #check if this is a plugin image and initialize corresponding data
        boot_data_offset = self.ivt_offset + \
                           (self.ivt_boot_data - self.ivt_self)    
        self.boot_data_start, self.boot_data_length, self.boot_data_plugin = \
                              struct.unpack_from("<III",
                                                 self.file_data,
                                                 boot_data_offset)
        
    def dump_ivt_data(self):
        print("ivt_length={:X}\nivt_entry={:X}\nivt_dcd={:X}\n"
              "ivt_boot_data={:X}\nivt_self={:X}\nivt_csf={:X}".format(
            self.ivt_length,
            self.ivt_entry,
            self.ivt_dcd,
            self.ivt_boot_data,
            self.ivt_self,
            self.ivt_csf))

    def dump_dcd_data(self, prettyFormatter=None):

        if(self.dcd_data==None):
            print("No DCD data in table")
            return

        if(prettyFormatter and isinstance(prettyFormatter,ImxRegInterpreter) is False):
            raise Exception("Wrong register interpreter given")
        
        offset = 0

        while(offset < (self.dcd_length-4)):
            tag, length, parameter = struct.unpack_from(
                ">BHB",self.dcd_data,offset)
            print("tag={:#04x} length={:#06x} param={:#04x}".format(tag, length, parameter))
            length = length - 4 #write command header
            offset = offset + 4 #read offset
            if(tag == 0xCC):
                #dump all write commands
                while(length):
                    address,value = struct.unpack_from(
                        ">II",self.dcd_data,offset)
                    offset = offset+8
                    length = length-8
                    if(prettyFormatter):
                        prettyFormatter.interpret(address,value)
                    else:
                        binvalue="{:032b}".format(value)
                        binvalue= binvalue[0:4]+" "+\
                            binvalue[4:8]+" "+\
                            binvalue[8:12]+" "+\
                            binvalue[12:16]+" "+\
                            binvalue[16:20]+" "+\
                            binvalue[20:24]+" "+\
                            binvalue[24:28]+" "+\
                            binvalue[28:32]
                        print("{:#010x}={:#010x}({})".format(address,value,binvalue))
                        #print("{:#010x} = {:#010x}".format(address,value))
                    
            else:
                #So far other tags than 0xCC not seen in use. Raise exception.
                raise Exception("DCD command with tag {:#02x} is not supported".\
                                format(tag))
            
                                
class Imx6RegInterpreter(ImxRegInterpreter):
    _imx6_regs = {
        #MMDC registers of interest
        0x021B0000:"MMDC0_MDCTL",
        0x021B4000:"MMDC1_MDCTL",
        0x021B0004:"MMDC0_MDPDC",
        0x021B4004:"MMDC1_MDPDC",
        0x021B0008:"MMDC0_MDOTC",
        0x021B4008:"MMDC1_MDOTC",
        0x021B000C:"MMDC0_MDCFG0",
        0x021B400C:"MMDC1_MDCFG0",
        0x021B0010:"MMDC0_MDCFG1",
        0x021B4010:"MMDC1_MDCFG1",
        0x021B0014:"MMDC0_MDCFG2",
        0x021B4014:"MMDC1_MDCFG2",
        0x021B0018:"MMDC0_MDMISC",
        0x021B4018:"MMDC1_MDMISC",
        0x021B001C:"MMDC0_MDSCR",
        0x021B401C:"MMDC1_MDSCR",
        0x021B0020:"MMDC0_MDREF",
        0x021B4020:"MMDC1_MDREF",
        0x021B002C:"MMDC0_MDRWD",
        0x021B402C:"MMDC1_MDRWD",
        0x021B0030:"MMDC0_MDOR",
        0x021B4030:"MMDC1_MDOR",
        0x021B0038:"MMDC0_MDCFG3LP",
        0x021B4038:"MMDC1_MDCFG3LP",
        0x021B003C:"MMDC0_MDMR4",
        0x021B403C:"MMDC1_MDMR4",
        0x021B0040:"MMDC0_MDASP",
        0x021B4040:"MMDC1_MDASP",
        0x021B0404:"MMDC1_MAPSR",
        0x021B4404:"MMDC1_MAPSR",
        0x021B0800:"MMDC0_MPZQHWCTRL",
        0x021B4800:"MMDC1_MPZQHWCTRL",
        0x021B0804:"MMDC0_MPZQSWCTRL",
        0x021B4804:"MMDC1_MPZQSWCTRL",
        0x021B0808:"MMDC0_MPWLGCR",
        0x021B4808:"MMDC1_MPWLGCR",
        0x021B080C:"MMDC0_MPWLDECTRL0",
        0x021B480C:"MMDC1_MPWLDECTRL0",
        0x021B0810:"MMDC0_MPWLDECTRL1",
        0x021B4810:"MMDC1_MPWLDECTRL1",
        0x021B0818:"MMDC0_MPODTCTRL",
        0x021B4818:"MMDC1_MPODTCTRL",
        0x021B081C:"MMDC0_MPRDDQBY0DL",
        0x021B481C:"MMDC1_MPRDDQBY0DL",
        0x021B0820:"MMDC0_MPRDDQBY1DL",
        0x021B4820:"MMDC1_MPRDDQBY1DL",
        0x021B0824:"MMDC0_MPRDDQBY2DL",
        0x021B4824:"MMDC1_MPRDDQBY2DL",
        0x021B0828:"MMDC0_MPRDDQBY3DL",
        0x021B4828:"MMDC1_MPRDDQBY3DL",
        0x021B082C:"MMDC0_MPWRDQBY0DL",
        0x021B482C:"MMDC1_MPWRDQBY0DL",
        0x021B0830:"MMDC0_MPWRDQBY1DL",
        0x021B4830:"MMDC1_MPWRDQBY1DL",
        0x021B0834:"MMDC0_MPWRDQBY2DL",
        0x021B4834:"MMDC1_MPWRDQBY2DL",
        0x021B0838:"MMDC0_MPWRDQBY3DL",
        0x021B4838:"MMDC1_MPWRDQBY3DL",
        0x021B083C:"MMDC0_MPDGCTRL0",
        0x021B483C:"MMDC1_MPDGCTRL0",
        0x021B0840:"MMDC0_MPDGCTRL1",
        0x021B4840:"MMDC1_MPDGCTRL1",
        0x021B0848:"MMDC0_MPRDDLCTL",
        0x021B4848:"MMDC1_MPRDDLCTL",
        0x021B0850:"MMDC0_MPWRDLCTL",
        0x021B4850:"MMDC1_MPWRDLCTL",
        0x021B0858:"MMDC0_MPSDCTRL",
        0x021B4858:"MMDC1_MPSDCTRL",
        0x021B0860:"MMDC0_MPRDDLHWCTL",
        0x021B4860:"MMDC1_MPRDDLHWCTL",
        0x021B0864:"MMDC0_MPWRDLHWCTL",
        0x021B4864:"MMDC1_MPWRDLHWCTL",
        0x021B088C:"MMDC0_MPPDCMPR1",
        0x021B488C:"MMDC1_MPPDCMPR1",
        0x021B0890:"MMDC0_MPPDCMPR2",
        0x021B4890:"MMDC1_MPPDCMPR2",
        0x021B08B8:"MMDC0_MPMUR0",
        0x021B48B8:"MMDC1_MPMUR0",
        0x021B08BC:"MMDC0_MPWRCADL",
        0x021B48BC:"MMDC1_MPWRCADL",
        }
    
    def interpret(self, register, value):
        binvalue="{:032b}".format(value)
        binvalue= binvalue[0:4]+" "+\
                    binvalue[4:8]+" "+\
                    binvalue[8:12]+" "+\
                    binvalue[12:16]+" "+\
                    binvalue[16:20]+" "+\
                    binvalue[20:24]+" "+\
                    binvalue[24:28]+" "+\
                    binvalue[28:32]
        if (register in Imx6RegInterpreter._imx6_regs):    
            print("{:18} = {:#010x}({})".\
                  format(Imx6RegInterpreter._imx6_regs[register],value,binvalue))
            detvalue=self._interpret_mx6_details(register,value)
            if(detvalue):
                print("\t"+" ".join(detvalue))
        else:
            print("{:#010x}={:#010x}({})".format(register,value,binvalue))

    
    
    #provide some more info on interesting registers for debugging
    #memory and startup issues.
    def _interpret_mx6_details(self,register,value):
        retvalue=[]
        if((register&(~0x00004000))==0x021B0000):
            #MMDCx_MDCTL
            retvalue.append("SDE_0="+str((value&0x80000000)>>31))
            retvalue.append("SDE_1="+str((value&0x40000000)>>30))
            retvalue.append("ROW="+str(((value&0x07000000)>>24)+11))
            retvalue.append("COL="+str(((value&0x00700000)>>20)+9))
            retvalue.append("BL="+str(((value&0x00080000)>>19)*4+4))
            retvalue.append("DSIZ="+str(16<<((value&0x00030000)>>16)))
        elif((register&(~0x00004000))==0x021B0004):
            #MMDCx_MDPDC
            if(value&0x70000000):
                retvalue.append("PRCT_1="+str(1<<((value&0x70000000)>>28)))
            else:
                retvalue.append("PRCT_1=DIS")
            if(value&0x07000000):
                retvalue.append("PRCT_0="+str(1<<((value&0x07000000)>>24)))
            else:
                retvalue.append("PRCT_0=DIS")
            retvalue.append("tCKE="+str(((value&0x00070000)>>16)+1))
            if(value&0x0000F000):
                retvalue.append("PWDT_1="+str(16<<(((value&0x0000F000)>>12)-1)))
            else:
                retvalue.append("PWDT_1=DIS")
            if(value&0x00000F00):
                retvalue.append("PWDT_0="+str(16<<(((value&0x00000F00)>>8)-1)))
            else:
                retvalue.append("PWDT_0=DIS")
            retvalue.append("SLOW_PD="+str((value&0x00000080)>>7))
            retvalue.append("tCKSRX="+str((value&0x00000038)>>3))
            retvalue.append("tCKSRE="+str(value&0x00000007))
        elif((register&(~0x00004000))==0x021B0008):
            #MMDCx_MDOTC
            retvalue.append("tAOFPD="+str(((value&0x38000000)>>27)+1))
            retvalue.append("tAONPD="+str(((value&0x03000000)>>24)+1))
            retvalue.append("tANPD="+str(((value&0x00F00000)>>20)+1))
            retvalue.append("tAXPD="+str(((value&0x000F0000)>>16)+1))
        elif((register&(~0x00004000))==0x021B000C):
            #MMDCx_MDCFG0
            retvalue.append("tRFC="+str(((value&0xFF000000)>>24)+1))
            retvalue.append("tXS="+str(((value&0x00FF0000)>>16)+1))
            retvalue.append("tXP="+str(((value&0x0000E000)>>13)+1))
            retvalue.append("tXPDLL="+str(((value&0x00001E00)>>9)+1))
            retvalue.append("tFAW="+str(((value&0x000001F0)>>4)+1))
            retvalue.append("tCL="+str((value&0x0000000F)+3))
        elif((register&(~0x00004000))==0x021B0010):
            #MMDCx_MDCFG1
            retvalue.append("tRCD="+str(((value&0xE0000000)>>29)+1))
            retvalue.append("tRP="+str(((value&0x1C000000)>>26)+1))
            retvalue.append("tRC="+str(((value&0x03E00000)>>21)+1))
            retvalue.append("tRAS="+str(((value&0x001F0000)>>16)+1))
            if(value&0x00008000):
                retvalue.append("tRPA=tRP+1")
            else:
                retvalue.append("tRPA=tRP")
            retvalue.append("tWR="+str(((value&0x00000E00)>>9)+1))
            retvalue.append("tMRD="+str(((value&0x000001E0)>>5)+1))
            retvalue.append("_tCWL="+hex((value&0x00000007)+1))
        elif((register&(~0x00004000))==0x021B0014):
            #MMDCx_MDCFG2
            retvalue.append("tDLLK="+str(((value&0x01FF0000)>>16)+1))
            retvalue.append("tRTP="+str(((value&0x000001C0)>>6)+1))
            retvalue.append("tWTR="+str(((value&0x00000038)>>3)+1))
            retvalue.append("tRRD="+str((value&0x00000007)+1))
        elif((register&(~0x00004000))==0x021B0018):
            #MMDCx_MDMISC
            retvalue.append("CALIB_PER_CS="+str((value&0x00100000)>>20))
            retvalue.append("ADDR_MIRROR="+str((value&0x00080000)>>19))
            retvalue.append("LHD="+str((value&0x00040000)>>18))
            retvalue.append("WALAT="+str((value&0x00030000)>>16))
            retvalue.append("BI_ON="+str((value&0x00001000)>>12))
            retvalue.append("LPDDR2_S2="+str((value&0x00000800)>>11))
            retvalue.append("MIF3_MODE="+hex((value&0x00000600)>>9))
            retvalue.append("RALAT="+str((value&0x000001C0)>>6))
            retvalue.append("DDR_4_BANK="+str((value&0x00000020)>>5))
            ddr_type=(value&0x00000018)>>3
            if(ddr_type==0):
                retvalue.append("DDR_TYPE=DDR3")
            elif(ddr_type==1):
                retvalue.append("DDR_TYPE=LPDDR2")
            else:
                retvalue.append("DDR_TYPE=RESERVED")
            retvalue.append("LPDDR2_2CH="+str((value&0x00000004)>>2))
            retvalue.append("RST="+str((value&0x00000002)>>1))
        elif((register&(~0x00004000))==0x021B001C):
            #MMDCx_MDSCR
            retvalue.append("CMD_ADDR_MSB_MR_OP="+hex((value&0xFF000000)>>24))
            retvalue.append("CMD_ADDR_LSB_MR_OP="+hex((value&0x00FF0000)>>16))
            retvalue.append("CON_REQ="+str((value&0x00008000)>>15))
            retvalue.append("WL_EN="+str((value&0x00000200)>>9))
            cmd=(value&0x00000070)>>4
            if(cmd==0):
                retvalue.append("CMD=normal op")
            elif(cmd==1):
                retvalue.append("CMD=precharge all")
            elif(cmd==2):
                retvalue.append("CMD=auto refresh cmd")
            elif(cmd==3):
                retvalue.append("CMD=LMR/MRW")
            elif(cmd==4):
                retvalue.append("CMD=ZQ calibration")
            elif(cmd==5):
                retvalue.append("CMD=precharge all")
            elif(cmd==6):
                retvalue.append("CMD=MRR")
            elif(cmd==7):
                retvalue.append("CMD=RESERVED")
            retvalue.append("CMD_CS="+str((value&0x00000008)>>3))
            retvalue.append("CMD_BA="+str(value&0x00000007))
        elif((register&(~0x00004000))==0x021B0040):
            #MMDCx_MDASP
            cs0_end=value&0x0000007F
            cs0_end_mbit=(cs0_end+1)*256
            cs0_end_mbyte=cs0_end_mbit/8
            retvalue.append("CS0_END="+hex(cs0_end)+"("+str(cs0_end_mbit)+\
                            "Mb, "+str(cs0_end_mbyte)+"MB)")
        return retvalue
        
        

if __name__ == "__main__":

    import argparse
    import sys

    if(sys.version_info[0] == 2):
        print("You're using Python 2.x, which is not supported by this script")
        quit()

    parser = argparse.ArgumentParser(description="\
                                     Parse and dump IVT/DCD data \
                                     from a given i.MX binary. \
                                     ")
    parser.add_argument("imxbinary", help="filename containing i.MX binary \
                        data")
    parser.add_argument("-d","--device", type=str, choices=["imx53",\
                                                            "imx6dq",\
                                                            "imx6sdl",\
                                                            "imx6sx",\
                                                            "imx6sl"],\
                        default="imx6dq", help="choose target device")

    args = parser.parse_args()

    f = ImxBin(args.imxbinary)
    print("=================IVT data===================")
    f.dump_ivt_data()
    print("=================DCD data===================")
    interpr = None
    if(args.device.startswith("imx6")):
        interpr = Imx6RegInterpreter()
    f.dump_dcd_data(interpr)
    
