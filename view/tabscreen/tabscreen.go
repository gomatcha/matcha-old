package tabscreen

import (
	"fmt"
	"image"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb"
	tabnavpb "github.com/overcyn/mochi/pb/view/tabnav"
	"github.com/overcyn/mochi/view"
)

type Screen struct {
	comm.Storer
	store         *comm.AsyncStore
	screens       []view.Screen
	selectedIndex int
}

func NewScreen() *Screen {
	st := &comm.AsyncStore{}
	return &Screen{Storer: st, store: st}
}

func (s *Screen) NewView(ctx *view.Context, key string) view.View {
	return NewView(ctx, key, s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.store.Update()
	s.screens = ss
}

func (s *Screen) Children() []view.Screen {
	return s.screens
}

func (s *Screen) SetSelectedIndex(idx int) {
	if idx != s.selectedIndex {
		s.store.Update()
		s.selectedIndex = idx
	}
}

func (s *Screen) SelectedIndex() int {
	return s.selectedIndex
}

type View struct {
	*view.Embed
	screen *Screen
}

func NewView(ctx *view.Context, key string, s *Screen) *View {
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
		pbevent := &tabnavpb.Event{}
		err := proto.Unmarshal(data, pbevent)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		v.screen.Lock()
		defer v.screen.Unlock()
		v.screen.SetSelectedIndex(int(pbevent.SelectedIndex))
	}

	screenspb := []*tabnavpb.Screen{}
	for idx, i := range v.screen.Children() {
		chld := i.View(ctx, strconv.Itoa(idx))

		var button *TabButton
		if childView, ok := chld.(ChildView); ok {
			button = childView.TabButton(ctx)
		} else {
			button = &TabButton{
				Title: "Title",
			}
		}

		screenspb = append(screenspb, &tabnavpb.Screen{
			Id:           int64(chld.Id()),
			Title:        button.Title,
			Icon:         pb.ImageEncode(button.Icon),
			SelectedIcon: pb.ImageEncode(button.SelectedIcon),
			Badge:        button.Badge,
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
		NativeViewName: "github.com/overcyn/mochi/view/tabscreen",
		NativeViewState: &tabnavpb.TabNav{
			Screens:       screenspb,
			SelectedIndex: int64(v.screen.SelectedIndex()),
			EventFunc:     funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}

type ChildView interface {
	view.View
	TabButton(*view.Context) *TabButton
}

type TabButton struct {
	Title        string
	Icon         image.Image
	SelectedIcon image.Image
	Badge        string
}

func WithTabButton(s view.Screen, button *TabButton) view.Screen {
	return &tabButtonScreen{
		Screen: s,
		button: button,
	}
}

type tabButtonScreen struct {
	view.Screen
	button *TabButton
}

func (s *tabButtonScreen) View(ctx *view.Context, key string) view.View {
	return &tabButtonView{
		View:   s.Screen.View(ctx, key),
		button: s.button,
	}
}

type tabButtonView struct {
	view.View
	button *TabButton
}

func (v *tabButtonView) TabButton(*view.Context) *TabButton {
	return v.button
}
