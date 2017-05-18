package example

import (
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/table New", func() *view.ViewController {
		return view.NewViewController(func(c view.Config) view.View {
			return New(c)
		}, 0)
	})
}

type TableView struct {
	*view.Embed
}

func New(c view.Config) *TableView {
	v, ok := c.Prev.(*TableView)
	if !ok {
		v = &TableView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TableView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: internal.GreenColor},
	}

	childLayouter := &table.Layout{}
	childViews := []view.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx.Get(i + 1000))
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: internal.RedColor}
		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx.Get(10))
	scrollChild.Painter = &paint.Style{BackgroundColor: internal.WhiteColor}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childViews

	scrollView := scrollview.New(ctx.Get(11))
	scrollView.Painter = &paint.Style{BackgroundColor: internal.CyanColor}
	scrollView.ContentView = scrollChild
	m.Add(scrollView)
	_ = l.Add(scrollView, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(400))
		s.HeightEqual(constraint.Const(400))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return m
}

type TableCell struct {
	*view.Embed
	String  string
	Painter paint.Painter
}

func NewTableCell(c view.Config) *TableCell {
	v, ok := c.Prev.(*TableCell)
	if !ok {
		v = &TableCell{}
		v.Embed = c.Embed
	}
	return v
}

func (v *TableCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	n := &view.Model{}
	n.Layouter = l
	n.Painter = v.Painter

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx.Get(1))
	textView.String = v.String
	textView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	n.Add(textView)
	l.Add(textView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(10))
		s.RightEqual(l.Right().Add(-10))
		s.CenterYEqual(l.CenterY())
	})

	return n
}
