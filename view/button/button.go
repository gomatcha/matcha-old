package button

import (
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	pbbutton "github.com/overcyn/matcha/pb/view/button"
	"github.com/overcyn/matcha/text"
	"github.com/overcyn/matcha/view"
)

type layouter struct {
	styledText *text.StyledText
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	const padding = 10.0
	size := l.styledText.Size(layout.Pt(0, 0), ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}

type Button struct {
	*view.Embed
	Text    string
	OnPress func(*Button)
}

func New(ctx *view.Context, key string) *Button {
	if v, ok := ctx.Prev(key).(*Button); ok {
		return v
	}
	return &Button{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *Button) Build(ctx *view.Context) *view.Model {
	style := &text.Style{}
	style.SetAlignment(text.AlignmentCenter)
	style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	t := text.New(v.Text)
	st := text.NewStyledText(t)
	st.Set(style, 0, 0)

	return &view.Model{
		Layouter:       &layouter{styledText: st},
		NativeViewName: "github.com/overcyn/matcha/view/button",
		NativeViewState: &pbbutton.View{
			StyledText: st.MarshalProtobuf(),
		},
		NativeFuncs: map[string]interface{}{
			"OnPress": func() {
				if v.OnPress != nil {
					v.OnPress(v)
				}
			},
		},
	}
}
