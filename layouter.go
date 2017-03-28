package mochi

type LayoutContext struct {
	MinSize  Size
	MaxSize  Size
	ChildKeys []interface{}
	node *Node
}

type Layouter interface {
	Layout(ctx *LayoutContext) Guide
}

func (l *LayoutContext)LayoutChild(k interface{}, minSize, maxSize Size) Guide {
	n := l.node.nodeChildren[k]
	if n == nil {
		return Guide{}
	}
	return n.Layout(minSize, maxSize)
}

type FullLayout struct {
	needsLayoutFunc func()
}

func (l *FullLayout) NeedsLayoutFunc(f func()) {
	l.needsLayoutFunc = f
}

func (l *FullLayout) Layout(ctx *LayoutContext) Guide {
	g := Guide{Frame: Rect{Size: ctx.MinSize}}
	for k := range ctx.ChildKeys {
		ConstrainChild(ctx, k, Insets{}, []Constraint{
			{Top, Equal, g.Top()},
			{Bottom, Equal, g.Bottom()},
			{Left, Equal, g.Left()},
			{Right, Equal, g.Right()},
		})
	}
	return g
}

// Guides

type Guide struct {
	Frame  Rect
	Insets Insets
}

func (g Guide) Left() float64 {
	return g.Frame.Origin.X
}
func (g Guide) Right() float64 {
	return g.Frame.Origin.X + g.Frame.Size.Width
}
func (g Guide) Top() float64 {
	return g.Frame.Origin.Y
}
func (g Guide) Bottom() float64 {
	return g.Frame.Origin.Y + g.Frame.Size.Height
}
func (g Guide) Width() float64 {
	return g.Frame.Size.Width
}
func (g Guide) Height() float64 {
	return g.Frame.Size.Height
}
func (g Guide) CenterX() float64 {
	return g.Frame.Origin.X + g.Frame.Size.Width/2
}
func (g Guide) CenterY() float64 {
	return g.Frame.Origin.Y + g.Frame.Size.Height/2
}
