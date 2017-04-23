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
// 	marker       mochi.Marker
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
	marker       mochi.Marker
	PaintOptions mochi.PaintOptions
	Image        image.Image
	image        image.Image
	bytes        []byte
}

func NewImageView(cfg mochi.Config) *ImageView {
	v, ok := cfg.Prev.(*ImageView)
	if !ok {
		v = &ImageView{}
		v.marker = cfg.Marker
	}
	return v
}

func (v *ImageView) Update(ctx *mochi.ViewContext) *mochi.Node {
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
