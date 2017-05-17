package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochibridge"
)

type textViewLayouter struct {
	formattedText *text.Text
}

func (l *textViewLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	size := textSize(l.formattedText, ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

func (l *textViewLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *textViewLayouter) Unnotify(chan struct{}) {
	// no-op
}

func textSize(t *text.Text, max layout.Point) layout.Point {
	return mochibridge.Root().Call("sizeForAttributedString:minSize:maxSize:", mochibridge.Interface(t), nil, mochibridge.Interface(max)).ToInterface().(layout.Point)
}

type TextView struct {
	*view.Embed
	String string
	Style  *text.Style
	Text   *text.Text

	// String     string
	// Text       text.Text
	// Attributes text.Attributes
	Painter paint.Painter
}

func New(c view.Config) *TextView {
	v, ok := c.Prev.(*TextView)
	if !ok {
		v = &TextView{}
		v.Embed = c.Embed
		v.Style = &text.Style{}
	}
	return v
}

func (v *TextView) Build(ctx *view.Context) *view.Model {
	ft := v.Text
	if ft == nil {
		ft = &text.Text{}
		ft.SetString(v.String)
		ft.SetStyle(v.Style)
	}

	n := &view.Model{}
	n.Layouter = &textViewLayouter{formattedText: ft}
	n.Painter = v.Painter
	n.BridgeName = "github.com/overcyn/mochi/view/textview"
	n.BridgeState = struct {
		Text *text.Text
	}{
		Text: ft,
	}
	return n
}
