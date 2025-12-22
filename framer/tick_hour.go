package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Hour(qnt int) time.Time {
	return ticker.Hour(t.fra.tic).Tick(qnt).Time()
}
