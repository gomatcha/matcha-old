package mochi

import (
	"sync"
)

type Bridge struct {
	Name  string
	State interface{}
}

type Node struct {
	Children map[interface{}]View
	Layouter Layouter
	Painter  Painter
	Bridge   Bridge

	// Context map[string] interface{}
	// Handlers map[interface{}]Handler
	// Accessibility
	// Gesture Recognizers
	// OnAboutToScrollIntoView??
	// LayoutData?
}

func (n *Node) Set(k interface{}, v View) {
	if n.Children == nil {
		n.Children = map[interface{}]View{}
	}
	n.Children[k] = v
}

type RenderNode struct {
	Children map[interface{}]*RenderNode
	Layouter Layouter
	Painter  Painter
	Bridge   Bridge

	LayoutGuide  Guide
	PaintOptions PaintOptions
}

func (n *RenderNode) LayoutRoot(minSize Point, maxSize Point) {
	g := n.Layout(minSize, maxSize)
	g.Frame = g.Frame.Add(Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	n.LayoutGuide = g
}

func (n *RenderNode) Layout(minSize Point, maxSize Point) Guide {
	// Create the LayoutContext
	ctx := &LayoutContext{
		MinSize:   minSize,
		MaxSize:   maxSize,
		ChildKeys: []interface{}{},
		node:      n,
	}
	for i := range n.Children {
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
		n.Children[k].LayoutGuide = v
	}
	return g
}

func (n *RenderNode) Paint() {
	if p := n.Painter; p != nil {
		n.PaintOptions = p.PaintOptions()
	}
	for _, v := range n.Children {
		v.Paint()
	}
}

type Config struct {
	Prev  View
	Embed *Embed
}

type BuildContext struct {
	keyPath  []interface{}
	view     View
	node     *Node
	markerId int
	children map[interface{}]*BuildContext
}

func NewBuildContext(v View) *BuildContext {
	return &BuildContext{
		view: v,
	}
}

func (ctx *BuildContext) Get(k interface{}) Config {
	var prev View
	if chl := ctx.children[k]; chl != nil {
		prev = chl.view
	}
	return Config{
		Prev: prev,
		Embed: &Embed{
			mu:      &sync.Mutex{},
			keyPath: append([]interface{}(nil), k),
		},
	}
}

func (ctx *BuildContext) RenderNode() *RenderNode {
	n := &RenderNode{}
	n.Layouter = ctx.node.Layouter
	n.Painter = ctx.node.Painter
	n.Bridge = ctx.node.Bridge
	n.Children = map[interface{}]*RenderNode{}

	for k, v := range ctx.children {
		n.Children[k] = v.RenderNode()
	}
	return n
}

func (ctx *BuildContext) Build() {
	// Generate the new node.
	node := ctx.view.Build(ctx)

	// Build new children from the node.
	prevChildren := ctx.children
	children := map[interface{}]*BuildContext{}
	for k, v := range node.Children {
		chlCtx := &BuildContext{}
		chlCtx.keyPath = append([]interface{}(nil), ctx.keyPath)
		chlCtx.keyPath = append(chlCtx.keyPath, k)
		chlCtx.view = v
		children[k] = chlCtx
	}

	// Diff the children.
	addedKeys := []interface{}{}
	removedKeys := []interface{}{}
	updatedKeys := []interface{}{}
	unupdatedKeys := []interface{}{}
	for k, prevChlCtx := range prevChildren {
		chlCtx, ok := children[k]
		if !ok {
			removedKeys = append(removedKeys, k)
		} else if prevChlCtx.view == chlCtx.view {
			unupdatedKeys = append(unupdatedKeys, k)
		} else {
			updatedKeys = append(updatedKeys, k)
		}
	}
	for k := range children {
		_, ok := prevChildren[k]
		if !ok {
			addedKeys = append(addedKeys, k)
		}
	}

	// Pass properties from the old child to the new if it was unupdated.
	for _, k := range unupdatedKeys {
		prevChlCtx := prevChildren[k]
		chlCtx := children[k]
		chlCtx.markerId = prevChlCtx.markerId
		chlCtx.node = prevChlCtx.node
		chlCtx.children = prevChlCtx.children
	}

	// Recursively update children.
	ctx.children = children
	ctx.node = node
	for _, i := range children {
		i.Build()
	}
}
