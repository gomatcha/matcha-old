package textinput

import (
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type TextInput struct {
	*view.Embed
	Text    text.Text
	Style   text.Style
	Painter paint.Painter
}

func New(c view.Config) *TextInput {
	v, ok := c.Prev.(*TextInput)
	if !ok {
		v = &TextInput{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TextInput) Build(ctx *view.Context) *view.Model {
	n := &view.Model{}
	n.Painter = v.Painter
	n.BridgeName = "github.com/overcyn/mochi/view/textinput TextInput"
	n.BridgeState = struct {
		Text    *text.Text
		OnPress func()
	}{
	// Text:    ft,
	// OnPress: v.OnPress,
	}
	return n
}
