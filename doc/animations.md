* Inspiration
Flutter https://medium.com/dartlang/zero-to-one-with-flutter-43b13fd7b354
https://flutter.io/animations/
ReactNative
Pop
https://github.com/matthewcheok/POP-MCAnimate
* Hardware acceleration? https://www.sitepoint.com/introduction-to-hardware-acceleration-css-animations/



Animation Curve. Protocol that maps doubles nominally in the range 0.0-1.0 to doubles nominally in the range 0.0-1.0
Tween. Maps a double value nominally in the range 0.0-1.0 to a typed value (e.g. a Color, or another double). 

Chaining and composing animations
Playing animations in parell or sequence.

Animation.Value{}.WithCurve()

PanGestureRecognizer.OnUpdate(func(x, y int){
    value.Set(p.X)
})

Needs to update the layouter and the painter.
Animations should be serializable so they don't have to cross the bridge as often.

type Animation interface {
    Value()
}

type Animatable interface {
    SetValue(float64)
}

chl1 := basicview.New(ctx.Get(chl1id))
chl1.PaintOptions.BackgroundColor = mochi.RedColor
n.Set(chl1id, chl1)
g1 := l.Add(chl1id, func(s *constraint.Solver) {
    s.WidthEqual(constraint.Animater(a))
    s.HeightEqual(constraint.Animater(a))
})
-

chl1.Painter = mochi.PaintOptions{
    BackgroundColor: mochi.RedColor `mochi:animatable`
}

chl1.Painter = Painter.WithAnimater(func (p *Painter){
    return p
})

Anything that only modifies the layout and the painter in a predefined manner can be animated.
The whole point is to keep animations separate from the "main thread" so the main thread doesn't cause lag.
We could unify everything, but that removes the performance benefit of async painting and layout.

Painter
Animatable<Float>

Animate.Float
Animate.Int
Animate.Image
Animate.Interface
Animate.Size
Animate.Rect
Animate.Transform
Animate.Point
Animate.Shadow
Animate.Path
Animate.Font
Animate.Layouter
Animate.Painter
Animate.Date

type animate.Float interface {
    Value() float
    Update() <-chan bool
}

type animate.Image interface {
    Value() image.Image
}

Image {
    AnimatedImage animate.Image
}

type ImageView struct {
    AnimatedImage chan-> image.Image
}
<!-- type Value interface {
    float->
    Value() float
    // some sort of way to listen.
}
 -->
animate.NewSource(duration: 0)

type animate.Curve interface {
}

Source().WithCurveFunc(EaseIn.Curve)

func EaseIn.Curve(in float64) float64
func EaseOut.Curve(in float64) float64
func EaseInOut.Curve(in float64) float64

for {
switch {
case <-update:
    
}
}