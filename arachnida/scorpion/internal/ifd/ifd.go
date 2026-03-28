package ifd

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type IFDEntry struct {
	Tag uint16
	Fmt uint16
	N   uint32
	Val uint32
}

var fmtSize = map[uint16]uint32{
	1: 1, 2: 1, 3: 2, 4: 4, 5: 8, 6: 1,
	7: 1, 8: 2, 9: 4, 10: 8, 11: 4, 12: 4,
}

func (ifd *IFDEntry) getValue(base []byte) ([]byte, uint32) {
	size := fmtSize[ifd.Fmt] * ifd.N
	if size > 4 {
		return base[ifd.Val : ifd.Val+size], size
	}
	value := []byte{
		uint8(ifd.Val),
		uint8(ifd.Val) << 1,
		uint8(ifd.Val) << 2,
		uint8(ifd.Val) << 3,
	}
	return value, size
}

func (ifd *IFDEntry) FormatIFD(base []byte, order binary.ByteOrder) string {
	entry := fmt.Sprintf("%s:", IFDTags[ifd.Tag])

	value, size := ifd.getValue(base)
	valueArr := [][]byte{}
	for i := uint32(0); i < ifd.N; i++ {
		valueArr = append(valueArr, value[i*size:i*size+size])
	}
	fmt.Println(value, valueArr, size)
	switch ifd.Fmt {
	case 1, 3, 4: // unsigned numbers
		for _, x := range valueArr {
			y := order.Uint64(x)
			entry = entry + ", " + strconv.FormatUint(y, 10)
		}
	case 6, 8, 9: // signed numbers
		for _, x := range valueArr {
			y := int64(order.Uint64(x))
			entry = entry + ", " + strconv.FormatInt(y, 10)
		}
	case 2: // string
		entry = entry + string(value)
	case 5:
		for _, x := range valueArr {
			// [][]byte, cada []byte se debe transformar en dos unsigned logns
			n := order.Uint16(x[:2])
			d := order.Uint16(x[2:])
		}
	case 7:
		entry = entry + "undef."
	case 10: // signed rationa;
	case 11, 12: // floats
	}
	entry = entry + "\n"
	return entry
}
