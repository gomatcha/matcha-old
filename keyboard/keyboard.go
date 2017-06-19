package keyboard

type key struct{}
type textKey struct{}

var Key = key{}
var HelperKey = textKey{}

type Responder struct {
}

// func (g *Responder) Next() {
// }

// func (g *Responder) Prev() {
// }

func (g *Responder) Show() {

}

func (g *Responder) Dismiss() {
}

func (g *Responder) Visible() bool {
	return true
}

// func (g *Responder) Notifier() *comm.BoolNotifier {
// }

// type Helper struct {
// 	responder *Responder
// 	notifyId  comm.Id
// }

// type Middleware struct {
// }

// func (m *Middleware) Build(ctx *view.Context, next *view.Model) {
// 	// Get previous helper and unsubscribe.
// 	var prevHelper *Helper
// 	if prevModel := ctx.PrevModel(); prevModel != nil && prevModel.Values != nil {
// 		prevHelper, _ = prevModel.Values[HelperKey].(*Helper)
// 	}
// 	if prevHelper != nil {
// 		prevHelper.responder.Unnotify(prevHelper.notifyId)
// 	}

// 	// Get new helper.
// 	helper, ok := next.Values[HelperKey].(*Helper)
// 	if !ok {
// 		return
// 	}
// 	helper.notifyId = helper.responder.Notify(func() {
// 		// Subscribe
// 		ctx.Update()
// 	})

// 	funcId := ctx.NewFuncId()
// 	f := func(data []byte) {
// 		// keyboard visibility

// 		if true {
// 			helper.responder.Show()
// 		} else {
// 			helper.responder.Hide()
// 		}

// 	}

// 	pb, ok := next.NativeViewState.(*textinput.View)
// 	if !ok {
// 		return
// 	}
// 	pb.KeyboardVisible = helper.responder.Visible()
// 	pb.OnKeyboard = funcId

// 	next.NativeFuncs[funcId] = f
// }
