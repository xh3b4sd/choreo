package ticker

import "time"

type second struct {
	tic time.Time
}

func Second(tic time.Time) Interface {
	return &second{
		tic: tic,
	}
}

func (s *second) Round(mul int) Interface {
	{
		s.tic = rndSec(s.tic, mul)
	}

	return s
}

func (s *second) Tick(qnt int) Interface {
	{
		s.tic = s.tic.Add(time.Duration(qnt) * time.Second)
	}

	return s
}

func (s *second) Time() time.Time {
	return s.tic
}

func rndSec(tic time.Time, mul int) time.Time {
	var yea, mon, day = tic.Date()

	var zer = time.Date(yea, mon, day, 0, 0, 0, 0, tic.Location())

	return zer.Add(tic.Sub(zer).Round(time.Duration(mul) * time.Second))
}
