package mochi

import (
	"fmt"
	"github.com/overcyn/mochi/internal"
	"mochi/bridge"
	"strings"
	"sync"
	"time"
)

type Id int64

type View interface {
	Build(*BuildContext) *Node
	// Lifecyle(*Stage)
	Id() Id
	Lock()
	Unlock()
}

type Embed struct {
	mu   *sync.Mutex
	id   Id
	root *BuildContextRoot
}

func (e *Embed) Build(ctx *BuildContext) *Node {
	return &Node{}
}

func (e *Embed) Id() Id {
	return e.id
}

func (e *Embed) Lock() {
	e.mu.Lock()
}

func (e *Embed) Unlock() {
	e.mu.Unlock()
}

func (e *Embed) Update(key interface{}) {
	e.root.Update(key)
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
	Children map[interface{}]*RenderNode
	Layouter Layouter
	Painter  Painter
	Bridge   Bridge

	LayoutGuide  *Guide
	PaintOptions PaintOptions
}

func (n *RenderNode) LayoutRoot(minSize Point, maxSize Point) {
	g := n.Layout(minSize, maxSize)
	g.Frame = g.Frame.Add(Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	n.LayoutGuide = &g
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
		guide := v
		n.Children[k].LayoutGuide = &guide
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

func (n *RenderNode) DebugString() string {
	all := []string{}
	for _, i := range n.Children {
		lines := strings.Split(i.DebugString(), "\n")
		for idx, line := range lines {
			lines[idx] = "|	" + line
		}
		all = append(all, lines...)
	}

	str := fmt.Sprintf("{%p Bridge:%v LayoutGuide:%v PaintOptions:%v}", n, n.Bridge, n.LayoutGuide, n.PaintOptions)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}

func (n *RenderNode) Copy() *RenderNode {
	children := map[interface{}]*RenderNode{}
	for k, v := range n.Children {
		children[k] = v
	}
	copy := &RenderNode{
		Children:     children,
		Layouter:     n.Layouter,
		Painter:      n.Painter,
		Bridge:       n.Bridge,
		LayoutGuide:  n.LayoutGuide,
		PaintOptions: n.PaintOptions,
	}
	return copy
}

type ViewController struct {
	id           int
	mu           *sync.Mutex
	buildContext *BuildContext
	renderNode   *RenderNode
	stopped      bool
	size         Point
	ticker       *internal.Ticker
}

func NewViewController(f func(Config) View, id int) *ViewController {
	vc := &ViewController{
		mu:           &sync.Mutex{},
		buildContext: NewBuildContext(f),
		ticker:       internal.NewTicker(time.Hour * 99999),
		id:           id,
	}

	// start run loop
	vc.ticker.NotifyFunc(func() {
		vc.mu.Lock()
		defer vc.mu.Unlock()

		if vc.stopped {
			return
		}
		rn := vc.renderNode
		rn.LayoutRoot(Pt(0, 0), vc.size)
		rn.Paint()

		bridge.Root().Call("updateId:withRenderNode:", bridge.Int64(int64(id)), bridge.Interface(rn))
	})

	return vc
}

func (vc *ViewController) Render() *RenderNode {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.buildContext.Build()
	vc.renderNode = vc.buildContext.RenderNode()

	rn := vc.renderNode.Copy()
	rn.LayoutRoot(Pt(0, 0), vc.size)
	rn.Paint()
	return rn
}

func (vc *ViewController) SetSize(p Point) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.size = p
}

func (vc *ViewController) Stop() {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.stopped = true
}

type Config struct {
	Prev  View
	Embed *Embed
}

type BuildContextRoot struct {
	ctx       *BuildContext
	viewCache map[Id]View
	maxId     Id
}

func (root *BuildContextRoot) Update(key interface{}) {
	root.ctx.needsUpdate = true
	bridge.Root().Call("rerender")
}

func (root *BuildContextRoot) NewId() Id {
	root.maxId += 1
	return root.maxId
}

type BuildContext struct {
	view        View
	node        *Node
	children    map[interface{}]*BuildContext
	root        *BuildContextRoot
	needsUpdate bool
}

func NewBuildContext(f func(Config) View) *BuildContext {
	ctx := &BuildContext{}
	e := &Embed{
		mu: &sync.Mutex{},
	}
	cfg := Config{Embed: e}
	ctx.view = f(cfg)
	ctx.root = &BuildContextRoot{ctx: ctx}
	ctx.needsUpdate = true
	e.root = ctx.root
	return ctx
}

func (ctx *BuildContext) Get(k interface{}) Config {
	var prev View
	if chl := ctx.children[k]; chl != nil {
		prev = chl.view
	} else {
		fmt.Println("No Prev", ctx.view, k)
	}

	return Config{
		Prev: prev,
		Embed: &Embed{
			mu:   &sync.Mutex{},
			root: ctx.root,
			id:   ctx.root.NewId(),
		},
	}
}

func (ctx *BuildContext) RenderNode() *RenderNode {
	n := &RenderNode{
		Layouter: ctx.node.Layouter,
		Painter:  ctx.node.Painter,
		Bridge:   ctx.node.Bridge,
		Children: map[interface{}]*RenderNode{},
	}
	for k, v := range ctx.children {
		rn := v.RenderNode()
		n.Children[k] = rn
	}
	return n
}

func (ctx *BuildContext) Build() {
	if ctx.needsUpdate {
		ctx.needsUpdate = false

		// Generate the new node.
		node := ctx.view.Build(ctx)

		// Build new children from the node.
		prevChildren := ctx.children
		children := map[interface{}]*BuildContext{}
		for k, v := range node.Children {
			chlCtx := &BuildContext{}
			chlCtx.view = v
			chlCtx.root = ctx.root
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
			chlCtx.node = prevChlCtx.node
			chlCtx.children = prevChlCtx.children
		}

		// Mark all children as needing update since we updated
		for _, i := range children {
			i.needsUpdate = true
		}

		ctx.children = children
		ctx.node = node
	}

	// Recursively update children.
	for _, i := range ctx.children {
		i.Build()
	}
}

func (ctx *BuildContext) DebugString() string {
	all := []string{}
	for _, i := range ctx.children {
		lines := strings.Split(i.DebugString(), "\n")
		for idx, line := range lines {
			lines[idx] = "|	" + line
		}
		all = append(all, lines...)
	}

	str := fmt.Sprintf("{%p view:%v node:%p}", ctx, ctx.view, ctx.node)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}
