package scrollview

import (
	"math"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
)

type Direction int

const (
	Horizontal Direction = 1 << iota
	Vertical
)

type ScrollView struct {
	*view.Embed
	Directions                     Direction
	ScrollEnabled                  bool // TODO(KD): replace with Directions flag
	ShowsHorizontalScrollIndicator bool // TODO(KD): replace with Directions flag
	ShowsVerticalScrollIndicator   bool
	ContentView                    view.View
	Painter                        paint.Painter
}

func New(ctx *view.Context, key interface{}) *ScrollView {
	if v, ok := ctx.Prev(key).(*ScrollView); ok {
		return v
	}
	return &ScrollView{
		Embed:                          view.NewEmbed(ctx.NewId(key)),
		Directions:                     Vertical,
		ShowsHorizontalScrollIndicator: true,
		ShowsVerticalScrollIndicator:   true,
		ScrollEnabled:                  true,
	}
}

func (v *ScrollView) Build(ctx *view.Context) *view.Model {
	children := []view.View{}
	if v.ContentView != nil {
		children = append(children, v.ContentView)
	}

	return &view.Model{
		Children: children,
		Painter:  v.Painter,
		Layouter: &scrollViewLayouter{
			Directions: v.Directions,
		},
		NativeViewName: "github.com/overcyn/mochi/view/scrollview",
		NativeViewState: &pb.ScrollView{
			ScrollEnabled:                  v.ScrollEnabled,
			ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
			ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
		},
	}
}

type scrollViewLayouter struct {
	Directions Direction
}

func (l *scrollViewLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	gs := map[mochi.Id]layout.Guide{}
	if len(ctx.ChildIds) > 0 {
		minSize := ctx.MinSize
		if l.Directions&Horizontal == Horizontal {
			minSize.X = 0
		}
		if l.Directions&Vertical == Vertical {
			minSize.Y = 0
		}

		g := ctx.LayoutChild(ctx.ChildIds[0], minSize, layout.Pt(math.Inf(1), math.Inf(1)))
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
