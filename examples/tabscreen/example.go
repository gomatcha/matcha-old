package tabscreen

import (
	"image/color"

	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/env"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/touch"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/tabscreen"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/overcyn/matcha/examples/tabscreen New", func() *view.Root {
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
		Title:        "Title 1",
		Badge:        "badge",
		Icon:         env.MustLoadImage("TabCamera"),
		SelectedIcon: env.MustLoadImage("TabCameraFilled"),
	}

	screen2 := NewTouchScreen(app, colornames.Red)
	options2 := &tabscreen.TabButton{
		Title:        "Title 2",
		Icon:         env.MustLoadImage("TabMap"),
		SelectedIcon: env.MustLoadImage("TabMapFilled"),
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

func (app *App) View(ctx *view.Context) view.View {
	return app.TabScreen().View(ctx)
}

func (app *App) TabScreen() *tabscreen.Screen {
	return app.tabScreen
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
	app    *App
	Color  color.Color
	button *tabscreen.TabButton
}

func NewTouchView(ctx *view.Context, key string, app *App) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: view.NewEmbed(ctx.NewId(key)),
		app:   app,
		button: &tabscreen.TabButton{
			Title:        "Testing",
			Icon:         env.MustLoadImage("TabSearch"),
			SelectedIcon: env.MustLoadImage("TabSearchFilled"),
		},
	}
}

func (v *TouchView) Build(ctx *view.Context) *view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		OnTouch: func(e *touch.TapEvent) {
			v.button.Title = "Updated"
			v.Update()

			// v.app.Lock()
			// defer v.app.Unlock()

			// v.app.TabScreen().SetSelectedIndex(0)
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}

func (v *TouchView) TabButton(*view.Context) *tabscreen.TabButton {
	return v.button
}
