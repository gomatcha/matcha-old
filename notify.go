package mochi

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
}
