package jitter

func (j *Jitter[T]) Percent(num T) T {
	return T(float64(num) * (1 + j.per*(2*j.ran.Float64()-1)))
}
