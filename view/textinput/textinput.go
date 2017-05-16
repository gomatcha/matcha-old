package textinput

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
)

type TextInput struct {
	*view.Embed
	Text         text.Text
	Style        text.Style
	PaintOptions paint.PaintStyle
}

func New(c view.Config) *TextInput {
	v, ok := c.Prev.(*TextInput)
	if !ok {
		v = &TextInput{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TextInput) Build(ctx *view.BuildContext) *view.ViewModel {
	n := &view.ViewModel{}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi/view/textinput TextInput"
	n.Bridge.State = struct {
		Text    *text.Text
		OnPress func()
	}{
	// Text:    ft,
	// OnPress: v.OnPress,
	}
	return n
}
