package textinput

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/matcha/keyboard"
	"github.com/overcyn/matcha/pb/view/textinput"
	"github.com/overcyn/matcha/text"
	"github.com/overcyn/matcha/view"
)

// View mutates the Text and StyledText fields in place.
type View struct {
	*view.Embed
	Text      *text.Text
	Style     *text.Style
	Responder *keyboard.Responder

	// TODO(KD):
	// StyledText *text.StyledText
	// Cursor position?
	// Keyboard visibility?

	OnChange func(*View)
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Text:  text.New(""),
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	// st := v.StyledText
	var st *text.StyledText
	if st == nil {
		st = text.NewStyledText(v.Text)
		st.Set(v.Style, 0, 0)
	}

	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		pbevent := &textinput.Event{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		_ = v.Text.UnmarshalProtobuf(pbevent.StyledText.Text)
		if v.OnChange != nil {
			v.OnChange(v)
		}
	}

	return &view.Model{
		Values: map[interface{}]interface{}{
			keyboard.HelperKey: v.Responder,
		},
		NativeViewName: "github.com/overcyn/matcha/view/textinput",
		NativeViewState: &textinput.View{
			StyledText: st.MarshalProtobuf(),
			OnUpdate:   funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}
