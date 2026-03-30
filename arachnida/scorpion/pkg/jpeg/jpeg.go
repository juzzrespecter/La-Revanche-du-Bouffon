package jpeg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"scorpion/internal/ifd"
)

// Parseo recursivo de las tablas IFD
//
// * base:  bytearray que parte desde la cabecera TIFF
// * o:     offset desde el comienzo de la tabla
// * order: byte order
func parseIFDTable(base []byte, o uint32, order binary.ByteOrder) (string, uint32, error) {
	var fmtTag string

	n := order.Uint16(base[o : o+2])
	ifdOffset := uint32(0)
	for i := 0; i < int(n); i++ {
		ifdOffset = o + 12*uint32(i) + 2
		ifdRawEntry := base[ifdOffset : ifdOffset+12]
		ifdEntry := &ifd.IFDEntry{
			Tag: order.Uint16(ifdRawEntry[:2]),
			Fmt: order.Uint16(ifdRawEntry[2:4]),
			N:   order.Uint32(ifdRawEntry[4:8]),
			Val: order.Uint32(ifdRawEntry[8:]),
		}
		switch ifdEntry.Tag {
		case 0x8769, 0x8825: // SubIFDs
			tags, o, err := parseIFDTable(base, ifdEntry.Val, order)
			if err != nil {
				return "", 0, err
			}
			fmtTag = fmtTag + tags
			if o == 0 {
				return fmtTag, o, nil
			}
		}
		fmtTag = fmtTag + ifdEntry.FormatIFD(base, order)
	}
	x := order.Uint32(base[ifdOffset+12 : ifdOffset+12+8])
	return fmtTag, x, nil
}

func parseExif(f io.Reader) (string, error) {
	buffer := make([]byte, 8)
	err := binary.Read(f, binary.LittleEndian, buffer)
	if err != nil {
		return "", err
	}
	exifHdr := buffer[2:]
	if !bytes.Equal(exifHdr, []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00}) {
		return "", nil
	}
	base, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	// TIFF Header
	var order binary.ByteOrder
	switch string(base[:2]) {
	case "MM":
		order = binary.BigEndian
	case "II":
		order = binary.LittleEndian
	default:
		return "", fmt.Errorf("invalid byte order")
	}
	if order.Uint16(base[2:4]) != 42 {
		return "", fmt.Errorf("invalid TIFF magic number")
	}
	tiffOffset := order.Uint32(base[4:8])
	nextOffset := tiffOffset
	var tags string
	for nextOffset != 0 {
		tableTags, nextOffset, err := parseIFDTable(base, tiffOffset, order)
		if err != nil {
			return "", err
		}
		tags = tags + tableTags
		if nextOffset == 0 {
			break
		}
	}
	return tags, nil
}

func parseJfif(f io.Reader) (string, error) {
	return "TODO", nil
}

func Jpeg(f io.Reader) (string, error) {
	var jpegInfo string

	magic := make([]byte, 2)
	f.Read(magic)
	if !bytes.Equal(magic, []byte{0xFF, 0xD8}) {
		return "", fmt.Errorf("not a png file")
	}
	appMarker := make([]byte, 2)
	f.Read(appMarker)
	switch {
	case bytes.Equal(appMarker, []byte{0xff, 0xe1}):
		info, err := parseExif(f)
		if err != nil {
			return "", err
		}
		jpegInfo = info
	case bytes.Equal(appMarker, []byte{0xff, 0xe1}):
		info, err := parseJfif(f)
		if err != nil {
			return "", err
		}
		jpegInfo = info
	default:
		return "", nil
	}
	return jpegInfo, nil
}
