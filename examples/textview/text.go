package textview

import (
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/textview New", func() *view.ViewController {
		return view.NewViewController(func(c view.Config) view.View {
			return New(c)
		}, 0)
	})
}

type TextView struct {
	*view.Embed
}

func New(c view.Config) *TextView {
	v, ok := c.Prev.(*TextView)
	if !ok {
		v = &TextView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TextView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: internal.GreenColor},
	}

	chl := textview.New(ctx.Get(5))
	chl.Painter = &paint.Style{BackgroundColor: internal.RedColor}
	chl.String = "Subtitle"
	chl.Style.SetAlignment(text.AlignmentCenter)
	chl.Style.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl.Style.SetStrikethroughColor(internal.MagentaColor)
	chl.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl.Style.SetUnderlineColor(internal.GreenColor)
	chl.Style.SetFont(text.Font{
		Family: "American Typewriter",
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
