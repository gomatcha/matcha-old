package view

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/layout/full"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochibridge"
)

var marshallers []func(*Model) (string, []byte)

func RegisterMarshaller(f func(*Model) (string, []byte)) {
	marshallers = append(marshallers, f)
}

type RenderNode struct {
	Id           mochi.Id
	BuildId      int64
	LayoutId     int64
	PaintId      int64
	Children     map[mochi.Id]*RenderNode
	BridgeName   string
	BridgeState  interface{}
	Values       map[string][]byte
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

type Root struct {
	id         int
	mu         *sync.Mutex
	root       *root
	renderNode *RenderNode
	size       layout.Point
	ticker     *internal.Ticker
}

func NewRoot(f func(Config) View, id int) *Root {
	vc := &Root{
		mu:     &sync.Mutex{},
		root:   newRoot(f),
		ticker: internal.NewTicker(time.Hour * 99999),
		id:     id,
	}

	// Start run loop.
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

func (vc *Root) Render() *RenderNode {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.root.update(vc.size)
	rn := vc.root.renderNode()
	return rn
}

func (vc *Root) SetSize(p layout.Point) {
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
	node      *node
	prevIds   map[viewCacheKey]mochi.Id
	prevNodes map[mochi.Id]*node
}

func (ctx *Context) Get(key interface{}) Config {
	return Config{
		Prev: ctx.Prev(key),
		Embed: &Embed{
			root: ctx.node.root,
			id:   ctx.NewId(key),
		},
	}
}

func (ctx *Context) Prev(key interface{}) View {
	cacheKey := viewCacheKey{key: key, id: ctx.node.id}
	prevId := ctx.prevIds[cacheKey]
	prevCtx := ctx.prevNodes[prevId]
	if prevCtx != nil {
		return prevCtx.view
	}
	return nil
}

func (ctx *Context) NewId(key interface{}) mochi.Id {
	cacheKey := viewCacheKey{key: key, id: ctx.node.id}
	id := ctx.node.root.newId()
	ctx.node.root.ids[cacheKey] = id
	return id
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
	nodes       map[mochi.Id]*node
	maxId       mochi.Id
	updateFlags map[mochi.Id]updateFlag
}

func newRoot(f func(Config) View) *root {
	root := &root{}

	id := root.newId()
	cfg := Config{Embed: &Embed{
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

	// Lock the entire tree.
	root.node.lock()
	defer root.node.unlock()

	var flag updateFlag
	for _, v := range root.updateFlags {
		flag |= v
	}

	if flag.needsBuild() {
		root.buildLocked()
	}
	if flag.needsLayout() {
		root.layoutLocked(layout.Pt(0, 0), size)
	}
	if flag.needsPaint() {
		root.paintLocked()
	}
	root.updateFlags = map[mochi.Id]updateFlag{}
}

func (root *root) renderNode() *RenderNode {
	root.mu.Lock()
	defer root.mu.Unlock()

	return root.node.renderNode()
}

func (root *root) buildLocked() {
	prevIds := root.ids
	prevNodes := root.nodes

	root.ids = map[viewCacheKey]mochi.Id{}
	root.nodes = map[mochi.Id]*node{}
	root.node.build(prevIds, prevNodes)

	keys := map[mochi.Id]viewCacheKey{}
	for k, v := range root.ids {
		keys[v] = k
	}
	for k, v := range prevIds {
		keys[v] = k
	}

	ids := map[viewCacheKey]mochi.Id{}
	for k := range root.nodes {
		ids[keys[k]] = k
	}
	root.ids = ids
}

func (root *root) layoutLocked(minSize layout.Point, maxSize layout.Point) {
	g := root.node.layout(minSize, maxSize)
	g.Frame = g.Frame.Add(layout.Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	root.node.layoutGuide = &g
}

func (root *root) paintLocked() {
	root.node.paint()
}

func (root *root) newId() mochi.Id {
	root.maxId += 1
	return root.maxId
}

type node struct {
	id    mochi.Id
	root  *root
	view  View
	stage Stage

	buildId   int64
	buildChan chan struct{}
	buildDone chan struct{}
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

func (n *node) renderNode() *RenderNode {
	values := map[string][]byte{}
	for _, f := range marshallers {
		key, value := f(n.viewModel)
		if key != "" {
			values[key] = value
		}
	}

	rn := &RenderNode{
		Id:           n.id,
		BuildId:      n.buildId,
		LayoutId:     n.layoutId,
		PaintId:      n.paintId,
		Children:     map[mochi.Id]*RenderNode{},
		BridgeName:   n.viewModel.BridgeName,
		BridgeState:  n.viewModel.BridgeState,
		Values:       values,
		LayoutGuide:  n.layoutGuide,
		PaintOptions: n.paintOptions,
	}
	for k, v := range n.children {
		rn.Children[k] = v.renderNode()
	}
	return rn
}

func (n *node) build(prevIds map[viewCacheKey]mochi.Id, prevNodes map[mochi.Id]*node) {
	if n.root.updateFlags[n.id].needsBuild() {
		n.buildId += 1

		// Send lifecycle event to new children.
		if n.stage == StageDead {
			n.view.Lifecycle(n.stage, StageVisible)
			n.stage = StageVisible
		}

		// Generate the new viewModel.
		viewModel := n.view.Build(&Context{node: n, prevIds: prevIds, prevNodes: prevNodes})

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

			// Lock new child.
			view.Lock()

			children[id] = &node{
				id:   id,
				view: view,
				root: n.root,
			}
		}
		// Reuse old context for unupdated keys.
		for _, id := range unchangedIds {
			children[id] = n.children[id]
		}
		// Send lifecycle event to removed childern.
		for _, id := range removedIds {
			n.children[id].done()
		}

		// Mark all children as needing rebuild since we rebuilt.
		for k := range children {
			n.root.updateFlags[k] |= buildFlag
		}

		// Watch for build changes
		if n.buildChan == nil {
			buildChan := n.view.Notify()
			if buildChan != nil {
				buildDone := make(chan struct{})
				go func(id mochi.Id) {
				loop:
					for {
						select {
						case <-buildChan:
							n.root.addFlag(id, buildFlag)
							buildChan <- struct{}{}
						case <-buildDone:
							break loop
						}
					}
				}(n.view.Id())
				n.buildChan = buildChan
				n.buildDone = buildDone
			}
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
		i.build(prevIds, prevNodes)

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

func (n *node) done() {
	n.view.Lifecycle(n.stage, StageDead)
	n.stage = StageDead

	if n.buildChan != nil {
		n.view.Unnotify(n.buildChan)
		close(n.buildDone)
	}
	if n.layoutChan != nil {
		n.viewModel.Layouter.Unnotify(n.layoutChan)
		close(n.layoutDone)
	}
	if n.paintChan != nil {
		n.viewModel.Painter.Unnotify(n.paintChan)
		close(n.paintDone)
	}
	n.view.Unlock()

	for _, i := range n.children {
		i.done()
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

func (n *node) lock() {
	n.view.Lock()
	for _, i := range n.children {
		i.lock()
	}
}

func (n *node) unlock() {
	n.view.Unlock()
	for _, i := range n.children {
		i.unlock()
	}
}
