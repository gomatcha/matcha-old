package basic

// import (
// 	"fmt"
// 	"github.com/overcyn/mochi"
// )

// func main() {
// 	fmt.Println("Hello, 世界")
// mochi.Display(nil)
// }

// type AnimatedView struct {
// 	Animator Animator
// }

// func AddAnimatedView(p, n *Node, k interface{}) *AnimatedView {
// 	v, ok = p.Children[k].(*AnimatedView)
// 	if !ok {
// 		v = &AnimatedView{}
// 	}
// 	return v
// }

// func (v *AnimatedView)Update(prev, next *Node) {
// 	next.Layouter = v.UpdateLayouter(prev.Layouter)
// }

// func (v *AnimatedView)OnMount() {
// 	v.Animator.OnUpdate = func() {
// 		v.NeedsRepaint()
// 	}
// }

// func (v *AnimatedView)UpdatePaint(prev *Painter) {

// }

// func (v *AnimatedView)UpdateLayout(prev *Layouter) {
// 	next := Layouter()
// 	next.Point.X = v.Animator.Tween
// 	return next
// }

// func (v *AnimatedView)UpdateChildren(prev map[string]Component) {
// 	return prev
// }

// type TodoView struct {
// 	Items []string
// 	Input string
// }

// const (
// 	labelId = "todo.label"
// 	listId = "todo.list"
// 	textFieldId = "todo.textField"
// 	buttonId = "todo.button"
// 	scrollId = "todo.scroll"
// 	wrapperId = "todo.centerWrapper"
// )

// func NewTodoView(v interface{}) {
// 	todoView, ok := v.(*TodoView)
// 	if !ok {
// 		todoView = (*TodoView){}
// 	}
// 	return todoView
// }

// func (v *TodoView) Update(p *Node) *Node {
// 	l := &constraint.System{}
// 	l.Update(func(constraint.Solver *s) {
// 		s.HeightEqual(constraint.Const(40))
// 		s.Equal(l.Max())
// 	})

// 	n := &Node{}
// 	n.layouter = l

// 	// Wrapper
// 	wrap := l.Add(wrapId)
// 	var prev constraint.Guide
// 	{
// 		// Label
// 		chl := NewLabel(p.Get(labelId))
// 		chl.Text = "TODO"
// 		n.Set(labelId, chl)

// 		prev = l.Add(labelID)
// 		prev.Solve(func(constraint.Solver *s){
// 			s.TopEqual(l.Top())
// 			s.BotLess(l.Bot())
// 		})
// 	}
// 	{
// 		// List
// 		chl := NewList(p.Get(listId))
// 		chl.Items = v.Items
// 		n.Set(listID, chl)

// 		prev = l.Add(labelID)
// 		prev.Solve(func(constraint.Solver *s){
// 			s.TopEqual(prev.Bot())
// 			s.BotLess(l.Bot())
// 		})
// 	}
// 	{
// 		// Text
// 		chl := NewTextField(p.Get(textFieldId))
// 		chl.Input = v.Input
// 		chl.OnChange = func(str string) {
// 			v.Input = str
// 			v.NeedsUpdate()
// 		}
// 		n.Add(textFieldId, chl)

// 		cst := constraint.New(wrap)
// 		cst.TopEqual(prev.Top())
// 		cst.BotLess(wrap.Bot())
// 		cst.Equal(wrap)
// 		l.Solve(textId, cst)
// 		prev = cst
// 	}
// 	{
// 		// Button
// 		chl := NewButton(p.Get(buttonId))
// 		chl.OnClick = func() {
// 			if v.Input == "" {
// 				return
// 			}
// 			append(v.Items, v.Input)
// 			v.Input = ""
// 			v.NeedsUpdate()
// 		}
// 		n.Add(buttonId, chl)

// 		cst := constraint.New()
// 		cst.TopEqual(prev.Top())
// 		cst.BotLess(wrap.Bot())
// 		cst.Equal(wrap)
// 		l.Solve(buttonId, cst)
// 		prev = cst
// 	}

// 	// ScrollView
// 	scrollView := NewScrollView(p.Get(scrollId))
// 	contentView := NewTextField(scrollView.ContentView)
// 	scrollView.ContentView = contentView

// 	// Layout ScrollView??

// 	// Root
// 	root.BotEqual(text.Bot().Add(constraint.Const(10)))
// 	l.Solve(nil, root)

// 	return n
// }

// func (v *TodoView) UpdateLayout(p Layouter) Layouter {
// 	l := &constraint.NewLayouter()

// 	root := l.Add(nil)
// 	root.HeightEqual(constraint.Const(40))
// 	root.Equal(l.Max())
// 	root.Solve()

// 	label := l.Add(labelId)
// 	label.BotLess(root.Bot())
// 	label.Equal(root)
// 	label.Solve()

// 	list := l.Add(root)
// 	list.TopEqual(label.Top())
// 	list.BotLess(root.Bot())
// 	list.Equal(root)
// 	list.Solve()

// 	text := l.Add(root)
// 	text.TopEqual(list.Top())
// 	text.BotLess(root.Bot())
// 	text.Equal(root)
// 	text.Solve()

// 	root.BotEqual(text.Bot().Add(constraint.Const(10)))
// 	root.Solve()
// 	return l
// }

// Handlers usually have no state, but in some cases they may need to (ie gesture recognizers need to build up state)
// Do layouts need to hold onto state? Do we want to support multiple layouters?

// type ScrollView struct{
// 	headerView View
// 	contentView View
// 	footerView View
// }

// func (v *ScrollView) Update(ctx *UpdateContext) map[string]View {
// 	c := map[string]View{}
// 	c["header"] = v.HeaderViewFunc()
// 	c["content"] = v.ContentViewFunc()
// 	c["footer"] = v.FooterViewFunc()
// 	return c
// }

// func (v *ScrollView) Layout(ctx *LayoutContext) Guide map[string]Guide {

// }

// type TableView struct{
// 	headerString string
// 	contentString string
// 	footerString string

// 	scrollNeedsUpdate bool
// }

// func (v *TableView) SetHeaderString(s string) {
// 	v.headerString = s
// 	v.headerNeedsUpdate = true
// 	v.NeedsUpdate()
// }

// func (v *TableView) SetFooterString(s string) {
// 	v.footerString = s
// 	v.NeedsUpdate()
// }

// func (v *TableView) SetContentString(s string) {
// 	v.footerString = s
// 	v.NeedsUpdate()

// 	scrollNeedsUpdate = true
// }

// func (v *TableView) Update(ctx *UpdateContext) map[string]View {
// 	c := make(map[string]View)

// 	scroll, new := NewScrollView(ctx.children["scroll"])
// 	scroll.SetHeaderViewFunc(func(prev View) View {
// 		t := NewTextField(View)
// 		if v.headerNeedsUpdate {
// 			v.headerNeedsUpdate = false
// 			t.SetString(headerString)
// 		}
// 		return t
// 	})

// 	scroll := NewScrollView(ctx.children["scroll"])
// 	if scrollNeedsUpdate {
// 		t := NewTextField(scroll.HeaderContext().children["header"])
// 		t.SetString(headerString)
// 		t.SetFont(font.System(12))
// 		t.SetFontColor(color.Red())
// 		scroll.SetHeaderView(t)
// 		scrollNeedsUpdate = false
// 	}
// 	c["scroll"] = scroll

// 	scroll := NewScrollView(ctx.children["scroll"])
// 	t, new := NewTextField(scroll.HeaderContext().children["header"])
// 	t.SetString(headerString)
// 	scroll.SetHeaderView(t)
// 	ctx.autoPropogateUpdate["scroll"] = dirty

// 	scroll := NewScrollView(ctx.children["scroll"])
// 	t := NewTextField(scroll.HeaderContext().children["header"])
// 	t.String = headerString
// 	t.Font = font.System(12)
// 	t.FontColor = color.Red()
// 	scroll.HeaderView = t
// 	ctx.AutoUpdate["scroll"] = dirty

// 	// if scrollDirty || new {
// 	// 	scroll.SetHeaderViewFunc(func(prev View) View {
// 	// 		t := NewTextField(View)
// 	// 		t.SetString(headerString)
// 	// 		return t
// 	// 	})
// 	// }

// 	scroll.SetContentViewFunc(func(prev View) View {
// 		t, new := NewTextField(View)
// 		if new || v.headerNeedsUpdate {
// 			v.headerNeedsUpdate = false
// 			t.SetString(footerString)
// 		}
// 		return t
// 	})
// 	c["scroll"] := NewScrollView(ctx.children["scroll"], scrollDirty, func(ScrollView *s){
// 		scroll.SetHeaderViewFunc(func(prev View) View {
// 			t := NewTextField(View)
// 			t.SetString(headerString)
// 			return t
// 		})
// 		s.SetFooterViewFunc(func(prev View) View {
// 			t := NewTextField(View)
// 			t.SetString(contentString)
// 			return t
// 		})
// 	})

// 	footerUpdater.Clean(func() {
// 		scroll.SetFooterViewFunc(func(prev View) View {
// 			t := NewTextField(View)
// 			t.SetString(contentString)
// 			return t
// 		})
// 	})
// 	c["scroll"] = scroll

// 	return c
// }

// func (v *TextField) Init() *TextField {
// 	if v == nil {
// 		return &TextField{}
// 	}
// 	return v
// }

// func (v *TableView) Update() map[string]View {
// 	c := make(map[string]View)

// 	textField = NewTextField(p["1"])

// 	//

// 	textField, _ := p["1"].(*TextField)
// 	textField = textField.Init()
// 	c["1"] = textField

// 	// v := p.Get("1", &TextField{}).(*TextField)
// 	// c["1"] = textField

// 	// textField := p["1"].extract(&TextField{})
// 	textField := &TextField{}.Copy(p["1"])

// 	// textField := p["1"].(*TextField) ?: &TextField{}
// 	// c["1"] = textField

// 	// var textField TextField = &TextField{}
// 	// if textField, ok := p["1"].(*TextField); !ok {
// 	// 	textField = &TextField{}
// 	// }
// 	// c["1"] = textField

// 	textField, ok := p["1"].(*TextField)
// 	if !ok {
// 		textField = &TextField{}
// 	}
// 	c["1"] = textField

// 	// if textField, ok := p["1"].(*TextField); !ok { textField = &TextField{} }
// 	// c["1"] = textField

// 	scroll = &ScrollView{
// 		Children: c,
// 		Layouter: &mochi.TableLayouter{},
// 	}

// 	return map[string]View{"scroll": scroll}
// }

// type ComplexView struct {}
// func (v *ComplexView) Update(ctx *UpdateContext) map[string]View {
// 	c := make(map[string]View)
// 	l := &mochi.AbsoluteLayouter{}

// 	circle, _ := ctx.children["1"].(*PathView)
// 	circle = circle.Init()
// 	circle.Path = Circle(0, 0, 10)
// 	c["1"] = circle
// 	// l["1"] = Guide{
// 	// 	Frame: Fr(0, 0, 0, 0),
// 	// 	Insets: In(0, 0, 0, 0),
// 	// }

// 	tri, _ := ctx.children["2"].(*PathView)
// 	tri = tri.Init()
// 	tri.Path = Triangle(0, 2, 4)
// 	tri.Color = color.Red()
// 	tri.BorderWidth = 4
// 	c["2"] = tri
// 	// l["2"] = Guide{
// 	// 	Frame: Fr(0, 0, 0, 0)
// 	// }

// 	v.eventHandlers = []Event{
// 		onClick{ func(){v.highlight = !v.highlight; v.NeedsUpdate()} }
// 	}

// 	v.layouter = l
// 	return c
// }

// func (v *TableView) Update(ctx *UpdateContext) map[string]View {
// 	c := map[string]View{}

// 	scroll := NewScrollView(ctx.Pop("scroll"))
// 	sc := map[string]View{}
// 	{
// 		textField := NewTextField(ctx.Pop("1"))
// 		textField.SetString("apple")
// 		sc["1"] = textField

// 		textField := NewTextField(ctx.Pop("2"))
// 		textField.SetString("baby")
// 		sc["2"] = textField

// 		textField := NewTextField(ctx.Pop("3"))
// 		textField.SetString("cabbage")
// 		sc["3"] = textField

// 		textField := NewTextField(ctx.Pop("4"))
// 		textField.SetString("doggo")
// 		sc["4"] = textField
// 	}
// 	scroll.SetContentView(sc)
// 	c["scroll"] = scroll

// 	return c
// }

// type ComplexView struct {}
// func (v *ComplexView) Update() map[string]View {
// 	c := make(map[string]View)
// 	l := &mochi.AbsoluteLayouter{}

// 	c["1"] = &PathView{
// 		path: Circle(0, 0, 10)
// 	}
// 	l.guides["1"] = Guide{
// 		Frame: Fr(0, 0, 0, 0),
// 		Insets: In(0, 0, 0, 0),
// 	}

// 	c["2"] = &PathView{
// 		Path: Triangle(0, 2, 4),
// 		Color: color.Red(),
// 		BorderWidth: 4,
// 	}
// 	l.guides["2"] = Guide{
// 		Frame: Fr(0, 0, 0, 0)
// 	}

// 	v.layouter = l
// 	return c
// }

// 	l.ConstrainChild(labelId, []Constraint{
// 		constraint.TopEq(l.Parent().Top()),
// 		constraint.BotLess(l.Parent().Bot())
// 		constraint.RightEq(l.Parent().Right())
// 		constraint.LeftEq(l.Parent().Right())
// 	})

// 	l.ConstrainChild(listId, func(s *constraint.Solver) {
// 		s.TopEq = l.Child(labelId).Top()
// 		s.BotLess = l.Parent().Bot()
// 		s.RightEq = l.Parent().Right()
// 		s.LeftEq = l.Parent().Left()

// 		return &constraint.Solver{
// 			.TopEq: l.Child(labelId).Top(),
// 			.BotLess: l.Parent().Bot()
// 			.RightEq: l.Parent().Right()
// 			.LeftEq: l.Parent().Left()
// 		}
// 	})

// 	l.ConstrainChild(labelId, []Constraint{
// 		layout.Eq{layout.Top, l.Guide(labelId).Top() - 5}
// 		layout.TopEqual(l.Guide(labelId).Top() - 5),
// 		layout.BotInsetLess(l.Guide(labelId).Top() - 5),
// 	})

// 	l.Constrain(labelId, []Constraint{
// 		layout.TopEqTop("", 0),
// 		layout.BotLsBot("", 0)})

// 	l.Constrain(labelId, []Constraint{
// 		layout.Top().Eq().Top(labelId)
// 	})

// 	l.Constrain(labelId, []Constraint{
// 		layout.ITeqIB("labelId", 0),
// 		layout.BotInsetLessBottom("", 0)})

// 	l.Constrain(listId, []Constraint{
// 		{lay.TopIn, lay.Eq, lay.TopIn, "", 0},
// 		{lay.BotIn, lay.Ls, lay.BotIn, "", 0},
// 		{lay.LftIn, lay.Eq, lay.LftIn, "", 0},
// 		{lay.RgtIn, lay.Eq, lay.RgtIn, "", 0}})

// 	l.Constrain(buttonId, []Constraint{
// 		layout.Eq(layout.TopIn, layout.TopIn, listId),
// 		layout.Ls(layout.BotIn, layout.BotIn, "")}

// 	return l
// }

// func (v *Todo) Update(ctx *UpdateContext) map[string]View {
// 	c := make(map[string]View)

// 	label := NewLabel(ctx.children[labelId])
// 	label.Text = "TODO"
// 	c[labelId] = llabel

// 	list := NewList(ctx.children[listId])
// 	list.Items = v.Items
// 	c[listId] = list

// 	text := NewTextField(ctx.children[textFieldId])
// 	text.OnChange = func(str string) {
// 		v.Input = str
// 	}
// 	c[textFieldId] = text

// 	button := NewButton(ctx.children[buttonId])
// 	button.OnClick = func() {
// 		append(v.Items, v.Input)
// 		v.NeedsUpdate()
// 	}
// 	c[buttonId] = button
// }

// func (v *Todo) Layout(ctx *LayoutContext) (Guide, map[string]Guide) {
// 	g := Guide{Frame: Rect{Size: ctx.MinSize}}
// 	c := map[string]Guide{}

// 	c[labelId] = ConstrainChild(ctx.children[labelId], Insets{}, []Constraint{
// 		{Top, Equal, g.Top()},
// 		{Bottom, Less, g.Bottom()},
// 		{Left, Equal, g.Left()},
// 		{Right, Equal, g.Right()},
// 	})

// 	c[listId] = ConstrainChild(ctx.children[listId], Insets{}, []Constraint{
// 		{Top, Equal, c[labelId].Top()},
// 		{Bottom, Less, g.Bottom()},
// 		{Left, Equal, g.Left()},
// 		{Right, Equal, g.Right()},
// 	})

// 	c[textFieldId] = ConstrainChild(ctx.children[textFieldId], Insets{}, []Constraint{
// 		{Top, Equal, c[listId].Top()},
// 		{Bottom, Less, g.Bottom()},
// 		{Left, Equal, g.Left()},
// 		{Right, Equal, g.Right()},
// 	})

// 	c[buttonId] = ConstrainChild(ctx.children[buttonId], Insets{}, []Constraint{
// 		{Top, Equal, c[textFieldId].Top()},
// 		{Bottom, Less, g.Bottom()},
// 		{Left, Equal, g.Left()},
// 		{Right, Equal, g.Right()},
// 	})
// 	return g, chl
// }

// <todo>
//   <h3>TODO</h3>

//   <ul>
//     <li each={ item, i in items }>{ item }</li>
//   </ul>

//   <form onsubmit={ handleSubmit }>
//     <input ref="input">
//     <button>Add #{ items.length + 1 }</button>
//   </form>

//   this.items = []

//   handleSubmit(e) {
//     e.preventDefault()
//     var input = this.refs.input
//     this.items.push(input.value)
//     input.value = ''
//   }
// </todo>
