package comm

type InterfaceValue struct {
	value interface{}
	batch GroupNotifier
}

func (v *InterfaceValue) Notify(f func()) Id {
	return v.batch.Notify(f)
}

func (v *InterfaceValue) Unnotify(id Id) {
	v.batch.Unnotify(id)
}

func (v *InterfaceValue) Value() interface{} {
	return v.value
}

func (v *InterfaceValue) SetValue(val interface{}) {
	v.value = val
	v.batch.Signal()
}

type BoolValue struct {
	value bool
	batch GroupNotifier
}

func NewBoolValue(val bool) *BoolValue {
	v := &BoolValue{}
	v.SetValue(val)
	return v
}

func (v *BoolValue) Notify(f func()) Id {
	return v.batch.Notify(f)
}

func (v *BoolValue) Unnotify(id Id) {
	v.batch.Unnotify(id)
}

func (v *BoolValue) Value() bool {
	return v.value
}

func (v *BoolValue) SetValue(val bool) {
	v.value = val
	v.batch.Signal()
}

type IntValue struct {
	value int
	batch GroupNotifier
}

func NewIntValue(val int) *IntValue {
	v := &IntValue{}
	v.SetValue(val)
	return v
}

func (v *IntValue) Notify(f func()) Id {
	return v.batch.Notify(f)
}

func (v *IntValue) Unnotify(id Id) {
	v.batch.Unnotify(id)
}

func (v *IntValue) Value() int {
	return v.value
}

func (v *IntValue) SetValue(val int) {
	v.value = val
	v.batch.Signal()
}

type Float64Value struct {
	value float64
	batch GroupNotifier
}

func NewFloat64Value(val float64) *Float64Value {
	v := &Float64Value{}
	v.SetValue(val)
	return v
}

func (v *Float64Value) Notify(f func()) Id {
	return v.batch.Notify(f)
}

func (v *Float64Value) Unnotify(id Id) {
	v.batch.Unnotify(id)
}

func (v *Float64Value) Value() float64 {
	return v.value
}

func (v *Float64Value) SetValue(val float64) {
	v.value = val
	v.batch.Signal()
}
