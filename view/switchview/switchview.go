package switchview // import "gomatcha.io/matcha/view/switchview"

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/switchview"
	"gomatcha.io/matcha/view"
)

type View struct {
	*view.Embed
	Value         bool
	OnValueChange func(value bool)
	PaintStyle    *paint.Style
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(31))
		s.WidthEqual(constraint.Const(51))
		s.TopEqual(l.MaxGuide().Top())
		s.LeftEqual(l.MaxGuide().Left())
	})

	painter := paint.Painter(nil)
	if v.PaintStyle != nil {
		painter = v.PaintStyle
	}
	return &view.Model{
		Painter:        painter,
		Layouter:       l,
		NativeViewName: "gomatcha.io/matcha/view/switch",
		NativeViewState: &switchview.View{
			Value: v.Value,
		},
		NativeFuncs: map[string]interface{}{
			"OnChange": func(data []byte) {
				event := &switchview.Event{}
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
