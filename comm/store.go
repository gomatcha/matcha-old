package comm

import (
	"sync"

	"github.com/overcyn/mochi"
)

type Storer interface {
	sync.Locker
	mochi.Notifier
	Stores() map[string]Storer
	StoreNotifier(string) mochi.Notifier
}

type asyncStoreNotifier struct {
	store *AsyncStore
	key   string
}

func (s *asyncStoreNotifier) Notify() chan struct{} {
	s.store.chansMu.Lock()
	defer s.store.chansMu.Unlock()

	c := make(chan struct{})
	if s.store.keyChans == nil {
		s.store.keyChans = map[string][]chan struct{}{}
	}
	chans := s.store.keyChans[s.key]
	chans = append(chans, c)
	s.store.keyChans[s.key] = chans
	return c
}

func (s *asyncStoreNotifier) Unnotify(c chan struct{}) {
	s.store.chansMu.Lock()
	defer s.store.chansMu.Unlock()

	chans := []chan struct{}{}
	for _, i := range s.store.chans {
		if i != c {
			chans = append(chans, i)
		}
	}
	s.store.chans = chans
}

type AsyncStore struct {
	mu     sync.Mutex
	locked bool
	stores map[string]Storer

	updated     bool
	updatedKeys []string

	chansMu  sync.Mutex
	chans    []chan struct{}
	keyChans map[string][]chan struct{}
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

func (s *AsyncStore) StoreNotifier(key string) mochi.Notifier {
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
		s.chansMu.Lock()
		defer s.chansMu.Unlock()

		if updated {
			for _, i := range s.chans {
				i <- struct{}{}
				<-i
			}
		}
		for _, i := range updatedKeys {
			for _, j := range s.keyChans[i] {
				j <- struct{}{}
				<-j
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

func (s *AsyncStore) Notify() chan struct{} {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	c := make(chan struct{})
	s.chans = append(s.chans, c)
	return c
}

func (s *AsyncStore) Unnotify(c chan struct{}) {
	s.chansMu.Lock()
	defer s.chansMu.Unlock()

	chans := []chan struct{}{}
	for _, i := range s.chans {
		if i != c {
			chans = append(chans, i)
		}
	}
	s.chans = chans
}

func (s *AsyncStore) Update() {
	s.updated = true
}
