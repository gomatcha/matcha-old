package animate

import (
	"github.com/overcyn/mochi"
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
			}
		}
	}()
	return t
}

func (t *Timing) Stop() {
	t.ticker.Stop()
}

type value struct {
	chans     []chan<- struct{}
	mu        *sync.Mutex
	value     interface{}
	notifiers []mochi.Notifier
	onNotify  func()
}

func (v *value) Notify(c chan<- struct{}) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.chans = append(v.chans, c)
}

func (v *value) Unnotify(c chan<- struct{}) {
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

type UnitNotifier interface {
	mochi.Notifier
	Value() float64
}

type unitNotifier struct {
	watcher      UnitNotifier
	interpolater UnitInterpolater
}

func (w *unitNotifier) Notify(c chan<- struct{}) {
	w.watcher.Notify(c)
}

func (w *unitNotifier) Unnotify(c chan<- struct{}) {
	w.watcher.Unnotify(c)
}

func (w *unitNotifier) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type UnitValue struct {
	v *value
}

func (v *UnitValue) Notify(c chan<- struct{}) {
	v.v.Notify(c)
}

func (v *UnitValue) Unnotify(c chan<- struct{}) {
	v.v.Unnotify(c)
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

func InterpolatedUnit(w UnitNotifier, l UnitInterpolater) UnitNotifier {
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

type FloatNotifier interface {
	mochi.Notifier
	Value() float64
}

type floatNotifier struct {
	watcher      FloatNotifier
	interpolater FloatInterpolater
}

func (w *floatNotifier) Notify(c chan<- struct{}) {
	w.watcher.Notify(c)
}

func (w *floatNotifier) Unnotify(c chan<- struct{}) {
	w.watcher.Unnotify(c)
}

func (w *floatNotifier) Value() float64 {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type FloatValue struct {
	v         *value
	notifiers []FloatNotifier
	done      []chan struct{}
}

func NewFloatValue() *FloatValue {
	return &FloatValue{
		v: &value{},
	}
}

func (v *FloatValue) Watch(n FloatNotifier) {
	v.v.mu.Lock()
	defer v.v.mu.Unlock()

	c := make(chan struct{})
	done := make(chan struct{})
	v.notifiers = append(v.notifiers, n)
	v.done = append(v.done, done)
	n.Notify(c)

	go func() {
	loop:
		for {
			select {
			case <-c:
				v.Set(n.Value())
			case <-done:
				n.Unnotify(c)
				break loop
			}
		}
	}()
}

func (v *FloatValue) Unwatch(n FloatNotifier) {
	v.v.mu.Lock()
	defer v.v.mu.Unlock()

	notifiers := []FloatNotifier{}
	done := []chan struct{}{}
	for idx, i := range v.notifiers {
		if i == n {
			v.done[idx] <- struct{}{}
		} else {
			notifiers = append(notifiers, i)
			done = append(done, v.done[idx])
		}
	}
}

func (v *FloatValue) Notify(c chan<- struct{}) {
	v.v.Notify(c)
}

func (v *FloatValue) Unnotify(c chan<- struct{}) {
	v.v.Unnotify(c)
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

func FloatInterpolate(w UnitNotifier, l FloatInterpolater) FloatNotifier {
	return &floatNotifier{
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
// func onMount() {
// 	// I want multiple timings to be able to update the animate.UnitValue()
// 	// And I want multiple things to be able to watch animate.UnitValue
// 	unitN := animate.UnitInterpolate(animate.NewTiming(10), animate.LinearEase{})
// 	floatN = animate.FloatInterpolate(w, animate.FloatLerp{0, 10})

// 	value := &animate.FloatValue{}
// 	value.Watch(floatN)
// 	value.Unwatch(floatN)
// 	// value.Notify

// 	view.Rect.X(value)
// }
