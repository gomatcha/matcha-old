package basicview

import (
	"fmt"
	// "github.com/overcyn/mochi"
	"github.com/overcyn/mochi"
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
	children := map[mochi.Id]view.View{}
	for _, i := range v.Children {
		children[i.Id()] = i
	}

	return &view.Model{
		Children: children,
		Painter:  v.Painter,
		Layouter: v.Layouter,
	}
}

func (v *BasicView) String() string {
	return fmt.Sprintf("&BasicView{%p}", v)
}
