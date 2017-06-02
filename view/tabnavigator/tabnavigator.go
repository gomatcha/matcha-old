package tabnavigator

import (
	"image"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type TabNavigator struct {
	*view.Embed
	views    []view.View
	options  []*TabOptions
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

	for _, i := range n.views {
		l.Add(i, func(s *constraint.Solver) {
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
		Children:       n.views,
		Layouter:       l,
		NativeViewName: "github.com/overcyn/mochi/view/tabnavigator",
		// NativeViewState: ft.MarshalProtobuf(),
	}
}

func (n *TabNavigator) Views() []view.View {
	return n.views
}

func (n *TabNavigator) SetViews(vs []view.View) {
	// unsubscribe from old views
	for _, opt := range n.options {
		n.notifier.Unsubscribe(opt)
	}

	// subscribe to new views
	opts := []*TabOptions{}
	for _, i := range vs {
		var opt *TabOptions
		tabber, ok := i.(Tabber)
		if ok {
			opt = tabber.TabOptions()
		} else {
			opt = &TabOptions{}
		}
		opts = append(opts, opt)
		n.notifier.Subscribe(opt)
	}
	n.options = opts
	n.views = vs
}

type TabOptions struct {
	store store.Store

	title        string
	icon         image.Image
	selectedIcon image.Image
	badge        string
}

func (opt *TabOptions) SetTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.title = v
}

func (opt *TabOptions) Title() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.title
}

func (opt *TabOptions) SetIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.icon = v
}

func (opt *TabOptions) Icon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.icon
}

func (opt *TabOptions) SetSelectedIcon(v image.Image) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.selectedIcon = v
}

func (opt *TabOptions) SelectedIcon() image.Image {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.selectedIcon
}

func (opt *TabOptions) SetBadge(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.badge = v
}

func (opt *TabOptions) Badge() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.badge
}

func (opt *TabOptions) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *TabOptions) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}

type Tabber interface {
	view.View
	TabOptions() *TabOptions
}
