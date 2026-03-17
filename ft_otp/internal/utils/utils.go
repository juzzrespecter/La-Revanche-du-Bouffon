package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrArgParse = errors.New("error: one must be provided: -g, -k")
var ErrKeyLength = errors.New("key: string too short (must be >64 characters)")
var ErrKeyEncode = errors.New("key: must be encoded in base16")
var ErrMismatchKey = errors.New("")

func ReadInput() ([]byte, error) {
	rd := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce key")
	key, err := rd.ReadBytes(byte('\n'))
	if err != nil {
		return nil, err
	}
	return key, nil
}
