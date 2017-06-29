package animate

import (
	"time"

	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/internal"
	"github.com/overcyn/matcha/view"
)

type Ticker struct {
	ticker *internal.Ticker
}

func NewTicker(duration time.Duration) *Ticker {
	return &Ticker{
		ticker: internal.NewTicker(duration),
	}
}

func (t *Ticker) Notify(f func()) comm.Id {
	return t.ticker.Notify(f)
}

func (t *Ticker) Unnotify(id comm.Id) {
	t.ticker.Unnotify(id)
}

func (t *Ticker) Value() float64 {
	return t.ticker.Value()
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
}

type Value struct {
	value     float64
	batch     comm.BatchNotifier
	animation *animation
}

func (v *Value) Notify(f func()) comm.Id {
	return v.batch.Notify(f)
}

func (v *Value) Unnotify(id comm.Id) {
	v.batch.Unnotify(id)
}

func (v *Value) Value() float64 {
	return v.value
}

func (v *Value) setValue(val float64) {
	v.value = val
	v.batch.Update()
}

func (v *Value) Run(a Animation, onComplete func()) (cancelFunc func()) {
	if v.animation != nil {
		v.animation.cancel()
	}

	start := time.Now()
	an := &animation{animation: a, onComplete: onComplete, ticker: NewTicker(time.Hour * 99)}
	an.tickerId = an.ticker.Notify(func() {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()

		d := time.Now().Sub(start)
		a.SetTime(d)
		v.setValue(a.Value())
		if d > a.Duration() {
			an.cancel()
		}
	})
	v.animation = an

	return func() {
		an.cancel()
	}
}

type Animation interface {
	SetTime(time.Duration)
	Duration() time.Duration
	Value() float64
}

// type Animation2D interface {
// 	Duration() time.Duration
// 	SetTime(*time.Duration)
// 	Values() [2]float64
// }

// type AnimationND interface {
// 	Duration() time.Duration
// 	SetTime(*time.Duration)
// 	Values() []float64
// }

type animation struct {
	cancelled  bool
	animation  Animation
	ticker     *Ticker
	tickerId   comm.Id
	onComplete func()
	value      *Value
}

func (a *animation) cancel() {
	if a.cancelled {
		return
	}

	a.ticker.Unnotify(a.tickerId)
	a.value.animation = nil
	if a.onComplete != nil {
		a.onComplete()
	}
	a.cancelled = true
}

type Basic struct {
	Start        float64
	End          float64
	Ease         FloatInterpolater
	TimeInterval time.Duration
	time         time.Duration
}

func (a *Basic) Duration() time.Duration {
	return a.TimeInterval
}

func (a *Basic) SetTime(t time.Duration) {
	a.time = t
}

func (a *Basic) Value() float64 {
	if a.TimeInterval == 0 {
		return a.End
	}
	ratio := float64(a.time) / float64(a.TimeInterval)
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1 {
		ratio = 1
	}
	if a.Ease != nil {
		ratio = a.Ease.Interpolate(ratio)
	}
	return a.Start + ratio*(a.End-a.Start)
}

// type Spring struct {
// 	Start     float64
// 	End       float64
// 	Velocity  float64
// 	Stiffness float64
// 	Dampening float64
// }

// func (a *Spring) Duration() time.Duration {
// 	return time.Duration(1)
// }

// func (a *Spring) SetTime(t time.Duration) {
// }

// func (a *Spring) Value() float64 {
// 	return 0
// }

// type Decay struct {
// 	Start        float64
// 	End          float64
// 	Velocity     float64 // units/second
// 	Deceleration float64
// }

// func (a *Decay) Duration() time.Duration {
// 	return time.Duration(1)
// }

// func (a *Decay) SetTime(t time.Duration) {
// }

// func (a *Decay) Value() float64 {
// 	return 0
// }

// func Run(a Animation, v *Value, onComplete func()) (cancelFunc func()) {
// 	if v.animation != nil {
// 		v.animation.cancel()
// 	}

// 	start := time.Now()
// 	an := &animation{animation: a, onComplete: onComplete, ticker: NewTicker(time.Hour * 99)}
// 	an.tickerId = an.ticker.Notify(func() {
// 		view.MainMu.Lock()
// 		defer view.MainMu.Unlock()

// 		d := time.Now().Sub(start)
// 		a.SetTime(d)
// 		v.setValue(a.Value())
// 		if d > a.Duration() {
// 			an.cancel()
// 		}
// 	})
// 	v.animation = an

// 	return func() {
// 		an.cancel()
// 	}
// }

// func Reverse(a animation) animation {
// }

// func Delay(a animation) animation {
// }

// func Repeat(a animation) animation {
// }

// cancelFunc := animate.RunMultiple(func(t *animate.Timing) {
// 	a := &animate.Basic{Start:animate.Current, End:3.0, Duration: 4.0}
// 	a := &animate.Bounce{Start:animate.Current, End:3.0, Velocity: 2.0}

// 	t.Run(a, value, nil)
// 	t.Run(b, value, nil)

// 	t = t.After(a.Duration())

// 	a2 := animate.Basic{Start:1.0, End:3.0, Duration: 4.0, Value:Value}
// 	t.Run(a, value, nil)
// }, nil)
