package animate

import (
	"fmt"
	"math"

	"github.com/gomatcha/matcha/comm"
	"golang.org/x/mobile/exp/sprite/clock"
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

var (
	DefaultEase      FloatInterpolater = CubicBezierEase{0.25, 0.1, 0.25, 1}
	DefaultInEase    FloatInterpolater = CubicBezierEase{0.42, 0, 1, 1}
	DefaultOutEase   FloatInterpolater = CubicBezierEase{0, 0, 0.58, 1}
	DefaultInOutEase FloatInterpolater = CubicBezierEase{0.42, 0, 0.58, 1}
)

type CubicBezierEase struct {
	X0, Y0, X1, Y1 float64
}

func (e CubicBezierEase) Interpolate(a float64) float64 {
	f := clock.CubicBezier(float32(e.X0), float32(e.Y0), float32(e.X1), float32(e.Y1))
	t := f(0, 100000, clock.Time(a*100000))
	fmt.Println(a, t)
	return float64(t) / 100000
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
	ExpIn  float64
	ExpOut float64
}

func (e PolyInOutEase) Interpolate(a float64) float64 {
	if a < 0.5 {
		return math.Pow(a, e.ExpIn)
	} else {
		return 1 - math.Pow(1-a, e.ExpOut)
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
