package mochi

import "image/color"

type Painter interface {
	PaintOptions() PaintOptions
	Notifier
}

type PaintOptions struct {
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

type AnimatedPaintStyle struct {
	Alpha           Float64Notifier
	BackgroundColor ColorNotifier
	BorderColor     ColorNotifier
	BorderWidth     Float64Notifier
	CornerRadius    Float64Notifier
	ShadowOpacity   Float64Notifier
	ShadowRadius    Float64Notifier
	ShadowOffset    Float64Notifier
	ShadowColor     ColorNotifier
}

func (p PaintOptions) PaintOptions() PaintOptions {
	return p
}

func (p PaintOptions) Notify() chan struct{} {
	return nil // no-op
}

func (p PaintOptions) Unnotify(chan struct{}) {
	// no-op
}
