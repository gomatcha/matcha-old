package radix

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	r := NewRadix()
	r.Insert([]int64{0, 10})

	fmt.Println(r.String())

	// t.Error("Ticker did not trigger")
}
