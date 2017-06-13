package comm

// import (
// 	"testing"
// )

// func TestInterface(t *testing.T) {
// 	v := &Interface{}
// 	go func() {
// 		v.Set(1, nil)
// 		v.Set(2, nil)
// 	}()
// 	c := v.Notify()
// 	<-c
// 	if v.Get(nil) != 1 {
// 		t.Error("Get not 1")
// 	}
// 	c <- struct{}{}

// 	<-c
// 	if v.Get(nil) != 2 {
// 		t.Error("Get not 2")
// 	}
// 	c <- struct{}{}

// }
