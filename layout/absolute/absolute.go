package absolute // import "gomatcha.io/matcha/layout/absolute"

import (
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/view"
)

type Layouter struct {
	Guide       layout.Guide
	childGuides map[matcha.Id]layout.Guide
	views       []view.View
}

func (l *Layouter) Add(v view.View, g layout.Guide) {
	if l.childGuides == nil {
		l.childGuides = map[matcha.Id]layout.Guide{}
	}
	l.childGuides[v.Id()] = g
	l.views = append(l.views, v)
}

func (l *Layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	return l.Guide, l.childGuides
}

func (l *Layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *Layouter) Unnotify(id comm.Id) {
	// no-op
}
