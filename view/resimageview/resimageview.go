package resimageview

import (
	"image"
	"image/color"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/env"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
	pbenv "github.com/overcyn/mochi/pb/env"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/imageview"
)

type layouter struct {
	bounds     image.Rectangle
	resizeMode imageview.ResizeMode
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rect{Max: ctx.MaxSize}}
	switch l.resizeMode {
	case imageview.ResizeModeFit:
		imgRatio := float64(l.bounds.Dx()) / float64(l.bounds.Dy())
		maxRatio := ctx.MaxSize.X / ctx.MaxSize.Y
		if imgRatio > maxRatio {
			g.Frame.Max = layout.Pt(ctx.MaxSize.X, ctx.MaxSize.X/imgRatio)
		} else {
			g.Frame.Max = layout.Pt(ctx.MaxSize.Y/imgRatio, ctx.MaxSize.Y)
		}
	case imageview.ResizeModeFill:
		fallthrough
	case imageview.ResizeModeStretch:
		g.Frame.Max = ctx.MaxSize
	case imageview.ResizeModeCenter:
		g.Frame.Max = layout.Pt(float64(l.bounds.Dx()), float64(l.bounds.Dy()))
	}
	return g, nil
}

func (l *layouter) Notify() chan struct{} {
	return nil // no-op
}

func (l *layouter) Unnotify(chan struct{}) {
	// no-op
}

type View struct {
	*view.Embed
	Painter    paint.Painter
	Resource   *env.Resource
	ResizeMode imageview.ResizeMode
	Tint       color.Color
}

func New(ctx *view.Context, key interface{}) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	// Default to Center if we don't have an image
	bounds := image.Rect(0, 0, 0, 0)
	resizeMode := imageview.ResizeModeCenter
	var resPb *pbenv.Resource = nil
	if v.Resource != nil {
		size := v.Resource.Size()
		bounds = image.Rect(0, 0, int(size.X), int(size.Y))
		resizeMode = v.ResizeMode
		resPb = v.Resource.MarshalProtobuf()
	}

	return &view.Model{
		Layouter:       &layouter{bounds: bounds, resizeMode: resizeMode},
		Painter:        v.Painter,
		NativeViewName: "github.com/overcyn/mochi/view/imageview",
		NativeViewState: &pb.ImageView{
			Resource:   resPb,
			ResizeMode: v.ResizeMode.MarshalProtobuf(),
			Tint:       pb.ColorEncode(v.Tint),
		},
	}
}
