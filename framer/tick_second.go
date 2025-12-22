package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Second(qnt int) time.Time {
	return ticker.Second(t.fra.tic).Tick(qnt).Time()
}
