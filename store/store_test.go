package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	tx := NewReadTx()
	s := &Store{}
	s.ReadKey(0, tx)
	tx.Commit()
}

func TestMultiRead(t *testing.T) {
	tx1 := NewReadTx()
	tx2 := NewReadTx()
	tx3 := NewReadTx()

	s := &Store{}
	s.ReadKey(0, tx1)
	s.ReadKey(0, tx2)
	s.ReadKey(0, tx3)

	tx1.Commit()
	tx2.Commit()
	tx3.Commit()
}

func TestNotify(t *testing.T) {
	s := &Store{}
	n := s.Notifier(0)
	go func() {
		tx1 := NewWriteTx()
		s.WriteKey(0, tx1)
		tx1.Commit()
	}()

	c := n.Notify()
	<-c
	c <- struct{}{}
	n.Unnotify(c)
}

func TestReadWrite(t *testing.T) {
	s := &Store{}

	go func() {
		tx1 := NewReadTx()
		s.ReadKey(0, tx1)
		tx1.Commit()
	}()

	tx2 := NewWriteTx()
	s.ReadKey(0, tx2)
	tx2.Commit()
}
