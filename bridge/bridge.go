package bridge

import (
	"fmt"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/constraint"
	"reflect"
)

func Run() *Value {
	v := &NestedView{}
	n := mochi.Display(v)
	for _, i := range n.NodeChildren {
		fmt.Println("blah", i.LayoutGuide)
	}
	return &Value{reflect.ValueOf(n)}
}

type NestedView struct {
}

func (v *NestedView) Update(p *mochi.Node) *mochi.Node {
	l := constraint.New()
	n := mochi.NewNode()
	n.Layouter = l
	n.PaintOptions.BackgroundColor = mochi.GreenColor

	chl1id := "1"
	chl1 := mochi.NewBasicView(p.Get(chl1id))
	chl1.PaintOptions.BackgroundColor = mochi.RedColor
	n.Set(chl1id, chl1)
	g1 := l.AddGuide(chl1id, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2id := "2"
	chl2 := mochi.NewBasicView(p.Get(chl2id))
	chl2.PaintOptions.BackgroundColor = mochi.YellowColor
	n.Set(chl2id, chl2)
	g2 := l.AddGuide(chl2id, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3id := "3"
	chl3 := mochi.NewBasicView(p.Get(chl3id))
	chl3.PaintOptions.BackgroundColor = mochi.BlueColor
	n.Set(chl3id, chl3)
	g3 := l.AddGuide(chl3id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4id := "4"
	chl4 := mochi.NewBasicView(p.Get(chl4id))
	chl4.PaintOptions.BackgroundColor = mochi.CyanColor
	n.Set(chl4id, chl4)
	_ = l.AddGuide(chl4id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})
	_ = g3
	return n
}

func (n *NestedView) NeedsUpdate() {
	// ?
}

// End Example

type ValueSlice struct {
	v []reflect.Value
}

func NewValueSlice() *ValueSlice {
	return &ValueSlice{[]reflect.Value{}}
}

func (v *ValueSlice) Append(n *Value) {
	v.v = append(v.v, n.v)
}

func (v *ValueSlice) Len() int {
	return len(v.v)
}

func (v *ValueSlice) Index(i int) *Value {
	return &Value{v.v[i]}
}

type Value struct {
	v reflect.Value
}

// String returns the string v's underlying value, as a string. String is a special case because of Go's String method convention. Unlike the other getters, it does not panic if v's Kind is not String. Instead, it returns a string of the form "<T value>" where T is v's type. The fmt package treats Values specially. It does not call their String method implicitly but instead prints the concrete values they hold.
func (v *Value) String_() string {
	return v.v.String()
}

// Bool returns v's underlying value. It panics if v's kind is not Bool.
func (v *Value) Bool_() bool {
	return v.v.Bool()
}

// Bytes returns v's underlying value. It panics if v's underlying value is not a slice of bytes.
func (v *Value) Bytes_() []byte {
	return v.v.Bytes()
}

// Float returns v's underlying value, as a float64. It panics if v's Kind is not Float32 or Float64
func (v *Value) Float_() float64 {
	return v.v.Float()
}

// Int returns v's underlying value, as an int64. It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
func (v *Value) Int_() int64 {
	return v.v.Int()
}

// Uint returns v's underlying value, as a uint64. It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
func (v *Value) Uint_() int64 {
	return int64(v.v.Uint())
}

// Index returns v's i'th element. It panics if v's Kind is not Array, Slice, or String or i is out of range.
func (v *Value) Index(i int) *Value {
	return &Value{v.v.Index(i)}
}

// Len returns v's length. It panics if v's Kind is not Array, Chan, Map, Slice, or String.
func (v *Value) Len() int {
	return v.v.Len()
}

// MapIndex returns the value associated with key in the map v. It panics if v's Kind is not Map. It returns the zero Value if key is not found in the map or if v represents a nil map. As in Go, the key's value must be assignable to the map's key type.
func (v *Value) MapIndex(key *Value) *Value {
	x := &Value{v.v.MapIndex(key.v)}
	return x
}

// MapKeys returns a slice containing all the keys present in the map, in unspecified order. It panics if v's Kind is not Map. It returns an empty slice if v represents a nil map.
func (v *Value) MapKeys() *ValueSlice {
	return &ValueSlice{v.v.MapKeys()}
}

// Call calls the function v with the arguments arg. The return values are wrapped in an array.
func (v *Value) Call(args *ValueSlice) *ValueSlice {
	return &ValueSlice{v.v.Call(args.v)}
}

// Calls a varadic function
func (v *Value) CallSlice(args *ValueSlice) *ValueSlice {
	return &ValueSlice{v.v.CallSlice(args.v)}
}

// FieldByName returns the struct field with the given name. It returns the zero Value if no field was found. It panics if v's Kind is not struct.
func (v *Value) FieldByName(name string) *Value {
	return &Value{v.v.FieldByName(name)}
}

// MethodByName returns a function value corresponding to the method of v with the given name. The arguments to a Call on the returned function should not include a receiver; the returned function will always use v as the receiver. It returns the zero Value if no method was found.
func (v *Value) MethodByName(name string) *Value {
	return &Value{v.v.MethodByName(name)}
}

// Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
func (v *Value) Elem() *Value {
	return &Value{v.v.Elem()}
}

// IsNil reports whether its argument v is nil. The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, IsNil panics. Note that IsNil is not always equivalent to a regular comparison with nil in Go. For example, if v was created by calling ValueOf with an uninitialized interface variable i, i==nil will be true but v.IsNil will panic as v will be the zero Value.
func (v *Value) IsNil() bool {
	return v.v.IsNil()
}

// Kind returns v's Kind. If v is the zero Value (IsValid returns false), Kind returns Invalid.
func (v *Value) Kind() int {
	return int(v.v.Kind())
}

// Copy returns a copy of v.
func (v *Value) Copy() *Value {
	return &Value{v.v}
}

func (v *Value) PtrEqual(v2 *Value) bool {
	return v.v.Interface() == v2.v.Interface()
}
