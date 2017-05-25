package scrollview

import (
	"math"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
)

type ScrollView struct {
	*view.Embed
	ScrollEnabled                  bool
	ShowsHorizontalScrollIndicator bool
	ShowsVerticalScrollIndicator   bool
	ContentView                    view.View
	Painter                        paint.Painter
}

func New(ctx *view.Context, key interface{}) *ScrollView {
	v, ok := ctx.Prev(key).(*ScrollView)
	if !ok {
		v = &ScrollView{
			Embed: view.NewEmbed(ctx.NewId(key)),
			ShowsHorizontalScrollIndicator: true,
			ShowsVerticalScrollIndicator:   true,
			ScrollEnabled:                  true,
		}
	}
	return v
}

func (v *ScrollView) Build(ctx *view.Context) *view.Model {
	children := map[mochi.Id]view.View{}
	if v.ContentView != nil {
		children[v.ContentView.Id()] = v.ContentView
	}

	return &view.Model{
		Children:   children,
		Painter:    v.Painter,
		Layouter:   &scrollViewLayouter{},
		NativeName: "github.com/overcyn/mochi/view/scrollview",
		NativeStateProtobuf: &pb.ScrollView{
			ScrollEnabled:                  v.ScrollEnabled,
			ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
			ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
		},
	}
}

type scrollViewLayouter struct {
}

func (l *scrollViewLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	gs := map[mochi.Id]layout.Guide{}
	if len(ctx.ChildIds) > 0 {
		g := ctx.LayoutChild(ctx.ChildIds[0], ctx.MinSize, layout.Pt(math.Inf(1), math.Inf(1)))
		g.Frame = g.Frame.Add(layout.Pt(-g.Frame.Min.X, -g.Frame.Min.Y))
		gs[ctx.ChildIds[0]] = g
	}
	return layout.Guide{
		Frame: layout.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y),
	}, gs
}

func (l *scrollViewLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *scrollViewLayouter) Unnotify(c chan struct{}) {
	// no-op
}
