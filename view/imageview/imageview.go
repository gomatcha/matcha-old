package imageview

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gomatcha/matcha"
	"github.com/gomatcha/matcha/comm"
	"github.com/gomatcha/matcha/env"
	"github.com/gomatcha/matcha/layout"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/pb"
	"github.com/gomatcha/matcha/pb/view/imageview"
	"github.com/gomatcha/matcha/view"
)

type ResizeMode int

const (
	ResizeModeFit ResizeMode = iota
	ResizeModeFill
	ResizeModeStretch
	ResizeModeCenter
)

func (m ResizeMode) MarshalProtobuf() imageview.ResizeMode {
	return imageview.ResizeMode(m)
}

type layouter struct {
	bounds     image.Rectangle
	scale      float64
	resizeMode ResizeMode
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rect{Max: ctx.MaxSize}}
	switch l.resizeMode {
	case ResizeModeFit:
		imgRatio := float64(l.bounds.Dx()) / l.scale / float64(l.bounds.Dy()) / l.scale
		maxRatio := ctx.MaxSize.X / ctx.MaxSize.Y
		if imgRatio > maxRatio {
			g.Frame.Max = layout.Pt(ctx.MaxSize.X, ctx.MaxSize.X/imgRatio)
		} else {
			g.Frame.Max = layout.Pt(ctx.MaxSize.Y/imgRatio, ctx.MaxSize.Y)
		}
	case ResizeModeFill:
		fallthrough
	case ResizeModeStretch:
		g.Frame.Max = ctx.MaxSize
	case ResizeModeCenter:
		g.Frame.Max = layout.Pt(float64(l.bounds.Dx())/l.scale, float64(l.bounds.Dy())/l.scale)
	}
	return g, nil
}

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}

type View struct {
	*view.Embed
	Image      image.Image
	ResizeMode ResizeMode
	Tint       color.Color
	PaintStyle *paint.Style
	image      image.Image
	pbImage    *pb.ImageOrResource
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	if v.Image != v.image {
		v.image = v.Image
		v.pbImage = env.ImageMarshalProtobuf(v.image)
	}

	// Default to Center if we don't have an image
	bounds := image.Rect(0, 0, 0, 0)
	resizeMode := ResizeModeCenter
	scale := 1.0
	if v.image != nil {
		bounds = v.image.Bounds()
		resizeMode = v.ResizeMode

		if res, ok := v.image.(*env.ImageResource); ok {
			scale = res.Scale()
		}
	}

	var painter paint.Painter
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		Layouter:       &layouter{bounds: bounds, resizeMode: resizeMode, scale: scale},
		NativeViewName: "github.com/gomatcha/matcha/view/imageview",
		NativeViewState: &imageview.View{
			Image:      v.pbImage,
			Scale:      scale,
			ResizeMode: v.ResizeMode.MarshalProtobuf(),
			Tint:       pb.ColorEncode(v.Tint),
		},
	}
}
