package framer

func (f *Framer) Tick() Tick {
	return Tick{
		fra: f,
	}
}
