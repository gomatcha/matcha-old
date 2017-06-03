package stacknavigator

import (
	"github.com/overcyn/mochi/layout/constraint"
	stacknavigatorpb "github.com/overcyn/mochi/pb/view/stacknavigator"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type StackNavigator struct {
	*view.Embed
	screens []Screen
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
	l := constraint.New()

	screenspb := []*stacknavigatorpb.Screen{}
	views := []view.View{}
	for _, i := range n.screens {
		screenpb, err := i.MarshalProtobuf()
		if err == nil {
			screenspb = append(screenspb, screenpb)
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
		NativeViewName: "github.com/overcyn/mochi/view/stacknavigator",
		NativeViewState: &stacknavigatorpb.StackNavigator{
			Screens: screenspb,
		},
	}
}

func (n *StackNavigator) Screens() []Screen {
	return n.screens
}

func (n *StackNavigator) SetScreens(ss []Screen, animated bool) {
	n.screens = ss
	n.Update()
}

func (n *StackNavigator) Push(s Screen) {
	n.screens = append(n.screens, s)
	n.Update()
}

func (n *StackNavigator) Pop() {
	if len(n.screens) > 0 {
		n.screens = n.screens[:len(n.screens)-1]
		n.Update()
	}
}

type Screen struct {
	View    view.View
	Options *Options
}

func (s *Screen) MarshalProtobuf() (*stacknavigatorpb.Screen, error) {
	return &stacknavigatorpb.Screen{
		Id:    int64(s.View.Id()),
		Title: s.Options.title,
		CustomBackButtonTitle: len(s.Options.backButtonTitle) > 0,
		BackButtonTitle:       s.Options.backButtonTitle,
		BackButtonHidden:      s.Options.backButtonHidden,
		// TitleViewId:      s.Options.titleView.Id(),
		// RightViewIds
		// LeftViewIds:..
	}, nil
}

type Options struct {
	store            store.Store
	title            string
	backButtonTitle  string
	backButtonHidden bool
	titleView        view.View
	rightViews       []view.View
	leftViews        []view.View
	// BarHidden        bool
	// Bar height?
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

func (opt *Options) SetBackButtonTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonTitle = v
}

func (opt *Options) BackButtonTitle() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonTitle
}

func (opt *Options) SetBackButtonHidden(v bool) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonHidden = v
}

func (opt *Options) BackButtonHidden() bool {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonHidden
}

func (opt *Options) SetTitleView(v view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.titleView = v
}

func (opt *Options) TitleView() view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.titleView
}

func (opt *Options) SetRightViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.rightViews = v
}

func (opt *Options) RightViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.rightViews
}

func (opt *Options) SetLeftViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.leftViews = v
}

func (opt *Options) LeftViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.leftViews
}

func (opt *Options) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *Options) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}
