package mochi

import (
	"sync"
)

type View interface {
	Build(*BuildContext) *Node
	// Lifecyle(*Stage)
	Key() interface{}
	Lock()
	Unlock()
}

type Embed struct {
	mu      *sync.Mutex
	keyPath []interface{}
}

func (e *Embed) Build(ctx *BuildContext) *Node {
	return &Node{}
}

func (e *Embed) Key() interface{} {
	return e.keyPath[len(e.keyPath)-1]
}

func (e *Embed) Lock() {
	e.mu.Lock()
}

func (e *Embed) Unlock() {
	e.mu.Unlock()
}

func (e *Embed) Update(key interface{}) {
	// e.marks.Update(key)
}

type BasicView struct {
	*Embed
	PaintOptions PaintOptions
}

func NewBasicView(c Config) *BasicView {
	v, ok := c.Prev.(*BasicView)
	if !ok {
		v = &BasicView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *BasicView) Build(ctx *BuildContext) *Node {
	n := &Node{}
	n.Painter = v.PaintOptions
	return n
}
