package mochi

type Rect struct {
	Min, Max Point
}

func Rt(x0, y0, x1, y1 float64) Rect {
	return Rect{Min: Point{X: x0, Y: y0}, Max: Point{X: x1, Y: y1}}
}

func (r Rect) Add(p Point) Rect {
	n := r
	n.Min.X += p.X
	n.Min.Y += p.Y
	n.Max.X += p.X
	n.Max.Y += p.Y
	return n
}

type Point struct {
	X float64
	Y float64
}

func Pt(x, y float64) Point {
	return Point{X: x, Y: y}
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
