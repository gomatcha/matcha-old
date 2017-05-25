package scrollview

import (
	"math"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
)

const bridgeName = "github.com/overcyn/mochi/view/scrollview"

func init() {
	view.RegisterBridgeMarshaller(bridgeName, func(state interface{}) (proto.Message, error) {
		return state.(proto.Message), nil
	})
}

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
	n := &view.Model{}
	n.Painter = v.Painter
	n.Layouter = &scrollViewLayouter{}

	if v.ContentView != nil {
		n.Add(v.ContentView)
	}

	n.BridgeName = bridgeName
	n.BridgeState = &pb.ScrollView{
		ScrollEnabled:                  v.ScrollEnabled,
		ShowsHorizontalScrollIndicator: v.ShowsHorizontalScrollIndicator,
		ShowsVerticalScrollIndicator:   v.ShowsVerticalScrollIndicator,
	}
	return n
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
