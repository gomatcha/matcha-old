package mochi

import (
	_ "fmt"
)

type View interface {
	Update(n *Node) *Node
	NeedsUpdate()
}

type BasicView struct {
	PaintOptions PaintOptions
}

func NewBasicView(p interface{}) *BasicView {
	return &BasicView{}
}

func (v *BasicView) Update(p *Node) *Node {
	n := &Node{}
	n.PaintOptions = v.PaintOptions
	return n
}

func (v *BasicView) NeedsUpdate() {
	// ?
}
