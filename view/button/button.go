package button // import "gomatcha.io/matcha/view/button"

import (
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/paint"
	pbbutton "gomatcha.io/matcha/pb/view/button"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
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
	Text       string
	OnPress    func()
	PaintStyle *paint.Style
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

	painter := paint.Painter(nil)
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		Layouter:       &layouter{styledText: st},
		NativeViewName: "gomatcha.io/matcha/view/button",
		NativeViewState: &pbbutton.View{
			StyledText: st.MarshalProtobuf(),
		},
		NativeFuncs: map[string]interface{}{
			"OnPress": func() {
				if v.OnPress != nil {
					v.OnPress()
				}
			},
		},
	}
}
