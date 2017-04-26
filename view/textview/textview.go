package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
)

type textViewLayouter struct {
	formattedText *text.Text
}

func (l *textViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	size := l.formattedText.Size(ctx.MaxSize)
	g := mochi.Guide{Frame: mochi.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

type TextView struct {
	*mochi.Embed
	String string
	Format *text.Format
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
		v.Format = &text.Format{}
	}
	return v
}

func (v *TextView) Build(ctx *mochi.BuildContext) *mochi.Node {
	ft := v.Text
	if ft == nil {
		ft = &text.Text{}
		ft.SetString(v.String)
		ft.SetFormat(v.Format)
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
