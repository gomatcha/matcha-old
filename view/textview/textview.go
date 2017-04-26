package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
)

type textViewLayouter struct {
	formattedText *text.FormattedText
}

func (l *textViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	size := l.formattedText.Size(ctx.MaxSize)
	g := mochi.Guide{Frame: mochi.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

type TextView struct {
	*mochi.Embed
	Text          string
	Format        *text.Format
	FormattedText *text.FormattedText
	PaintOptions  mochi.PaintOptions
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
	ft := v.FormattedText
	if ft == nil {
		ft = &text.FormattedText{}
		ft.SetString(v.Text)
		ft.SetFormat(v.Format)
	}

	n := &mochi.Node{}
	n.Layouter = &textViewLayouter{formattedText: ft}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi TextView"
	n.Bridge.State = struct {
		FormattedText *text.FormattedText
	}{
		FormattedText: ft,
	}
	return n
}
