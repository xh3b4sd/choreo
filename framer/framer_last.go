package framer

func (f *Framer) Last() bool {
	return !f.tic.Before(f.max)
}
