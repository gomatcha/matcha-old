package textinput

import (
	"github.com/overcyn/mochi/pb/view/textinput"
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
	if v, ok := ctx.Prev(key).(*View); ok {
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

	funcId := ctx.NewFuncId()
	f := func() {
		if v.OnChange != nil {
			v.OnChange(v)
		}
	}

	return &view.Model{
		NativeViewName: "github.com/overcyn/mochi/view/textinput",
		NativeViewState: &textinput.View{
			StyledText: st.MarshalProtobuf(),
			OnUpdate:   funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}
