package touch

import (
	"fmt"
	"time"

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
		return view.NewRoot(New(nil, nil))
	})
}

type TouchView struct {
	*view.Embed
	counter      int
	pressCounter int
}

func New(ctx *view.Context, key interface{}) *TouchView {
	if v, ok := ctx.Prev(key).(*TouchView); ok {
		return v
	}
	return &TouchView{
		Embed: view.NewEmbed(ctx.NewId(key)),
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
		go v.Update()
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
	g2 := l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
	})

	chl3 := NewPressChildView(ctx, 3)
	chl3.OnPress = func() {
		v.Lock()
		defer v.Unlock()

		fmt.Println("On Press")
		v.pressCounter += 1
		go v.Update()
	}
	g3 := l.Add(chl3, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := textview.New(ctx, 4)
	chl4.Painter = &paint.Style{BackgroundColor: colornames.Red}
	chl4.String = fmt.Sprintf("Press: %v", v.pressCounter)
	chl4.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})
	l.Add(chl4, func(s *constraint.Solver) {
		s.TopEqual(g3.Bottom())
		s.LeftEqual(g3.Left())
	})

	return &view.Model{
		Children: []view.View{chl1, chl2, chl3, chl4},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}

type PressChildView struct {
	*view.Embed
	OnPress func()
}

func NewPressChildView(ctx *view.Context, key interface{}) *PressChildView {
	if v, ok := ctx.Prev(key).(*PressChildView); ok {
		return v
	}
	return &PressChildView{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *PressChildView) Build(ctx *view.Context) *view.Model {
	tap := &touch.PressRecognizer{
		MinDuration: time.Second / 2,
		OnEvent: func(e *touch.PressEvent) {
			v.Lock()
			defer v.Unlock()
			v.OnPress()
		},
	}

	return &view.Model{
		Painter: &paint.Style{BackgroundColor: colornames.Blue},
		Values: map[interface{}]interface{}{
			touch.Key: []touch.Recognizer{tap},
		},
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
		RecognizedFunc: func(e *touch.TapEvent) {
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
