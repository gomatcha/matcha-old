package pb

import (
	"image/color"
)

func ColorEncode(c color.Color) *Color {
	r, g, b, a := c.RGBA()
	return &Color{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}
}
