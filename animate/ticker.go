package animate

import (
	"time"

	"github.com/overcyn/mochi/internal"
)

type Ticker struct {
	ticker *internal.Ticker
}

func NewTicker(duration time.Duration) *Ticker {
	return &Ticker{
		ticker: internal.NewTicker(duration),
	}
}

func (t *Ticker) Notify() chan struct{} {
	return t.ticker.Notify()
}

func (t *Ticker) Unnotify(c chan struct{}) {
	t.ticker.Unnotify(c)
}

func (t *Ticker) Value() float64 {
	return t.ticker.Value()
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
}
