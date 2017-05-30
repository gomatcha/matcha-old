package paint

import (
	"image/color"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
)

type Painter interface {
	PaintStyle() Style
	mochi.Notifier
}

type Style struct {
	Transparency    float64
	BackgroundColor color.Color
	BorderColor     color.Color
	BorderWidth     float64
	CornerRadius    float64
	ShadowRadius    float64
	ShadowOffset    layout.Point
	ShadowColor     color.Color
	// Transform?
	// Mask
}

func (s *Style) EncodeProtobuf() *pb.PaintStyle {
	return &pb.PaintStyle{
		Transparency:    s.Transparency,
		BackgroundColor: pb.ColorEncode(s.BackgroundColor),
		BorderColor:     pb.ColorEncode(s.BorderColor),
		BorderWidth:     s.BorderWidth,
		CornerRadius:    s.CornerRadius,
		ShadowRadius:    s.ShadowRadius,
		ShadowOffset:    s.ShadowOffset.EncodeProtobuf(),
		ShadowColor:     pb.ColorEncode(s.ShadowColor),
	}
}

func (s *Style) PaintStyle() Style {
	return *s
}

func (s *Style) Notify() chan struct{} {
	return nil // no-op
}

func (s *Style) Unnotify(chan struct{}) {
	// no-op
}

type AnimatedStyle struct {
	Style           Style
	Transparency    mochi.Float64Notifier
	BackgroundColor mochi.ColorNotifier
	BorderColor     mochi.ColorNotifier
	BorderWidth     mochi.Float64Notifier
	CornerRadius    mochi.Float64Notifier
	ShadowRadius    mochi.Float64Notifier
	// ShadowOffset    mochi.Float64Notifier
	ShadowColor mochi.ColorNotifier

	batchNotifiers map[chan struct{}]*mochi.BatchNotifier
}

func (as *AnimatedStyle) PaintStyle() Style {
	s := as.Style
	if as.Transparency != nil {
		s.Transparency = as.Transparency.Value()
	}
	if as.BackgroundColor != nil {
		s.BackgroundColor = as.BackgroundColor.Value()
	}
	if as.BorderColor != nil {
		s.BorderColor = as.BorderColor.Value()
	}
	if as.BorderWidth != nil {
		s.BorderWidth = as.BorderWidth.Value()
	}
	if as.CornerRadius != nil {
		s.CornerRadius = as.CornerRadius.Value()
	}
	if as.ShadowRadius != nil {
		s.ShadowRadius = as.ShadowRadius.Value()
	}
	// if as.ShadowOffset != nil {
	// 	s.ShadowOffset = as.ShadowOffset.Value()
	// }
	if as.ShadowColor != nil {
		s.ShadowColor = as.ShadowColor.Value()
	}
	// fmt.Println(s.BackgroundColor)
	return s
}

func (as *AnimatedStyle) Notify() chan struct{} {
	if as.batchNotifiers == nil {
		as.batchNotifiers = map[chan struct{}]*mochi.BatchNotifier{}
	}

	ns := []mochi.Notifier{}
	if as.Transparency != nil {
		ns = append(ns, as.Transparency)
	}
	if as.BackgroundColor != nil {
		ns = append(ns, as.BackgroundColor)
	}
	if as.BorderColor != nil {
		ns = append(ns, as.BorderColor)
	}
	if as.BorderWidth != nil {
		ns = append(ns, as.BorderWidth)
	}
	if as.CornerRadius != nil {
		ns = append(ns, as.CornerRadius)
	}
	if as.ShadowRadius != nil {
		ns = append(ns, as.ShadowRadius)
	}
	// if as.ShadowOffset != nil {
	// 	ns = append(ns, as.ShadowOffset)
	// }
	if as.ShadowColor != nil {
		ns = append(ns, as.ShadowColor)
	}

	n := mochi.NewBatchNotifier(ns...)
	c := n.Notify()
	if c != nil {
		as.batchNotifiers[c] = n
	}
	return c
}

func (as *AnimatedStyle) Unnotify(c chan struct{}) {
	if c == nil {
		return
	}
	n := as.batchNotifiers[c]
	if n != nil {
		n.Unnotify(c)
		delete(as.batchNotifiers, c)
	}
}
