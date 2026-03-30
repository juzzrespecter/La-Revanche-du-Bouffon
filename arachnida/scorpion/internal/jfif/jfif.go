package jfif

import "fmt"

type JFIFApp0 struct {
	Length      uint16
	Id          []byte
	Version     uint16
	DensityUnit uint8
	Xdensity    uint16
	Ydensity    uint16
	Xthumbnail  uint8
	Ythumbnail  uint8
}

var densityUnits = map[uint8]string{
	0: "",
	1: "pixels per inch",
	2: "pixels per cm",
}

func (i *JFIFApp0) GetMetadata() string {
	units, ok := densityUnits[i.DensityUnit]
	if !ok {
		units = "(unknown)"
	}
	maj, min := uint8(i.Version), uint8(i.Version<<1)
	return fmt.Sprintf(
		"JFIF version:                  %d.0%d\n"+
			"X Density:                     %d %s\n"+
			"Y Density:                     %d %s\n"+
			"X Thumbnail:                   %d %s\n"+
			"Y Thumbnail:                   %d %s\n",
		maj, min,
		i.Xdensity, units,
		i.Ydensity, units,
		i.Xthumbnail, units,
		i.Ythumbnail, units)
}

func (i *JFIFApp0) GetThumbnailLength() int {
	return int(i.Xthumbnail) * int(i.Ythumbnail) * 3
}
