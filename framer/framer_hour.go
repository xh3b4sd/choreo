package framer

import "github.com/xh3b4sd/choreo/framer/ticker"

func (f *Framer) Hour(qnt int) {
	f.tic = ticker.Hour(f.tic).Tick(qnt).Time()
}
