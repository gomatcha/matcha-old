package stackscreen

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb/view/stacknav"
	"github.com/overcyn/mochi/view"
)

type Screen struct {
	comm.Storer
	store   *comm.AsyncStore
	screens []view.Screen
}

func NewScreen() *Screen {
	st := &comm.AsyncStore{}
	return &Screen{
		Storer: st,
		store:  st,
	}
}

func (s *Screen) NewView(ctx *view.Context, key interface{}) view.View {
	return New(ctx, key, s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.store.Update()

	s.screens = ss
}

func (s *Screen) Children() []view.Screen {
	return s.screens
}

func (s *Screen) Push(vs view.Screen) {
	s.store.Update()

	s.screens = append(s.screens, vs)
}

func (s *Screen) Pop() {
	s.store.Update()

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
	embed.Subscribe(s)
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
		v.Embed.Unsubscribe(v.screen)

		v.screen.Lock()
		chl := v.screen.Children()[:len(pbevent.Id)]
		v.screen.SetChildren(chl...)
		v.screen.Unlock()

		v.Embed.Subscribe(v.screen)
	}

	screenspb := []*stacknav.Screen{}
	for idx, i := range v.screen.Children() {
		chld := i.NewView(ctx, idx)

		var bar *StackBar
		if childView, ok := chld.(ChildView); ok {
			bar = childView.StackBar(ctx)
		} else {
			bar = &StackBar{
				Title: "Title",
			}
		}

		screenspb = append(screenspb, &stacknav.Screen{
			Id:    int64(chld.Id()),
			Title: bar.Title,
			CustomBackButtonTitle: len(bar.BackButtonTitle) > 0,
			BackButtonTitle:       bar.BackButtonTitle,
			BackButtonHidden:      bar.BackButtonHidden,
		})

		l.Add(chld, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height())
		})
	}

	return &view.Model{
		Children:       l.Views(),
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

type ChildView interface {
	view.View
	StackBar(*view.Context) *StackBar
}

type StackBar struct {
	Title            string
	BackButtonTitle  string
	BackButtonHidden bool
	TitleView        view.View // TODO(KD):
	RightViews       []view.View
	LeftViews        []view.View
	BarHidden        bool
	// Bar height?
}

func WithStackBar(s view.Screen, bar *StackBar) view.Screen {
	return &stackScreen{
		Screen:   s,
		stackBar: bar,
	}
}

type stackScreen struct {
	view.Screen
	stackBar *StackBar
}

func (s *stackScreen) NewView(ctx *view.Context, key interface{}) view.View {
	return &stackView{
		View:     s.Screen.NewView(ctx, key),
		stackBar: s.stackBar,
	}
}

type stackView struct {
	view.View
	stackBar *StackBar
}

func (s *stackView) StackBar(*view.Context) *StackBar {
	return s.stackBar
}
