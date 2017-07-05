package settings

import (
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/stackscreen"
	"gomatcha.io/matcha/view/switchview"
)

type WifiStore struct {
	store.Store
	enabled            bool
	networks           []*WifiNetworkStore
	currentNetworkSSID string
}

func NewWifiStore() *WifiStore {
	s := &WifiStore{}
	s.SetEnabled(true)

	n1 := &WifiNetworkStore{}
	n1.SetNetwork(WifiNetwork{
		SSID: "XfinityWifi",
	})

	n2 := &WifiNetworkStore{}
	n2.SetNetwork(WifiNetwork{
		SSID: "Bluestone",
	})

	n3 := &WifiNetworkStore{}
	n3.SetNetwork(WifiNetwork{
		SSID: "Starbucks",
	})

	n4 := &WifiNetworkStore{}
	n4.SetNetwork(WifiNetwork{
		SSID: "FastMesh Wifi",
	})

	s.SetNetworks([]*WifiNetworkStore{n1, n2, n3, n4})
	s.SetCurrentNetworkSSID(n4.Network().SSID)
	return s
}

func (s *WifiStore) SetEnabled(v bool) {
	s.enabled = v
	s.Update()
}

func (s *WifiStore) Enabled() bool {
	return s.enabled
}

func (s *WifiStore) SetNetworks(ns []*WifiNetworkStore) {
	for _, i := range s.networks {
		s.Delete(i.Network().SSID)
	}
	for _, i := range ns {
		s.Set(i.Network().SSID, i)
	}
	s.networks = ns
	s.Update()
}

func (s *WifiStore) Networks() []*WifiNetworkStore {
	return s.networks
}

func (s *WifiStore) SetCurrentNetworkSSID(v string) {
	s.currentNetworkSSID = v
	s.Update()
}

func (s *WifiStore) CurrentNetworkSSID() string {
	return s.currentNetworkSSID
}

type WifiNetworkStore struct {
	store.Store
	network WifiNetwork
}

func (s *WifiNetworkStore) Network() WifiNetwork {
	return s.network
}

func (s *WifiNetworkStore) SetNetwork(v WifiNetwork) {
	s.network = v
	s.Update()
}

type WifiNetwork struct {
	SSID   string
	Locked bool
	Signal int
}

type WifiView struct {
	*view.Embed
	app       *App
	wifiStore *WifiStore
}

func NewWifiView(ctx *view.Context, key string, app *App, wifiStore *WifiStore) *WifiView {
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
		l.Add(spacer, nil)

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
			var currentNetwork WifiNetwork
			for _, i := range v.wifiStore.Networks() {
				if i.Network().SSID == currentSSID {
					currentNetwork = i.Network()
					break
				}
			}

			cell2 := NewBasicCell(ctx, "current")
			cell2.Title = currentNetwork.SSID
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
				network := i.Network()
				if network.SSID != v.wifiStore.CurrentNetworkSSID() {
					cell := NewBasicCell(ctx, "network"+network.SSID)
					cell.Title = network.SSID
					cell.OnTap = func() {
						v.wifiStore.Lock()
						defer v.wifiStore.Unlock()

						v.wifiStore.SetCurrentNetworkSSID(network.SSID)
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
