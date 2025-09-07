package jitter

// Number allows every numerical type, as well as any implementation using those
// underlying types to be jittered, e.g. time.Duration.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}
