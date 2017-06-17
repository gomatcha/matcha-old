package settings

import (
	"github.com/overcyn/matcha/layout/table"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/basicview"
	"github.com/overcyn/matcha/view/scrollview"
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
