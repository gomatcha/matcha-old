package constraint

import (
	"fmt"
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view"
	"math"
)

type comparison int

const (
	equal comparison = iota
	greater
	less
)

func (c comparison) String() string {
	switch c {
	case equal:
		return "="
	case greater:
		return ">"
	case less:
		return "<"
	}
	return ""
}

type attribute int

const (
	leftAttr attribute = iota
	rightAttr
	topAttr
	bottomAttr
	widthAttr
	heightAttr
	centerXAttr
	centerYAttr
)

func (a attribute) String() string {
	switch a {
	case leftAttr:
		return "Left"
	case rightAttr:
		return "Right"
	case topAttr:
		return "Top"
	case bottomAttr:
		return "Bottom"
	case widthAttr:
		return "Width"
	case heightAttr:
		return "Height"
	case centerXAttr:
		return "CenterX"
	case centerYAttr:
		return "CenterY"
	}
	return ""
}

type Anchor struct {
	anchor anchor
}

func (a *Anchor) Add(v float64) *Anchor {
	return &Anchor{
		offsetAnchor{
			offset:     v,
			underlying: a.anchor,
		},
	}
}

func (a *Anchor) Multiply(v float64) *Anchor {
	return &Anchor{
		multiplierAnchor{
			multiplier: v,
			underlying: a.anchor,
		},
	}
}

type anchor interface {
	value(*System) float64
}

type multiplierAnchor struct {
	multiplier float64
	underlying anchor
}

func (a multiplierAnchor) value(sys *System) float64 {
	return a.underlying.value(sys) * a.multiplier
}

type offsetAnchor struct {
	offset     float64
	underlying anchor
}

func (a offsetAnchor) value(sys *System) float64 {
	return a.underlying.value(sys) + a.offset
}

type constAnchor float64

func (a constAnchor) value(sys *System) float64 {
	return float64(a)
}

type notifierAnchor struct {
	n mochi.Float64Notifier
}

func (a notifierAnchor) value(sys *System) float64 {
	return a.n.Value()
}

type guideAnchor struct {
	guide     *Guide
	attribute attribute
}

func (a guideAnchor) value(sys *System) float64 {
	var g layout.Guide
	switch a.guide.id {
	case rootId:
		g = *sys.Guide.mochiGuide
	case minId:
		g = *sys.min.mochiGuide
	case maxId:
		g = *sys.max.mochiGuide
	default:
		g = *sys.children[a.guide.id].mochiGuide
	}

	// if g == nil {
	// 	return 0
	// }

	switch a.attribute {
	case leftAttr:
		return g.Left()
	case rightAttr:
		return g.Right()
	case topAttr:
		return g.Top()
	case bottomAttr:
		return g.Bottom()
	case widthAttr:
		return g.Width()
	case heightAttr:
		return g.Height()
	case centerXAttr:
		return g.CenterX()
	case centerYAttr:
		return g.CenterY()
	}
	return 0
}

func Const(f float64) *Anchor {
	return &Anchor{constAnchor(f)}
}

func Notifier(n mochi.Float64Notifier) *Anchor {
	return &Anchor{notifierAnchor{n}}
}

type Guide struct {
	id         mochi.Id
	system     *System
	children   map[mochi.Id]*Guide
	mochiGuide *layout.Guide
}

func (g *Guide) Top() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: topAttr}}
}

func (g *Guide) Right() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: rightAttr}}
}

func (g *Guide) Bottom() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: bottomAttr}}
}

func (g *Guide) Left() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: leftAttr}}
}

func (g *Guide) Width() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: widthAttr}}
}

func (g *Guide) Height() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: heightAttr}}
}

func (g *Guide) CenterX() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: centerXAttr}}
}

func (g *Guide) CenterY() *Anchor {
	return &Anchor{guideAnchor{guide: g, attribute: centerYAttr}}
}

func (g *Guide) Add(view view.View, solveFunc func(*Solver)) *Guide {
	id := view.Id()
	chl := &Guide{
		id:         id,
		system:     g.system,
		children:   map[mochi.Id]*Guide{},
		mochiGuide: nil,
	}
	s := &Solver{id: id}
	if solveFunc != nil {
		solveFunc(s)
	}
	g.children[id] = chl
	g.system.solvers = append(g.system.solvers, s)

	// Add any new notifier anchors to our notifier list.
	for _, i := range s.constraints {
		if anchor, ok := i.anchor.(notifierAnchor); ok {
			g.system.notifiers = append(g.system.notifiers, anchor.n)
		}
	}
	return chl
}

func (g *Guide) Solve(solveFunc func(*Solver)) {
	s := &Solver{id: g.id}
	if solveFunc != nil {
		solveFunc(s)
	}
	g.system.solvers = append(g.system.solvers, s)

	// Add any new notifier anchors to our notifier list.
	for _, i := range s.constraints {
		if anchor, ok := i.anchor.(notifierAnchor); ok {
			g.system.notifiers = append(g.system.notifiers, anchor.n)
		}
	}
}

type constraint struct {
	attribute  attribute
	comparison comparison
	anchor     anchor
}

func (c constraint) String() string {
	return fmt.Sprintf("%v%v%v", c.attribute, c.comparison, c.anchor)
}

type Solver struct {
	id          mochi.Id
	constraints []constraint
}

func (s *Solver) solve(sys *System, ctx *layout.Context) {
	cr := newConstrainedRect()

	for _, i := range s.constraints {
		copy := cr

		// Generate the range from constraint
		var r _range
		switch i.comparison {
		case equal:
			r = _range{min: i.anchor.value(sys), max: i.anchor.value(sys)}
		case greater:
			r = _range{min: i.anchor.value(sys), max: math.Inf(1)}
		case less:
			r = _range{min: math.Inf(-1), max: i.anchor.value(sys)}
		}

		// Update the solver
		switch i.attribute {
		case leftAttr:
			copy.left = copy.left.intersect(r)
		case rightAttr:
			copy.right = copy.right.intersect(r)
		case topAttr:
			copy.top = copy.top.intersect(r)
		case bottomAttr:
			copy.bottom = copy.bottom.intersect(r)
		case widthAttr:
			copy.width = copy.width.intersect(r)
		case heightAttr:
			copy.height = copy.height.intersect(r)
		case centerXAttr:
			copy.centerX = copy.centerX.intersect(r)
		case centerYAttr:
			copy.centerY = copy.centerY.intersect(r)
		}

		// Validate that the new system is well-formed. Otherwise ignore the changes.
		if !copy.isValid() {
			continue
		}
		cr = copy
	}

	// Get parent guide.
	var parent layout.Guide
	if s.id == rootId {
		parent = *sys.min.mochiGuide
	} else {
		parent = *sys.Guide.mochiGuide
	}

	// Solve for width & height.
	var width, height float64
	var g layout.Guide
	if s.id == rootId {
		g = layout.Guide{}
		width, _ = cr.solveWidth(parent.Width())
		height, _ = cr.solveHeight(parent.Height())
	} else {
		// Solve for width and height? Should we set cr to this?
		_, widthCR := cr.solveWidth(0)
		_, heightCR := cr.solveHeight(0)

		g = ctx.LayoutChild(s.id, layout.Pt(widthCR.width.min, heightCR.height.min), layout.Pt(widthCR.width.max, heightCR.height.max))
		width = g.Width()
		height = g.Height()

		if width < cr.width.min || height < cr.height.min || width > cr.width.max || height > cr.height.max {
			fmt.Printf("constraints: child guide is outside of bounds. Min:%v Max:%v Actual:%v", layout.Pt(cr.width.min, cr.height.min), layout.Pt(cr.width.max, cr.height.max), layout.Pt(width, height))
			width = cr.width.min
			height = cr.height.min
		}
	}

	// Solve for centerX & centerY using new width & height.
	cr.width = cr.width.intersect(_range{min: width, max: width})
	cr.height = cr.height.intersect(_range{min: height, max: height})
	if !cr.isValid() {
		panic("constraints: system inconsistency")
	}
	var centerX, centerY float64
	if s.id == rootId {
		centerX = width / 2
		centerY = height / 2
	} else {
		centerX, _ = cr.solveCenterX(parent.CenterX())
		centerY, _ = cr.solveCenterY(parent.CenterY())
	}

	// Set zIndex
	g.ZIndex = sys.zIndex
	sys.zIndex += 1

	// Update the guide and the system.
	g.Frame = layout.Rt(centerX-width/2, centerY-height/2, centerX+width/2, centerY+height/2)
	if s.id == rootId {
		sys.Guide.mochiGuide = &g
	} else {
		sys.Guide.children[s.id].mochiGuide = &g
	}
}

func (s *Solver) TopEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: topAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) TopLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: topAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) TopGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: topAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) RightEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: rightAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) RightLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: rightAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) RightGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: rightAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) BottomEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: bottomAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) BottomLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: bottomAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) BottomGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: bottomAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) LeftEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: leftAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) LeftLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: leftAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) LeftGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: leftAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) WidthEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: widthAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) WidthLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: widthAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) WidthGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: widthAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) HeightEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: heightAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) HeightLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: heightAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) HeightGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: heightAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) CenterXEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerXAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) CenterXLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerXAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) CenterXGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerXAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) CenterYEqual(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerYAttr, comparison: equal, anchor: a.anchor})
}

func (s *Solver) CenterYLess(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerYAttr, comparison: less, anchor: a.anchor})
}

func (s *Solver) CenterYGreater(a *Anchor) {
	s.constraints = append(s.constraints, constraint{attribute: centerYAttr, comparison: greater, anchor: a.anchor})
}

func (s *Solver) String() string {
	return fmt.Sprintf("Solver{%v, %v}", s.id, s.constraints)
}

type systemId int

const (
	rootId mochi.Id = -1 * iota
	minId
	maxId
)

type System struct {
	*Guide
	min            *Guide
	max            *Guide
	solvers        []*Solver
	zIndex         int
	notifiers      []mochi.Notifier
	batchNotifiers map[chan struct{}]*mochi.BatchNotifier
}

func New() *System {
	sys := &System{}
	sys.Guide = &Guide{id: rootId, system: sys, children: map[mochi.Id]*Guide{}}
	sys.min = &Guide{id: minId, system: sys, children: map[mochi.Id]*Guide{}}
	sys.max = &Guide{id: maxId, system: sys, children: map[mochi.Id]*Guide{}}
	sys.batchNotifiers = map[chan struct{}]*mochi.BatchNotifier{}
	return sys
}

func (sys *System) MinGuide() *Guide {
	return sys.min
}

func (sys *System) MaxGuide() *Guide {
	return sys.max
}

func (sys *System) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	sys.min.mochiGuide = &layout.Guide{
		Frame: layout.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y),
	}
	sys.max.mochiGuide = &layout.Guide{
		Frame: layout.Rt(0, 0, ctx.MaxSize.X, ctx.MaxSize.Y),
	}
	sys.Guide.mochiGuide = &layout.Guide{
		Frame: layout.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y),
	}
	// TODO(Kevin): reset all guides

	for _, i := range sys.solvers {
		i.solve(sys, ctx)
	}

	g := *sys.Guide.mochiGuide
	gs := map[mochi.Id]layout.Guide{}
	for k, v := range sys.Guide.children {
		gs[k] = *v.mochiGuide
	}
	return g, gs
}

func (sys *System) Notify() chan struct{} {
	n := mochi.NewBatchNotifier(sys.notifiers...)
	c := n.Notify()
	if c != nil {
		sys.batchNotifiers[c] = n
	}
	return c
}

func (sys *System) Unnotify(c chan struct{}) {
	if c == nil {
		return
	}
	n := sys.batchNotifiers[c]
	if n == nil {
		panic("Cannot unnotify unknown chan")
	}
	n.Unnotify(c)
	delete(sys.batchNotifiers, c)
}

type _range struct {
	min float64
	max float64
}

func (r _range) intersectMin(v float64) _range {
	r.min = math.Max(r.min, v)
	return r
}

func (r _range) intersectMax(v float64) _range {
	r.max = math.Min(r.max, v)
	return r
}

func (r _range) intersect(r2 _range) _range {
	return _range{min: math.Max(r.min, r2.min), max: math.Min(r.max, r2.max)}
}

func (r _range) isValid() bool {
	return r.max >= r.min
}

func (r _range) nearest(v float64) float64 {
	// return a sane value even if range is invalid
	if r.max < r.min {
		r.max, r.min = r.min, r.max
	}
	switch {
	case r.min == r.max:
		return r.min
	case r.min >= v:
		return r.min
	case r.max <= v:
		return r.max
	default:
		return v
	}
}

type constrainedRect struct {
	left, right, top, bottom, width, height, centerX, centerY _range
}

func newConstrainedRect() constrainedRect {
	all := _range{min: math.Inf(-1), max: math.Inf(1)}
	pos := _range{min: 0, max: math.Inf(1)}
	return constrainedRect{
		left: all, right: all, top: all, bottom: all, width: pos, height: pos, centerX: all, centerY: all,
	}
}

func (cr constrainedRect) isValid() bool {
	_, r1 := cr.solveWidth(0)
	_, r2 := cr.solveHeight(0)
	_, r3 := cr.solveCenterX(0)
	_, r4 := cr.solveCenterY(0)
	return r1.width.isValid() && r2.height.isValid() && r3.centerX.isValid() && r4.centerY.isValid()
}

func (r constrainedRect) solveWidth(b float64) (float64, constrainedRect) {
	centerXMax, centerXMin := r.centerX.max, r.centerX.min
	rightMax, rightMin := r.right.max, r.right.min
	leftMax, leftMin := r.left.max, r.left.min

	// Width = (Right - CenterX) * 2
	if !math.IsInf(centerXMin, 0) && !math.IsInf(rightMax, 0) {
		r.width = r.width.intersectMax((rightMax - centerXMin) * 2)
	}
	if !math.IsInf(centerXMax, 0) && !math.IsInf(rightMin, 0) {
		r.width = r.width.intersectMin((rightMin - centerXMax) * 2)
	}

	// Width = Right - Left
	if !math.IsInf(rightMax, 0) && !math.IsInf(leftMin, 0) {
		r.width = r.width.intersectMax(rightMax - leftMin)
	}
	if !math.IsInf(rightMin, 0) && !math.IsInf(leftMax, 0) {
		r.width = r.width.intersectMin(rightMin - leftMax)
	}

	// Width = (CenterX - Left) * 2
	if !math.IsInf(centerXMax, 0) && !math.IsInf(leftMin, 0) {
		r.width = r.width.intersectMax((centerXMax - leftMin) * 2)
	}
	if !math.IsInf(centerXMin, 0) && !math.IsInf(leftMax, 0) {
		r.width = r.width.intersectMin((centerXMin - leftMax) * 2)
	}

	return r.width.nearest(b), r
}

func (r constrainedRect) solveCenterX(b float64) (float64, constrainedRect) {
	rightMax, rightMin := r.right.max, r.right.min
	leftMax, leftMin := r.left.max, r.left.min
	widthMax, widthMin := r.width.max, r.width.min

	// CenterX = (Right + Left)/2
	if !math.IsInf(rightMax, 0) && !math.IsInf(leftMax, 0) {
		r.centerX = r.centerX.intersectMax((rightMax + leftMax) / 2)
	}
	if !math.IsInf(rightMin, 0) && !math.IsInf(leftMin, 0) {
		r.centerX = r.centerX.intersectMin((rightMin + leftMin) / 2)
	}

	// CenterX = Right - Width / 2
	if !math.IsInf(rightMax, 0) && !math.IsInf(widthMin, 0) {
		r.centerX = r.centerX.intersectMax(rightMax - widthMin/2)
	}
	if !math.IsInf(rightMin, 0) && !math.IsInf(widthMax, 0) {
		r.centerX = r.centerX.intersectMin(rightMin - widthMax/2)
	}

	// CenterX = Left + Width / 2
	if !math.IsInf(leftMax, 0) && !math.IsInf(widthMax, 0) {
		r.centerX = r.centerX.intersectMax(leftMax + widthMax/2)
	}
	if !math.IsInf(leftMin, 0) && !math.IsInf(widthMin, 0) {
		r.centerX = r.centerX.intersectMin(leftMin + widthMin/2)
	}

	return r.centerX.nearest(b), r
}

func (r constrainedRect) solveHeight(b float64) (float64, constrainedRect) {
	centerYMax, centerYMin := r.centerY.max, r.centerY.min
	bottomMax, bottomMin := r.bottom.max, r.bottom.min
	topMax, topMin := r.top.max, r.top.min

	// height = (bottom - centerY) * 2
	if !math.IsInf(centerYMin, 0) && !math.IsInf(bottomMax, 0) {
		r.height = r.height.intersectMax((bottomMax - centerYMin) * 2)
	}
	if !math.IsInf(centerYMax, 0) && !math.IsInf(bottomMin, 0) {
		r.height = r.height.intersectMin((bottomMin - centerYMax) * 2)
	}

	// height = bottom - top
	if !math.IsInf(bottomMax, 0) && !math.IsInf(topMin, 0) {
		r.height = r.height.intersectMax(bottomMax - topMin)
	}
	if !math.IsInf(bottomMin, 0) && !math.IsInf(topMax, 0) {
		r.height = r.height.intersectMin(bottomMin - topMax)
	}

	// height = (centerY - top) * 2
	if !math.IsInf(centerYMax, 0) && !math.IsInf(topMin, 0) {
		r.height = r.height.intersectMax((centerYMax - topMin) * 2)
	}
	if !math.IsInf(centerYMin, 0) && !math.IsInf(topMax, 0) {
		r.height = r.height.intersectMin((centerYMin - topMax) * 2)
	}

	return r.height.nearest(b), r
}

func (r constrainedRect) solveCenterY(b float64) (float64, constrainedRect) {
	bottomMax, bottomMin := r.bottom.max, r.bottom.min
	topMax, topMin := r.top.max, r.top.min
	heightMax, heightMin := r.height.max, r.height.min

	// centerY = (bottom + top)/2
	if !math.IsInf(bottomMax, 0) && !math.IsInf(topMax, 0) {
		r.centerY = r.centerY.intersectMax((bottomMax + topMax) / 2)
	}
	if !math.IsInf(bottomMin, 0) && !math.IsInf(topMin, 0) {
		r.centerY = r.centerY.intersectMin((bottomMin + topMin) / 2)
	}

	// centerY = bottom - height / 2
	if !math.IsInf(bottomMax, 0) && !math.IsInf(heightMin, 0) {
		r.centerY = r.centerY.intersectMax(bottomMax - heightMin/2)
	}
	if !math.IsInf(bottomMin, 0) && !math.IsInf(heightMax, 0) {
		r.centerY = r.centerY.intersectMin(bottomMin - heightMax/2)
	}

	// centerY = top + height / 2
	if !math.IsInf(topMax, 0) && !math.IsInf(heightMax, 0) {
		r.centerY = r.centerY.intersectMax(topMax + heightMax/2)
	}
	if !math.IsInf(topMin, 0) && !math.IsInf(heightMin, 0) {
		r.centerY = r.centerY.intersectMin(topMin + heightMin/2)
	}

	return r.centerY.nearest(b), r
}

func (r constrainedRect) String() string {
	return fmt.Sprintf("{left:%v, right:%v, top:%v, bottom:%v, width:%v, height:%v, centerX:%v, centerY:%v}", r.left, r.right, r.top, r.bottom, r.width, r.height, r.centerX, r.centerY)
}
