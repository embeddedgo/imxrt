// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains descriptor data from Teensyduino Core Library.

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

var descriptors = map[uint32][]byte{
	0x0100_0000: deviceDescr[:],
	0x0200_0000: usbConfigDescr480[:],
	0x0300_0000: string0[:],
	0x0301_0409: usbManufacturerName[:],
	0x0302_0409: usbProductName[:],
	0x0303_0409: usbSerialNumber[:],
	0x0700_0000: usbConfigDescr12[:],
	0x0600_0000: qualifierDescr[:],
}

const (
	ep0pktSize = 64
	vendorID   = 0x16C0
	productID  = 0x0483
)

const deviceDescrStr = "" +
	"\x16" + // bLength
	"\x01" + // bDescriptorType
	"\x00\x02" + // bcdUSB
	"\x00" + // bDeviceClass
	"\x00" + // bDeviceSubClass
	"\x00" + // bDeviceProtocol
	"\x40" + // bMaxPacketSize0
	"\xc0\x16" + // idVendor
	"\x83\x04" + // idProduct
	"\x81\x02" + // bcdDevice
	"\x01" + // iManufacturer
	"\x02" + // iProduct
	"\x03" + // iSerialNumber
	"\x01" // bNumConfigurations

var deviceDescr = [18]byte{
	18,   // bLength
	1,    // bDescriptorType
	0, 2, // bcdUSB, 0x0200 means USB 2.0
	0,                              // bDeviceClass, 0 means see interface class
	0,                              // bDeviceSubClass
	0,                              // bDeviceProtocol
	ep0pktSize,                     // bMaxPacketSize0
	vendorID & 0xff, vendorID >> 8, // idVendor
	productID & 0xff, productID >> 8, // idProduct
	0x81, 0x02, // bcdDevice
	1, // iManufacturer
	2, // iProduct
	3, // iSerialNumber
	1, // bNumConfigurations
}

var qualifierDescr = [10]byte{
	10,         // bLength
	6,          // bDescriptorType
	0x00, 0x02, // bcdUSB
	0,          // bDeviceClass
	0,          // bDeviceSubClass
	0,          // bDeviceProtocol
	ep0pktSize, // bMaxPacketSize0
	1,          // bNumConfigurations
	0,          // bReserved
}

const (
	intNum = 4

	acm0_StatusInt = 0
	acm0_DataInt   = 1
	acm0_StatusIN  = 1 | 0x80
	acm0_DataIN    = 2 | 0x80
	acm0_DataOUT   = 2

	acm1_StatusInt = 2
	acm1_DataInt   = 3
	acm1_StatusIN  = 3 | 0x80
	acm1_DataIN    = 4 | 0x80
	acm1_DataOUT   = 4

	acmStatusSize  = 16
	acmDataSize480 = 512
	acmDataSize12  = 64
)

const usbConfigDescrLen = 141

var usbConfigDescr480 = [usbConfigDescrLen]byte{
	9,                        // bLength;
	2,                        // bDescriptorType;
	usbConfigDescrLen & 0xff, // wTotalLength
	usbConfigDescrLen >> 8,
	intNum, // bNumInterfaces
	1,      // bConfigurationValue, use 1 to select this conig
	0,      // iConfiguration, 0 means no string descriptor for this conf
	0xC0,   // bmAttributes, Self Powered
	50,     // bMaxPower, 50 * 2 mA = 100 mA

	// interface association descriptor, USB ECN, Table 9-Z
	8,              // bLength
	11,             // bDescriptorType
	acm0_StatusInt, // bFirstInterface
	2,              // bInterfaceCount
	0x02,           // bFunctionClass
	0x02,           // bFunctionSubClass
	0x01,           // bFunctionProtocol
	0,              // iFunction

	// configuration for 480 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,              // bLength
	4,              // bDescriptorType
	acm0_StatusInt, // bInterfaceNumber
	0,              // bAlternateSetting
	1,              // bNumEndpoints
	0x02,           // bInterfaceClass
	0x02,           // bInterfaceSubClass
	0x01,           // bInterfaceProtocol
	0,              // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	5,          // bFunctionLength
	0x24,       // bDescriptorType
	0x00,       // bDescriptorSubtype
	0x10, 0x01, // bcdCDC

	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	5,    // bFunctionLength
	0x24, // bDescriptorType
	0x01, // bDescriptorSubtype
	0x01, // bmCapabilities
	1,    // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	4,    // bFunctionLength
	0x24, // bDescriptorType
	0x02, // bDescriptorSubtype
	0x06, // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	5,              // bFunctionLength
	0x24,           // bDescriptorType
	0x06,           // bDescriptorSubtype
	acm0_StatusInt, // bMasterInterface
	acm0_DataInt,   // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                             // bLength
	5,                                             // bDescriptorType
	acm0_StatusIN,                                 // bEndpointAddress
	0x03,                                          // bmAttributes (0x03=intr)
	byte(acmStatusSize), byte(acmStatusSize >> 8), // wMaxPacketSize
	5, // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,            // bLength
	4,            // bDescriptorType
	acm0_DataInt, // bInterfaceNumber
	0,            // bAlternateSetting
	2,            // bNumEndpoints
	0x0A,         // bInterfaceClass
	0x00,         // bInterfaceSubClass
	0x00,         // bInterfaceProtocol
	0,            // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                          // bLength
	5,                                          // bDescriptorType
	acm0_DataIN,                                // bEndpointAddress
	0x02,                                       // bmAttributes (0x02=bulk)
	acmDataSize480 & 0xff, acmDataSize480 >> 8, // wMaxPacketSize
	0, // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                          // bLength
	5,                                          // bDescriptorType
	acm0_DataOUT,                               // bEndpointAddress
	0x02,                                       // bmAttributes (0x02=bulk)
	acmDataSize480 & 0xff, acmDataSize480 >> 8, // wMaxPacketSize
	0, // bInterval

	// interface association descriptor, USB ECN, Table 9-Z
	8,              // bLength
	11,             // bDescriptorType
	acm1_StatusInt, // bFirstInterface
	2,              // bInterfaceCount
	0x02,           // bFunctionClass
	0x02,           // bFunctionSubClass
	0x01,           // bFunctionProtocol
	0,              // iFunction

	// configuration for 480 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,              // bLength
	4,              // bDescriptorType
	acm1_StatusInt, // bInterfaceNumber
	0,              // bAlternateSetting
	1,              // bNumEndpoints
	0x02,           // bInterfaceClass
	0x02,           // bInterfaceSubClass
	0x01,           // bInterfaceProtocol
	0,              // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	5,          // bFunctionLength
	0x24,       // bDescriptorType
	0x00,       // bDescriptorSubtype
	0x10, 0x01, // bcdCDC

	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	5,    // bFunctionLength
	0x24, // bDescriptorType
	0x01, // bDescriptorSubtype
	0x01, // bmCapabilities
	1,    // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	4,    // bFunctionLength
	0x24, // bDescriptorType
	0x02, // bDescriptorSubtype
	0x06, // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	5,              // bFunctionLength
	0x24,           // bDescriptorType
	0x06,           // bDescriptorSubtype
	acm1_StatusInt, // bMasterInterface
	acm1_DataInt,   // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                             // bLength
	5,                                             // bDescriptorType
	acm1_StatusIN,                                 // bEndpointAddress
	0x03,                                          // bmAttributes (0x03=intr)
	byte(acmStatusSize), byte(acmStatusSize >> 8), // wMaxPacketSize
	5, // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,            // bLength
	4,            // bDescriptorType
	acm1_DataInt, // bInterfaceNumber
	0,            // bAlternateSetting
	2,            // bNumEndpoints
	0x0A,         // bInterfaceClass
	0x00,         // bInterfaceSubClass
	0x00,         // bInterfaceProtocol
	0,            // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                          // bLength
	5,                                          // bDescriptorType
	acm1_DataIN,                                // bEndpointAddress
	0x02,                                       // bmAttributes (0x02=bulk)
	acmDataSize480 & 0xff, acmDataSize480 >> 8, // wMaxPacketSize
	0, // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                          // bLength
	5,                                          // bDescriptorType
	acm1_DataOUT,                               // bEndpointAddress
	0x02,                                       // bmAttributes (0x02=bulk)
	acmDataSize480 & 0xff, acmDataSize480 >> 8, // wMaxPacketSize
	0, // bInterval
}

var usbConfigDescr12 = [usbConfigDescrLen]byte{
	// configuration descriptor, USB spec 9.6.3, page 264-266, Table 9-10
	9,                        // bLength;
	2,                        // bDescriptorType;
	usbConfigDescrLen & 0xff, // wTotalLength
	usbConfigDescrLen >> 8,
	intNum, // bNumInterfaces
	1,      // bConfigurationValue
	0,      // iConfiguration
	0xC0,   // bmAttributes
	50,     // bMaxPower

	// interface association descriptor, USB ECN, Table 9-Z
	8,              // bLength
	11,             // bDescriptorType
	acm0_StatusInt, // bFirstInterface
	2,              // bInterfaceCount
	0x02,           // bFunctionClass
	0x02,           // bFunctionSubClass
	0x01,           // bFunctionProtocol
	0,              // iFunction

	// configuration for 12 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,              // bLength
	4,              // bDescriptorType
	acm0_StatusInt, // bInterfaceNumber
	0,              // bAlternateSetting
	1,              // bNumEndpoints
	0x02,           // bInterfaceClass
	0x02,           // bInterfaceSubClass
	0x01,           // bInterfaceProtocol
	0,              // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	5,          // bFunctionLength
	0x24,       // bDescriptorType
	0x00,       // bDescriptorSubtype
	0x10, 0x01, // bcdCDC
	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	5,    // bFunctionLength
	0x24, // bDescriptorType
	0x01, // bDescriptorSubtype
	0x01, // bmCapabilities
	1,    // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	4,    // bFunctionLength
	0x24, // bDescriptorType
	0x02, // bDescriptorSubtype
	0x06, // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	5,              // bFunctionLength
	0x24,           // bDescriptorType
	0x06,           // bDescriptorSubtype
	acm0_StatusInt, // bMasterInterface
	acm0_DataInt,   // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                // bLength
	5,                // bDescriptorType
	acm0_StatusIN,    // bEndpointAddress
	0x03,             // bmAttributes (0x03=intr)
	acmStatusSize, 0, // wMaxPacketSize
	16, // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,            // bLength
	4,            // bDescriptorType
	acm0_DataInt, // bInterfaceNumber
	0,            // bAlternateSetting
	2,            // bNumEndpoints
	0x0A,         // bInterfaceClass
	0x00,         // bInterfaceSubClass
	0x00,         // bInterfaceProtocol
	0,            // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                        // bLength
	5,                                        // bDescriptorType
	acm0_DataIN,                              // bEndpointAddress
	0x02,                                     // bmAttributes (0x02=bulk)
	acmDataSize12 & 0xff, acmDataSize12 >> 8, // wMaxPacketSize
	0, // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                        // bLength
	5,                                        // bDescriptorType
	acm0_DataOUT,                             // bEndpointAddress
	0x02,                                     // bmAttributes (0x02=bulk)
	acmDataSize12 & 0xff, acmDataSize12 >> 8, // wMaxPacketSize
	0, // bInterval

	// interface association descriptor, USB ECN, Table 9-Z
	8,              // bLength
	11,             // bDescriptorType
	acm1_StatusInt, // bFirstInterface
	2,              // bInterfaceCount
	0x02,           // bFunctionClass
	0x02,           // bFunctionSubClass
	0x01,           // bFunctionProtocol
	0,              // iFunction

	// configuration for 12 Mbit/sec speed
	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,              // bLength
	4,              // bDescriptorType
	acm1_StatusInt, // bInterfaceNumber
	0,              // bAlternateSetting
	1,              // bNumEndpoints
	0x02,           // bInterfaceClass
	0x02,           // bInterfaceSubClass
	0x01,           // bInterfaceProtocol
	0,              // iInterface

	// CDC Header Functional Descriptor, CDC Spec 5.2.3.1, Table 26
	5,          // bFunctionLength
	0x24,       // bDescriptorType
	0x00,       // bDescriptorSubtype
	0x10, 0x01, // bcdCDC
	// Call Management Functional Descriptor, CDC Spec 5.2.3.2, Table 27
	5,    // bFunctionLength
	0x24, // bDescriptorType
	0x01, // bDescriptorSubtype
	0x01, // bmCapabilities
	1,    // bDataInterface

	// Abstract Control Management Functional Descriptor, CDC Spec 5.2.3.3, Table 28
	4,    // bFunctionLength
	0x24, // bDescriptorType
	0x02, // bDescriptorSubtype
	0x06, // bmCapabilities

	// Union Functional Descriptor, CDC Spec 5.2.3.8, Table 33
	5,              // bFunctionLength
	0x24,           // bDescriptorType
	0x06,           // bDescriptorSubtype
	acm1_StatusInt, // bMasterInterface
	acm1_DataInt,   // bSlaveInterface0

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                // bLength
	5,                // bDescriptorType
	acm1_StatusIN,    // bEndpointAddress
	0x03,             // bmAttributes (0x03=intr)
	acmStatusSize, 0, // wMaxPacketSize
	16, // bInterval

	// interface descriptor, USB spec 9.6.5, page 267-269, Table 9-12
	9,            // bLength
	4,            // bDescriptorType
	acm1_DataInt, // bInterfaceNumber
	0,            // bAlternateSetting
	2,            // bNumEndpoints
	0x0A,         // bInterfaceClass
	0x00,         // bInterfaceSubClass
	0x00,         // bInterfaceProtocol
	0,            // iInterface

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                        // bLength
	5,                                        // bDescriptorType
	acm1_DataIN,                              // bEndpointAddress
	0x02,                                     // bmAttributes (0x02=bulk)
	acmDataSize12 & 0xff, acmDataSize12 >> 8, // wMaxPacketSize
	0, // bInterval

	// endpoint descriptor, USB spec 9.6.6, page 269-271, Table 9-13
	7,                                        // bLength
	5,                                        // bDescriptorType
	acm1_DataOUT,                             // bEndpointAddress
	0x02,                                     // bmAttributes (0x02=bulk)
	acmDataSize12 & 0xff, acmDataSize12 >> 8, // wMaxPacketSize
	0, // bInterval
}

// string0 descriptor lists supportedd languages
var string0 = [4]byte{
	4,
	3,
	9, 4, // English (United States)
}

var usbManufacturerName = [24]byte{
	24,
	3,
	'E', 0,
	'm', 0,
	'b', 0,
	'e', 0,
	'd', 0,
	'd', 0,
	'e', 0,
	'd', 0,
	' ', 0,
	'G', 0,
	'o', 0,
}

var usbProductName = [28]byte{
	28,
	3,
	'C', 0,
	'o', 0,
	'n', 0,
	's', 0,
	'o', 0,
	'l', 0,
	'e', 0,
	' ', 0,
	'+', 0,
	' ', 0,
	'A', 0,
	'U', 0,
	'X', 0,
}

var usbSerialNumber = [20]byte{
	20,
	3,
	'1', 0,
	'2', 0,
	'3', 0,
	'4', 0,
	'5', 0,
	'6', 0,
	'7', 0,
	'8', 0,
	'9', 0,
}
