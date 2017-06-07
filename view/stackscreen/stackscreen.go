package stackscreen

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb/view/stacknav"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type Screen struct {
	store   store.Store
	screens []view.Screen
}

func (s *Screen) Store() *store.Store {
	return &s.store
}

func (s *Screen) Lock() {
	s.store.Lock()
}

func (s *Screen) Unlock() {
	s.store.Unlock()
}

func (s *Screen) NewView(ctx *view.Context, key interface{}) view.View {
	return New(ctx, key, s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.store.Write()

	s.screens = ss
}

func (s *Screen) Children() []view.Screen {
	return s.screens
}

func (s *Screen) Push(vs view.Screen) {
	s.store.Write()

	s.screens = append(s.screens, vs)
}

func (s *Screen) Pop() {
	s.store.Write()

	if len(s.screens) > 0 {
		s.screens = s.screens[:len(s.screens)-1]
	}
}

type View struct {
	*view.Embed
	screen *Screen
}

func New(ctx *view.Context, key interface{}, s *Screen) *View {
	if v, ok := ctx.Prev(key).(*View); ok && v.screen == s {
		return v
	}

	embed := view.NewEmbed(ctx.NewId(key))
	embed.Subscribe(&s.store)
	return &View{
		Embed:  embed,
		screen: s,
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	v.screen.Lock()
	defer v.screen.Unlock()

	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		view.MainMu.Lock()
		defer view.MainMu.Unlock()

		pbevent := &stacknav.StackEvent{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		// Don't update the view for this
		v.Embed.Unsubscribe(&v.screen.store)

		v.screen.Lock()
		chl := v.screen.Children()[:len(pbevent.Id)]
		v.screen.SetChildren(chl...)
		v.screen.Unlock()

		v.Embed.Subscribe(&v.screen.store)
	}

	screenspb := []*stacknav.Screen{}
	chlds := []view.View{}
	for idx, i := range v.screen.Children() {
		chld := i.NewView(ctx, idx)

		var options *Options
		if optionsView, ok := chld.(*optionsView); ok {
			options = optionsView.options
		} else {
			options = &Options{
				Title: "Stack Title",
			}
		}

		screenspb = append(screenspb, &stacknav.Screen{
			Id:    int64(chld.Id()),
			Title: options.Title,
			CustomBackButtonTitle: len(options.BackButtonTitle) > 0,
			BackButtonTitle:       options.BackButtonTitle,
			BackButtonHidden:      options.BackButtonHidden,
		})

		chlds = append(chlds, chld)
		l.Add(chld, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height())
		})
	}

	return &view.Model{
		Children:       chlds,
		Layouter:       l,
		NativeViewName: "github.com/overcyn/mochi/view/stacknav",
		NativeViewState: &stacknav.StackNav{
			Screens:   screenspb,
			EventFunc: funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}

// TODO(KD): add middleware to read nativeValues{Key:Options} from view.Model
type key struct{}

var Key = key{}

type Options struct {
	Title            string
	BackButtonTitle  string
	BackButtonHidden bool
	// TitleView        view.View
	// RightViews       []view.View
	// LeftViews        []view.View
	// BarHidden        bool
	// Bar height?
}

func WithOptions(s view.Screen, opt *Options) view.Screen {
	return &optionsScreen{
		Screen:  s,
		options: opt,
	}
}

type optionsScreen struct {
	view.Screen
	options *Options
}

func (s *optionsScreen) NewView(ctx *view.Context, key interface{}) view.View {
	return &optionsView{
		View:    s.Screen.NewView(ctx, key),
		options: s.options,
	}
}

type optionsView struct {
	view.View
	options *Options
}
