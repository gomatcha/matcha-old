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

const (
	urlImageViewId int = iota
)

type URLImageView struct {
	marker       mochi.Marker
	PaintOptions mochi.PaintOptions
	URL          string
}

func NewURLImageView(p interface{}) *URLImageView {
	v, ok := p.(*URLImageView)
	if !ok {
		v = &URLImageView{}
	}
	return v
}

func (v *URLImageView) Mount(m mochi.Marker) {
	v.marker = m
}

func (v *URLImageView) Update(p *mochi.Node) *mochi.Node {
	n := mochi.NewNode()
	n.PaintOptions = v.PaintOptions

	chl := NewImageView(p.Get(urlImageViewId))
	chl.PaintOptions.BackgroundColor = mochi.RedColor
	n.Set(urlImageViewId, chl)

	return n
}

func (v *URLImageView) Unmount() {
	v.marker = nil
}

// ImageView

type ImageView struct {
	PaintOptions mochi.PaintOptions
	Image        image.Image
	image        image.Image
	bytes        []byte
}

func NewImageView(p interface{}) *ImageView {
	v, ok := p.(*ImageView)
	if !ok {
		v = &ImageView{}
	}
	return v
}

func (v *ImageView) Mount(m mochi.Marker) {
}

func (v *ImageView) Update(p *mochi.Node) *mochi.Node {
	n := mochi.NewNode()

	if v.Image != v.image {
		v.image = v.Image

		buf := &bytes.Buffer{}
		err := bmp.Encode(buf, v.image)
		if err != nil {
			fmt.Println("ImageView encoding error:", err)
		}
		v.bytes = buf.Bytes()
	}

	n.PaintOptions = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi ImageView"
	n.Bridge.State = v.bytes
	return n
}

func (v *ImageView) Unmount() {
}
