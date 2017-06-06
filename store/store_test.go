package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	s := &Store{}
	s.Lock()
	s.ReadKey(0)
	s.Unlock()
}

// func TestMultiRead(t *testing.T) {
// 	s := &Store{}
// 	s.RLock()
// 	s.RLock()
// 	s.RLock()

// 	s.ReadKey(0)
// 	s.ReadKey(0)
// 	s.ReadKey(0)

// 	s.RUnlock()
// 	s.RUnlock()
// 	s.RUnlock()
// }

func TestNotify(t *testing.T) {
	s := &Store{}
	n := s.Notifier(0)
	go func() {
		s.Lock()
		s.WriteKey(0)
		s.Unlock()
	}()

	c := n.Notify()
	<-c
	c <- struct{}{}
	n.Unnotify(c)
}

// func TestReadWrite(t *testing.T) {
// 	s := &Store{}

// 	go func() {
// 		s.RLock()
// 		s.ReadKey(0)
// 		s.RUnlock()
// 	}()

// 	s.Lock()
// 	s.ReadKey(0)
// 	s.Unlock()
// }
