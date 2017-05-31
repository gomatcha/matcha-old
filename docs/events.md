type TodoView struct {
    model *TodoViewModel
    operations chan Operations
    marker chan<- Mochi.Marker
}

type TodoViewModel {
    Items []string
    Input string
}

const (
    labelId = "todo.label"
    listId = "todo.list"
    textFieldId = "todo.textField"
    buttonId = "todo.button"
)

func NewTodoView(v interface{}) {
    todoView, ok := v.(*TodoView)
    if !ok {
        todoView = (*TodoView){}
    }
    return todoView
}

func (v *TodoView) Load(sig chan<- Mochi.Signal) {
    v.sig = sig
}

func (v *TodoView) Unload() {
    
}

// func (v *TodoView) Run(sig <-chan Mochi.Signal, nodeChan chan<- *Node) {
//     for {
//         select {
//         case s := <-sig:
//             nodeChan <- v.Render()

//         case op := <-v.operations
//             op()
//         }
//     }
// }

func (v *TodoView)setModel(m *TodoViewModel) {
    v.operations <- func() {
        v.model = m
        v.marker<- Mochi.Update()
    }
}

func (v *TodoView) Update(p *Node) *Node {
    l := &constraint.System{}
    n := &Node{}
    n.layouter = l

    var prev *constraint.Guide
    {
        // Label
        chl := NewLabel(p.Get(labelId))
        chl.Text = "TODO"
        n.Set(labelId, chl)

        prev = l.Add(labelId, func(constraint.Solver *s){
            s.WidthEqual(l.Width().Multiply(0.5))
            s.HeightEqual(l.Height().Multiply(0.5))
            s.TopEqual(constraint.Const(10))
            s.BottomEqual(constraint.Const(10))
        })
    }
    {
        // List
        chl := NewList(p.Get(listId))
        chl.Items = v.Items
        n.Set(listID, chl)

        prev = l.Add(listId, func(constraint.Solver *s){
            s.TopEqual(prev.Bot())
            s.BotLess(l.Bot())
        })
    }
    {
        // Text Input
        // if v.textListener {
        //     v.textListener.Close()
        // }
        
        chl := NewTextField(p.Get(textFieldId))
        chl.Input = v.Input
        chl.OnChange = func(s string) {
            v.Lock()
            defer v.Unlock()
            
            v.Input = s
            v.Update(nil)
        }
        
        chl.OnChange = func(s string) {
            v.UpdateFunc(func() {
                v.Input = s
            })
        }
        
        // l := sig.NewLoop(v.onChange)
        // l.Chan = chl.OnChange()
        // l.Func = func(s string){
        //     v.Lock()
        //     defer v.Unlock()
            
        //     v.Input = s
        //     v.Update()
        // })
        // l.Listen()
        
        // chl := NewTextField(p.Get(textFieldId))
        // chl.Input = v.Input
        
        // v.onChange := NewListener(v.onChange)
        // onChange.Chan = chl.OnChange()
        // onChange.Sig = sig
        // onChange.Func = func(s string){
        //     v.Input = s
        //     v.sig <-Mochi.Update{}
        // })
        // onChange.Listen()
        
        // v.Listener = sig.NewListener(chl.OnChange(), func(){
        //     v.Input = blah
        //     v.sig <-Mochi.Update{}
        // })
        
        // chl := NewTextField(p.Get(textFieldId))
        // chl.Input = v.Input
        // v.Listener = NewStringListener(chl.OnChange(), sig, func(){
        //     v.Input = blah
        //     v.sig <-Mochi.Update{}
        // })
        
        // chl := NewTextField(p.Get(textFieldId))
        // chl.Input = v.Input
        // l := chl.TextListener()
        // go func() {
        //     x, ok := <- l.Updates()
        //     if !ok {
        //         return
        //     }
        //     v.sig <- Mochi.Func{func() {
        //         // if l.Open {
        //         v.Input = x
        //         v.sig <-Mochi.Update{}
        //         // }
        //     }}
        // }()
        // v.textListener = l
        
        
        // chl.OnChange = func(str string) {
            // v.sig <- Mochi.Func{func() {
            //     v.Input = str
            //     v.sig <- Mochi.Update{}
            // }}
        // }
        
        // cancel := make(chan interface{})
        // input := make(chan chan string)
        // chl.OnChange(input, cancel)
        // go func () {
        //     select {
        //     case in := <-input
        //         v.sig <- Mochi.Func{func() {
        //             blah = 
        //             // v.Input = in
        //             // v.sig <-Mochi.Update{}
        //         }}
        //     case cancel
        //     }
        // }
        
        // if v.textListener {
        //     v.textListener.Close()
        //     v.textListener = nil
        // }
        
        // v.textListener := chl.OnChange()
        // go func(l Listener) {
        //     blah := <- l.Listen()
        //     v.sig <- Mochi.Func{func() {
        //         if v.textListener = l {
        //             v.Input = in
        //             v.sig <-Mochi.Update{}
        //         }
        //     }}
        // }(v.textListener)
        
        // cancel := make(chan interface{})
        // input := make(chan string)
        // chl.OnChange(input, cancel)
        // go func () {
        //     select {
        //     case in := <-input
        //         v.sig <- Mochi.Func{func() {
        //             v.Input = in
        //             v.sig <-Mochi.Update{}
        //         }}
        //     case cancel
        //     }
        // }
        
        // go func() {
        //     blah := <-chl.Callback()
        //     v.Lock()
        //     defer v.Unlock()
            
        //     v.Input = str
        //     v.NeedsUpdate()
        // }
        // chl.OnChange = func(str string) {
        //     v.Input = str
        //     v.NeedsUpdate()
        // }
        n.Set(textFieldId, chl)

        prev = l.Add(textFieldId, func(constraint.Solver *s){
            s.TopEqual(prev.Bot())
            s.BotLess(l.Bot())
        })
    }
    {
        // Button
        chl := NewButton(p.Get(buttonId))
        cancel := make(chan interface{})
        go func() {
            <-chl.OnClick()
            
            select {
            case <-chl.OnClick():
                // v.sig <- Mochi.Func{func() {
                if v.Input == "" {
                    return
                }
                v.Items = append(v.Items, v.Input)
                v.Input = ""
                v.sig <- Mochi.Update{}
                // }}
            case <-cancel:
            }
        }
        // chl.OnClick = func() {
        //     if v.Input == "" {
        //         return
        //     }
        //     append(v.Items, v.Input)
        //     v.Input = ""
        //     v.NeedsUpdate()
        // }
        n.Set(buttonId, chl)

        prev = l.Add(buttonId, func(constraint.Solver *s){
            s.TopEqual(prev.Bot())
            s.BotLess(l.Bot())
        })
    }

    l.Solve(func(constraint.Solver *s){
        s.BotEqual(prev.Bot())
    })
    return n
}