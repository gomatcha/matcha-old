package touch

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochibridge"
)

func init() {
	// view.RegisterVisitor(&visitor{})
	mochibridge.RegisterFunc("github.com/overcyn/mochi/touch TapRecognizer.Recognized", func(id int64) {

	})
}

type visitorKey struct{}
type key struct{}

var VisitorKey = visitorKey{}
var Key = key{}

type cacheKey struct {
	viewId mochi.Id
	key    interface{}
}

type Root struct {
	prevRecognizers map[int64]Recognizer
	prevIds         map[cacheKey]int64
	recognizers     map[int64]Recognizer
	ids             map[cacheKey]int64
	maxId           int64
}

func (r *Root) BeforeBuild() {
	r.prevRecognizers = r.recognizers
	r.prevIds = r.ids
	r.recognizers = map[int64]Recognizer{}
	r.ids = map[cacheKey]int64{}
}

func (r *Root) Build(id mochi.Id, prev, next *view.Model) {
	rs, _ := next.Values[Key].([]Recognizer)
	pbRecognizers := &pb.RecognizerList{}
	for _, i := range rs {
		str, msg := i.EncodeProtobuf()

		var pbAny *any.Any
		if a, err := ptypes.MarshalAny(msg); err == nil {
			pbAny = a
			continue
		}

		pbRecognizers.RecognizerNames = append(pbRecognizers.RecognizerNames, str)
		pbRecognizers.Recognizers = append(pbRecognizers.Recognizers, pbAny)

		// Add to recognizer list.
		r.recognizers[i.Id()] = i
	}
	next.NativeValues["github.com/overcyn/mochi/touch"] = pbRecognizers
}

func (r *Root) AfterBuild() {
	keys := map[int64]cacheKey{}
	for k, v := range r.ids {
		keys[v] = k
	}
	for k, v := range r.prevIds {
		keys[v] = k
	}

	ids := map[cacheKey]int64{}
	for k := range r.recognizers {
		ids[keys[k]] = k
	}
	r.ids = ids
}

func (r *Root) Prev(id mochi.Id, key interface{}) Recognizer {
	prevId := r.ids[cacheKey{viewId: id, key: key}]
	return r.recognizers[prevId]
}

func (r *Root) NewId(id mochi.Id, key interface{}) int64 {
	r.maxId += 1
	r.ids[cacheKey{viewId: id, key: key}] = r.maxId
	return r.maxId
}

type Recognizer interface {
	Id() int64
	EncodeProtobuf() (string, proto.Message)
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

type TapRecognizer struct {
	id             int64
	Count          int
	RecognizedFunc func(e *TapEvent)
}

func NewTapRecognizer(ctx *view.Context, key interface{}) *TapRecognizer {
	// root := ctx.Visitor(VisitorKey).(visitor)

	// r, ok := root.Prev(ctx.Id(), key).(*TapRecognizer)
	// if !ok {
	// 	r = &TapRecognizer{
	// 		id: root.NewId(ctx.Id(), key),
	// 	}
	// }
	// return root
	return nil
}

func (r *TapRecognizer) Id() int64 {
	return r.id
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
