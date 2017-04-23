package mochi

import (
	_ "fmt"
)

type Config struct {
	Prev   View
	Marker Marker
}

type View interface {
	Update(*ViewContext) *Node
	// Key() interface{}
	// Lifecycle(Stage)
}

type Stage int

type Marker interface {
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
	marker       Marker
	PaintOptions PaintOptions
}

func NewBasicView(c Config) *BasicView {
	v, ok := c.Prev.(*BasicView)
	if !ok {
		v = &BasicView{}
		v.marker = c.Marker
	}
	return v
}

func (v *BasicView) Update(ctx *ViewContext) *Node {
	n := &Node{}
	n.Painter = v.PaintOptions
	return n
}
