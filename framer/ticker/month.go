package ticker

import "time"

type month struct {
	tic time.Time
}

func Month(tic time.Time) Interface {
	return &month{
		tic: tic,
	}
}

func (m *month) Round(mul int) Interface {
	{
		m.tic = rndMon(m.tic, mul)
	}

	return m
}

func (m *month) Tick(qnt int) Interface {
	{
		m.tic = m.tic.AddDate(0, qnt, 0)
	}

	return m
}

func (m *month) Time() time.Time {
	return m.tic
}

func rndMon(tic time.Time, mul int) time.Time {
	var y, m, d = tic.Date()

	if d >= 15 {
		y, m, _ = time.Date(y, m, 1, 0, 0, 0, 0, tic.Location()).AddDate(0, 1, 0).Date()
	}

	var x = (((y*12 + (int(m) - 1)) + mul/2) / mul) * mul

	return time.Date(x/12, time.Month(x%12+1), 1, 0, 0, 0, 0, tic.Location())
}
