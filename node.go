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
	e.root.updateFlags[e.id] = buildFlag
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
	Id           Id
	Children     map[Id]*RenderNode
	Bridge       Bridge
	LayoutGuide  *Guide
	PaintOptions PaintOptions
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

		var flag updateFlag
		for _, v := range vc.root.updateFlags {
			flag |= v
		}

		if flag.needsBuild() {
			vc.root.build()
		}
		if flag.needsLayout() {
			vc.root.layout(Pt(0, 0), vc.size)
		}
		if flag.needsPaint() {
			vc.root.paint()
		}
		vc.root.updateFlags = map[Id]updateFlag{}

		rn := vc.root.node.renderNode()
		mochibridge.Root().Call("updateId:withRenderNode:", mochibridge.Int64(int64(id)), mochibridge.Interface(rn))
	})
	return vc
}

func (vc *ViewController) Render() *RenderNode {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.root.build()
	vc.root.layout(Pt(0, 0), vc.size)
	vc.root.paint()
	return vc.root.node.renderNode()
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

type updateFlag int

const (
	buildFlag updateFlag = 1 << iota
	layoutFlag
	paintFlag
)

func (f updateFlag) needsBuild() bool {
	return f&buildFlag != 0
}

func (f updateFlag) needsLayout() bool {
	return f&buildFlag != 0 || f&layoutFlag != 0
}

func (f updateFlag) needsPaint() bool {
	return f&buildFlag != 0 || f&layoutFlag != 0 || f&paintFlag != 0
}

type root struct {
	node        *node
	ids         map[viewCacheKey]Id
	prevIds     map[viewCacheKey]Id
	nodes       map[Id]*node
	prevNodes   map[Id]*node
	maxId       Id
	updateFlags map[Id]updateFlag
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
	root.updateFlags = map[Id]updateFlag{id: buildFlag}
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
}

func (root *root) paint() {
	root.node.paint()
}

func (root *root) layout(minSize Point, maxSize Point) {
	g := root.node.layout(minSize, maxSize)
	g.Frame = g.Frame.Add(Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	root.node.layoutGuide = &g
}

func (root *root) newId() Id {
	root.maxId += 1
	return root.maxId
}

type node struct {
	id   Id
	root *root
	view View

	viewModel    *ViewModel
	children     map[Id]*node
	layoutGuide  *Guide
	paintOptions PaintOptions
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
		Id:           n.id,
		Children:     map[Id]*RenderNode{},
		Bridge:       n.viewModel.Bridge,
		LayoutGuide:  n.layoutGuide,
		PaintOptions: n.paintOptions,
	}
	for k, v := range n.children {
		rn.Children[k] = v.renderNode()
	}
	return rn
}

func (n *node) build() {
	if n.root.updateFlags[n.id].needsBuild() {
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

		// Mark all children as needing rebuild since we rebuilt
		for k := range children {
			n.root.updateFlags[k] |= buildFlag
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

func (n *node) paint() {
	if n.root.updateFlags[n.id].needsPaint() {
		if p := n.viewModel.Painter; p != nil {
			n.paintOptions = p.PaintOptions()
		} else {
			n.paintOptions = PaintOptions{}
		}
	}

	// Recursively update children
	for _, v := range n.children {
		v.paint()
	}
}

func (n *node) layout(minSize Point, maxSize Point) Guide {
	// Create the LayoutContext
	ctx := &LayoutContext{
		MinSize:  minSize,
		MaxSize:  maxSize,
		ChildIds: []Id{},
		node:     n,
	}
	for i := range n.children {
		ctx.ChildIds = append(ctx.ChildIds, i)
	}

	// Perform layout
	layouter := n.viewModel.Layouter
	if layouter == nil {
		layouter = &FullLayout{}
	}
	g, gs := layouter.Layout(ctx)
	g = g.fit(ctx)

	// Assign guides to children
	for k, v := range gs {
		guide := v
		n.children[k].layoutGuide = &guide
	}
	return g
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
