/*
Package constraint implements touch recognizers.

Create the touch recognizer in the Build function.

 tap := &touch.TapRecognizer{
 	Count: 1,
 	OnTouch: func(e *touch.TapEvent) {
 		// Respond to touch events. This callback occurs on main thread.
 		fmt.Println("view touched")
 	},
 }

Attach the recognizer to the view.

 return &view.Model{
 	Values: map[interface{}]interface{}{
 		touch.Key:[]touch.Recognizer{tap},
 	},
 }
*/
package touch

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"gomatcha.io/matcha/internal"
	"gomatcha.io/matcha/layout"
	pbtouch "gomatcha.io/matcha/pb/touch"
	"gomatcha.io/matcha/view"
)

func init() {
	internal.RegisterMiddleware(func() interface{} { return &middleware{} })
}

type key struct{}
type _idKey struct{}

var Key = key{}
var idKey = _idKey{}

type middleware struct {
	maxId int64
}

var maxFuncId int64 = 0

// NewFuncId generates a new func identifier for serialization.
func newFuncId() int64 {
	return atomic.AddInt64(&maxFuncId, 1)
}

func (r *middleware) Build(ctx *view.Context, next *view.Model) {
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
			if !i.equal(v) {
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
	pbRecognizers := &pbtouch.RecognizerList{}
	allFuncs := map[string]interface{}{}
	for k, v := range ids {
		msg, funcs := v.marshalProtobuf(ctx)
		pbAny, err := ptypes.MarshalAny(msg)
		if err != nil {
			continue
		}

		pbRecognizer := &pbtouch.Recognizer{
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
	next.NativeValues["gomatcha.io/matcha/touch"] = pbRecognizers

	if next.NativeFuncs == nil {
		next.NativeFuncs = map[string]interface{}{}
	}
	for k, v := range allFuncs {
		next.NativeFuncs[k] = v
	}
}

type Recognizer interface {
	marshalProtobuf(ctx *view.Context) (proto.Message, map[string]interface{})
	equal(Recognizer) bool
}

type TapEvent struct {
	Timestamp time.Time
	Position  layout.Point
}

func (e *TapEvent) unmarshalProtobuf(ev *pbtouch.TapEvent) error {
	t, _ := ptypes.Timestamp(ev.Timestamp)
	e.Timestamp = t
	e.Position.UnmarshalProtobuf(ev.Position)
	return nil
}

type TapRecognizer struct {
	Count   int
	OnTouch func(*TapEvent)
}

func (r *TapRecognizer) equal(a Recognizer) bool {
	b, ok := a.(*TapRecognizer)
	if !ok {
		return false
	}
	return r.Count == b.Count
}

func (r *TapRecognizer) marshalProtobuf(ctx *view.Context) (proto.Message, map[string]interface{}) {
	funcId := newFuncId()
	f := func(data []byte) {
		pbevent := &pbtouch.TapEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		event := &TapEvent{}
		if err := event.unmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}

		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &pbtouch.TapRecognizer{
			Count:          int64(r.Count),
			RecognizedFunc: funcId,
		}, map[string]interface{}{
			strconv.Itoa(int(funcId)): f,
		}
}

type EventKind int

const (
	EventKindPossible EventKind = iota
	EventKindChanged
	EventKindFailed
	EventKindRecognized
)

type PressEvent struct {
	Kind      EventKind
	Timestamp time.Time
	Position  layout.Point
	Duration  time.Duration
}

func (e *PressEvent) unmarshalProtobuf(ev *pbtouch.PressEvent) error {
	d, err := ptypes.Duration(ev.Duration)
	if err != nil {
		return err
	}
	t, err := ptypes.Timestamp(ev.Timestamp)
	if err != nil {
		return err
	}
	e.Kind = EventKind(ev.Kind)
	e.Timestamp = t
	e.Position.UnmarshalProtobuf(ev.Position)
	e.Duration = d
	return nil
}

type PressRecognizer struct {
	MinDuration time.Duration
	OnTouch     func(e *PressEvent)
}

func (r *PressRecognizer) equal(a Recognizer) bool {
	b, ok := a.(*PressRecognizer)
	if !ok {
		return false
	}
	return r.MinDuration == b.MinDuration
}

func (r *PressRecognizer) marshalProtobuf(ctx *view.Context) (proto.Message, map[string]interface{}) {
	funcId := newFuncId()
	f := func(data []byte) {
		event := &PressEvent{}
		pbevent := &pbtouch.PressEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		if err := event.unmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}
		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &pbtouch.PressRecognizer{
			MinDuration: ptypes.DurationProto(r.MinDuration),
			FuncId:      funcId,
		}, map[string]interface{}{
			strconv.Itoa(int(funcId)): f,
		}
}

type ButtonEvent struct {
	Timestamp time.Time
	Inside    bool
	Kind      EventKind
}

func (e *ButtonEvent) unmarshalProtobuf(ev *pbtouch.ButtonEvent) error {
	t, err := ptypes.Timestamp(ev.Timestamp)
	if err != nil {
		return err
	}
	e.Timestamp = t
	e.Inside = ev.Inside
	e.Kind = EventKind(ev.Kind)
	return nil
}

type ButtonRecognizer struct {
	OnTouch       func(e *ButtonEvent)
	IgnoresScroll bool
}

func (r *ButtonRecognizer) equal(a Recognizer) bool {
	_, ok := a.(*ButtonRecognizer)
	if !ok {
		return false
	}
	return true
}

func (r *ButtonRecognizer) marshalProtobuf(ctx *view.Context) (proto.Message, map[string]interface{}) {
	funcId := newFuncId()
	f := func(data []byte) {
		event := &ButtonEvent{}
		pbevent := &pbtouch.ButtonEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		if err := event.unmarshalProtobuf(pbevent); err != nil {
			fmt.Println("error", err)
			return
		}

		if r.OnTouch != nil {
			r.OnTouch(event)
		}
	}

	return &pbtouch.ButtonRecognizer{
			OnEvent:       funcId,
			IgnoresScroll: r.IgnoresScroll,
		}, map[string]interface{}{
			strconv.Itoa(int(funcId)): f,
		}
}
