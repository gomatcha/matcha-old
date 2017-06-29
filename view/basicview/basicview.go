package basicview

import (
	"fmt"

	"github.com/gomatcha/matcha/layout"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/view"
)

type BasicView struct {
	*view.Embed
	Painter  paint.Painter
	Layouter layout.Layouter
	Children []view.View
}

func New(ctx *view.Context, key string) *BasicView {
	if v, ok := ctx.Prev(key).(*BasicView); ok {
		return v
	}
	return &BasicView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *BasicView) Build(ctx *view.Context) *view.Model {
	return &view.Model{
		Children: v.Children,
		Painter:  v.Painter,
		Layouter: v.Layouter,
	}
}

func (v *BasicView) String() string {
	return fmt.Sprintf("&BasicView{%p}", v)
}
