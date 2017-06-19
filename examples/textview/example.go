package textview

import (
	"fmt"

	"github.com/overcyn/matcha/keyboard"
	"github.com/overcyn/matcha/layout/constraint"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/text"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/textinput"
	"github.com/overcyn/matcha/view/textview"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/overcyn/matcha/examples/textview New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return New(ctx, "")
		}))
	})
}

type TextView struct {
	*view.Embed
	text      *text.Text
	responder *keyboard.Responder
}

func New(ctx *view.Context, key string) *TextView {
	if v, ok := ctx.Prev(key).(*TextView); ok {
		return v
	}
	return &TextView{
		Embed:     view.NewEmbed(ctx.NewId(key)),
		text:      text.New("blah"),
		responder: &keyboard.Responder{},
	}
}

func (v *TextView) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageVisible) {
		v.responder.Show()
		fmt.Println("show", v.responder.Visible())
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
	chlP := view.WithPainter(chl, &paint.Style{BackgroundColor: colornames.Blue})
	chlG := l.Add(chlP, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
	})

	input := textinput.New(ctx, "input")
	input.Text = v.text
	input.Responder = v.responder
	input.OnChange = func(input *textinput.View) {
		v.Update()
	}
	inputP := view.WithPainter(input, &paint.Style{BackgroundColor: colornames.Yellow})
	l.Add(inputP, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(200))
		s.LeftEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(100))
	})

	reverse := textview.New(ctx, "reverse")
	reverse.String = Reverse(v.text.String())
	l.Add(reverse, func(s *constraint.Solver) {
		s.TopEqual(chlG.Bottom())
		s.LeftEqual(chlG.Left())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
