package constraints

import (
	"github.com/gomatcha/matcha/layout/constraint"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/view"
	"github.com/gomatcha/matcha/view/basicview"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/gomatcha/matcha/examples/constraints New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return New(ctx, "")
		}))
	})
}

type ConstraintsView struct {
	*view.Embed
}

func New(ctx *view.Context, key string) *ConstraintsView {
	if v, ok := ctx.Prev(key).(*ConstraintsView); ok {
		return v
	}
	return &ConstraintsView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *ConstraintsView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl1 := basicview.New(ctx, "1")
	chl1.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		// s.TopEqualC(0) // TODO(KD): Possible API? for doing const guides?
		// s.LeftEqualC(0)
		// s.WidthEqualC(100)
		// s.HeightEqualC(100)

		// s.TopEqual(100)
		// s.TopEqualAnchor(g1.Left())

		// s.TopEqual(s.Const(0)) // TODO(KD): Possible API? for doing const guides?
		// s.LeftEqual(s.Const(0))
		// s.WidthEqual(s.Const(100))
		// s.HeightEqual(s.Const(100))

		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2 := basicview.New(ctx, "2")
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		// s.TopEqualAnchor(g1.Bottom())
		// s.LeftEqualAnchor(g1.Left())
		// s.WidthEqual(300)
		// s.HeightEqual(300)

		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx, "3")
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx, "4")
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Magenta}
	_ = l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
