package scrollview

import (
	"github.com/overcyn/mochi"
	"math"
)

type ScrollView struct {
	*mochi.Embed
	ScrollEnabled                  bool
	ShowsHorizontalScrollIndicator bool
	ShowsVerticalScrollIndicator   bool
	ContentView                    mochi.View
	PaintOptions                   mochi.PaintOptions
}

func New(c mochi.Config) *ScrollView {
	v, ok := c.Prev.(*ScrollView)
	if !ok {
		v = &ScrollView{}
		v.Embed = c.Embed
		v.ShowsHorizontalScrollIndicator = true
		v.ShowsVerticalScrollIndicator = true
		v.ScrollEnabled = true
	}
	return v
}

func (v *ScrollView) Build(ctx *mochi.BuildContext) *mochi.ViewModel {
	n := &mochi.ViewModel{}
	n.Painter = v.PaintOptions

	n.Layouter = &scrollViewLayouter{}

	if v.ContentView != nil {
		n.Add(v.ContentView)
	}

	n.Bridge.Name = "github.com/overcyn/mochi/view/scrollview"
	n.Bridge.State = struct {
		ScrollEnabled                  bool
		ShowsHorizontalScrollIndicator bool
		ShowsVerticalScrollIndicator   bool
	}{
		ScrollEnabled:                  v.ScrollEnabled,
		ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
		ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
	}
	return n
}

type scrollViewLayouter struct {
}

func (l *scrollViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[mochi.Id]mochi.Guide) {
	gs := map[mochi.Id]mochi.Guide{}
	if len(ctx.ChildIds) > 0 {
		g := ctx.LayoutChild(ctx.ChildIds[0], ctx.MinSize, mochi.Pt(math.Inf(1), math.Inf(1)))
		g.Frame = g.Frame.Add(mochi.Pt(-g.Frame.Min.X, -g.Frame.Min.Y))
		gs[ctx.ChildIds[0]] = g
	}
	return mochi.Guide{
		Frame: mochi.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y),
	}, gs
}

func (l *scrollViewLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *scrollViewLayouter) Unnotify(c chan struct{}) {
	// no-op
}
