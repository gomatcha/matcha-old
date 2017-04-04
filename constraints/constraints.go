package constraints

import (
	"github.com/overcyn/mochi"
	"math"
)

type comparison int

const (
	equal comparison = iota
	greater
	less
)

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

type Anchor struct {
	anchor anchor
}

func (a Anchor) Add(v float64) Anchor {
	return Anchor{
		offsetAnchor{
			offset:     v,
			underlying: a.anchor,
		},
	}
}

func (a Anchor) Multiply(v float64) Anchor {
	return Anchor{
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

type guideAnchor struct {
	guide     *Guide
	attribute attribute
}

func (a guideAnchor) value(sys *System) float64 {
	var g mochi.Guide
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

type Guide struct {
	id         interface{}
	system     *System
	children   map[interface{}]*Guide
	mochiGuide *mochi.Guide
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

func (g *Guide) AddGuide(id interface{}, solveFunc func(*Solver)) *Guide {
	chl := &Guide{
		id:         id,
		system:     g.system,
		children:   map[interface{}]*Guide{},
		mochiGuide: nil,
	}
	s := &Solver{id: id}
	if solveFunc != nil {
		solveFunc(s)
	}
	g.children[id] = chl
	g.system.solvers = append(g.system.solvers, s)
	return chl
}

type constraint struct {
	attribute  attribute
	comparison comparison
	anchor     anchor
}

type Solver struct {
	id          interface{}
	constraints []constraint
}

func (s *Solver) solve(sys *System, ctx *mochi.LayoutContext) {
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
	var parent mochi.Guide
	if s.id == rootId {
		parent = *sys.max.mochiGuide
	} else {
		parent = *sys.Guide.mochiGuide
	}

	// Solve for width & height.
	var width, height float64
	var g mochi.Guide
	if s.id == rootId {
		g = mochi.Guide{}
		width, _ = cr.solveWidth(parent.Width())
		height, _ = cr.solveHeight(parent.Height())
	} else {
		g = ctx.LayoutChild(s.id, mochi.Pt(cr.width.min, cr.height.min), mochi.Pt(cr.width.max, cr.height.max))
		width = g.Width()
		height = g.Height()
	}

	// Solve for centerX & centerY using new width & height.
	cr.width = cr.width.intersect(_range{min: width, max: width})
	cr.height = cr.height.intersect(_range{min: height, max: height})
	if !cr.isValid() {
		panic("constraints: system inconsistency")
	}
	centerX, _ := cr.solveCenterX(parent.CenterY())
	centerY, _ := cr.solveCenterY(parent.CenterY())

	// Update the guide and the system.
	g.Frame = mochi.Rt(centerX-width/2, centerY-height/2, centerX+width/2, centerY+height/2)
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

type systemId int

const (
	rootId systemId = iota
	minId
	maxId
)

type System struct {
	*Guide
	min     *Guide
	max     *Guide
	solvers []*Solver
}

func New() *System {
	sys := &System{}
	sys.min = &Guide{id: minId, system: sys}
	sys.max = &Guide{id: maxId, system: sys}
	sys.Guide = &Guide{id: rootId, system: sys}
	return sys
}

func (sys *System) MinGuide() *Guide {
	return sys.min
}

func (sys *System) MaxGuide() *Guide {
	return sys.max
}

func (sys *System) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	// TODO(Kevin): reset all guides

	for _, i := range sys.solvers {
		i.solve(sys, ctx)
	}

	g := *sys.Guide.mochiGuide
	gs := map[interface{}]mochi.Guide{}
	for k, v := range sys.Guide.children {
		gs[k] = *v.mochiGuide
	}
	return g, gs
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
	return r.max < r.min
}

func (r _range) nearest(v float64) float64 {
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

func (r constrainedRect) isValid() bool {
	_, ok := r.solveWidth(0)
	_, ok2 := r.solveHeight(0)
	_, ok3 := r.solveCenterX(0)
	_, ok4 := r.solveCenterY(0)
	return ok && ok2 && ok3 && ok4
}

func (r constrainedRect) solveWidth(b float64) (float64, bool) {
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

	return r.width.nearest(b), r.width.isValid()
}

func (r constrainedRect) solveCenterX(b float64) (float64, bool) {
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

	return r.centerX.nearest(b), r.centerX.isValid()
}

func (r constrainedRect) solveHeight(b float64) (float64, bool) {
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

	return r.height.nearest(b), r.height.isValid()
}

func (r constrainedRect) solveCenterY(b float64) (float64, bool) {
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

	return r.centerY.nearest(b), r.centerY.isValid()
}
