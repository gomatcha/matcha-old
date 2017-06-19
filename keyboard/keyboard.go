package keyboard

import (
	"github.com/overcyn/matcha/internal/radix"
	"github.com/overcyn/matcha/view"
)

type key struct{}
type textKey struct{}

var Key = key{}
var HelperKey = textKey{}

type Responder struct {
}

// func (g *Responder) Next() {
// }

// func (g *Responder) Prev() {
// }

func (g *Responder) Show() {

}

func (g *Responder) Dismiss() {
}

func (g *Responder) Visible() bool {
	return true
}

// func (g *Responder) Notifier() *comm.BoolNotifier {
// }

type Middleware struct {
	radix *radix.Radix
}

func NewMiddleware() *Middleware {
	return &Middleware{radix: radix.NewRadix()}
}

func (m *Middleware) Build(ctx *view.Context, next *view.Model) {
	responder, ok := next.Values[Key].(*Responder)
	path := []int64{}
	for _, i := range ctx.Path() {
		path = append(path, int64(i))
	}

	if ok {
		n := m.radix.Insert(path)
		n.Value = responder
	} else {
		m.radix.Delete(path)
	}
}
