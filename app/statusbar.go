package app

import (
	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/internal"
	"gomatcha.io/matcha/internal/radix"
	"gomatcha.io/matcha/pb/app"
	"gomatcha.io/matcha/view"
)

const (
	// TODO(KD): If there are multiple views with a StatusBarKey, a random one will be chosen...
	// values[app.StatusBarKey] = &StatusBar{Hidden:false, Style:StatusBarStyleLight}
	StatusBarKey = "gomatcha.io/matcha/app statusbar"

	// If any view has set the ActivityIndicatorKey set to true, the activity indicator will be visible.
	// values[app.ActivityIndicatorKey] = true
	ActivityIndicatorKey = "gomatcha.io/matcha/app activity"
)

type StatusBarStyle int

const (
	StatusBarStyleDefault StatusBarStyle = iota
	StatusBarStyleLight
	StatusBarStyleDark
)

type StatusBar struct {
	Hidden bool
	Style  StatusBarStyle
}

type statusBarMiddleware struct {
	radix *radix.Radix
}

// func (m *statusBarMiddleware) Build(ctx *view.Context, model *view.Model) {
// 	path := idSliceToIntSlice(ctx.Path())

// 	add := false
// 	val, ok := model.Values[ActivityIndicatorKey]
// 	if ok {
// 		if val2, ok := val.(bool); ok {
// 			add = val2
// 		}
// 	}
// 	if add {
// 		m.radix.Insert(path)
// 	} else {
// 		m.radix.Delete(path)
// 	}
// }

// func (m *statusBarMiddleware) MarshalProtobuf() proto.Message {
// 	visible := false
// 	m.radix.Range(func(path []int64, node *radix.Node) {
// 		visible = true
// 	})
// 	return &app.ActivityIndicator{
// 		Visible: visible,
// 	}
// }

// func (m *statusBarMiddleware) Key() string {
// 	return ActivityIndicatorKey
// }

func init() {
	internal.RegisterMiddleware(func() interface{} {
		return &statusBarMiddleware{
			radix: radix.NewRadix(),
		}
	})
}

type activityIndicatorMiddleware struct {
	radix *radix.Radix
}

func (m *activityIndicatorMiddleware) Build(ctx *view.Context, model *view.Model) {
	path := idSliceToIntSlice(ctx.Path())

	add := false
	val, ok := model.Values[ActivityIndicatorKey]
	if ok {
		if val2, ok := val.(bool); ok {
			add = val2
		}
	}
	if add {
		m.radix.Insert(path)
	} else {
		m.radix.Delete(path)
	}
}

func (m *activityIndicatorMiddleware) MarshalProtobuf() proto.Message {
	visible := false
	m.radix.Range(func(path []int64, node *radix.Node) {
		visible = true
	})
	return &app.ActivityIndicator{
		Visible: visible,
	}
}

func (m *activityIndicatorMiddleware) Key() string {
	return ActivityIndicatorKey
}

func idSliceToIntSlice(ids []matcha.Id) []int64 {
	ints := make([]int64, len(ids))
	for idx, i := range ids {
		ints[idx] = int64(i)
	}
	return ints
}
