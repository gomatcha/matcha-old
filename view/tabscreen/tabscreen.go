package tabscreen

import (
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
	s.store.Write()
	s.selectedIndex = idx
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

	// funcId := ctx.NewFuncId()
	// f := func(data []byte) {
	// 	view.MainMu.Lock()
	// 	defer view.MainMu.Unlock()

	// 	pbevent := &stacknav.StackEvent{}
	// 	err := proto.Unmarshal(data, pbevent)
	// 	if err != nil {
	// 		fmt.Println("error", err)
	// 		return
	// 	}

	// 	v.screen.Lock()
	// 	defer v.screen.Unlock()
	// 	chl := v.screen.Children()[:len(pbevent.Id)]
	// 	v.screen.SetChildren(chl...)
	// }

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
			// EventFunc: funcId,
		},
		// NativeFuncs: map[int64]interface{}{
		// 	funcId: f,
		// },
	}
}

// type TabNav struct {
// 	*view.Embed
// 	screens []*Screen
// 	// notifier *mochi.BatchNotifier
// }

// func New(ctx *view.Context, key interface{}) *TabNav {
// 	if v, ok := ctx.Prev(key).(*TabNav); ok {
// 		return v
// 	}
// 	return &TabNav{
// 		Embed: view.NewEmbed(ctx.NewId(key)),
// 		// notifier: &mochi.BatchNotifier{},
// 	}
// }

// func (n *TabNav) Build(ctx *view.Context) *view.Model {
// 	l := constraint.New()

// 	tabspb := []*tabnavpb.Screen{}
// 	views := []view.View{}
// 	for _, i := range n.screens {
// 		tabpb, err := i.MarshalProtobuf()
// 		if err == nil {
// 			tabspb = append(tabspb, tabpb)
// 		}

// 		views = append(views, i.View())
// 		l.Add(i.View(), func(s *constraint.Solver) {
// 			s.TopEqual(constraint.Const(0))
// 			s.LeftEqual(constraint.Const(0))
// 			s.WidthEqual(l.MaxGuide().Width())
// 			s.HeightEqual(l.MaxGuide().Height())
// 		})
// 	}

// 	l.Solve(func(s *constraint.Solver) {
// 		s.WidthEqual(l.MaxGuide().Width())
// 		s.HeightEqual(l.MaxGuide().Height())
// 	})

// 	return &view.Model{
// 		Children:       views,
// 		Layouter:       l,
// 		NativeViewName: "github.com/overcyn/mochi/view/tabnav",
// 		NativeViewState: &tabnavpb.TabNav{
// 			Screens: tabspb,
// 		},
// 	}
// }

// func (n *TabNav) Screens() []*Screen {
// 	return n.screens
// }

// func (n *TabNav) SetScreens(ss []*Screen) {
// 	// // unsubscribe from old views
// 	// for _, i := range n.tabs {
// 	// 	n.notifier.Unsubscribe(i.Options)
// 	// }

// 	// // subscribe to new views
// 	// for _, i := range tabs {
// 	// 	n.notifier.Subscribe(i.Options)
// 	// }
// 	n.screens = ss
// }

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

// func (opt *Screen) SetView(v view.View) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.view = v
// }

// func (opt *Screen) View() view.View {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.view
// }

// func (opt *Screen) SetTitle(v string) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.title = v
// }

// func (opt *Screen) Title() string {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.title
// }

// func (opt *Screen) SetIcon(v image.Image) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.icon = v
// }

// func (opt *Screen) Icon() image.Image {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.icon
// }

// func (opt *Screen) SetSelectedIcon(v image.Image) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.selectedIcon = v
// }

// func (opt *Screen) SelectedIcon() image.Image {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.selectedIcon
// }

// func (opt *Screen) SetBadge(v string) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.badge = v
// }

// func (opt *Screen) Badge() string {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.badge
// }

// func (opt *Screen) Notify() chan struct{} {
// 	return opt.store.Notify()
// }

// func (opt *Screen) Unnotify(c chan struct{}) {
// 	opt.store.Unnotify(c)
// }
