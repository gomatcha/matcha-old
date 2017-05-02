package animation

import (
	// "github.com/overcyn/mochi"
	"sync"
	"time"
)

type Timing struct {
	ticker *time.Ticker
	after  <-chan time.Time
	value  UnitValue
	start  time.Time
}

func NewTiming(duration time.Duration, value UnitValue) *Timing {
	t := &Timing{
		ticker: time.NewTicker(time.Second / 30),
		after:  time.After(duration),
		value:  value,
		start:  time.Now(),
	}
	go func() {
	loop:
		for {
			select {
			case <-t.after:
				t.ticker.Stop()
				break loop
			case <-t.ticker.C:
				fraction := float64(time.Since(t.start) / duration)
				if fraction > 1 {
					fraction = 1
				} else if fraction < 0 {
					fraction = 0
				}
				t.value.Set(fraction)
			}
		}
	}()
	return t
}

func (t *Timing) Stop() {
	t.ticker.Stop()
}

type Watcher interface {
	Watch(chan<- struct{})
	Unwatch(chan<- struct{})
}

type value struct {
	chans []chan<- struct{}
	mu    *sync.Mutex
	value interface{}
}

func (v *value) Watch(c chan<- struct{}) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.chans = append(v.chans, c)
}

func (v *value) Unwatch(c chan<- struct{}) {
	v.mu.Lock()
	defer v.mu.Unlock()

	chans := make([]chan<- struct{}, 0, len(v.chans))
	for _, i := range chans {
		if i != c {
			chans = append(chans, i)
		}
	}
	v.chans = chans
}

func (v *value) Value() interface{} {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.value
}

func (v *value) Set(a interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.value = a
	for _, i := range v.chans {
		i <- struct{}{}
	}
}

type UnitWatcher interface {
	Watcher
	Value() float64
}

type unitWatcher struct {
	watcher      UnitWatcher
	interpolater UnitInterpolater
}

func (w *unitWatcher) Watch(c chan<- struct{}) {
	w.watcher.Watch(c)
}

func (w *unitWatcher) Unwatch(c chan<- struct{}) {
	w.watcher.Unwatch(c)
}

func (w *unitWatcher) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type UnitValue struct {
	v *value
}

func (v *UnitValue) Watch(c chan<- struct{}) {
	v.v.Watch(c)
}

func (v *UnitValue) Unwatch(c chan<- struct{}) {
	v.v.Unwatch(c)
}

func (v *UnitValue) Value() float64 {
	return v.v.Value().(float64)
}

func (v *UnitValue) Set(a float64) {
	if a > 1 {
		a = 1
	} else if a < 0 {
		a = 0
	}
	v.v.Set(a)
}

type UnitInterpolater interface {
	Interpolate(float64) float64
}

func InterpolatedUnit(w UnitWatcher, l UnitInterpolater) UnitWatcher {
	return &unitWatcher{
		watcher:      w,
		interpolater: l,
	}
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

type FloatWatcher interface {
	Watcher
	Value() float64
}

type floatWatcher struct {
	watcher      FloatWatcher
	interpolater FloatInterpolater
}

func (w *floatWatcher) Watch(c chan<- struct{}) {
	w.watcher.Watch(c)
}

func (w *floatWatcher) Unwatch(c chan<- struct{}) {
	w.watcher.Unwatch(c)
}

func (w *floatWatcher) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type FloatValue struct {
	v *value
}

func (v *FloatValue) Watch(c chan<- struct{}) {
	v.v.Watch(c)
}

func (v *FloatValue) Unwatch(c chan<- struct{}) {
	v.v.Unwatch(c)
}

func (v *FloatValue) Value() float64 {
	return v.v.Value().(float64)
}

func (v *FloatValue) Set(a float64) {
	v.v.Set(a)
}

type FloatInterpolater interface {
	Interpolate(float64) float64
}

func FloatInterpolated(w UnitWatcher, l FloatInterpolater) FloatWatcher {
	return &floatWatcher{
		watcher:      w,
		interpolater: l,
	}
}

type FloatLerp struct {
	start, end float64
}

func (f FloatLerp) Interpolate(a float64) float64 {
	return f.start + (f.end-f.start)*a
}

// value := animate.UnitValue()
// timing := animate.NewTiming(10, value)
func onMount() {
	// I want multiple timings to be able to update the animate.UnitValue()
	// And I want multiple things to be able to watch animate.UnitValue
	source := animate.FloatInterpolate(animate.NewTiming(10), animate.FloatLerp{0, 10})
	value := &animate.FloatValue{}
	// value.Watch

	view.Rect.X(value)
}

// type IntSource interface {
// 	Source
// 	Value() int
// }

// type PointSource interface {
// 	Source
// 	Value() mochi.Point
// }

// type RectSource interface {
// 	Source
// 	Value() mochi.Rect
// }

// type InterfaceSource interface {
// 	Source
// 	Value() mochi.Rect
// }

// type ColorSource interface {
// 	Source
// 	Value() color.Color
// }

// type ColorInterpolater interface {
//     Interpolate(float64) color.Color
// }

// // func InterpolateColor(s UnitSource, start, end color.Color) ColorSource {
// // 	return nil
// // }
// //

// curved := animate.UnitInterpolated(unitSource, animate.EaseIn)
// source := animate.FloatInterpolatedWatcher(unitWatcher, &animate.FloatLerp{start: 0, end: 10})
// source := animate.ColorInterpolated(unitSource, &animate.HSLInterpolater{})

// // colorSource := animate.WithColorInterpolater(animate.WithCurve(source, animate.EaseInEaseOutCurve), &animate.HSLInterpolater{start: color.Red, end: color.Blue})
// // animate.InterpolateFloat(source, 5, 10)

// // animate.NewFloatSource(unitSource, animate.FloatInterpolater{start: start, end: end})

// // animate.InterpolateFloat(source, animate.FloatInterpolater{start: start, end: end})
// // animateFloat.Interpolate(source, 5, 10)
// // animateFloat.Interpolater(source, animateFloat.Interpolater{start: start, end: end})

// // animate.InterpolatePoint(source, animate.PointInterpolater{start: startPoint, end: endPoint})
// // animate.InterpolateColor(source, start, end)
// // source.WithInterpolater(animate.FloatInterpolater{start, end})

// // type IntInterpolaterFunc func(float64) int

// // type IntInterpolater interface {
// // 	Interpolate(a float64) int
// // }

// // func WithIntInterpolater(s UnitSource, lerp IntInterpolater) IntSource {
// // }

// // func WithIntInterpolater2(s UnitSource, start, end int) IntSource {
// // }

// // type ColorInterpolater interface {
// // 	Interpolate(a float64) color.Color
// // }

// // func WithColorInterpolater(s UnitSource, lerp ColorInterpolater) ColorSource {
// // }

// // func ColorInterpolater {
// // }

// // With easing function
// func WithCurve(s UnitSource, f func(float64) float64) UnitSource {
// }

// func EaseInCurve(a float64) float64 {
// 	return a
// }

// func EaseOutCurve(a float64) float64 {
// 	return a
// }

// func EaseInEaseOutCurve(a float64) float64 {
// 	return a
// }

// func BezierCurve(a float64, c1x, c1y, c2x, c2y float64) float64 {
// 	return a
// }

// func LinearCurve(a float64) float64 {
// 	return a
// }
