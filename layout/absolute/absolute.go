package absolute

import (
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/view"
)

type Layout struct {
	Guide       layout.Guide
	childGuides map[matcha.Id]layout.Guide
	views       []view.View
}

func (l *Layout) Add(v view.View, g layout.Guide) {
	if l.childGuides == nil {
		l.childGuides = map[matcha.Id]layout.Guide{}
	}
	l.childGuides[v.Id()] = g
	l.views = append(l.views, v)
}

func (l *Layout) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	return l.Guide, l.childGuides
}

func (l *Layout) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *Layout) Unnotify(id comm.Id) {
	// no-op
}
