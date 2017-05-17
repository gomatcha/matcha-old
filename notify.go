package mochi

import (
	"image/color"
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

type batchSubscription struct {
	done chan struct{}
	c    chan struct{}
}

type BatchNotifier struct {
	notifiers map[Notifier]*batchSubscription
	chans     []chan struct{}
}

func NewBatchNotifier(n ...Notifier) *BatchNotifier {
	notifiers := map[Notifier]*batchSubscription{}
	for _, i := range n {
		notifiers[i] = &batchSubscription{}
	}

	return &BatchNotifier{
		notifiers: notifiers,
	}
}

func (n *BatchNotifier) Notify() chan struct{} {
	c := make(chan struct{})
	n.chans = append(n.chans, c)

	if len(n.chans) == 1 {
		n.resubscribe()
	}
	return c
}

func (n *BatchNotifier) Unnotify(c chan struct{}) {
	if c == nil {
		return
	}

	chans := []chan struct{}{}
	for _, i := range n.chans {
		if i != c {
			chans = append(chans, c)
		}
	}
	if len(chans) != len(n.chans)-1 {
		panic("Cant unnotify unknown chan")
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
						for _, i := range n.chans {
							i <- struct{}{}
							<-i
						}
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
