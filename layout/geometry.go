package layout

import (
	"fmt"

	"github.com/overcyn/mochi/pb"
)

type Rect struct {
	Min, Max Point
}

func Rt(x0, y0, x1, y1 float64) Rect {
	return Rect{Min: Point{X: x0, Y: y0}, Max: Point{X: x1, Y: y1}}
}

func (r Rect) EncodeProtobuf() *pb.Rect {
	return &pb.Rect{
		Min: r.Min.EncodeProtobuf(),
		Max: r.Max.EncodeProtobuf(),
	}
}

func (r Rect) Add(p Point) Rect {
	n := r
	n.Min.X += p.X
	n.Min.Y += p.Y
	n.Max.X += p.X
	n.Max.Y += p.Y
	return n
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{%v, %v, %v, %v}", r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
}

type Point struct {
	X float64
	Y float64
}

func Pt(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (p Point) EncodeProtobuf() *pb.Point {
	return &pb.Point{
		X: p.X,
		Y: p.Y,
	}
}

type Insets struct {
	Top    float64
	Left   float64
	Bottom float64
	Right  float64
}

func In(top, left, bottom, right float64) Insets {
	return Insets{Top: top, Left: left, Bottom: bottom, Right: right}
}

func (in Insets) EncodeProtobuf() *pb.Insets {
	return &pb.Insets{
		Top:    in.Top,
		Left:   in.Left,
		Bottom: in.Bottom,
		Right:  in.Right,
	}
}
