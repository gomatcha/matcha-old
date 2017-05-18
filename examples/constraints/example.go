package constraints

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/constraints New", func() *view.ViewController {
		return view.NewViewController(func(c view.Config) view.View {
			return New(c)
		}, 0)
	})
}

type ConstraintsView struct {
	*view.Embed
}

func New(c view.Config) *ConstraintsView {
	v, ok := c.Prev.(*ConstraintsView)
	if !ok {
		v = &ConstraintsView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *ConstraintsView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}

	chl1 := basicview.New(ctx, 1)
	chl1.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	m.Add(chl1)
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	m.Add(chl2)
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	m.Add(chl3)
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Magenta}
	m.Add(chl4)
	_ = l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return m
}
