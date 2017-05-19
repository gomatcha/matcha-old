package touch

import (
	"time"

	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view"
)

// func init() {
// 	view.RegisterMarshaller(func(m *view.Model) (string, []byte) {
// 		value, ok := m.Values[key{}]
// 		if !ok {
// 			return "", nil
// 		}
// 		recognizers, ok := value.([]Recognizer)
// 		if !ok {
// 			panic("Value is not a []Recognizer")
// 		}
// 		_ = recognizers
// 		return "github.com/overcyn/mochi/touch", nil
// 	})
// }

type key struct{}

func Key() interface{} {
	return key{}
}

type Recognizer interface {
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

type TapRecognizer struct {
	Count          int
	RecognizedFunc func(e *TapEvent)
}

func NewTapRecognizer(ctx *view.Context, key interface{}) *TapRecognizer {
	return nil
}

type PressEvent struct {
	Timestamp time.Time
	Position  layout.Point
	Duration  time.Duration
}

type PressRecognizer struct {
	MinDuration   time.Duration
	BeganFunc     func(e *PressEvent)
	EndFunc       func(e *PressEvent)
	CancelledFunc func(e *PressEvent)
	ChangedFunc   func(e *PressEvent)
}

type PanEvent struct {
	Timestamp time.Time
	Position  layout.Point
	Velocity  layout.Point
}

type PanRecognizer struct {
	BeganFunc     func(e *PanEvent)
	EndFunc       func(e *PanEvent)
	CancelledFunc func(e *PanEvent)
	ChangedFunc   func(e *PanEvent)
}

// type SwipeRecognizer struct {
// }

// type PinchRecognizer struct {
// }

// type EdgePanRecognizer struct {
// }

// type RotationGesture struct {
// }
