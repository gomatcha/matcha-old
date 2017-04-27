package button

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
	"mochi/bridge"
)

func textSize(t *text.Text, max mochi.Point) mochi.Point {
	return bridge.Root().Call("sizeForAttributedString:minSize:maxSize:", bridge.Interface(t), nil, bridge.Interface(max)).ToInterface().(mochi.Point)
}

const padding = 10.0

type buttonLayouter struct {
	formattedText *text.Text
}

func (l *buttonLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	size := textSize(l.formattedText, ctx.MaxSize)
	g := mochi.Guide{Frame: mochi.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

type Button struct {
	*mochi.Embed
	Text         string
	PaintOptions mochi.PaintOptions
	OnPress      func()
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
	ft := &text.Text{}
	ft.SetString(v.Text)
	ft.Style().SetAlignment(text.AlignmentCenter)
	ft.Style().SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})

	n := &mochi.Node{}
	n.Layouter = &buttonLayouter{formattedText: ft}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi/view/button"
	n.Bridge.State = struct {
		Text    *text.Text
		OnPress func()
	}{
		Text:    ft,
		OnPress: v.OnPress,
	}
	return n
}
