package example

import (
	"fmt"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	_ "github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/button"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"reflect"
	"time"
)

type GoRoot struct {
}

func (b *GoRoot) NewViewController(id int) *mochi.ViewController {
	return mochi.NewViewController(func(c mochi.Config) mochi.View {
		return NewNestedView(c)
	}, id)
}

func init() {
	mochibridge.SetGoRoot(&GoRoot{})
	mochibridge.RegisterType("mochi.Point", reflect.TypeOf(mochi.Point{}))
	mochibridge.RegisterType("mochi.Rect", reflect.TypeOf(mochi.Rect{}))
}

const (
	chl1id int = 1000 + iota
	chl2id
	chl3id
	chl4id
	chl5id
	chl6id
	chl7id
	chl8id
	chl9id
	scrollId      = "scroll"
	scrollChildId = "scrollChild"
)

type NestedView struct {
	*mochi.Embed
	counter     int
	ticker      *animate.Ticker
	floatTicker mochi.Float64Notifier
}

func NewNestedView(c mochi.Config) *NestedView {
	v, ok := c.Prev.(*NestedView)
	if !ok {
		v = &NestedView{}
		v.Embed = c.Embed
		v.ticker = animate.NewTicker(time.Second * 5)
		// v.floatTicker = animate.FloatInterpolate(v.ticker, animate.FloatLerp{Start: 0, End: 150})
		// fmt.Println("Float ticker", v.floatTicker.Value())
	}
	return v
}

func (v *NestedView) Build(ctx *mochi.BuildContext) *mochi.ViewModel {
	m := &mochi.ViewModel{}

	l := constraint.New()
	m.Layouter = l

	p := &mochi.PaintStyle{}
	p.BackgroundColor = mochi.GreenColor
	m.Painter = p

	chl1 := basicview.New(ctx.Get("red"))
	// chl1.Painter = &paint.Style{BackgroundColor: mochi.RedColor}
	// chl1.Painter = &paint.AnimatedStyle{BackgroundColor: mochi.RedColor}
	chl1.Painter = &mochi.PaintStyle{BackgroundColor: mochi.RedColor}
	m.Add(chl1)
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
		// s.WidthEqual(constraint.Notifier(v.floatTicker))
		// s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	chl2 := basicview.New(ctx.Get(chl2id))
	chl2.Painter = &mochi.PaintStyle{BackgroundColor: mochi.YellowColor}
	m.Add(chl2)
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx.Get(chl3id))
	chl3.Painter = &mochi.PaintStyle{BackgroundColor: mochi.BlueColor}
	m.Add(chl3)
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx.Get(chl4id))
	chl4.Painter = &mochi.PaintStyle{BackgroundColor: mochi.MagentaColor}
	m.Add(chl4)
	g4 := l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(ctx.Get(chl5id))
	chl5.Painter = &mochi.PaintStyle{BackgroundColor: mochi.CyanColor}
	chl5.String = "Subtitle"
	chl5.Style.SetAlignment(text.AlignmentCenter)
	chl5.Style.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl5.Style.SetStrikethroughColor(mochi.MagentaColor)
	chl5.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl5.Style.SetUnderlineColor(mochi.GreenColor)
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

	chl6 := textview.New(ctx.Get(chl6id))
	chl6.Painter = &mochi.PaintStyle{BackgroundColor: mochi.RedColor}
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

	chl8 := imageview.NewURLImageView(ctx.Get(chl8id))
	chl8.Painter = &mochi.PaintStyle{BackgroundColor: mochi.CyanColor}
	chl8.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl8.ResizeMode = imageview.ResizeModeFit
	// chl8 := imageview.NewImageView(ctx.Get(chl8id))
	// chl8.PaintOptions.BackgroundColor = mochi.CyanColor
	// chl8.ResizeMode = imageview.ResizeModeFit
	m.Add(chl8)
	g8 := l.Add(chl8, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	chl9 := button.New(ctx.Get(chl9id))
	chl9.Painter = &mochi.PaintStyle{BackgroundColor: mochi.WhiteColor}
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
	childViews := []mochi.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx.Get(i))
		childView.String = "TEST TEST"
		childView.Painter = &mochi.PaintStyle{BackgroundColor: mochi.RedColor}
		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx.Get(scrollChildId))
	scrollChild.Painter = &mochi.PaintStyle{BackgroundColor: mochi.WhiteColor}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childViews

	chl10 := scrollview.New(ctx.Get(scrollId))
	chl10.Painter = &mochi.PaintStyle{BackgroundColor: mochi.CyanColor}
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

const (
	textId int = 1000
)

type TableCell struct {
	*mochi.Embed
	String  string
	Painter mochi.Painter
}

func NewTableCell(c mochi.Config) *TableCell {
	v, ok := c.Prev.(*TableCell)
	if !ok {
		v = &TableCell{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TableCell) Build(ctx *mochi.BuildContext) *mochi.ViewModel {
	l := constraint.New()
	n := &mochi.ViewModel{}
	n.Layouter = l
	n.Painter = v.Painter

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx.Get(textId))
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
