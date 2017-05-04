package internal

import (
	"fmt"
	"mochi/bridge"
	"sync"
	"time"
)

func init() {
	bridge.RegisterFunc("github.com/overcyn/mochi/animate screenUpdate", screenUpdate)
}

var tickers = struct {
	ts     map[int]*Ticker
	mu     *sync.Mutex
	maxKey int
}{
	ts:     map[int]*Ticker{},
	mu:     &sync.Mutex{},
	maxKey: 0,
}

func screenUpdate() {
	fmt.Println("start")
	bridge.Root().Call("goWantsUpdate")
	tickers.mu.Lock()
	defer tickers.mu.Unlock()

	for _, i := range tickers.ts {
		i.send()
	}
	fmt.Println("done")
}

type Ticker struct {
	key      int
	mu       *sync.Mutex
	chans    []chan<- struct{}
	timer    *time.Timer
	start    time.Time
	duration time.Duration
}

func NewTicker(duration time.Duration) *Ticker {
	tickers.mu.Lock()
	defer tickers.mu.Unlock()

	mu := &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	tickers.maxKey += 1
	t := &Ticker{
		key:      tickers.maxKey,
		mu:       mu,
		start:    time.Now(),
		duration: duration,
	}
	t.timer = time.AfterFunc(duration, func() {
		t.Stop()
	})
	tickers.ts[t.key] = t
	return t
}

func (t *Ticker) Notify(c chan<- struct{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.chans = append(t.chans, c)
}

func (t *Ticker) Unnotify(c chan<- struct{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	chans := []chan<- struct{}{}
	for _, i := range t.chans {
		if c != i {
			chans = append(chans, c)
		} else {
			break
		}
	}
	t.chans = chans
}

func (t *Ticker) Value() float64 {
	v := float64(time.Since(t.start) / t.duration)
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	return v
}

func (t *Ticker) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	tickers.mu.Lock()
	defer tickers.mu.Unlock()

	delete(tickers.ts, t.key)
}

func (t *Ticker) send() {
	t.mu.Lock()
	chans := t.chans
	t.mu.Unlock()

	for _, i := range chans {
		select {
		case i <- struct{}{}:
		default:
		}
	}
}
