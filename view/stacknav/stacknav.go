package stacknav

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb/view/stacknav"
	"github.com/overcyn/mochi/store"
	"github.com/overcyn/mochi/view"
)

type StackNav struct {
	*view.Embed
	screens []*Screen
}

func New(ctx *view.Context, key interface{}) *StackNav {
	if v, ok := ctx.Prev(key).(*StackNav); ok {
		return v
	}
	return &StackNav{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (n *StackNav) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	screenspb := []*stacknav.Screen{}
	views := []view.View{}
	for _, i := range n.screens {
		screenpb, err := i.MarshalProtobuf()
		if err == nil {
			screenspb = append(screenspb, screenpb)
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
		NativeViewName: "github.com/overcyn/mochi/view/stacknav",
		NativeViewState: &stacknav.StackNav{
			Screens: screenspb,
		},
	}
}

func (n *StackNav) Screens() []*Screen {
	return n.screens
}

func (n *StackNav) SetScreens(ss []*Screen, animated bool) {
	n.screens = ss
	n.Update()
}

func (n *StackNav) Push(s *Screen) {
	n.screens = append(n.screens, s)
	n.Update()
}

func (n *StackNav) Pop() {
	if len(n.screens) > 0 {
		n.screens = n.screens[:len(n.screens)-1]
		n.Update()
	}
}

type Screen struct {
	store            store.Store
	view             view.View
	title            string
	backButtonTitle  string
	backButtonHidden bool
	titleView        view.View
	rightViews       []view.View
	leftViews        []view.View
	// BarHidden        bool
	// Bar height?
}

func (s *Screen) MarshalProtobuf() (*stacknav.Screen, error) {
	return &stacknav.Screen{
		Id:    int64(s.view.Id()),
		Title: s.title,
		CustomBackButtonTitle: len(s.backButtonTitle) > 0,
		BackButtonTitle:       s.backButtonTitle,
		BackButtonHidden:      s.backButtonHidden,
		// TitleViewId:      s.titleView.Id(), // TODO(KD):
		// RightViewIds
		// LeftViewIds:..
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

func (opt *Screen) SetBackButtonTitle(v string) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonTitle = v
}

func (opt *Screen) BackButtonTitle() string {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonTitle
}

func (opt *Screen) SetBackButtonHidden(v bool) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.backButtonHidden = v
}

func (opt *Screen) BackButtonHidden() bool {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.backButtonHidden
}

func (opt *Screen) SetTitleView(v view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.titleView = v
}

func (opt *Screen) TitleView() view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.titleView
}

func (opt *Screen) SetRightViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.rightViews = v
}

func (opt *Screen) RightViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.rightViews
}

func (opt *Screen) SetLeftViews(v []view.View) {
	tx := store.NewWriteTx()
	defer tx.Commit()

	opt.store.Write(tx)
	opt.leftViews = v
}

func (opt *Screen) LeftViews() []view.View {
	tx := store.NewReadTx()
	defer tx.Commit()

	opt.store.Read(tx)
	return opt.leftViews
}

func (opt *Screen) Notify() chan struct{} {
	return opt.store.Notify()
}

func (opt *Screen) Unnotify(c chan struct{}) {
	opt.store.Unnotify(c)
}
