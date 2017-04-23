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

func (n *RenderNode) layout(minSize Point, maxSize Point) Guide {
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

func (n *RenderNode) getPaintOptions() {
	if p := n.Painter; p != nil {
		n.PaintOptions = p.PaintOptions()
	}
	for _, v := range n.Children {
		v.getPaintOptions()
	}
}

type ViewContext struct {
	keyPath  []interface{}
	view     View
	node     *Node
	markerId int
	children map[interface{}]*ViewContext
}

func (ctx *ViewContext) Get(k interface{}) Config {
	marker := Marker{
		keyPath: ctx.keyPath,
		id:      ctx.markerId,
	}

	return Config{} // KD: TODO
	// if n == nil {
	// 	return nil
	// }
	// return n.Children[k]
}

func (ctx *ViewContext) RenderNode() *RenderNode {
	renderNode := &RenderNode{}
	renderNode.Layouter = ctx.node.Layouter
	renderNode.Painter = ctx.node.Painter
	renderNode.Bridge = ctx.node.Bridge
	renderNode.Children = map[interface{}]*RenderNode{}

	for k, v := range ctx.children {
		renderNode.Children[k] = v.RenderNode()
	}
	return renderNode
}

func (ctx *ViewContext) Update() {
	// Generate the new node.
	node := ctx.view.Update(ctx)

	// Build new children from the node.
	prevChildren := ctx.children
	children := map[interface{}]*ViewContext{}
	for k, v := range node.Children {
		chlCtx := &ViewContext{}
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
	for k := range unupdatedKeys {
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
		i.Update()
	}
}

func Display(v View) *RenderNode {
	ctx := &ViewContext{}
	ctx.view = v
	ctx.Update()
	renderNode := ctx.RenderNode()
	renderNode.layout(Pt(0, 0), Pt(1000, 1000))
	renderNode.getPaintOptions()
	return renderNode
}
