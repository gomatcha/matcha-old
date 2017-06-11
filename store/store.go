package store

import (
	"sync"

	"github.com/overcyn/mochi"
)

type storeNotifier struct {
	store *Store
	paths [][]int64
}

func (s *storeNotifier) Notify() chan struct{} {
	return s.store.notifyPaths(s.paths)
}

func (s *storeNotifier) Unnotify(c chan struct{}) {
	s.store.unnotifyPaths(c)
}

type Store struct {
	mu        sync.RWMutex
	locked    bool
	rlocked   int
	parent    *Store
	parentKey int64
	children  map[int64]*Store

	updated           bool
	updatedPaths      [][]int64
	updatedStorePaths [][]int64

	chansMu   sync.Mutex
	chans     []chan struct{}
	pathChans map[chan struct{}][][]int64
}

func (s *Store) root() *Store {
	root := s
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (s *Store) Add(key int64) (*Store, error) {
	// TODO(KD): check if the key already exists.
	chl := &Store{}
	s.Set(key, chl)
	return chl, nil
}

func (s *Store) Set(key int64, chl *Store) {
	s.updateStore([]int64{key})

	chl.lock()
	if chl.parent != nil {
		panic("Store2.Set() child already has parent")
	}
	chl.parent = s
	chl.parentKey = key
	if s.children == nil {
		s.children = map[int64]*Store{}
	}
	s.children[key] = chl
}

func (s *Store) Delete(key int64) {
	s.updateStore([]int64{key})

	chl, ok := s.children[key]
	if !ok {
		return
	}
	if chl.parent != s {
		panic("Store2.RemoveChild() child is not parent")
	}
	delete(s.children, chl.parentKey)
	chl.parent = nil
	chl.parentKey = 0
	chl.unlock()
}

func (s *Store) isLocked() bool {
	return s.locked
}

func (s *Store) isRLocked() bool {
	return s.rlocked > 0
}

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

	// notify any direct listeners for all writes
	s.chansMu.Lock()
	defer s.chansMu.Unlock()
	if s.updated {
		for _, c := range s.chans {
			c <- struct{}{}
			<-c
		}
	}
	s.updated = false

	// notify paths of all writes.
	for c, paths := range s.pathChans {
		if matchPaths(s.updatedPaths, paths) {
			c <- struct{}{}
			<-c
		}
	}
	s.updatedPaths = nil

	for c, paths := range s.pathChans {
		if matchPaths2(s.updatedPaths, paths) {
			c <- struct{}{}
			<-c
		}
	}
	s.updatedStorePaths = nil
}

func (s *Store) RLock() {
	s.rlock()
}

func (s *Store) rlock() {
	if s.parent != nil {
		s.parent.rlock()
	} else {
		s.mu.RLock()
		s.rlocked += 1
	}
}

func (s *Store) RUnlock() {
	s.runlock()
}

func (s *Store) runlock() {
	if s.parent != nil {
		s.parent.rlock()
	} else {
		s.rlocked -= 1
		s.mu.RUnlock()
	}
}

func (s *Store) Update() {
	if !s.isLocked() {
		panic("Update called on unlocked store")
	}

	// Skip write if already happened
	if s.updated {
		return
	}

	s.updated = true
	if s.parent != nil {
		path := []int64{s.parentKey}
		s.parent.update(path)
	}
}

func (s *Store) update(path []int64) {
	if !s.isLocked() {
		panic("Update called on unlocked store")
	}

	s.updated = true
	s.updatedPaths = append(s.updatedPaths, path)
	if s.parent != nil {
		newPath := append([]int64(nil), path...)
		s.parent.update(newPath)
	}
}

func (s *Store) updateStore(path []int64) {
	if !s.isLocked() {
		panic("Update called on unlocked store")
	}

	s.updated = true
	s.updatedPaths = append(s.updatedPaths, path)
	s.updatedStorePaths = append(s.updatedStorePaths, path)
	if s.parent != nil {
		newPath := append([]int64(nil), path...)
		s.parent.updateStore(newPath)
	}
}

func (s *Store) Notifier(childStores ...*Store) mochi.Notifier {
	paths := [][]int64{}
	for _, i := range childStores {
		path := s.path(i, nil)
		if path == nil {
			panic("Store.Notifier(): Passed store is not a child of the store")
		}
		paths = append(paths, path)
	}

	return &storeNotifier{
		store: s,
		paths: paths,
	}
}

// Assumes locked
func (s *Store) path(child *Store, curr []int64) []int64 {
	if s == child {
		return curr
	}

	for k, v := range s.children {
		rlt := v.path(child, append(curr, k))
		if rlt != nil {
			return rlt
		}
	}
	return nil
}

func (s *Store) Notify() chan struct{} {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	c := make(chan struct{})
	s.chans = append(s.chans, c)
	return c
}

func (s *Store) Unnotify(c chan struct{}) {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	copy := []chan struct{}{}
	for _, i := range s.chans {
		if i != c {
			copy = append(copy, i)
		}
	}
	s.chans = copy
}

func (s *Store) notifyPaths(paths [][]int64) chan struct{} {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	c := make(chan struct{})
	if s.pathChans == nil {
		s.pathChans = map[chan struct{}][][]int64{}
	}
	s.pathChans[c] = paths
	return c
}

func (s *Store) unnotifyPaths(c chan struct{}) {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	delete(s.pathChans, c)
}

func matchPaths(updatedPaths [][]int64, monitoredPaths [][]int64) bool {
	for _, i := range updatedPaths {
		for _, j := range monitoredPaths {
			if matchPath(i, j) {
				return true
			}
		}
	}
	return false
}

func matchPath(updatedPath []int64, monitoredPath []int64) bool {
	if len(monitoredPath) > len(updatedPath) {
		return false
	}
	for i := range monitoredPath {
		if monitoredPath[i] != updatedPath[i] {
			return false
		}
	}
	return true
}

func matchPaths2(updatedPaths [][]int64, monitoredPaths [][]int64) bool {
	for _, i := range updatedPaths {
		for _, j := range monitoredPaths {
			if matchPath(i, j) {
				return true
			}
		}
	}
	return false
}

func matchPath2(updatedPath []int64, monitoredPath []int64) bool {
	if len(monitoredPath) > len(updatedPath) {
		for i := range updatedPath {
			if monitoredPath[i] != updatedPath[i] {
				return false
			}
		}
	} else {
		for i := range monitoredPath {
			if monitoredPath[i] != updatedPath[i] {
				return false
			}
		}
	}
	return true
}
