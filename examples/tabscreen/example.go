package tabscreen

import (
	"image/color"

	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/tabscreen"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/tabscreen New", func() *view.Root {
		app := NewApp()
		app.Lock()
		defer app.Unlock()
		return view.NewRoot(app.NewView(nil, nil))
	})
}

type App struct {
	store     store.Store
	tabScreen *tabscreen.Screen
}

func NewApp() *App {
	app := &App{}
	app.Lock()
	defer app.Unlock()

	screen1 := NewTouchScreen(app, colornames.Blue)
	options1 := &tabscreen.Options{
		Title: "Title 1",
	}

	screen2 := NewTouchScreen(app, colornames.Red)
	options2 := &tabscreen.Options{
		Title: "Title 2",
	}

	screen3 := NewTouchScreen(app, colornames.Yellow)
	screen4 := NewTouchScreen(app, colornames.Green)

	app.tabScreen = &tabscreen.Screen{}
	app.store.Set(0, app.tabScreen.Store())
	app.tabScreen.SetSelectedIndex(1)
	app.tabScreen.SetChildren(
		tabscreen.WithOptions(screen1, options1),
		tabscreen.WithOptions(screen2, options2),
		screen3,
		screen4,
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
	return app.TabScreen().NewView(ctx, key)
}

func (app *App) TabScreen() *tabscreen.Screen {
	return app.tabScreen
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
		OnTouch: func(e *touch.TapEvent) {
			v.app.Lock()
			defer v.app.Unlock()

			v.app.TabScreen().SetSelectedIndex(0)
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
