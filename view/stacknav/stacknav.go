package stacknav

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
	s.store.Read()
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

		v.screen.Lock()
		defer v.screen.Unlock()
		chl := v.screen.Children()[:len(pbevent.Id)]
		v.screen.SetChildren(chl...)
	}

	screenspb := []*stacknav.Screen{}
	chlds := []view.View{}
	for idx, i := range v.screen.Children() {
		chld := i.NewView(ctx, idx)
		screenspb = append(screenspb, &stacknav.Screen{
			Id:    int64(chld.Id()),
			Title: "Stack Title",
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

// type Screen struct {
// 	store            store.Store
// 	view             view.View
// 	title            string
// 	backButtonTitle  string
// 	backButtonHidden bool
// 	titleView        view.View
// 	rightViews       []view.View
// 	leftViews        []view.View
// 	// BarHidden        bool
// 	// Bar height?
// }

// func (s *Screen) MarshalProtobuf() (*stacknav.Screen, error) {
// return &stacknav.Screen{
// 	Id:    int64(s.view.Id()),
// 	Title: s.title,
// 	CustomBackButtonTitle: len(s.backButtonTitle) > 0,
// 	BackButtonTitle:       s.backButtonTitle,
// 	BackButtonHidden:      s.backButtonHidden,
// 	// TitleViewId:      s.titleView.Id(), // TODO(KD):
// 	// RightViewIds
// 	// LeftViewIds:..
// }, nil
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

// func (opt *Screen) SetBackButtonTitle(v string) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.backButtonTitle = v
// }

// func (opt *Screen) BackButtonTitle() string {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.backButtonTitle
// }

// func (opt *Screen) SetBackButtonHidden(v bool) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.backButtonHidden = v
// }

// func (opt *Screen) BackButtonHidden() bool {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.backButtonHidden
// }

// func (opt *Screen) SetTitleView(v view.View) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.titleView = v
// }

// func (opt *Screen) TitleView() view.View {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.titleView
// }

// func (opt *Screen) SetRightViews(v []view.View) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.rightViews = v
// }

// func (opt *Screen) RightViews() []view.View {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.rightViews
// }

// func (opt *Screen) SetLeftViews(v []view.View) {
// 	tx := store.NewWriteTx()
// 	defer tx.Commit()

// 	opt.store.Write(tx)
// 	opt.leftViews = v
// }

// func (opt *Screen) LeftViews() []view.View {
// 	tx := store.NewReadTx()
// 	defer tx.Commit()

// 	opt.store.Read(tx)
// 	return opt.leftViews
// }

// func (opt *Screen) Notify() chan struct{} {
// 	return opt.store.Notify()
// }

// func (opt *Screen) Unnotify(c chan struct{}) {
// 	opt.store.Unnotify(c)
// }
