package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Minute(qnt int) time.Time {
	return ticker.Minute(t.fra.tic).Tick(qnt).Time()
}
