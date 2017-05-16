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

func (p *Style) PaintStyle() Style {
	return *p
}

func (p *Style) Notify() chan struct{} {
	return nil // no-op
}

func (p *Style) Unnotify(chan struct{}) {
	// no-op
}

type AnimatedStyle struct {
	Style Style

	// Alpha           Float64Notifier
	// BackgroundColor ColorNotifier
	// BorderColor     ColorNotifier
	// BorderWidth     Float64Notifier
	// CornerRadius    Float64Notifier
	// ShadowOpacity   Float64Notifier
	// ShadowRadius    Float64Notifier
	// ShadowOffset    Float64Notifier
	// ShadowColor     ColorNotifier
}

func (p *AnimatedStyle) PaintStyle() Style {
	return p.Style
}

func (p *AnimatedStyle) Notify() chan struct{} {
	return nil // no-op
}

func (p *AnimatedStyle) Unnotify(chan struct{}) {
	// no-op
}
