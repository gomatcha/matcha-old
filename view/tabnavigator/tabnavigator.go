package tabnavigator

import (
	"image"

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
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (n *TabNavigator) Build(ctx *view.Context) *view.Model {
	return &view.Model{}
}

func (n *TabNavigator) Views() []view.View {
	return n.views
}

func (n *TabNavigator) SetViews(vs []view.View, animated bool) {
	n.views = vs

	opts := []TabOptions{}
	notifers := []mochi.Notifiers{}
	for _, i := range n.views {
		opt, ok := i.(*TabOptions)
		if !ok {
			opt
		}
		opts = append(opts, opt)
		notifiers := append(notifiers, opt)
	}
	n.notifier = mochi.NewBatchNotifier(notifiers)
	n.options = opts
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
