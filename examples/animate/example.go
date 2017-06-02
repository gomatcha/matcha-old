package animate

import (
	"fmt"
	"time"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/animate New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type AnimateView struct {
	*view.Embed
	ticker      *animate.Ticker
	floatTicker mochi.Float64Notifier
	colorTicker mochi.ColorNotifier

	floatTickerFunc chan struct{}
	constraintFunc  chan struct{}
}

func New(ctx *view.Context, key interface{}) *AnimateView {
	if v, ok := ctx.Prev(key).(*AnimateView); ok {
		return v
	}
	ticker := animate.NewTicker(time.Second * 4)
	return &AnimateView{
		Embed:       view.NewEmbed(ctx.NewId(key)),
		ticker:      ticker,
		floatTicker: animate.FloatInterpolate(ticker, animate.FloatLerp{Start: 0, End: 500}),
		colorTicker: animate.ColorInterpolate(ticker, animate.RGBALerp{Start: colornames.Red, End: colornames.Yellow}),
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

	chl := basicview.New(ctx, 1)
	chl.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Notifier(v.floatTicker))
		s.HeightEqual(constraint.Notifier(v.floatTicker))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})
	return &view.Model{
		Children: []view.View{chl},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
