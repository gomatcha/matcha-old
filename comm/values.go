package comm

// type Interface struct {
// 	store Store
// 	value interface{}
// }

// func (s *Interface) Notify() chan struct{} {
// 	return s.store.Notify()
// }

// func (s *Interface) Unnotify(c chan struct{}) {
// 	s.store.Unnotify(c)
// }

// func (s *Interface) Value() interface{} {
// 	return s.Get(nil)
// }

// func (s *Interface) Get(tx *Tx) interface{} {
// 	if tx == nil {
// 		tx = NewReadTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Read(tx)
// 	return s.value
// }

// func (s *Interface) Set(v interface{}, tx *Tx) {
// 	if tx == nil {
// 		tx = NewWriteTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Write(tx)
// 	s.value = v
// }

// type Bool struct {
// 	store Store3
// 	value bool
// }

// func (s *Bool) Notify() chan struct{} {
// 	return s.store.Notify()
// }

// func (s *Bool) Unnotify(c chan struct{}) {
// 	s.store.Unnotify(c)
// }

// func (s *Bool) Value() bool {
// 	return s.Get(nil)
// }

// func (s *Bool) Get(tx *Tx) bool {
// 	if tx == nil {
// 		tx = NewReadTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Read(tx)
// 	return s.value
// }

// func (s *Bool) Set(v bool, tx *Tx) {
// 	if tx == nil {
// 		tx = NewWriteTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Write(tx)
// 	s.value = v
// }

// // type Int struct{}
// // type Uint struct{}
// // type Int64 struct{}
// // type Uint64 struct{}

// type Float64 struct {
// 	store Store3
// 	value float64
// }

// func (s *Float64) Notify() chan struct{} {
// 	return s.store.Notify()
// }

// func (s *Float64) Unnotify(c chan struct{}) {
// 	s.store.Unnotify(c)
// }

// func (s *Float64) Value() float64 {
// 	return s.Get(nil)
// }

// func (s *Float64) Get(tx *Tx) float64 {
// 	if tx == nil {
// 		tx = NewReadTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Read(tx)
// 	return s.value
// }

// func (s *Float64) Set(v float64, tx *Tx) {
// 	if tx == nil {
// 		tx = NewWriteTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Write(tx)
// 	s.value = v
// }

// type String struct {
// 	store Store3
// 	value string
// }

// func (s *String) Notify() chan struct{} {
// 	return s.store.Notify()
// }

// func (s *String) Unnotify(c chan struct{}) {
// 	s.store.Unnotify(c)
// }

// func (s *String) Value() string {
// 	return s.Get(nil)
// }

// func (s *String) Get(tx *Tx) string {
// 	if tx == nil {
// 		tx = NewReadTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Read(tx)
// 	return s.value
// }

// func (s *String) Set(v string, tx *Tx) {
// 	if tx == nil {
// 		tx = NewWriteTx()
// 		defer tx.Commit()
// 	}
// 	s.store.Write(tx)
// 	s.value = v
// }

// // type Bytes struct{}