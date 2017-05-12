package animate

import (
	"github.com/overcyn/mochi"
)

// type value struct {
// 	chans     []chan struct{}
// 	mu        *sync.Mutex
// 	value     interface{}
// 	notifiers []mochi.Notifier
// 	done      []chan struct{}
// }

// func (v *value) Notify() chan struct{} {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	c := make(chan struct{})
// 	v.chans = append(v.chans, c)
// 	return c
// }

// func (v *value) Unnotify(c chan struct{}) {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	chans := make([]chan struct{}, 0, len(v.chans))
// 	for _, i := range chans {
// 		if i != c {
// 			chans = append(chans, i)
// 		}
// 	}
// 	v.chans = chans
// }

// func (v *value) Value() interface{} {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	return v.value
// }

// func (v *value) Set(a interface{}) {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	v.value = a
// 	for _, i := range v.chans {
// 		i <- struct{}{}
// 		<-i
// 	}
// }

// func (v *value) Watch(n mochi.Notifier, f func() interface{}) {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	done := make(chan struct{})
// 	v.notifiers = append(v.notifiers, n)
// 	v.done = append(v.done, done)
// 	c := n.Notify()

// 	// setup a go routine waiting for notifications from n.
// 	go func() {
// 	loop:
// 		for {
// 			select {
// 			case <-c:
// 				v.Set(f())
// 				c <- struct{}{}
// 			case <-done:
// 				n.Unnotify(c)
// 				break loop
// 			}
// 		}
// 	}()
// }

// func (v *value) Unwatch(n mochi.Notifier) {
// 	v.mu.Lock()
// 	defer v.mu.Unlock()

// 	notifiers := []mochi.Notifier{}
// 	done := []chan struct{}{}
// 	for idx, i := range v.notifiers {
// 		if i == n {
// 			v.done[idx] <- struct{}{}
// 		} else {
// 			notifiers = append(notifiers, i)
// 			done = append(done, v.done[idx])
// 		}
// 	}
// }

type UnitNotifier interface {
	mochi.Notifier
	Value() float64
}

type unitNotifier struct {
	watcher      UnitNotifier
	interpolater UnitInterpolater
}

func (w *unitNotifier) Notify() chan struct{} {
	return w.watcher.Notify()
}

func (w *unitNotifier) Unnotify(c chan struct{}) {
	w.watcher.Unnotify(c)
}

func (w *unitNotifier) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

// type UnitValue struct {
// 	v *value
// }

// func (v *UnitValue) Watch(n UnitNotifier) {
// 	v.v.Watch(n, func() interface{} {
// 		return n.Value()
// 	})
// }

// func (v *UnitValue) Unwatch(n UnitNotifier) {
// 	v.Unwatch(n)
// }

// func (v *UnitValue) Notify() chan struct{} {
// 	return v.v.Notify()
// }

// func (v *UnitValue) Unnotify(c chan struct{}) {
// 	v.v.Unnotify(c)
// }

// func (v *UnitValue) Value() float64 {
// 	return v.v.Value().(float64)
// }

// func (v *UnitValue) Set(a float64) {
// 	if a > 1 {
// 		a = 1
// 	} else if a < 0 {
// 		a = 0
// 	}
// 	v.v.Set(a)
// }

type UnitInterpolater interface {
	Interpolate(float64) float64
}

func UnitInterpolate(w UnitNotifier, l UnitInterpolater) UnitNotifier {
	return &unitNotifier{
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

type FloatInterpolater interface {
	Interpolate(float64) float64
}

func FloatInterpolate(w UnitNotifier, l FloatInterpolater) mochi.Float64Notifier {
	return &floatInterpolater{
		watcher:      w,
		interpolater: l,
	}
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
