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
	screens []view.Screen
}

func New() *Screen {
	return &Screen{}
}

func (s *Screen) View(ctx *view.Context) view.View {
	return newView(ctx, "", s)
}

func (s *Screen) SetChildren(ss ...view.Screen) {
	s.Signal()

	s.screens = ss
}

func (s *Screen) Children() []view.Screen {
	return s.screens
}

func (s *Screen) Push(vs view.Screen) {
	s.Signal()

	s.screens = append(s.screens, vs)
}

func (s *Screen) Pop() {
	s.Signal()

	if len(s.screens) > 0 {
		s.screens = s.screens[:len(s.screens)-1]
	}
}

type stackView struct {
	*view.Embed
	screen   *Screen
	children []view.View
}

func newView(ctx *view.Context, key string, s *Screen) *stackView {
	if v, ok := ctx.Prev(key).(*stackView); ok && v.screen == s {
		return v
	}

	embed := view.NewEmbed(ctx.NewId(key))
	embed.Subscribe(s)
	return &stackView{
		Embed:  embed,
		screen: s,
	}
}

func (v *stackView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	v.screen.Lock()
	defer v.screen.Unlock()

	// Unsubscribe from old views
	for _, i := range v.children {
		v.Unsubscribe(i)
	}

	v.children = []view.View{}
	childrenPb := []*stacknav.ChildView{}
	for idx, i := range v.screen.Children() {
		ctx := ctx.WithPrefix(strconv.Itoa(idx))
		chld := i.View(ctx.WithPrefix("view"))

		var bar *Bar
		if childView, ok := chld.(ChildView); ok {
			bar = childView.StackBar(ctx.WithPrefix("bar"))
		} else {
			bar = &Bar{
				Title: "Title",
			}
		}

		barV := &barView{
			Embed: view.NewEmbed(ctx.NewId(strconv.Itoa(idx))),
			bar:   bar,
		}
		l.Add(barV, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(constraint.Const(44))
		})

		v.Subscribe(chld)
		v.children = append(v.children, chld)
		childrenPb = append(childrenPb, &stacknav.ChildView{
			ViewId: int64(chld.Id()),
			BarId:  int64(barV.Id()),
		})

		l.Add(chld, func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height().Add(-64)) // TODO(KD): Respect bar actual height, shorter when rotated, etc...
		})
	}

	return &view.Model{
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
				chl := v.screen.Children()[:len(pbevent.Id)]
				v.screen.SetChildren(chl...)
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
	*view.Embed
	bar *Bar
}

func (v *barView) Build(ctx *view.Context) *view.Model {
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

	return &view.Model{
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
