// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usb

import (
	"unicode/utf16"
	"unsafe"
)

type StringDescriptor struct {
	Len  uint8
	Type uint8
	data [128]uint16
}

func NewStringDescriptor(s string) *typeStringDescriptor {
	data := utf16.Encode([]rune(s))
	d := (*StringDescriptor)(unsafe.Pointer(&make([]uint16, len(data)+1)[0]))
	d.Len = len(data) * 2
	d.Type = 3
	copy(d.data[:], data)
	return d
}

func (d *StringDescriptor) String() string {
	return string(utf16.Decode(d.data[:d.Len/2]))
}

func (d *StringDescriptor) UTF16() []uint16 {
	return d.data[:d.Len/2]
}
