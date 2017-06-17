package textview

import (
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/text"
	"github.com/overcyn/matcha/view"
)

type layouter struct {
	styledText *text.StyledText
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	size := l.styledText.Size(layout.Pt(0, 0), ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}

type View struct {
	*view.Embed
	String     string
	Style      *text.Style
	StyledText *text.StyledText
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
		Style: &text.Style{},
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	st := v.StyledText
	if st == nil {
		t := text.New(v.String)
		st = text.NewStyledText(t)
		st.Set(v.Style, 0, 0)
	}

	return &view.Model{
		Layouter:        &layouter{styledText: st},
		NativeViewName:  "github.com/overcyn/matcha/view/textview",
		NativeViewState: st.MarshalProtobuf(),
	}
}
