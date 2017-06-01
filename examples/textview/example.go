package textview

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/textview New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type TextView struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *TextView {
	if v, ok := ctx.Prev(key).(*TextView); ok {
		return v
	}
	return &TextView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TextView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}

	chl := textview.New(ctx, 5)
	chl.Painter = &paint.Style{BackgroundColor: colornames.Red}
	chl.String = "Subtitle"
	chl.Style.SetAlignment(text.AlignmentCenter)
	chl.Style.SetStrikethroughStyle(text.StrikethroughStyleDouble)
	chl.Style.SetStrikethroughColor(colornames.Blue)
	chl.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl.Style.SetUnderlineColor(colornames.Blue)
	chl.Style.SetTextColor(colornames.Yellow)
	chl.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Face:   "Bold",
		Size:   20,
	})
	m.Add(chl)
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return m
}
