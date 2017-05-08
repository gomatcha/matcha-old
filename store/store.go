package store

import (
	"github.com/overcyn/mochi"
	"sync"
)

func Read(f func(tx *Tx)) {
	tx := &Tx{kind: txKindRead}
	tx.begin()
	f(tx)
	tx.end(nil)
}

func Write(f func(tx *Tx)) {
	tx := &Tx{kind: txKindReadWrite}
	tx.begin()
	f(tx)
	tx.end(nil)
}

func Notifier(f func(tx *Tx)) mochi.Notifier {
	return &notifier{f: f}
}

type notifier struct {
	mu    *sync.Mutex
	close chan struct{}
	chans []chan struct{}
	f     func(tx *Tx)
}

func (n *notifier) Notify(c chan struct{}) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.close == nil {
		goChan := make(chan struct{})
		close := make(chan struct{})

		tx := &Tx{kind: txKindRead}
		tx.begin()
		n.f(tx)
		tx.end(goChan)

		go func() {
		loop:
			for {
				select {
				case <-goChan:
					n.mu.Lock()
					for _, i := range n.chans {
						i <- struct{}{}
						<-i
					}
					n.mu.Unlock()
					goChan <- struct{}{}
				case <-close:
					break loop
				}
			}
		}()
		n.close = close
	}

	n.chans = append(n.chans, c)
}

func (n *notifier) Unnotify(c chan struct{}) {
	n.mu.Lock()
	defer n.mu.Unlock()

	// remove channel
	chans := []chan struct{}{}
	for _, i := range n.chans {
		if c != i {
			chans = append(chans, i)
		}
	}
	n.chans = chans

	// Stop go routine of no channels left.
	if len(chans) == 0 {
		n.close <- struct{}{}
		n.close = nil
	}
}

type txKind int

const (
	txKindRead txKind = iota
	txKindReadWrite
)

type txState int

const (
	txStateBefore txState = iota
	txStateDuring
	txStateAfter
)

type txAccess struct {
	store *Store
	key   interface{}
}

type Tx struct {
	state  txState
	kind   txKind
	stores []*Store
	reads  []txAccess
	writes []txAccess
}

func (tx *Tx) begin() {
	if tx.state != txStateBefore {
		panic("Starting already started transaction")
	}
	tx.state = txStateDuring
}

func (tx *Tx) end(c chan struct{}) map[*Store]struct{} {
	if tx.state != txStateDuring {
		panic("Ending already unstarted or completed transaction")
	}
	tx.state = txStateAfter

	// Lock stores
	for _, i := range tx.stores {
		i.mu.Lock()
		defer i.mu.Unlock()
		i.tx = nil
	}

	// Notify all listeners of updates
	for _, i := range tx.writes {
		for _, j := range tx.stores {
			if i.store == j {
				j.unlockedUpdate(i.key)
				break
			}
		}
	}

	// Add new listener if necessary
	if c == nil {
		return nil
	}
	s := map[*Store]struct{}{}
	for _, i := range tx.reads {
		for _, j := range tx.stores {
			if i.store == j {
				s[j] = struct{}{}
				j.unlockedNotify(i.key, c)
				break
			}
		}
	}
	return s
}

func (tx *Tx) Include(s *Store) {
	if tx.state != txStateDuring {
		panic("Using transaction outside of block")
	}

	// Give the store a reference to the transaction
	s.mu.Lock()
	s.tx = tx
	s.mu.Unlock()

	// Add the store to our list of transactions
	found := false
	for _, i := range tx.stores {
		if i == s {
			found = true
			break
		}
	}
	if !found {
		tx.stores = append(tx.stores, s)
	}
}

type Store struct {
	mu    *sync.Mutex
	tx    *Tx
	chans map[interface{}][]chan struct{}
}

func (s *Store) Write(key interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tx == nil {
		panic("Store.Write() called outside of transaction")
	}
	if s.tx.kind != txKindReadWrite {
		panic("Store.Write() called in readonly transaction")
	}
	s.tx.writes = append(s.tx.writes, txAccess{store: s, key: key})
}

func (s *Store) Read(key interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tx == nil {
		panic("Store.Read() called outside of transaction")
	}
	s.tx.reads = append(s.tx.reads, txAccess{store: s, key: key})
}

func (s *Store) unlockedUpdate(key interface{}) {
	for _, i := range s.chans[key] {
		i <- struct{}{}
		<-i
	}
}

func (s *Store) unlockedNotify(key interface{}, c chan struct{}) {
	s.chans[key] = append(s.chans[key], c)
}

func (s *Store) unlockedUnnotify(c chan struct{}) {
loop:
	for k, v := range s.chans {
		for idx, i := range v {
			if i == c {
				s.chans[k] = append(v[:idx], v[idx+1:]...)
				break loop
			}
		}
	}
}

// type MapStore interface {
// 	Set(interface{})
// 	Get(interface{})
// 	Del(interface{})
// 	Len()
// }

// type ArrayStore interface {
// 	At(int)
// 	Set(int)
// 	Insert(int)
// 	Del(int)
// 	Len()
// }

// func Tx(f func(), s ...Store) {
// }

// type store struct {
// }

// func (s *store) TxBegin() {
// }

// func (s *store) TxEnd() {
// }

// func (s *store) Value() interface{} {
//     return s.value
// }

// func (s *Store) SetValue(v interface{}) {
//     s.value = v
// }

// type storeNotifier struct {
//     observes []Storer
// }

// type Storer interface {
//     Store() Store
// }

// // How to observe datastructures and get notifications
// type Store struct {
//     // TxBegin()
//     // TxEnd()
//     // Value() interface{}
//     // SetValue(interface{})
//     // Tx(func(s Store)) Notifier
// }

// func NewStore() Store {
//     return nil
// }

// type Int struct {
//     Store2
//     value int
// }

// func (s *Int) TxBegin() {
//     (*Store)(s).TxBegin()
// }

// func (s *Int) TxEnd() {
//     (*Store)(s).TxEnd()
// }

// func (s *Int) Value() int {
//     return s.value
// }

// func (s *Int) SetValue(v int) {
//     s.value = v
// }

// // type InterfaceStore interface {
// //     Value() interface{}
// // }

// // type interfaceStore struct {
// //     value interface{}
// // }

// // func (s *interfaceStore) Value() interface{} {
// //     return s.value
// // }

// // func (s *interfaceStore) SetValue(v interface{}) {
// //     s.value = v
// // }

// func NewInterfaceStore(v interface{}) InterfaceStore {
//     return &interfaceStore{v}
// }

// type MapStore interface {
//     Set(interface{}) Store
//     Get(interface{}) Store
//     Del(interface{})
//     Keys() []interface{}

//     Notifier
// }

// func NewMapStore(s Store) MapStore {
//     return nil
// }

// type ArrayStore interface {
// Len() int
// At(int) interface{}
// Set(int, interface{})
// Insert(int, interface{})
// Del(int)

//     Notifier
// }

// func NewArrayStore(s Store) ArrayStore {
//     return nil
// }

// type SetStore interface {
//     Add(Store)
//     Del(Store)
//     Contains(Store) bool
//     Len() int
//     Elems() []Store
//     Notifier
//     Store
// }

// type UsersStore struct {
//     store Store
//     users [string]*UserStore
// }

// func (s *UsersStore) Add(id string, u *UserStore) {
//     s.store.Set(id)
//     s.store.Use(u)
//     // s.store.Tx(u)
//     s.users[id] = u
// }

// func (s *UserStore) Get(id string) (*UserStore, bool) {
//     s.store.Get(id)
//     return users[id]
// }

// func (s *UsersStore) Del(id string) {
//     s.store.Set(id)
//     del(users[id])
// }

// type UserStore struct {
//     store     Store
//     firstName string
//     lastName  string
// }

// func NewUserStore(s Store) (*UserStore, err) {
//     return &UserStore{
//         mapStore: mapStore,
//     }
// }

// func (s *UserStore) SetId(id string) {
//     s.store.Set("id")
//     s.id = id
// }

// func (s *UserStore) Id() string {
//     s.store.Get("id")
//     return s.id
// }

// func a() {
//     var fullName StringNotifier = store.ReadTx(func() string {
//         user := userStore.Get(id)
//         if user {
//             return user.FirstName() + user.LastName()
//         }
//         return ""
//     }, userStore) // block access to other stores somehow??

//     store.Read(func(tx store.Tx) {
//         tx.Include(usersStore)

//         user := userStore.Get(id)
//         if user {
//             return user.FirstName() + user.LastName()
//         }
//         return ""
//     })

//     store.Write(func(tx store.Tx) {
//         tx.Include(usersStore)

//         user := NewUserStore()
//         tx.Include(user)
//         user.SetFirstName("dog")
//         user.SetLastName("cat")

//         usersStore.Add(user)
//     })
// }

// type FamilyStore Store

// func NewFamilyStore(m Store) FamilyStore {
//  return &FamilyStore{
//      m: m,
//  }
// }

// func (f *FamilyStore) Members() []Members {
//  return nil
// }

// func (f *FamilyStore) Len() int {
//  return s.Len()
// }

// At(int) (bool, interface{})
// Set(int, interface{})
// Insert(int, interface{})
// Del(int)
// Notifier(int) Notifier
// Notifier
// Store

// func (f *FamilyStore) Len() int {
//  return 0
// }

// func (f *FamilyStore) At() *Member {
// }

// func (f *FamilyStore) NotifierAt(idx int) Notifier {
//  return nil
// }

// func (f *FamilyStore) Notify(chan struct{}) {
// }

// func (f *FamilyStore) Unnotify(chan struct{}) {
// }

// // Mutable arrays

// func (f *FamilyStore) Remove(m *Member, idx int) {
// }

// func (f *FamilyStore) Set(m *Member, idx int) {
// }

// func (f *FamilyStore) Insert(m *Member) {
// }

// // Investigate Redis and Group Cache, Socket.io
// // Maybe rename Notfier to Notify/Unnotify?

// type FamilyModel struct {
// }

// func (m *FamilyModel) Members() []PersonModel {
// }

// func (m *FamilyModel) SetMembers(s []PersonModel) {
// }

// func (m *FamilyModel) MembersNotifier() Notifier {
// }

// type PersonModel struct {
// }

// func (m *PersonModel) FirstName() StringNotifier {
// }

// func (m *PersonModel) SetFirstName(s string) {
// }

// func (m *PersonModel) LastName() StringNotifier {
// }

// func (m *PersonModel) SetLastName() StringNotifier {
// }

// func (m *PersonModel) Tx(f func(*PersonModel)) {
// }

// model.Tx(func (m *PersonModel){
//     m.SetLastName("Kevin")
//     m.SetFirstName("Dang")
// })

// // How can we do transactions with ValueObjects?
// // PreSet()
// // Send()
// // Maybe make it so you can't transaction across objects

// type NameModel struct {
//  state NameState
// }

// func (m *NameModel) SetState(s NameState) {
// }

// func (m *NameModel) State() NameState {

// }

// func (m *NameModel) NotifierForPath(p string) {
// }

// type NameState struct {
//  FirstName StringValue `mochi: observable`
//  LastName  StringValue `mochi: observable`
// }

// type NameValue2 struct {
//  firstName StringValue
//  lastName  StringValue
// }

// func (v *NameValue) Notify(c chan struct{}) {
// }

// func (v *NameValue) Unnotify(c chan struct{}) {

// }

// func (v *NameValue) FullName() string {
//  return v.firstName + v.lastName
// }
