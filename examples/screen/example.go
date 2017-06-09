package screen

import (
	"image/color"

	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/stackscreen"
	"github.com/overcyn/mochi/view/tabscreen"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/screen New", func() *view.Root {
		app := NewApp()
		app.Lock()
		defer app.Unlock()
		return view.NewRoot(app.NewView(nil, nil))
	})
}

type App struct {
	store        store.Store
	tabScreen    *tabscreen.Screen
	stackScreen1 *stackscreen.Screen
	stackScreen2 *stackscreen.Screen
	stackScreen3 *stackscreen.Screen
	stackScreen4 *stackscreen.Screen
}

func NewApp() *App {
	app := &App{}
	app.Lock()
	defer app.Unlock()

	app.stackScreen1 = &stackscreen.Screen{}
	app.store.Set(1, app.stackScreen1.Store())
	app.stackScreen1.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen2 = &stackscreen.Screen{}
	app.store.Set(2, app.stackScreen2.Store())
	app.stackScreen2.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen3 = &stackscreen.Screen{}
	app.store.Set(3, app.stackScreen3.Store())
	app.stackScreen3.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.stackScreen4 = &stackscreen.Screen{}
	app.store.Set(4, app.stackScreen4.Store())
	app.stackScreen4.SetChildren(
		NewTouchScreen(app, colornames.Green),
	)

	app.tabScreen = &tabscreen.Screen{}
	app.store.Set(5, app.tabScreen.Store())
	app.tabScreen.SetChildren(
		app.stackScreen1,
		app.stackScreen2,
		app.stackScreen3,
		app.stackScreen4,
	)
	return app
}

func (app *App) Lock() {
	app.store.Lock()
}

func (app *App) Unlock() {
	app.store.Unlock()
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

func (app *App) NewView(ctx *view.Context, key interface{}) view.View {
	return app.tabScreen.NewView(ctx, key)
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
