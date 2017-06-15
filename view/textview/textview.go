package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type layouter struct {
	styledText *text.StyledText
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
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

type TextView struct {
	*view.Embed
	String     string
	Style      *text.Style
	StyledText *text.StyledText
}

func New(ctx *view.Context, key interface{}) *TextView {
	v, ok := ctx.Prev(key).(*TextView)
	if !ok {
		v = &TextView{
			Embed: view.NewEmbed(ctx.NewId(key)),
			Style: &text.Style{},
		}
	}
	return v
}

func (v *TextView) Build(ctx *view.Context) *view.Model {
	st := v.StyledText
	if st == nil {
		st = text.NewStyledText(text.New(v.String))
		st.Set(v.Style, 0, 0)
	}

	return &view.Model{
		Layouter:        &layouter{styledText: st},
		NativeViewName:  "github.com/overcyn/mochi/view/textview",
		NativeViewState: st.MarshalProtobuf(),
	}
}
