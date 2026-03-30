package ifd

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
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

func (ifd *IFDEntry) getValue(base []byte) []byte {
	size := fmtSize[ifd.Fmt] * ifd.N
	if size > 4 {
		return base[ifd.Val : ifd.Val+size]
	}
	value := []byte{
		uint8(ifd.Val),
		uint8(ifd.Val) << 1,
		uint8(ifd.Val) << 2,
		uint8(ifd.Val) << 3,
	}
	return value
}

func (ifd *IFDEntry) FormatIFD(base []byte, order binary.ByteOrder) string {
	value := ifd.getValue(base)
	valueArr := [][]byte{}
	e := fmtSize[ifd.Fmt]
	for i := uint32(0); i < ifd.N; i++ {
		valueArr = append(valueArr, value[i*e:i*e+e])
	}
	//fmt.Printf("IFDEntry getValues: \nNumber of values: %d\nValue array: %v (length: %d)\nArray of values (offset): %v\nActual size of value: %d\n\n",
	//	ifd.N, value, len(value), valueArr, size,
	//)
	var valueString []string
	switch ifd.Fmt {
	case 1, 3, 4: // unsigned numbers
		for _, x := range valueArr {
			y := order.Uint16(x)
			valueString = append(valueString, strconv.FormatUint(uint64(y), 10))
		}
	case 6, 8, 9: // signed numbers
		for _, x := range valueArr {
			y := int64(order.Uint64(x))
			valueString = append(valueString, strconv.FormatInt(y, 10))
		}
	case 2: // string
		valueString = append(valueString, string(value))
	case 5: // unsigned rational
		for _, x := range valueArr {
			n := float64(order.Uint16(x[:4]))
			d := float64(order.Uint16(x[4:]))
			valueString = append(valueString, strconv.FormatFloat(n/d, 'f', -1, 64))
		}
	case 7:
		valueString = append(valueString, "undef.")
	case 10: // signed rational
		for _, x := range valueArr {
			n := float64(int16(order.Uint16(x[:4])))
			d := float64(int16(order.Uint16(x[4:])))

			valueString = append(valueString, strconv.FormatFloat(n/d, 'f', -1, 64))
		}
	case 11, 12: // floats
		for _, x := range valueArr {
			bits := binary.LittleEndian.Uint64(x)
			float := math.Float64frombits(bits)
			valueString = append(valueString, strconv.FormatFloat(float, 'f', -1, 64))
		}

	}

	return fmt.Sprintf("%-*s%s\n", 31, IFDTags[ifd.Tag]+":", strings.Join(valueString, ","))
}
