package touch

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view"
)

// func init() {
// 	view.RegisterValueMarshaller(Key(), func(a interface{}) (string, proto.Message) {
// 		rs := a.([]Recognizer)
// 		return a
// 	})
// }

type key struct{}

func Key() interface{} {
	return key{}
}

// type context struct {
// 	prev []Recognizer
// }

// func newContext(vctx *view.Context) *context {
// 	prev, _ := ctx.PrevValue(key).([]Recognizer)
// 	return &context{
// 		prev: prev,
// 	}
// }

type Recognizer interface {
	id() int64
	encodeProtobuf() (string, proto.Message)
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

type TapRecognizer struct {
	key            interface{}
	Count          int
	RecognizedFunc func(e *TapEvent)
}

func NewTapRecognizer(ctx *view.Context, key interface{}) *TapRecognizer {
	return nil
}

// 	var r *TapRecognizer
// 	found := false

// 	prev, _ := ctx.PrevValue(key).([]Recognizer)
// 	for _, i := range prev {
// 		if tap, ok := i.(*TapRecognizer); ok && i.key() == key {
// 			r = tap
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		r = &TapRecognizer{}
// 	}
// 	return r
// }

type PressEvent struct {
	Timestamp time.Time
	Position  layout.Point
	Duration  time.Duration
}

type PressRecognizer struct {
	key           interface{}
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
	key           interface{}
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
