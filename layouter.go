package mochi

type LayoutContext struct {
	MinSize   Point
	MaxSize   Point
	ChildKeys []interface{}
	node      *Node
}

type Layouter interface {
	Layout(ctx *LayoutContext) (Guide, map[interface{}]Guide)
}

func (l *LayoutContext) LayoutChild(k interface{}, minSize, maxSize Point) Guide {
	n := l.node.NodeChildren[k]
	return n.layout(minSize, maxSize)
}

// Full Layout

type FullLayout struct {
	needsLayoutFunc func()
}

func (l *FullLayout) NeedsLayoutFunc(f func()) {
	l.needsLayoutFunc = f
}

func (l *FullLayout) Layout(ctx *LayoutContext) (Guide, map[interface{}]Guide) {
	g := Guide{Frame: Rect{Max: ctx.MinSize}}
	gs := map[interface{}]Guide{}
	for k := range ctx.ChildKeys {
		gs[k] = ctx.LayoutChild(k, ctx.MinSize, ctx.MinSize)
	}
	return g, gs
}

// Guides

type Guide struct {
	Frame  Rect
	Insets Insets
	ZIndex int
	// Transform?
}

func (g Guide) Left() float64 {
	return g.Frame.Min.X
}
func (g Guide) Right() float64 {
	return g.Frame.Max.X
}
func (g Guide) Top() float64 {
	return g.Frame.Min.Y
}
func (g Guide) Bottom() float64 {
	return g.Frame.Max.Y
}
func (g Guide) Width() float64 {
	return g.Frame.Max.X - g.Frame.Min.X
}
func (g Guide) Height() float64 {
	return g.Frame.Max.Y - g.Frame.Min.Y
}
func (g Guide) CenterX() float64 {
	return (g.Frame.Min.X - g.Frame.Max.X) / 2
}
func (g Guide) CenterY() float64 {
	return (g.Frame.Min.Y - g.Frame.Max.Y) / 2
}
