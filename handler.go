package mochi

type Handler interface {
    Handle(event interface{})
}
