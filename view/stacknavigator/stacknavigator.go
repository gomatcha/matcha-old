package stacknavigator

import (
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type StackNavigator struct {
	*view.Embed
	views []view.View
}

func New(ctx *view.Context, key interface{}) *StackNavigator {
	if v, ok := ctx.Prev(key).(*StackNavigator); ok {
		return v
	}
	return &StackNavigator{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (n *StackNavigator) Build(ctx *view.Context) *view.Model {
	return &view.Model{}
}

func (n *StackNavigator) Set(vs []view.View) {

}

func (n *StackNavigator) Push(v view.View) {

}

func (n *StackNavigator) Pop() {
}

type StackOptions struct {
	store store.Store

	title            string
	backButtonTitle  string
	backButtonHidden bool
	titleView        view.View
	rightViews       []view.View
	leftViews        []view.View
	// BarHidden        bool
	// Bar height?
}

func (opt *StackOptions) SetTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.title = v
}

func (opt *StackOptions) Title() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.title
}

func (opt *StackOptions) SetBackButtonTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonTitle = v
}

func (opt *StackOptions) BackButtonTitle() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonTitle
}

func (opt *StackOptions) SetBackButtonHidden(v bool) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonHidden = v
}

func (opt *StackOptions) BackButtonHidden() bool {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonHidden
}

func (opt *StackOptions) SetTitleView(v view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.titleView = v
}

func (opt *StackOptions) TitleView() view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.titleView
}

func (opt *StackOptions) SetRightViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.rightViews = v
}

func (opt *StackOptions) RightViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.rightViews
}

func (opt *StackOptions) SetLeftViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.leftViews = v
}

func (opt *StackOptions) LeftViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.leftViews
}

func (opt *StackOptions) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *StackOptions) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}

type Stacker interface {
	view.View
	StackOptions() *StackOptions
}
