package pb

import (
	"bytes"
	"image"
	"image/color"

	"golang.org/x/image/bmp"
)

func ColorEncode(c color.Color) *Color {
	if c == nil {
		return nil
	}
	r, g, b, a := c.RGBA()
	return &Color{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}
}

func ImageEncode(img image.Image) *Image {
	buf := &bytes.Buffer{}
	err := bmp.Encode(buf, img)
	if err != nil {
		return nil
	}

	bounds := img.Bounds()
	return &Image{
		Width:  int64(bounds.Max.X - bounds.Min.X),
		Height: int64(bounds.Max.Y - bounds.Min.Y),
		Data:   buf.Bytes(),
	}
}
