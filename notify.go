package mochi

type Notifier interface {
	Notify() chan struct{}
	Unnotify(chan struct{})
}

// type BatchNotifier struct {
// 	notifiers []notifier
// }

// func (n *BatchNotifier) Notify() chan struct{} {
// }
