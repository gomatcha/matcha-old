package stacknavigator

import (
	"github.com/overcyn/mochi/view"
)

type StackNavigator struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *StackNavigator {
	if v, ok := ctx.Prev(key).(*StackNavigator); ok {
		return v
	}
	return &StackNavigator{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *StackNavigator) Build(ctx *view.Context) *view.Model {
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
