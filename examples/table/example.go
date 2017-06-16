package example

import (
	"strconv"

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
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context, key string) view.View {
			return New(ctx, key)
		}))
	})
}

type TableView struct {
	*view.Embed
}

func New(ctx *view.Context, key string) *TableView {
	if v, ok := ctx.Prev(key).(*TableView); ok {
		return v
	}
	return &TableView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	childLayouter := &table.Layout{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, strconv.Itoa(i))
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: colornames.Red}
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx, "a")
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childLayouter.Views()

	scrollView := scrollview.New(ctx, "b")
	scrollView.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	scrollView.ContentView = scrollChild
	_ = l.Add(scrollView, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(400))
		s.HeightEqual(constraint.Const(400))
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}

}

type TableCell struct {
	*view.Embed
	String  string
	Painter paint.Painter
}

func NewTableCell(ctx *view.Context, key string) *TableCell {
	if v, ok := ctx.Prev(key).(*TableCell); ok {
		return v
	}
	return &TableCell{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TableCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(50))
	})

	textView := textview.New(ctx, "a")
	textView.String = v.String
	textView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	l.Add(textView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(10))
		s.RightEqual(l.Right().Add(-10))
		s.CenterYEqual(l.CenterY())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  v.Painter,
	}
}
