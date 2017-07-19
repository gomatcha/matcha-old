package comm

import (
	"sync"
)

type Group struct {
	mu    sync.Mutex
	subs  map[Notifier]Id
	funcs map[Id]func()
	maxId Id
}

func (bn *Group) Subscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	// Multiple subscriptions on the same object are ignored.
	if _, ok := bn.subs[n]; ok {
		return
	}

	id := n.Notify(func() {
		bn.mu.Lock()
		defer bn.mu.Unlock()

		for _, f := range bn.funcs {
			f()
		}
	})
	if bn.subs == nil {
		bn.subs = map[Notifier]Id{}
	}
	bn.subs[n] = id
}

func (bn *Group) Unsubscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	id, ok := bn.subs[n]
	if !ok {
		return
	}
	n.Unnotify(id)
	delete(bn.subs, n)
}

func (bn *Group) Notify(f func()) Id {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	bn.maxId += 1
	bn.funcs[bn.maxId] = f
	return bn.maxId
}

func (bn *Group) Unnotify(id Id) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	if _, ok := bn.funcs[id]; !ok {
		panic("comm.Unnotify(): on unknown id")
	}
	delete(bn.funcs, id)
}

func (bn *Group) Signal() {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	for _, f := range bn.funcs {
		f()
	}
}
