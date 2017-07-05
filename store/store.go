package store

import (
	"sync"

	"gomatcha.io/matcha/comm"
)

type Storer interface {
	sync.Locker
	comm.Notifier
	Stores() map[string]Storer
	StoreNotifier(string) comm.Notifier
}

type storeNotifier struct {
	store *Store
	key   string
}

func (s *storeNotifier) Notify(f func()) comm.Id {
	s.store.funcsMu.Lock()
	defer s.store.funcsMu.Unlock()

	if s.store.keyFuncs == nil {
		s.store.keyFuncs = map[string]map[comm.Id]func(){}
	}

	funcs := s.store.keyFuncs[s.key]
	if funcs == nil {
		funcs = map[comm.Id]func(){}
		s.store.keyFuncs[s.key] = funcs
	}

	s.store.maxId += 1
	funcs[s.store.maxId] = f
	return s.store.maxId
}

func (s *storeNotifier) Unnotify(id comm.Id) {
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

type Store struct {
	mu     sync.Mutex
	locked bool
	stores map[string]Storer

	updated     bool
	updatedKeys []string

	funcsMu  sync.Mutex
	funcs    map[comm.Id]func()
	keyFuncs map[string]map[comm.Id]func()
	maxId    comm.Id
}

func (s *Store) Set(key string, chl Storer) {
	if s.stores == nil {
		s.stores = map[string]Storer{}
	}
	s.stores[key] = chl
	s.updatedKeys = append(s.updatedKeys, key)
}

func (s *Store) Delete(key string) {
	delete(s.stores, key)
	s.updatedKeys = append(s.updatedKeys, key)
}

func (s *Store) StoreNotifier(key string) comm.Notifier {
	return &storeNotifier{store: s, key: key}
}

func (s *Store) Stores() map[string]Storer {
	return s.stores
}

func (s *Store) Lock() {
	s.mu.Lock()
	s.locked = true
	for _, i := range s.stores {
		i.Lock()
	}
	s.updated = false
	s.updatedKeys = nil
}

func (s *Store) Unlock() {
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

func (s *Store) Notify(f func()) comm.Id {
	s.funcsMu.Lock()
	defer s.funcsMu.Unlock()

	if s.funcs == nil {
		s.funcs = map[comm.Id]func(){}
	}

	s.maxId += 1
	s.funcs[s.maxId] = f
	return s.maxId
}

func (s *Store) Unnotify(id comm.Id) {
	s.funcsMu.Lock()
	defer s.funcsMu.Unlock()

	if s.funcs == nil {
		panic("store.Unnotify(): Unknown id")
	}

	delete(s.funcs, id)
}

func (s *Store) Update() {
	s.updated = true
}
