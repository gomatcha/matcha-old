package tabnav

import (
	"image"

	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb"
	tabnavpb "github.com/overcyn/mochi/pb/view/tabnav"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type TabNav struct {
	*view.Embed
	screens []*Screen
	// notifier *mochi.BatchNotifier
}

func New(ctx *view.Context, key interface{}) *TabNav {
	if v, ok := ctx.Prev(key).(*TabNav); ok {
		return v
	}
	return &TabNav{
		Embed: view.NewEmbed(ctx.NewId(key)),
		// notifier: &mochi.BatchNotifier{},
	}
}

func (n *TabNav) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	tabspb := []*tabnavpb.Screen{}
	views := []view.View{}
	for _, i := range n.screens {
		tabpb, err := i.MarshalProtobuf()
		if err == nil {
			tabspb = append(tabspb, tabpb)
		}

		views = append(views, i.View())
		l.Add(i.View(), func(s *constraint.Solver) {
			s.TopEqual(constraint.Const(0))
			s.LeftEqual(constraint.Const(0))
			s.WidthEqual(l.MaxGuide().Width())
			s.HeightEqual(l.MaxGuide().Height())
		})
	}

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})

	return &view.Model{
		Children:       views,
		Layouter:       l,
		NativeViewName: "github.com/overcyn/mochi/view/tabnav",
		NativeViewState: &tabnavpb.TabNav{
			Screens: tabspb,
		},
	}
}

func (n *TabNav) Screens() []*Screen {
	return n.screens
}

func (n *TabNav) SetScreens(ss []*Screen) {
	// // unsubscribe from old views
	// for _, i := range n.tabs {
	// 	n.notifier.Unsubscribe(i.Options)
	// }

	// // subscribe to new views
	// for _, i := range tabs {
	// 	n.notifier.Subscribe(i.Options)
	// }
	n.screens = ss
}

type Screen struct {
	store        store.Store
	view         view.View
	title        string
	icon         image.Image
	selectedIcon image.Image
	badge        string
}

func (tab *Screen) MarshalProtobuf() (*tabnavpb.Screen, error) {
	return &tabnavpb.Screen{
		Id:           int64(tab.view.Id()),
		Title:        tab.title,
		Icon:         pb.ImageEncode(tab.icon),
		SelectedIcon: pb.ImageEncode(tab.selectedIcon),
		Badge:        tab.badge,
	}, nil
}

func (opt *Screen) SetView(v view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.view = v
}

func (opt *Screen) View() view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.view
}

func (opt *Screen) SetTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.title = v
}

func (opt *Screen) Title() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.title
}

func (opt *Screen) SetIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.icon = v
}

func (opt *Screen) Icon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.icon
}

func (opt *Screen) SetSelectedIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.selectedIcon = v
}

func (opt *Screen) SelectedIcon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.selectedIcon
}

func (opt *Screen) SetBadge(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.badge = v
}

func (opt *Screen) Badge() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.badge
}

func (opt *Screen) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *Screen) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}
