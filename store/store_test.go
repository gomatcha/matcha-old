package store

// import (
// 	"testing"
// )

// func TestStore(t *testing.T) {
// 	tx := NewReadTx()
// 	s := NewStore()
// 	s.Read(0, tx)
// 	tx.Commit()
// }

// func TestMultiRead(t *testing.T) {
// 	tx1 := NewReadTx()
// 	tx2 := NewReadTx()
// 	tx3 := NewReadTx()

// 	s := NewStore()
// 	s.Read(0, tx1)
// 	s.Read(0, tx2)
// 	s.Read(0, tx3)

// 	tx1.Commit()
// 	tx2.Commit()
// 	tx3.Commit()
// }

// func TestNotify(t *testing.T) {
// 	s := NewStore()
// 	n := s.Notifier(0)
// 	go func() {
// 		tx1 := NewWriteTx()
// 		s.Write(0, tx1)
// 		tx1.Commit()
// 	}()

// 	c := n.Notify()
// 	<-c
// 	c <- struct{}{}
// }

// func TestReadWrite(t *testing.T) {
// 	s := NewStore()

// 	go func() {
// 		tx1 := NewReadTx()
// 		s.Read(0, tx1)
// 		tx1.Commit()
// 	}()

// 	tx2 := NewWriteTx()
// 	s.Read(0, tx2)
// 	tx2.Commit()
// }
