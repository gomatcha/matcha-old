package button

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
)

const padding = 10.0

type buttonLayouter struct {
	formattedText *text.FormattedText
}

func (l *buttonLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	size := l.formattedText.Size(ctx.MaxSize)
	g := mochi.Guide{Frame: mochi.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

type Button struct {
	*mochi.Embed
	Text         string
	PaintOptions mochi.PaintOptions
	// OnPress
}

func New(c mochi.Config) *Button {
	v, ok := c.Prev.(*Button)
	if !ok {
		v = &Button{}
		v.Embed = c.Embed
	}
	return v
}

func (v *Button) Build(ctx *mochi.BuildContext) *mochi.Node {
	ft := &text.FormattedText{}
	ft.SetString(v.Text)
	ft.Format().SetAlignment(text.AlignmentCenter)
	ft.Format().SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})

	n := &mochi.Node{}
	n.Layouter = &buttonLayouter{formattedText: ft}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi/view/button Button"
	n.Bridge.State = struct {
		FormattedText *text.FormattedText
	}{
		FormattedText: ft,
	}
	return n
}
