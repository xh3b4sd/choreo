package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Day(qnt int) time.Time {
	return ticker.Day(t.fra.tic).Tick(qnt).Time()
}
