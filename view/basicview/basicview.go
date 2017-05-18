package basicview

import (
	"fmt"
	// "github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
)

type BasicView struct {
	*view.Embed
	Painter  paint.Painter
	Layouter layout.Layouter
	Children []view.View
}

func New(ctx *view.Context, key interface{}) *BasicView {
	v, ok := ctx.Prev(key).(*BasicView)
	if !ok {
		v = &BasicView{
			Embed: view.NewEmbed(ctx.NewId(key)),
		}
	}
	return v
}

func (v *BasicView) Build(ctx *view.Context) *view.Model {
	n := &view.Model{
		Painter:  v.Painter,
		Layouter: v.Layouter,
	}
	for _, i := range v.Children {
		n.Add(i)
	}
	return n
}

func (v *BasicView) String() string {
	return fmt.Sprintf("&BasicView{%p}", v)
}
