package screen

import (
	"image/color"

	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/touch"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/stackscreen"
	"github.com/overcyn/matcha/view/tabscreen"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/overcyn/matcha/examples/screen New", func() *view.Root {
		return view.NewRoot(NewApp())
	})
}

type App struct {
	comm.Storer
	store        *comm.AsyncStore
	tabScreen    *tabscreen.Screen
	stackScreen1 *stackscreen.Screen
	stackScreen2 *stackscreen.Screen
	stackScreen3 *stackscreen.Screen
	stackScreen4 *stackscreen.Screen
}

func NewApp() *App {
	st := &comm.AsyncStore{}
	app := &App{Storer: st, store: st}

	app.stackScreen1 = stackscreen.New()
	app.store.Set("1", app.stackScreen1)
	app.stackScreen1.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen2 = stackscreen.New()
	app.store.Set("2", app.stackScreen2)
	app.stackScreen2.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen3 = stackscreen.New()
	app.store.Set("3", app.stackScreen3)
	app.stackScreen3.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen4 = stackscreen.New()
	app.store.Set("4", app.stackScreen4)
	app.stackScreen4.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.tabScreen = tabscreen.New()
	app.store.Set("5", app.tabScreen)
	app.tabScreen.SetChildren(
		app.stackScreen1,
		app.stackScreen2,
		app.stackScreen3,
		app.stackScreen4,
	)
	return app
}

func (app *App) CurrentStackScreen() *stackscreen.Screen {
	switch app.tabScreen.SelectedIndex() {
	case 0:
		return app.stackScreen1
	case 1:
		return app.stackScreen2
	case 2:
		return app.stackScreen3
	case 3:
		return app.stackScreen4
	}
	return nil
}

func (app *App) View(ctx *view.Context) view.View {
	return app.tabScreen.View(ctx)
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
}

func NewTouchView(ctx *view.Context, key string, app *App) *TouchView {
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

			v.app.CurrentStackScreen().Push(NewTouchScreen(v.app, colornames.Blue))
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
