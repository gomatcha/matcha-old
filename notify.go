package mochi

type Id int64

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
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
