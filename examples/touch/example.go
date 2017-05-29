package touch

import (
	"fmt"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/textview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/touch New", func() *view.Root {
		return view.NewRoot(func(c view.Config) view.View {
			return New(c)
		}, 0)
	})
}

type TouchView struct {
	*view.Embed
	counter int
}

func New(c view.Config) *TouchView {
	if v, ok := c.Prev.(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: c.Embed,
	}
}

func (v *TouchView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl1 := NewTouchChildView(ctx, 1)
	chl1.OnTouch = func() {
		v.Lock()
		defer v.Unlock()

		fmt.Println("On touch")
		v.counter += 1
		go v.Update(nil)
	}
	g1 := l.Add(chl1, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2 := textview.New(ctx, 2)
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Red}
	chl2.String = fmt.Sprintf("Counter: %v", v.counter)
	chl2.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
	})

	l.Solve(func(s *constraint.Solver) {
		s.WidthEqual(l.MaxGuide().Width())
		s.HeightEqual(l.MaxGuide().Height())
	})

	return &view.Model{
		Children: map[mochi.Id]view.View{chl1.Id(): chl1, chl2.Id(): chl2},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type TouchChildView struct {
	*view.Embed
	OnTouch func()
}

func NewTouchChildView(ctx *view.Context, key interface{}) *TouchChildView {
	if v, ok := ctx.Prev(key).(*TouchChildView); ok {
		return v
	}
	return &TouchChildView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TouchChildView) Build(ctx *view.Context) *view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		RecognizedFunc: func() {
			v.Lock()
			defer v.Unlock()
			v.OnTouch()
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: colornames.Blue},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
