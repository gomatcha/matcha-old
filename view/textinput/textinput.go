package textinput

import (
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type View struct {
	*view.Embed
	String  string
	Style   text.Style
	Text    text.Text
	Painter paint.Painter
}

func New(ctx *view.Context, key interface{}) *View {
	if v, ok := ctx.Prev(key).(*View); !ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	n := &view.Model{}
	n.Painter = v.Painter
	n.NativeViewName = "github.com/overcyn/mochi/view/textinput"
	// n.BridgeState = struct {
	// 	Text    *text.Text
	// 	OnPress func()
	// }{
	// Text:    ft,
	// OnPress: v.OnPress,
	// }
	return n
}
