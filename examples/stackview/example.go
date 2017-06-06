package stackview

import (
	"image/color"

	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/stacknav"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/stackview New", func() *view.Root {
		app := NewApp()
		app.Lock()
		defer app.Unlock()
		return view.NewRoot(app.NewView(nil, nil))
	})
}

type App struct {
	store       store.Store
	stackScreen *stacknav.Screen
}

func NewApp() *App {
	app := &App{
		stackScreen: &stacknav.Screen{},
	}
	app.Lock()
	defer app.Unlock()

	app.store.AddChild(app.StackScreen().Store(), "set")
	app.stackScreen.SetChildren(
		NewTouchScreen(app, colornames.Blue),
		NewTouchScreen(app, colornames.Red),
		NewTouchScreen(app, colornames.Yellow),
		NewTouchScreen(app, colornames.Green),
	)
	return app
}

func (app *App) Lock() {
	app.store.Lock()
}

func (app *App) Unlock() {
	app.store.Unlock()
}

func (app *App) NewView(ctx *view.Context, key interface{}) view.View {
	return app.StackScreen().NewView(ctx, key)
}

func (app *App) StackScreen() *stacknav.Screen {
	app.store.Read()
	return app.stackScreen
}

func NewTouchScreen(app *App, c color.Color) view.Screen {
	return view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
		chl := NewTouchView(ctx, key, app)
		chl.Color = c
		return chl
	})
}

type TouchView struct {
	*view.Embed
	app   *App
	Color color.Color
}

func NewTouchView(ctx *view.Context, key interface{}, app *App) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: view.NewEmbed(ctx.NewId(key)),
		app:   app,
	}
}

func (v *TouchView) Build(ctx *view.Context) *view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		RecognizedFunc: func(e *touch.TapEvent) {
			v.app.Lock()
			defer v.app.Unlock()

			v.app.StackScreen().Push(NewTouchScreen(v.app, colornames.Blue))
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
