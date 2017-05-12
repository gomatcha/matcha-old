package store

import (
	"github.com/overcyn/mochi"
)

func Notifier(f func(tx *Tx)) mochi.Notifier {
	return nil
}

// func InterfaceNotifier(f func(tx *Tx) interface{}) mochi.InterfaceNotifier {
// 	return nil
// }

// func BoolNotifier(f func(tx *Tx) bool) mochi.BoolNotifier {
// 	return nil
// }

// func IntNotifier(f func(tx *Tx) int) mochi.IntNotifier {
// 	return nil
// }

// func UintNotifier(f func(tx *Tx) uint) mochi.UintNotifier {
// 	return nil
// }

// func Int64Notifier(f func(tx *Tx) int64) mochi.Int64Notifier {
// 	return nil
// }

// func Uint64Notifier(f func(tx *Tx) uint64) mochi.Uint64Notifier {
// 	return nil
// }

// func Float64Notifier(f func(tx *Tx) float64) mochi.Float64Notifier {
// 	return nil
// }

// func StringNotifier(f func(tx *Tx) string) mochi.StringNotifier {
//  return nil
// }

// func BytesNotifier(f func(tx *Tx) []byte) mochi.ByteNotifier {
// 	return nil
// }

// type value struct {
//  chans     []chan struct{}
//  mu        *sync.Mutex
//  value     interface{}
//  notifiers []mochi.Notifier
//  done      []chan struct{}
// }

// func (v *value) Notify() chan struct{} {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  c := make(chan struct{})
//  v.chans = append(v.chans, c)
//  return c
// }

// func (v *value) Unnotify(c chan struct{}) {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  chans := make([]chan struct{}, 0, len(v.chans))
//  for _, i := range chans {
//      if i != c {
//          chans = append(chans, i)
//      }
//  }
//  v.chans = chans
// }

// func (v *value) Value() interface{} {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  return v.value
// }

// func (v *value) Set(a interface{}) {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  v.value = a
//  for _, i := range v.chans {
//      i <- struct{}{}
//      <-i
//  }
// }

// func (v *value) Watch(n mochi.Notifier, f func() interface{}) {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  done := make(chan struct{})
//  v.notifiers = append(v.notifiers, n)
//  v.done = append(v.done, done)
//  c := n.Notify()

//  // setup a go routine waiting for notifications from n.
//  go func() {
//  loop:
//      for {
//          select {
//          case <-c:
//              v.Set(f())
//              c <- struct{}{}
//          case <-done:
//              n.Unnotify(c)
//              break loop
//          }
//      }
//  }()
// }

// func (v *value) Unwatch(n mochi.Notifier) {
//  v.mu.Lock()
//  defer v.mu.Unlock()

//  notifiers := []mochi.Notifier{}
//  done := []chan struct{}{}
//  for idx, i := range v.notifiers {
//      if i == n {
//          v.done[idx] <- struct{}{}
//      } else {
//          notifiers = append(notifiers, i)
//          done = append(done, v.done[idx])
//      }
//  }
// }
