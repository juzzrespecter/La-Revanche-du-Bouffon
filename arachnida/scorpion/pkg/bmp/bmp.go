package bmp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type BMPHeader struct {
	//Magic  uint16
	Size   uint32
	Res_1  uint16 // reserved
	Res_2  uint16 // reserved
	Offset uint32
}

type DIBHeader struct {
	Size            uint32
	ImageWidth      uint32
	ImageHeight     uint32
	ColorPlanes     uint16
	BitsPerPixel    uint16
	Compression     uint32
	RawImageSize    uint32
	PixelPerMetreX  uint32
	PixelPerMetreY  uint32
	ColorN          uint32
	ImportantColorN uint32
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

func parseBMPHeader(f io.Reader) (string, error) {
	bmpHdr := &BMPHeader{}
	err := binary.Read(f, binary.LittleEndian, bmpHdr)
	if err != nil {
		return "", err
	}
	bmpHdrInfo := fmt.Sprintf(
		"BMP Image size:              %d\n"+
			"Bitmap offset:               %p\n",
		bmpHdr.Size,
		&bmpHdr.Offset)
	return bmpHdrInfo, nil
}

func parseDIBHeader(f io.Reader) (string, error) {
	dibHdr := &DIBHeader{}
	err := binary.Read(f, binary.LittleEndian, dibHdr)
	if err != nil {
		return "", err
	}
	dibHdrInfo := fmt.Sprintf(
		"Version:                     %s\n"+
			"Image width:                 %d\n"+
			"Image height:                %d\n"+
			"Color planes:                %d\n"+
			"Bits per pixel:              %d\n"+
			"Compression:                 %s\n"+
			"Raw image size:              %d\n"+
			"Pixel per metre X:           %d\n"+
			"Pixel per metre Y:           %d\n"+
			"Colors:                      %d\n"+
			"Important colors:            %d\n",
		BMPVersion[dibHdr.Size],
		dibHdr.ImageHeight,
		dibHdr.ImageWidth,
		dibHdr.ColorPlanes,
		dibHdr.BitsPerPixel,
		compressionMethods[dibHdr.Compression],
		dibHdr.RawImageSize,
		dibHdr.PixelPerMetreX,
		dibHdr.PixelPerMetreY,
		dibHdr.ColorN,
		dibHdr.ImportantColorN,
	)
	return dibHdrInfo, nil
}

func Bmp(f io.Reader, file string) (string, error) {
	magic := make([]byte, 2)
	f.Read(magic)
	if !bytes.Equal(magic, []byte{0x42, 0x4D}) {
		return "", fmt.Errorf("%s: not a bmp file", file)
	}
	bmpHeaderInfo, err := parseBMPHeader(f)
	if err != nil {
		return "", nil
	}
	dibHdrInfo, err := parseDIBHeader(f)
	if err != nil {
		return "", nil
	}
	return bmpHeaderInfo + dibHdrInfo, nil
}
