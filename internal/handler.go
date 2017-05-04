// package mochi

// import (
// 	"time"
// )

// type Handler interface {
// 	Handle(event Event) bool
// }

// type Event interface {
// 	Bubbling() bool
// }

// type OnTapHandler struct {
// 	HandleFunc func(event interface{})
// }

// func (h *OnTapHandler) Handle(event interface{}) {
// 	// switch i := e.(type) {
// 	// case TouchEvent:

// 	// }
// }

// type Type int

// const (
// 	Direct Type = iota
// 	Indirect
// 	Stylus
// )

// type TouchKind int

// const (
// 	Begin TouchKind = iota
// 	Move
// 	Stationary
// 	End
// 	Cancel
// )

// type TouchEvent struct {
// 	// Window          Window
// 	Timestamp time.Time
// 	Location  Point
// 	Force     float64
// 	Kind      TouchKind
// 	Type      Type
// 	Radius    float64
// 	// RadiusTolerance float64
// 	// AltitudeAngle   float64
// 	// AzimuthalAngle  float64
// }

// type KeyKind int

// // const (
// // 	Begin KeyKind = iota
// // 	End
// // )

// type KeyCode int

// type Modifiers int

// const (
// 	Shift Modifiers = 1 << iota
// 	Control
// 	Alt
// 	Meta // Command
// )

// type KeyEvent struct {
// 	Timestamp time.Time
// 	Rune      rune
// 	Code      KeyCode
// 	Kind      KeyKind
// }
