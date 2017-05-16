package animate

import (
	"github.com/overcyn/mochi"
)

type FloatInterpolater interface {
	Interpolate(float64) float64
}

func FloatInterpolate(w mochi.Float64Notifier, l FloatInterpolater) mochi.Float64Notifier {
	return &floatInterpolater{
		watcher:      w,
		interpolater: l,
	}
}

type floatInterpolater struct {
	watcher      mochi.Float64Notifier
	interpolater FloatInterpolater
}

func (w *floatInterpolater) Notify() chan struct{} {
	return w.watcher.Notify()
}

func (w *floatInterpolater) Unnotify(c chan struct{}) {
	w.watcher.Unnotify(c)
}

func (w *floatInterpolater) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type LinearEase struct {
}

func (e *LinearEase) Interpolate(a float64) float64 {
	return a
}

type PolyInEase struct {
}

func (e *PolyInEase) Interpolate(a float64) float64 {
	return a
}

type PolyOutEase struct {
}

func (e *PolyOutEase) Interpolate(a float64) float64 {
	return a
}

type PolyInOutEase struct {
}

func (e *PolyInOutEase) Interpolate(a float64) float64 {
	return a
}

type FloatLerp struct {
	Start, End float64
}

func (f FloatLerp) Interpolate(a float64) float64 {
	return f.Start + (f.End-f.Start)*a
}

// value := animate.UnitValue()
// Ticker := animate.NewTicker(10, value)
// func onMount() {
// 	// I want multiple Tickers to be able to update the animate.UnitValue()
// 	// And I want multiple things to be able to watch animate.UnitValue
// 	unitN := animate.UnitInterpolate(animate.NewTicker(10), animate.LinearEase{})
// 	floatN = animate.FloatInterpolate(w, animate.FloatLerp{0, 10})

// 	value := &animate.FloatValue{}
// 	value.Watch(floatN)
// 	value.Unwatch(floatN)
// 	// value.Notify

// 	view.Rect.X(value)
// }
