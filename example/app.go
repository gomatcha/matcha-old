package example

import (
	"fmt"
	"reflect"
	"time"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout"
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

type GoRoot struct {
}

func (b *GoRoot) NewViewController(id int) *view.ViewController {
	return view.NewViewController(func(c view.Config) view.View {
		return NewNestedView(c)
	}, id)
}

func init() {
	mochibridge.SetGoRoot(&GoRoot{})
	mochibridge.RegisterType("layout.Point", reflect.TypeOf(layout.Point{}))
	mochibridge.RegisterType("layout.Rect", reflect.TypeOf(layout.Rect{}))
}

type NestedView struct {
	*view.Embed
	counter     int
	ticker      *animate.Ticker
	floatTicker mochi.Float64Notifier
	colorTicker mochi.ColorNotifier
}

func NewNestedView(c view.Config) *NestedView {
	v, ok := c.Prev.(*NestedView)
	if !ok {
		v = &NestedView{}
		v.Embed = c.Embed
		v.ticker = animate.NewTicker(time.Second * 5)
		v.floatTicker = animate.FloatInterpolate(v.ticker, animate.FloatLerp{Start: 0, End: 150})
		v.colorTicker = animate.ColorInterpolate(v.ticker, animate.RGBALerp{Start: internal.RedColor, End: internal.YellowColor})
	}
	return v
}

func (v *NestedView) Build(ctx *view.Context) *view.Model {
	m := &view.Model{}

	l := constraint.New()
	m.Layouter = l

	p := &paint.Style{}
	p.BackgroundColor = internal.GreenColor
	m.Painter = p

	chl1 := basicview.New(ctx.Get(1))
	chl1.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	m.Add(chl1)
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Notifier(v.floatTicker))
		s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	chl2 := basicview.New(ctx.Get(2))
	chl2.Painter = &paint.Style{BackgroundColor: internal.YellowColor}
	m.Add(chl2)
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx.Get(3))
	chl3.Painter = &paint.Style{BackgroundColor: internal.BlueColor}
	m.Add(chl3)
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx.Get(4))
	chl4.Painter = &paint.Style{BackgroundColor: internal.MagentaColor}
	m.Add(chl4)
	g4 := l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(ctx.Get(5))
	chl5.Painter = &paint.Style{BackgroundColor: internal.CyanColor}
	chl5.String = "Subtitle"
	chl5.Style.SetAlignment(text.AlignmentCenter)
	chl5.Style.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl5.Style.SetStrikethroughColor(internal.MagentaColor)
	chl5.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl5.Style.SetUnderlineColor(internal.GreenColor)
	chl5.Style.SetFont(text.Font{
		Family: "American Typewriter",
		Face:   "Bold",
		Size:   20,
	})
	m.Add(chl5)
	g5 := l.Add(chl5, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl6 := textview.New(ctx.Get(6))
	chl6.Painter = &paint.Style{BackgroundColor: internal.RedColor}
	chl6.String = fmt.Sprintf("Counter: %v", v.counter)
	chl6.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	m.Add(chl6)
	g6 := l.Add(chl6, func(s *constraint.Solver) {
		s.BottomEqual(g5.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl8 := imageview.NewURLImageView(ctx.Get(8))
	chl8.Painter = &paint.Style{BackgroundColor: internal.CyanColor}
	chl8.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl8.ResizeMode = imageview.ResizeModeFit
	m.Add(chl8)
	g8 := l.Add(chl8, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	chl9 := button.New(ctx.Get(9))
	chl9.Painter = &paint.Style{BackgroundColor: internal.WhiteColor}
	chl9.Text = "Button"
	chl9.OnPress = func() {
		v.Lock()
		defer v.Unlock()

		fmt.Println("On Click")
		v.counter += 1
		v.Update(nil)
	}
	m.Add(chl9)
	_ = l.Add(chl9, func(s *constraint.Solver) {
		s.BottomEqual(g8.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	childLayouter := &table.Layout{}
	childViews := []view.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx.Get(i + 1000))
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: internal.RedColor}
		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx.Get(10))
	scrollChild.Painter = &paint.Style{BackgroundColor: internal.WhiteColor}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childViews

	chl10 := scrollview.New(ctx.Get(11))
	chl10.Painter = &paint.Style{BackgroundColor: internal.CyanColor}
	chl10.ContentView = scrollChild
	m.Add(chl10)
	_ = l.Add(chl10, func(s *constraint.Solver) {
		s.TopEqual(g4.Bottom())
		s.LeftEqual(g4.Left())
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return m
}

type TableCell struct {
	*view.Embed
	String  string
	Painter paint.Painter
}

func NewTableCell(c view.Config) *TableCell {
	v, ok := c.Prev.(*TableCell)
	if !ok {
		v = &TableCell{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TableCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	n := &view.Model{}
	n.Layouter = l
	n.Painter = v.Painter

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx.Get(1))
	textView.String = v.String
	textView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	n.Add(textView)
	l.Add(textView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(10))
		s.RightEqual(l.Right().Add(-10))
		s.CenterYEqual(l.CenterY())
	})

	return n
}
