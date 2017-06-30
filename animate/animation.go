package animate

import (
	"time"

	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/internal"
	"gomatcha.io/matcha/view"
)

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
	an := &animation{animation: a, onComplete: onComplete, ticker: internal.NewTicker(time.Hour * 99), value: v}
	an.tickerId = an.ticker.Notify(func() {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()
		if an.cancelled {
			return
		}

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
	start    float64
	end      float64
	ease     FloatInterpolater
	duration time.Duration
	time     time.Duration
}

func (a *Basic) SetStart(v float64) {
	a.start = v
}

func (a *Basic) Start() float64 {
	return a.start
}

func (a *Basic) SetEnd(v float64) {
	a.end = v
}

func (a *Basic) End() float64 {
	return a.end
}

func (a *Basic) SetEase(v FloatInterpolater) {
	a.ease = v
}

func (a *Basic) Ease() FloatInterpolater {
	return a.ease
}

func (a *Basic) SetDuration(v time.Duration) {
	a.duration = v
}

func (a *Basic) Duration() time.Duration {
	return a.duration
}

func (a *Basic) Time() time.Duration {
	return a.time
}

func (a *Basic) SetTime(t time.Duration) {
	a.time = t
}

func (a *Basic) Value() float64 {
	if a.duration == 0 {
		return a.end
	}
	ratio := float64(a.time) / float64(a.duration)
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1 {
		ratio = 1
	}
	if a.ease != nil {
		ratio = a.ease.Interpolate(ratio)
	}
	return a.start + ratio*(a.end-a.start)
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
