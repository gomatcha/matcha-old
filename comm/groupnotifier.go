package comm

import (
	"sync"
)

type GroupNotifier struct {
	mu    sync.Mutex
	subs  map[Notifier]Id
	funcs map[Id]func()
	maxId Id
}

func (bn *GroupNotifier) Subscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

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

func (bn *GroupNotifier) Unsubscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	id, ok := bn.subs[n]
	if !ok {
		return
	}
	n.Unnotify(id)
	delete(bn.subs, n)
}

func (bn *GroupNotifier) Notify(f func()) Id {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	bn.maxId += 1
	bn.funcs[bn.maxId] = f
	return bn.maxId
}

func (bn *GroupNotifier) Unnotify(id Id) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	delete(bn.funcs, id)
}

func (bn *GroupNotifier) Signal() {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	for _, f := range bn.funcs {
		f()
	}
}
