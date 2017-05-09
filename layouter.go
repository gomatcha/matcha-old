package mochi

type LayoutContext struct {
	MinSize  Point
	MaxSize  Point
	ChildIds []Id
	node     *node
}

type NotifyLayouter interface {
	Layouter
	Notifier
}

type Layouter interface {
	Layout(ctx *LayoutContext) (Guide, map[Id]Guide)
}

func (l *LayoutContext) LayoutChild(id Id, minSize, maxSize Point) Guide {
	n := l.node.children[id] // TODO(KD): FIX!!!!!!!!!!
	return n.layout(minSize, maxSize)
}

// Full Layout

type FullLayout struct {
	needsLayoutFunc func()
}

func (l *FullLayout) NeedsLayoutFunc(f func()) {
	l.needsLayoutFunc = f
}

func (l *FullLayout) Layout(ctx *LayoutContext) (Guide, map[Id]Guide) {
	g := Guide{Frame: Rect{Max: ctx.MinSize}}
	gs := map[Id]Guide{}
	for _, id := range ctx.ChildIds {
		gs[id] = ctx.LayoutChild(id, ctx.MinSize, ctx.MinSize)
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
	return (g.Frame.Max.X - g.Frame.Min.X) / 2
}
func (g Guide) CenterY() float64 {
	return (g.Frame.Max.Y - g.Frame.Min.Y) / 2
}

// Fit adjusts the frame of the guide to be within MinSize and MaxSize of the LayoutContext.
func (g Guide) fit(ctx *LayoutContext) Guide {
	if g.Width() < ctx.MinSize.X {
		g.Frame.Max.X = ctx.MinSize.X - g.Frame.Min.X
	}
	if g.Height() < ctx.MinSize.Y {
		g.Frame.Max.Y = ctx.MinSize.Y - g.Frame.Min.Y
	}
	if g.Width() > ctx.MaxSize.X {
		g.Frame.Max.X = ctx.MaxSize.X - g.Frame.Max.X
	}
	if g.Height() > ctx.MaxSize.Y {
		g.Frame.Max.Y = ctx.MaxSize.Y - g.Frame.Max.Y
	}
	return g
}
