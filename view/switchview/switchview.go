package switchview

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/pb/view/switchview"
	"github.com/overcyn/mochi/view"
)

type View struct {
	*view.Embed
	Value         bool
	OnValueChange func(*View)
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

	funcId := ctx.NewFuncId()
	f := func(data []byte) {
		event := &switchview.Event{}
		err := proto.Unmarshal(data, event)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		v.Value = event.Value
		if v.OnValueChange != nil {
			v.OnValueChange(v)
		}
	}

	return &view.Model{
		Layouter:       l,
		NativeViewName: "github.com/overcyn/mochi/view/switch",
		NativeViewState: &switchview.View{
			Value:         v.Value,
			OnValueChange: funcId,
		},
		NativeFuncs: map[int64]interface{}{
			funcId: f,
		},
	}
}
