package examples

import (
	"gomatcha.io/bridge"
	"gomatcha.io/matcha/examples/animate"
	"gomatcha.io/matcha/examples/complex"
	"gomatcha.io/matcha/examples/constraints"
	"gomatcha.io/matcha/examples/custom"
	"gomatcha.io/matcha/examples/imageview"
	"gomatcha.io/matcha/examples/paint"
	"gomatcha.io/matcha/examples/screen"
	"gomatcha.io/matcha/examples/settings"
	"gomatcha.io/matcha/examples/table"
	"gomatcha.io/matcha/examples/textview"
	"gomatcha.io/matcha/examples/touch"
	"gomatcha.io/matcha/view"
)

func init() {
	bridge.RegisterFunc("gomatcha.io/matcha/examples New", func(str string) *view.Root {
		return New(str)
	})
}

func New(str string) *view.Root {
	var v func(ctx *view.Context) view.View
	var s view.Screen

	switch str {
	case "animate":
		v = func(ctx *view.Context) view.View { return animate.New(ctx, "") }
	case "complex":
		v = func(ctx *view.Context) view.View { return complex.New(ctx, "") }
	case "constraints":
		v = func(ctx *view.Context) view.View { return constraints.New(ctx, "") }
	case "custom":
		v = func(ctx *view.Context) view.View { return custom.New(ctx, "") }
	case "imageview":
		v = func(ctx *view.Context) view.View { return imageview.New(ctx, "") }
	case "paint":
		v = func(ctx *view.Context) view.View { return paint.New(ctx, "") }
	case "screen":
		s = screen.NewApp()
	case "settings":
		s = settings.NewApp()
	// case "stackscreen":
	// 	s = stackscreen.NewApp()
	case "table":
		v = func(ctx *view.Context) view.View { return table.New(ctx, "") }
	// case "tabscreen":
	// 	s = tabscreen.NewApp()
	case "textview":
		v = func(ctx *view.Context) view.View { return textview.New(ctx, "") }
	// case "todo":
	// 	stack := &ss.Screen{}
	// 	stack.SetChildren(view.ScreenFunc(func(ctx *view.Context) view.View {
	// 		app := todo.NewAppView(ctx, "")
	// 		app.Todos = []*todo.Todo{
	// 			&todo.Todo{Title: "Title1"},
	// 			&todo.Todo{Title: "Title2"},
	// 			&todo.Todo{Title: "Title3"},
	// 		}
	// 		return app
	// 	}))
	// 	s = stack
	case "touch":
		v = func(ctx *view.Context) view.View { return touch.New(ctx, "") }
	}
	if s != nil {
		return view.NewRoot(s)
	}
	return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
		return v(ctx)
	}))
}
