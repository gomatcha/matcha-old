package screen

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"

	"gomatcha.io/bridge"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/touch"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/stackscreen"
	"gomatcha.io/matcha/view/tabscreen"
)

func init() {
	bridge.RegisterFunc("gomatcha.io/matcha/examples/screen New", func() *view.Root {
		return view.NewRoot(NewAppView())
	})
}

type App struct {
	tabs   *tabscreen.Tabs
	stack1 *stackscreen.Stack
	stack2 *stackscreen.Stack
	stack3 *stackscreen.Stack
	stack4 *stackscreen.Stack
}

func NewAppView() view.View {
	app := &App{
		stack1: &stackscreen.Stack{},
		stack2: &stackscreen.Stack{},
		stack3: &stackscreen.Stack{},
		stack4: &stackscreen.Stack{},
		tabs:   &tabscreen.Tabs{},
	}

	// Configure the stacks
	app.stack1.SetViews(NewTouchView(nil, "", app))
	app.stack2.SetViews(NewTouchView(nil, "", app))
	app.stack3.SetViews(NewTouchView(nil, "", app))
	app.stack4.SetViews(NewTouchView(nil, "", app))

	// Configure the tabs
	stackview1 := stackscreen.New(nil, "")
	stackview1.Stack = app.stack1
	stackview2 := stackscreen.New(nil, "")
	stackview2.Stack = app.stack2
	stackview3 := stackscreen.New(nil, "")
	stackview3.Stack = app.stack3
	stackview4 := stackscreen.New(nil, "")
	stackview4.Stack = app.stack4
	app.tabs.SetViews(
		stackview1,
		stackview2,
		stackview3,
		stackview4,
	)

	// Return tabview
	v := tabscreen.New(nil, "")
	v.Tabs = app.tabs
	return v
}

func (app *App) CurrentStackScreen() *stackscreen.Stack {
	switch app.tabs.SelectedIndex() {
	case 0:
		return app.stack1
	case 1:
		return app.stack2
	case 2:
		return app.stack3
	case 3:
		return app.stack4
	}
	return nil
}

type TouchView struct {
	view.Embed
	app   *App
	Color color.Color
}

func NewTouchView(ctx *view.Context, key string, app *App) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: ctx.NewEmbed(key),
		Color: colornames.White,
		app:   app,
	}
}

func (v *TouchView) Build(ctx *view.Context) view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		OnTouch: func(e *touch.TapEvent) {
			child := NewTouchView(nil, "", v.app)
			child.Color = colornames.Red
			v.app.CurrentStackScreen().Push(child)
			fmt.Println("child", child)
		},
	}

	return view.Model{
		Painter: &paint.Style{BackgroundColor: v.Color},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
