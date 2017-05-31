package stacknavigator

import (
	"github.com/overcyn/mochi/view"
)

type StackNavigator struct {
	*view.Embed
	Views []view.View
}

func New(ctx *view.Context, key interface{}) *StackNavigator {
	if v, ok := ctx.Prev(key).(*StackNavigator); ok {
		return v
	}
	return &StackNavigator{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (n *StackNavigator) Build(ctx *view.Context) *view.Model {
	return &view.Model{}
}

func (n *StackNavigator) Push(v view.View) {

}

func (n *StackNavigator) Pop() {
}
