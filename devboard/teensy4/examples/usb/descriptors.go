// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains descriptor data from the Teensyduino Core Library.

/* Teensyduino Core Library
 * http://www.pjrc.com/teensy/
 * Copyright (c) 2017 PJRC.COM, LLC.
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * 1. The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * 2. If the Software is incorporated into a build system that allows
 * selection among a list of target devices, then similar target
 * devices manufactured by PJRC.COM must be included in the list of
 * target devices and selectable in the same manner.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

var descriptors = map[uint32]string{
	0x0100_0000: deviceDescr,
	0x0200_0000: usbConfigDescr480,
	0x0300_0000: string0,
	0x0301_0409: usbManufacturerName,
	0x0302_0409: usbProductName,
	0x0303_0409: usbSerialNumber,
	0x0600_0000: qualifierDescr,
	0x0700_0000: usbConfigDescr12,
}

const (
	ctrlPktSize = "\x40"     // 64
	vendorID    = "\xC0\x16" // 0x16C0 "Van Ooijen Technische Informatica"
	productID   = "\x83\x04" // 0x0483 "Teensyduino Serial"
)

const deviceDescr = "" +
	"\x12" + // bLength
	"\x01" + // bDescriptorType
	"\x00\x02" + // bcdUSB
	"\x00" + // bDeviceClass
	"\x00" + // bDeviceSubClass
	"\x00" + // bDeviceProtocol
	ctrlPktSize + // bMaxPacketSize0
	vendorID + // idVendor
	productID + // idProduct
	"\x81\x02" + // bcdDevice
	"\x01" + // iManufacturer
	"\x02" + // iProduct
	"\x03" + // iSerialNumber
	"\x01" // bNumConfigurations

const qualifierDescr = "" +
	"\x0a" + // bLength
	"\x06" + // bDescriptorType
	"\x00\x02" + // bcdUSB
	"\x00" + // bDeviceClass
	"\x00" + // bDeviceSubClass
	"\x00" + // bDeviceProtocol
	ctrlPktSize + // bMaxPacketSize0
	"\x01" + // bNumConfigurations
	"\x00" // bReserved

const (
	intNum = "\x04"

	acm0_StatusInt = "\x00"
	acm0_DataInt   = "\x01"
	acm0_StatusIN  = "\x81" // 1 IN
	acm0_DataIN    = "\x82" // 2 IN
	acm0_DataOUT   = "\x02" // 2 OUT

	acm1_StatusInt = "\x02"
	acm1_DataInt   = "\x03"
	acm1_StatusIN  = "\x83" // 3 IN
	acm1_DataIN    = "\x84" // 4 IN
	acm1_DataOUT   = "\x04" // 4 OUT

	acmStatusSize  = "\x10\x00" // 16
	acmDataSize480 = "\x00\x02" // 512
	acmDataSize12  = "\x40\x00" // 64
)

const usbConfigDescrLen = "\x8d\x00"

const usbConfigDescr480 = "" +
	"\x09" + // bLength;
	"\x02" + // bDescriptorType;
	usbConfigDescrLen + // wTotalLength
	intNum + // bNumInterfaces
	"\x01" + // bConfigurationValue, use 1 to select this conig
	"\x00" + // iConfiguration, 0 means no string descriptor for this conf
	"\xC0" + // bmAttributes, Self Powered
	"\x32" + // bMaxPower, 50 * 2 mA = 100 mA

	// interface association descriptor, USB ECN, Table 9-Z
	"\x08" + // bLength
	"\x0b" + // bDescriptorType
	acm0_StatusInt + // bFirstInterface
	"\x02" + // bInterfaceCount
	"\x02" + // bFunctionClass
	"\x02" + // bFunctionSubClass
	"\x01" + // bFunctionProtocol
	"\x00" + // iFunction

	// configuration for 480 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm0_StatusInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x01" + // bNumEndpoints
	"\x02" + // bInterfaceClass
	"\x02" + // bInterfaceSubClass
	"\x01" + // bInterfaceProtocol
	"\x00" + // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x00" + // bDescriptorSubtype
	"\x10\x01" + // bcdCDC

	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x01" + // bDescriptorSubtype
	"\x01" + // bmCapabilities
	"\x01" + // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	"\x04" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x02" + // bDescriptorSubtype
	"\x06" + // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x06" + // bDescriptorSubtype
	acm0_StatusInt + // bMasterInterface
	acm0_DataInt + // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_StatusIN + // bEndpointAddress
	"\x03" + // bmAttributes (0x03=intr)
	acmStatusSize + // wMaxPacketSize
	"\x05" + // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm0_DataInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x02" + // bNumEndpoints
	"\x0A" + // bInterfaceClass
	"\x00" + // bInterfaceSubClass
	"\x00" + // bInterfaceProtocol
	"\x00" + // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_DataIN + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize480 + // wMaxPacketSize
	"\x00" + // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_DataOUT + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize480 + // wMaxPacketSize
	"\x00" + // bInterval

	// interface association descriptor, USB ECN, Table 9-Z
	"\x08" + // bLength
	"\x0b" + // bDescriptorType
	acm1_StatusInt + // bFirstInterface
	"\x02" + // bInterfaceCount
	"\x02" + // bFunctionClass
	"\x02" + // bFunctionSubClass
	"\x01" + // bFunctionProtocol
	"\x00" + // iFunction

	// configuration for 480 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm1_StatusInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x01" + // bNumEndpoints
	"\x02" + // bInterfaceClass
	"\x02" + // bInterfaceSubClass
	"\x01" + // bInterfaceProtocol
	"\x00" + // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x00" + // bDescriptorSubtype
	"\x10\x01" + // bcdCDC

	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x01" + // bDescriptorSubtype
	"\x01" + // bmCapabilities
	"\x01" + // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	"\x04" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x02" + // bDescriptorSubtype
	"\x06" + // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x06" + // bDescriptorSubtype
	acm1_StatusInt + // bMasterInterface
	acm1_DataInt + // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_StatusIN + // bEndpointAddress
	"\x03" + // bmAttributes (0x03=intr)
	acmStatusSize + // wMaxPacketSize
	"\x05" + // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm1_DataInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x02" + // bNumEndpoints
	"\x0A" + // bInterfaceClass
	"\x00" + // bInterfaceSubClass
	"\x00" + // bInterfaceProtocol
	"\x00" + // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_DataIN + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize480 + // wMaxPacketSize
	"\x00" + // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_DataOUT + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize480 + // wMaxPacketSize
	"\x00" // bInterval

const usbConfigDescr12 = "" +
	// configuration descriptor, USB spec 9.6.3, page 264-266, Table 9-10
	"\x09" + // bLength;
	"\x02" + // bDescriptorType;
	usbConfigDescrLen + // wTotalLength
	intNum + // bNumInterfaces
	"\x01" + // bConfigurationValue
	"\x00" + // iConfiguration
	"\xC0" + // bmAttributes
	"\x32" + // bMaxPower

	// interface association descriptor, USB ECN, Table 9-Z
	"\x08" + // bLength
	"\x0B" + // bDescriptorType
	acm0_StatusInt + // bFirstInterface
	"\x02" + // bInterfaceCount
	"\x02" + // bFunctionClass
	"\x02" + // bFunctionSubClass
	"\x01" + // bFunctionProtocol
	"\x00" + // iFunction

	// configuration for 12 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm0_StatusInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x01" + // bNumEndpoints
	"\x02" + // bInterfaceClass
	"\x02" + // bInterfaceSubClass
	"\x01" + // bInterfaceProtocol
	"\x00" + // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x00" + // bDescriptorSubtype
	"\x10\x01" + // bcdCDC
	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x01" + // bDescriptorSubtype
	"\x01" + // bmCapabilities
	"\x01" + // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	"\x04" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x02" + // bDescriptorSubtype
	"\x06" + // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x06" + // bDescriptorSubtype
	acm0_StatusInt + // bMasterInterface
	acm0_DataInt + // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_StatusIN + // bEndpointAddress
	"\x03" + // bmAttributes (0x03=intr)
	acmStatusSize + // wMaxPacketSize
	"\x10" + // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm0_DataInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x02" + // bNumEndpoints
	"\x0A" + // bInterfaceClass
	"\x00" + // bInterfaceSubClass
	"\x00" + // bInterfaceProtocol
	"\x00" + // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_DataIN + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize12 + // wMaxPacketSize
	"\x00" + // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm0_DataOUT + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize12 + // wMaxPacketSize
	"\x00" + // bInterval

	// interface association descriptor, USB ECN, Table 9-Z
	"\x08" + // bLength
	"\x0B" + // bDescriptorType
	acm1_StatusInt + // bFirstInterface
	"\x02" + // bInterfaceCount
	"\x02" + // bFunctionClass
	"\x02" + // bFunctionSubClass
	"\x01" + // bFunctionProtocol
	"\x00" + // iFunction

	// configuration for 12 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm1_StatusInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x01" + // bNumEndpoints
	"\x02" + // bInterfaceClass
	"\x02" + // bInterfaceSubClass
	"\x01" + // bInterfaceProtocol
	"\x00" + // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x00" + // bDescriptorSubtype
	"\x10\x01" + // bcdCDC
	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x01" + // bDescriptorSubtype
	"\x01" + // bmCapabilities
	"\x01" + // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	"\x04" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x02" + // bDescriptorSubtype
	"\x06" + // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	"\x05" + // bFunctionLength
	"\x24" + // bDescriptorType
	"\x06" + // bDescriptorSubtype
	acm1_StatusInt + // bMasterInterface
	acm1_DataInt + // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_StatusIN + // bEndpointAddress
	"\x03" + // bmAttributes (0x03=intr)
	acmStatusSize + // wMaxPacketSize
	"\x10" + // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	"\x09" + // bLength
	"\x04" + // bDescriptorType
	acm1_DataInt + // bInterfaceNumber
	"\x00" + // bAlternateSetting
	"\x02" + // bNumEndpoints
	"\x0A" + // bInterfaceClass
	"\x00" + // bInterfaceSubClass
	"\x00" + // bInterfaceProtocol
	"\x00" + // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_DataIN + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize12 + // wMaxPacketSize
	"\x00" + // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	"\x07" + // bLength
	"\x05" + // bDescriptorType
	acm1_DataOUT + // bEndpointAddress
	"\x02" + // bmAttributes (0x02=bulk)
	acmDataSize12 + // wMaxPacketSize
	"\x00" // bInterval

// string0 descriptor lists supportedd languages
const string0 = "" +
	"\x04" +
	"\x03" +
	"\x09\x04" // English (United States)

const usbManufacturerName = "" +
	"\x18" + // bLength
	"\x03" + // bDescriptorType
	"E\x00" +
	"m\x00" +
	"b\x00" +
	"e\x00" +
	"d\x00" +
	"d\x00" +
	"e\x00" +
	"d\x00" +
	" \x00" +
	"G\x00" +
	"o\x00"

const usbProductName = "" +
	"\x1C" + // bLength
	"\x03" + // bDescriptorType
	"C\x00" +
	"o\x00" +
	"n\x00" +
	"s\x00" +
	"o\x00" +
	"l\x00" +
	"e\x00" +
	" \x00" +
	"+\x00" +
	" \x00" +
	"A\x00" +
	"U\x00" +
	"X\x00"

const usbSerialNumber = "" +
	"\x14" + // bLength
	"\x03" + // bDescriptorType
	"1\x00" +
	"2\x00" +
	"3\x00" +
	"4\x00" +
	"5\x00" +
	"6\x00" +
	"7\x00" +
	"8\x00" +
	"9\x00"
