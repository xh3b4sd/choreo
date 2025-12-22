package ticker

import "time"

type day struct {
	tic time.Time
}

func Day(tic time.Time) Interface {
	return &day{
		tic: tic,
	}
}

func (d *day) Round(mul int) Interface {
	{
		d.tic = rndDay(d.tic, mul)
	}

	return d
}

func (d *day) Tick(qnt int) Interface {
	{
		d.tic = d.tic.AddDate(0, 0, qnt)
	}

	return d
}

func (d *day) Time() time.Time {
	return d.tic
}

func rndDay(tic time.Time, mul int) time.Time {
	var yea, mon, _ = tic.Date()

	var zer = time.Date(yea, mon, 1, 0, 0, 0, 0, tic.Location())

	return zer.Add(tic.Sub(zer).Round(time.Duration(mul) * 24 * time.Hour))
}
