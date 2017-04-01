package mochi

import (
	_ "fmt"
)

type View interface {
	Update(n *Node) *Node
	NeedsUpdate()
}

type NestedView struct {
}

func (v *NestedView) Update(p *Node) *Node {
	l := NewAbsoluteLayout()
	n := NewNode()
	n.Layouter = l

	chl1id := "1"
	chl1 := NewBasicView(p.Get(chl1id))
	chl1.PaintOptions.BackgroundColor = redColor
	n.Set(chl1id, chl1)
	l.ChildGuides[chl1id] = Guide{Frame: Rt(0, 0, 100, 100)}

	chl2id := "2"
	chl2 := NewBasicView(p.Get(chl2id))
	chl2.PaintOptions.BackgroundColor = greenColor
	n.Set(chl2id, chl2)
	l.ChildGuides[chl2id] = Guide{Frame: Rt(0, 100, 100, 200)}

	chl3id := "3"
	chl3 := NewBasicView(p.Get(chl3id))
	chl3.PaintOptions.BackgroundColor = blueColor
	n.Set(chl3id, chl3)
	l.ChildGuides[chl3id] = Guide{Frame: Rt(0, 200, 100, 300)}

	n.PaintOptions.BackgroundColor = greenColor

	return n
}

func (n *NestedView) NeedsUpdate() {
	// ?
}

type BasicView struct {
	PaintOptions PaintOptions
}

func NewBasicView(p interface{}) *BasicView {
	return &BasicView{}
	// v, ok := p.(*BasicView)
	// if !ok {
	// 	v = &BasicView{}
	// }
	// return v
}

func (v *BasicView) Update(p *Node) *Node {
	n := &Node{}
	n.PaintOptions = v.PaintOptions
	return n
}

func (v *BasicView) NeedsUpdate() {
	// ?
}
