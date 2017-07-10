package todo

import (
	"image/color"
	"strconv"

	"golang.org/x/image/colornames"

	"gomatcha.io/bridge"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/layout/table"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/basicview"
	"gomatcha.io/matcha/view/scrollview"
	"gomatcha.io/matcha/view/textview"
)

func init() {
	bridge.RegisterFunc("gomatcha.io/matcha/examples/todo New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {
			app := NewAppView(ctx, "")
			app.Todos = []*Todo{
				&Todo{Title: "Title1"},
				&Todo{Title: "Title2"},
				&Todo{Title: "Title3"},
			}
			return app
		}))
	})
}

type Todo struct {
	Title     string
	Completed bool
}

type AppView struct {
	*view.Embed
	Todos []*Todo
}

func NewAppView(ctx *view.Context, key string) *AppView {
	if v, ok := ctx.Prev(key).(*AppView); ok {
		return v
	}
	return &AppView{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *AppView) Build(ctx *view.Context) *view.Model {
	l := &table.Layouter{}

	for idx, todo := range v.Todos {
		todoView := NewTodoView(ctx, strconv.Itoa(idx))
		todoView.Todo = todo
		todoView.OnDelete = func() {
			v.Todos = append(v.Todos[:idx], v.Todos[idx+1:]...)
			v.Signal()
		}
		todoView.OnComplete = func(complete bool) {
			v.Todos[idx].Completed = complete
			v.Signal()
		}
		l.Add(todoView, nil)
	}

	scrollView := scrollview.New(ctx, "scrollView")
	scrollView.ContentChildren = l.Views()
	scrollView.ContentLayouter = l
	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: colornames.White},
	}
}

type TodoView struct {
	*view.Embed
	Todo       *Todo
	OnDelete   func()
	OnComplete func(check bool)
}

func NewTodoView(ctx *view.Context, key string) *TodoView {
	if v, ok := ctx.Prev(key).(*TodoView); ok {
		return v
	}
	return &TodoView{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *TodoView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.Height(50)
		s.WidthEqual(l.MaxGuide().Width())
	})

	checkbox := NewCheckbox(ctx, "checkbox")
	checkbox.Value = v.Todo.Completed
	checkbox.OnValueChange = func(value bool) {
		if v.OnComplete != nil {
			v.OnComplete(value)
		}
	}
	checkboxGuide := l.Add(checkbox, func(s *constraint.Solver) {
		s.CenterYEqual(l.CenterY())
		s.LeftEqual(l.Left().Add(15))
	})

	deleteButton := NewDeleteButton(ctx, "delete")
	deleteButton.OnPress = func() {
		if v.OnDelete != nil {
			v.OnDelete()
		}
	}
	deleteGuide := l.Add(deleteButton, func(s *constraint.Solver) {
		s.CenterYEqual(l.CenterY())
		s.RightEqual(l.Right().Add(-15))
	})

	titleView := textview.New(ctx, "title")
	titleView.String = v.Todo.Title
	titleView.Style = nil //...
	l.Add(titleView, func(s *constraint.Solver) {
		s.CenterYEqual(l.CenterY())
		s.LeftEqual(checkboxGuide.Right().Add(15))
		s.RightEqual(deleteGuide.Left().Add(-15))
	})

	separator := basicview.New(ctx, "separator")
	separator.Painter = &paint.Style{BackgroundColor: color.RGBA{203, 202, 207, 255}}
	l.Add(separator, func(s *constraint.Solver) {
		s.Height(1)
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
		s.BottomEqual(l.Bottom())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
	}
}

type Checkbox struct {
	*view.Embed
	Value         bool
	OnValueChange func(value bool)
}

func NewCheckbox(ctx *view.Context, key string) *Checkbox {
	if v, ok := ctx.Prev(key).(*Checkbox); ok {
		return v
	}
	return &Checkbox{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *Checkbox) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.Width(25)
		s.Height(25)
	})

	return &view.Model{
		Painter:  &paint.Style{BackgroundColor: colornames.Red},
		Layouter: l,
	}
}

type DeleteButton struct {
	*view.Embed
	OnPress func()
}

func NewDeleteButton(ctx *view.Context, key string) *DeleteButton {
	if v, ok := ctx.Prev(key).(*DeleteButton); ok {
		return v
	}
	return &DeleteButton{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *DeleteButton) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.Width(25)
		s.Height(25)
	})

	return &view.Model{
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
		Layouter: l,
	}
}
