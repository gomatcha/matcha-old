package scrollview

import (
	"fmt"
	"math"

	"github.com/gogo/protobuf/proto"

	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/scrollview"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
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

	// DefaultOffset  layout.Point
	OffsetNotifier layout.PointNotifier
	OnScroll       func(offset layout.Point)

	ContentChildren []view.View
	ContentPainter  paint.Painter
	ContentLayouter layout.Layouter
	PaintStyle      *paint.Style
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
	child := basicview.New(ctx, "child")
	child.Children = v.ContentChildren
	child.Layouter = v.ContentLayouter
	child.Painter = v.ContentPainter

	var painter paint.Painter
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Children: []view.View{child},
		Painter:  painter,
		Layouter: &layouter{
			Directions: v.Directions,
		},
		NativeViewName: "gomatcha.io/matcha/view/scrollview",
		NativeViewState: &scrollview.View{
			ScrollEnabled:                  v.ScrollEnabled,
			ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
			ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
			ScrollEvents:                   v.OnScroll != nil,
		},
		NativeFuncs: map[string]interface{}{
			"OnScroll": func(data []byte) {
				event := &scrollview.ScrollEvent{}
				err := proto.Unmarshal(data, event)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				if v.OnScroll != nil {
					var offset layout.Point
					(&offset).UnmarshalProtobuf(event.ContentOffset)
					v.OnScroll(offset)
				}
			},
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
