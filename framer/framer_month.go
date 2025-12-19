package framer

import "github.com/xh3b4sd/choreo/framer/ticker"

func (f *Framer) Month(qnt int) {
	f.tic = ticker.Month(f.tic).Tick(qnt)
}
