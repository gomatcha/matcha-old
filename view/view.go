package view

import (
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
)

type Screen interface {
	sync.Locker
	NewView(*Context, interface{}) View
}

type ScreenFunc func(*Context, interface{}) View

func (f ScreenFunc) NewView(ctx *Context, key interface{}) View {
	return f(ctx, key)
}

func (f ScreenFunc) Lock() {
	// no-op
}

func (f ScreenFunc) Unlock() {
	// no-op
}

type View interface {
	Build(*Context) *Model
	Lifecycle(from, to Stage)
	Id() mochi.Id
	mochi.Notifier
}

type Embed struct {
	mu            sync.Mutex
	id            mochi.Id
	batchNotifier mochi.BatchNotifier
}

func NewEmbed(id mochi.Id) *Embed {
	return &Embed{id: id}
}

func (e *Embed) Build(ctx *Context) *Model {
	return &Model{}
}

func (e *Embed) Id() mochi.Id {
	return e.id
}

func (e *Embed) Lifecycle(from, to Stage) {
	// no-op
}

func (e *Embed) Notify() chan struct{} {
	return e.batchNotifier.Notify()
}

func (e *Embed) Unnotify(c chan struct{}) {
	e.batchNotifier.Unnotify(c)
}

func (e *Embed) Subscribe(n mochi.Notifier) {
	e.batchNotifier.Subscribe(n)
}

func (e *Embed) Unsubscribe(n mochi.Notifier) {
	e.batchNotifier.Unsubscribe(n)
}

func (e *Embed) Update() {
	e.batchNotifier.Signal()
}

type Stage int

const (
	StageDead Stage = iota
	StageMounted
	StagePrepreload
	StagePreload
	StageVisible
)

func EntersStage(from, to, s Stage) bool {
	return from < s && to >= s
}

func ExitsStage(from, to, s Stage) bool {
	return from >= s && to < s
}

type Model struct {
	Children []View
	Layouter layout.Layouter
	Painter  paint.Painter
	Values   map[interface{}]interface{}

	NativeViewName  string
	NativeViewState proto.Message
	NativeValues    map[string]proto.Message
	NativeFuncs     map[int64]interface{}
}
