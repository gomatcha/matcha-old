package keyboard

import (
	"github.com/overcyn/matcha/comm"
)

// type key struct{}

// var Key = key{}

type Responder struct {
	visible bool
	value   comm.Value
}

// func (g *Responder) Next() {
// }

// func (g *Responder) Prev() {
// }

func (g *Responder) Show() {
	if !g.visible {
		g.visible = true
		g.value.Signal()
	}
}

func (g *Responder) Dismiss() {
	if g.visible {
		g.visible = false
		g.value.Signal()
	}
}

func (g *Responder) Visible() bool {
	return g.visible
}

func (g *Responder) Notify(f func()) comm.Id {
	return g.value.Notify(f)
}

func (g *Responder) Unnotify(id comm.Id) {
	g.value.Unnotify(id)
}

// type Middleware struct {
// 	radix *radix.Radix
// }

// func NewMiddleware() *Middleware {
// 	return &Middleware{radix: radix.NewRadix()}
// }

// func (m *Middleware) Build(ctx *view.Context, next *view.Model) {
// 	responder, ok := next.Values[Key].(*Responder)
// 	path := []int64{}
// 	for _, i := range ctx.Path() {
// 		path = append(path, int64(i))
// 	}

// 	if ok {
// 		n := m.radix.Insert(path)
// 		n.Value = responder
// 	} else {
// 		m.radix.Delete(path)
// 	}
// }
