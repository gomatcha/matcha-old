package imageview

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/pb"
	"github.com/overcyn/matcha/view"
)

type ResizeMode int

const (
	ResizeModeFit ResizeMode = iota
	ResizeModeFill
	ResizeModeStretch
	ResizeModeCenter
)

func (m ResizeMode) MarshalProtobuf() pb.ResizeMode {
	return pb.ResizeMode(m)
}

type layouter struct {
	bounds     image.Rectangle
	resizeMode ResizeMode
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rect{Max: ctx.MaxSize}}
	switch l.resizeMode {
	case ResizeModeFit:
		imgRatio := float64(l.bounds.Dx()) / float64(l.bounds.Dy())
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
		g.Frame.Max = layout.Pt(float64(l.bounds.Dx()), float64(l.bounds.Dy()))
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
	image      image.Image
	pbImage    *pb.Image
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
		v.pbImage = pb.ImageEncode(v.image)
	}

	// Default to Center if we don't have an image
	bounds := image.Rect(0, 0, 0, 0)
	resizeMode := ResizeModeCenter
	if v.image != nil {
		bounds = v.image.Bounds()
		resizeMode = v.ResizeMode
	}

	return &view.Model{
		Layouter:       &layouter{bounds: bounds, resizeMode: resizeMode},
		NativeViewName: "github.com/overcyn/matcha/view/imageview",
		NativeViewState: &pb.ImageView{
			Image:      v.pbImage,
			ResizeMode: v.ResizeMode.MarshalProtobuf(),
			Tint:       pb.ColorEncode(v.Tint),
		},
	}
}
