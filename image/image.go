package image

import (
	"bytes"
	"fmt"
	"github.com/overcyn/mochi"
	"golang.org/x/image/bmp"
	"image"
	// "image/color"
	// "mochi/bridge"
)

// const (
// 	urlImageViewId int = iota
// )

// type URLImageView struct {
// 	marker       mochi.Updater
// 	PaintOptions mochi.PaintOptions
// 	URL          string
// }

// func NewURLImageView(p interface{}) *URLImageView {
// 	v, ok := p.(*URLImageView)
// 	if !ok {
// 		v = &URLImageView{}
// 	}
// 	return v
// }

// func (v *URLImageView) Update(p *mochi.Node) *mochi.Node {
// 	n := mochi.NewNode()
// 	n.Painter = &mochi.BasicPainter{v.PaintOptions}

// 	chl := NewImageView(p.Get(urlImageViewId))
// 	chl.PaintOptions.BackgroundColor = mochi.RedColor
// 	n.Set(urlImageViewId, chl)

// 	return n
// }

// ImageView

type ImageView struct {
	*mochi.Embed
	PaintOptions mochi.PaintOptions
	Image        image.Image
	image        image.Image
	bytes        []byte
}

func NewImageView(c mochi.Config) *ImageView {
	v, ok := c.Prev.(*ImageView)
	if !ok {
		v = &ImageView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *ImageView) Build(ctx *mochi.BuildContext) *mochi.Node {
	n := &mochi.Node{}

	if v.Image != v.image {
		v.image = v.Image

		buf := &bytes.Buffer{}
		err := bmp.Encode(buf, v.image)
		if err != nil {
			fmt.Println("ImageView encoding error:", err)
		}
		v.bytes = buf.Bytes()
	}

	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi ImageView"
	n.Bridge.State = v.bytes
	return n
}
