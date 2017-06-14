package touch

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/pb/touch"
	"github.com/overcyn/mochi/view"
)

func init() {
	view.RegisterMiddleware(&Middleware{})
}

type key struct{}
type _idKey struct{}

var Key = key{}
var idKey = _idKey{}

type Middleware struct {
	maxId int64
}

func (r *Middleware) Build(ctx *view.Context, next *view.Model) {
	var prevIds map[int64]Recognizer
	if prevModel := ctx.PrevModel(); prevModel != nil && prevModel.Values != nil {
		prevIds, _ = prevModel.Values[idKey].(map[int64]Recognizer)
	}

	ids := map[int64]Recognizer{}

	var rs []Recognizer
	rs1, ok := next.Values[Key]
	if ok {
		rs2, ok := rs1.([]Recognizer)
		if !ok {
			fmt.Println("Value for recognizer key is not a []touch.Recognizer")
		} else {
			rs = rs2
		}
	}

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

	// Add new list back to next.
	if next.Values == nil {
		next.Values = map[interface{}]interface{}{}
	}
	next.Values[idKey] = ids

	// Serialize into protobuf.
	pbRecognizers := &pb.RecognizerList{}
	allFuncs := map[int64]interface{}{}
	for k, v := range ids {
		msg, funcs := v.MarshalProtobuf(ctx)
		pbAny, err := ptypes.MarshalAny(msg)
		if err != nil {
			continue
		}

		pbRecognizer := &pb.Recognizer{
			Id:         k,
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
		next.NativeFuncs = map[int64]interface{}{}
	}
	for k, v := range allFuncs {
		next.NativeFuncs[k] = v
	}
}

type Recognizer interface {
	MarshalProtobuf(ctx *view.Context) (proto.Message, map[int64]interface{})
	Equal(Recognizer) bool
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

func (e *TapEvent) UnmarshalProtobuf(pbevent *pb.TapEvent) error {
	t, _ := ptypes.Timestamp(pbevent.Timestamp)
	e.Timestamp = t
	e.Position.UnmarshalProtobuf(pbevent.Position)
	return nil
}

type TapRecognizer struct {
	Count   int
	OnTouch func(*TapEvent)
}

func (r *TapRecognizer) Equal(a Recognizer) bool {
	b, ok := a.(*TapRecognizer)
	if !ok {
		return false
	}
	return r.Count == b.Count
}

func (r *TapRecognizer) MarshalProtobuf(ctx *view.Context) (proto.Message, map[int64]interface{}) {
	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()

		pbevent := &pb.TapEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		event := &TapEvent{}
		if err := event.UnmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}

		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &pb.TapRecognizer{
			Count:          int64(r.Count),
			RecognizedFunc: funcId,
		}, map[int64]interface{}{
			funcId: f,
		}
}

type EventKind int

const (
	EventKindPossible EventKind = iota
	EventKindChanged
	EventKindFailed
	EventKindRecognized

	// EventKindPossible // discrete events
	// EventKindRecognized
	// EventKindFailed
)

type PressEvent struct {
	Kind      EventKind
	Timestamp time.Time
	Position  layout.Point
	Duration  time.Duration
}

func (e *PressEvent) UnmarshalProtobuf(pbevent *pb.PressEvent) error {
	d, err := ptypes.Duration(pbevent.Duration)
	if err != nil {
		return err
	}
	t, err := ptypes.Timestamp(pbevent.Timestamp)
	if err != nil {
		return err
	}
	e.Kind = EventKind(pbevent.Kind)
	e.Timestamp = t
	e.Position.UnmarshalProtobuf(pbevent.Position)
	e.Duration = d
	return nil
}

type PressRecognizer struct {
	MinDuration time.Duration
	OnTouch     func(e *PressEvent)
}

func (r *PressRecognizer) Equal(a Recognizer) bool {
	b, ok := a.(*PressRecognizer)
	if !ok {
		return false
	}
	return r.MinDuration == b.MinDuration
}

func (r *PressRecognizer) MarshalProtobuf(ctx *view.Context) (proto.Message, map[int64]interface{}) {
	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()

		event := &PressEvent{}
		pbevent := &pb.PressEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		if err := event.UnmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}
		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &pb.PressRecognizer{
			MinDuration: ptypes.DurationProto(r.MinDuration),
			FuncId:      funcId,
		}, map[int64]interface{}{
			funcId: f,
		}
}

type ButtonEvent struct {
	Timestamp time.Time
	Inside    bool
	Kind      EventKind
}

func (e *ButtonEvent) UnmarshalProtobuf(pbevent *touch.ButtonEvent) error {
	t, err := ptypes.Timestamp(pbevent.Timestamp)
	if err != nil {
		return err
	}
	e.Timestamp = t
	e.Inside = pbevent.Inside
	e.Kind = EventKind(pbevent.Kind)
	return nil
}

type ButtonRecognizer struct {
	OnTouch       func(e *ButtonEvent)
	IgnoresScroll bool
}

func (r *ButtonRecognizer) Equal(a Recognizer) bool {
	_, ok := a.(*ButtonRecognizer)
	if !ok {
		return false
	}
	return true
}

func (r *ButtonRecognizer) MarshalProtobuf(ctx *view.Context) (proto.Message, map[int64]interface{}) {
	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()

		event := &ButtonEvent{}
		pbevent := &touch.ButtonEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		if err := event.UnmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}

		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &touch.ButtonRecognizer{
			OnEvent:       funcId,
			IgnoresScroll: r.IgnoresScroll,
		}, map[int64]interface{}{
			funcId: f,
		}
}

// type PanEvent struct {
// 	Timestamp time.Time
// 	Position  layout.Point
// 	Velocity  layout.Point
// }

// type PanRecognizer struct {
// 	key      interface{}
// 	OnBegin  func(e *PanEvent)
// 	OnEnd    func(e *PanEvent)
// 	OnCancel func(e *PanEvent)
// 	OnChange func(e *PanEvent)
// }

// type SwipeRecognizer struct {
// }

// type PinchRecognizer struct {
// }

// type EdgePanRecognizer struct {
// }

// type RotationGesture struct {
// }
