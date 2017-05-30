package textview

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type textViewLayouter struct {
	formattedText *text.Text
}

func (l *textViewLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	size := l.formattedText.Size(layout.Pt(0, 0), ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X, size.Y)}
	return g, nil
}

func (l *textViewLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *textViewLayouter) Unnotify(chan struct{}) {
	// no-op
}

type TextView struct {
	*view.Embed
	String string
	Style  *text.Style
	Text   *text.Text

	// String     string
	// Text       text.Text
	// Attributes text.Attributes
	Painter paint.Painter
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
	ft := v.Text
	if ft == nil {
		ft = &text.Text{}
		ft.SetString(v.String)
		ft.SetStyle(v.Style)
	}

	return &view.Model{
		Layouter:        &textViewLayouter{formattedText: ft},
		Painter:         v.Painter,
		NativeViewName:  "github.com/overcyn/mochi/view/textview",
		NativeViewState: ft.MarshalProtobuf(),
	}
}
