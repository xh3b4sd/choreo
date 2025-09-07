package jitter

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Per is the fraction of jitter added or removed at random. E.g. 0.25 will
	// modify the given input by a random percentage between -25% and +25%.
	// Values <= 0 and >= 1 will cause a runtime panic.
	Per float64
}

type Jitter[T Number] struct {
	per float64
	ran *rand.Rand
}

func New[T Number](c Config) *Jitter[T] {
	if c.Per <= 0 || c.Per >= 1 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Per must be within (0, 1)", c)))
	}

	return &Jitter[T]{
		per: c.Per,
		ran: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
