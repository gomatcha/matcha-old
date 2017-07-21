package stackscreen

import (
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/pb/view/stacknav"
	"gomatcha.io/matcha/store"
	"gomatcha.io/matcha/view"
)

type Screen struct {
	store.Node
	ids      []int64
	children map[int64]view.Screen
	maxId    int64
}

func (s *Screen) View(ctx *view.Context) view.View {
	s.Lock()
	defer s.Unlock()
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

func (s *Screen) setChildIds(ids []int64) {
	s.Signal()

	prevChildren := s.children
	s.children = map[int64]view.Screen{}
	s.ids = []int64{}
	for _, i := range ids {
		if child, ok := prevChildren[i]; ok {
			s.children[i] = child
			s.ids = append(s.ids, i)
		}
	}
}

func (s *Screen) Children() []view.Screen {
	children := []view.Screen{}
	for _, i := range s.ids {
		children = append(children, s.children[i])
	}
	return children
}

func (s *Screen) Push(vs view.Screen) {
	s.Signal()

	if s.children == nil {
		s.children = map[int64]view.Screen{}
	}
	s.maxId += 1
	s.children[s.maxId] = vs
	s.ids = append(s.ids, s.maxId)
}

func (s *Screen) Pop() {
	s.Signal()

	if len(s.ids) > 0 {
		id := s.ids[len(s.ids)-1]
		s.ids = s.ids[:len(s.ids)-1]
		delete(s.children, id)
	}
}

type stackView struct {
	view.Embed
	screen   *Screen
	children map[int64]view.View
	ids      []int64
}

func newView(ctx *view.Context, key string, s *Screen) *stackView {
	if v, ok := ctx.Prev(key).(*stackView); ok && v.screen == s {
		return v
	}

	v := &stackView{
		Embed:  ctx.NewEmbed(key),
		screen: s,
	}
	v.Subscribe(s)
	return v
}

func (v *stackView) Build(ctx *view.Context) view.Model {
	v.screen.Lock()
	defer v.screen.Unlock()

	l := constraint.New()

	children := map[int64]view.View{}
	childrenPb := []*stacknav.ChildView{}
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

		// Create the bar.
		var bar *Bar
		if childView, ok := chld.(ChildView); ok {
			bar = childView.StackBar(ctx.WithPrefix("bar" + key))
		} else {
			bar = &Bar{
				Title: "Title",
			}
		}

		// Add the bar.
		barV := &barView{
			Embed: view.NewEmbed(ctx.NewId(key)),
			bar:   bar,
		}
		l.Add(barV, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(constraint.Const(44))
		})

		// Add the child.
		l.Add(chld, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height().Add(-64)) // TODO(KD): Respect bar actual height, shorter when rotated, etc...
		})

		// Add ids to protobuf.
		childrenPb = append(childrenPb, &stacknav.ChildView{
			ViewId:   int64(chld.Id()),
			BarId:    int64(barV.Id()),
			ScreenId: i,
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
		NativeViewName: "gomatcha.io/matcha/view/stacknav",
		NativeViewState: &stacknav.View{
			Children: childrenPb,
		},
		NativeFuncs: map[string]interface{}{
			"OnChange": func(data []byte) {
				pbevent := &stacknav.StackEvent{}
				err := proto.Unmarshal(data, pbevent)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				v.screen.Lock()
				v.screen.setChildIds(pbevent.Id)
				v.screen.Unlock()
			},
		},
	}
}

type ChildView interface {
	view.View
	StackBar(*view.Context) *Bar // TODO(KD): Doesn't this make it harder to wrap??
}

type barView struct {
	view.Embed
	bar *Bar
}

func (v *barView) Build(ctx *view.Context) view.Model {
	l := constraint.New()

	// iOS does the layouting for us. We just need the correct sizes.
	titleViewId := int64(0)
	if v.bar.TitleView != nil {
		titleViewId = int64(v.bar.TitleView.Id())
		l.Add(v.bar.TitleView, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.HeightLess(l.MaxGuide().Height())
			s.WidthLess(l.MaxGuide().Width())
		})
	}

	rightViewIds := []int64{}
	for _, i := range v.bar.RightViews {
		rightViewIds = append(rightViewIds, int64(i.Id()))
		l.Add(i, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.HeightLess(l.MaxGuide().Height())
			s.WidthLess(l.MaxGuide().Width())
		})
	}
	leftViewIds := []int64{}
	for _, i := range v.bar.LeftViews {
		leftViewIds = append(leftViewIds, int64(i.Id()))
		l.Add(i, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.HeightLess(l.MaxGuide().Height())
			s.WidthLess(l.MaxGuide().Width())
		})
	}

	return view.Model{
		Layouter:       l,
		Children:       l.Views(),
		NativeViewName: "gomatcha.io/matcha/view/stacknav Bar",
		NativeViewState: &stacknav.Bar{
			Title: v.bar.Title,
			CustomBackButtonTitle: len(v.bar.BackButtonTitle) > 0,
			BackButtonTitle:       v.bar.BackButtonTitle,
			BackButtonHidden:      v.bar.BackButtonHidden,
			TitleViewId:           titleViewId,
			RightViewIds:          rightViewIds,
			LeftViewIds:           leftViewIds,
		},
	}
}

type Bar struct {
	Title            string
	BackButtonTitle  string
	BackButtonHidden bool

	TitleView  view.View
	RightViews []view.View
	LeftViews  []view.View
}

func WithBar(s view.Screen, bar *Bar) view.Screen {
	return &screenWrapper{
		Screen:   s,
		stackBar: bar,
	}
}

type screenWrapper struct {
	view.Screen
	stackBar *Bar
}

func (s *screenWrapper) View(ctx *view.Context) view.View {
	return &viewWrapper{
		View:     s.Screen.View(ctx),
		stackBar: s.stackBar,
	}
}

type viewWrapper struct {
	view.View
	stackBar *Bar
}

func (s *viewWrapper) StackBar(*view.Context) *Bar {
	return s.stackBar
}
