package imageview

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/gomatcha/matcha/env"
	"github.com/gomatcha/matcha/layout/constraint"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/view"
	"github.com/gomatcha/matcha/view/imageview"
	"github.com/gomatcha/matcha/view/urlimageview"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/gomatcha/matcha/examples/imageview New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return New(ctx, "")
		}))
	})
}

type ImageView struct {
	*view.Embed
}

func New(ctx *view.Context, key string) *ImageView {
	if v, ok := ctx.Prev(key).(*ImageView); ok {
		return v
	}
	return &ImageView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *ImageView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl := urlimageview.New(ctx, "0")
	chl.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl.ResizeMode = imageview.ResizeModeStretch
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
		s.WidthLess(constraint.Const(200))
		s.HeightLess(constraint.Const(200))
	})

	chl2 := imageview.New(ctx, "1")
	chl2.Image = env.MustLoadImage("TableArrow")
	chl2.ResizeMode = imageview.ResizeModeFit
	l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(300))
		s.LeftEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
