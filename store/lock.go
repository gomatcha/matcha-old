package store

import (
	"sync"

	"github.com/overcyn/mochi"
)

type storeNotifier2 struct {
	store *Store
	key   interface{}
}

func (s *storeNotifier2) Notify() chan struct{} {
	return s.store.NotifyKey(s.key)
}

func (s *storeNotifier2) Unnotify(c chan struct{}) {
	s.store.UnnotifyKey(s.key, c)
}

type Store struct {
	mu     sync.RWMutex
	locked bool
	// rlocked   int
	parent    *Store
	parentKey interface{}
	children  map[interface{}]*Store
	// reads     []interface{}
	writes    []interface{}
	writeRoot bool

	chansMu sync.Mutex
	chans   map[interface{}][]chan struct{}
}

func (s *Store) root() *Store {
	root := s
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (s *Store) AddChild(chl *Store, key interface{}) {
	s.WriteKey(key)
	chl.lock()
	if chl.parent != nil {
		panic("Store.AddChild() child already has parent")
	}
	chl.parent = s
	chl.parentKey = key

	if s.children == nil {
		s.children = map[interface{}]*Store{}
	}
	s.children[key] = chl
}

func (s *Store) RemoveChild(chl *Store) {
	s.WriteKey(chl.parentKey)
	if chl.parent != s {
		panic("Store.RemoveChild() child is not parent")
	}
	delete(s.children, chl.parentKey)
	chl.unlock()
	chl.parent = nil
	chl.parentKey = nil
}

func (s *Store) isLocked() bool {
	return s.locked
}

// func (s *Store) isRLocked() bool {
// 	if s.parent != nil {
// 		return s.parent.isLocked()
// 	}
// 	return s.rlocked > 0
// }

func (s *Store) Lock() {
	s.root().lock()
}

func (s *Store) lock() {
	s.mu.Lock()
	s.locked = true
	for _, i := range s.children {
		i.lock()
	}
}

func (s *Store) Unlock() {
	s.root().unlock()
}

func (s *Store) unlock() {
	s.mu.Unlock()
	s.locked = false
	for _, i := range s.children {
		i.unlock()
	}

	// notify for all writes
	s.chansMu.Lock()
	defer s.chansMu.Unlock()
	if s.writeRoot {
		s.writes = append(s.writes, rootKeyVar)
	}
	for _, i := range s.writes {
		cs, ok := s.chans[i]
		for _, c := range cs {
			if ok {
				c <- struct{}{}
				<-c
			}
		}
	}
	s.writes = nil
	s.writeRoot = false
}

// func (s *Store) RLock() {
// 	s.rlock()
// }

// func (s *Store) rlock() {
// 	if s.parent != nil {
// 		s.parent.rlock()
// 	} else {
// 		s.mu.RLock()
// 		s.rlocked += 1
// 	}
// }

// func (s *Store) RUnlock() {
// 	s.runlock()
// }

// func (s *Store) runlock() {
// 	if s.parent != nil {
// 		s.parent.rlock()
// 	} else {
// 		s.rlocked -= 1
// 		s.mu.RUnlock()
// 	}
// }

func (s *Store) Write() {
	if !s.isLocked() {
		panic("WriteKey called on unlocked store")
	}

	// Skip write if already happend
	if s.writeRoot {
		return
	}
	s.writeRoot = true

	if s.parent != nil {
		s.parent.WriteKey(s.parentKey)
	}
}

func (s *Store) Read() {
	if !s.isLocked() {
		panic("WriteKey called on unlocked store")
	}
}

func (s *Store) WriteKey(key interface{}) {
	if !s.isLocked() {
		panic("WriteKey called on unlocked store")
	}

	// Skip write if already happend
	for _, i := range s.writes {
		if i == key {
			return
		}
	}
	s.writes = append(s.writes, key)

	if s.parent != nil {
		s.parent.WriteKey(s.parentKey)
	}
}

func (s *Store) ReadKey(key interface{}) {
	if !s.isLocked() {
		panic("WriteKey called on unlocked store")
	}
	// s.reads = append(s.reads, key)
}

func (s *Store) Notifier(key interface{}) mochi.Notifier {
	return &storeNotifier2{
		store: s,
		key:   key,
	}
}

func (s *Store) Notify() chan struct{} {
	return s.NotifyKey(rootKeyVar)
}

func (s *Store) Unnotify(c chan struct{}) {
	s.UnnotifyKey(rootKeyVar, c)
}

func (s *Store) NotifyKey(k interface{}) chan struct{} {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	c := make(chan struct{})
	if s.chans == nil {
		s.chans = map[interface{}][]chan struct{}{}
	}
	s.chans[k] = append(s.chans[k], c)
	return c
}

func (s *Store) UnnotifyKey(k interface{}, c chan struct{}) {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	chans := s.chans[k]
	copy := []chan struct{}{}
	for _, i := range chans {
		if i != c {
			copy = append(copy, i)
		}
	}
	s.chans[k] = copy
}
