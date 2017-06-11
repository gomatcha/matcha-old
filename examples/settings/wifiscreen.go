package settings

/*
import (
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/switchview"
)

type WifiController struct {
	// store.Storer
	store          *store.Store
	enabled        bool
	networks       map[int64]*WifiNetwork
	currentNetwork int
}

func NewWifiStore() *WifiController {
	return nil
	// st := &store.Store{}
	// s := &WifiController{Storer: st, store: st}
	// s.SetEnabled(true)

	// net1 := s.AddNetwork(1)
	// net1.SetId(1)
	// net1.SetName("XfinityWifi")

	// net2 := s.AddNetwork(2)
	// net2.SetId(2)
	// net2.SetName("Bluestone")
	// return s
}

func (s *WifiController) SetEnabled(v bool) {
	s.enabled = v
	s.store.Update()
}

func (s *WifiController) Enabled() bool {
	return s.enabled
}

func (s *WifiController) AddNetwork(id int64) (*WifiNetwork, error) {
	st, err := s.store.Add(id)
	if err != nil {
		return err
	}
	return NewWifiNetwork(st), nil
}

func (s *WifiController) RemoveNetwork(id int64) error {
}

// func (s *WifiController) SetNetworks(ns map[int64]*WifiNetwork) {
// 	s.store.Update()
// 	for k, _ := range s.networks {
// 		s.store.Delete(k)
// 	}

// 	for k, v := range ns {
// 		s.store.Set(k, v.Store())
// 	}
// 	s.networks = ns
// }

func (s *WifiController) Networks() map[int64]*WifiNetwork {
	return s.networks
}

func (s *WifiController) SetCurrentNetwork(v int) {
	s.currentNetwork = v
	s.store.Update()
}

func (s *WifiController) CurrentNetwork() int {
	return s.currentNetwork
}

type WifiNetwork struct {
	// store.Storer
	store  *store.Store
	id     int64
	name   string
	locked bool
	signal int
}

func NewWifiNetwork() *WifiNetwork {
	st := &store.Store{}
	return &WifiController{Storer: st, store: st}
}

func (n *WifiNetwork) SetId(v int64) {
	n.id = v
	n.store.Update()
}

func (n *WifiNetwork) Id() int64 {
	return n.id
}

func (n *WifiNetwork) SetName(v string) {
	n.name = v
	n.store.Update()
}

func (n *WifiNetwork) Name() string {
	return n.name
}

func (n *WifiNetwork) SetLocked(v bool) {
	n.locked = v
	n.store.Update()
}

func (n *WifiNetwork) Locked() bool {
	return n.locked
}

func (n *WifiNetwork) SetSignal(v int) {
	n.signal = v
	n.store.Update()
}

func (n *WifiNetwork) Signal() int {
	return n.signal
}

type WifiView struct {
	*view.Embed
	app       *App
	wifiStore *WifiController
}

func NewWifiView(ctx *view.Context, key interface{}, app *App, wifiStore *WifiController) *WifiView {
	if v, ok := ctx.Prev(key).(*WifiView); ok {
		return v
	}
	v := &WifiView{
		Embed:     view.NewEmbed(ctx.NewId(key)),
		app:       app,
		wifiStore: wifiStore,
	}
	v.Subscribe(wifiStore.Store())
	return v
}

func (v *WifiView) Build(ctx *view.Context) *view.Model {
	v.wifiStore.Lock()
	defer v.wifiStore.Unlock()

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

		for _, v := range v.wifiStore.Networks() {
			cell := NewBasicCell(ctx, "network"+v.Name())
			cell.Title = v.Name()
			group = append(group, cell)
		}

		cell6 := NewBasicCell(ctx, "other")
		cell6.Title = "Other..."
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
		group := []view.View{}

		switchView := switchview.New(ctx, 10)
		cell1 := NewBasicCell(ctx, 11)
		cell1.Title = "Ask to Join Networks"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, "c", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
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
*/
