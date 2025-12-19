package framer

import (
	"time"
)

func (t Tick) Time() time.Time {
	return t.fra.tic
}
