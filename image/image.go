package text

import (
	"github.com/overcyn/mochi"
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

type ImageView struct {
	marker       mochi.Marker
	PaintOptions mochi.PaintOptions
	ImageBytes   []byte
}

func NewImageView(p interface{}) *ImageView {
	v, ok := p.(*ImageView)
	if !ok {
		v = &ImageView{}
	}
	return v
}

func (v *ImageView) Mount(m mochi.Marker) {
	v.marker = m
}

func (v *ImageView) Update(p *mochi.Node) *mochi.Node {
	n := mochi.NewNode()
	n.PaintOptions = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi ImageView"
	n.Bridge.State = v.ImageBytes
	return n
}

func (v *ImageView) Unmount() {
	v.marker = nil
}
