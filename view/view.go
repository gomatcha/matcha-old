package view

import (
	"sync"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/paint"
)

type Screen interface {
	sync.Locker
	View(*Context) View
}

// ScreenFunc is an adapter to allow the use of ordinary functions as a Screen.
type ScreenFunc func(*Context) View

// View calls f(ctx, key).
func (f ScreenFunc) View(ctx *Context) View {
	return f(ctx)
}

// Lock is a no-op.
func (f ScreenFunc) Lock() {
	// no-op
}

// Unlock is a no-op.
func (f ScreenFunc) Unlock() {
	// no-op
}

type View interface {
	Build(*Context) *Model
	Lifecycle(from, to Stage)
	Id() matcha.Id
	comm.Notifier
}

// Embed is a convenience struct that provides a default implementation of View. It also wraps a comm.GroupNotifier.
type Embed struct {
	mu       sync.Mutex
	id       matcha.Id
	notifier comm.GroupNotifier
}

// NewEmbed creates a new Embed with the given Id.
func NewEmbed(id matcha.Id) *Embed {
	return &Embed{id: id}
}

// Build is an empty implementation of View's Build method.
func (e *Embed) Build(ctx *Context) *Model {
	return &Model{}
}

// Id returns the id passed into NewEmbed
func (e *Embed) Id() matcha.Id {
	return e.id
}

// Lifecycle is an empty implementation of View's Lifecycle method.
func (e *Embed) Lifecycle(from, to Stage) {
	// no-op
}

// Notify calls Notify(id) on the underlying comm.GroupNotifier.
func (e *Embed) Notify(f func()) comm.Id {
	return e.notifier.Notify(f)
}

// Unnotify calls Unnotify(id) on the underlying comm.GroupNotifier.
func (e *Embed) Unnotify(id comm.Id) {
	e.notifier.Unnotify(id)
}

// Subscribe calls Subscribe(n) on the underlying comm.GroupNotifier.
func (e *Embed) Subscribe(n comm.Notifier) {
	e.notifier.Subscribe(n)
}

// Unsubscribe calls Unsubscribe(n) on the underlying comm.GroupNotifier.
func (e *Embed) Unsubscribe(n comm.Notifier) {
	e.notifier.Unsubscribe(n)
}

// Update calls Signal() on the underlying comm.GroupNotifier.
func (e *Embed) Signal() {
	e.notifier.Signal()
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

// EntersStage returns true if from<s and to>=s.
func EntersStage(from, to, s Stage) bool {
	return from < s && to >= s
}

// ExitsStage returns true if from>=s and to<s.
func ExitsStage(from, to, s Stage) bool {
	return from >= s && to < s
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
	NativeFuncs     map[string]interface{}
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
