package example

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/image/colornames"

	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/layout/table"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochi/view/button"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochi/view/scrollview"
	"github.com/overcyn/mochi/view/switchview"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochi/view/urlimageview"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/complex New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return NewNestedView(ctx, "")
		}))
	})
}

type NestedView struct {
	*view.Embed
	counter int
	ticker  *animate.Ticker
}

func NewNestedView(ctx *view.Context, key string) *NestedView {
	if v, ok := ctx.Prev(key).(*NestedView); ok {
		return v
	}
	return &NestedView{
		Embed:  view.NewEmbed(ctx.NewId(key)),
		ticker: animate.NewTicker(time.Second * 5),
	}
}

func (v *NestedView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	value := animate.FloatInterpolate(v.ticker, animate.FloatLerp{Start: 0, End: 150})
	color := animate.ColorInterpolate(v.ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow})

	chl1 := basicview.New(ctx, "1")
	chl1.Painter = &paint.AnimatedStyle{BackgroundColor: color}
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Notifier(value))
		s.HeightEqual(constraint.Notifier(value))
	})

	chl2 := basicview.New(ctx, "2")
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Yellow}
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := basicview.New(ctx, "3")
	chl3.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := basicview.New(ctx, "4")
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Magenta}
	g4 := l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := textview.New(ctx, "a")
	chl5.String = "Subtitle"
	chl5.Style.SetAlignment(text.AlignmentCenter)
	chl5.Style.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl5.Style.SetStrikethroughColor(colornames.Magenta)
	chl5.Style.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl5.Style.SetUnderlineColor(colornames.Green)
	chl5.Style.SetFont(text.Font{
		Family: "American Typewriter",
		Face:   "Bold",
		Size:   20,
	})
	chl5p := view.WithPainter(chl5, &paint.Style{BackgroundColor: colornames.Cyan})

	g5 := l.Add(chl5p, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl6 := textview.New(ctx, "6")
	chl6.String = fmt.Sprintf("Counter: %v", v.counter)
	chl6.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	chl6p := view.WithPainter(chl6, &paint.Style{BackgroundColor: colornames.Red})
	g6 := l.Add(chl6p, func(s *constraint.Solver) {
		s.BottomEqual(g5.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	chl8 := button.New(ctx, "8")
	chl8.Text = "Button"
	chl8.OnPress = func(b *button.Button) {
		fmt.Println("On Click")
		v.counter += 1
		v.Update()
	}
	g8 := l.Add(chl8, func(s *constraint.Solver) {
		s.BottomEqual(g6.Top())
		s.RightEqual(g2.Right().Add(-15))
	})

	if v.counter%2 == 0 {
		chl9 := urlimageview.New(ctx, "7")
		chl9.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
		chl9.ResizeMode = imageview.ResizeModeFit
		pChl9 := view.WithPainter(chl9, &paint.Style{BackgroundColor: colornames.Cyan})

		_ = l.Add(pChl9, func(s *constraint.Solver) {
			s.BottomEqual(g8.Top())
			s.RightEqual(g2.Right().Add(-15))
			s.WidthEqual(constraint.Const(200))
			s.HeightEqual(constraint.Const(200))
		})
	}
	chl11 := switchview.New(ctx, "12")
	chl11.OnValueChange = func(a *switchview.View) {
		fmt.Println("switch tapped", a.Value)
	}
	_ = l.Add(chl11, func(s *constraint.Solver) {
		s.LeftEqual(g6.Right())
		s.TopEqual(g6.Top())
	})

	childLayouter := &table.Layout{}
	for i := 0; i < 20; i++ {
		childView := NewTableCell(ctx, "a"+strconv.Itoa(i))
		childView.String = "TEST TEST"
		childView.Painter = &paint.Style{BackgroundColor: colornames.Red}
		childLayouter.Add(childView)
	}

	scrollChild := basicview.New(ctx, "9")
	scrollChild.Painter = &paint.Style{BackgroundColor: colornames.White}
	scrollChild.Layouter = childLayouter
	scrollChild.Children = childLayouter.Views()

	chl10 := scrollview.New(ctx, "10")
	chl10.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	chl10.ContentView = scrollChild
	_ = l.Add(chl10, func(s *constraint.Solver) {
		s.TopEqual(g4.Bottom())
		s.LeftEqual(g4.Left())
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
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

	textView := textview.New(ctx, "1")
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
