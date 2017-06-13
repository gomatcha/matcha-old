package mochi

import (
	"image/color"
)

type Id int64

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
}

type ColorNotifier interface {
	Notifier
	Value() color.Color
}

type InterfaceNotifier interface {
	Notifier
	Value() interface{}
}

type BoolNotifier interface {
	Notifier
	Value() bool
}

type IntNotifier interface {
	Notifier
	Value() int
}

type UintNotifier interface {
	Notifier
	Value() uint
}

type Int64Notifier interface {
	Notifier
	Value() int64
}

type Uint64Notifier interface {
	Notifier
	Value() uint64
}

type Float64Notifier interface {
	Notifier
	Value() float64
}

type StringNotifier interface {
	Notifier
	Value() string
}

type BytesNotifier interface {
	Notifier
	Value() []byte
}

func NotifyFunc(n Notifier, f func()) (done chan struct{}) {
	c := n.Notify()
	done = make(chan struct{})
	go func() {
	loop:
		for {
			select {
			case <-c:
				f()
				c <- struct{}{}
			case <-done:
				break loop
			}
		}
	}()
	return done
}
