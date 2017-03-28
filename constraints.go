package mochi

import (
	"math"
)

type Comparison int

const (
	Equal Comparison = iota
	Greater
	Less
)

type Attribute int

const (
	Left Attribute = iota
	Right
	Top
	Bottom
	Width
	Height
	CenterX
	CenterY
)

type Constraint struct {
	Attribute  Attribute
	Comparison Comparison
	Value      float64
}

func Constrain(ctx *LayoutContext, in Insets, c []Constraint) Guide {
	return ConstrainChild(ctx, nil, in, c)
}

func ConstrainChild(ctx *LayoutContext, key interface{}, in Insets, c []Constraint) Guide {
	solver := newConstraintSolver()
	for _, i := range c {
		copy := solver

		// Generate the range from constraint
		var r _range
		switch i.Comparison {
		case Equal:
			r = _range{min: i.Value, max: i.Value}
		case Greater:
			r = _range{min: i.Value, max: math.Inf(1)}
		case Less:
			r = _range{min: math.Inf(-1), max: i.Value}
		}

		// Update the solver
		switch i.Attribute {
		case Left:
			solver.left = solver.left.intersect(r)
		case Right:
			solver.right = solver.right.intersect(r)
		case Top:
			solver.top = solver.top.intersect(r)
		case Bottom:
			solver.bottom = solver.bottom.intersect(r)
		case Width:
			solver.width = solver.width.intersect(r)
		case Height:
			solver.height = solver.height.intersect(r)
		case CenterX:
			solver.centerX = solver.centerX.intersect(r)
		case CenterY:
			solver.centerY = solver.centerY.intersect(r)
		}

		// Solve and validate that the new system is well-formed. Otherwise ignore the changes.
		copy = copy.solve()
		if !copy.isValid() {
			continue
		}
		solver = copy
	}

	if key != nil {
		guide := ctx.LayoutChild(key, Sz(solver.width.min, solver.height.min), Sz(solver.width.max, solver.height.max))

		copy := solver
		copy.width = copy.width.intersect(_range{min: guide.Frame.Size.Width, max: guide.Frame.Size.Width})
		copy.height = copy.height.intersect(_range{min: guide.Frame.Size.Height, max: guide.Frame.Size.Height})
		copy = copy.solve()
		if !copy.isValid() {
			solver = copy
		}
	}

	return Guide{Frame: solver.rect(), Insets: in}
}

type _range struct {
	min float64
	max float64
}

func (r _range) intersect(r2 _range) _range {
	return _range{min: math.Max(r.min, r2.min), max: math.Min(r.max, r2.max)}
}

func (r _range) nearest(v float64) float64 {
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

type constraintSolver struct {
	left, right, top, bottom, width, height, centerX, centerY _range
}

func newConstraintSolver() constraintSolver {
	all := _range{min: math.Inf(-1), max: math.Inf(1)}
	pos := _range{min: 0, max: math.Inf(1)}
	return constraintSolver{
		left: all, right: all, top: all, bottom: all, width: pos, height: pos, centerX: all, centerY: all,
	}
}

func (r constraintSolver) isValid() bool {
	if r.left.max > r.left.min ||
		r.right.max > r.right.min ||
		r.top.max > r.top.min ||
		r.bottom.max > r.bottom.min ||
		r.width.max > r.width.min ||
		r.height.max > r.height.min ||
		r.centerX.max > r.centerX.min ||
		r.centerY.max > r.centerY.min ||
		r.width.max < 0 ||
		r.width.min < 0 ||
		r.height.max < 0 ||
		r.height.min < 0 {
		return false
	}
	return true
}

// Layout along the x-axis is determined by the `Left`, `Right`, `CenterX`, and `Width` attributes.
// And any x-axis attribute can be determined given two other x-axis attributes. We can therefore
// solve this constraint system by giving each attribute a unique priority and iteratively
// updating higher priority attributes from every possible combination of lower priority attributes.
// Our priorities from highest to lowest are: `Width`, `Left`, `Right`, `CenterX`.
// And for the y-axis are: `Height`, `Top`, `Bottom`, `CenterY`.
func (r constraintSolver) solve() constraintSolver {
	// Left = CenterX * 2 - Right
	left := _range{min: math.Inf(-1), max: math.Inf(1)}
	if !math.IsInf(r.centerX.min, 0) && !math.IsInf(r.right.max, 0) {
		left.min = r.centerX.min*2 - r.right.max
	}
	if !math.IsInf(r.centerX.max, 0) && !math.IsInf(r.right.min, 0) {
		left.max = r.centerX.max*2 - r.right.min
	}
	r.left = r.left.intersect(left)

	// Width = Right - Left
	width := _range{min: 0, max: math.Inf(1)}
	if !math.IsInf(r.right.max, 0) && !math.IsInf(r.left.min, 0) {
		width.max = r.right.max - r.left.min
	}
	if !math.IsInf(r.right.min, 0) && !math.IsInf(r.left.max, 0) {
		width.min = r.right.min - r.left.max
	}
	r.width = r.width.intersect(width)

	// Width = (CenterX - Left) * 2
	width = _range{min: 0, max: math.Inf(1)}
	if !math.IsInf(r.centerX.max, 0) && !math.IsInf(r.left.min, 0) {
		width.max = (r.centerX.max - r.left.min) * 2
	}
	if !math.IsInf(r.centerX.min, 0) && !math.IsInf(r.left.max, 0) {
		width.min = (r.centerX.min - r.left.max) * 2
	}
	r.width = r.width.intersect(width)

	// Top = CenterY * 2 - Bottom
	top := _range{min: 0, max: math.Inf(1)}
	if !math.IsInf(r.centerY.min, 0) && !math.IsInf(r.bottom.max, 0) {
		left.min = r.centerY.min*2 - r.bottom.max
	}
	if !math.IsInf(r.centerY.max, 0) && !math.IsInf(r.bottom.min, 0) {
		left.max = r.centerY.max*2 - r.bottom.min
	}
	r.top = r.top.intersect(top)

	// Height = Bottom - Top
	height := _range{min: 0, max: math.Inf(1)}
	if !math.IsInf(r.bottom.max, 0) && !math.IsInf(r.top.min, 0) {
		height.max = r.bottom.max - r.top.min
	}
	if !math.IsInf(r.bottom.min, 0) && !math.IsInf(r.top.max, 0) {
		height.min = r.bottom.min - r.top.max
	}
	r.height = r.height.intersect(height)

	// Height = (CenterY - Top) * 2
	height = _range{min: 0, max: math.Inf(1)}
	if !math.IsInf(r.centerY.max, 0) && !math.IsInf(r.top.min, 0) {
		height.max = (r.centerY.max - r.top.min) * 2
	}
	if !math.IsInf(r.centerY.min, 0) && !math.IsInf(r.top.max, 0) {
		height.min = (r.centerY.min - r.top.max) * 2
	}
	r.height = r.height.intersect(height)

	return r
}

// Assumes `constraintSolver` is valid. Returns the smallest possible size, and the origin closest to (0, 0).
func (r constraintSolver) rect() Rect {
	return Rt(r.left.nearest(0), r.top.nearest(0), r.width.nearest(0), r.height.nearest(0))
}

// Left
// Right
// Top
// Bottom
// Width
// Height
// CenterX
// CenterY

// InLft
// InRgt
// InTop
// InBot
// InWid
// InHei
// InCnx
// InCny

// Gr
// Eq
// Ls

// Lft
// Rgt
// Top
// Bot
// Wid
// Hei
// Cnx
// Cny

// LftIn
// RgtIn
// TopIn
// BotIn
// WidIn
// HeiIn
// CnxIn
// CnyIn

// {lay.LftIn, lay.Eq, lay.RgtIn, "", 0}
