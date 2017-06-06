package mochi

import (
	"fmt"
	"image/color"
	"sync"
)

type Id int64

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
}

type ColorNotifier interface {
	Notifier
	Value() color.Color
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

func NotifyFunc(n Notifier, f func()) (done chan struct{}) {
	c := n.Notify()
	done = make(chan struct{})
	go func() {
	loop:
		for {
			select {
			case <-c:
				f()
				c <- struct{}{}
			case <-done:
				break loop
			}
		}
	}()
	return done
}

type batchSubscription struct {
	notifier Notifier
	done     chan struct{}
	c        chan struct{}
}

type BatchNotifier struct {
	mu            sync.Mutex
	subscriptions []*batchSubscription
	chans         []chan struct{}
}

func (bn *BatchNotifier) Subscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	bn.subscriptions = append(bn.subscriptions, &batchSubscription{notifier: n})
	bn.resubscribeLocked()
}

func (bn *BatchNotifier) Unsubscribe(n Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	found := false
	subs := []*batchSubscription{}
	for _, sub := range bn.subscriptions {
		if found || sub.notifier != n {
			subs = append(subs, sub)
		} else {
			if sub.c != nil {
				n.Unnotify(sub.c)
				close(sub.done)
			}
			found = true
		}
	}
	if !found {
		panic("Cant unobserve unknown notifier")
	}
	bn.subscriptions = subs

	bn.resubscribeLocked()
}

func (bn *BatchNotifier) Notify() chan struct{} {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	c := make(chan struct{})
	bn.chans = append(bn.chans, c)

	bn.resubscribeLocked()
	return c
}

func (bn *BatchNotifier) Unnotify(c chan struct{}) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	chans := []chan struct{}{}
	for _, i := range bn.chans {
		if i != c {
			chans = append(chans, c)
		}
	}
	if len(chans) != len(bn.chans)-1 {
		panic("Cant unnotify unknown chan")
	}
	bn.chans = chans

	bn.resubscribeLocked()
}

func (bn *BatchNotifier) Signal() {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	for _, i := range bn.chans {
		i <- struct{}{}
		<-i
	}
}

func (bn *BatchNotifier) resubscribeLocked() {
	// If we have no chans, remove all subscribers.
	if len(bn.chans) == 0 {
		for _, sub := range bn.subscriptions {
			if sub.c != nil {
				sub.notifier.Unnotify(sub.c)
				close(sub.done)
				sub.c = nil
				sub.done = nil
			}
		}
		return
	}

	for _, sub := range bn.subscriptions {
		// Subscribe to any notifier that isn't yet subscribed
		if sub.c == nil {
			c := sub.notifier.Notify()
			done := make(chan struct{})

			go func() {
			loop:
				for {
					select {
					case <-c:
						fmt.Println("send1")
						bn.mu.Lock()
						for _, i := range bn.chans {
							fmt.Println("send2", i)
							i <- struct{}{}
							fmt.Println("send3", i)
							<-i
							fmt.Println("send4", i)
						}
						c <- struct{}{}
						fmt.Println("send5")
						bn.mu.Unlock()
						fmt.Println("send6")
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

// func (bn *BatchNotifier) String() string {
// 	ns := []Notifier{}
// 	for k := range bn.subscriptions {
// 		ns = append(ns, k)
// 	}
// 	return fmt.Sprintf("&BatchNotifier{%p observing:%v chans:%v}", bn, ns, len(bn.chans))
// }
