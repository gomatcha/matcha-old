package paint

import (
	"github.com/overcyn/mochi"
	"image/color"
)

type Painter interface {
	PaintStyle() Style
	mochi.Notifier
}

type Style struct {
	Alpha           float64
	BackgroundColor color.Color
	BorderColor     color.Color
	BorderWidth     float64
	CornerRadius    float64
	ShadowOpacity   float64
	ShadowRadius    float64
	ShadowOffset    float64
	ShadowColor     color.Color
	// Transform?
	// Mask
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
	Alpha           mochi.Float64Notifier
	BackgroundColor mochi.ColorNotifier
	BorderColor     mochi.ColorNotifier
	BorderWidth     mochi.Float64Notifier
	CornerRadius    mochi.Float64Notifier
	ShadowOpacity   mochi.Float64Notifier
	ShadowRadius    mochi.Float64Notifier
	ShadowOffset    mochi.Float64Notifier
	ShadowColor     mochi.ColorNotifier

	batchNotifiers map[chan struct{}]*mochi.BatchNotifier
}

func (as *AnimatedStyle) PaintStyle() Style {
	s := as.Style
	if as.Alpha != nil {
		s.Alpha = as.Alpha.Value()
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
	if as.ShadowOpacity != nil {
		s.ShadowOpacity = as.ShadowOpacity.Value()
	}
	if as.ShadowRadius != nil {
		s.ShadowRadius = as.ShadowRadius.Value()
	}
	if as.ShadowOffset != nil {
		s.ShadowOffset = as.ShadowOffset.Value()
	}
	if as.ShadowColor != nil {
		s.ShadowColor = as.ShadowColor.Value()
	}
	return s
}

func (as *AnimatedStyle) Notify() chan struct{} {
	return nil // no-op
}

func (as *AnimatedStyle) Unnotify(chan struct{}) {
	// no-op
}
