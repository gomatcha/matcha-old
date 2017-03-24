package mochi

type Rect struct {
	Origin Point
	Size   Size
}

func Rt(x, y, w, h float64) Rect {
	return Rect{Origin: Point{X: x, Y: y}, Size: Size{Width: w, Height: h}}
}

type Size struct {
	Width  float64
	Height float64
}

func Sz(w, h float64) Size {
	return Size{Width: w, Height: h}
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
