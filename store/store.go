package store

import (
	"github.com/overcyn/mochi"
	"sync"
)

type storeNotifier struct {
	store *Store
	key   interface{}
}

func (s *storeNotifier) Notify() chan struct{} {
	s.store.chansMu.Lock()
	defer s.store.chansMu.Unlock()

	c := make(chan struct{})
	s.store.chans[s.key] = append(s.store.chans[s.key], c)
	return c
}

func (s *storeNotifier) Unnotify(c chan struct{}) {
	s.store.chansMu.Lock()
	defer s.store.chansMu.Unlock()

	chans := s.store.chans[s.key]
	copy := make([]chan struct{}, 0, len(chans)-1)
	for _, i := range chans {
		if i != c {
			copy = append(copy, i)
		}
	}
	s.store.chans[s.key] = copy
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

func NewStore() *Store {
	return &Store{
		chans:  map[interface{}][]chan struct{}{},
		readTx: map[*Tx]struct{}{},
	}
}

func (s *Store) Write(key interface{}, tx *Tx) {
	if tx == nil {
		panic("Store.Write() called outside of a transaction")
	}
	if tx.kind != txKindWrite {
		panic("Store.Write() called in readonly transaction")
	}

	s.lock(tx)
	tx.writes = append(tx.writes, txAccess{store: s, key: key})
}

func (s *Store) Read(key interface{}, tx *Tx) {
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
