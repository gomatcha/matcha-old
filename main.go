package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

type View struct {
	Frame    Rect
	Bounds   Rect
	Layouter Layouter
}

type Rect struct {
	Size  Size
	Point Point
}

type Size struct {
	Width  float64
	Height float64
}

type Point struct {
	Width  float64
	Height float64
}

type Layouter interface {
	NeedsLayout() bool
	Layout(minSize Size, maxSize Size) Size
}
