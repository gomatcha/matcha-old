package stacknavigator

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	stacknav "github.com/overcyn/mochi/view/stacknav"
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
	screen1 := &stacknav.Screen{}
	screen1.SetTitle("title1")
	screen1.SetView(chl1)

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Red}
	screen2 := &stacknav.Screen{}
	screen2.SetTitle("title2")
	screen2.SetView(chl2)

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	screen3 := &stacknav.Screen{}
	screen3.SetTitle("title3")
	screen3.SetView(chl3)

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Green}
	screen4 := &stacknav.Screen{}
	screen4.SetTitle("title4")
	screen4.SetView(chl4)

	stack := stacknav.New(ctx, 100)
	stack.SetScreens([]*stacknav.Screen{screen1, screen2, screen3, screen4}, false)
	l.Add(stack, func(s *constraint.Solver) {
		s.WidthEqual(l.Width())
		s.HeightEqual(l.Height())
	})

	return &view.Model{
		Children: []view.View{stack},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
