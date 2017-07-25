package settings

import (
	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/env"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/touch"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/imageview"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/stackview"
	"gomatcha.io/matcha/view/switchview"
)

type WifiView struct {
	view.Embed
	app  *App
	wifi *Wifi
}

func NewWifiView(ctx *view.Context, key string, app *App) *WifiView {
	if v, ok := ctx.Prev(key).(*WifiView); ok {
		return v
	}
	app.Store.Lock()
	defer app.Store.Unlock()
	v := &WifiView{
		Embed: ctx.NewEmbed(key),
		app:   app,
		wifi:  app.Store.WifiStore(),
	}
	v.Subscribe(v.wifi)
	return v
}

func (v *WifiView) Build(ctx *view.Context) view.Model {
	v.wifi.Lock()
	defer v.wifi.Unlock()

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		switchView := switchview.New(ctx, "switch")
		switchView.Value = v.wifi.Enabled()
		switchView.OnValueChange = func(value bool) {
			v.wifi.Lock()
			defer v.wifi.Unlock()
			v.wifi.SetEnabled(!v.wifi.Enabled())
		}

		cell1 := NewBasicCell(ctx, "wifi")
		cell1.Title = "Wi-Fi"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		if v.wifi.CurrentSSID() != "" && v.wifi.Enabled() {
			cell2 := NewBasicCell(ctx, "current")
			cell2.Title = v.wifi.CurrentSSID()
			group = append(group, cell2)
		}

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}

	if v.wifi.Enabled() {
		{
			ctx := ctx.WithPrefix("2")
			group := []view.View{}

			spacer := NewSpacerHeader(ctx, "spacer")
			spacer.Title = "Choose a Network..."
			l.Add(spacer, nil)

			for _, i := range v.wifi.Networks() {
				network := i
				ssid := network.SSID()

				// Don't show the current network in this list.
				if ssid == v.wifi.CurrentSSID() {
					continue
				}

				info := NewInfoButton(ctx, "networkbutton"+ssid)
				info.PaintStyle = &paint.Style{BackgroundColor: colornames.Red}
				info.OnPress = func() {
					v.app.Stack.Push(NewWifiNetworkView(nil, "", v.app, network))
				}

				cell := NewBasicCell(ctx, "network"+ssid)
				cell.Title = ssid
				cell.AccessoryView = info
				cell.OnTap = func() {
					v.wifi.Lock()
					defer v.wifi.Unlock()

					v.wifi.SetCurrentSSID(ssid)
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

			spacer := NewSpacer(ctx, "spacer")
			l.Add(spacer, nil)

			switchView := switchview.New(ctx, "switch")
			cell1 := NewBasicCell(ctx, "join")
			cell1.Title = "Ask to Join Networks"
			cell1.AccessoryView = switchView

			for _, i := range AddSeparators(ctx, []view.View{cell1}) {
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
	app     *App
	network *WifiNetwork
}

func NewWifiNetworkView(ctx *view.Context, key string, app *App, network *WifiNetwork) *WifiNetworkView {
	if v, ok := ctx.Prev(key).(*WifiNetworkView); ok {
		return v
	}
	v := &WifiNetworkView{
		Embed:   ctx.NewEmbed(key),
		app:     app,
		network: network,
	}
	v.Subscribe(network)
	return v
}

func (v *WifiNetworkView) Build(ctx *view.Context) view.Model {
	v.network.Lock()
	defer v.network.Unlock()
	props := v.network.Properties()

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "forget")
		cell1.Title = "Forget This Network"

		for _, i := range AddSeparators(ctx, []view.View{cell1}) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("2")

		spacer := NewSpacerHeader(ctx, "spacer")
		spacer.Title = "IP Address"
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "ip")
		cell1.Title = "IP Address"
		cell1.Subtitle = props.IPAddress

		cell2 := NewBasicCell(ctx, "subnet")
		cell2.Title = "Subnet Mask"
		cell2.Subtitle = props.SubnetMask

		cell3 := NewBasicCell(ctx, "router")
		cell3.Title = "Router"
		cell3.Subtitle = props.Router

		cell4 := NewBasicCell(ctx, "dns")
		cell4.Title = "DNS"
		cell4.Subtitle = props.DNS

		cell5 := NewBasicCell(ctx, "clientid")
		cell5.Title = "Client ID"
		cell5.Subtitle = props.ClientID

		for _, i := range AddSeparators(ctx, []view.View{cell1, cell2, cell3, cell4, cell5}) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("3")

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "renew")
		cell1.Title = "Renew Lease"

		for _, i := range AddSeparators(ctx, []view.View{cell1}) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("4")

		spacer := NewSpacerHeader(ctx, "spacer")
		spacer.Title = "HTTP Proxy"
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "renew")
		cell1.Title = "Renew Lease"

		for _, i := range AddSeparators(ctx, []view.View{cell1}) {
			l.Add(i, nil)
		}
	}
	{
		ctx := ctx.WithPrefix("5")

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		cell1 := NewBasicCell(ctx, "manage")
		cell1.Title = "Manage This Network"

		for _, i := range AddSeparators(ctx, []view.View{cell1}) {
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
	v.network.Lock()
	defer v.network.Unlock()

	return &stackview.Bar{
		Title: v.network.SSID(),
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
