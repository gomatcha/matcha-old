package store

// import (
// 	"errors"
// 	"sync"

// 	"github.com/overcyn/mochi"
// )

// type Storer interface {
// 	sync.Locker
// 	mochi.Notifier
// 	Stores() map[string]Storer
// 	StoreNotifier() mochi.Notifier
// }

// type Store2 struct {
// 	mu     sync.Mutex
// 	locked bool
// 	stores map[string]Storer

// 	updated     bool
// 	updatedKeys []string

// 	chansMu sync.Mutex
// 	chans   []chan struct{}
// }

// func (s *Store2) Add(key string) (*Store2, error) {
// 	_, ok := s.stores[key]
// 	if ok {
// 		return nil, errors.New("Store.Add(): Key already exists")
// 	}
// 	chl := &Store2{}
// 	s.Set(key, chl)
// 	return chl
// }

// func (s *Store2) Set(key string, chl Storer) {
// 	s.stores[key] = chl
// 	s.updatedKeys = append(s.updatedKeys, key)
// }

// func (s *Store2) Delete(key string) {
// 	delete(s.stores, key)
// 	s.updatedKeys = append(s.updatedKeys, key)
// }

// func (s *Store2) StoreNotifier() mochi.Notifier {
// 	return nil
// }

// func (s *Store2) Stores() map[string]Storer {
// 	return s.stores
// }

// func (s *Store2) Lock() {
// 	s.mu.Lock()
// 	s.locked = true
// 	for _, i := range s.stores {
// 		i.Lock()
// 	}
// }

// func (s *Store2) Unlock() {
// 	s.locked = false
// 	for _, i := range s.stores {
// 		i.Unlock()
// 	}
// 	s.mu.Unlock()
// }

// func (s *Store2) Notify() chan struct{} {

// }

// func (s *Store2) Unnotify(c chan struct{}) {

// }

// func (s *Store2) Update() {}
