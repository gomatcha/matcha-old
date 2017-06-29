package scrollview // import "gomatcha.io/matcha/view/scrollview"

import (
	"math"

	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/scrollview"
	"gomatcha.io/matcha/view"
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

func New(ctx *view.Context, key string) *ScrollView {
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
		Layouter: &layouter{
			Directions: v.Directions,
		},
		NativeViewName: "gomatcha.io/matcha/view/scrollview",
		NativeViewState: &scrollview.View{
			ScrollEnabled:                  v.ScrollEnabled,
			ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
			ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
		},
	}
}

type layouter struct {
	Directions Direction
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	gs := map[matcha.Id]layout.Guide{}
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

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}
