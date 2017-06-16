package tabscreen

import (
	"image/color"

	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/tabscreen"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/tabscreen New", func() *view.Root {
		return view.NewRoot(NewApp())
	})
}

type App struct {
	comm.Storer
	store     *comm.AsyncStore
	tabScreen *tabscreen.Screen
}

func NewApp() *App {
	st := &comm.AsyncStore{}
	app := &App{
		Storer: st,
		store:  st,
	}

	screen1 := NewTouchScreen(app, colornames.Blue)
	options1 := &tabscreen.TabButton{
		Title: "Title 1",
	}

	screen2 := NewTouchScreen(app, colornames.Red)
	options2 := &tabscreen.TabButton{
		Title: "Title 2",
	}

	screen3 := NewTouchScreen(app, colornames.Yellow)
	screen4 := NewTouchScreen(app, colornames.Green)

	app.tabScreen = tabscreen.NewScreen()
	app.store.Set("0", app.tabScreen)
	app.tabScreen.SetSelectedIndex(1)
	app.tabScreen.SetChildren(
		tabscreen.WithTabButton(screen1, options1),
		tabscreen.WithTabButton(screen2, options2),
		screen3,
		screen4,
	)
	return app
}

func (app *App) View(ctx *view.Context, key interface{}) view.View {
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
