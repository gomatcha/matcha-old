package touch

import (
	"fmt"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/touch"
	"github.com/overcyn/mochi/view"
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
}

func New(c view.Config) *TouchView {
	v, ok := c.Prev.(*TouchView)
	if !ok {
		v = &TouchView{
			Embed: c.Embed,
		}
	}
	return v
}

func (v *TouchView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl := NewTouchChildView(ctx, 1)
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

	return &view.Model{
		Children: map[mochi.Id]view.View{chl.Id(): chl},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type TouchChildView struct {
	*view.Embed
}

func NewTouchChildView(ctx *view.Context, key interface{}) *TouchChildView {
	v, ok := ctx.Prev(key).(*TouchChildView)
	if !ok {
		v = &TouchChildView{
			Embed: view.NewEmbed(ctx.NewId(key)),
		}
	}
	return v
}

func (v *TouchChildView) Build(ctx *view.Context) *view.Model {
	tap := &touch.TapRecognizer{
		Count: 1,
		RecognizedFunc: func() {
			fmt.Println("touched")
			// do something
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: colornames.Blue},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
	}
}
