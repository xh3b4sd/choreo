package backoff

import "time"

// Backoff returns the default token indexed backoff durations as a drop-in
// option.
//
//	[ 1, 1, 2, 4, 8, 16, 32, 64 ]
func Backoff() []time.Duration {
	return []time.Duration{
		1 * time.Second,
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		16 * time.Second,
		32 * time.Second,
		64 * time.Second,
	}
}
