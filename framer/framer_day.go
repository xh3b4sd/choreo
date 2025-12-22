package framer

import "github.com/xh3b4sd/choreo/framer/ticker"

func (f *Framer) Day(qnt int) {
	f.tic = ticker.Day(f.tic).Tick(qnt).Time()
}
