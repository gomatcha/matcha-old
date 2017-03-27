package mochi

type View interface {
	Update(*UpdateContext) map[string]View
	// Painter() Painter   // immutable object that paints
	Layouter() Layouter // immutable object that layouts
	Handler() Handler   // immutable object that handles events
}

type UpdateContext struct {
	Children map[string]View
}

// Immutable tree...
type node struct {
	children map[string]*node
	layouter Layouter
	guide    *Guide
}

func nodeFromView(view View) *node {
	chl := make(map[string]*node)
	for k, v := range view.Update(nil) {
		chl[k] = nodeFromView(v)
	}

	return &node{
		children: chl,
		layouter: view.Layouter(),
	}
}

func (n *node) layout(maxSize Size, minSize Size) {
	// Create the LayoutContext
	chl := make(map[string]*LayoutChild)
	for k, v := range n.children {
		chl[k] = &LayoutChild{
			node: v,
		}
	}
	ctx := &LayoutContext{
		MaxSize:  maxSize,
		MinSize:  minSize,
		Children: chl,
	}

	// Perform layout
	l := n.layouter
	if l == nil {
		l = &FullLayout{}
	}
	g, _ := l.Layout(ctx)
	n.guide = &g
}

func Display(v View) *node {
	node := nodeFromView(v)
	_ = node

	return nil

	// Generate a mutable tree
	// Copy into an immutable tree
	// Run a layout pass on the immutable tree
	// Run a paint pass on the immutable tree
}
