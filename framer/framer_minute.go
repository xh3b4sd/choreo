package framer

import "github.com/xh3b4sd/choreo/framer/ticker"

func (f *Framer) Minute(qnt int) {
	f.tic = ticker.Minute(f.tic).Tick(qnt).Time()
}
