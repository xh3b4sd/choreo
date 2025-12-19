package timeout

import "time"

type Config struct {
	Clo chan struct{}
	Dur time.Duration
}

type Limit struct {
	clo chan struct{}
	dur time.Duration
}

func New(c Config) *Limit {
	return &Limit{
		clo: c.Clo,
		dur: c.Dur,
	}
}
