package animate

import "github.com/overcyn/mochi/comm"

type FloatInterpolater interface {
	Interpolate(float64) float64
}

func FloatInterpolate(w comm.Float64Notifier, l FloatInterpolater) comm.Float64Notifier {
	return &floatInterpolater{
		watcher:      w,
		interpolater: l,
	}
}

type floatInterpolater struct {
	watcher      comm.Float64Notifier
	interpolater FloatInterpolater
}

func (w *floatInterpolater) Notify(f func()) comm.Id {
	return w.watcher.Notify(f)
}

func (w *floatInterpolater) Unnotify(id comm.Id) {
	w.watcher.Unnotify(id)
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
