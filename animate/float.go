package animate

import (
	"math"

	"github.com/overcyn/matcha/comm"
)

type FloatInterpolater interface {
	Interpolate(float64) float64
}

func floatInterpolate(w comm.Float64Notifier, l FloatInterpolater) comm.Float64Notifier {
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

func (e LinearEase) Interpolate(a float64) float64 {
	return a
}

func (e LinearEase) Notifier(a comm.Float64Notifier) comm.Float64Notifier {
	return floatInterpolate(a, e)
}

type PolyInEase struct {
	Exp float64
}

func (e PolyInEase) Interpolate(a float64) float64 {
	return math.Pow(a, e.Exp)
}

func (e PolyInEase) Notifier(a comm.Float64Notifier) comm.Float64Notifier {
	return floatInterpolate(a, e)
}

type PolyOutEase struct {
	Exp float64
}

func (e PolyOutEase) Interpolate(a float64) float64 {
	return 1 - math.Pow(1-a, e.Exp)
}

func (e PolyOutEase) Notifier(a comm.Float64Notifier) comm.Float64Notifier {
	return floatInterpolate(a, e)
}

type PolyInOutEase struct {
	Exp float64
}

func (e PolyInOutEase) Interpolate(a float64) float64 {
	if a < 0.5 {
		return math.Pow(a, e.Exp)
	} else {
		return 1 - math.Pow(1-a, e.Exp)
	}
}

func (e PolyInOutEase) Notifier(a comm.Float64Notifier) comm.Float64Notifier {
	return floatInterpolate(a, e)
}

type FloatLerp struct {
	Start, End float64
}

func (f FloatLerp) Interpolate(a float64) float64 {
	return f.Start + (f.End-f.Start)*a
}

func (e FloatLerp) Notifier(a comm.Float64Notifier) comm.Float64Notifier {
	return floatInterpolate(a, e)
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
