package ticker

import (
	"time"
)

type Fake struct{}

func (f Fake) Close() {}

func (f Fake) Reset() {}

func (f Fake) Ticks() <-chan time.Time {
	return nil
}
