package touch

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochibridge"
)

func init() {
	view.RegisterMiddleware(&Root{})
	mochibridge.RegisterFunc("github.com/overcyn/mochi/touch TapRecognizer.Recognized", func(viewId, touchId int64) {
		fmt.Println("tapRecognized")
	})
}

type key struct{}

var Key = key{}

type cacheKey struct {
	viewId mochi.Id
	key    interface{}
}

type Root struct {
	prevIds map[mochi.Id]map[int64]Recognizer
	ids     map[mochi.Id]map[int64]Recognizer
	maxId   int64
}

func (r *Root) BeforeBuild() {
	r.prevIds = r.ids
	r.ids = map[mochi.Id]map[int64]Recognizer{}
}

func (r *Root) Build(id mochi.Id, prev, next *view.Model) {
	prevIds, _ := r.prevIds[id]
	ids := map[int64]Recognizer{}
	rs, _ := next.Values[Key].([]Recognizer)

	// Diff prev and next recognizers
	for _, i := range rs {
		found := false
		for k, v := range prevIds {
			// Check that the id has not already been used.
			if _, ok := ids[k]; ok {
				continue
			}

			// Check that the recognizers are equal.
			if !i.Equal(v) {
				continue
			}

			ids[k] = i
			found = true
		}

		// Generate a new id if we don't have a previous one.
		if !found {
			r.maxId += 1
			ids[r.maxId] = i
		}
	}

	if len(ids) == 0 {
		return
	}

	// Add new list back to root.
	r.ids[id] = ids

	// Serialize into protobuf.
	pbRecognizers := &pb.RecognizerList{}
	for k, v := range ids {
		str, msg := v.EncodeProtobuf()
		pbAny, err := ptypes.MarshalAny(msg)
		if err != nil {
			continue
		}

		pbRecognizer := &pb.Recognizer{
			Id:         k,
			Name:       str,
			Recognizer: pbAny,
		}
		pbRecognizers.Recognizers = append(pbRecognizers.Recognizers, pbRecognizer)
	}

	if next.NativeValues == nil {
		next.NativeValues = map[string]proto.Message{}
	}
	next.NativeValues["github.com/overcyn/mochi/touch"] = pbRecognizers
}

func (r *Root) AfterBuild() {
}

type Recognizer interface {
	EncodeProtobuf() (string, proto.Message)
	Equal(Recognizer) bool
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

type TapRecognizer struct {
	Count          int
	RecognizedFunc func(e *TapEvent)
}

func (r *TapRecognizer) Equal(a Recognizer) bool {
	b, ok := a.(*TapRecognizer)
	if !ok {
		return false
	}
	return r.Count == b.Count
}

func (r *TapRecognizer) EncodeProtobuf() (string, proto.Message) {
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
