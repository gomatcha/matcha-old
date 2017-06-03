package stacknavigator

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/stacknavigator"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/stacknavigator New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type StackView struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *StackView {
	if v, ok := ctx.Prev(key).(*StackView); ok {
		return v
	}
	return &StackView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *StackView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl1 := basicview.New(ctx, 1)
	chl1.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	opt1 := &stacknavigator.Options{}
	opt1.SetTitle("title1")

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Red}
	opt2 := &stacknavigator.Options{}
	opt2.SetTitle("title2")

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	opt3 := &stacknavigator.Options{}
	opt3.SetTitle("title3")

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Green}
	opt4 := &stacknavigator.Options{}
	opt4.SetTitle("title4")

	stacknav := stacknavigator.New(ctx, 100)
	stacknav.SetScreens([]stacknavigator.Screen{
		stacknavigator.Screen{
			View:    chl1,
			Options: opt1,
		},
		stacknavigator.Screen{
			View:    chl2,
			Options: opt2,
		},
		stacknavigator.Screen{
			View:    chl3,
			Options: opt3,
		},
		stacknavigator.Screen{
			View:    chl4,
			Options: opt4,
		},
	}, false)
	l.Add(stacknav, func(s *constraint.Solver) {
		s.WidthEqual(l.Width())
		s.HeightEqual(l.Height())
	})

	return &view.Model{
		Children: []view.View{stacknav},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
