package mochi

import (
	_ "fmt"
	_ "image/color"
)

type Bridge struct {
	Name  string
	State interface{}
}

type Node struct {
	Children map[interface{}]View
	Layouter Layouter
	Painter  Painter
	Handlers map[interface{}]Handler

	Bridge Bridge

	// Context map[string] interface{}
	// Accessibility
	// Gesture Recognizers
	// OnAboutToScrollIntoView??
	// LayoutData?

	// These should be hidden, probs
	NodeChildren map[interface{}]*Node
	LayoutGuide  Guide
	PaintOptions PaintOptions
}

func NewNode() *Node {
	n := &Node{}
	n.Children = map[interface{}]View{}
	n.Handlers = map[interface{}]Handler{}
	n.NodeChildren = map[interface{}]*Node{}
	return n
}

func nodeFromView(view View, prev *Node) *Node {
	node := view.Update(prev)
	for k, v := range node.Children {
		var prevNode *Node
		if prev != nil {
			prevNode = prev.NodeChildren[k]
		}
		node.NodeChildren[k] = nodeFromView(v, prevNode)
	}
	return node
}

func (n *Node) Get(k interface{}) View {
	if n == nil {
		return nil
	}
	return n.Children[k]
}

func (n *Node) Set(k interface{}, v View) {
	n.Children[k] = v
}

func (n *Node) layout(minSize Point, maxSize Point) Guide {
	// Create the LayoutContext
	ctx := &LayoutContext{
		MinSize:   minSize,
		MaxSize:   maxSize,
		ChildKeys: []interface{}{},
		node:      n,
	}
	for i := range n.NodeChildren {
		ctx.ChildKeys = append(ctx.ChildKeys, i)
	}

	// Perform layout
	layouter := n.Layouter
	if layouter == nil {
		layouter = &FullLayout{}
	}
	g, gs := layouter.Layout(ctx)
	g = g.fit(ctx)

	// Assign guides to children
	for k, v := range gs {
		n.NodeChildren[k].LayoutGuide = v
	}
	return g
}

func (n *Node) getPaintOptions() {
	if p := n.Painter; p != nil {
		n.PaintOptions = p.PaintOptions()
	}
	for _, v := range n.NodeChildren {
		v.getPaintOptions()
	}
}

func Display(v View) *Node {
	node := nodeFromView(v, nil)
	node.layout(Pt(0, 0), Pt(1000, 1000))
	node.getPaintOptions()
	return node
}
