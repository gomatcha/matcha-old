package store

import (
	"sync"

	"github.com/overcyn/mochi"
)

type storeNotifier struct {
	store *Store
	key   interface{}
}

func (s *storeNotifier) Notify() chan struct{} {
	return s.store.NotifyKey(s.key)
}

func (s *storeNotifier) Unnotify(c chan struct{}) {
	s.store.UnnotifyKey(s.key, c)
}

type Store struct {
	writeMu sync.Mutex
	readMu  sync.RWMutex

	txMu    sync.Mutex
	readTx  map[*Tx]struct{}
	writeTx *Tx

	chansMu sync.Mutex
	chans   map[interface{}][]chan struct{}
}

var rootKeyVar rootKey // TODO(KD): any change to the store should affect this key. And any change to the key should affect all watchers of the store.

type rootKey struct{}

func (s *Store) Write(tx *Tx) {
	s.WriteKey(rootKeyVar, tx)
}

func (s *Store) Read(tx *Tx) {
	s.ReadKey(rootKeyVar, tx)
}

func (s *Store) WriteKey(key interface{}, tx *Tx) {
	if tx == nil {
		panic("Store.Write() called outside of a transaction")
	}
	if tx.kind != txKindWrite {
		panic("Store.Write() called in readonly transaction")
	}

	s.lock(tx)
	tx.writes = append(tx.writes, txAccess{store: s, key: key})
}

func (s *Store) ReadKey(key interface{}, tx *Tx) {
	if tx == nil {
		panic("Store.Read() is called outside of a transaction")
	}

	s.lock(tx)
	tx.reads = append(tx.reads, txAccess{store: s, key: key})
}

func (s *Store) Notifier(key interface{}) mochi.Notifier {
	return &storeNotifier{
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

func (s *Store) lock(tx *Tx) {
	// If we have not used this store in the transaction, lock the store and set the Tx.
	if tx.kind == txKindRead {
		s.txMu.Lock()
		_, contains := s.readTx[tx]
		s.txMu.Unlock()

		if !contains {
			s.readMu.RLock()
			tx.stores = append(tx.stores, s)

			s.txMu.Lock()
			if s.readTx == nil {
				s.readTx = map[*Tx]struct{}{}
			}
			s.readTx[tx] = struct{}{}
			s.txMu.Unlock()
		}
	} else {
		s.txMu.Lock()
		equal := s.writeTx == tx
		s.txMu.Unlock()

		if !equal {
			s.writeMu.Lock()
			s.readMu.Lock()
			tx.stores = append(tx.stores, s)

			s.txMu.Lock()
			s.writeTx = tx
			s.txMu.Unlock()
		}
	}
}

func (s *Store) writeCommit1(tx *Tx) {
	if tx == nil {
		panic("Store.writeCommit1() called outside of a transaction")
	}
	if tx.kind != txKindWrite {
		panic("Store.writeCommit1() called in a readonly transaction")
	}

	s.txMu.Lock()
	if s.writeTx != tx {
		panic("Store.commit() called with an unknown transaction")
	}
	s.writeTx = nil
	s.txMu.Unlock()

	s.readMu.Unlock()
}

func (s *Store) writeCommit2(tx *Tx) {
	if tx == nil {
		panic("Store.writeCommit2() called outside of a transaction")
	}
	if tx.kind != txKindWrite {
		panic("Store.writeCommit2() called in a readonly transaction")
	}

	s.writeMu.Unlock()
}

func (s *Store) readCommit(tx *Tx) {
	if tx == nil {
		panic("Store.readCommit() called outside of a transaction")
	}
	if tx.kind != txKindRead {
		panic("Store.readCommit() called in a write transaction")
	}

	s.txMu.Lock()
	if _, ok := s.readTx[tx]; !ok {
		panic("Store.commit() called with an unknown transaction")
	}
	delete(s.readTx, tx)
	s.txMu.Unlock()

	s.readMu.RUnlock()
}

type txKind int

const (
	txKindRead txKind = iota
	txKindWrite
)

type txAccess struct {
	store *Store
	key   interface{}
}

type Tx struct {
	commited bool
	kind     txKind
	stores   []*Store
	reads    []txAccess
	writes   []txAccess
}

func NewReadTx() *Tx {
	return &Tx{kind: txKindRead}
}

func NewWriteTx() *Tx {
	return &Tx{kind: txKindWrite}
}

func (tx *Tx) Commit() {
	if tx.commited {
		panic("Commiting already commited Tx")
	}
	tx.commited = true

	if tx.kind == txKindRead {
		for _, i := range tx.stores {
			i.readCommit(tx)
		}
	} else {
		for _, i := range tx.stores {
			i.writeCommit1(tx)
		}

		// Get chans to update.
		// TODO(KD): there should be a lock on i.store.chans
		chans := map[chan struct{}]struct{}{}
		for _, i := range tx.writes {
			for _, j := range i.store.chans[i.key] {
				chans[j] = struct{}{}
			}
		}

		// Notify chans of update.
		for c := range chans {
			c <- struct{}{}
			<-c
		}

		for _, i := range tx.stores {
			i.writeCommit2(tx)
		}
	}
}
