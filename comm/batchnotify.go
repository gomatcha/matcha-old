package comm

import (
	"sync"

	"github.com/overcyn/mochi"
)

type batchSubscription struct {
	notifier mochi.Notifier
	done     chan struct{}
	c        chan struct{}
}

type BatchNotifier struct {
	mu            sync.Mutex
	subscriptions []*batchSubscription
	chans         []chan struct{}
}

func (bn *BatchNotifier) Subscribe(n mochi.Notifier) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	bn.subscriptions = append(bn.subscriptions, &batchSubscription{notifier: n})
	bn.resubscribeLocked()
}

func (bn *BatchNotifier) Unsubscribe(n mochi.Notifier) {
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
						bn.mu.Lock()
						for _, i := range bn.chans {
							i <- struct{}{}
							<-i
						}
						c <- struct{}{}
						bn.mu.Unlock()
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
