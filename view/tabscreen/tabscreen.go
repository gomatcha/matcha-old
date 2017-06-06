package tabscreen

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout/constraint"
	tabnavpb "github.com/overcyn/mochi/pb/view/tabnav"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type Screen struct {
	store         store.Store
	screens       []view.Screen
	selectedIndex int
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
	return NewView(ctx, key, s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.store.Write()
	s.screens = ss
}

func (s *Screen) Children() []view.Screen {
	s.store.Read()
	return s.screens
}

func (s *Screen) SetSelectedIndex(idx int) {
	if idx != s.selectedIndex {
		s.store.Write()
		s.selectedIndex = idx
	}
}

func (s *Screen) SelectedIndex() int {
	s.store.Read()
	return s.selectedIndex
}

type View struct {
	*view.Embed
	screen *Screen
}

func NewView(ctx *view.Context, key interface{}, s *Screen) *View {
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
	chlds := []view.View{}
	for idx, i := range v.screen.Children() {
		chld := i.NewView(ctx, idx)
		screenspb = append(screenspb, &tabnavpb.Screen{
			Id:    int64(chld.Id()),
			Title: "Tab Title",
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

// type Screen struct {
// 	store        store.Store3
// 	view         view.View
// 	title        string
// 	icon         image.Image
// 	selectedIcon image.Image
// 	badge        string
// }

// func (tab *Screen) MarshalProtobuf() (*tabnavpb.Screen, error) {
// 	return &tabnavpb.Screen{
// 		Id:           int64(tab.view.Id()),
// 		Title:        tab.title,
// 		Icon:         pb.ImageEncode(tab.icon),
// 		SelectedIcon: pb.ImageEncode(tab.selectedIcon),
// 		Badge:        tab.badge,
// 	}, nil
// }
