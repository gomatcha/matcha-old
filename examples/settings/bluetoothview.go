package settings

import (
	"github.com/gomatcha/matcha/layout/table"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/view"
	"github.com/gomatcha/matcha/view/basicview"
	"github.com/gomatcha/matcha/view/scrollview"
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
	l := &table.Layouter{}
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
