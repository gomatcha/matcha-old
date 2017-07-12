package stackscreen

import (
	"image/color"

	"golang.org/x/image/colornames"
	"gomatcha.io/bridge"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/touch"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
	"gomatcha.io/matcha/view/stackscreen"
)

func init() {
	bridge.RegisterFunc("gomatcha.io/matcha/examples/stackscreen New", func() *view.Root {
		return view.NewRoot(NewApp())
	})
}

type App struct {
	store.Node
	stackScreen *stackscreen.Screen
}

func NewApp() *App {
	app := &App{}

	screen1 := NewTouchScreen(app, colornames.Blue)
	bar1 := &stackscreen.Bar{
		Title: "Title 1",
	}

	screen2 := NewTouchScreen(app, colornames.Red)
	bar2 := &stackscreen.Bar{
		Title: "Title 2",
	}

	screen3 := NewTouchScreen(app, colornames.Yellow)
	screen4 := NewTouchScreen(app, colornames.Green)

	app.stackScreen = &stackscreen.Screen{}
	app.Set("stackscreen", app.stackScreen)
	app.stackScreen.SetChildren(
		stackscreen.WithBar(screen1, bar1),
		stackscreen.WithBar(screen2, bar2),
		screen3,
		screen4,
	)
	return app
}

func (app *App) View(ctx *view.Context) view.View {
	return app.StackScreen().View(ctx)
}

func (app *App) StackScreen() *stackscreen.Screen {
	return app.stackScreen
}

func NewTouchScreen(app *App, c color.Color) view.Screen {
	return view.ScreenFunc(func(ctx *view.Context) view.View {
		chl := NewTouchView(ctx, "", app)
		chl.Color = c
		return chl
	})
}

type TouchView struct {
	*view.Embed
	app   *App
	Color color.Color
	bar   *stackscreen.Bar
}

func NewTouchView(ctx *view.Context, key string, app *App) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: ctx.NewEmbed(key),
		app:   app,
	}
}

func (v *TouchView) Build(ctx *view.Context) *view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		OnTouch: func(e *touch.TapEvent) {
			// v.bar.Title = "Updated"
			// v.Signal()

			v.app.Lock()
			defer v.app.Unlock()

			v.app.StackScreen().Push(NewTouchScreen(v.app, colornames.Purple))
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}

func (v *TouchView) StackBar(ctx *view.Context) *stackscreen.Bar {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.HeightEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(100))
	})

	titleView := basicview.New(ctx, "axbaba")
	titleView.Painter = &paint.Style{BackgroundColor: colornames.Red}
	titleView.Layouter = l

	l2 := constraint.New()
	l2.Solve(func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.HeightEqual(constraint.Const(50))
		s.WidthEqual(constraint.Const(50))
	})
	rightView := basicview.New(ctx, "right")
	rightView.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	rightView.Layouter = l2

	l3 := constraint.New()
	l3.Solve(func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.HeightEqual(constraint.Const(50))
		s.WidthEqual(constraint.Const(50))
	})
	leftView := basicview.New(ctx, "left")
	leftView.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	leftView.Layouter = l3

	return &stackscreen.Bar{
		Title:      "Title",
		TitleView:  titleView,
		RightViews: []view.View{rightView},
		LeftViews:  []view.View{leftView},
	}
}
