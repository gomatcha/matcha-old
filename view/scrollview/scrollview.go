package scrollview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"math"
)

type ScrollView struct {
	*view.Embed
	ScrollEnabled                  bool
	ShowsHorizontalScrollIndicator bool
	ShowsVerticalScrollIndicator   bool
	ContentView                    view.View
	Painter                        paint.Painter
}

func New(c view.Config) *ScrollView {
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

func (v *ScrollView) Build(ctx *view.BuildContext) *view.ViewModel {
	n := &view.ViewModel{}
	n.Painter = v.Painter
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

func (l *scrollViewLayouter) Layout(ctx *layout.LayoutContext) (layout.Guide, map[mochi.Id]layout.Guide) {
	gs := map[mochi.Id]layout.Guide{}
	if len(ctx.ChildIds) > 0 {
		g := ctx.LayoutChild(ctx.ChildIds[0], ctx.MinSize, mochi.Pt(math.Inf(1), math.Inf(1)))
		g.Frame = g.Frame.Add(mochi.Pt(-g.Frame.Min.X, -g.Frame.Min.Y))
		gs[ctx.ChildIds[0]] = g
	}
	return layout.Guide{
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
