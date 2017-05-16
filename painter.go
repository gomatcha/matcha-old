package mochi

import "image/color"

type Painter interface {
	PaintStyle() PaintStyle
	Notifier
}

type PaintStyle struct {
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

func (p *PaintStyle) PaintStyle() PaintStyle {
	return *p
}

func (p *PaintStyle) Notify() chan struct{} {
	return nil // no-op
}

func (p *PaintStyle) Unnotify(chan struct{}) {
	// no-op
}

type AnimatedPaintStyle struct {
	Style PaintStyle

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

func (p *AnimatedPaintStyle) PaintStyle() PaintStyle {
	return p.Style
}

func (p *AnimatedPaintStyle) Notify() chan struct{} {
	return nil // no-op
}

func (p *AnimatedPaintStyle) Unnotify(chan struct{}) {
	// no-op
}
