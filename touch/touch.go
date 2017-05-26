package touch

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
)

func init() {
	view.RegisterVisitor(&visitor{})
}

type key struct{}

var Key = key{}

type visitor struct {
}

func (v *visitor) Visit(m *view.Model) {
	rs, _ := m.Values[Key].([]Recognizer)

	pbRecognizers := &pb.RecognizerList{}
	for _, i := range rs {
		str, msg := i.encodeProtobuf()

		var pbAny *any.Any
		if a, err := ptypes.MarshalAny(msg); err == nil {
			pbAny = a
			continue
		}

		pbRecognizers.RecognizerNames = append(pbRecognizers.RecognizerNames, str)
		pbRecognizers.Recognizers = append(pbRecognizers.Recognizers, pbAny)
	}

	m.NativeValues["github.com/overcyn/mochi/touch"] = pbRecognizers
}

type Recognizer interface {
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

func (r *TapRecognizer) encodeProtobuf() (string, proto.Message) {
	return "github.com/overcyn/mochi/touch TapRecognizer", &pb.TapRecognizer{
		Count: int64(r.Count),
	}
}

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
