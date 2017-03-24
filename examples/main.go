package main

import (
	"fmt"
	_ "github.com/overcyn/mochi"
)

func main() {
	fmt.Println("Hello, 世界")
	// mochi.Display(nil)
}

// type CustomView struct {}

// func (v *CustomView) Render() Renderer {
// 	return &ScrollView{
// 		Children: map[string]View{
// 			"1": &TextField{},
// 			"2": &TextField{},
// 			"3": &TextField{},
// 			"4": &TextField{},
// 		},
// 		OnScroll: func() {
// 		},
// 	}
// }

type ComplexDrawing struct {}
func (v *ComplexDrawing) Render() Renderer {
	c := make(map[string]View)
	r := &mochi.AbsoluteLayouter{}

	c["1"] = &PathView{
		path: Circle(0, 0, 10)
	}
	r.guides["1"] = Rt()

	c["2"] = &Path{}

	layouter = r
	return
}

type Rect struct {
	// fill: '#06538e',
	// width: 125,
	// height: 125,
	// stroke: 'red',
	// strokeDashArray: [5, 5]
}

// type View struct {
// 	Frame       layout.Rect
// 	Bounds      layout.Rect
// 	Children    map[string]View
// 	needsRender func()
// }

// func (v *View) NeedsRenderFunc(f func()) {
// 	v.needsRender = f
// }

// func (v *View) Render() Renderer {
// 	return nil
// }
