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
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/internal"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/layout/full"
	"github.com/overcyn/matcha/paint"
	pb "github.com/overcyn/matcha/pb/view"
	"github.com/overcyn/matchabridge"
)

var MainMu sync.Mutex

var maxId int64

// Middleware is called on the result of View.Build(*context).
type Middleware interface {
	Build(*Context, *Model)
}

var middlewaresMu sync.Mutex
var middlewares = []func() Middleware{}

// RegisterMiddleware adds v to the list of default middleware that Root starts with.
func RegisterMiddleware(v func() Middleware) {
	middlewaresMu.Lock()
	defer middlewaresMu.Unlock()

	middlewares = append(middlewares, v)
}

func defaultMiddlewares() []func() Middleware {
	middlewaresMu.Lock()
	defer middlewaresMu.Unlock()

	return middlewares
}

type Root struct {
	id     comm.Id
	root   *root
	size   layout.Point
	ticker *internal.Ticker
}

// NewRoot initializes a Root with screen s.
func NewRoot(s Screen) *Root {
	r := &Root{
		root: newRoot(s),
		id:   comm.Id(atomic.AddInt64(&maxId, 1)),
	}
	r.start()
	return r
}

func (r *Root) start() {
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
		matchabridge.Bridge().Call("updateId:withProtobuf:", matchabridge.Int64(int64(id)), matchabridge.Bytes(pb))

		fmt.Println(r.root.node.debugString())
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

func (r *Root) Call(funcId string, viewId int64, args []reflect.Value) []reflect.Value {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.root.call(funcId, viewId, args)
}

// Id returns the unique identifier for r.
func (r *Root) Id() comm.Id {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.id
}

// Size returns the size of r.
func (r *Root) Size() layout.Point {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.size
}

// SetSize sets the size of r.
func (r *Root) SetSize(p layout.Point) {
	MainMu.Lock()
	defer MainMu.Unlock()

	r.size = p
}

// Middlewares returns the Middleware that are applied to r's views.
func (r *Root) Middlewares() []Middleware {
	MainMu.Lock()
	defer MainMu.Unlock()

	return r.root.middlewares
}

// SetMiddlewares sets the list of Middleware applied to all of r's views.
func (r *Root) SetMiddlewares(rs []Middleware) {
	MainMu.Lock()
	defer MainMu.Unlock()

	r.root.middlewares = rs
}

type viewCacheKey struct {
	id  matcha.Id
	key string
}

// Context specifies the supporting context for building a View.
type Context struct {
	prefix string
	parent *Context

	node      *node
	prevIds   map[viewCacheKey]matcha.Id
	prevNodes map[matcha.Id]*node
	skipBuild map[matcha.Id]struct{}
}

// Prev returns the view returned by the last call to Build with the given key.
func (ctx *Context) Prev(key string) View {
	return ctx.prev(key, "")
}

func (ctx *Context) prev(key string, prefix string) View {
	if ctx.parent != nil {
		return ctx.parent.prev(key, ctx.prefix+"|"+prefix)
	}
	if ctx.node == nil {
		return nil
	}
	if prefix != "" {
		key = prefix + "|" + key
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

	return v
}

// PrevModel returns the last result of View.Build().
func (ctx *Context) PrevModel() *Model {
	if ctx.parent != nil {
		return ctx.PrevModel()
	}
	if ctx.node == nil {
		return nil
	}
	return ctx.node.model
}

// NewId generates a new identifier for a given key.
func (ctx *Context) NewId(key string) matcha.Id {
	return ctx.newId(key, "")
}

func (ctx *Context) newId(key string, prefix string) matcha.Id {
	if ctx.parent != nil {
		return ctx.parent.newId(key, ctx.prefix+"|"+prefix)
	}
	if prefix != "" {
		key = prefix + "|" + key
	}

	id := matcha.Id(atomic.AddInt64(&maxId, 1))
	if ctx.node != nil {
		cacheKey := viewCacheKey{key: key, id: ctx.node.id}
		if _, ok := ctx.node.root.ids[cacheKey]; ok {
			fmt.Println("Context.NewId(): key has already been used", key)
		}
		ctx.node.root.ids[cacheKey] = id
	}
	return id
}

// SkipBuild marks the child ids as not needing to be rebuilt.
func (ctx *Context) SkipBuild(ids []matcha.Id) {
	if ctx.parent != nil {
		ctx.parent.SkipBuild(ids)
		return
	}

	if ctx.skipBuild == nil {
		ctx.skipBuild = map[matcha.Id]struct{}{}
	}
	for _, i := range ids {
		ctx.skipBuild[i] = struct{}{}
	}
}

// WithPrefix returns a new Context. Calls to this Prev and NewId on this context will be prepended with key.
func (ctx *Context) WithPrefix(key string) *Context {
	return &Context{prefix: key, parent: ctx}
}

// Id returns the identifier associated with the build context.
func (ctx *Context) Id() matcha.Id {
	if ctx.parent != nil {
		return ctx.parent.Id()
	}
	if ctx.node == nil {
		return 0
	}
	return ctx.node.id
}

func (ctx *Context) Path() []matcha.Id {
	if ctx.parent != nil {
		return ctx.parent.Path()
	}
	if ctx.node == nil {
		return []matcha.Id{0}
	}
	return nil
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
	ids         map[viewCacheKey]matcha.Id
	nodes       map[matcha.Id]*node
	middlewares []Middleware

	flagMu      sync.Mutex
	updateFlags map[matcha.Id]updateFlag
}

func newRoot(s Screen) *root {
	s.Lock()
	defer s.Unlock()

	v := s.View(&Context{})
	id := v.Id()

	root := &root{}
	root.node = &node{
		id:   id,
		path: []matcha.Id{id},
		view: v,
		root: root,
	}
	root.updateFlags = map[matcha.Id]updateFlag{v.Id(): buildFlag}
	for _, i := range defaultMiddlewares() {
		root.middlewares = append(root.middlewares, i())
	}
	return root
}

func (root *root) addFlag(id matcha.Id, f updateFlag) {
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
	root.updateFlags = map[matcha.Id]updateFlag{}
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

	root.ids = map[viewCacheKey]matcha.Id{}
	root.nodes = map[matcha.Id]*node{
		root.node.id: root.node,
	}

	// Rebuild
	root.node.build(prevIds, prevNodes)

	keys := map[matcha.Id]viewCacheKey{}
	for k, v := range root.ids {
		keys[v] = k
	}
	for k, v := range prevIds {
		keys[v] = k
	}

	ids := map[viewCacheKey]matcha.Id{}
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

func (root *root) call(funcId string, viewId int64, args []reflect.Value) []reflect.Value {
	node, ok := root.nodes[matcha.Id(viewId)]
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
	id    matcha.Id
	path  []matcha.Id
	root  *root
	view  View
	stage Stage

	buildId       int64
	buildNotify   bool
	buildNotifyId comm.Id
	model         *Model
	children      map[matcha.Id]*node

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

func (n *node) build(prevIds map[viewCacheKey]matcha.Id, prevNodes map[matcha.Id]*node) {
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
		viewModelChildren := map[matcha.Id]View{} // TODO: Do this without maps.
		for _, i := range viewModel.Children {
			viewModelChildren[i.Id()] = i
		}

		// Call middleware
		for _, i := range n.root.middlewares {
			i.Build(ctx, viewModel)
		}

		// Diff the old children (n.children) with new children (viewModelChildren).
		addedIds := []matcha.Id{}
		removedIds := []matcha.Id{}
		unchangedIds := []matcha.Id{}
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

		children := map[matcha.Id]*node{}
		// Add build contexts for new children.
		for _, id := range addedIds {
			var view View
			for _, i := range viewModelChildren {
				if i.Id() == id {
					view = i
					break
				}
			}

			path := make([]matcha.Id, len(n.path)+1)
			copy(path, n.path)
			path[len(n.path)] = id

			children[id] = &node{
				id:   id,
				path: path,
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
		ChildIds: []matcha.Id{},
		LayoutFunc: func(id matcha.Id, minSize, maxSize layout.Point) layout.Guide {
			return n.children[id].layout(minSize, maxSize)
		},
	}
	for i := range n.children {
		ctx.ChildIds = append(ctx.ChildIds, i)
	}

	// Perform layout
	layouter := n.model.Layouter
	if layouter == nil {
		layouter = &full.Layouter{}
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

	str := fmt.Sprintf("{%p Id:%v View:%v Node:%p Layout:%v}", n, n.id, n.view, n.model, n.layoutGuide.Frame)
	if len(all) > 0 {
		str += "\n" + strings.Join(all, "\n")
	}
	return str
}
