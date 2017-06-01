package imageview

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/imageview New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type ImageView struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *ImageView {
	if v, ok := ctx.Prev(key).(*ImageView); ok {
		return v
	}
	return &ImageView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *ImageView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}

	chl := imageview.NewURLImageView(ctx, 0)
	chl.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	chl.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl.ResizeMode = imageview.ResizeModeFit
	m.Add(chl)
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return m
}
