package backoff

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Bac is the list of token indexed backoff durations that apply on execution
	// failure. The backoff duration is selected by token index, starting at index
	// 0. Every execution failure increments the token index, which increases the
	// applied backoff, given an ascending list of durations. Every execution
	// success decrements the token index which decreases the applied backoff
	// again. The default list in seconds looks like the following.
	Bac []time.Duration
	// Inf is the infinity option to specify whether the last token index ends the
	// backoff. By default, once the last token index was applied on failure, any
	// subsequent failure will not cause another retry, but instead causes
	// Token.Backoff to return that last observed error.
	Inf bool
}

type Token struct {
	bac []time.Duration
	ind int
	inf bool
	slp func(time.Duration)
}

func New(c Config) *Token {
	if len(c.Bac) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Bac must not be empty", c)))
	}

	return &Token{
		bac: c.Bac,
		ind: 0,
		inf: c.Inf,
		slp: time.Sleep,
	}
}
