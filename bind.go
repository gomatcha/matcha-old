package mochi

type BoundNode struct {
	Children map[interface{}]View
	Layouter Layouter
	Painter  Painter
	Handlers map[interface{}]Handler

	// Context map[string] interface{}
	// Accessibility
	// Gesture Recognizers
	// OnAboutToScrollIntoView??
	// LayoutData?

	nodeChildren map[interface{}]*Node
	layoutGuide  Guide
	paintOptions PaintOptions
}
