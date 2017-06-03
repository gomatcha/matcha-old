package tabnavigator

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/tabnav"
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

	chl1 := basicview.New(ctx, 1)
	chl1.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	screen1 := &tabnav.Screen{}
	screen1.SetTitle("title1")
	screen1.SetView(chl1)

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Red}
	screen2 := &tabnav.Screen{}
	screen2.SetTitle("title2")
	screen2.SetView(chl2)

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	screen3 := &tabnav.Screen{}
	screen3.SetTitle("title3")
	screen3.SetView(chl3)

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Green}
	screen4 := &tabnav.Screen{}
	screen4.SetTitle("title4")
	screen4.SetView(chl4)

	tab := tabnav.New(ctx, 100)
	tab.SetScreens([]*tabnav.Screen{screen1, screen2, screen3, screen4})
	l.Add(tab, func(s *constraint.Solver) {
		s.WidthEqual(l.Width())
		s.HeightEqual(l.Height())
	})

	return &view.Model{
		Children: []view.View{tab},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
