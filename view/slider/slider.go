package slider

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/matcha"
	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/pb/view/slider"
	"github.com/overcyn/matcha/view"
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
	Value         float64
	MaxValue      float64
	MinValue      float64
	OnValueChange func(value float64)
	Enabled       bool
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed:    view.NewEmbed(ctx.NewId(key)),
		MaxValue: 1,
		MinValue: 0,
		Enabled:  true,
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	painter := paint.Painter(nil)
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		Layouter:       &layouter{},
		NativeViewName: "github.com/overcyn/matcha/view/slider",
		NativeViewState: &slider.View{
			Value:    v.Value,
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

				v.Value = event.Value
				if v.OnValueChange != nil {
					v.OnValueChange(v.Value)
				}
			},
		},
	}
}
