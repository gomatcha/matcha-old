package mochi

type LayoutContext struct {
	MinSize  Size
	MaxSize  Size
	Children map[string]*LayoutChild
}

type LayoutChild struct {
	node *node
}

func (n *LayoutChild) Layout(min Size, max Size) Guide {
	return Guide{}
}

type Layouter interface {
	// Return your preferred size/insets and your children's frames. The origin of your guide is ignored.
	Layout(ctx *LayoutContext) (Guide, map[string]Guide)
	// Hook for layout engine to pass in a update callback. Do not call directly.
	NeedsLayoutFunc(func())
}

type FullLayout struct {
	needsLayoutFunc func()
}

func (l *FullLayout) NeedsLayoutFunc(f func()) {
	l.needsLayoutFunc = f
}

func (l *FullLayout) Layout(ctx *LayoutContext) (Guide, map[string]Guide) {
	g := Guide{Frame: Rect{Size: ctx.MinSize}}
	chl := map[string]Guide{}
	for k, child := range ctx.Children {
		chl[k] = ConstrainChild(child, Insets{}, []Constraint{
			{Top, Equal, g.Top()},
			{Bottom, Equal, g.Bottom()},
			{Left, Equal, g.Left()},
			{Right, Equal, g.Right()},
		})
	}
	return g, chl
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
