package animate

import (
	"time"

	"github.com/overcyn/mochi/comm"
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

func (t *Ticker) Notify(f func()) comm.Id {
	return comm.Id(t.ticker.NotifyFunc(f))
}

func (t *Ticker) Unnotify(id comm.Id) {
	t.ticker.UnnotifyFunc(int(id))
}

func (t *Ticker) Value() float64 {
	return t.ticker.Value()
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
}
