package settings

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
	store.Storer
	store              *store.AsyncStore
	enabled            bool
	networks           []*WifiNetwork
	currentNetworkSSID string
}

func NewWifiStore() *WifiController {
	st := &store.AsyncStore{}
	s := &WifiController{Storer: st, store: st}
	s.SetEnabled(true)

	n1 := NewWifiNetwork()
	n1.SetSSID("XfinityWifi")

	n2 := NewWifiNetwork()
	n2.SetSSID("Bluestone")

	n3 := NewWifiNetwork()
	n3.SetSSID("Starbucks")

	n4 := NewWifiNetwork()
	n4.SetSSID("FastMesh Wifi")

	s.SetNetworks([]*WifiNetwork{n1, n2, n3, n4})
	s.SetCurrentNetworkSSID(n4.SSID())
	return s
}

func (s *WifiController) SetEnabled(v bool) {
	s.enabled = v
	s.store.Update()
}

func (s *WifiController) Enabled() bool {
	return s.enabled
}

func (s *WifiController) SetNetworks(ns []*WifiNetwork) {
	s.store.Update()

	for _, i := range s.networks {
		s.store.Delete(i.SSID())
	}
	for _, i := range ns {
		s.store.Set(i.SSID(), i)
	}
	s.networks = ns
}

func (s *WifiController) Networks() []*WifiNetwork {
	return s.networks
}

func (s *WifiController) SetCurrentNetworkSSID(v string) {
	s.currentNetworkSSID = v
	s.store.Update()
}

func (s *WifiController) CurrentNetworkSSID() string {
	return s.currentNetworkSSID
}

type WifiNetwork struct {
	store.Storer
	store  *store.AsyncStore
	ssid   string
	locked bool
	signal int
}

func NewWifiNetwork() *WifiNetwork {
	st := &store.AsyncStore{}
	return &WifiNetwork{Storer: st, store: st}
}

func (n *WifiNetwork) SetSSID(v string) {
	n.ssid = v
	n.store.Update()
}

func (n *WifiNetwork) SSID() string {
	return n.ssid
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
	v.Subscribe(wifiStore)
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
		switchView.Value = v.wifiStore.Enabled()
		switchView.OnValueChange = func(sv *switchview.View) {
			v.wifiStore.Lock()
			defer v.wifiStore.Unlock()

			v.wifiStore.SetEnabled(sv.Value)
		}

		group := []view.View{}
		cell1 := NewBasicCell(ctx, 0)
		cell1.Title = "Wi-Fi"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		currentSSID := v.wifiStore.CurrentNetworkSSID()
		if currentSSID != "" && v.wifiStore.Enabled() {
			var currentNetwork *WifiNetwork
			for _, i := range v.wifiStore.Networks() {
				if i.SSID() == currentSSID {
					currentNetwork = i
					break
				}
			}

			cell2 := NewBasicCell(ctx, 1)
			cell2.Title = currentNetwork.SSID()
			group = append(group, cell2)
		}

		for _, i := range AddSeparators(ctx, "a", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}

	if v.wifiStore.Enabled() {
		{
			spacer := NewSpacer(ctx, "spacer2")
			chlds = append(chlds, spacer)
			l.Add(spacer)
		}
		{
			group := []view.View{}

			for _, i := range v.wifiStore.Networks() {
				if i.SSID() != v.wifiStore.CurrentNetworkSSID() {
					ssid := i.SSID()
					cell := NewBasicCell(ctx, "network"+i.SSID())
					cell.Title = i.SSID()
					cell.OnTap = func() {
						v.wifiStore.Lock()
						defer v.wifiStore.Unlock()

						v.wifiStore.SetCurrentNetworkSSID(ssid)
					}
					group = append(group, cell)
				}
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
