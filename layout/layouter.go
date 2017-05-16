package layout

import (
	"github.com/overcyn/mochi"
)

type Context struct {
	MinSize    mochi.Point
	MaxSize    mochi.Point
	ChildIds   []mochi.Id
	LayoutFunc func(mochi.Id, mochi.Point, mochi.Point) Guide
	// node     *node
}

type Layouter interface {
	Layout(ctx *Context) (Guide, map[mochi.Id]Guide)
	mochi.Notifier
}

func (l *Context) LayoutChild(id mochi.Id, minSize, maxSize mochi.Point) Guide {
	return l.LayoutFunc(id, minSize, maxSize)
}

type Guide struct {
	Frame  mochi.Rect
	Insets mochi.Insets
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
func (g Guide) Fit(ctx *Context) Guide {
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
