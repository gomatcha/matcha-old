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
	store.Node
	selectedIndex int
	ids           []int64
	children      map[int64]view.Screen
	maxId         int64
}

func (s *Screen) View(ctx *view.Context) view.View {
	return newView(ctx, "", s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.Signal()

	s.children = map[int64]view.Screen{}
	s.ids = []int64{}
	for _, i := range ss {
		s.maxId += 1
		s.ids = append(s.ids, s.maxId)
		s.children[s.maxId] = i
	}
}

func (s *Screen) Children() []view.Screen {
	children := []view.Screen{}
	for _, i := range s.ids {
		children = append(children, s.children[i])
	}
	return children
}

func (s *Screen) SetSelectedIndex(idx int) {
	if idx != s.selectedIndex {
		s.Signal()
		s.selectedIndex = idx
	}
}

func (s *Screen) SelectedIndex() int {
	return s.selectedIndex
}

type tabView struct {
	view.Embed
	screen   *Screen
	children map[int64]view.View
	ids      []int64
}

func newView(ctx *view.Context, key string, s *Screen) *tabView {
	if v, ok := ctx.Prev(key).(*tabView); ok && v.screen == s {
		return v
	}

	v := &tabView{
		Embed:  ctx.NewEmbed(key),
		screen: s,
	}
	v.Subscribe(s)
	return v
}

func (v *tabView) Build(ctx *view.Context) view.Model {
	v.screen.Lock()
	defer v.screen.Unlock()

	l := constraint.New()

	children := map[int64]view.View{}
	childrenPb := []*tabnavpb.ChildView{}
	v.ids = append([]int64(nil), v.screen.ids...)
	for _, i := range v.ids {
		key := strconv.Itoa(int(i))

		// Create the child if necessary and subscribe to it.
		chld, ok := v.children[i]
		if !ok {
			chld = v.screen.children[i].View(ctx.WithPrefix("view" + key))
			children[i] = chld
			v.Subscribe(chld)
		} else {
			children[i] = chld
			delete(v.children, i)
		}

		// Create the button
		var button *Button
		if childView, ok := chld.(ChildView); ok {
			button = childView.TabButton(ctx)
		} else {
			button = &Button{
				Title: "Title",
			}
		}

		// Add the child.
		l.Add(chld, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height())
		})

		// Add to protobuf.
		childrenPb = append(childrenPb, &tabnavpb.ChildView{
			Id:           int64(chld.Id()),
			Title:        button.Title,
			Icon:         env.ImageMarshalProtobuf(button.Icon),
			SelectedIcon: env.ImageMarshalProtobuf(button.SelectedIcon),
			Badge:        button.Badge,
		})
	}

	// Unsubscribe from old views
	for _, chld := range v.children {
		v.Unsubscribe(chld)
	}
	v.children = children

	return view.Model{
		Children:       l.Views(),
		Layouter:       l,
		NativeViewName: "gomatcha.io/matcha/view/tabscreen",
		NativeViewState: &tabnavpb.View{
			Screens:       childrenPb,
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
