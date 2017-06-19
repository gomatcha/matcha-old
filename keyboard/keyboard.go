package keyboard

type key struct{}

var Key = key{}

type Responder struct {
}

func (g *Responder) Next() {
}

func (g *Responder) Prev() {
}

func (g *Responder) Show() {
}

func (g *Responder) Dismiss() {
}

func (g *Responder) Visible() bool {
}

type Middleware struct {
}
