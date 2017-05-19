package layout

import (
	"fmt"

	"github.com/overcyn/mochi/layout/encoding"
	capnp "zombiezen.com/go/capnproto2"
)

type Rect struct {
	Min, Max Point
}

func Rt(x0, y0, x1, y1 float64) Rect {
	return Rect{Min: Point{X: x0, Y: y0}, Max: Point{X: x1, Y: y1}}
}

func (r Rect) MarshalCapnp(s *capnp.Segment) (encoding.Rect, error) {
	rect, err := encoding.NewRect(s)
	if err != nil {
		return encoding.Rect{}, err
	}

	min, err := r.Min.MarshalCapnp(s)
	if err != nil {
		return encoding.Rect{}, err
	}
	if err = rect.SetMin(min); err != nil {
		return encoding.Rect{}, err
	}

	max, err := r.Max.MarshalCapnp(s)
	if err != nil {
		return encoding.Rect{}, err
	}
	if err = rect.SetMax(max); err != nil {
		return encoding.Rect{}, err
	}
	return rect, nil
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

func (p Point) MarshalCapnp(s *capnp.Segment) (encoding.Point, error) {
	point, err := encoding.NewPoint(s)
	if err != nil {
		return encoding.Point{}, err
	}
	point.SetX(p.X)
	point.SetY(p.Y)
	return point, nil
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

func (in Insets) MarshalCapnp(s *capnp.Segment) (encoding.Insets, error) {
	insets, err := encoding.NewInsets(s)
	if err != nil {
		return encoding.Insets{}, err
	}
	insets.SetTop(in.Top)
	insets.SetLeft(in.Left)
	insets.SetBottom(in.Bottom)
	insets.SetRight(in.Right)
	return insets, nil
}
