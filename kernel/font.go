package main

import (
	_ "embed"
	"unsafe"
)

//go:embed zap-light16.psf
var zaplight16 []byte

const (
	PSF1_MAGIC0 = 0x36
	PSF1_MAGIC1 = 0x04

	fontWidth  = 8
	fontHeight = 16
)

type PSF1_HEADER struct {
	magic    [2]byte
	mode     byte
	charsize byte
}

type Font struct {
	header          PSF1_HEADER
	glyphBufferSize uint32
}

var font Font

func initPSF(screen []uint32, pitch int) {
	font.header = *(*PSF1_HEADER)(unsafe.Pointer(unsafe.SliceData(zaplight16)))
	if font.header.magic[0] != PSF1_MAGIC0 || font.header.magic[1] != PSF1_MAGIC1 {
		hcf()
	}
	font.glyphBufferSize = uint32(font.header.charsize) * 256
	if font.header.mode == 1 { // 512 glyph mode
		font.glyphBufferSize = uint32(font.header.charsize) * 512
	}
}

func drawString(screen []uint32, width, pitch int, str string, xOff, yOff int, fgColor, bgColor uint32) {
	// TODO: support UTF8
	for i := 0; i < len(str); i++ {
		drawChar(screen, width, pitch, str[i], xOff+(i*fontWidth), yOff, fgColor, bgColor)
	}
}

func drawChar(screen []uint32, width, pitch int, c byte, xOff, yOff int, fgColor, bgColor uint32) {
	const glyphStartIndex = 4 // skip the header
	fontPtr := zaplight16[uint32(c)*uint32(font.header.charsize)+glyphStartIndex:]
	for y := yOff; y < yOff+fontHeight; y++ {
		for x := xOff; x < xOff+fontWidth; x++ {
			if (fontPtr[y-yOff] & (0b10000000 >> (x - xOff))) > 0 {
				putPixel(screen, width, pitch, x, y, fgColor)
			} else {
				putPixel(screen, width, pitch, x, y, bgColor)
			}
		}
	}
}
