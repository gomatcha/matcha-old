package slider

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha"
	"gomatcha.io/matcha/comm"
	"gomatcha.io/matcha/layout"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/slider"
	"gomatcha.io/matcha/view"
)

type layouter struct {
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[matcha.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rt(0, 0, ctx.MinSize.X, 31)}
	return g, nil
}

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}

type View struct {
	*view.Embed
	PaintStyle    *paint.Style
	DefaultValue  float64
	Value         *comm.Float64Value
	MaxValue      float64
	MinValue      float64
	OnValueChange func(value float64)
	OnSubmit      func(value float64) // TODO(KD): naming? OnBeginEdit, OnEndEdit? OnValueBegin, OnValueEnd
	Enabled       bool
	initialized   bool
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed:    ctx.NewEmbed(key),
		MaxValue: 1,
		MinValue: 0,
		Enabled:  true,
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	val := 0.0
	if !v.initialized {
		v.initialized = true
		val = v.DefaultValue
	}
	if v.Value != nil {
		val = v.Value.Value()
	}

	painter := paint.Painter(nil)
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		Layouter:       &layouter{},
		NativeViewName: "gomatcha.io/matcha/view/slider",
		NativeViewState: &slider.View{
			Value:    val,
			MaxValue: v.MaxValue,
			MinValue: v.MinValue,
			Enabled:  v.Enabled,
		},
		NativeFuncs: map[string]interface{}{
			"OnValueChange": func(data []byte) {
				event := &slider.Event{}
				err := proto.Unmarshal(data, event)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				if v.OnValueChange != nil {
					v.OnValueChange(event.Value)
				}
			},
			"OnSubmit": func(data []byte) {
				event := &slider.Event{}
				err := proto.Unmarshal(data, event)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				if v.OnSubmit != nil {
					v.OnSubmit(event.Value)
				}
			},
		},
	}
}
