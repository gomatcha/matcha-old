package animate

import (
	"time"

	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/internal"
	"gomatcha.io/matcha/view"
)

type Value struct {
	value     float64
	batch     comm.GroupNotifier
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

func (v *Value) SetValue(val float64) {
	v.setValue(val)
	if v.animation != nil {
		v.animation.cancel()
	}
}

func (v *Value) setValue(val float64) {
	v.value = val
	v.batch.Update()
}

func (v *Value) Run(a Animation) (cancelFunc func()) {
	if v.animation != nil {
		v.animation.cancel()
	}

	start := time.Now()
	an := &animation{animation: a, ticker: internal.NewTicker(time.Hour * 99), value: v}
	an.tickerId = an.ticker.Notify(func() {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()
		if an.cancelled {
			return
		}

		d := time.Now().Sub(start)

		v.setValue(a.Tick(d))
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
	Duration() time.Duration
	Tick(time.Duration) float64
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
	ticker     *internal.Ticker
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
	Start float64
	End   float64
	Ease  FloatInterpolater
	Dur   time.Duration // Duration
}

func (a *Basic) Duration() time.Duration {
	return a.Dur
}

func (a *Basic) Tick(t time.Duration) float64 {
	if a.Dur == 0 {
		return a.End
	}
	ratio := float64(t) / float64(a.Dur)
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

// func Reverse(a animation) animation {
// }

// func Delay(a animation) animation {
// }

// func Repeat(a animation) animation {
// }
