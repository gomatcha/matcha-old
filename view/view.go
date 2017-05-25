package view

import (
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
)

type View interface {
	Build(*Context) *Model
	Lifecycle(from, to Stage)
	Id() mochi.Id
	sync.Locker
	mochi.Notifier
}

type Embed struct {
	mu    sync.Mutex
	id    mochi.Id
	root  *root
	chans []chan struct{}
}

func NewEmbed(id mochi.Id) *Embed {
	return &Embed{id: id}
}

func (e *Embed) Build(ctx *Context) *Model {
	return &Model{}
}

func (e *Embed) Id() mochi.Id {
	return e.id
}

func (e *Embed) Lifecycle(from, to Stage) {
	// no-op
}

func (e *Embed) Lock() {
	e.mu.Lock()
}

func (e *Embed) Unlock() {
	e.mu.Unlock()
}

func (e *Embed) Notify() chan struct{} {
	c := make(chan struct{})
	e.chans = append(e.chans, c)
	return c
}

func (e *Embed) Unnotify(c chan struct{}) {
	chans := []chan struct{}{}
	for _, i := range e.chans {
		if i != c {
			chans = append(chans, i)
		}
	}
	e.chans = chans
}

func (e *Embed) Update(key interface{}) {
	for _, i := range e.chans {
		i <- struct{}{}
		<-i
	}
}

func (e *Embed) String() string {
	return fmt.Sprintf("&Embed{id:%v}", e.id)
}

type Stage int

const (
	StageDead Stage = iota
	StageMounted
	StagePrepreload
	StagePreload
	StageVisible
)

func EntersStage(from, to, s Stage) bool {
	return from < s && to >= s
}

func ExitsStage(from, to, s Stage) bool {
	return from >= s && to < s
}

type Model struct {
	Children            map[mochi.Id]View
	Layouter            layout.Layouter
	Painter             paint.Painter
	Values              map[interface{}]interface{}
	NativeName          string
	NativeStateProtobuf proto.Message
}

func (n *Model) Add(v View) {
	if n.Children == nil {
		n.Children = map[mochi.Id]View{}
	}
	n.Children[v.Id()] = v
}
