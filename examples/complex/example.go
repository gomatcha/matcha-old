package example

import (
	"fmt"
	"time"

	"golang.org/x/image/colornames"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/button"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/textview"
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
	childViews := []view.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, i+1000)
		childView.String = "TEST TEST"
		// childView.OnClick = func() {
		// 	v.Lock()
		// 	defer v.Unlock()

		// 	child := NewTableView(nil, nil)
		// 	v.Nav.StackNav().Push(child.StackScreen())
		// }

		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	content := basicview.New(ctx, 9)
	content.Painter = &paint.Style{BackgroundColor: colornames.White}
	content.Layouter = childLayouter
	content.Children = childViews

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
	floatTicker mochi.Float64Notifier
	colorTicker mochi.ColorNotifier
}

func NewNestedView(ctx *view.Context, key interface{}) *NestedView {
	if v, ok := ctx.Prev(key).(*NestedView); ok {
		return v
	}
	ticker := animate.NewTicker(time.Second * 5)
	return &NestedView{
		Embed:       view.NewEmbed(ctx.NewId(key)),
		ticker:      ticker,
		floatTicker: animate.FloatInterpolate(ticker, animate.FloatLerp{Start: 0, End: 150}),
		colorTicker: animate.ColorInterpolate(ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow}),
	}
}

func (v *NestedView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chls := []view.View{}
	chl1 := basicview.New(ctx, 1)
	chl1.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	chls = append(chls, chl1)
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Notifier(v.floatTicker))
		s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	chl2 := basicview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	chls = append(chls, chl2)
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx, 3)
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	chls = append(chls, chl3)
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Magenta}
	chls = append(chls, chl4)
	g4 := l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(nil, nil) // test no context
	chl5.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
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
	chls = append(chls, chl5)
	g5 := l.Add(chl5, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl6 := textview.New(ctx, 6)
	chl6.Painter = &paint.Style{BackgroundColor: colornames.Red}
	chl6.String = fmt.Sprintf("Counter: %v", v.counter)
	chl6.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	chls = append(chls, chl6)
	g6 := l.Add(chl6, func(s *constraint.Solver) {
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
	chls = append(chls, chl8)
	g8 := l.Add(chl8, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	if v.counter%2 == 0 {
		chl9 := imageview.NewURLImageView(ctx, 7)
		chl9.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
		chl9.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
		chl9.ResizeMode = imageview.ResizeModeFit
		chls = append(chls, chl9)
		_ = l.Add(chl9, func(s *constraint.Solver) {
			s.BottomEqual(g8.Top())
			s.RightEqual(g2.Right().Add(-15))
			s.WidthEqual(constraint.Const(200))
			s.HeightEqual(constraint.Const(200))
		})
	}

	childLayouter := &table.Layout{}
	childViews := []view.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, i+1000)
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: colornames.Red}
		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx, 9)
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childViews

	chl10 := scrollview.New(ctx, 10)
	chl10.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	chl10.ContentView = scrollChild
	chls = append(chls, chl10)
	_ = l.Add(chl10, func(s *constraint.Solver) {
		s.TopEqual(g4.Bottom())
		s.LeftEqual(g4.Left())
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	return &view.Model{
		Children: chls,
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
		Children: []view.View{textView},
		Layouter: l,
		Painter:  v.Painter,
	}
}
