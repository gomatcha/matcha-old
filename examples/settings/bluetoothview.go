package settings

import (
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/switchview"
)

type BluetoothStore struct {
	store.Node
	bluetooth Bluetooth
}

func NewBluetoothStore() *BluetoothStore {
	n1 := NewBluetoothDeviceStore("JBL Charge 3")
	n2 := NewBluetoothDeviceStore("Kevin's AirPods")
	n3 := NewBluetoothDeviceStore("Kevin's Apple Watch")
	n4 := NewBluetoothDeviceStore("Honda Fit")

	s := &BluetoothStore{}
	s.SetBluetooth(Bluetooth{
		Enabled: true,
		Devices: []*BluetoothDeviceStore{n1, n2, n3, n4},
	})
	return s
}

func (s *BluetoothStore) Bluetooth() Bluetooth {
	return s.bluetooth
}

func (s *BluetoothStore) SetBluetooth(v Bluetooth) {
	for _, i := range s.bluetooth.Devices {
		ssid := i.Device().SSID
		s.Delete(ssid)
	}
	for _, i := range v.Devices {
		ssid := i.Device().SSID
		s.Set(ssid, i)
	}

	s.bluetooth = v
	s.Signal()
}

type Bluetooth struct {
	Enabled bool
	Devices []*BluetoothDeviceStore
}

type BluetoothDeviceStore struct {
	store.Node
	device BluetoothDevice
	ssid   string
}

func NewBluetoothDeviceStore(ssid string) *BluetoothDeviceStore {
	return &BluetoothDeviceStore{
		ssid: ssid,
		device: BluetoothDevice{
			SSID: ssid,
		},
	}
}

func (s *BluetoothDeviceStore) Device() BluetoothDevice {
	return s.device
}

func (s *BluetoothDeviceStore) SetDevice(v BluetoothDevice) {
	s.device = v
	s.device.SSID = s.ssid // Don't allow the SSID to change.
	s.Signal()
}

type BluetoothDevice struct {
	SSID      string
	Connected bool
}

type BluetoothView struct {
	view.Embed
	app       *App
	bluetooth *BluetoothStore
}

func NewBluetoothView(ctx *view.Context, key string, app *App) *BluetoothView {
	if v, ok := ctx.Prev(key).(*BluetoothView); ok {
		return v
	}
	app.Store.Lock()
	defer app.Store.Unlock()
	v := &BluetoothView{
		Embed:     ctx.NewEmbed(key),
		app:       app,
		bluetooth: app.Store.BluetoothStore(),
	}
	v.Subscribe(v.bluetooth)
	return v
}

func (v *BluetoothView) Build(ctx *view.Context) view.Model {
	v.bluetooth.Lock()
	defer v.bluetooth.Unlock()
	bt := v.bluetooth.Bluetooth()

	l := &table.Layouter{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer, nil)

		switchView := switchview.New(ctx, "switch")
		switchView.Value = bt.Enabled
		switchView.OnValueChange = func(value bool) {
			v.bluetooth.Lock()
			defer v.bluetooth.Unlock()

			bt := v.bluetooth.Bluetooth()
			bt.Enabled = value
			v.bluetooth.SetBluetooth(bt)
		}

		cell1 := NewBasicCell(ctx, "wifi")
		cell1.Title = "Bluetooth"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}
	if bt.Enabled {
		ctx := ctx.WithPrefix("2")
		group := []view.View{}

		spacer := NewSpacerHeader(ctx, "spacer")
		spacer.Title = "My Devices"
		l.Add(spacer, nil)

		for _, i := range bt.Devices {
			deviceStore := i
			d := deviceStore.Device()
			cell := NewBasicCell(ctx, "network"+d.SSID)
			cell.Title = d.SSID
			if d.Connected {
				cell.Subtitle = "Connected"
			} else {
				cell.Subtitle = "Not Connected"
			}
			cell.OnTap = func() {
				deviceStore.Lock()
				defer deviceStore.Unlock()

				d := deviceStore.Device()
				d.Connected = !d.Connected
				deviceStore.SetDevice(d)
			}
			group = append(group, cell)
		}

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i, nil)
		}
	}

	scrollView := scrollview.New(ctx, "b")
	scrollView.ContentLayouter = l
	scrollView.ContentChildren = l.Views()

	return view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}
