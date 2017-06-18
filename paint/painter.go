package paint

import (
	"image"
	"image/color"

	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/pb"
	"github.com/overcyn/matcha/pb/paint"
)

type Painter interface {
	Paint(*image.RGBA) // does nothing atm
	PaintStyle() Style
	comm.Notifier
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

func (s *Style) MarshalProtobuf() *paint.Style {
	return &paint.Style{
		Transparency:    s.Transparency,
		BackgroundColor: pb.ColorEncode(s.BackgroundColor),
		BorderColor:     pb.ColorEncode(s.BorderColor),
		BorderWidth:     s.BorderWidth,
		CornerRadius:    s.CornerRadius,
		ShadowRadius:    s.ShadowRadius,
		ShadowOffset:    s.ShadowOffset.MarshalProtobuf(),
		ShadowColor:     pb.ColorEncode(s.ShadowColor),
	}
}

func (s *Style) Paint(img *image.RGBA) {
}

func (s *Style) PaintStyle() Style {
	return *s
}

func (s *Style) Notify(func()) comm.Id {
	return 0 // no-op
}

func (s *Style) Unnotify(id comm.Id) {
	// no-op
}

type notifier struct {
	notifier *comm.BatchNotifier
	id       comm.Id
}

type AnimatedStyle struct {
	Style           Style
	Transparency    comm.Float64Notifier
	BackgroundColor comm.ColorNotifier
	BorderColor     comm.ColorNotifier
	BorderWidth     comm.Float64Notifier
	CornerRadius    comm.Float64Notifier
	ShadowRadius    comm.Float64Notifier
	// ShadowOffset    comm.Float64Notifier
	ShadowColor comm.ColorNotifier

	maxId          comm.Id
	batchNotifiers map[comm.Id]notifier
}

func (as *AnimatedStyle) Paint(img *image.RGBA) {
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

func (as *AnimatedStyle) Notify(f func()) comm.Id {
	n := &comm.BatchNotifier{}

	if as.Transparency != nil {
		n.Subscribe(as.Transparency)
	}
	if as.BackgroundColor != nil {
		n.Subscribe(as.BackgroundColor)
	}
	if as.BorderColor != nil {
		n.Subscribe(as.BorderColor)
	}
	if as.BorderWidth != nil {
		n.Subscribe(as.BorderWidth)
	}
	if as.CornerRadius != nil {
		n.Subscribe(as.CornerRadius)
	}
	if as.ShadowRadius != nil {
		n.Subscribe(as.ShadowRadius)
	}
	// if as.ShadowOffset != nil {
	// 	n.Subscribe(as.ShadowOffset)
	// }
	if as.ShadowColor != nil {
		n.Subscribe(as.ShadowColor)
	}

	as.maxId += 1
	if as.batchNotifiers == nil {
		as.batchNotifiers = map[comm.Id]notifier{}
	}
	as.batchNotifiers[as.maxId] = notifier{
		notifier: n,
		id:       n.Notify(f),
	}
	return as.maxId
}

func (as *AnimatedStyle) Unnotify(id comm.Id) {
	n, ok := as.batchNotifiers[id]
	if ok {
		n.notifier.Unnotify(n.id)
		delete(as.batchNotifiers, id)
	}
}
