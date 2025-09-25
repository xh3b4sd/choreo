package ticker

import "time"

type Interface interface {
	// Close turns off a ticker by calling Ticker.Stop on the underlying ticker
	// implementation. Once Ticker.Close was called, no more ticks will be
	// delivered.
	Close()

	// Reset stops a ticker and resets its period to the underlying wait duration.
	// The next tick will arrive at Ticker.Ticks after this wait duration elapses
	// again.
	Reset()

	// Ticks returns the underlying Ticker.C channel on which the ticks are
	// delivered.
	Ticks() <-chan time.Time
}
