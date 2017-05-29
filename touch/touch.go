package touch

import (
	"reflect"
	"sync"
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
}

type key struct{}

var Key = key{}

type cacheKey struct {
	viewId mochi.Id
	key    interface{}
}

type Root struct {
	mu      sync.Mutex
	prevIds map[mochi.Id]map[int64]Recognizer
	ids     map[mochi.Id]map[int64]Recognizer
	maxId   int64
}

func (r *Root) BeforeBuild() {
	r.prevIds = r.ids
	r.ids = map[mochi.Id]map[int64]Recognizer{}
}

func (r *Root) Build(ctx *view.Context, next *view.Model) {
	id := ctx.Id()
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
	allFuncs := map[int64]reflect.Value{}
	for k, v := range ids {
		str, msg, funcs := v.EncodeProtobuf(ctx)
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
		for k2, v2 := range funcs {
			allFuncs[k2] = v2
		}
	}

	if next.NativeValues == nil {
		next.NativeValues = map[string]proto.Message{}
	}
	next.NativeValues["github.com/overcyn/mochi/touch"] = pbRecognizers

	if next.NativeFuncs == nil {
		next.NativeFuncs = map[int64]reflect.Value{}
	}
	for k, v := range allFuncs {
		next.NativeFuncs[k] = v
	}
}

func (r *Root) AfterBuild() {
	r.prevIds = nil
}

type Recognizer interface {
	EncodeProtobuf(ctx *view.Context) (string, proto.Message, map[int64]reflect.Value)
	Equal(Recognizer) bool
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

type TapRecognizer struct {
	Count          int
	RecognizedFunc func()
}

func (r *TapRecognizer) Equal(a Recognizer) bool {
	b, ok := a.(*TapRecognizer)
	if !ok {
		return false
	}
	return r.Count == b.Count
}

func (r *TapRecognizer) EncodeProtobuf(ctx *view.Context) (string, proto.Message, map[int64]reflect.Value) {
	funcId := ctx.NewFuncId()

	return "github.com/overcyn/mochi/touch TapRecognizer", &pb.TapRecognizer{
			Count:          int64(r.Count),
			RecognizedFunc: funcId,
		}, map[int64]reflect.Value{
			funcId: reflect.ValueOf(r.RecognizedFunc),
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
