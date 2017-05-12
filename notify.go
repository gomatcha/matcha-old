package mochi

import (
	"sync"
)

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
}

type InterfaceNotifier interface {
	Notifier
	Value() interface{}
}

type BoolNotifier interface {
	Notifier
	Value() bool
}

type IntNotifier interface {
	Notifier
	Value() int
}

type UintNotifier interface {
	Notifier
	Value() uint
}

type Int64Notifier interface {
	Notifier
	Value() int64
}

type Uint64Notifier interface {
	Notifier
	Value() uint64
}

type Float64Notifier interface {
	Notifier
	Value() float64
}

type StringNotifier interface {
	Notifier
	Value() string
}

type BytesNotifier interface {
	Notifier
	Value() []byte
}

type batchSubscription struct {
	done chan struct{}
	c    chan struct{}
}

type BatchNotifier struct {
	mu        sync.Mutex
	notifiers map[Notifier]*batchSubscription
	chans     []chan struct{}
}

func NewBatchNotifier() *BatchNotifier {
	return &BatchNotifier{
		notifiers: map[Notifier]*batchSubscription{},
	}
}

func (n *BatchNotifier) Add(ns Notifier) {
	n.mu.Lock()
	defer n.mu.Unlock()

	_, ok := n.notifiers[ns]
	if ok {
		return
	}
	n.notifiers[ns] = &batchSubscription{}
	n.resubscribe()
}

func (n *BatchNotifier) Remove(ns Notifier) {
	n.mu.Lock()
	defer n.mu.Unlock()

	sub, ok := n.notifiers[ns]
	if !ok {
		return
	}

	if sub.c != nil {
		ns.Unnotify(sub.c)
		close(sub.done)
	}
	delete(n.notifiers, ns)
}

func (n *BatchNotifier) Notify() chan struct{} {
	n.mu.Lock()
	defer n.mu.Unlock()

	c := make(chan struct{})
	n.chans = append(n.chans, c)

	if len(n.chans) == 1 {
		n.resubscribe()
	}
	return c
}

func (n *BatchNotifier) Unnotify(c chan struct{}) {
	n.mu.Lock()
	defer n.mu.Unlock()

	chans := []chan struct{}{}
	for _, i := range n.chans {
		if i != c {
			chans = append(chans, c)
		}
	}
	n.chans = chans

	if len(n.chans) == 0 {
		n.resubscribe()
	}
}

func (n *BatchNotifier) resubscribe() {
	// If we have no chans, remove all subscribers.
	if len(n.chans) == 0 {
		for notifier, sub := range n.notifiers {
			if sub.c != nil {
				notifier.Unnotify(sub.c)
				close(sub.done)
				sub.c = nil
				sub.done = nil
			}
		}
		return
	}

	for notifier, sub := range n.notifiers {
		if sub.c == nil {
			c := notifier.Notify()
			done := make(chan struct{})

			go func() {
			loop:
				for {
					select {
					case <-c:
						n.mu.Lock()
						for _, i := range n.chans {
							i <- struct{}{}
							<-i
						}
						n.mu.Unlock()
						c <- struct{}{}
					case <-done:
						break loop
					}
				}
			}()

			sub.c = c
			sub.done = done
		}
	}
}
