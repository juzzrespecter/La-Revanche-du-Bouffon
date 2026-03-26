package bmp

import "io"

type BMPHeader struct {
	magic  uint16
	size   uint32
	res_1  uint16 // reserved
	res_2  uint16 // reserved
	offset uint32
}

type DIBHeader struct {
	size            uint32
	imageWidth      uint32
	imageHeight     uint32
	colorPlanes     uint16
	bitsPerPixel    uint16
	compression     uint32
	rawImageSize    uint32
	pixelPerMetreX  uint32
	pixelPerMetreY  uint32
	colorN          uint32
	importantColorN uint32
}

var BMPVersion = map[uint32]string{
	12:  "BITMAPCOREHEADER",
	64:  "OS22XBITMAPHEADER",
	16:  "OS22XBITMAPHEADER",
	40:  "BITMAPINFOHEADER",
	52:  "BITMAPV2INFOHEADER",
	56:  "BITMAPV3INFOHEADER",
	108: "BITMAPV4HEADER",
	124: "BITMAPV5HEADER",
}

var compressionMethods = map[uint32]string{
	0:  "none",
	1:  "RLE 8-bit/pixel",
	2:  "RLE 4-bit/pixel",
	3:  "OS22XBITMAPHEADER: Huffman 1D",
	4:  "OS22XBITMAPHEADER: RLE-24",
	5:  "OS22XBITMAPHEADER: RLE-24",
	6:  "RGBA bit field masks",
	11: "none",
	12: "RLE-8",
	13: "RLE-4",
}

func parseBMPHeader(f *io.Reader) {

}

func parseDIBHeader() {

}

func Bmp(f io.Reader) (string, error) {
	return "TODO", nil
}
