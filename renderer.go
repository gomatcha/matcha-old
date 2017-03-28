package mochi

type View interface {
	Update(n *Node) *Node
	// UpdateLayouter(n *Layouter) *Layouter
	// UptadePainter(n *Painter) *Painter
	// UpdateHandlers(prev, next *Node)

	NeedsUpdate()
	// NeedsRehandle()
	// NeedsRelayout()
	// NeedsRepaint()
}

type Node struct {
	Children map[interface{}]View
	Layouter Layouter
	Painter  Painter
	Handlers map[interface{}]Handler

	// Context map[string] interface{}
	// Accessibility
	// Gesture Recognizers
	// OnAboutToScrollIntoView??
	// LayoutData?

	nodeChildren map[interface{}]*Node
	layoutGuide  Guide
	paintOptions PaintOptions
}

func NewNode() *Node {
	n := &Node{}
	n.Children = map[interface{}]View{}
	n.Handlers = map[interface{}]Handler{}
	n.nodeChildren = map[interface{}]*Node{}
	return n
}

func nodeFromView(view View, prev *Node) *Node {
	node := view.Update(prev)
	for k, v := range node.Children {
		node.nodeChildren[k] = nodeFromView(v, prev.nodeChildren[k])
	}
	return node
}

func (n *Node) layout(maxSize Size, minSize Size) Guide {
	// Create the LayoutContext
	ctx := &LayoutContext{
		MaxSize:   maxSize,
		MinSize:   minSize,
		ChildKeys: []interface{}{},
		node:      n,
	}
	for i := range n.nodeChildren {
		ctx.ChildKeys = append(ctx.ChildKeys, i)
	}

	// Perform layout
	layouter := n.Layouter
	if layouter == nil {
		layouter = &FullLayout{}
	}
	n.layoutGuide = layouter.Layout(ctx)
	return n.layoutGuide
}

func (n *Node) getPaintOptions() {
	if p := n.Painter; p != nil {
		n.paintOptions = p.PaintOptions()
	}
	for _, v := range n.nodeChildren {
		v.getPaintOptions()
	}
}

func Display(v View) *Node {
	node := nodeFromView(v, nil)
	node.layout(Sz(0, 0), Sz(0, 0))
	node.getPaintOptions()
	return node
}
