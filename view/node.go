package view

import (
	"fmt"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/layout/full"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochibridge"
	"strings"
	"sync"
	"time"
)

type View interface {
	Build(*Context) *Model
	// Lifecyle(*Stage)
	Id() mochi.Id
	Lock()
	Unlock()
}

type Embed struct {
	mu   *sync.Mutex
	id   mochi.Id
	root *root
}

func (e *Embed) Build(ctx *Context) *Model {
	return &Model{}
}

func (e *Embed) Id() mochi.Id {
	return e.id
}

func (e *Embed) Lock() {
	e.mu.Lock()
}

func (e *Embed) Unlock() {
	e.mu.Unlock()
}

func (e *Embed) Update(key interface{}) {
	e.root.addFlag(e.id, buildFlag)
}

type Bridge struct {
	Name  string
	State interface{}
}

type Model struct {
	Children map[mochi.Id]View
	Layouter layout.Layouter
	Painter  paint.Painter
	Bridge   Bridge

	// Context map[string] interface{}
	// Handlers map[interface{}]Handler
	// Accessibility
	// Gesture Recognizers
	// OnAboutToScrollIntoView??
	// LayoutData?
}

func (n *Model) Add(v View) {
	if n.Children == nil {
		n.Children = map[mochi.Id]View{}
	}
	n.Children[v.Id()] = v
}

type RenderNode struct {
	Id           mochi.Id
	BuildId      int64
	LayoutId     int64
	PaintId      int64
	Children     map[mochi.Id]*RenderNode
	Bridge       Bridge
	LayoutGuide  *layout.Guide
	PaintOptions paint.Style
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

	str := fmt.Sprintf("{%p Id:%v BuildId:%v LayoutId:%v PaintId:%v LayoutGuide:%v PaintOptions:%v}", n, n.Id, n.BuildId, n.LayoutId, n.PaintId, n.LayoutGuide, n.PaintOptions)
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
	size       layout.Point
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

		vc.root.update(vc.size)
		rn := vc.root.renderNode()
		mochibridge.Root().Call("updateId:withRenderNode:", mochibridge.Int64(int64(id)), mochibridge.Interface(rn))
		// fmt.Println(rn.DebugString())
	})
	return vc
}

func (vc *ViewController) Render() *RenderNode {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.root.update(vc.size)
	rn := vc.root.renderNode()
	return rn
}

func (vc *ViewController) SetSize(p layout.Point) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.size = p
}

type Config struct {
	Prev  View
	Embed *Embed
}

type viewCacheKey struct {
	id  mochi.Id
	key interface{}
}

type Context struct {
	node *node
}

func (ctx *Context) Get(key interface{}) Config {
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
	mu          sync.Mutex
	node        *node
	ids         map[viewCacheKey]mochi.Id
	prevIds     map[viewCacheKey]mochi.Id
	nodes       map[mochi.Id]*node
	prevNodes   map[mochi.Id]*node
	maxId       mochi.Id
	updateFlags map[mochi.Id]updateFlag
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
	root.updateFlags = map[mochi.Id]updateFlag{id: buildFlag}
	return root
}

func (root *root) addFlag(id mochi.Id, f updateFlag) {
	root.mu.Lock()
	defer root.mu.Unlock()

	root.updateFlags[id] |= f
}

func (root *root) update(size layout.Point) {
	root.mu.Lock()
	defer root.mu.Unlock()

	var flag updateFlag
	for _, v := range root.updateFlags {
		flag |= v
	}
	// fmt.Println("RunLoop", flag.needsBuild(), flag.needsLayout(), flag.needsPaint())

	if flag.needsBuild() {
		root.build()
	}
	if flag.needsLayout() {
		root.layout(layout.Pt(0, 0), size)
	}
	if flag.needsPaint() {
		root.paint()
	}
	root.updateFlags = map[mochi.Id]updateFlag{}
}

func (root *root) renderNode() *RenderNode {
	root.mu.Lock()
	defer root.mu.Unlock()

	return root.node.renderNode()
}

func (root *root) build() {
	root.prevIds = root.ids
	root.ids = map[viewCacheKey]mochi.Id{}
	root.prevNodes = root.nodes
	root.nodes = map[mochi.Id]*node{}
	root.node.build()

	keys := map[mochi.Id]viewCacheKey{}
	for k, v := range root.ids {
		keys[v] = k
	}
	for k, v := range root.prevIds {
		keys[v] = k
	}

	ids := map[viewCacheKey]mochi.Id{}
	for k := range root.nodes {
		ids[keys[k]] = k
	}
	root.ids = ids
}

func (root *root) paint() {
	root.node.paint()
}

func (root *root) layout(minSize layout.Point, maxSize layout.Point) {
	g := root.node.layout(minSize, maxSize)
	g.Frame = g.Frame.Add(layout.Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	root.node.layoutGuide = &g
}

func (root *root) newId() mochi.Id {
	root.maxId += 1
	return root.maxId
}

type node struct {
	id   mochi.Id
	root *root
	view View

	buildId   int64
	viewModel *Model
	children  map[mochi.Id]*node

	layoutId    int64
	layoutChan  chan struct{}
	layoutDone  chan struct{}
	layoutGuide *layout.Guide

	paintId      int64
	paintChan    chan struct{}
	paintDone    chan struct{}
	paintOptions paint.Style
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
		BuildId:      n.buildId,
		LayoutId:     n.layoutId,
		PaintId:      n.paintId,
		Children:     map[mochi.Id]*RenderNode{},
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
		n.buildId += 1

		// Generate the new viewModel.
		viewModel := n.view.Build(&Context{node: n})

		// Diff the old children (n.children) with new children (viewModel.Children).
		addedIds := []mochi.Id{}
		removedIds := []mochi.Id{}
		unchangedIds := []mochi.Id{}
		for id := range n.children {
			if _, ok := viewModel.Children[id]; !ok {
				removedIds = append(removedIds, id)
			} else {
				unchangedIds = append(unchangedIds, id)
			}
		}
		for id := range viewModel.Children {
			if _, ok := n.children[id]; !ok {
				addedIds = append(addedIds, id)
			}
		}

		children := map[mochi.Id]*node{}
		// Add build contexts for new children.
		for _, id := range addedIds {
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
				children: map[mochi.Id]*node{},
				root:     n.root,
			}
		}
		// Reuse old context for unupdated keys.
		for _, id := range unchangedIds {
			children[id] = n.children[id]
		}

		// Mark all children as needing rebuild since we rebuilt.
		for k := range children {
			n.root.updateFlags[k] |= buildFlag
		}

		// Watch for layout changes.
		if n.layoutChan != nil {
			n.viewModel.Layouter.Unnotify(n.layoutChan)
			close(n.layoutDone)
			n.layoutChan = nil
			n.layoutDone = nil
		}
		if viewModel.Layouter != nil {
			layoutChan := viewModel.Layouter.Notify()
			if layoutChan != nil {
				layoutDone := make(chan struct{})
				go func() {
				loop:
					for {
						select {
						case <-layoutChan:
							n.root.addFlag(n.id, layoutFlag)
							layoutChan <- struct{}{}
						case <-layoutDone:
							break loop
						}
					}
				}()
				n.layoutChan = layoutChan
				n.layoutDone = layoutDone
			}
		}

		// Watch for paint changes.
		if n.paintChan != nil {
			n.viewModel.Painter.Unnotify(n.paintChan)
			close(n.paintDone)
			n.paintChan = nil
			n.paintDone = nil
		}
		if viewModel.Painter != nil {
			paintChan := viewModel.Painter.Notify()
			if paintChan != nil {
				paintDone := make(chan struct{})
				go func() {
				loop:
					for {
						select {
						case <-paintChan:
							n.root.addFlag(n.id, paintFlag)
							paintChan <- struct{}{}
						case <-paintDone:
							break loop
						}
					}
				}()
				n.paintChan = paintChan
				n.paintDone = paintDone
			}
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

func (n *node) layout(minSize layout.Point, maxSize layout.Point) layout.Guide {
	n.layoutId += 1

	// Create the LayoutContext
	ctx := &layout.Context{
		MinSize:  minSize,
		MaxSize:  maxSize,
		ChildIds: []mochi.Id{},
		LayoutFunc: func(id mochi.Id, minSize, maxSize layout.Point) layout.Guide {
			// n :=  // TODO(KD): FIX!!!!!!!!!!
			return n.children[id].layout(minSize, maxSize)
		},
	}
	for i := range n.children {
		ctx.ChildIds = append(ctx.ChildIds, i)
	}

	// Perform layout
	layouter := n.viewModel.Layouter
	if layouter == nil {
		layouter = &full.Layout{}
	}
	g, gs := layouter.Layout(ctx)
	g = g.Fit(ctx)

	// Assign guides to children
	for k, v := range gs {
		guide := v
		n.children[k].layoutGuide = &guide
	}
	return g
}

func (n *node) paint() {
	if n.root.updateFlags[n.id].needsPaint() {
		n.paintId += 1

		if p := n.viewModel.Painter; p != nil {
			n.paintOptions = p.PaintStyle()
		} else {
			n.paintOptions = paint.Style{}
		}
	}

	// Recursively update children
	for _, v := range n.children {
		v.paint()
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
