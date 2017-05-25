package button

import (
	"fmt"
	"sync"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/pb"
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
		button.Lock()
		defer button.Unlock()
		if button.OnPress == nil {
			return
		}
		button.OnPress()
	})
}

var buttonMu sync.Mutex
var buttons = map[mochi.Id]*Button{}

func textSize(t *text.Text, max layout.Point) layout.Point {
	return mochibridge.Root().Call("sizeForAttributedString:minSize:maxSize:", mochibridge.Interface(t), nil, mochibridge.Interface(max)).ToInterface().(layout.Point)
}

type buttonLayouter struct {
	formattedText *text.Text
}

func (l *buttonLayouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	const padding = 10.0
	size := textSize(l.formattedText, ctx.MaxSize)
	g := layout.Guide{Frame: layout.Rt(0, 0, size.X+padding*2, size.Y+padding*2)}
	return g, nil
}

func (l *buttonLayouter) Notify() chan struct{} {
	// no-op
	return nil
}

func (l *buttonLayouter) Unnotify(chan struct{}) {
	// no-op
}

type Button struct {
	*view.Embed
	Text    string
	Painter paint.Painter
	OnPress func()
}

func New(ctx *view.Context, key interface{}) *Button {
	v, ok := ctx.Prev(key).(*Button)
	if !ok {
		v = &Button{
			Embed: view.NewEmbed(ctx.NewId(key)),
		}
	}
	return v
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
	ft := &text.Text{}
	ft.SetString(v.Text)
	ft.Style().SetAlignment(text.AlignmentCenter)
	ft.Style().SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   20,
	})

	return &view.Model{
		Layouter:   &buttonLayouter{formattedText: ft},
		Painter:    v.Painter,
		NativeName: "github.com/overcyn/mochi/view/button",
		NativeStateProtobuf: &pb.Button{
			Text: ft.EncodeProtobuf(),
		},
	}
}

func (v *Button) String() string {
	return fmt.Sprintf("&Button{id:%v text:%v}", v.Id(), v.Text)
}
