package store

import (
	"testing"
)

func TestMatchPath(t *testing.T) {
	if !matchPath([]int64{}, []int64{}) {
		t.Fail()
	}
	if !matchPath([]int64{0}, []int64{0}) {
		t.Fail()
	}
	if matchPath([]int64{1}, []int64{0}) {
		t.Fail()
	}
	if !matchPath([]int64{0, 1, 2, 3}, []int64{0}) {
		t.Fail()
	}
	if matchPath([]int64{0}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if matchPath([]int64{1}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if !matchPath([]int64{0, 1, 2, 3}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
}

func TestMatchPath2(t *testing.T) {
	if !matchPath2([]int64{}, []int64{}) {
		t.Fail()
	}
	if !matchPath2([]int64{0}, []int64{0}) {
		t.Fail()
	}
	if matchPath2([]int64{1}, []int64{0}) {
		t.Fail()
	}
	if !matchPath2([]int64{0, 1, 2, 3}, []int64{0}) {
		t.Fail()
	}
	if !matchPath2([]int64{0}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if matchPath2([]int64{1}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
	if !matchPath2([]int64{0, 1, 2, 3}, []int64{0, 1, 2, 3}) {
		t.Fail()
	}
}
