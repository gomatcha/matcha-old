package touch

import (
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/basicview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/touch New", func() *view.ViewController {
		return view.NewViewController(func(c view.Config) view.View {
			return New(c)
		}, 0)
	})
}

type AnimateView struct {
	*view.Embed
}

func New(c view.Config) *AnimateView {
	v, ok := c.Prev.(*AnimateView)
	if !ok {
		v = &AnimateView{
			Embed: c.Embed,
		}
	}
	return v
}

func (v *AnimateView) Build(ctx *view.Context) *view.Model {
	tap := touch.NewTapRecognizer(ctx, 1)
	tap.RecognizedFunc = func(e *touch.TapEvent) {
		// do something
	}

	l := constraint.New()
	m := &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
		Values: map[interface{}]interface{}{
			touch.Key(): []touch.Recognizer{tap},
		},
	}

	chl := basicview.New(ctx, 1)
	chl.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	m.Add(chl)
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})

	// r := &touch.Recognizer{}
	// pan := touch.NewPanRecognizer(ctx, 1)
	// pan.BeganFunc = func() {
	// }
	// pan.CancelledFunc = func() {
	// }
	// r.add(pan)

	// v.store.Observe(pan.ChangedNotifier)

	return m
}
