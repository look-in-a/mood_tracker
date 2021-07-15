package graphics

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	"mood/moods"
)

func createIcon(c color.NRGBA) []byte {
	width := 200
	height := 200

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}

	buf := []byte{}
	f := bytes.NewBuffer(buf)
	png.Encode(f, img)

	return []byte(f.String())
}

func createIcons(m moods.Mood) (on, off []byte) {
	return createIcon(m.ColorON), createIcon(m.ColorOFF)
}
