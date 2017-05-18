package animate

import (
	"fmt"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/animate"
	"github.com/overcyn/mochi/internal"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochibridge"
	"time"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/animate New", func() *view.ViewController {
		return view.NewViewController(func(c view.Config) view.View {
			return New(c)
		}, 0)
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

func New(c view.Config) *AnimateView {
	v, ok := c.Prev.(*AnimateView)
	if !ok {
		v = &AnimateView{}
		v.Embed = c.Embed
		v.ticker = animate.NewTicker(time.Second * 4)
		v.floatTicker = animate.FloatInterpolate(v.ticker, animate.FloatLerp{Start: 0, End: 500})
		v.colorTicker = animate.ColorInterpolate(v.ticker, animate.RGBALerp{Start: internal.RedColor, End: internal.YellowColor})

		_ = mochi.NotifyFunc(v.ticker, func() {
			fmt.Println("Ticker update")
		})
		v.floatTickerFunc = mochi.NotifyFunc(v.floatTicker, func() {
			fmt.Println("Float update")
		})
		// ticker := mochi.NotifyFunc(v.colorTicker, func() {
		// 	fmt.Println("Float 2 update")
		// })

		// time.AfterFunc(time.Second*2, func() {
		// 	close(ticker)
		// 	mochi.NotifyFunc(v.colorTicker, func() {
		// 		fmt.Println("Float 3 update")
		// 	})
		// })
	}
	return v
}

func (v *AnimateView) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageVisible) {
		time.AfterFunc(time.Second*2, func() {
			fmt.Println("Update")
			v.Update(nil)
		})
	}
}

func (v *AnimateView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: internal.GreenColor},
	}

	chl := basicview.New(ctx.Get(1))
	chl.Painter = &paint.Style{BackgroundColor: internal.BlueColor}
	// chl.Painter = &paint.AnimatedStyle{BackgroundColor: v.colorTicker}
	m.Add(chl)
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		// s.WidthEqual(constraint.Notifier(v.floatTicker))
		// s.HeightEqual(constraint.Notifier(v.floatTicker))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})

	// if v.floatTickerFunc != nil {
	// 	close(v.floatTickerFunc)
	// }
	// v.floatTickerFunc = mochi.NotifyFunc(v.floatTicker, func() {
	// 	fmt.Println("Float update")
	// })
	if v.constraintFunc != nil {
		close(v.constraintFunc)
	}
	v.constraintFunc = mochi.NotifyFunc(l, func() {
		fmt.Println("Constraint update")
	})

	return m
}
