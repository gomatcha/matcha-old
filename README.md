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
<!--
## Blah
### 

func (v *TodoView) Update(p *Node) *Node {
	n := &Node{}

	label := NewLabel(p.Get(labelId))
	label.Text = "TODO"
	n.Add(labelId, label)

	list := NewList(p.Get(listId))
	list.Items = v.Items
	n.Add(listID, list)

	text := NewTextField(p.Get(textFieldId))
	text.Input = v.Input
	text.OnChange = func(str string) {
		v.Input = str
		v.NeedsUpdate()
	}
	n.Add(textFieldId, textField)

	button := NewButton(p.Get(buttonId))
	button.OnClick = func() {
		if v.Input == "" {
			return
		}
		append(v.Items, v.Input)
		v.Input = ""
		v.NeedsUpdate()
	}
	n.Add(textFieldId, textField)
	scrollView := NewScrollView(p.Get(scrollId))
	contentView := NewTextField(scrollView.ContentView)
	scrollView.ContentView = contentView
}

## Drawing

What is the minimum api necessary for drawing? 
* Groups
* Shapes
* Shadow
* Fill Color
* Gradient
* Mask

## Layout

Layout should happen on a background thread. Parent always knows where the child is. Does this include 3d transforms? Rotations?

Ignore Transforms and rotations for now.

## Event 

Mouse, keyboard and touch input is handled by event handlers attached to each view. Events start at the handler deepest in the view hierarchy. Handlers are given an option to bubble the event further upwards or capture it. Multi-touch events will behave similar to UIGestureRecognizer.

Once a gesture recognizer recognizes a view, it should, start capturing events on the way down. 

What is the purpose of sending events down before going up? We can do UserInteractionEnabled easily. Once a gesture recognizer has begun recognizing an event, it can prevent other recognizers from accidentally triggering.

A event handlers should be able to track all input regardless of position. And event handlers should be able to cancel other event handlers.

### UIGestureRecognizer

iOS has a great API in UIGestureRecognizer. It does have the complexity of `canBePreventedByGestureRecognizer:` and `canPreventGestureRecognizer:`. Is there a way to do this in a declarative manner? We could give each event handler a `priority` value. Or we could refer to other gesture recognizers by a keypath. How does this work with the view tree? Is there a way we could reduce the scope of UIGestureRecognizer, to give us wins in other areas?

### Use Case: Double Tap and Single Tap

In Safari, double tap zoom the page while a single tap opens a link. Even if the single tap event handler is activated, it must wait for the double tap handler to verify that it is complete. How do we choose which handler to prefer? We could base it on the one that took a longer time to respond. Are there other options we could use to determine the winner? One possibility is we use the eventHandler array order to hint at priority?

### Use Case: Tap Drag and Drag Tap

Similar to the double Tap and single Tap in Safari, you could imagine event handlers initiating at separate times. Again how do we determine the winner?

### Use Case: Button inside Button vs Button in ScrollView

If you have a button inside of another touchable area. The inner button should take priority. However if you have a button in a scrollView, and the scroll drags then the scroll should take priority.

Alternately we could use a signal outside of the event system to cancel the button press. Or the eventHandlers could mediate between themselves? 

Press and Hold =  press -> hold
Double tap = press -> release -> press -> release
Button = press -> hold / drag -> release
Scroll = press -> hold / drag

- ScrollView : Scroll, Double Tap
	- Button : Button, Press and Hold

Scenario 1: press -> release. Only the button will be triggered. But double tap is waiting to trigger?
Scenario 2: press -> hold. Press and Hold will be triggered. Button will be waiting to trigger.
We can give Press and Hold priority by ordering them.
Scenario 3: press -> drag. Scroll will be triggered.

It seems that we are waiting for all gesture recognizers to get out of the possible state. At which point the first to complete is the winner.

Possible -> Began -> Ended/Cancelled

### Other Use Cases
* UserInteractionEnabled = False
* Scrolling
* Button activation cancelled by vertical scrolling but not horizontal
* Pinch to zoom
* Highlighting a view and dragging outside of it and back in.
* Horizontal swipe on tableview to show delete button
* Touch driven animations. AKA swipe back to navigate.

How do I do this in an abstract manner, that doesn't need built in support similar to our constraint system? Also note the gesture recognizer lower in the heirarchy wins.

We could wrap it the event, and rebubble. There needs to be some synchronization mechanism, so that when an event completes, it notifies the other gesture recognizers that they should cancel. The gesture recognizers need to intercept the events before anything else hit. GestureEvent separate from a touch event? 

GestureCompletedEvent {}

GestureEvent {
	Events []Event
	Possible bool
	Complete func()complete
}

GestureEvent flows through the system. If any are possible, then do nothing. When a gesture flows through and possible is still false, but complete is true. Call the completion() and send through a GestureCompleteEvent. 

## Animations

## Updating

What if we didn't need to call NeedsUpdate? We have it so that when a component modifies itself, it can trigger a rerender. Also to give opportunity to stop the update from flowing downwards.
Instead of calling setters, NeedsUpdate will automatically flow through the entire tree. We can stop by calling, DoesntNeedsUpdate()?

Rather than calling setState(). We instead mark v.NeedsUpdate(). And instead of passing in components, we assume you don't modify components except in the Update() func.-->
