package settings

import (
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/stackscreen"
	"gomatcha.io/matcha/view/switchview"
)

type WifiController struct {
	comm.Storer
	store              *comm.AsyncStore
	enabled            bool
	networks           []*WifiNetwork
	currentNetworkSSID string
}

func NewWifiStore() *WifiController {
	st := &comm.AsyncStore{}
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
	comm.Storer
	store  *comm.AsyncStore
	ssid   string
	locked bool
	signal int
}

func NewWifiNetwork() *WifiNetwork {
	st := &comm.AsyncStore{}
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

func NewWifiView(ctx *view.Context, key string, app *App, wifiStore *WifiController) *WifiView {
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

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l, nil.Add(spacer)

		switchView := switchview.New(ctx, "switch")
		switchView.Value = v.wifiStore.Enabled()
		switchView.OnValueChange = func(value bool) {
			v.wifiStore.Lock()
			defer v.wifiStore.Unlock()

			v.wifiStore.SetEnabled(value)
		}

		cell1 := NewBasicCell(ctx, "wifi")
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

			cell2 := NewBasicCell(ctx, "current")
			cell2.Title = currentNetwork.SSID()
			group = append(group, cell2)
		}

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}

	if v.wifiStore.Enabled() {
		{
			ctx := ctx.WithPrefix("2")
			group := []view.View{}

			spacer := NewSpacerHeader(ctx, "spacer")
			spacer.Title = "Choose a Network..."
			l.Add(spacer, nil)

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

			cell1 := NewBasicCell(ctx, "other")
			cell1.Title = "Other..."
			group = append(group, cell1)

			for _, i := range AddSeparators(ctx, group) {
				l.Add(i, nil)
			}
		}
		{
			ctx := ctx.WithPrefix("3")
			group := []view.View{}

			spacer := NewSpacer(ctx, "spacer")
			l.Add(spacer, nil)

			switchView := switchview.New(ctx, "switch")
			cell1 := NewBasicCell(ctx, "join")
			cell1.Title = "Ask to Join Networks"
			cell1.AccessoryView = switchView
			group = append(group, cell1)

			for _, i := range AddSeparators(ctx, group) {
				l.Add(i, nil)
			}
		}
		{
			spacer := NewSpacerDescription(ctx, "spacerDescr")
			spacer.Description = "Known networks will be joined automatically. If no known networks are available, you will have to manually join a network."
			l.Add(spacer, nil)
		}
	}

	scrollChild := basicview.New(ctx, "scrollChild")
	scrollChild.Layouter = l
	scrollChild.Children = l.Views()

	scrollView := scrollview.New(ctx, "scroll")
	scrollView.ContentView = scrollChild

	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

func (v *WifiView) StackBar(ctx *view.Context) *stackscreen.Bar {
	return &stackscreen.Bar{Title: "Wi-Fi"}
}
