package full

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
)

type Layout struct {
}

func (l *Layout) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	g := layout.Guide{Frame: mochi.Rect{Max: ctx.MinSize}}
	gs := map[mochi.Id]layout.Guide{}
	for _, id := range ctx.ChildIds {
		gs[id] = ctx.LayoutChild(id, ctx.MinSize, ctx.MinSize)
	}
	return g, gs
}

func (l *Layout) Notify() chan struct{} {
	return nil // no-op
}

func (l *Layout) Unnotify(chan struct{}) {
	// no-op
}
