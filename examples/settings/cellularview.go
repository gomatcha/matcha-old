package settings

import (
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
	"gomatcha.io/matcha/view/scrollview"
)

type CellularView struct {
	*view.Embed
	app *App
}

func NewCellularView(ctx *view.Context, key string, app *App) *CellularView {
	if v, ok := ctx.Prev(key).(*CellularView); ok {
		return v
	}
	return &CellularView{Embed: view.NewEmbed(ctx.NewId(key)), app: app}
}

func (v *CellularView) Build(ctx *view.Context) *view.Model {
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
