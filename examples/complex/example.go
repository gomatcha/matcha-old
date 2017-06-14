package example

import (
	"fmt"
	"time"

	"golang.org/x/image/colornames"

	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/button"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/switchview"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochi/view/urlimageview"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/complex New", func() *view.Root {
		return view.NewRoot(NewNestedView(nil, nil))
	})
}

type TableView struct {
	*view.Embed
}

func NewTableView(ctx *view.Context, key interface{}) *TableView {
	if v, ok := ctx.Prev(key).(*TableView); ok {
		return v
	}
	return &TableView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	childLayouter := &table.Layout{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, i+1000)
		childView.String = "TEST TEST"
		// childView.OnClick = func() {
		// 	v.Lock()
		// 	defer v.Unlock()

		// 	child := NewTableView(nil, nil)
		// 	v.Nav.StackNav().Push(child.StackScreen())
		// }

		childLayouter.Add(childView)
	}

	content := basicview.New(ctx, 9)
	content.Painter = &paint.Style{BackgroundColor: colornames.White}
	content.Layouter = childLayouter
	content.Children = childLayouter.Views()

	scroll := scrollview.New(ctx, 10)
	scroll.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	scroll.ContentView = content
	l.Add(scroll, func(s *constraint.Solver) {
		s.WidthEqual(l.Width())
		s.HeightEqual(l.Height())
	})

	return &view.Model{
		Children: []view.View{scroll},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type NestedView struct {
	*view.Embed
	counter     int
	ticker      *animate.Ticker
	floatTicker comm.Float64Notifier
	colorTicker comm.ColorNotifier
}

func NewNestedView(ctx *view.Context, key interface{}) *NestedView {
	if v, ok := ctx.Prev(key).(*NestedView); ok {
		return v
	}
	ticker := animate.NewTicker(time.Second * 5)
	floatTicker := animate.FloatInterpolate(ticker, animate.FloatLerp{Start: 0, End: 150})
	return &NestedView{
		Embed:       view.NewEmbed(ctx.NewId(key)),
		ticker:      ticker,
		floatTicker: floatTicker,
		colorTicker: animate.ColorInterpolate(ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow}),
	}
}

func (v *NestedView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl1 := basicview.New(ctx, 1)
	chl1.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Notifier(v.floatTicker))
		s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Magenta}
	g4 := l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(nil, nil) // test no context
	chl5.String = "Subtitle"
	chl5.Style.SetAlignment(text.AlignmentCenter)
	chl5.Style.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl5.Style.SetStrikethroughColor(colornames.Magenta)
	chl5.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl5.Style.SetUnderlineColor(colornames.Green)
	chl5.Style.SetFont(text.Font{
		Family: "American Typewriter",
		Face:   "Bold",
		Size:   20,
	})
	chl5p := view.WithPainter(chl5, &paint.Style{BackgroundColor: colornames.Cyan})

	g5 := l.Add(chl5p, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl6 := textview.New(ctx, 6)
	chl6.String = fmt.Sprintf("Counter: %v", v.counter)
	chl6.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	chl6p := view.WithPainter(chl6, &paint.Style{BackgroundColor: colornames.Red})
	g6 := l.Add(chl6p, func(s *constraint.Solver) {
		s.BottomEqual(g5.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl8 := button.New(ctx, 8)
	chl8.Text = "Button"
	chl8.OnPress = func() {
		fmt.Println("On Click")
		v.counter += 1
		v.Update()
	}
	g8 := l.Add(chl8, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	if v.counter%2 == 0 {
		chl9 := urlimageview.New(ctx, 7)
		chl9.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
		chl9.ResizeMode = imageview.ResizeModeFit
		pChl9 := view.WithPainter(chl9, &paint.Style{BackgroundColor: colornames.Cyan})

		_ = l.Add(pChl9, func(s *constraint.Solver) {
			s.BottomEqual(g8.Top())
			s.RightEqual(g2.Right().Add(-15))
			s.WidthEqual(constraint.Const(200))
			s.HeightEqual(constraint.Const(200))
		})
	}
	chl11 := switchview.New(ctx, 12)
	chl11.OnValueChange = func(a *switchview.View) {
		fmt.Println("switch tapped", a.Value)
	}
	_ = l.Add(chl11, func(s *constraint.Solver) {
		s.LeftEqual(g6.Right())
		s.TopEqual(g6.Top())
	})

	childLayouter := &table.Layout{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, i+1000)
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: colornames.Red}
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx, 9)
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childLayouter.Views()

	chl10 := scrollview.New(ctx, 10)
	chl10.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	chl10.ContentView = scrollChild
	_ = l.Add(chl10, func(s *constraint.Solver) {
		s.TopEqual(g4.Bottom())
		s.LeftEqual(g4.Left())
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type TableCell struct {
	*view.Embed
	String  string
	Painter paint.Painter
}

func NewTableCell(ctx *view.Context, key interface{}) *TableCell {
	if v, ok := ctx.Prev(key).(*TableCell); ok {
		return v
	}
	return &TableCell{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx, 1)
	textView.String = v.String
	textView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	l.Add(textView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(10))
		s.RightEqual(l.Right().Add(-10))
		s.CenterYEqual(l.CenterY())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  v.Painter,
	}
}
