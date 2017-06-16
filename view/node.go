package view

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	google_protobuf "github.com/golang/protobuf/ptypes/any"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/layout/full"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochibridge"
)

var MainMu sync.Mutex

var maxId int64
var maxFuncId int64

type Middleware interface {
	Build(*Context, *Model)
}

var middlewaresMu sync.Mutex
var middlewares = []Middleware{}

func RegisterMiddleware(v Middleware) {
	middlewaresMu.Lock()
	defer middlewaresMu.Unlock()

	middlewares = append(middlewares, v)
}

func defaultMiddlewares() []Middleware {
	middlewaresMu.Lock()
	defer middlewaresMu.Unlock()

	return middlewares
}

type Root struct {
	id     int64
	root   *root
	size   layout.Point
	ticker *internal.Ticker
}

func NewRoot(s Screen) *Root {
	r := &Root{
		root: newRoot(s),
		id:   atomic.AddInt64(&maxId, 1),
	}
	r.Start()
	return r
}

func (r *Root) Start() {
	MainMu.Lock()
	defer MainMu.Unlock()

	if r.ticker != nil {
		return
	}

	id := r.id
	r.ticker = internal.NewTicker(time.Hour * 99999)
	_ = r.ticker.Notify(func() {
		MainMu.Lock()
		defer MainMu.Unlock()

		if !r.root.update(r.size) {
			// nothing changed
			return
		}

		pb, err := r.root.MarshalProtobuf2()
		if err != nil {
			fmt.Println("err", err)
			return
		}
		mochibridge.Bridge().Call("updateId:withProtobuf:", mochibridge.Int64(id), mochibridge.Bytes(pb))
	})
}

func (r *Root) Stop() {
	MainMu.Lock()
	defer MainMu.Unlock()

	if r.ticker == nil {
		return
	}
	r.ticker.Stop()
}

func (r *Root) Call(funcId int64, viewId int64, args []reflect.Value) []reflect.Value {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.root.call(funcId, viewId, args)
}

func (r *Root) Id() int64 {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.id
}

func (r *Root) Size() layout.Point {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.size
}

func (r *Root) SetSize(p layout.Point) {
	MainMu.Lock()
	defer MainMu.Unlock()

	r.size = p
}

func (r *Root) Middlewares() []Middleware {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.root.middlewares
}

func (r *Root) SetMiddlewares(rs []Middleware) {
	MainMu.Lock()
	defer MainMu.Unlock()

	r.root.middlewares = rs
}

type viewCacheKey struct {
	id  mochi.Id
	key interface{}
}

type Context struct {
	node      *node
	prevIds   map[viewCacheKey]mochi.Id
	prevNodes map[mochi.Id]*node
	skipBuild map[mochi.Id]struct{}
}

func (ctx *Context) Prev(key interface{}) View {
	if ctx == nil {
		return nil
	}
	cacheKey := viewCacheKey{key: key, id: ctx.node.id}
	prevId := ctx.prevIds[cacheKey]
	prevNode := ctx.prevNodes[prevId]
	if prevNode == nil {
		return nil
	}

	v := prevNode.view
	for {
		if pv, ok := v.(*painterView); ok {
			v = pv.View
			continue
		} else if vv, ok := v.(*valuesView); ok {
			v = vv.View
			continue
		}
		break
	}

	return prevNode.view
}

func (ctx *Context) PrevModel() *Model {
	return ctx.node.model
}

func (ctx *Context) NewId(key interface{}) mochi.Id {
	id := mochi.Id(atomic.AddInt64(&maxId, 1))
	if ctx != nil {
		cacheKey := viewCacheKey{key: key, id: ctx.node.id}
		if _, ok := ctx.node.root.ids[cacheKey]; ok {
			fmt.Println("Context.NewId(): key has already been used", key)
		}
		ctx.node.root.ids[cacheKey] = id
	}
	return id
}

func (ctx *Context) NewFuncId() int64 {
	return atomic.AddInt64(&maxFuncId, 1)
}

func (ctx *Context) SkipBuild(ids []mochi.Id) {
	if ctx.skipBuild == nil {
		ctx.skipBuild = map[mochi.Id]struct{}{}
	}
	for _, i := range ids {
		ctx.skipBuild[i] = struct{}{}
	}
}

func (ctx *Context) Id() mochi.Id {
	return ctx.node.id
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
	ids         map[viewCacheKey]mochi.Id
	nodes       map[mochi.Id]*node
	middlewares []Middleware

	flagMu      sync.Mutex
	updateFlags map[mochi.Id]updateFlag
}

func newRoot(s Screen) *root {
	s.Lock()
	defer s.Unlock()

	v := s.View(nil, "root")

	root := &root{}
	root.node = &node{
		id:   v.Id(),
		view: v,
		root: root,
	}
	root.updateFlags = map[mochi.Id]updateFlag{v.Id(): buildFlag}
	root.middlewares = defaultMiddlewares()
	return root
}

func (root *root) addFlag(id mochi.Id, f updateFlag) {
	root.flagMu.Lock()
	defer root.flagMu.Unlock()

	root.updateFlags[id] |= f
}

func (root *root) update(size layout.Point) bool {
	root.flagMu.Lock()
	defer root.flagMu.Unlock()

	var flag updateFlag
	for _, v := range root.updateFlags {
		flag |= v
	}

	updated := false
	if flag.needsBuild() {
		root.build()
		updated = true
	}
	if flag.needsLayout() {
		root.layout(size, size)
		updated = true
	}
	if flag.needsPaint() {
		root.paint()
		updated = true
	}
	root.updateFlags = map[mochi.Id]updateFlag{}
	return updated
}

func (root *root) MarshalProtobuf2() ([]byte, error) {
	return proto.Marshal(root.MarshalProtobuf())
}

func (root *root) MarshalProtobuf() *pb.Root {
	return &pb.Root{
		Node: root.node.MarshalProtobuf(),
	}
}

func (root *root) build() {
	prevIds := root.ids
	prevNodes := root.nodes

	root.ids = map[viewCacheKey]mochi.Id{}
	root.nodes = map[mochi.Id]*node{
		root.node.id: root.node,
	}

	// Rebuild
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
		key, ok := keys[k]
		if ok {
			ids[key] = k
		}
	}
	root.ids = ids
}

func (root *root) layout(minSize layout.Point, maxSize layout.Point) {
	g := root.node.layout(minSize, maxSize)
	g.Frame = g.Frame.Add(layout.Pt(-g.Frame.Min.X, -g.Frame.Min.Y)) // Move Frame.Min to the origin.
	root.node.layoutGuide = &g
}

func (root *root) paint() {
	root.node.paint()
}

func (root *root) call(funcId int64, viewId int64, args []reflect.Value) []reflect.Value {
	node, ok := root.nodes[mochi.Id(viewId)]
	if !ok || node.model == nil {
		return nil
	}

	f, ok := node.model.NativeFuncs[funcId]
	if !ok {
		return nil
	}
	v := reflect.ValueOf(f)

	return v.Call(args)
}

type node struct {
	id    mochi.Id
	root  *root
	view  View
	stage Stage

	buildId       int64
	buildNotify   bool
	buildNotifyId comm.Id
	model         *Model
	children      map[mochi.Id]*node

	layoutId       int64
	layoutNotify   bool
	layoutNotifyId comm.Id
	layoutGuide    *layout.Guide

	paintId       int64
	paintNotify   bool
	paintNotifyId comm.Id
	paintOptions  paint.Style
}

func (n *node) MarshalProtobuf() *pb.Node {
	children := []*pb.Node{}
	for _, v := range n.children {
		children = append(children, v.MarshalProtobuf())
	}

	var nativeViewState *any.Any
	if a, err := ptypes.MarshalAny(n.model.NativeViewState); err == nil {
		nativeViewState = a
	}

	nativeValues := map[string]*google_protobuf.Any{}
	for k, v := range n.model.NativeValues {
		a, err := ptypes.MarshalAny(v)
		if err != nil {
			fmt.Println("Error enocding native value: ", err)
			continue
		}
		nativeValues[k] = a
	}

	return &pb.Node{
		Id:          int64(n.id),
		BuildId:     n.buildId,
		LayoutId:    n.layoutId,
		PaintId:     n.paintId,
		Children:    children,
		LayoutGuide: n.layoutGuide.MarshalProtobuf(),
		PaintStyle:  n.paintOptions.MarshalProtobuf(),
		BridgeName:  n.model.NativeViewName,
		BridgeValue: nativeViewState,
		Values:      nativeValues,
	}
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
		ctx := &Context{node: n, prevIds: prevIds, prevNodes: prevNodes}
		viewModel := n.view.Build(ctx)
		viewModelChildren := map[mochi.Id]View{} // TODO: Do this without maps.
		for _, i := range viewModel.Children {
			viewModelChildren[i.Id()] = i
		}

		// Call middleware
		for _, i := range n.root.middlewares {
			i.Build(ctx, viewModel)
		}

		// Diff the old children (n.children) with new children (viewModelChildren).
		addedIds := []mochi.Id{}
		removedIds := []mochi.Id{}
		unchangedIds := []mochi.Id{}
		for id := range n.children {
			if _, ok := viewModelChildren[id]; !ok {
				removedIds = append(removedIds, id)
			} else {
				unchangedIds = append(unchangedIds, id)
			}
		}
		for id := range viewModelChildren {
			if _, ok := n.children[id]; !ok {
				addedIds = append(addedIds, id)
			}
		}

		children := map[mochi.Id]*node{}
		// Add build contexts for new children.
		for _, id := range addedIds {
			var view View
			for _, i := range viewModelChildren {
				if i.Id() == id {
					view = i
					break
				}
			}

			children[id] = &node{
				id:   id,
				view: view,
				root: n.root,
			}

			// Mark as needing rebuild
			n.root.updateFlags[id] |= buildFlag
		}
		// Reuse old context for unupdated keys.
		for _, id := range unchangedIds {
			children[id] = n.children[id]

			// Mark as needing rebuild
			if _, ok := ctx.skipBuild[id]; !ok {
				n.root.updateFlags[id] |= buildFlag
			}
		}
		// Send lifecycle event to removed childern.
		for _, id := range removedIds {
			n.children[id].done()
		}

		// Watch for build changes, if we haven't
		if !n.buildNotify {
			n.buildNotifyId = n.view.Notify(func() {
				n.root.addFlag(n.id, buildFlag)
			})
			n.buildNotify = true
		}

		// Watch for layout changes.
		if n.layoutNotify {
			n.model.Layouter.Unnotify(n.layoutNotifyId)
			n.layoutNotify = false
		}
		if viewModel.Layouter != nil {
			n.layoutNotifyId = viewModel.Layouter.Notify(func() {
				n.root.addFlag(n.id, layoutFlag)
			})
			n.layoutNotify = true
		}

		// Watch for paint changes.
		if n.paintNotify {
			n.model.Painter.Unnotify(n.paintNotifyId)
			n.paintNotify = false
		}
		if viewModel.Painter != nil {
			n.paintNotifyId = viewModel.Painter.Notify(func() {
				n.root.addFlag(n.id, paintFlag)
			})
			n.paintNotify = true
		}

		n.children = children
		n.model = viewModel
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
	layouter := n.model.Layouter
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

		if p := n.model.Painter; p != nil {
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

	if n.buildNotify {
		n.view.Unnotify(n.buildNotifyId)
	}
	if n.layoutNotify {
		n.model.Layouter.Unnotify(n.layoutNotifyId)
	}
	if n.paintNotify {
		n.model.Painter.Unnotify(n.paintNotifyId)
	}

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

	str := fmt.Sprintf("{%p Id:%v View:%v Node:%p}", n, n.id, n.view, n.model)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}
