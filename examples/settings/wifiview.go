package settings

import (
	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/env"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/touch"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/imageview"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/stackview"
	"gomatcha.io/matcha/view/switchview"
)

type WifiStore struct {
	store.Node
	wifi Wifi
}

func NewWifiStore() *WifiStore {
	n1 := NewWifiNetworkStore("XfinityWifi")
	n2 := NewWifiNetworkStore("Bluestone")
	n3 := NewWifiNetworkStore("Starbucks")
	n4 := NewWifiNetworkStore("FastMesh Wifi")

	s := &WifiStore{}
	s.SetWifi(Wifi{
		Enabled:     true,
		Networks:    []*WifiNetworkStore{n1, n2, n3, n4},
		CurrentSSID: n4.Network().SSID,
	})
	return s
}

func (s *WifiStore) SetWifi(v Wifi) {
	for _, i := range s.wifi.Networks {
		s.Delete(i.Network().SSID)
	}
	for _, i := range v.Networks {
		s.Set(i.Network().SSID, i)
	}

	s.wifi = v
	s.Signal()
}

func (s *WifiStore) Wifi() Wifi {
	return s.wifi
}

func (s *WifiStore) NetworkWithSSID(ssid string) *WifiNetworkStore {
	for _, i := range s.Wifi().Networks {
		if i.Network().SSID == ssid {
			return i
		}
	}
	return nil
}

type Wifi struct {
	Enabled     bool
	Networks    []*WifiNetworkStore
	CurrentSSID string
}

type WifiNetworkStore struct {
	store.Node
	ssid    string
	network WifiNetwork
}

func NewWifiNetworkStore(ssid string) *WifiNetworkStore {
	return &WifiNetworkStore{
		ssid: ssid,
		network: WifiNetwork{
			SSID:          ssid,
			IPAddress:     "10.0.1.25",
			SubnetMask:    "255.255.255.0",
			Router:        "10.0.1.1",
			DNS:           "10.0.1.1",
			SearchDomains: "hsd1.or.comcast.net.",
			ClientID:      "",
		},
	}
}

func (s *WifiNetworkStore) Network() WifiNetwork {
	return s.network
}

func (s *WifiNetworkStore) SetNetwork(v WifiNetwork) {
	s.network = v
	s.network.SSID = s.ssid // Don't allow the network's ssid to change.
	s.Signal()
}

type WifiNetwork struct {
	SSID          string
	Locked        bool
	Signal        int
	IPAddress     string
	SubnetMask    string
	Router        string
	DNS           string
	SearchDomains string
	ClientID      string
}

type WifiView struct {
	view.Embed
	app       *App
	wifiStore *WifiStore
}

func NewWifiView(ctx *view.Context, key string, app *App) *WifiView {
	if v, ok := ctx.Prev(key).(*WifiView); ok {
		return v
	}
	app.Store.Lock()
	defer app.Store.Unlock()
	v := &WifiView{
		Embed:     ctx.NewEmbed(key),
		app:       app,
		wifiStore: app.Store.WifiStore(),
	}
	v.Subscribe(v.wifiStore)
	return v
}

func (v *WifiView) Build(ctx *view.Context) view.Model {
	v.wifiStore.Lock()
	defer v.wifiStore.Unlock()
	wifi := v.wifiStore.Wifi()

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		switchView := switchview.New(ctx, "switch")
		switchView.Value = wifi.Enabled
		switchView.OnValueChange = func(value bool) {
			v.wifiStore.Lock()
			defer v.wifiStore.Unlock()

			wifi := v.wifiStore.Wifi()
			wifi.Enabled = value
			v.wifiStore.SetWifi(wifi)
		}

		cell1 := NewBasicCell(ctx, "wifi")
		cell1.Title = "Wi-Fi"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		if wifi.CurrentSSID != "" && wifi.Enabled {
			cell2 := NewBasicCell(ctx, "current")
			cell2.Title = wifi.CurrentSSID
			group = append(group, cell2)
		}

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}

	if wifi.Enabled {
		{
			ctx := ctx.WithPrefix("2")
			group := []view.View{}

			spacer := NewSpacerHeader(ctx, "spacer")
			spacer.Title = "Choose a Network..."
			l.Add(spacer, nil)

			for _, i := range wifi.Networks {
				networkStore := i
				network := i.Network()

				// Don't show the current network in this list.
				if network.SSID == wifi.CurrentSSID {
					continue
				}

				info := NewInfoButton(ctx, "networkbutton"+network.SSID)
				info.PaintStyle = &paint.Style{BackgroundColor: colornames.Red}
				info.OnPress = func() {
					v.app.Stack.Push(NewWifiNetworkView(nil, "", v.app, networkStore))
				}

				cell := NewBasicCell(ctx, "network"+network.SSID)
				cell.Title = network.SSID
				cell.AccessoryView = info
				cell.OnTap = func() {
					v.wifiStore.Lock()
					defer v.wifiStore.Unlock()

					wifi := v.wifiStore.Wifi()
					wifi.CurrentSSID = network.SSID
					v.wifiStore.SetWifi(wifi)
				}
				group = append(group, cell)
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

	scrollView := scrollview.New(ctx, "scroll")
	scrollView.ContentChildren = l.Views()
	scrollView.ContentLayouter = l

	return view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

func (v *WifiView) StackBar(ctx *view.Context) *stackview.Bar {
	return &stackview.Bar{Title: "Wi-Fi"}
}

type WifiNetworkView struct {
	view.Embed
	app          *App
	networkStore *WifiNetworkStore
}

func NewWifiNetworkView(ctx *view.Context, key string, app *App, network *WifiNetworkStore) *WifiNetworkView {
	if v, ok := ctx.Prev(key).(*WifiNetworkView); ok {
		return v
	}
	v := &WifiNetworkView{
		Embed:        ctx.NewEmbed(key),
		app:          app,
		networkStore: network,
	}
	v.Subscribe(network)
	return v
}

func (v *WifiNetworkView) Build(ctx *view.Context) view.Model {
	v.networkStore.Lock()
	defer v.networkStore.Unlock()
	network := v.networkStore.Network()

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "forget")
		cell1.Title = "Forget This Network"
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("2")
		group := []view.View{}

		spacer := NewSpacerHeader(ctx, "spacer")
		spacer.Title = "IP Address"
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "ip")
		cell1.Title = "IP Address"
		cell1.Subtitle = network.IPAddress
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, "subnet")
		cell2.Title = "Subnet Mask"
		cell2.Subtitle = network.SubnetMask
		group = append(group, cell2)

		cell3 := NewBasicCell(ctx, "router")
		cell3.Title = "Router"
		cell3.Subtitle = network.Router
		group = append(group, cell3)

		cell4 := NewBasicCell(ctx, "dns")
		cell4.Title = "DNS"
		cell4.Subtitle = network.DNS
		group = append(group, cell4)

		cell5 := NewBasicCell(ctx, "clientid")
		cell5.Title = "Client ID"
		cell5.Subtitle = network.ClientID
		group = append(group, cell5)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("3")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "renew")
		cell1.Title = "Renew Lease"
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("4")
		group := []view.View{}

		spacer := NewSpacerHeader(ctx, "spacer")
		spacer.Title = "HTTP Proxy"
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "renew")
		cell1.Title = "Renew Lease"
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("5")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "manage")
		cell1.Title = "Manage This Network"
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	spacer := NewSpacer(ctx, "spacer")
	l.Add(spacer, nil)

	scrollView := scrollview.New(ctx, "scroll")
	scrollView.ContentChildren = l.Views()
	scrollView.ContentLayouter = l

	return view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

func (v *WifiNetworkView) StackBar(*view.Context) *stackview.Bar {
	v.networkStore.Lock()
	defer v.networkStore.Unlock()
	return &stackview.Bar{
		Title: v.networkStore.Network().SSID,
	}
}

type InfoButton struct {
	view.Embed
	OnPress    func()
	PaintStyle *paint.Style
}

func NewInfoButton(ctx *view.Context, key string) *InfoButton {
	if v, ok := ctx.Prev(key).(*InfoButton); ok {
		return v
	}
	return &InfoButton{
		Embed: ctx.NewEmbed(key),
	}
}

func (v *InfoButton) Build(ctx *view.Context) view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.Width(35)
		s.Height(44)
	})

	img := imageview.New(ctx, "image")
	img.Image = env.MustLoadImage("Info")
	l.Add(img, func(s *constraint.Solver) {
		s.Width(22)
		s.Height(22)
		s.RightEqual(l.Right())
	})

	button := &touch.ButtonRecognizer{
		OnTouch: func(e *touch.ButtonEvent) {
			if e.Kind == touch.EventKindRecognized && v.OnPress != nil {
				v.OnPress()
			}
		},
	}

	return view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  v.PaintStyle,
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{button},
		},
	}
}
