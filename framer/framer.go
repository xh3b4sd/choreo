package framer

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Min time.Time
	Max time.Time
}

type Framer struct {
	min time.Time
	max time.Time

	tic time.Time
}

func New(c Config) *Framer {
	if c.Min.IsZero() {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Min must not be empty", c)))
	}
	if c.Max.IsZero() {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Max must not be empty", c)))
	}
	if !c.Max.After(c.Min) {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Max must be after %T.Min", c, c)))
	}

	return &Framer{
		min: c.Min,
		max: c.Max,

		tic: c.Min,
	}
}
