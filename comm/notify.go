package comm

import (
	"image/color"

	"github.com/overcyn/mochi"
)

type ColorNotifier interface {
	mochi.Notifier
	Value() color.Color
}

type InterfaceNotifier interface {
	mochi.Notifier
	Value() interface{}
}

type BoolNotifier interface {
	mochi.Notifier
	Value() bool
}

type IntNotifier interface {
	mochi.Notifier
	Value() int
}

type UintNotifier interface {
	mochi.Notifier
	Value() uint
}

type Int64Notifier interface {
	mochi.Notifier
	Value() int64
}

type Uint64Notifier interface {
	mochi.Notifier
	Value() uint64
}

type Float64Notifier interface {
	mochi.Notifier
	Value() float64
}

type StringNotifier interface {
	mochi.Notifier
	Value() string
}

type BytesNotifier interface {
	mochi.Notifier
	Value() []byte
}
