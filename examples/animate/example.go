package animate

import (
	"fmt"
	"time"

	"github.com/gomatcha/matcha/layout/constraint"
	"github.com/gomatcha/matcha/paint"
	"github.com/gomatcha/matcha/view"
	"github.com/gomatcha/matcha/view/basicview"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/gomatcha/matcha/examples/animate New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			return New(ctx, "")
		}))
	})
}

type AnimateView struct {
	*view.Embed
	// ticker      *animate.Ticker
	// floatTicker comm.Float64Notifier
	// colorTicker comm.ColorNotifier

	// floatTickerFunc chan struct{}
	// constraintFunc  chan struct{}
}

func New(ctx *view.Context, key string) *AnimateView {
	if v, ok := ctx.Prev(key).(*AnimateView); ok {
		return v
	}
	// ticker := animate.NewTicker(time.Second * 4)
	return &AnimateView{
		Embed: view.NewEmbed(ctx.NewId(key)),
		// ticker:      ticker,
		// floatTicker: animate.FloatInterpolate(ticker, animate.FloatLerp{Start: 0, End: 500}),
		// colorTicker: animate.ColorInterpolate(ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow}),
	}
}

func (v *AnimateView) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageVisible) {
		time.AfterFunc(time.Second*2, func() {
			fmt.Println("Update")
			v.Update()
		})
	}
}

func (v *AnimateView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl := basicview.New(ctx, "")
	// chl.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		// s.WidthEqual(constraint.Notifier(v.floatTicker))
		// s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
