package button

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout"
	pbbutton "github.com/overcyn/mochi/pb/button"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type layouter struct {
	styledText *text.StyledText
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
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

	funcId := ctx.NewFuncId()
	f := func() {
		if v.OnPress != nil {
			v.OnPress(v)
		}
	}

	return &view.Model{
		Layouter:       &layouter{styledText: st},
		NativeViewName: "github.com/overcyn/mochi/view/button",
		NativeViewState: &pbbutton.Button{
			StyledText: st.MarshalProtobuf(),
			OnPress:    funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}
