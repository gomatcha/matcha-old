package settings

import (
	"github.com/overcyn/matcha/layout/table"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/basicview"
	"github.com/overcyn/matcha/view/scrollview"
)

type BluetoothView struct {
	*view.Embed
	app *App
}

func NewBluetoothView(ctx *view.Context, key string, app *App) *BluetoothView {
	if v, ok := ctx.Prev(key).(*BluetoothView); ok {
		return v
	}
	return &BluetoothView{Embed: view.NewEmbed(ctx.NewId(key)), app: app}
}

func (v *BluetoothView) Build(ctx *view.Context) *view.Model {
	l := &table.Layout{}
	chlds := []view.View{}

	scrollChild := basicview.New(ctx, "a")
	scrollChild.Layouter = l
	scrollChild.Children = chlds

	scrollView := scrollview.New(ctx, "b")
	scrollView.ContentView = scrollChild

	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}
