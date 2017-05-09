package basicview

import (
	"fmt"
	"github.com/overcyn/mochi"
)

type BasicView struct {
	*mochi.Embed
	PaintOptions mochi.PaintOptions
	Layouter     mochi.Layouter
	Children     []mochi.View
}

func New(c mochi.Config) *BasicView {
	v, ok := c.Prev.(*BasicView)
	if !ok {
		v = &BasicView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *BasicView) Build(ctx *mochi.BuildContext) *mochi.ViewModel {
	n := &mochi.ViewModel{}
	n.Painter = v.PaintOptions
	n.Layouter = v.Layouter
	for _, i := range v.Children {
		n.Add(i)
	}
	return n
}

func (v *BasicView) String() string {
	return fmt.Sprintf("&BasicView{%p}", v)
}
