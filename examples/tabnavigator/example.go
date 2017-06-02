package tabnavigator

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/tabnavigator"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/tabnavigator New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type TabView struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *TabView {
	if v, ok := ctx.Prev(key).(*TabView); ok {
		return v
	}
	return &TabView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TabView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl1 := NewTabChildView(ctx, 1)
	chl1.Painter = &paint.Style{BackgroundColor: colornames.Red}

	chl2 := NewTabChildView(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Blue}

	chl3 := NewTabChildView(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Yellow}

	chl4 := NewTabChildView(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Green}

	tabnav := tabnavigator.New(ctx, 100)
	tabnav.SetViews([]view.View{chl1, chl2, chl3, chl4})
	l.Add(tabnav, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(500))
		s.HeightEqual(constraint.Const(500))
	})

	return &view.Model{
		Children: []view.View{tabnav},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type TabChildView struct {
	*view.Embed
	Painter paint.Painter
}

func NewTabChildView(ctx *view.Context, key interface{}) *TabChildView {
	if v, ok := ctx.Prev(key).(*TabChildView); ok {
		return v
	}
	return &TabChildView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TabChildView) Build(ctx *view.Context) *view.Model {
	return &view.Model{
		Painter: v.Painter,
	}
}
