package tabnavigator

import (
	"image"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb"
	tabnavigatorpb "github.com/overcyn/mochi/pb/view/tabnavigator"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type TabNavigator struct {
	*view.Embed
	tabs     []Tab
	notifier *mochi.BatchNotifier
}

func New(ctx *view.Context, key interface{}) *TabNavigator {
	if v, ok := ctx.Prev(key).(*TabNavigator); ok {
		return v
	}
	return &TabNavigator{
		Embed:    view.NewEmbed(ctx.NewId(key)),
		notifier: &mochi.BatchNotifier{},
	}
}

func (n *TabNavigator) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	tabspb := []*tabnavigatorpb.Tab{}
	views := []view.View{}
	for _, i := range n.tabs {
		tabpb, err := i.MarshalProtobuf()
		if err == nil {
			tabspb = append(tabspb, tabpb)
		}

		views = append(views, i.View)
		l.Add(i.View, func(s *constraint.Solver) {
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
		NativeViewName: "github.com/overcyn/mochi/view/tabnavigator",
		NativeViewState: &tabnavigatorpb.TabNavigator{
			Tabs: tabspb,
		},
	}
}

func (n *TabNavigator) Tabs() []Tab {
	return n.tabs
}

func (n *TabNavigator) SetTabs(tabs []Tab) {
	// // unsubscribe from old views
	// for _, i := range n.tabs {
	// 	n.notifier.Unsubscribe(i.Options)
	// }

	// // subscribe to new views
	// for _, i := range tabs {
	// 	n.notifier.Subscribe(i.Options)
	// }
	n.tabs = tabs
}

type Tab struct {
	View    view.View
	Options *Options
}

func (tab *Tab) MarshalProtobuf() (*tabnavigatorpb.Tab, error) {
	return &tabnavigatorpb.Tab{
		Id:           int64(tab.View.Id()),
		Title:        tab.Options.title,
		Icon:         pb.ImageEncode(tab.Options.icon),
		SelectedIcon: pb.ImageEncode(tab.Options.selectedIcon),
		Badge:        tab.Options.badge,
	}, nil
}

type Options struct {
	store        store.Store
	title        string
	icon         image.Image
	selectedIcon image.Image
	badge        string
}

func (opt *Options) SetTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.title = v
}

func (opt *Options) Title() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.title
}

func (opt *Options) SetIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.icon = v
}

func (opt *Options) Icon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.icon
}

func (opt *Options) SetSelectedIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.selectedIcon = v
}

func (opt *Options) SelectedIcon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.selectedIcon
}

func (opt *Options) SetBadge(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.badge = v
}

func (opt *Options) Badge() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.badge
}

func (opt *Options) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *Options) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}
