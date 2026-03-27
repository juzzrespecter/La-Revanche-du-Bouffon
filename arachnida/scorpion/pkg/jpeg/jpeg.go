package jpeg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

type IFDEntry struct {
	Tag uint16
	Fmt uint16
	N   uint32
	Val uint32
}

func formatIFD(ifd *IFDEntry) string {
	fmtSize := map[uint16]uint32{
		1: 1, 2: 1, 3: 2, 4: 4, 5: 8,
	}
	x := fmt.Sprintf(
		"Tag: %d\n"+
			"Fmt: %d\n"+
			"N  : %d\n"+
			"Val: %d\n"+
			"Cos: %s\n",
		ifd.Tag, ifd.Fmt, ifd.N, ifd.Val, IFDTags[ifd.Tag],
	)
	size := ifd.N * fmtSize[ifd.Fmt]
	var value string
	if size <= 4 {
		switch ifd.Fmt {
		case 1, 3, 4:
			value = strconv.FormatUint(uint64(ifd.Val), 10)
		case 2:
			value = string(ifd.Val)
		}
	} else {
		value = "Offset"
	}

	// format
	fmt.Printf("%s      : %s\n", IFDTags[ifd.Tag], value)
	return x
}

func parseExif(f io.Reader) (string, error) {
	buffer := make([]byte, 8)
	err := binary.Read(f, binary.LittleEndian, buffer)
	if err != nil {
		return "", err
	}
	exifHdr := buffer[2:]
	if !bytes.Equal(exifHdr, []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00}) {
		fmt.Println("Not an exif header, silently exit") // borrame
		return "", nil
	}
	tiffHdr := make([]byte, 8)
	err = binary.Read(f, binary.LittleEndian, tiffHdr)
	if err != nil {
		return "", err
	}
	var byteOrder binary.ByteOrder
	switch string(tiffHdr[:2]) {
	case "MM":
		byteOrder = binary.BigEndian
	case "II":
		byteOrder = binary.LittleEndian
	default:
		return "", fmt.Errorf("invalid byte order")
	}
	if byteOrder.Uint16(tiffHdr[2:4]) != 42 {
		return "", fmt.Errorf("invalid TIFF magic number")
	}
	offset := byteOrder.Uint32(tiffHdr[4:]) - 8
	if offset != 0 {
		if _, err := io.CopyN(io.Discard, f, int64(offset)); err != nil {
			return "", fmt.Errorf("offset error: %w", err)
		}
	}
	for {
		var n int16
		if err := binary.Read(f, byteOrder, &n); err != nil {
			return "", err
		}
		for i := 0; i < int(n); i++ {
			entryBuffer := make([]byte, 16)
			if err := binary.Read(f, byteOrder, entryBuffer); err != nil {
				return "", err
			}
			entry := &IFDEntry{
				Tag: byteOrder.Uint16(entryBuffer[:2]),
				Fmt: byteOrder.Uint16(entryBuffer[2:4]),
				N:   byteOrder.Uint32(entryBuffer[4:8]),
				Val: byteOrder.Uint32(entryBuffer[8:]),
			}
			formatIFD(entry)
		}
		var offset uint32
		if err := binary.Read(f, byteOrder, &offset); err != nil {
			return "", err
		}
		if offset == 0 {
			break
		}
		fmt.Println("offset ", offset)
		offset = offset - 8
		if _, err := io.CopyN(io.Discard, f, int64(offset)); err != nil {
			return "", fmt.Errorf("offset error: %w, %d", err, offset)
		}
	}

	return "TODO", nil
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
