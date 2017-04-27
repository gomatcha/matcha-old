package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
	"mochi/bridge"
)

type textViewLayouter struct {
	formattedText *text.Text
}

func (l *textViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	size := textSize(l.formattedText, ctx.MaxSize)
	g := mochi.Guide{Frame: mochi.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

func textSize(t *text.Text, max mochi.Point) mochi.Point {
	return bridge.Root().Call("sizeForAttributedString:minSize:maxSize:", bridge.Interface(t), nil, bridge.Interface(max)).ToInterface().(mochi.Point)
}

type TextView struct {
	*mochi.Embed
	String string
	Style  *text.Style
	Text   *text.Text

	// String     string
	// Text       text.Text
	// Attributes text.Attributes
	PaintOptions mochi.PaintOptions
}

func New(c mochi.Config) *TextView {
	v, ok := c.Prev.(*TextView)
	if !ok {
		v = &TextView{}
		v.Embed = c.Embed
		v.Style = &text.Style{}
	}
	return v
}

func (v *TextView) Build(ctx *mochi.BuildContext) *mochi.Node {
	ft := v.Text
	if ft == nil {
		ft = &text.Text{}
		ft.SetString(v.String)
		ft.SetStyle(v.Style)
	}

	n := &mochi.Node{}
	n.Layouter = &textViewLayouter{formattedText: ft}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi TextView"
	n.Bridge.State = struct {
		Text *text.Text
	}{
		Text: ft,
	}
	return n
}
