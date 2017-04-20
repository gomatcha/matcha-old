package bridge

import (
	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/constraint"
	"github.com/overcyn/mochi/text"
	"mochi/bridge"
)

type GoRoot struct {
}

func (b *GoRoot) Display() *mochi.Node {
	v := &NestedView{}
	n := mochi.Display(v)
	return n
}

func init() {
	bridge.SetGoRoot(&GoRoot{})
}

const (
	chl1id int = iota
	chl2id
	chl3id
	chl4id
	chl5id
)

type NestedView struct {
	marker mochi.Marker
}

func (v *NestedView) Mount(m mochi.Marker) {
	v.marker = m
}

func (v *NestedView) Update(p *mochi.Node) *mochi.Node {
	l := constraint.New()
	n := mochi.NewNode()
	n.Layouter = l
	n.PaintOptions.BackgroundColor = mochi.GreenColor

	chl1 := mochi.NewBasicView(p.Get(chl1id))
	chl1.PaintOptions.BackgroundColor = mochi.RedColor
	n.Set(chl1id, chl1)
	g1 := l.AddGuide(chl1id, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(0))
		s.LeftEqual(constraint.Const(0))
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl2 := mochi.NewBasicView(p.Get(chl2id))
	chl2.PaintOptions.BackgroundColor = mochi.YellowColor
	n.Set(chl2id, chl2)
	g2 := l.AddGuide(chl2id, func(s *constraint.Solver) {
		s.TopEqual(g1.Bottom())
		s.LeftEqual(g1.Left())
		s.WidthEqual(constraint.Const(300))
		s.HeightEqual(constraint.Const(300))
	})

	chl3 := mochi.NewBasicView(p.Get(chl3id))
	chl3.PaintOptions.BackgroundColor = mochi.BlueColor
	n.Set(chl3id, chl3)
	g3 := l.AddGuide(chl3id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g2.Left())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(100))
	})

	chl4 := mochi.NewBasicView(p.Get(chl4id))
	chl4.PaintOptions.BackgroundColor = mochi.MagentaColor
	n.Set(chl4id, chl4)
	_ = l.AddGuide(chl4id, func(s *constraint.Solver) {
		s.TopEqual(g2.Bottom())
		s.LeftEqual(g3.Right())
		s.WidthEqual(constraint.Const(50))
		s.HeightEqual(constraint.Const(50))
	})

	chl5 := text.NewTextView(p.Get(chl5id))
	chl5.PaintOptions.BackgroundColor = mochi.CyanColor
	chl5.Text = "poop"
	chl5.Format.SetAlignment(text.AlignmentCenter)
	chl5.Format.SetStrikethroughStyle(text.StrikethroughStyleSingle)
	chl5.Format.SetStrikethroughColor(mochi.MagentaColor)
	chl5.Format.SetUnderlineStyle(text.UnderlineStyleDouble)
	chl5.Format.SetUnderlineColor(mochi.GreenColor)
	chl5.Format.SetFont(text.Font{
		Family: "American Typewriter",
		Face:   "Bold",
		Size:   20,
	})
	n.Set(chl5id, chl5)
	_ = l.AddGuide(chl5id, func(s *constraint.Solver) {
		s.BottomEqual(g2.Bottom())
		s.RightEqual(g2.Right())
		s.WidthEqual(constraint.Const(100))
		s.HeightEqual(constraint.Const(50))
	})

	return n
}

func (v *NestedView) Unmount() {
	v.marker = nil
}
