package ihdr

import "fmt"

type IHDRChunk struct {
	Width             uint32
	Height            uint32
	BitDepth          uint8
	ColorType         uint8
	CompressionMethod uint8
	FilterMethod      uint8
	InterlaceMethod   uint8
}

var colorTypeValues = map[uint8]string{
	0: "Grayscale",
	2: "Red, green and blue",
	3: "Indexed",
	4: "Grayscale and alpha",
	6: "Red, green ,blue and alpha",
}

var interlaceValues = map[uint8]string{
	0: "Noninterlaced",
	1: "Adam7 interlace",
}

func (chunk *IHDRChunk) GetMetadata() string {
	ctv := colorTypeValues[chunk.ColorType]
	iv := interlaceValues[chunk.InterlaceMethod]

	return fmt.Sprintf("Width:                         %d\n"+
		"Height:                        %d\n"+
		"Bit depth:                     %d\n"+
		"Color type:                    %s\n"+
		"Compression method:            Inflate/Deflate\n"+
		"Filter method:                 Adaptive\n"+
		"Interlace:                     %s\n",
		int32(chunk.Width),
		int32(chunk.Height),
		chunk.BitDepth,
		ctv,
		iv,
	)
}
