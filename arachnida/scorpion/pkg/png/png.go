package png

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"scorpion/internal/ihdr"
)

// Byte order quilombo: https://www.w3.org/TR/png/#7Integers-and-byte-order
func Png(f io.Reader) (string, error) {
	magic := make([]byte, 8)
	f.Read(magic)
	if !bytes.Equal(magic, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "", fmt.Errorf("not a png file")
	}
	img, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	if !bytes.Equal(img[4:8], []byte{0x49, 0x48, 0x44, 0x52}) {
		return "", fmt.Errorf("Invalid first IHDR chunk")
	}
	ihdrChunk := &ihdr.IHDRChunk{
		Width:             binary.BigEndian.Uint32(img[8:12]),
		Height:            binary.BigEndian.Uint32(img[12:16]),
		BitDepth:          uint8(img[16]),
		ColorType:         uint8(img[17]),
		CompressionMethod: uint8(img[18]),
		FilterMethod:      uint8(img[19]),
		InterlaceMethod:   uint8(img[20]),
	}
	metadata := ihdrChunk.GetMetadata()
	offset := 25

	for {
		if len(img) < offset+8 {
			return metadata, nil
		}
		chunkLength := binary.BigEndian.Uint32(img[offset : offset+4])
		chunkTag := string(img[offset+4 : offset+8])
		switch chunkTag {
		case "IEND":
			return metadata, nil
		case "iTXt", "tEXt", "zTXt":
			text := img[offset+8 : offset+8+int(chunkLength)]
			textArray := bytes.Split(text, []byte{0})
			key := string(textArray[0])
			val := string(bytes.Trim(textArray[1], string([]byte{0})))
			metadata = metadata + fmt.Sprintf("%-31s%s\n", key, val)
			offset = offset + 12 + int(chunkLength)
		default:
			offset = offset + 12 + int(chunkLength)
		}
	}
}
