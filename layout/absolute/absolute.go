package absolute

import (
	"github.com/overcyn/mochi"
)

type Layout struct {
	Guide       mochi.Guide
	ChildGuides map[mochi.Id]mochi.Guide
}

func (l *Layout) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[mochi.Id]mochi.Guide) {
	return l.Guide, l.ChildGuides
}

func (l *Layout) Notify(c chan struct{}) {
	// no-op
}

func (l *Layout) Unnotify(chan struct{}) {
	// no-op
}
