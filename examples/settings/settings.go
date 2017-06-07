package settings

import (
	"image"

	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/stackscreen"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/settings New", func() *view.Root {
		app := NewApp()
		app.Lock()
		defer app.Unlock()
		return view.NewRoot(app.NewView(nil, nil))
	})
}

type App struct {
	store       store.Store
	stackScreen *stackscreen.Screen
}

func NewApp() *App {
	app := &App{}
	app.Lock()
	defer app.Unlock()

	rootScreen := view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
		return NewRootView(ctx, key, app)
	})
	rootScreen2 := stackscreen.WithOptions(rootScreen, &stackscreen.Options{
		Title: "Settings",
	})

	app.stackScreen = &stackscreen.Screen{}
	app.store.Set(0, app.stackScreen.Store())
	app.stackScreen.SetChildren(rootScreen2)
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

func (app *App) StackScreen() *stackscreen.Screen {
	return app.stackScreen
}

type RootView struct {
	*view.Embed
	app *App
}

func NewRootView(ctx *view.Context, key interface{}, app *App) *RootView {
	if v, ok := ctx.Prev(key).(*RootView); ok {
		return v
	}
	return &RootView{Embed: view.NewEmbed(ctx.NewId(key)), app: app}
}

func (v *RootView) Build(ctx *view.Context) *view.Model {
	l := &table.Layout{}
	chlds := []view.View{}

	{
		group := []view.View{}
		airplaneCell := NewBasicCell(ctx, 0)
		airplaneCell.Title = "Airplane Mode"
		airplaneCell.Subtitle = "Enabled"
		group = append(group, airplaneCell)

		wifiCell := NewBasicCell(ctx, 1)
		wifiCell.Title = "Wi-Fi"
		wifiCell.Subtitle = "Home Wifi"
		group = append(group, wifiCell)

		bluetoothCell := NewBasicCell(ctx, 2)
		bluetoothCell.Title = "Bluetooth"
		bluetoothCell.Subtitle = "On"
		group = append(group, bluetoothCell)

		cellularCell := NewBasicCell(ctx, 3)
		cellularCell.Title = "Cellular"
		group = append(group, cellularCell)

		for _, i := range AddSeparators(ctx, "a", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}

	scrollChild := basicview.New(ctx, 6)
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = l
	scrollChild.Children = chlds

	scrollView := scrollview.New(ctx, 5)
	scrollView.ContentView = scrollChild

	return &view.Model{
		Children: []view.View{scrollView},
		// Layouter: l,
		Painter: &paint.Style{BackgroundColor: colornames.Lightgray},
	}
}

type separatorKey struct {
	index int
	key   interface{}
}

func AddSeparators(ctx *view.Context, key interface{}, vs []view.View) []view.View {
	newViews := []view.View{}

	top := NewSeparator(ctx, separatorKey{-1, key})
	newViews = append(newViews, top)

	for idx, i := range vs {
		newViews = append(newViews, i)

		if idx != len(vs)-1 { // Don't add short separator after last view
			sep := NewSeparator(ctx, separatorKey{idx, key})
			sep.LeftPadding = 25
			newViews = append(newViews, sep)
		}
	}

	bot := NewSeparator(ctx, separatorKey{-2, key})
	newViews = append(newViews, bot)
	return newViews
}

type Separator struct {
	*view.Embed
	LeftPadding float64
}

func NewSeparator(ctx *view.Context, key interface{}) *Separator {
	if v, ok := ctx.Prev(key).(*Separator); ok {
		return v
	}
	return &Separator{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *Separator) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(0.5))
		s.WidthEqual(l.MaxGuide().Width())
	})

	chl := basicview.New(ctx, 0)
	chl.Painter = &paint.Style{BackgroundColor: colornames.Gray}
	l.Add(chl, func(s *constraint.Solver) {
		s.HeightEqual(l.Height())
		s.LeftEqual(l.Left().Add(v.LeftPadding))
		s.RightEqual(l.Right())
	})

	return &view.Model{
		Children: []view.View{chl},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.White},
	}
}

type BasicCell struct {
	*view.Embed
	Icon     image.Image
	Title    string
	Subtitle string
}

func NewBasicCell(ctx *view.Context, key interface{}) *BasicCell {
	if v, ok := ctx.Prev(key).(*BasicCell); ok {
		return v
	}
	return &BasicCell{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *BasicCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
		s.WidthEqual(l.MaxGuide().Width())
	})

	chlds := []view.View{}

	iconView := imageview.NewImageView(ctx, 0)
	iconView.Image = v.Icon
	iconView.ResizeMode = imageview.ResizeModeFill
	chlds = append(chlds, iconView)

	iconGuide := l.Add(iconView, func(s *constraint.Solver) {
		s.WidthEqual(constraint.Const(20))
		s.HeightEqual(constraint.Const(20))
		s.LeftEqual(l.Left().Add(10))
		s.CenterYEqual(l.CenterY())
	})

	var subtitleGuide *constraint.Guide
	if len(v.Subtitle) > 0 {
		subtitleView := textview.New(ctx, 2)
		subtitleView.String = v.Subtitle
		subtitleView.Style.SetFont(text.Font{
			Family: "Helvetica Neue",
			Size:   22,
		})
		subtitleView.Style.SetTextColor(colornames.Gray)
		chlds = append(chlds, subtitleView)

		subtitleGuide = l.Add(subtitleView, func(s *constraint.Solver) {
			s.RightEqual(l.Right().Add(-10))
			s.LeftGreater(l.Left())
			s.CenterYEqual(l.CenterY())
		})
	}

	titleView := textview.New(ctx, 1)
	titleView.String = v.Title
	titleView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   22,
	})
	chlds = append(chlds, titleView)

	titleGuide := l.Add(titleView, func(s *constraint.Solver) {
		s.LeftEqual(iconGuide.Right().Add(10))
		s.CenterYEqual(l.CenterY())
		if subtitleGuide != nil {
			s.RightLess(subtitleGuide.Left())
		} else {
			s.RightLess(l.Right())
		}
	})
	_ = titleGuide

	return &view.Model{
		Children: chlds,
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.White},
	}
}
