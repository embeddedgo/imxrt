// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"strings"
	"time"

	"github.com/embeddedgo/imxrt/devboard/teensy4/board/pins"
	"github.com/embeddedgo/imxrt/hal/lpuart"
	"github.com/embeddedgo/imxrt/hal/lpuart/lpuart1"
	"github.com/embeddedgo/imxrt/hal/system/console/uartcon"
	"github.com/embeddedgo/imxrt/hal/usb"
	"github.com/embeddedgo/imxrt/hal/usb/serial"
)

var usbd *usb.Device

func cdcSetLineCoding(cr *usb.ControlRequest) int {
	fmt.Printf("cdcSetLineCoding: %+v\r\n", *cr)
	fmt.Printf(" -interface: %d\r\n", cr.Index)
	d := cr.Data
	if len(d) < 7 {
		fmt.Println(" len(d) =", len(d))
		return 0
	}
	fmt.Printf(" -baudrate:  %d\r\n", d[0]+d[1]<<8+d[2]<<16+d[3]<<24)
	fmt.Printf(" -stop bits: %.1f\r\n", float32(d[4]+2)/2)
	par := d[5]
	if par > 5 {
		par = 5
	}
	par *= 4
	fmt.Printf(" -parity:    %s\r\n", "noneodd evenmarkspacunkn"[par:par+4])
	fmt.Printf(" -data bits: %d\r\n", d[6])
	return 0
}

func cdcSetControlLineState(cr *usb.ControlRequest) int {
	fmt.Printf("cdcSetControlLineState: %+v\r\n", *cr)
	return 0
}

var txt = []byte(`
000102030405060708090a0b0c0d0e0f0g0h0i0j0k0l0m0n0o0p0r0s0t0u0v0w0x0y0z
101112131415161718191a1b1c1d1e1f1g1h1i1j1k1l1m1n1o1p1r1s1t1u1v1w1x1y1z
202122232425262728292a2b2c2d2e2f2g2h2i2j2k2l2m2n2o2p2r2s2t2u2v2w2x2y2z
303132333435363738393a3b3c3d3e3f3g3h3i3j3k3l3m3n3o3p3r3s3t3u3v3w3x3y3z
404142434445464748494a4b4c4d4e4f4g4h4i4j4k4l4m4n4o4p4r4s4t4u4v4w4x4y4z
505152535455565758595a5b5c5d5e5f5g5h5i5j5k5l5m5n5o5p5r5s5t5u5v5w5x5y5z
606162636465666768696a6b6c6d6e6f6g6h6i6j6k6l6m6n6o6p6r6s6t6u6v6w6x6y6z
707172737475767778797a7b7c7d7e7f7g7h7i7j7k7l7m7n7o7p7r7s7t7u7v7w7x7y7z
808182838485868788898a8b8c8d8e8f8g8h8i8j8k8l8m8n8o8p8r8s8t8u8v8w8x8y8z
909192939495969798999a9b9c9d9e9f9g9h9i9j9k9l9m9n9o9p9r9s9t9u9v9w9x9y9z
a0a1a2a3a4a5a6a7a8a9aaabacadaeafagahaiajakalamanaoaparasatauavawaxayaz
b0b1b2b3b4b5b6b7b8b9babbbcbdbebfbgbhbibjbkblbmbnbobpbrbsbtbubvbwbxbybz
c0c1c2c3c4c5c6c7c8c9cacbcccdcecfcgchcicjckclcmcncocpcrcsctcucvcwcxcycz
d0d1d2d3d4d5d6d7d8d9dadbdcdddedfdgdhdidjdkdldmdndodpdrdsdtdudvdwdxdydz
e0e1e2e3e4e5e6e7e8e9eaebecedeeefegeheiejekelemeneoepereseteuevewexeyez
f0f1f2f3f4f5f6f7f8f9fafbfcfdfefffgfhfifjfkflfmfnfofpfrfsftfufvfwfxfyfz
_0_1_2_3_4_5_6_7_8_9_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_r_s_t_u_v_w_x_y_z
`)

func main() {
	// IO pins
	conTx := pins.P24
	conRx := pins.P25

	// Serial console
	uartcon.Setup(lpuart1.Driver(), conRx, conTx, lpuart.Word8b, 115200, "UART1")
	fmt.Println("Start!")

	const (
		config = 1
		in     = 2 // input endpoint, host perspective, device Tx
		out    = 2 // output endopint, host prespective, device Rx
		maxPkt = 512
	)

	usbd = usb.NewDevice(1)
	usbd.Init(rtos.IntPrioLow, descriptors, false)
	se := serial.New(usbd, out, in, maxPkt, config)
	usbd.Enable()

	abuf := make([]byte, maxPkt*4)

usbNotReady:
	fmt.Println("Waiting for USB...")
	usbd.WaitConfig(config)
	fmt.Println("USB is ready.")

	time.Sleep(5 * time.Second)
	fmt.Println("Go!")

	for buf := txt; ; buf = buf[1 : len(buf)-1] {
		if len(buf) == 0 {
			buf = txt
		}
		if false {
			n, err := se.Read(abuf)
			if err != nil {
				if e, ok := err.(*usb.Error); ok && e.NotReady() {
					goto usbNotReady
				}
				fmt.Printf("\n!! Error:\n %v\n\n", err)
				continue
			}

			fmt.Printf("received %d bytes: %s\n", n, abuf[:n])

			if strings.TrimSpace(string(abuf[:n])) == "reset" {
				fmt.Println("* Reset! *")
				usbd.Disable()
				usbd.Enable()
				fmt.Println("* Go! *")
			}
		}

		n, err := se.Write(buf)
		if err != nil {
			if e, ok := err.(*usb.Error); ok && e.NotReady() {
				goto usbNotReady
			}
			fmt.Println(err)
			continue
		}

		fmt.Printf("sent %d bytes\n", n)
	}
}

//go:interrupthandler
func USB_OTG1_Handler() {
	usbd.ISR()
}
