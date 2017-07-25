package settings

import "gomatcha.io/matcha/store"

type Store struct {
	store.Node
	airplaneMode   bool
	wifiStore      *Wifi
	bluetoothStore *BluetoothStore
}

func NewStore() *Store {
	st := &Store{
		wifiStore:      NewWifi(),
		bluetoothStore: NewBluetoothStore(),
		airplaneMode:   false,
	}
	st.Set("wifi", st.wifiStore)
	st.Set("bluetooth", st.bluetoothStore)
	return st
}

func (st *Store) WifiStore() *Wifi {
	return st.wifiStore
}

func (st *Store) BluetoothStore() *BluetoothStore {
	return st.bluetoothStore
}

func (st *Store) AirplaneMode() bool {
	return st.airplaneMode
}

func (st *Store) SetAirplaneMode(v bool) {
	st.airplaneMode = v

	st.wifiStore.SetEnabled(!st.airplaneMode)

	bt := st.bluetoothStore.Bluetooth()
	bt.Enabled = !st.airplaneMode
	st.bluetoothStore.SetBluetooth(bt)

	st.Signal()
}

type Wifi struct {
	store.Node
	enabled     bool
	currentSSID string
	networks    []*WifiNetwork
}

func NewWifi() *Wifi {
	n1 := NewWifiNetwork("XfinityWifi")
	n2 := NewWifiNetwork("Bluestone")
	n3 := NewWifiNetwork("Starbucks")
	n4 := NewWifiNetwork("FastMesh Wifi")

	s := &Wifi{}
	s.SetEnabled(true)
	s.SetCurrentSSID(n4.SSID())
	s.SetNetworks([]*WifiNetwork{n1, n2, n3, n4})
	return s
}

func (s *Wifi) CurrentSSID() string {
	return s.currentSSID
}

func (s *Wifi) SetCurrentSSID(v string) {
	s.currentSSID = v
	s.Signal()
}

func (s *Wifi) Enabled() bool {
	return s.enabled
}

func (s *Wifi) SetEnabled(v bool) {
	s.enabled = v
	s.Signal()
}

func (s *Wifi) Networks() []*WifiNetwork {
	return s.networks
}

func (s *Wifi) SetNetworks(n []*WifiNetwork) {
	for _, i := range s.networks {
		s.Delete(i.SSID())
	}
	for _, i := range n {
		s.Set(i.SSID(), i)
	}

	s.networks = n
	s.Signal()
}

type WifiNetwork struct {
	store.Node
	props WifiNetworkProperties
}

func NewWifiNetwork(ssid string) *WifiNetwork {
	return &WifiNetwork{
		props: WifiNetworkProperties{
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

func (n *WifiNetwork) SSID() string {
	return n.props.SSID
}

func (n *WifiNetwork) Properties() WifiNetworkProperties {
	return n.props
}

func (n *WifiNetwork) SetProperties(v WifiNetworkProperties) {
	n.props = v
	n.Signal()
}

type WifiNetworkProperties struct {
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
