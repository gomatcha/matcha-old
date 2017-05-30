package pb

// import (
// 	"reflect"
// )

// type Encoder struct {
// 	funcs     map[int64]reflect.Value
// 	maxFuncId int64
// }

// func NewEncoder() *Encoder {
// 	return &Encoder{
// 		funcs: map[int64]reflect.Value{},
// 	}
// }

// func (e *Encoder) AddFunc(f interface{}) int64 {
// 	e.maxFuncId += 1
// 	e.funcs[e.maxFuncId] = reflect.ValueOf(f)
// 	return e.maxFuncId
// }
