package mochi

import (
	"fmt"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochibridge"
	"runtime/debug"
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
	Children map[Id]View
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

func (n *Node) Add(v View) {
	if n.Children == nil {
		n.Children = map[Id]View{}
	}
	n.Children[v.Id()] = v
}

type RenderNode struct {
	Children map[Id]*RenderNode
	Layouter Layouter
	Painter  Painter
	Bridge   Bridge

	Id           Id
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
		MinSize:  minSize,
		MaxSize:  maxSize,
		ChildIds: []Id{},
		node:     n,
	}
	for i := range n.Children {
		ctx.ChildIds = append(ctx.ChildIds, i)
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

	str := fmt.Sprintf("{%p Id:%v Bridge:%v LayoutGuide:%v PaintOptions:%v}", n, n.Id, n.Bridge, n.LayoutGuide, n.PaintOptions)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}

func (n *RenderNode) Copy() *RenderNode {
	children := map[Id]*RenderNode{}
	for k, v := range n.Children {
		children[k] = v.Copy()
	}
	copy := &RenderNode{
		Id:           n.Id,
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

		mochibridge.Root().Call("updateId:withRenderNode:", mochibridge.Int64(int64(id)), mochibridge.Interface(rn))
	})
	return vc
}

func (vc *ViewController) Render() *RenderNode {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.buildContext.Build()
	fmt.Println("VC", vc.buildContext.DebugString())
	vc.renderNode = vc.buildContext.RenderNode()

	rn := vc.renderNode.Copy()
	rn.LayoutRoot(Pt(0, 0), vc.size)
	rn.Paint()
	fmt.Println("WTF", rn.DebugString())
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
	mochibridge.Root().Call("rerender")
	debug.PrintStack()
}

func (root *BuildContextRoot) NewId() Id {
	root.maxId += 1
	return root.maxId
}

type BuildContext struct {
	id          Id
	view        View
	node        *Node
	children    map[Id]*BuildContext
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
	ctx.id = ctx.root.NewId()
	ctx.needsUpdate = true
	e.root = ctx.root
	return ctx
}

func (ctx *BuildContext) Get(k interface{}) Config {
	var prev View = nil
	// if chl := ctx.children[k]; chl != nil {
	// 	prev = chl.view
	// } else {
	// 	fmt.Println("No Prev", ctx.view, k)
	// }

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
		Id:       ctx.id,
		Layouter: ctx.node.Layouter,
		Painter:  ctx.node.Painter,
		Bridge:   ctx.node.Bridge,
		Children: map[Id]*RenderNode{},
	}
	for k, v := range ctx.children {
		n.Children[k] = v.RenderNode()
	}
	return n
}

func (ctx *BuildContext) Build() {
	if ctx.needsUpdate {
		ctx.needsUpdate = false

		// Generate the new node.
		node := ctx.view.Build(ctx)

		// Diff the old children (ctx.children) with new children (node.Children).
		addedKeys := []Id{}
		removedKeys := []Id{}
		unupdatedKeys := []Id{}
		for id := range ctx.children {
			if _, ok := node.Children[id]; !ok {
				removedKeys = append(removedKeys, id)
			} else {
				unupdatedKeys = append(unupdatedKeys, id)
			}
		}
		for id := range node.Children {
			if _, ok := ctx.children[id]; !ok {
				addedKeys = append(addedKeys, id)
			}
		}

		children := map[Id]*BuildContext{}
		// Add build contexts for new children
		for _, id := range addedKeys {
			var view View
			for _, i := range node.Children {
				if i.Id() == id {
					view = i
					break
				}
			}
			children[id] = &BuildContext{
				id:       id,
				view:     view,
				children: map[Id]*BuildContext{},
				root:     ctx.root,
			}
		}
		// Reuse old context for unupdated keys
		for _, id := range unupdatedKeys {
			children[id] = ctx.children[id]
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

	str := fmt.Sprintf("{%p Id:%v View:%v Node:%p}", ctx, ctx.id, ctx.view, ctx.node)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}
