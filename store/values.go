package store

import (
	"github.com/overcyn/mochi"
)

type Interface struct {
	store    *Store
	value    interface{}
	notifier mochi.Notifier
}

func NewInterface() *Interface {
	store := NewStore()
	return &Interface{
		store:    store,
		value:    nil,
		notifier: store.Notifier(0),
	}
}

func (s *Interface) Notify() chan struct{} {
	return s.notifier.Notify()
}

func (s *Interface) Unnotify(c chan struct{}) {
	s.notifier.Unnotify(c)
}

func (s *Interface) Value(tx *Tx) interface{} {
	if tx == nil {
		tx = NewReadTx()
		defer tx.Commit()
	}
	s.store.Read(0, tx)
	return s.value
}

func (s *Interface) Set(v interface{}, tx *Tx) {
	if tx == nil {
		tx = NewWriteTx()
		defer tx.Commit()
	}
	s.store.Write(0, tx)
	s.value = v
}

type String struct{}
type Bool struct{}
type Int struct{}
type Uint struct{}
type Int64 struct{}
type Uint64 struct{}
type Float64 struct{}
type Bytes struct{}
