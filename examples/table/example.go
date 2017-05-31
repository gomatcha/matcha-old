package example

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/table New", func() *view.Root {
		return view.NewRoot(New(nil, nil), 0)
	})
}

type TableView struct {
	*view.Embed
}

func New(ctx *view.Context, key interface{}) *TableView {
	if v, ok := ctx.Prev(key).(*TableView); ok {
		return v
	}
	return &TableView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}

	childLayouter := &table.Layout{}
	childViews := []view.View{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, i+1000)
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: colornames.Red}
		childViews = append(childViews, childView)
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx, 0)
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childViews

	scrollView := scrollview.New(ctx, 1)
	scrollView.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
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

func NewTableCell(ctx *view.Context, key interface{}) *TableCell {
	if v, ok := ctx.Prev(key).(*TableCell); ok {
		return v
	}
	return &TableCell{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	n := &view.Model{}
	n.Layouter = l
	n.Painter = v.Painter

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx, 1)
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
