package ticker

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Dur time.Duration
}

type Ticker struct {
	dur time.Duration
	tic *time.Ticker
}

func New(c Config) *Ticker {
	if c.Dur <= 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Bac must not be empty", c)))
	}

	return &Ticker{
		dur: c.Dur,
		tic: time.NewTicker(c.Dur),
	}
}

func (t *Ticker) Close() {
	t.tic.Stop()
}

func (t *Ticker) Reset() {
	t.tic.Reset(t.dur)
}

func (t *Ticker) Ticks() <-chan time.Time {
	return t.tic.C
}
