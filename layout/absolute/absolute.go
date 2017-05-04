package absolute

import (
	"github.com/overcyn/mochi"
)

type Layout struct {
	Guide       mochi.Guide
	ChildGuides map[interface{}]mochi.Guide
}

func (l *Layout) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	return l.Guide, l.ChildGuides
}
