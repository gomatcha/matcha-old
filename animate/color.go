package animate

import (
	"github.com/overcyn/mochi"
	"image/color"
)

type ColorInterpolater interface {
	Interpolate(float64) color.Color
}

func ColorInterpolate(w mochi.Float64Notifier, l ColorInterpolater) mochi.ColorNotifier {
	return &colorInterpolater{
		watcher:      w,
		interpolater: l,
	}
}

type colorInterpolater struct {
	watcher      mochi.Float64Notifier
	interpolater ColorInterpolater
}

func (w *colorInterpolater) Notify() chan struct{} {
	return w.watcher.Notify()
}

func (w *colorInterpolater) Unnotify(c chan struct{}) {
	w.watcher.Unnotify(c)
}

func (w *colorInterpolater) Value() color.Color {
	return w.interpolater.Interpolate(w.watcher.Value())
}

type RGBALerp struct {
	Start, End color.Color
}

func (e *RGBALerp) Interpolate(a float64) color.Color {
	r1, g1, b1, a1 := e.Start.RGBA()
	r2, g2, b2, a2 := e.End.RGBA()
	return color.RGBA64{
		R: uintInterpolate(r1, r2, a),
		G: uintInterpolate(g1, g2, a),
		B: uintInterpolate(b1, b2, a),
		A: uintInterpolate(a1, a2, a),
	}
}

func uintInterpolate(a, b uint32, c float64) uint16 {
	return uint16(a + uint32(float64(b-a)*c))
}
