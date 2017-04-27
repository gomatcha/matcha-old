package basicview

import (
	"github.com/overcyn/mochi"
)

type BasicView struct {
	*mochi.Embed
	PaintOptions mochi.PaintOptions
	Layouter     mochi.Layouter
}

func New(c mochi.Config) *BasicView {
	v, ok := c.Prev.(*BasicView)
	if !ok {
		v = &BasicView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *BasicView) Build(ctx *mochi.BuildContext) *mochi.Node {
	n := &mochi.Node{}
	n.Painter = v.PaintOptions
	n.Layouter = v.Layouter
	return n
}
