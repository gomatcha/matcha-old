package mochi

type View interface {
	Update(n *Node) *Node
	NeedsUpdate()
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

func (n *Node) layout(maxSize Point, minSize Point) Guide {
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
	g, gs := layouter.Layout(ctx)

	// Assign guides to children
	for k, v := range gs {
		n.nodeChildren[k].layoutGuide = v
	}
	return g
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
	node.layout(Pt(0, 0), Pt(0, 0))
	node.getPaintOptions()
	return node
}
