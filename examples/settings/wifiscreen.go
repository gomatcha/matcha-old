package settings

import (
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/switchview"
)

type WifiView struct {
	*view.Embed
	app *App
}

func NewWifiView(ctx *view.Context, key interface{}, app *App) *WifiView {
	if v, ok := ctx.Prev(key).(*WifiView); ok {
		return v
	}
	return &WifiView{Embed: view.NewEmbed(ctx.NewId(key)), app: app}
}

func (v *WifiView) Build(ctx *view.Context) *view.Model {
	l := &table.Layout{}
	chlds := []view.View{}
	{
		spacer := NewSpacer(ctx, "spacer1")
		chlds = append(chlds, spacer)
		l.Add(spacer)
	}
	{
		switchView := switchview.New(ctx, 6)

		group := []view.View{}
		cell1 := NewBasicCell(ctx, 0)
		cell1.Title = "Wi-Fi"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, 1)
		cell2.Title = "FastMesh Wifi"
		group = append(group, cell2)

		for _, i := range AddSeparators(ctx, "a", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}
	{
		spacer := NewSpacer(ctx, "spacer2")
		chlds = append(chlds, spacer)
		l.Add(spacer)
	}
	{
		group := []view.View{}
		cell3 := NewBasicCell(ctx, 2)
		cell3.Title = "HOME-ABCD"
		group = append(group, cell3)

		cell4 := NewBasicCell(ctx, 3)
		cell4.Title = "xfinitywifi"
		group = append(group, cell4)

		cell5 := NewBasicCell(ctx, 4)
		cell5.Title = "Starbucks Wifi"
		group = append(group, cell5)

		cell6 := NewBasicCell(ctx, 5)
		cell6.Title = "Other"
		group = append(group, cell6)

		for _, i := range AddSeparators(ctx, "b", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}
	{
		spacer := NewSpacer(ctx, "spacer3")
		chlds = append(chlds, spacer)
		l.Add(spacer)
	}
	{
		switchView := switchview.New(ctx, 10)

		cell1 := NewBasicCell(ctx, 11)
		cell1.Title = "Ask to Join Networks"
		cell1.AccessoryView = switchView
		chlds = append(chlds, cell1)
		l.Add(cell1)
	}

	scrollChild := basicview.New(ctx, -1)
	scrollChild.Layouter = l
	scrollChild.Children = chlds

	scrollView := scrollview.New(ctx, -2)
	scrollView.ContentView = scrollChild

	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}
