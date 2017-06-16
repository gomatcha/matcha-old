package view

import (
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
)

type Screen interface {
	sync.Locker
	View(*Context, interface{}) View
}

type ScreenFunc func(*Context, interface{}) View

func (f ScreenFunc) View(ctx *Context, key interface{}) View {
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
	comm.Notifier
}

// Embed is a convenience struct that provides a default implementation of View. It also wraps a comm.BatchNotifier.
type Embed struct {
	mu            sync.Mutex
	id            mochi.Id
	batchNotifier comm.BatchNotifier
}

// NewEmbed creates a new Embed with the given Id.
func NewEmbed(id mochi.Id) *Embed {
	return &Embed{id: id}
}

// Build is an empty implementation of View's Build method.
func (e *Embed) Build(ctx *Context) *Model {
	return &Model{}
}

// Id returns the id passed into NewEmbed
func (e *Embed) Id() mochi.Id {
	return e.id
}

// Lifecycle is an empty implementation of View's Lifecycle method.
func (e *Embed) Lifecycle(from, to Stage) {
	// no-op
}

// Notify calls Notify(id) on the underlying comm.BatchNotifier.
func (e *Embed) Notify(f func()) comm.Id {
	return e.batchNotifier.Notify(f)
}

// Unnotify calls Unnotify(id) on the underlying comm.BatchNotifier.
func (e *Embed) Unnotify(id comm.Id) {
	e.batchNotifier.Unnotify(id)
}

// Subscribe calls Subscribe(n) on the underlying comm.BatchNotifier.
func (e *Embed) Subscribe(n comm.Notifier) {
	e.batchNotifier.Subscribe(n)
}

// Unsubscribe calls Unsubscribe(n) on the underlying comm.BatchNotifier.
func (e *Embed) Unsubscribe(n comm.Notifier) {
	e.batchNotifier.Unsubscribe(n)
}

// Update calls Update() on the underlying comm.BatchNotifier.
func (e *Embed) Update() {
	e.batchNotifier.Update()
}

type Stage int

const (
	// StageDead marks views that are not attached to the view hierarchy.
	StageDead Stage = iota
	// StageMounted marks views that are in the view hierarchy but not visible.
	StageMounted
	// StageVisible marks views that are in the view hierarchy and visible.
	StageVisible
)

// EntersStage returns true if start<s and end>=s.
func EntersStage(start, end, s Stage) bool {
	return start < s && end >= s
}

// ExitsStage returns true if start>=s and end<s.
func ExitsStage(start, end, s Stage) bool {
	return start >= s && end < s
}

// Model describes the view and its children.
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

// WithPainter wraps the view v, and replaces its Model.Painter with p.
func WithPainter(v View, p paint.Painter) View {
	return &painterView{View: v, painter: p}
}

type painterView struct {
	View
	painter paint.Painter
}

func (v *painterView) Build(ctx *Context) *Model {
	m := v.View.Build(ctx)
	m.Painter = v.painter
	return m
}

// WithValues wraps the view v, and adds the given values to its Model.Values.
func WithValues(v View, vals map[interface{}]interface{}) View {
	return &valuesView{View: v, values: vals}
}

type valuesView struct {
	View
	values map[interface{}]interface{}
}

func (v *valuesView) Build(ctx *Context) *Model {
	m := v.View.Build(ctx)
	if m.Values == nil {
		m.Values = map[interface{}]interface{}{}
	}
	for k, val := range v.values {
		m.Values[k] = val
	}
	return m
}
