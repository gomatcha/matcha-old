package table

import (
	"math"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view"
)

type Direction int

const (
	DirectionFromTop Direction = iota
	DirectionFromBottom
	DirectionFromLeft
	DirectionFromRight
)

type Layout struct {
	Direction Direction // TODO(KD): Direction is ignored.
	Ids       []mochi.Id
}

func (l *Layout) Add(v view.View) {
	l.Ids = append(l.Ids, v.Id())
}

func (l *Layout) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	g := layout.Guide{}
	gs := map[mochi.Id]layout.Guide{}
	y := 0.0
	x := ctx.MinSize.X
	for i, id := range l.Ids {
		g := ctx.LayoutChild(id, layout.Pt(x, 0), layout.Pt(x, math.Inf(1)))
		g.Frame = layout.Rt(0, y, g.Width(), y+g.Height())
		g.ZIndex = i
		gs[id] = g
		y += g.Height()
	}
	g.Frame = layout.Rt(0, 0, x, y)
	return g, gs
}

func (l *Layout) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *Layout) Unnotify(chan struct{}) {
	// no-op
}
