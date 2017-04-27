package scrollview

import (
	"github.com/overcyn/mochi"
)

type ScrollView struct {
	*mochi.Embed
	ContentView  View
	PaintOptions mochi.PaintOptions
}

func New(c mochi.Config) *ScrollView {
	v, ok := c.Prev.(*ScrollView)
	if !ok {
		v = &ScrollView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *ScrollView) Build(ctx *mochi.BuildContext) *mochi.Node {
	n := &mochi.Node{}
	n.Layouter = &textViewLayouter{formattedText: ft}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi/view/scrollview"
	n.Bridge.State = nil
	return n
}
