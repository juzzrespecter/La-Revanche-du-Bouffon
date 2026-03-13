package utils

import "errors"

var ErrArgParse = errors.New("error: one must be provided: -g, -k")
var ErrKeyLength = errors.New("key: string too short (must be >64 characters)")
var ErrKeyEncode = errors.New("key: must be encoded in base64")
