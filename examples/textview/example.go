package textview

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/textinput"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/textview New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return New(ctx, "")
		}))
	})
}

type TextView struct {
	*view.Embed
}

func New(ctx *view.Context, key string) *TextView {
	if v, ok := ctx.Prev(key).(*TextView); ok {
		return v
	}
	return &TextView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TextView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl := textview.New(ctx, "a")
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
	chl2 := view.WithPainter(chl, &paint.Style{BackgroundColor: colornames.Blue})

	l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
	})

	input := textinput.New(ctx, "input")
	input.Text = text.New("blah")
	inputP := view.WithPainter(input, &paint.Style{BackgroundColor: colornames.Yellow})
	l.Add(inputP, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(200))
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
