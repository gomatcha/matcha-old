package textinput

import (
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

// View mutates it's Text and StyledText fields in place.
type View struct {
	*view.Embed
	Text       *text.Text
	Style      *text.Style
	StyledText *text.StyledText
	// Cursor position?
	// Keyboard visibility?

	OnChange func(*View)
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); !ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	st := v.StyledText
	if st == nil {
		st = text.NewStyledText(v.Text)
		st.Set(v.Style, 0, 0)
	}

	return &view.Model{
		NativeViewName:  "github.com/overcyn/mochi/view/textinput",
		NativeViewState: st.MarshalProtobuf(),
	}
}
