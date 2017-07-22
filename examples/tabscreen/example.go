package tabscreen

import (
	"image/color"

	"golang.org/x/image/colornames"
	"gomatcha.io/bridge"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/touch"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/tabscreen"
)

func init() {
	bridge.RegisterFunc("gomatcha.io/matcha/examples/tabscreen New", func() *view.Root {
		return view.NewRoot(NewAppView())
	})
}

type App struct {
	tabs *tabscreen.Tabs
}

func NewAppView() view.View {
	app := &App{}

	screen1 := NewTouchScreen(app, colornames.Blue)
	options1 := &tabscreen.Button{
		Title: "Title 1",
		Badge: "badge",
		// Icon:         env.MustLoadImage("TabCamera"),
		// SelectedIcon: env.MustLoadImage("TabCameraFilled"),
	}

	screen2 := NewTouchScreen(app, colornames.Red)
	options2 := &tabscreen.Button{
		Title: "Title 2",
		// Icon:         env.MustLoadImage("TabMap"),
		// SelectedIcon: env.MustLoadImage("TabMapFilled"),
	}

	screen3 := NewTouchScreen(app, colornames.Yellow)
	screen4 := NewTouchScreen(app, colornames.Green)

	app.tabs = &tabscreen.Tabs{}
	app.tabs.SetSelectedIndex(1)
	app.tabs.SetViews(
		tabscreen.WithButton(screen1, options1),
		tabscreen.WithButton(screen2, options2),
		screen3,
		screen4,
	)

	v := tabscreen.New(nil, "")
	v.Tabs = app.tabs
	return v
}

func NewTouchScreen(app *App, c color.Color) view.View {
	chl := NewTouchView(nil, "", app)
	chl.Color = c
	return chl
}

type TouchView struct {
	view.Embed
	app    *App
	Color  color.Color
	button *tabscreen.Button
}

func NewTouchView(ctx *view.Context, key string, app *App) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: ctx.NewEmbed(key),
		app:   app,
		button: &tabscreen.Button{
			Title: "Testing",
			// Icon:         env.MustLoadImage("TabSearch"),
			// SelectedIcon: env.MustLoadImage("TabSearchFilled"),
		},
	}
}

func (v *TouchView) Build(ctx *view.Context) view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		OnTouch: func(e *touch.TapEvent) {
			v.app.tabs.SetSelectedIndex(0)
			v.button.Title = "Updated"
			v.Signal()
		},
	}

	return view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}

func (v *TouchView) TabButton(*view.Context) *tabscreen.Button {
	return v.button
}
