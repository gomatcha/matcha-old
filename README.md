# Matcha - A cross platform mobile UI framework in Go

## Goals
* It should be straight forward to build general purpose UIs
	* Implementing swipe to delete and animated table views should not require complex logic
* Composable components over inheritance hierarchies
	* This is just good design. Components can evolve over time independently.
* Conceptually consistent
* Follow Go idioms while pushing a common set of idioms ourselves
	* There should be one right way to do things.
* Favor conciseness over extra hooks
	* This seems to be the general trend in popular frameworks
	* Most of the hooks I leave myself I never use.
* System desgin should be visible in the API. Minimize action at a distance
	* Using the API should help you learn how the framework works, rather than hide it away
* Scalability
	* Support large programs with large numbers of dependencies, with large teams of programmers working on them.

## Example

```go
type TodoView struct {
	*Embed
	Items []string
	Input string
}

func NewTodoView(ctx *view.Context, key string) *TodoView {
	if v, ok := ctx.Prev(key).(*TodoView); ok {
		return v
	}
	return &TodoView{
		Embed:  view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *TodoView) Update(p *Node) *Node {
	l := &constraint.New()
	n := &Node{}
	n.layouter = l

	var prev *constraint.Guide
	{
		// Label
		chl := NewLabel(ctx, "title")
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
		chl := NewTextField(p.Get(textFieldId))
		chl.Input = v.Input
		chl.OnChange = func(str string) {
			v.Input = str
			v.NeedsUpdate()
		}
		n.Set(textFieldId, chl)

		prev = l.Add(textFieldId, func(constraint.Solver *s){
			s.TopEqual(prev.Bot())
			s.BotLess(l.Bot())
		})
	}
	{
		// Button
		chl := NewButton(p.Get(buttonId))
		chl.OnClick = func() {
			if v.Input == "" {
				return
			}
			append(v.Items, v.Input)
			v.Input = ""
			v.NeedsUpdate()
		}
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
```

Similar Libraries

* https://github.com/murlokswarm/app
* https://github.com/andlabs/ui
* https://github.com/golang/exp/tree/master/shiny
* https://github.com/cztomczak/cef2go
* https://github.com/google/gxui
* https://github.com/go-ui/ui
* https://sciter.com