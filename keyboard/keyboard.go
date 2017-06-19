package keyboard

import (
	"github.com/overcyn/matcha"
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

type node struct {
	children  map[matcha.Id]*node
	responder *Responder
}

type Middleware struct {
	root *node
}

func (m *Middleware) Build(ctx *view.Context, next *view.Model) {
	resp, ok := next.Values[Key].(*Responder)
	path := ctx.Path()
	n := m.root.at(path)
	if n == nil {
		n = m.root.add(path)
	}
	n.responder = resp

	if !ok {
		for i := 0; i < len(path); i++ {
			p := path[0 : len(path)-i]
			n := m.root.at(p)
			if n != nil && n.responder == nil {
				m.root.delete(p)
			}
		}
	}
}

func (n *node) add(path []matcha.Id) *node {
	child, ok := n.children[path[0]]
	if !ok {
		child = &node{}
		n.children[path[0]] = child
	}
	if len(path) > 1 {
		return child.add(path[1:])
	} else {
		return child
	}
}

func (n *node) at(path []matcha.Id) *node {
	if len(path) == 0 {
		return n
	}
	child, ok := n.children[path[0]]
	if !ok {
		return nil
	}
	return child.at(path[1:])
}

func (n *node) delete(path []matcha.Id) {
	if len(path) == 1 {
		delete(n.children, path[0])
	} else {
		child, ok := n.children[path[0]]
		if !ok {
			return
		}
		child.delete(path[1:])
	}
}
