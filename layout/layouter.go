package layout

import (
	"reflect"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout/encoding"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterType("layout.Point", reflect.TypeOf(Point{}))
	mochibridge.RegisterType("layout.Rect", reflect.TypeOf(Rect{}))
}

type Context struct {
	MinSize    Point
	MaxSize    Point
	ChildIds   []mochi.Id
	LayoutFunc func(mochi.Id, Point, Point) Guide
}

type Layouter interface {
	Layout(ctx *Context) (Guide, map[mochi.Id]Guide)
	mochi.Notifier
}

func (l *Context) LayoutChild(id mochi.Id, minSize, maxSize Point) Guide {
	return l.LayoutFunc(id, minSize, maxSize)
}

type Guide struct {
	Frame  Rect
	Insets Insets
	ZIndex int
	// Transform?
}

func (g Guide) EncodeProtobuf() *encoding.Guide {
	return &encoding.Guide{
		Frame:  g.Frame.EncodeProtobuf(),
		Insets: g.Insets.EncodeProtobuf(),
		ZIndex: int64(g.ZIndex),
	}
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
