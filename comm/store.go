package comm

import (
	"sync"
)

type Storer interface {
	sync.Locker
	Notifier
	Stores() map[string]Storer
	StoreNotifier(string) Notifier
}

type asyncStoreNotifier struct {
	store *AsyncStore
	key   string
}

func (s *asyncStoreNotifier) Notify(f func()) Id {
	s.store.funcsMu.Lock()
	defer s.store.funcsMu.Unlock()

	if s.store.keyFuncs == nil {
		s.store.keyFuncs = map[string]map[Id]func(){}
	}

	funcs := s.store.keyFuncs[s.key]
	if funcs == nil {
		funcs = map[Id]func(){}
		s.store.keyFuncs[s.key] = funcs
	}

	s.store.maxId += 1
	funcs[s.store.maxId] = f
	return s.store.maxId
}

func (s *asyncStoreNotifier) Unnotify(id Id) {
	s.store.funcsMu.Lock()
	defer s.store.funcsMu.Unlock()

	if s.store.keyFuncs == nil {
		panic("Unknown id")
	}

	funcs := s.store.keyFuncs[s.key]
	if funcs == nil {
		panic("Unknown id")
	}

	delete(funcs, id)
}

type AsyncStore struct {
	mu     sync.Mutex
	locked bool
	stores map[string]Storer

	updated     bool
	updatedKeys []string

	funcsMu  sync.Mutex
	funcs    map[Id]func()
	keyFuncs map[string]map[Id]func()
	maxId    Id
}

func (s *AsyncStore) Set(key string, chl Storer) {
	if s.stores == nil {
		s.stores = map[string]Storer{}
	}
	s.stores[key] = chl
	s.updatedKeys = append(s.updatedKeys, key)
}

func (s *AsyncStore) Delete(key string) {
	delete(s.stores, key)
	s.updatedKeys = append(s.updatedKeys, key)
}

func (s *AsyncStore) StoreNotifier(key string) Notifier {
	return &asyncStoreNotifier{store: s, key: key}
}

func (s *AsyncStore) Stores() map[string]Storer {
	return s.stores
}

func (s *AsyncStore) Lock() {
	s.mu.Lock()
	s.locked = true
	for _, i := range s.stores {
		i.Lock()
	}
	s.updated = false
	s.updatedKeys = nil
}

func (s *AsyncStore) Unlock() {
	go func(updated bool, updatedKeys []string) {
		s.funcsMu.Lock()
		defer s.funcsMu.Unlock()

		if updated {
			for _, i := range s.funcs {
				i()
			}
		}
		for _, i := range updatedKeys {
			for _, j := range s.keyFuncs[i] {
				j()
			}
		}
	}(s.updated, s.updatedKeys)
	s.updated = false
	s.updatedKeys = nil

	s.locked = false
	for _, i := range s.stores {
		i.Unlock()
	}
	s.mu.Unlock()
}

func (s *AsyncStore) Notify(f func()) Id {
	s.funcsMu.Lock()
	defer s.funcsMu.Unlock()

	if s.funcs == nil {
		s.funcs = map[Id]func(){}
	}

	s.maxId += 1
	s.funcs[s.maxId] = f
	return s.maxId
}

func (s *AsyncStore) Unnotify(id Id) {
	s.funcsMu.Lock()
	defer s.funcsMu.Unlock()

	if s.funcs == nil {
		panic("store.Unnotify(): Unknown id")
	}

	delete(s.funcs, id)
}

func (s *AsyncStore) Update() {
	s.updated = true
}
