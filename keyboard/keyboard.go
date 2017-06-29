package keyboard // import "gomatcha.io/matcha/keyboard"

import (
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/pb/keyboard"
)

type Type int

const (
	DefaultType           Type = iota
	NumberType                 // Numbers
	NumberPunctuationType      // Numbers + Punctuation
	DecimalType                // Numbers + '.'
	PhoneType                  // Numbers + Phone keys
	ASCIIType                  // Ascii
	EmailType                  // Ascii + '@' + '.'
	URLType                    // Ascii + '.' + '/' + '.com'
	WebSearchType              // Ascii + '.' + 'go'
	NamePhoneType              // Ascii + Phone
)

func (t Type) MarshalProtobuf() keyboard.Type {
	return keyboard.Type(t)
}

type Appearance int

const (
	DefaultAppearance Appearance = iota
	LightAppearance
	DarkAppearance
)

func (a Appearance) MarshalProtobuf() keyboard.Appearance {
	return keyboard.Appearance(a)
}

type ReturnType int

const (
	DefaultReturnType ReturnType = iota
	GoReturnType
	GoogleReturnType
	JoinReturnType
	NextReturnType
	RouteReturnType
	SearchReturnType
	SendReturnType
	YahooReturnType
	DoneReturnType
	EmergencyCallReturnType
	ContinueReturnType
)

func (t ReturnType) MarshalProtobuf() keyboard.ReturnType {
	return keyboard.ReturnType(t)
}

type Responder struct {
	visible bool
	value   comm.Value
}

// func (g *Responder) Next() {
// }

// func (g *Responder) Prev() {
// }

func (g *Responder) Show() {
	if !g.visible {
		g.visible = true
		g.value.Signal()
	}
}

func (g *Responder) Dismiss() {
	if g.visible {
		g.visible = false
		g.value.Signal()
	}
}

func (g *Responder) Visible() bool {
	return g.visible
}

func (g *Responder) Notify(f func()) comm.Id {
	return g.value.Notify(f)
}

func (g *Responder) Unnotify(id comm.Id) {
	g.value.Unnotify(id)
}

// type key struct{}

// var Key = key{}

// type Middleware struct {
// 	radix *radix.Radix
// }

// func NewMiddleware() *Middleware {
// 	return &Middleware{radix: radix.NewRadix()}
// }

// func (m *Middleware) Build(ctx *view.Context, next *view.Model) {
// 	responder, ok := next.Values[Key].(*Responder)
// 	path := []int64{}
// 	for _, i := range ctx.Path() {
// 		path = append(path, int64(i))
// 	}

// 	if ok {
// 		n := m.radix.Insert(path)
// 		n.Value = responder
// 	} else {
// 		m.radix.Delete(path)
// 	}
// }
