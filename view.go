package mochi

import (
	_ "fmt"
)

type View interface {
    Mount(m Marker)
	Update(n *Node) *Node
    Unmount()
}

type Marker interface {
    Update()
    UpdateChild(interface{})
    Run()
}

type marker struct {
}
func (m *marker)Update() {
}
func (m *marker)UpdateChild(interface{}) {
}
func (m *marker)Run() {
}

type BasicView struct {
    marker Marker
	PaintOptions PaintOptions
}

func NewBasicView(p interface{}) *BasicView {
	return &BasicView{}
}

func (v *BasicView) Mount(m Marker) {
    v.marker = m
}

func (v *BasicView) Update(p *Node) *Node {
	n := &Node{}
	n.PaintOptions = v.PaintOptions
	return n
}

func (v *BasicView) Unmount() {
    v.marker = nil
}
