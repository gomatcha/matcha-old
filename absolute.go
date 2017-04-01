package mochi

type AbsoluteLayout struct {
	Guide       Guide
	ChildGuides map[interface{}]Guide
}

func NewAbsoluteLayout() *AbsoluteLayout {
	return &AbsoluteLayout{
		ChildGuides: map[interface{}]Guide{},
	}
}

func (l *AbsoluteLayout) Layout(ctx *LayoutContext) (Guide, map[interface{}]Guide) {
	return l.Guide, l.ChildGuides
}
