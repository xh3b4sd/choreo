package ticker

import "time"

type minute struct {
	tic time.Time
}

func Minute(tic time.Time) Interface {
	return &minute{
		tic: tic,
	}
}

func (m *minute) Round(mul int) Interface {
	{
		m.tic = rndMin(m.tic, mul)
	}

	return m
}

func (m *minute) Tick(qnt int) Interface {
	{
		m.tic = m.tic.Add(time.Duration(qnt) * time.Minute)
	}

	return m
}

func (m *minute) Time() time.Time {
	return m.tic
}

func rndMin(tic time.Time, mul int) time.Time {
	var yea, mon, day = tic.Date()

	var zer = time.Date(yea, mon, day, 0, 0, 0, 0, tic.Location())

	return zer.Add(tic.Sub(zer).Round(time.Duration(mul) * time.Minute))
}
