package absolute

import (
	"github.com/overcyn/mochi"
)

type Layout struct {
	Guide       layout.Guide
	ChildGuides map[mochi.Id]layout.Guide
}

func (l *Layout) Layout(ctx *layout.LayoutContext) (layout.Guide, map[mochi.Id]layout.Guide) {
	return l.Guide, l.ChildGuides
}

func (l *Layout) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *Layout) Unnotify(chan struct{}) {
	// no-op
}
