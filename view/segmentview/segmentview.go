package segmentview

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/pb/view/segmentview"
	"gomatcha.io/matcha/view"
)

type View struct {
	*view.Embed
	Enabled       bool
	Momentary     bool
	OnValueChange func(value int)
	Titles        []string
	Value         int
	PaintStyle    *paint.Style
}

func New(ctx *view.Context, key string) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{Embed: ctx.NewEmbed(key)}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.Height(29)
		s.WidthEqual(l.MaxGuide().Width())
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
		NativeViewName: "gomatcha.io/matcha/view/segmentview",
		NativeViewState: &segmentview.View{
			Value: int64(v.Value),
		},
		NativeFuncs: map[string]interface{}{
			"OnChange": func(data []byte) {
				event := &segmentview.Event{}
				err := proto.Unmarshal(data, event)
				if err != nil {
					fmt.Println("error", err)
					return
				}

				v.Value = int(event.Value)
				if v.OnValueChange != nil {
					v.OnValueChange(v.Value)
				}
			},
		},
	}
}
