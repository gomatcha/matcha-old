package comm

import "sync"

// func NotifyFunc(n Notifier, f func()) (done chan struct{}) {
// 	c := n.Notify()
// 	done = make(chan struct{})
// 	go func() {
// 	loop:
// 		for {
// 			select {
// 			case <-c:
// 				f()
// 				c <- struct{}{}
// 			case <-done:
// 				break loop
// 			}
// 		}
// 	}()
// 	return done
// }

type BatchNotifier2 struct {
	mu    sync.Mutex
	subs  map[Notifier]Id
	funcs map[Id]func()
	maxId Id
}

func (bn *BatchNotifier2) Subscribe(n Notifier) {
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

func (bn *BatchNotifier2) Unsubscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	id, ok := bn.subs[n]
	if !ok {
		return
	}
	n.Unnotify(id)
	delete(bn.subs, n)
}

func (bn *BatchNotifier2) Notify(f func()) Id {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	bn.maxId += 1
	bn.funcs[bn.maxId] = f
	return bn.maxId
}

func (bn *BatchNotifier2) Unnotify(id Id) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	if bn.funcs == nil {
		bn.funcs = map[Id]func(){}
	}
	delete(bn.funcs, id)
}

func (bn *BatchNotifier2) Update() {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	for _, f := range bn.funcs {
		f()
	}
}
