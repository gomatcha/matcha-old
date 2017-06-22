package full

import (
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
)

type Layouter struct {
}

func (l *Layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rect{Max: ctx.MinSize}}
	gs := map[matcha.Id]layout.Guide{}
	for _, id := range ctx.ChildIds {
		gs[id] = ctx.LayoutChild(id, ctx.MinSize, ctx.MinSize)
	}
	return g, gs
}

func (l *Layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *Layouter) Unnotify(id comm.Id) {
	// no-op
}
