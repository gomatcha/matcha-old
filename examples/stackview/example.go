package stackview

import (
	"image/color"

	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/stacknav"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/stackview New", func() *view.Root {
		app := NewApp()
		return view.NewRoot(app.NewView(nil, nil))
	})
}

type App struct {
	stackScreen *stacknav.Screen
}

func NewApp() *App {
	tx := store.NewWriteTx()
	defer tx.Commit()

	app := &App{}
	screen := &stacknav.Screen{}
	screen.SetChildren(tx,
		view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
			chl := NewTouchView(ctx, key, app)
			chl.Color = colornames.Green
			return chl
		}),
		view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
			chl := NewTouchView(ctx, key, app)
			chl.Color = colornames.Blue
			return chl
		}),
		view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
			chl := NewTouchView(ctx, key, app)
			chl.Color = colornames.Red
			return chl
		}),
		view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
			chl := NewTouchView(ctx, key, app)
			chl.Color = colornames.Yellow
			return chl
		}),
	)
	app.stackScreen = screen
	return app
}

func (app *App) NewView(ctx *view.Context, key interface{}) view.View {
	return app.StackScreen().NewView(ctx, key)
}

func (app *App) StackScreen() *stacknav.Screen {
	return app.stackScreen
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
	return &view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
	}
}
