package animate

import (
	"fmt"
	"time"

	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
)

type View struct {
	*view.Embed
	// ticker      *animate.Ticker
	// floatTicker comm.Float64Notifier
	// colorTicker comm.ColorNotifier

	// floatTickerFunc chan struct{}
	// constraintFunc  chan struct{}
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	// ticker := animate.NewTicker(time.Second * 4)
	return &View{
		Embed: ctx.NewEmbed(key),
		// ticker:      ticker,
		// floatTicker: animate.FloatInterpolate(ticker, animate.FloatLerp{Start: 0, End: 500}),
		// colorTicker: animate.ColorInterpolate(ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow}),
	}
}

func (v *View) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageVisible) {
		time.AfterFunc(time.Second*2, func() {
			fmt.Println("Update")
			v.Signal()
		})
	}
}

func (v *View) Build(ctx *view.Context) view.Model {
	l := constraint.New()

	chl := basicview.New(ctx, "")
	// chl.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		// s.WidthEqual(constraint.Notifier(v.floatTicker))
		// s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	return view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
