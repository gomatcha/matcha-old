package button

import (
	"fmt"
	"sync"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout"
	pbbutton "github.com/overcyn/mochi/pb/button"
	"github.com/overcyn/mochi/text"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochibridge"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/view/button OnPress", func(id int64) {
		buttonMu.Lock()
		defer buttonMu.Unlock()

		button := buttons[mochi.Id(id)]
		if button == nil {
			return
		}
		view.MainMu.Lock()
		defer view.MainMu.Unlock()
		if button.OnPress == nil {
			return
		}
		button.OnPress()
	})
}

var buttonMu sync.Mutex
var buttons = map[mochi.Id]*Button{}

type layouter struct {
	styledText *text.StyledText
}

func (l *layouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	const padding = 10.0
	size := l.styledText.Size(layout.Pt(0, 0), ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

func (l *layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l *layouter) Unnotify(id comm.Id) {
	// no-op
}

type Button struct {
	*view.Embed
	Text    string
	OnPress func()
}

func New(ctx *view.Context, key string) *Button {
	if v, ok := ctx.Prev(key).(*Button); ok {
		return v
	}
	return &Button{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *Button) Lifecycle(from, to view.Stage) {
	if view.EntersStage(from, to, view.StageMounted) {
		buttonMu.Lock()
		defer buttonMu.Unlock()

		buttons[v.Id()] = v
	} else if view.ExitsStage(from, to, view.StageMounted) {
		buttonMu.Lock()
		defer buttonMu.Unlock()

		delete(buttons, v.Id())
	}
}

func (v *Button) Build(ctx *view.Context) *view.Model {
	style := &text.Style{}
	style.SetAlignment(text.AlignmentCenter)
	style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})

	t := text.New(v.Text)

	st := text.NewStyledText(t)
	st.Set(style, 0, 0)

	return &view.Model{
		Layouter:       &layouter{styledText: st},
		NativeViewName: "github.com/overcyn/mochi/view/button",
		NativeViewState: &pbbutton.Button{
			StyledText: st.MarshalProtobuf(),
		},
	}
}

func (v *Button) String() string {
	return fmt.Sprintf("&Button{id:%v text:%v}", v.Id(), v.Text)
}
