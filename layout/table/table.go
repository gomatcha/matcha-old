package table

import (
	"github.com/overcyn/mochi"
	"math"
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
	Ids       []interface{}
}

func (l *Layout) Add(id interface{}) {
	l.Ids = append(l.Ids, id)
}

func (l *Layout) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	g := mochi.Guide{}
	gs := map[interface{}]mochi.Guide{}
	y := 0.0
	x := ctx.MinSize.X
	for i, id := range l.Ids {
		g := ctx.LayoutChild(id, mochi.Pt(x, 0), mochi.Pt(x, math.Inf(1)))
		g.Frame = mochi.Rt(0, y, g.Width(), y+g.Height())
		g.ZIndex = i
		gs[id] = g
		y += g.Height()
	}
	g.Frame = mochi.Rt(0, 0, x, y)
	return g, gs
}
