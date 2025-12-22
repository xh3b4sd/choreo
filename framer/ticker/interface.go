package ticker

import "time"

type Interface interface {
	// Round adjusts the underling time instance to the nearest multiple provided.
	Round(int) Interface

	// Tick adjusts the underling time instance by the provided quantity of units
	// defined by the underlying implementation. In other words, the minute
	// implementation moves the underlying time by the provided amount of minutes.
	Tick(int) Interface

	// Time returns the underlying time instance.
	Time() time.Time
}
