package tabscreen

import (
	"fmt"
	"image"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha/env"
	"gomatcha.io/matcha/layout/constraint"
	tabnavpb "gomatcha.io/matcha/pb/view/tabscreen"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/view"
)

type Screen struct {
	store.Storer
	store         *store.Store
	screens       []view.Screen
	selectedIndex int
}

func New() *Screen {
	st := &store.Store{}
	return &Screen{Storer: st, store: st}
}

func (s *Screen) View(ctx *view.Context) view.View {
	return newView(ctx, "", s)
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

type tabView struct {
	*view.Embed
	screen   *Screen
	children []view.View
}

func newView(ctx *view.Context, key string, s *Screen) *tabView {
	if v, ok := ctx.Prev(key).(*tabView); ok && v.screen == s {
		return v
	}

	embed := view.NewEmbed(ctx.NewId(key))
	embed.Subscribe(s)
	return &tabView{
		Embed:  embed,
		screen: s,
	}
}

func (v *tabView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	v.screen.Lock()
	defer v.screen.Unlock()

	// Unsubscribe from old views
	for _, i := range v.children {
		v.Unsubscribe(i)
	}

	v.children = []view.View{}
	screenspb := []*tabnavpb.ChildView{}
	for idx, i := range v.screen.Children() {
		chld := i.View(ctx.WithPrefix(strconv.Itoa(idx)))

		var button *Button
		if childView, ok := chld.(ChildView); ok {
			button = childView.TabButton(ctx)
		} else {
			button = &Button{
				Title: "Title",
			}
		}

		v.Subscribe(chld)
		v.children = append(v.children, chld)
		screenspb = append(screenspb, &tabnavpb.ChildView{
			Id:           int64(chld.Id()),
			Title:        button.Title,
			Icon:         env.ImageMarshalProtobuf(button.Icon),
			SelectedIcon: env.ImageMarshalProtobuf(button.SelectedIcon),
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
		NativeViewName: "gomatcha.io/matcha/view/tabscreen",
		NativeViewState: &tabnavpb.View{
			Screens:       screenspb,
			SelectedIndex: int64(v.screen.SelectedIndex()),
		},
		NativeFuncs: map[string]interface{}{
			"OnSelect": func(data []byte) {
				pbevent := &tabnavpb.Event{}
				err := proto.Unmarshal(data, pbevent)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				v.screen.Lock()
				defer v.screen.Unlock()
				v.screen.SetSelectedIndex(int(pbevent.SelectedIndex))
			},
		},
	}
}

type ChildView interface {
	view.View
	TabButton(*view.Context) *Button
}

type Button struct {
	Title        string
	Icon         image.Image
	SelectedIcon image.Image
	Badge        string
}

func WithButton(s view.Screen, button *Button) view.Screen {
	return &screenWrapper{
		Screen: s,
		button: button,
	}
}

type screenWrapper struct {
	view.Screen
	button *Button
}

func (s *screenWrapper) View(ctx *view.Context) view.View {
	return &viewWrapper{
		View:   s.Screen.View(ctx),
		button: s.button,
	}
}

type viewWrapper struct {
	view.View
	button *Button
}

func (v *viewWrapper) TabButton(*view.Context) *Button {
	return v.button
}
