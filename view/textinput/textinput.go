package textinput

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha/internal"
	"gomatcha.io/matcha/keyboard"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/textinput"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
)

// View mutates the Text and StyledText fields in place.
type View struct {
	*view.Embed
	PaintStyle         *paint.Style
	Text               *text.Text
	Style              *text.Style
	KeyboardType       keyboard.Type
	KeyboardAppearance keyboard.Appearance
	KeyboardReturnType keyboard.ReturnType
	Responder          *keyboard.Responder
	responder          *keyboard.Responder

	// TODO(KD):
	// StyledText *text.StyledText
	// Cursor position?

	OnChange func(*text.Text)
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
	var st *internal.StyledText
	if v.Text != nil {
		st = internal.NewStyledText(v.Text)
		st.Set(v.Style, 0, 0)
	}

	if v.Responder != v.responder {
		if v.responder != nil {
			v.Unsubscribe(v.responder)
		}

		v.responder = v.Responder
		if v.responder != nil {
			v.Subscribe(v.responder)
		}
	}

	focused := false
	if v.responder != nil {
		focused = v.responder.Visible()
	}

	painter := paint.Painter(nil)
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		NativeViewName: "gomatcha.io/matcha/view/textinput",
		NativeViewState: &textinput.View{
			StyledText:         st.MarshalProtobuf(),
			KeyboardType:       v.KeyboardType.MarshalProtobuf(),
			KeyboardAppearance: v.KeyboardAppearance.MarshalProtobuf(),
			KeyboardReturnType: v.KeyboardReturnType.MarshalProtobuf(),
			Focused:            focused,
		},
		NativeFuncs: map[string]interface{}{
			"OnChange": func(data []byte) {
				pbevent := &textinput.Event{}
				err := proto.Unmarshal(data, pbevent)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				_ = v.Text.UnmarshalProtobuf(pbevent.StyledText.Text)
				if v.OnChange != nil {
					v.OnChange(v.Text)
				}
			},
			"OnFocus": func(data []byte) {
				pbevent := &textinput.FocusEvent{}
				err := proto.Unmarshal(data, pbevent)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				if v.responder != nil {
					if pbevent.Focused {
						v.responder.Show()
					} else {
						v.responder.Dismiss()
					}
				}
			},
		},
	}
}
