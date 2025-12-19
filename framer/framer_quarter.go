package framer

import "github.com/xh3b4sd/choreo/framer/ticker"

func (f *Framer) Quarter(qnt int) {
	f.tic = ticker.Quarter(f.tic).Tick(qnt)
}
