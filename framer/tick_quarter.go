package framer

import (
	"time"

	"github.com/xh3b4sd/choreo/framer/ticker"
)

func (t Tick) Quarter(qnt int) time.Time {
	return ticker.Quarter(t.fra.tic).Tick(qnt)
}
