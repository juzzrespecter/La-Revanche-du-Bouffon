package png

import (
	"bytes"
	"fmt"
	"io"
)

func Png(f io.Reader) (string, error) {
	magic := make([]byte, 8)
	f.Read(magic)
	if !bytes.Equal(magic, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "", fmt.Errorf("not a png file")
	}
	return "TODO", nil
}
