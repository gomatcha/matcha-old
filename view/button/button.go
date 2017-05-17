package button

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochibridge"
)

func textSize(t *text.Text, max layout.Point) layout.Point {
	return mochibridge.Root().Call("sizeForAttributedString:minSize:maxSize:", mochibridge.Interface(t), nil, mochibridge.Interface(max)).ToInterface().(layout.Point)
}

const padding = 10.0

type buttonLayouter struct {
	formattedText *text.Text
}

func (l *buttonLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	size := textSize(l.formattedText, ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

func (l *buttonLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *buttonLayouter) Unnotify(chan struct{}) {
	// no-op
}

type Button struct {
	*view.Embed
	Text    string
	Painter paint.Painter
	OnPress func()
}

func New(c view.Config) *Button {
	v, ok := c.Prev.(*Button)
	if !ok {
		v = &Button{}
		v.Embed = c.Embed
	}
	return v
}

func (v *Button) Build(ctx *view.Context) *view.Model {
	ft := &text.Text{}
	ft.SetString(v.Text)
	ft.Style().SetAlignment(text.AlignmentCenter)
	ft.Style().SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})

	n := &view.Model{
		Layouter:   &buttonLayouter{formattedText: ft},
		Painter:    v.Painter,
		BridgeName: "github.com/overcyn/mochi/view/button",
		BridgeState: struct {
			Text    *text.Text
			OnPress func()
		}{
			Text:    ft,
			OnPress: v.OnPress,
		},
	}
	return n
}
