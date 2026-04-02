package ifd

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf16"
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
			n := float64(order.Uint32(x[:4]))
			d := float64(order.Uint32(x[4:]))
			valueString = append(valueString, strconv.FormatFloat(n/d, 'f', -1, 64))
		}
	case 7: // undefined
		switch ifd.Tag {
		case 0x927c: // MakerNote, unprintable
			break
		case 0x9000, 0xA000:
			fmt.Print("Tags fancy: ")
			valueString = append(valueString, string(valueArr[0]))
		case 0x9101:
			for _, x := range valueArr {
				valueString = append(valueString, fmt.Sprintf("%v", x))
			}
		case 0x9286:
			comment := []byte{}
			for _, x := range valueArr {
				comment = append(comment, x...)
			}
			if len(comment) > 8 {
				prefix := comment[:8]
				content := comment[8:]
				switch string(prefix) {
				case "ASCII\x00\x00\x00":
					valueString = append(valueString, string(content))
				case "UNICODE\x00":
					unicode16 := make([]uint16, len(comment)/2)
					for i := range len(comment) / 2 {
						order.PutUint16(content[i*2:i*2+2], unicode16[i])
					}
					valueString = append(valueString, string(utf16.Decode(unicode16)))
				default:
					valueString = append(valueString, hex.EncodeToString(comment))
				}
			} else {
				valueString = []string{""}
			}
		default:
			for _, x := range valueArr {
				isPrintable := func(n []byte) bool {
					for _, c := range n {
						if c > 126 || c < 32 {
							return false
						}
					}
					return true
				}
				if isPrintable(x) {
					valueString = append(valueString, string(x))
				} else {
					valueString = append(valueString, hex.EncodeToString(x))
				}
			}
		}

		valueString = append(valueString, "[UNDEF]") // borrame
	case 10: // signed rational
		for _, x := range valueArr {
			n := float64(int32(order.Uint32(x[:4])))
			d := float64(int32(order.Uint32(x[4:])))

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
