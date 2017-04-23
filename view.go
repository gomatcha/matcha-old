package mochi

import (
	_ "fmt"
)

type Config struct {
	Prev    View
	Updater Updater
}

type View interface {
	Build(*PaintContext) *Node
	// Key() interface{}
	// Lifecycle(Stage)
}

type Stage int

type Updater interface {
	Update()
	// UpdateChild(interface{})
	// Run()
}

type marker struct {
	KeyPath []interface{}
}

func (m *marker) Update() {
}

type BasicView struct {
	marker       Updater
	PaintOptions PaintOptions
}

func NewBasicView(c Config) *BasicView {
	v, ok := c.Prev.(*BasicView)
	if !ok {
		v = &BasicView{}
		v.marker = c.Updater
	}
	return v
}

func (v *BasicView) Build(ctx *PaintContext) *Node {
	n := &Node{}
	n.Painter = v.PaintOptions
	return n
}
