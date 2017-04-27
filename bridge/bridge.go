package bridge

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/constraint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/button"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/textview"
	// "mochi.io/mochi/layout/constraint"
	// "mochi.io/mochi/layout"

	"fmt"
	_ "image"
	"mochi/bridge"
	"reflect"
)

type GoRoot struct {
}

func (b *GoRoot) NewBuildContext() *mochi.BuildContext {
	ctx := mochi.NewBuildContext(func(c mochi.Config) mochi.View {
		return New(c)
	})
	return ctx
}

func init() {
	bridge.SetGoRoot(&GoRoot{})
	bridge.RegisterType("mochi.Point", reflect.TypeOf(mochi.Point{}))
	bridge.RegisterType("mochi.Rect", reflect.TypeOf(mochi.Rect{}))
}

const (
	chl1id int = iota
	chl2id
	chl3id
	chl4id
	chl5id
	chl6id
	chl7id
	chl8id
	chl9id
)

type NestedView struct {
	*mochi.Embed
	counter int
}

func New(c mochi.Config) *NestedView {
	v, ok := c.Prev.(*NestedView)
	if !ok {
		v = &NestedView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *NestedView) Build(ctx *mochi.BuildContext) *mochi.Node {
	n := &mochi.Node{}

	l := constraint.New()
	n.Layouter = l

	p := mochi.PaintOptions{}
	p.BackgroundColor = mochi.GreenColor
	n.Painter = p

	chl1 := basicview.New(ctx.Get(chl1id))
	chl1.PaintOptions.BackgroundColor = mochi.RedColor
	n.Set(chl1id, chl1)
	g1 := l.Add(chl1id, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2 := basicview.New(ctx.Get(chl2id))
	chl2.PaintOptions.BackgroundColor = mochi.YellowColor
	n.Set(chl2id, chl2)
	g2 := l.Add(chl2id, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx.Get(chl3id))
	chl3.PaintOptions.BackgroundColor = mochi.BlueColor
	n.Set(chl3id, chl3)
	g3 := l.Add(chl3id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx.Get(chl4id))
	chl4.PaintOptions.BackgroundColor = mochi.MagentaColor
	n.Set(chl4id, chl4)
	_ = l.Add(chl4id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(ctx.Get(chl5id))
	chl5.PaintOptions.BackgroundColor = mochi.CyanColor
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
	n.Set(chl5id, chl5)
	g5 := l.Add(chl5id, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl6 := textview.New(ctx.Get(chl6id))
	chl6.PaintOptions.BackgroundColor = mochi.RedColor
	chl6.String = fmt.Sprintf("Counter: %v", v.counter)
	chl6.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	n.Set(chl6id, chl6)
	g6 := l.Add(chl6id, func(s *constraint.Solver) {
		s.BottomEqual(g5.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	// img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// for x := 0; x < 100; x++ {
	// 	for y := 0; y < 100; y++ {
	// 		img.Set(x, y, mochi.MagentaColor)
	// 	}
	// }

	// chl7 := imageview.NewImageView(ctx.Get(chl7id))
	// chl7.PaintOptions.BackgroundColor = mochi.CyanColor
	// chl7.Image = img
	// n.Set(chl7id, chl7)
	// _ = l.Add(chl7id, func(s *constraint.Solver) {
	// 	s.BottomEqual(g6.Top())
	// 	s.RightEqual(g2.Right().Add(-15))
	// 	s.WidthEqual(constraint.Const(100))
	// 	s.HeightEqual(constraint.Const(100))
	// })

	chl8 := imageview.NewURLImageView(ctx.Get(chl8id))
	chl8.PaintOptions.BackgroundColor = mochi.CyanColor
	chl8.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl8.ResizeMode = imageview.ResizeModeFit
	n.Set(chl8id, chl8)
	g8 := l.Add(chl8id, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	chl9 := button.New(ctx.Get(chl9id))
	chl9.PaintOptions.BackgroundColor = mochi.WhiteColor
	chl9.Text = "Button"
	chl9.OnPress = func() {
		v.Lock()
		defer v.Unlock()

		fmt.Println("On Click")
		v.counter += 1
		v.Update(nil)
	}
	n.Set(chl9id, chl9)
	_ = l.Add(chl9id, func(s *constraint.Solver) {
		s.BottomEqual(g8.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return n
}
