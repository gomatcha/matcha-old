package comm

import (
	"fmt"
	"sync"
)

type BatchNotifier struct {
	mu    sync.Mutex
	subs  map[Notifier]Id
	funcs map[Id]func()
	maxId Id
}

func (bn *BatchNotifier) Subscribe(n Notifier) {
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

func (bn *BatchNotifier) Unsubscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	id, ok := bn.subs[n]
	if !ok {
		return
	}
	n.Unnotify(id)
	delete(bn.subs, n)
}

func (bn *BatchNotifier) Notify(f func()) Id {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	bn.maxId += 1
	bn.funcs[bn.maxId] = f
	return bn.maxId
}

func (bn *BatchNotifier) Unnotify(id Id) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	delete(bn.funcs, id)
}

func (bn *BatchNotifier) Update() {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	fmt.Println("Update", bn.funcs)
	for _, f := range bn.funcs {
		f()
	}
}
