package jpeg

import (
	"bytes"
	"fmt"
	"io"
)

var TIFFTags = &map[uint32]string{
	0x010e: "ImageDescription",
	0x010f: "Make",
	0x0110: "Model",
	0x0112: "Orientation",
	0x011a: "XResolution",
	0x011b: "YResolution",
	0x0128: "ResolutionUnit",
	0x0131: "Software",
	0x0132: "DateTime",
	0x013e: "WhitePoint",
	0x013f: "PrimaryChromaticities",
	0x0211: "YCbCrCoefficients",
	0x0213: "YCbCrPositioning",
	0x0214: "ReferenceBlackWhite",
	0x8298: "Copyright",
}

var SubIFDTags = &map[uint32]string{
	0x829a: "ExposureTime",
	0x829d: "FNumber",
	0x8822: "ExposureProgram",
	0x8827: "ISOSpeedRatings",
	0x9000: "ExifVersion",
	0x9003: "DateTimeOriginal",
	0x9004: "DateTimeDigitized",
	0x9101: "ComponentConfiguration",
	0x9102: "CompressedBitsPerPixel",
	0x9201: "ShutterSpeedValue",
	0x9202: "ApertureValue",
	0x9203: "BrightnessValue",
	0x9204: "ExposureBiasValue",
	0x9205: "MaxApertureValue",
	0x9206: "SubjectDistance",
	0x9207: "MeteringMode",
	0x9208: "LightSource",
	0x9209: "Flash",
	0x920a: "FocalLength",
	0x927c: "MakerNote",
	0x9286: "UserComment",
	0xa000: "FlashPixVersion",
	0xa001: "ColorSpace",
	0xa002: "ExifImageWidth",
	0xa003: "ExifImageHeight",
	0xa004: "RelatedSoundFile",
	0xa005: "ExifInteroperabilityOffset",
	0xa20e: "FocalPlaneXResolution",
	0xa20f: "FocalPlaneYResolution",
	0xa210: "FocalPlaneResolutionUnit",
	0xa217: "SensingMethod",
	0xa300: "FileSource",
	0xa301: "SceneType",
}

type APP1Data struct {
}

func parseExif(f io.Reader) (string, error) {
	//exifHdr := []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00}
	return "TODO", nil
}

func parseJfif(f io.Reader) (string, error) {
	return "TODO", nil
}

func Jpeg(f io.Reader, file string) (string, error) {
	magic := make([]byte, 2)
	f.Read(magic)
	if !bytes.Equal(magic, []byte{0xFF, 0xD8}) {
		return "", fmt.Errorf("%s: not a png file", file)
	}
	appMarker := make([]byte, 2)
	f.Read(appMarker)
	switch {
	case bytes.Equal(appMarker, []byte{0xff, 0xe1}):
		parseExif(f)
	case bytes.Equal(appMarker, []byte{0xff, 0xe1}):
		parseJfif(f)
	default:
		return "", nil
	}

	// comprobar app markers

	// ejemplo binario -> struct
	//	t := T{A: 0xEEFFEEFF, B: 3.14}
	//buf := &bytes.Buffer{}
	//err := binary.Write(buf, binary.BigEndian, t)
	//if err != nil {
	//    panic(err)
	//}
	//fmt.Println(buf.Bytes())
	return "TODO", nil
}
