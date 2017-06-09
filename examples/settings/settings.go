package settings

import (
	"image"
	"image/color"

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
	"github.com/overcyn/mochi/view/switchview"
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
		spacer := NewSpacer(ctx, "spacer1")
		chlds = append(chlds, spacer)
		l.Add(spacer)
	}
	{
		switchView := switchview.New(ctx, 6)

		group := []view.View{}
		cell1 := NewBasicCell(ctx, 0)
		cell1.Title = "Airplane Mode"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, 1)
		cell2.Title = "Wi-Fi"
		cell2.Subtitle = "Home Wifi"
		group = append(group, cell2)

		cell3 := NewBasicCell(ctx, 2)
		cell3.Title = "Bluetooth"
		cell3.Subtitle = "On"
		group = append(group, cell3)

		cell4 := NewBasicCell(ctx, 3)
		cell4.Title = "Cellular"
		group = append(group, cell4)

		cell5 := NewBasicCell(ctx, 4)
		cell5.Title = "Personal Hotspot"
		cell5.Subtitle = "Off"
		group = append(group, cell5)

		cell6 := NewBasicCell(ctx, 5)
		cell6.Title = "Carrier"
		cell6.Subtitle = "T-Mobile"
		group = append(group, cell6)

		for _, i := range AddSeparators(ctx, "a", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}
	{
		spacer := NewSpacer(ctx, "spacer2")
		chlds = append(chlds, spacer)
		l.Add(spacer)
	}
	{
		group := []view.View{}
		cell1 := NewBasicCell(ctx, 10)
		cell1.Title = "Notifications"
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, 11)
		cell2.Title = "Control Center"
		group = append(group, cell2)

		cell3 := NewBasicCell(ctx, 12)
		cell3.Title = "Do Not Disturb"
		group = append(group, cell3)

		for _, i := range AddSeparators(ctx, "b", group) {
			chlds = append(chlds, i)
			l.Add(i)
		}
	}

	scrollChild := basicview.New(ctx, 6)
	scrollChild.Layouter = l
	scrollChild.Children = chlds

	scrollView := scrollview.New(ctx, 5)
	scrollView.ContentView = scrollChild

	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

var (
	cellColor       = color.Gray{255}
	separatorColor  = color.Gray{200}
	backgroundColor = color.Gray{239}
	subtitleColor   = color.Gray{142}
	titleColor      = color.Gray{0}
)

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
			sep.LeftPadding = 60
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
	chl.Painter = &paint.Style{BackgroundColor: separatorColor}
	l.Add(chl, func(s *constraint.Solver) {
		s.HeightEqual(l.Height())
		s.LeftEqual(l.Left().Add(v.LeftPadding))
		s.RightEqual(l.Right())
	})

	return &view.Model{
		Children: []view.View{chl},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: cellColor},
	}
}

type Spacer struct {
	*view.Embed
	Height float64
}

func NewSpacer(ctx *view.Context, key interface{}) *Spacer {
	if v, ok := ctx.Prev(key).(*Spacer); ok {
		return v
	}
	return &Spacer{
		Embed:  view.NewEmbed(ctx.NewId(key)),
		Height: 35,
	}
}

func (v *Spacer) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(35))
		s.WidthEqual(l.MaxGuide().Width())
	})

	return &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

type BasicCell struct {
	*view.Embed
	Icon          image.Image
	Title         string
	Subtitle      string
	AccessoryView view.View
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
		s.HeightEqual(constraint.Const(44))
		s.WidthEqual(l.MaxGuide().Width())
	})

	chlds := []view.View{}

	iconView := imageview.New(ctx, 0)
	iconView.Image = v.Icon
	iconView.ResizeMode = imageview.ResizeModeFill
	iconView.Painter = &paint.Style{
		BackgroundColor: colornames.Lightgray,
		CornerRadius:    5,
	}
	chlds = append(chlds, iconView)

	iconGuide := l.Add(iconView, func(s *constraint.Solver) {
		s.WidthEqual(constraint.Const(30))
		s.HeightEqual(constraint.Const(30))
		s.LeftEqual(l.Left().Add(15))
		s.CenterYEqual(l.CenterY())
	})

	rightAnchor := l.Right()
	if v.AccessoryView != nil {
		chlds = append(chlds, v.AccessoryView)
		accessoryGuide := l.Add(v.AccessoryView, func(s *constraint.Solver) {
			s.RightEqual(rightAnchor.Add(-10))
			s.LeftGreater(l.Left())
			s.CenterYEqual(l.CenterY())
		})
		rightAnchor = accessoryGuide.Left()
	}

	if len(v.Subtitle) > 0 {
		subtitleView := textview.New(ctx, 2)
		subtitleView.String = v.Subtitle
		subtitleView.Style.SetFont(text.Font{
			Family: "Helvetica Neue",
			Size:   13,
		})
		subtitleView.Style.SetTextColor(subtitleColor)
		chlds = append(chlds, subtitleView)

		subtitleGuide := l.Add(subtitleView, func(s *constraint.Solver) {
			s.RightEqual(rightAnchor.Add(-10))
			s.LeftGreater(l.Left())
			s.CenterYEqual(l.CenterY())
		})
		rightAnchor = subtitleGuide.Left()
	}

	titleView := textview.New(ctx, 1)
	titleView.String = v.Title
	titleView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   13,
	})
	titleView.Style.SetTextColor(titleColor)
	chlds = append(chlds, titleView)

	titleGuide := l.Add(titleView, func(s *constraint.Solver) {
		s.LeftEqual(iconGuide.Right().Add(15))
		s.CenterYEqual(l.CenterY())
		s.RightLess(rightAnchor.Add(-10))
	})
	_ = titleGuide

	return &view.Model{
		Children: chlds,
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: cellColor},
	}
}
