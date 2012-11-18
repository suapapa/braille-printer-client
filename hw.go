// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"
	"unicode/utf8"
)

/*  Filename:    hw.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Description: Code to control braille label printer
 */

func Emboss(bs string) {
	rLen := utf8.RuneCountInString(bs)
	if rLen > 255 {
		log.Println("Too long string (%d chars)to print\n", rLen)
		return
	}

	log.Printf("Ebossing %s (%d)...", bs, rLen)

	conn, err := net.Dial("unixgram", "/tmp/blp_uds")
	if err != nil {
		log.Printf("Failed connect server: %s\n", err)
		return
	}
	defer conn.Close()

	pktBuf := make([]byte, 4+1+rLen) // header, len, data
	/* pktBuf := make([]byte, ) */
	pktBuf[0] = '#'
	pktBuf[1] = '#'
	pktBuf[2] = '#'
	pktBuf[3] = '#'
	pktBuf[4] = byte(rLen)

	j := 0;
	for _, bc := range bs {
		if bc&0xff00 != 0x2800 {
			log.Printf("%c is not braille character!\n", bc)
		}
		pktBuf[j+5] = byte(bc & 0xff)
		j += 1
	}

	conn.Write(pktBuf)
}
