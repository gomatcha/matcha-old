package textinput

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/text"
)

type TextInput struct {
	*mochi.Embed
	Text         text.Text
	Style        text.Style
	PaintOptions mochi.PaintOptions
}

func New(c mochi.Config) *TextInput {
	v, ok := c.Prev.(*TextInput)
	if !ok {
		v = &TextInput{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TextInput) Build(ctx *mochi.BuildContext) *mochi.ViewModel {
	n := &mochi.ViewModel{}
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
