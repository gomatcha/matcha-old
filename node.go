package mochi

import (
	"mochi/bridge"
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
	root    *BuildContext
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
	e.root.Update(e.keyPath, key)
}

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
	Children      map[interface{}]*RenderNode
	ChildrenSlice []*RenderNode
	Layouter      Layouter
	Painter       Painter
	Bridge        Bridge

	LayoutGuide  Guide
	PaintOptions PaintOptions

	BuildId int
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
	keyPath       []interface{}
	view          View
	node          *Node
	children      map[interface{}]*BuildContext
	childrenSlice []*BuildContext
	root          *BuildContext
	needsUpdate   bool
	buildId       int
}

func NewBuildContext(f func(Config) View) *BuildContext {
	ctx := &BuildContext{}
	e := &Embed{
		mu:      &sync.Mutex{},
		keyPath: nil,
		root:    ctx,
	}
	cfg := Config{Embed: e}
	ctx.view = f(cfg)
	ctx.root = ctx
	ctx.needsUpdate = true
	return ctx
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
			root:    ctx.root,
		},
	}
}

func (ctx *BuildContext) RenderNode() *RenderNode {
	n := &RenderNode{
		Layouter:      ctx.node.Layouter,
		Painter:       ctx.node.Painter,
		Bridge:        ctx.node.Bridge,
		Children:      map[interface{}]*RenderNode{},
		ChildrenSlice: []*RenderNode{},
		BuildId:       ctx.buildId,
	}
	for k, v := range ctx.children {
		rn := v.RenderNode()
		n.Children[k] = rn
		n.ChildrenSlice = append(n.ChildrenSlice, rn)
	}
	return n
}

func (ctx *BuildContext) Build() {
	if ctx.needsUpdate {
		ctx.needsUpdate = false
		ctx.buildId += 1

		// Generate the new node.
		node := ctx.view.Build(ctx)

		// Build new children from the node.
		prevChildren := ctx.children
		children := map[interface{}]*BuildContext{}
		childrenSlice := []*BuildContext{}
		for k, v := range node.Children {
			chlCtx := &BuildContext{}
			chlCtx.keyPath = append([]interface{}(nil), ctx.keyPath)
			chlCtx.keyPath = append(chlCtx.keyPath, k)
			chlCtx.view = v
			chlCtx.root = ctx.root
			chlCtx.needsUpdate = true
			children[k] = chlCtx
			childrenSlice = append(childrenSlice, chlCtx)
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
			chlCtx.node = prevChlCtx.node
			chlCtx.children = prevChlCtx.children
		}

		ctx.children = children
		ctx.childrenSlice = childrenSlice
		ctx.node = node
	}

	// Recursively update children.
	for _, i := range ctx.children {
		i.Build()
	}
}

func (ctx *BuildContext) Update(keyPath []interface{}, key interface{}) {
	ctx.needsUpdate = true
	bridge.Root().Call("rerender")
}
