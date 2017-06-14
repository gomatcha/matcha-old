package textinput

import (
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
)

type View struct {
	*view.Embed
	String string
	Style  text.Style
	Text   text.Text
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
	return &view.Model{
		NativeViewName: "github.com/overcyn/mochi/view/textinput",
	}
}
