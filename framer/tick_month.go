package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Month(qnt int) time.Time {
	return ticker.Month(t.fra.tic).Tick(qnt)
}
