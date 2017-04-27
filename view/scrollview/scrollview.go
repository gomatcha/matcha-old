package scrollview

import (
	"github.com/overcyn/mochi"
	"math"
)

const (
	chlid int = iota
)

type scrollViewLayouter struct {
}

func (l *scrollViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	gs := map[interface{}]mochi.Guide{}
	if len(ctx.ChildKeys) > 0 {
		g := ctx.LayoutChild(chlid, mochi.Pt(0, 0), mochi.Pt(math.Inf(1), math.Inf(1)))
		g.Frame = g.Frame.Add(mochi.Pt(-g.Frame.Min.X, -g.Frame.Min.Y))
		gs[chlid] = g
		println("layout scrollview", g.Frame.String())
	}
	return mochi.Guide{
		Frame: mochi.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y),
	}, gs
}

type ScrollView struct {
	*mochi.Embed
	ContentView  mochi.View
	PaintOptions mochi.PaintOptions
}

func New(c mochi.Config) *ScrollView {
	v, ok := c.Prev.(*ScrollView)
	if !ok {
		v = &ScrollView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *ScrollView) Build(ctx *mochi.BuildContext) *mochi.Node {
	n := &mochi.Node{}
	n.Painter = v.PaintOptions

	n.Layouter = &scrollViewLayouter{}

	if v.ContentView != nil {
		n.Set(chlid, v.ContentView)
		println(v.ContentView)
	}

	n.Bridge.Name = "github.com/overcyn/mochi/view/scrollview"
	// n.Bridge.State = struct {
	// 	Size mochi.Point
	// }{
	// 	Size: nil,
	// }
	return n
}
