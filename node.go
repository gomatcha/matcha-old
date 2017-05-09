package mochi

import (
	"fmt"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochibridge"
	"strings"
	"sync"
	"time"
)

type Id int64

type View interface {
	Build(*BuildContext) *ViewModel
	// Lifecyle(*Stage)
	Id() Id
	Lock()
	Unlock()
}

type Embed struct {
	mu   *sync.Mutex
	id   Id
	root *root
}

func (e *Embed) Build(ctx *BuildContext) *ViewModel {
	return &ViewModel{}
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
	e.root.updateIds[id] = struct{}{}
}

type Bridge struct {
	Name  string
	State interface{}
}

type ViewModel struct {
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

func (n *ViewModel) Add(v View) {
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
	id         int
	mu         *sync.Mutex
	root       *root
	renderNode *RenderNode
	stopped    bool
	size       Point
	ticker     *internal.Ticker
}

func NewViewController(f func(Config) View, id int) *ViewController {
	vc := &ViewController{
		mu:     &sync.Mutex{},
		root:   newRoot(f),
		ticker: internal.NewTicker(time.Hour * 99999),
		id:     id,
	}

	// start run loop
	vc.ticker.NotifyFunc(func() {
		vc.mu.Lock()
		defer vc.mu.Unlock()

		if vc.stopped {
			return
		}

		if len(vc.root.updateIds) > 0 {
			vc.root.build()
			vc.renderNode = vc.root.node.renderNode()
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

	vc.root.build()
	vc.renderNode = vc.root.node.renderNode()

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

type viewCacheKey struct {
	id  Id
	key interface{}
}

type BuildContext struct {
	node *node
}

func (ctx *BuildContext) Get(key interface{}) Config {
	return ctx.node.get(key)
}

type root struct {
	node      *node
	ids       map[viewCacheKey]Id
	prevIds   map[viewCacheKey]Id
	nodes     map[Id]*node
	prevNodes map[Id]*node
	maxId     Id
	updateIds map[Id]struct{}
}

func newRoot(f func(Config) View) *root {
	root := &root{}

	id := root.newId()
	cfg := Config{Embed: &Embed{
		mu:   &sync.Mutex{},
		root: root,
		id:   id,
	}}
	node := &node{
		id:   id,
		view: f(cfg),
		root: root,
	}
	root.node = node
	root.updateIds = map[Id]struct{}{id: struct{}{}}
	return root
}

func (root *root) build() {
	root.prevIds = root.ids
	root.ids = map[viewCacheKey]Id{}
	root.prevNodes = root.nodes
	root.nodes = map[Id]*node{}
	root.node.build()

	keys := map[Id]viewCacheKey{}
	for k, v := range root.ids {
		keys[v] = k
	}
	for k, v := range root.prevIds {
		keys[v] = k
	}

	ids := map[viewCacheKey]Id{}
	for k := range root.nodes {
		ids[keys[k]] = k
	}
	root.ids = ids
	root.updateIds = map[Id]struct{}{}
}

func (root *root) newId() Id {
	root.maxId += 1
	return root.maxId
}

type node struct {
	id          Id
	view        View
	viewModel   *ViewModel
	root        *root
	children    map[Id]*node
	needsUpdate bool
}

func (n *node) get(key interface{}) Config {
	cacheKey := viewCacheKey{key: key, id: n.id}
	id := n.root.newId()

	prevId := n.root.prevIds[cacheKey]
	prevCtx := n.root.prevNodes[prevId]
	var prevView View
	if prevCtx != nil {
		prevView = prevCtx.view
	}

	n.root.ids[cacheKey] = id
	return Config{
		Prev: prevView,
		Embed: &Embed{
			mu:   &sync.Mutex{},
			root: n.root,
			id:   id,
		},
	}
}

func (n *node) renderNode() *RenderNode {
	rn := &RenderNode{
		Id:       n.id,
		Layouter: n.viewModel.Layouter,
		Painter:  n.viewModel.Painter,
		Bridge:   n.viewModel.Bridge,
		Children: map[Id]*RenderNode{},
	}
	for k, v := range n.children {
		rn.Children[k] = v.renderNode()
	}
	return rn
}

func (n *node) build() {
	_, needsUpdate := n.root.updateIds[n.id]
	if needsUpdate {
		// Generate the new viewModel.
		viewModel := n.view.Build(&BuildContext{node: n})

		// Diff the old children (n.children) with new children (viewModel.Children).
		addedKeys := []Id{}
		removedKeys := []Id{}
		unupdatedKeys := []Id{}
		for id := range n.children {
			if _, ok := viewModel.Children[id]; !ok {
				removedKeys = append(removedKeys, id)
			} else {
				unupdatedKeys = append(unupdatedKeys, id)
			}
		}
		for id := range viewModel.Children {
			if _, ok := n.children[id]; !ok {
				addedKeys = append(addedKeys, id)
			}
		}

		children := map[Id]*node{}
		// Add build contexts for new children
		for _, id := range addedKeys {
			var view View
			for _, i := range viewModel.Children {
				if i.Id() == id {
					view = i
					break
				}
			}
			children[id] = &node{
				id:       id,
				view:     view,
				children: map[Id]*node{},
				root:     n.root,
			}
		}
		// Reuse old context for unupdated keys
		for _, id := range unupdatedKeys {
			children[id] = n.children[id]
		}

		// Mark all children as needing update since we updated
		for k := range children {
			n.root.updateIds[k] = struct{}{}
		}

		n.children = children
		n.viewModel = viewModel
	}

	// Recursively update children.
	for _, i := range n.children {
		i.build()

		// Also add to the root
		n.root.nodes[i.id] = i
	}
}

func (n *node) debugString() string {
	all := []string{}
	for _, i := range n.children {
		lines := strings.Split(i.debugString(), "\n")
		for idx, line := range lines {
			lines[idx] = "|	" + line
		}
		all = append(all, lines...)
	}

	str := fmt.Sprintf("{%p Id:%v View:%v Node:%p}", n, n.id, n.view, n.viewModel)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}
