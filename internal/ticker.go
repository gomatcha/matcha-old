package internal

import (
	"github.com/overcyn/mochibridge"
	"sync"
	"time"
)

var tickers = struct {
	ts     map[int]*Ticker
	mu     *sync.Mutex
	maxKey int
}{
	ts:     map[int]*Ticker{},
	mu:     &sync.Mutex{},
	maxKey: 0,
}

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/animate screenUpdate", screenUpdate)
}

func screenUpdate() {
	tickers.mu.Lock()
	defer tickers.mu.Unlock()

	for _, i := range tickers.ts {
		i.send()
	}
}

type Ticker struct {
	key      int
	mu       sync.Mutex
	chans    map[chan struct{}]struct{}
	funcs    []func()
	timer    *time.Timer
	start    time.Time
	duration time.Duration
}

func NewTicker(duration time.Duration) *Ticker {
	tickers.mu.Lock()
	defer tickers.mu.Unlock()

	tickers.maxKey += 1
	t := &Ticker{
		key:      tickers.maxKey,
		chans:    map[chan struct{}]struct{}{},
		start:    time.Now(),
		duration: duration,
	}
	t.timer = time.AfterFunc(duration, func() {
		t.Stop()
	})
	tickers.ts[t.key] = t
	return t
}

func (t *Ticker) Notify() chan struct{} {
	t.mu.Lock()
	defer t.mu.Unlock()

	c := make(chan struct{})
	t.chans[c] = struct{}{}
	return c
}

func (t *Ticker) Unnotify(c chan struct{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, ok := t.chans[c]
	if !ok {
		panic("Unnotify called with unknown chan")
	}
	delete(t.chans, c)
}

func (t *Ticker) NotifyFunc(f func()) int {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.funcs = append(t.funcs, f)
	return 0
}

func (t *Ticker) UnnotifyFunc(key int) {
	// TODO(KD):
}

func (t *Ticker) Value() float64 {
	v := float64(time.Since(t.start)) / float64(t.duration)
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	return v
}

func (t *Ticker) Stop() {
	tickers.mu.Lock()
	defer tickers.mu.Unlock()
	t.mu.Lock()
	defer t.mu.Unlock()

	delete(tickers.ts, t.key)
}

func (t *Ticker) send() {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i := range t.chans {
		i <- struct{}{}
		<-i
	}
	for _, i := range t.funcs {
		i()
	}
}
