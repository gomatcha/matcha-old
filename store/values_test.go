package store

import (
	"testing"
)

func TestInterface(t *testing.T) {
	v := NewInterface()
	go func() {
		v.Set(1, nil)
		v.Set(2, nil)
	}()
	c := v.Notify()
	<-c
	if v.Value(nil) != 1 {
		t.Error("Value not 1")
	}
	c <- struct{}{}

	<-c
	if v.Value(nil) != 2 {
		t.Error("Value not 2")
	}
	c <- struct{}{}

}
