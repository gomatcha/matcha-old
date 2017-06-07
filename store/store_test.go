package store

import (
	"testing"
)

func TestMatchPath(t *testing.T) {
	if !matchPath([]int64{}, []int64{}) {
		t.Fail()
	}
	if !matchPath([]int64{0}, []int64{0}) {
		t.Fail()
	}
	if matchPath([]int64{1}, []int64{0}) {
		t.Fail()
	}
	if !matchPath([]int64{0, 1, 2, 3}, []int64{0}) {
		t.Fail()
	}
	if matchPath([]int64{0}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if matchPath([]int64{1}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if !matchPath([]int64{0, 1, 2, 3}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
}

func TestMatchPath2(t *testing.T) {
	if !matchPath2([]int64{}, []int64{}) {
		t.Fail()
	}
	if !matchPath2([]int64{0}, []int64{0}) {
		t.Fail()
	}
	if matchPath2([]int64{1}, []int64{0}) {
		t.Fail()
	}
	if !matchPath2([]int64{0, 1, 2, 3}, []int64{0}) {
		t.Fail()
	}
	if !matchPath2([]int64{0}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if matchPath2([]int64{1}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if !matchPath2([]int64{0, 1, 2, 3}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
}

func TestStore(t *testing.T) {
	s := &Store{}
	s.Lock()
	s.Update()
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

// func TestNotify(t *testing.T) {
// 	s := &Store{}
// 	n := s.Notifier(0)
// 	go func() {
// 		s.Lock()
// 		s.WriteKey(0)
// 		s.Unlock()
// 	}()

// 	c := n.Notify()
// 	<-c
// 	c <- struct{}{}
// 	n.Unnotify(c)
// }

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
