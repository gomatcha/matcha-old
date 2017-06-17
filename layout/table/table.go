package table

import (
	"math"

	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/view"
)

// type Direction int

// const (
// 	DirectionFromTop Direction = iota
// 	DirectionFromBottom
// 	DirectionFromLeft
// 	DirectionFromRight
// )

type Layout struct {
	// Direction Direction // TODO(KD): Direction is ignored.
	ids   []matcha.Id
	views []view.View
}

func (l *Layout) Views() []view.View {
	return l.views
}

func (l *Layout) Add(v view.View) {
	l.ids = append(l.ids, v.Id())
	l.views = append(l.views, v)
}

func (l *Layout) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	g := layout.Guide{}
	gs := map[matcha.Id]layout.Guide{}
	y := 0.0
	x := ctx.MinSize.X
	for i, id := range l.ids {
		g := ctx.LayoutChild(id, layout.Pt(x, 0), layout.Pt(x, math.Inf(1)))
		g.Frame = layout.Rt(0, y, g.Width(), y+g.Height())
		g.ZIndex = i
		gs[id] = g
		y += g.Height()
	}
	g.Frame = layout.Rt(0, 0, x, y)
	return g, gs
}

func (l *Layout) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *Layout) Unnotify(id comm.Id) {
	// no-op
}
